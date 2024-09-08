import { ReactNode, useState, useEffect, createContext } from "react";
import { IContacts } from "../scripts/constants";

const ContactsContext = createContext<Array<IContacts>>([] as IContacts[]);

interface IContactsContextProps {
  children: ReactNode;
  contacts: Array<IContacts>;
}

export const ContactsContextProvider = ({
  children,
  contacts,
}: IContactsContextProps) => {
  const [contactList, setContactList] = useState<Array<IContacts>>(contacts);

  const saveContactList = (newContact: Array<IContacts>) => {
    useEffect(() => {
      setContactList(newContact);
    }, []);
  };

  return (
    <ContactsContext.Provider
      value={{ contactList: contactList, saveContactList }}
    >
      {children}
    </ContactsContext.Provider>
  );
};

export const ContactsConsumer = ContactsContext.Consumer;

export default ContactsContext;
