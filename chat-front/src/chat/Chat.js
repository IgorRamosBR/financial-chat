import React, { useEffect, useState, useCallback } from "react";
import { useAuthState } from "react-firebase-hooks/auth";
import { useHistory } from "react-router";
import { auth, logout, getCurrentUser } from "../firebase";
import "./Chat.css"

function Chat() {
    const [user, loading] = useAuthState(auth);
    const [name, setName] = useState("");
    const history = useHistory();
    const fetchUserName = useCallback( async () => {
        try {
            const query = await getCurrentUser(user)
            const data = await query.docs[0].data();
            setName(data.name);
        } catch (err) {
            console.error(err);
            alert("An error occured while fetching user data");
        }
    }, [user]);
    
    useEffect(() => {
        if (loading) return;
        if (!user) return history.replace("/");
        fetchUserName();
    }, [user, loading, history, fetchUserName]);
    return (
        <div className="chat">
            <div className="chat__container">
                <h1>Financial Chat </h1>
                <div><b>User: </b>{name}  </div>
                <div><b>Rooms: </b>
                    <select name="romms" id="rooms">
                        <option value="room1">room1</option>
                        <option value="room2">room2</option>
                        <option value="room3">room3</option>
                        <option value="room4">room4</option>
                    </select>
                </div>

                <div>
                    <div className={"chat__message_thirds"}>
                        <spam>{user?.email} - Camilla</spam>
                        <spam> Teste chat</spam>
                    </div>

                    <div className={"chat__message_user"}>
                        <spam>{user?.email} - Igor</spam>
                        <spam> Chat testado</spam>
                    </div>
                </div>
                <div className={"chat__container_send_message"}>
                    <input/>
                    <button className="chat__btn" onClick={logout}>
                        Send
                    </button>
                </div>
            </div>
        </div>
    );
}

export default Chat;