import { useContext } from "react";
import ChatsContext from "../context/ChatsContext";

const useChats = () => {
  const { chatList, addChatList, updateChatList } = useContext(ChatsContext);

  return { chatList, addChatList, updateChatList };
};

export default useChats;
