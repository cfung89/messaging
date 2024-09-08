import { useContext } from "react";
import ChatsContext from "../context/ChatsContext";

const useChats = () => {
  const context = useContext(ChatsContext);

  return context;
};

export default useChats;
