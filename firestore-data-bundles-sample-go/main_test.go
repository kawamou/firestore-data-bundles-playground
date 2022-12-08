package main

import (
	"testing"
)

func TestReplaceHeadNumber(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				target: "123{abc}256{abc}",
			},
			want: "[{abc}256{abc}",
		},
		{
			name: "",
			args: args{
				target: "1234{abc}256{abc}",
			},
			want: "1234{abc}256{abc}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceHeadNumber(tt.args.target); got != tt.want {
				t.Errorf("replaceHeadNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceMiddleNumber(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				target: "123{abc}256{abc}",
			},
			want: "123{abc},{abc}",
		},
		{
			name: "",
			args: args{
				target: "123{abc}2567{abc}",
			},
			want: "123{abc}2567{abc}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceMiddleNumber(tt.args.target); got != tt.want {
				t.Errorf("replaceMiddleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirestoreDataBundlesToJSONString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				str: "142{\"metadata\":{\"id\":\"1670488463147\",\"createTime\":{\"seconds\":\"1670488464\",\"nanos\":602042000},\"version\":1,\"totalDocuments\":2,\"totalBytes\":\"1202\"}}235{\"namedQuery\":{\"name\":\"latestBundles\",\"bundledQuery\":{\"parent\":\"projects/xxx/databases/(default)/documents\",\"structuredQuery\":{\"from\":[{\"collectionId\":\"bundles\"}]}},\"readTime\":{\"seconds\":\"1670488464\",\"nanos\":602042000}}}214{\"documentMetadata\":{\"name\":\"projects/xxx/databases/(default)/documents/bundles/xxx\",\"readTime\":{\"seconds\":\"1670488464\",\"nanos\":602042000},\"exists\":true,\"queries\":[\"latestBundles\"]}}262{\"document\":{\"name\":\"projects/xxx/databases/(default)/documents/bundles/xxx\",\"fields\":{\"data\":{\"stringValue\":\"bar\"}},\"createTime\":{\"seconds\":\"1670487756\",\"nanos\":768917000},\"updateTime\":{\"seconds\":\"1670487756\",\"nanos\":768917000}}}214{\"documentMetadata\":{\"name\":\"projects/xxx/databases/(default)/documents/bundles/xxx\",\"readTime\":{\"seconds\":\"1670488464\",\"nanos\":602042000},\"exists\":true,\"queries\":[\"latestBundles\"]}}262{\"document\":{\"name\":\"projects/xxx/databases/(default)/documents/bundles/xxx\",\"fields\":{\"data\":{\"stringValue\":\"foo\"}},\"createTime\":{\"seconds\":\"1670487756\",\"nanos\":773423000},\"updateTime\":{\"seconds\":\"1670487756\",\"nanos\":773423000}}}",
			},
			want: "[{\"metadata\":{\"id\":\"1670488463147\",\"createTime\":{\"seconds\":\"1670488464\",\"nanos\":602042000},\"version\":1,\"totalDocuments\":2,\"totalBytes\":\"1202\"}},{\"namedQuery\":{\"name\":\"latestBundles\",\"bundledQuery\":{\"parent\":\"projects/xxx/databases/(default)/documents\",\"structuredQuery\":{\"from\":[{\"collectionId\":\"bundles\"}]}},\"readTime\":{\"seconds\":\"1670488464\",\"nanos\":602042000}}},{\"documentMetadata\":{\"name\":\"projects/xxx/databases/(default)/documents/bundles/xxx\",\"readTime\":{\"seconds\":\"1670488464\",\"nanos\":602042000},\"exists\":true,\"queries\":[\"latestBundles\"]}},{\"document\":{\"name\":\"projects/xxx/databases/(default)/documents/bundles/xxx\",\"fields\":{\"data\":{\"stringValue\":\"bar\"}},\"createTime\":{\"seconds\":\"1670487756\",\"nanos\":768917000},\"updateTime\":{\"seconds\":\"1670487756\",\"nanos\":768917000}}},{\"documentMetadata\":{\"name\":\"projects/xxx/databases/(default)/documents/bundles/xxx\",\"readTime\":{\"seconds\":\"1670488464\",\"nanos\":602042000},\"exists\":true,\"queries\":[\"latestBundles\"]}},{\"document\":{\"name\":\"projects/xxx/databases/(default)/documents/bundles/xxx\",\"fields\":{\"data\":{\"stringValue\":\"foo\"}},\"createTime\":{\"seconds\":\"1670487756\",\"nanos\":773423000},\"updateTime\":{\"seconds\":\"1670487756\",\"nanos\":773423000}}}]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bundlesToJSONString(tt.args.str); got != tt.want {
				t.Errorf("bundlesToJSONString() = %v, want %v", got, tt.want)
			}
		})
	}
}
