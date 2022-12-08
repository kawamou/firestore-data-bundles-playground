package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
)

const UploadedBundleURL = "Firestore Data Bundlesを格納したGCS等のURL"

var (
	headNumber   = regexp.MustCompile("^[0-9]{3}{")
	middleNumber = regexp.MustCompile("}[0-9]{3}{")
)

func init() {}

func main() {
	c := NewBundlesClient()
	rets, _ := c.Get()
	for _, ret := range rets {
		log.Println(ret)
	}
}

type BundlesClient struct {
	c *http.Client
}

func (b *BundlesClient) Get() ([]any, error) {
	resp, err := b.c.Get(UploadedBundleURL)
	if err != nil {
		return nil, err
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bundlesToJSON(string(bytes))
}

func NewBundlesClient() *BundlesClient {
	return &BundlesClient{
		c: http.DefaultClient,
	}
}

// replaceHeadNumber はFirestore Data Bundlesに含まれる先頭の数字を置換します
func replaceHeadNumber(str string) string {
	return headNumber.ReplaceAllString(str, "[{")
}

// replaceMiddleNumber はFirestore Data Bundlesに含まれる途中の数字を置換します
func replaceMiddleNumber(str string) string {
	return middleNumber.ReplaceAllString(str, "},{")
}

// bundlesToJSONString はFirestore Data Bundlesに含まれる数字を置換しJSON文字列に変換します
// Firestore Data Bundlesはトップレベルに不要な数字を含んだJSON然としているため、JSONとして取り扱えるようにクレンジングします
// ex. "142{\"metadata\": ... }262{\"document\": ... }}}",
// ref. Firestore data bundles から余分な数字を取り除いて整形されたjsonでファイルに出力する（https://ta-watanabe.hatenablog.com/entry/2021/07/15/165732）
func bundlesToJSONString(str string) string {
	str = replaceHeadNumber(str)
	return replaceMiddleNumber(str) + "]"
}

// bundlesToJSON はFirestore Data BundlesをJSONに変換後、Goで扱うための型に変換します
func bundlesToJSON(str string) ([]any, error) {
	var ret []any
	formatted := bundlesToJSONString(str)
	if err := json.Unmarshal([]byte(formatted), &ret); err != nil {
		return nil, err
	}
	return ret, nil
}
