import { initializeApp } from "firebase-admin";
import firebase from "firebase/app";
import "firebase/firestore";
import "firebase/firestore/bundle";

export const initAdminProject = () => {
  return initializeApp();
};

export const initClientProject = () => {
  const firebaseConfig = {
    // YOUR CONFIG
  };

  return firebase.initializeApp(firebaseConfig);
};
