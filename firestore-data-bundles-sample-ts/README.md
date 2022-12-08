## 手順
#### セットアップ
```foo.sh
npm install

# ※ GCSやFirestoreを触れる権限のユーザ
gcloud auth application-default login
```

#### helper.ts内のCONFIGを書き換え

1. Firebaseコンソール
2. 歯車アイコン
3. プロジェクトの設定
4. マイアプリでウェブアプリを選択
5. SDKの設定と構成に記載の`firebaseConfig`をコピペ

#### 各定数の書き換え

1. `YOUR_BUCKET_NAME` in `uploadBundle.ts`
2. `UPLOADED_BUNDLE_URL` in `useBundle.ts`


#### 実行

```foo.sh
# 実行後Firestore上のbundlesコレクション生成（CREATE_INITIAL_DATA = true）
# GCS上にもファイル生成
npx ts-node uploadBundle.ts

npx ts-node useBundle.ts
```