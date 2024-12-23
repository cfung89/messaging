import {
  useState,
  useEffect,
  Dispatch,
  SetStateAction,
  BaseSyntheticEvent,
} from "react";

import { NavLink, useNavigate } from "react-router-dom";
import useChats from "../../hooks/useChats";
import useRClick from "../../hooks/useRClick";
import { CLRClick } from "./chatListStyles";

import "./ChatList.css";
import { IChats } from "../../scripts/constants";

interface IChatListProps {
  sidebarOpen: boolean;
  setSidebarOpen: Dispatch<SetStateAction<boolean>>;
}

const ChatList = ({ sidebarOpen, setSidebarOpen }: IChatListProps) => {
  const navigate = useNavigate();
  const [searchChat, setSearchChat] = useState("");
  const { chatList, addChatList } = useChats();

  // for right-click menu
  const { clicked, setClicked, points, setPoints } = useRClick();

  function handleSearchChange(event: BaseSyntheticEvent) {
    setSearchChat(event.target.value);
  }

  function onNewChat() {
    const id = self.crypto.randomUUID();
    const newChat = {
      name: "Unnamed Chat",
      url: `/chats/${id}`,
      id: id,
      msg: [],
    };
    addChatList(newChat);
    navigate(newChat.url);
  }

  useEffect(() => {
    const handleResize = () => {
      if (window.innerWidth < 600 && sidebarOpen) {
        setSidebarOpen(false);
      }
    };
    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, [sidebarOpen]);

  return (
    <div
      className={`sidebar ${sidebarOpen ? "sidebar-open" : "sidebar-closed"}`}
    >
      <table className="sidebar-table">
        <tbody>
          <tr>
            <td>{sidebarOpen && <div className="logo">Messaging</div>}</td>
            <td>
              <button
                className="sidebar-toggle"
                onClick={() => setSidebarOpen(!sidebarOpen)}
              >
                {sidebarOpen ? "Close" : "Open"}
              </button>
            </td>
          </tr>
        </tbody>
      </table>

      {sidebarOpen && (
        <>
          <input
            className="search-bar"
            value={searchChat}
            onChange={handleSearchChange}
          />
          <div className="sidebar-content">
            <ul>
              {chatList
                .filter((chat: IChats) =>
                  chat.name.toUpperCase().includes(searchChat.toUpperCase()),
                )
                .map((chat: IChats) => (
                  <li
                    onContextMenu={(e) => {
                      e.preventDefault();
                      setClicked(true);
                      setPoints({ x: e.pageX, y: e.pageY });
                    }}
                    key={chat.url}
                  >
                    <NavLink to={chat.url}>{chat.name}</NavLink>
                  </li>
                ))}
            </ul>
            {clicked && (
              <CLRClick $top={points.y} $left={points.x}>
                <ul>
                  <li>Rename</li>
                  <li>Delete</li>
                </ul>
              </CLRClick>
            )}
          </div>
          <div className="new-chat-container">
            <button
              type="submit"
              className="new-chat-button"
              onClick={() => onNewChat()}
            >
              New Chat
            </button>
          </div>
        </>
      )}
    </div>
  );
};

export default ChatList;
