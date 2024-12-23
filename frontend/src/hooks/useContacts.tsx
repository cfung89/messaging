import { useContext } from "react";
import ContactsContext from "../context/ContactsContext";

const useContacts = () => {
  const { contactList, saveContactList } = useContext(ContactsContext);

  return { contactList, saveContactList };
};

export default useContacts;
