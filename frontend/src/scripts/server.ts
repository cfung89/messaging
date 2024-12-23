import axios from "axios";
import { Dispatch } from "react";
import { PORT, IChats } from "./constants";

const getData = (setChats: Dispatch<IChats>) => {
  axios.get(`https://localhost:${PORT}/chats`).then((res) => {
    setChats(res.data);
  });
};

const login = () => {};

export default { getData, login };
