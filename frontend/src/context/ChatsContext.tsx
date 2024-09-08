import { ReactNode, useState, useEffect, createContext } from "react";
import { IChats } from "../scripts/constants";

const ChatsContext = createContext<Array<IChats>>([]);

interface IChatsContextProps {
  children: ReactNode;
  chats: Array<IChats>;
}

export const ChatsContextProvider = ({
  children,
  chats,
}: IChatsContextProps) => {
  const [chatList, setChatList] = useState<Array<IChats>>(chats);

  const saveChatList = (newChat: Array<IChats>) => {
    useEffect(() => {
      setChatList(newChat);
    }, []);
  };

  return (
    <ChatsContext.Provider value={{ chatList: chatList, saveChatList }}>
      {children}
    </ChatsContext.Provider>
  );
};

export const ChatsConsumer = ChatsContext.Consumer;

export default ChatsContext;
