import axios from "axios";
import { Dispatch } from "react";

const PORT = 8000;

interface IChat {
  name: string;
}

const getInitial = (setChats: Dispatch<IChat>) => {
  axios.get(`https://localhost:${PORT}/chats`).then((res) => {
    setChats(res.data);
  });
};

const upgradeWS = () => {
  const ws = new WebSocket(`ws://localhost:${PORT}/ws`);

  ws.addEventListener("open", (event) => {
    console.log("WebSocket is open");
    ws.send("Hello Server");
    console.log("Sending Hello");
  });

  ws.addEventListener("close", (event) => {
    console.log("WebSocket is closed");
  });

  ws.addEventListener("message", (event) => {
    const data = event.data;
    console.log("WebSocket data:", data);
  });

  ws.addEventListener("error", (event) => {
    console.error("WebSocket error:", event);
  });
};

export default { getInitial, upgradeWS };
