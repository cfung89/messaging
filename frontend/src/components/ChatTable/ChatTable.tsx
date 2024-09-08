import { IChats } from "../../scripts/constants";
import useChats from "../../hooks/useChats";

import "./ChatTable.css";

const ChatTable = () => {
  const { chatList } = useChats();
  function handleChatClick(chat: IChats) {
    console.log(chat);
  }

  return (
    <div className="grid-container">
      {chatList.map((chat) => {
        return (
          <div
            key={chat.name}
            className="grid-item"
            onClick={() => {
              handleChatClick(chat);
            }}
          >
            {chat.name}
          </div>
        );
      })}
    </div>
  );
};

export default ChatTable;
