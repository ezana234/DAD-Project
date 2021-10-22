import firebase from 'firebase/compat/app';
import "firebase/compat/firestore";
import "firebase/compat/auth";


// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
    apiKey: "AIzaSyDbtXQHM2v51zSuvRqfbOc-yHYAIGD0yLc",
    authDomain: "cfcapp-85ce5.firebaseapp.com",
    projectId: "cfcapp-85ce5",
    storageBucket: "cfcapp-85ce5.appspot.com",
    messagingSenderId: "601445591749",
    appId: "1:601445591749:web:778e0ed7aedb9b28421eac",
    measurementId: "G-K64N3Z0D67"
  };

  const firebaseApp = firebase.initializeApp(firebaseConfig);
  const db = firebaseApp.firestore();
  const auth = firebase.auth();

  export {db,auth};