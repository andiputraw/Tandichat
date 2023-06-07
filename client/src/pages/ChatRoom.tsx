import React from "react";
import ChatNavigation from "../components/chatroom/ChatNavigation";
import ChatMessage from "../components/chatroom/ChatMessage";

const ChatRoom = ({ onLogout }: { onLogout: Function }) => {

    return (
        <div className="max-h-screen w-screen bg-black text-blue-50 overflow-hidden flex">
            <ChatNavigation onLogout={onLogout}/>
            <ChatMessage />
        </div>
    )
}

export default ChatRoom