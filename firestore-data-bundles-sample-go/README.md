## これは何
- Firestore Data BundlesをGoで取り扱うための試み
- Firestore Data Bundlesを読み込むためのAPIはクライアントのSDK（Web, モバイル）にしか準備されていない
- 結論、Firestore Data Bundles文字列を整形して、JSONとして取り扱うマッチョな方法しか難しそう
- Experimentalなコード

## 実行方法
```shell
go run main.go
```

## 参考
- [Firestore data bundles から余分な数字を取り除いて整形されたjsonでファイルに出力する](https://ta-watanabe.hatenablog.com/entry/2021/07/15/165732)