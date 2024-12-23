import { ReactNode, useState, createContext } from "react";
import { IContacts } from "../scripts/constants";

const ContactsContext = createContext<{
  contactList: Array<IContacts>;
  saveContactList: (newContactList: IContacts) => void;
}>({
  contactList: [],
  saveContactList: () => {},
});

interface IContactsContextProps {
  children: ReactNode;
  contacts: Array<IContacts>;
}

export const ContactsContextProvider = ({
  children,
  contacts,
}: IContactsContextProps) => {
  const [contactList, setContactList] = useState<Array<IContacts>>(contacts);

  const saveContactList = (newContactList: IContacts) => {
    setContactList([...contactList, newContactList]);
  };

  return (
    <ContactsContext.Provider value={{ contactList, saveContactList }}>
      {children}
    </ContactsContext.Provider>
  );
};

export const ContactsConsumer = ContactsContext.Consumer;

export default ContactsContext;
