import {
  useState,
  useEffect,
  Dispatch,
  SetStateAction,
  BaseSyntheticEvent,
} from "react";

import { NavLink } from "react-router-dom";
import useChats from "../../hooks/useChats";

import "./ChatList.css";

interface IChatListProps {
  sidebarOpen: boolean;
  setSidebarOpen: Dispatch<SetStateAction<boolean>>;
}

const ChatList = ({ sidebarOpen, setSidebarOpen }: IChatListProps) => {
  const [searchChat, setSearchChat] = useState("");
  const { chatList } = useChats();

  function handleSearchChange(event: BaseSyntheticEvent) {
    setSearchChat(event.target.value);
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

  if (chatList.length === 0) {
    return <div>No chats</div>;
  } else {
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
            <nav className="sidebar-content">
              <ul>
                {chatList
                  .filter((chat) =>
                    chat.name.toUpperCase().includes(searchChat.toUpperCase()),
                  )
                  .map((chat) => (
                    <li key={chat.name}>
                      <NavLink to={chat.url}>{chat.name}</NavLink>
                    </li>
                  ))}
              </ul>
            </nav>
          </>
        )}
      </div>
    );
  }
};

export default ChatList;
