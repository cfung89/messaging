import { IContacts } from "../../scripts/constants";
import useContacts from "../../hooks/useContacts";

import "./Contacts.css";

const Contacts = () => {
  const { contactList } = useContacts();

  function handleContactClick(contact: IContacts) {
    console.log(contact);
  }

  if (contactList.length === 0) {
    return <div>No contacts</div>;
  } else {
    return (
      <div className="grid-container">
        {contactList.map((contact) => {
          return (
            <div
              key={contact.name}
              className="grid-item"
              onClick={() => {
                handleContactClick(contact);
              }}
            >
              {contact.name}
            </div>
          );
        })}
      </div>
    );
  }
};

export default Contacts;
