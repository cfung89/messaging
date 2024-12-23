import { IChats } from "../../scripts/constants";
import { Outlet, useParams, useNavigate } from "react-router-dom";
import useChats from "../../hooks/useChats";

import "./ChatTable.css";

const ChatTable = () => {
  const { chatList } = useChats();
  const { chatID } = useParams();
  const navigate = useNavigate();

  function handleChatClick(chat: IChats) {
    navigate(chat.url);
  }

  return (
    <div>
      {chatID ? (
        <Outlet />
      ) : (
        <table className="table-container">
          <tbody>
            {chatList.map((chat: IChats) => {
              return (
                <tr key={chat.id}>
                  <th
                    className="table-item"
                    onClick={() => {
                      handleChatClick(chat);
                    }}
                  >
                    {chat.name}
                  </th>
                </tr>
              );
            })}
          </tbody>
        </table>
      )}
    </div>
  );
};

export default ChatTable;
