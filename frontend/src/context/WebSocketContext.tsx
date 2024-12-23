import { useEffect, useRef, ReactNode, createContext } from "react";
import { PORT, IMessage } from "../scripts/constants";

const WebSocketContext = createContext<{
  sendMessage: (message: IMessage) => void;
}>({
  sendMessage: () => {},
});

export const WebSocketContextProvider = ({
  children,
}: {
  children: ReactNode;
}) => {
  const ws = useRef<WebSocket | null>(null);
  useEffect(() => {
    ws.current = new WebSocket(`ws.current://localhost:${PORT}/ws`);
    ws.current.addEventListener("open", (event: Event) => {
      console.log("WebSocket is open: ", event);
      ws.current?.send("Hello Server");
    });
    ws.current.addEventListener("close", (event: CloseEvent) => {
      console.log("WebSocket is closed: ", event);
    });
    ws.current.addEventListener("message", (event: MessageEvent) => {
      console.log("WebSocket message received: ", event.data);
    });
    ws.current.addEventListener("error", (event: Event) => {
      console.error("WebSocket error: ", event);
    });
    return () => {
      if (ws.current) {
        console.log("Cleaning up WebSocket connection");
        ws.current.close();
        ws.current = null;
      }
    };
  }, []);

  const sendMessage = (message: IMessage) => {
    if (ws.current && ws.current.readyState === WebSocket.OPEN) {
      ws.current.send(JSON.stringify(message));
    } else {
      console.error("WebSocket not open");
    }
  };

  return (
    <WebSocketContext.Provider value={{ sendMessage }}>
      {children}
    </WebSocketContext.Provider>
  );
};

export const WebSocketConsumer = WebSocketContext.Consumer;

export default WebSocketContext;

// const data: IMessage = event.data;
// console.log("WebSocket data:", data);

// const { chatList, addChatList, updateChatList } = useChats();
// const newChat: IChats | undefined = chatList.find(
//   (chat) => chat.id === data.id,
// );

// if (newChat === undefined) {
//   const newChat = {
//     name: "Unnamed Chat",
//     url: `/chats/${data.roomId}`,
//     id: data.roomId,
//     msg: [data],
//   };
//   addChatList(newChat);
// } else {
//   newChat.msg.push(data);
//   updateChatList(newChat);
// }
