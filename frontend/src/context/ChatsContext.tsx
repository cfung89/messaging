import { ReactNode, useState, createContext } from "react";
import { IChats } from "../scripts/constants";

const ChatsContext = createContext<{
  chatList: Array<IChats>;
  addChatList: (newChat: IChats) => void;
  updateChatList: (newChat: IChats) => void;
}>({
  //default values
  chatList: [],
  addChatList: () => {},
  updateChatList: () => {},
});

interface IChatsContextProps {
  children: ReactNode;
  chats: Array<IChats>;
}

export const ChatsContextProvider = ({
  children,
  chats,
}: IChatsContextProps) => {
  const [chatList, setChatList] = useState<Array<IChats>>(chats);

  const addChatList = (newChat: IChats) => {
    setChatList([...chatList, newChat]);
  };

  const updateChatList = (newChat: IChats) => {
    const idx = chatList.findIndex((x) => x.id === newChat.id);
    setChatList([
      ...chatList.slice(0, idx),
      newChat,
      ...chatList.slice(idx + 1),
    ]);
  };

  return (
    <ChatsContext.Provider value={{ chatList, addChatList, updateChatList }}>
      {children}
    </ChatsContext.Provider>
  );
};

export const ChatsConsumer = ChatsContext.Consumer;

export default ChatsContext;
