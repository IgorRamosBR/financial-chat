import { initializeApp } from 'firebase/app';
import { getAuth, createUserWithEmailAndPassword, signInWithEmailAndPassword } from "firebase/auth";
import { getFirestore, collection, addDoc, getDocs, query, where } from 'firebase/firestore/lite';

const firebaseConfig = {
    apiKey: "AIzaSyAZzQEMuRS42RdqF34yTkXMZSYOSYoQJYo",
    authDomain: "financial-chat.firebaseapp.com",
    projectId: "financial-chat",
    storageBucket: "financial-chat.appspot.com",
    messagingSenderId: "70087257448",
    appId: "1:70087257448:web:1026ef93b3a0b06fd67add"
};

const app = initializeApp(firebaseConfig);
const auth = getAuth();
const db = getFirestore(app);

const signIn = (email, password) => {
    signInWithEmailAndPassword(auth, email, password)
        .then(() => {
            console.log("Logged with success")
        })
        .catch((error) => {
            console.error(error);
            alert(error.message)

        });
};

const registerWithEmailAndPassword = (name, email, password) => {
    createUserWithEmailAndPassword(auth, email, password)
        .then((userCredential) => {
            // Signed in
            const user = userCredential.user;
            console.log(user)
             addDoc(collection(db, "users"), {
                uid: user.uid,
                name,
                authProvider: "local",
                email,
            });
        })
        .catch((error) => {
            console.error(error);
            alert(error.message)
        });
};

const getCurrentUser = (user) => {
    return getDocs(query(collection(db, "users"), where("uid", "==", user?.uid)));
}

const logout = () => {
    auth.signOut();
};

export {
    auth,
    db,
    signIn,
    registerWithEmailAndPassword,
    getCurrentUser,
    logout,
};