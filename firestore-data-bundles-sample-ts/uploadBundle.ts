import { firestore, storage } from "firebase-admin";
import { initAdminProject } from "./helper";
import * as fs from "fs";

initAdminProject();

const BUCKET_NAME = "YOUR_BUCKET_NAME";
const CREATE_INITIAL_DATA = false;

const main = async () => {
  const db = firestore();
  const timestamp = Date.now();

  if (CREATE_INITIAL_DATA) {
    [{ data: "foo" }, { data: "bar" }].forEach(async (document) => {
      await db.collection("bundles").add(document);
    });
  }

  const snapshots = await db.collection("bundles").get();

  const buffer = await db
    .bundle(timestamp.toString())
    .add("latestBundles", snapshots)
    .build();
  const bundledFilePath = `./${timestamp}.txt`;
  fs.writeFileSync(bundledFilePath, buffer);
  const destination = `latestBundles.txt`;
  await storage().bucket(BUCKET_NAME).upload(bundledFilePath, {
    destination,
  });
  console.log(
    `uploaded to https://storage.googleapis.com/${BUCKET_NAME}/${destination}`
  );

  process.exit(0);
};

main();
