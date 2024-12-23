import { useEffect, useRef } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { IMessage, IChats } from "../../scripts/constants";
import NotFound from "../NotFound/NotFound";
import useChats from "../../hooks/useChats";
import "./Chat.css";

const Chat = () => {
  const { chatID } = useParams();
  const navigate = useNavigate();
  const { chatList, updateChatList } = useChats();
  const messageRef = useRef(null);
  const newChat: IChats | undefined = chatList.find(
    (chat) => chat.id === chatID,
  );

  useEffect(() => {
    if (newChat === undefined) {
      navigate("/notfound");
    }
  }, [chatID]);

  if (newChat === undefined) {
    return <NotFound />;
  }

  const handleSubmit: React.FormEventHandler<any> = (e) => {
    e.preventDefault();

    if (messageRef.current === null || messageRef.current.value === "") {
      return;
    }

    const currentTime = new Date().toLocaleString();
    newChat.msg = [
      ...newChat.msg,
      {
        id: self.crypto.randomUUID(),
        time: currentTime,
        sender: true,
        content: messageRef.current.value,
      },
    ];
    updateChatList(newChat);

    // send message through websocket to server
    // console.log(messageRef.current.value);

    messageRef.current.value = null;
  };

  const messageEnterSubmit = (e: any) => {
    if (e.ctrlKey && e.key === "Enter") {
      handleSubmit(e);
    }
  };

  return (
    <div>
      <br />
      <div className="message-container">
        {newChat.msg.map((msg: IMessage) => (
          <div
            key={msg.id}
            className={msg.sender ? "message-send" : "message-receive"}
          >
            {msg.content}
          </div>
        ))}
      </div>
      <form onSubmit={handleSubmit}>
        <textarea
          className="message-box"
          onKeyUp={messageEnterSubmit}
          placeholder="Enter message..."
          ref={messageRef}
        />
        <button type="submit">Send</button>
      </form>
    </div>
  );
};

export default Chat;
