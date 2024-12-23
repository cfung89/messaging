export const NAVBAR: Array<INavLink> = [
  { name: "Home", url: "/" },
  { name: "Contacts", url: "/contacts" },
  { name: "Chats", url: "/chats" },
];

export const PORT = 8000;

export interface INavLink {
  name: string;
  url: string;
}

export interface IChats {
  name: string;
  url: string;
  id: string;
  msg: Array<IMessage>;
}

export interface IContacts {
  name: string;
  id: string;
}

export interface IMessage {
  id: string;
  roomId: string;
  time: string;
  sender: string;
  content: string;
}
