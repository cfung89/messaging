import { useContext } from "react";
import ContactsContext from "../context/ContactsContext";

const useContacts = () => {
  const context = useContext(ContactsContext);

  return context;
};

export default useContacts;
