export const NAVBAR: Array<INavLink> = [
  { name: "Home", url: "/" },
  { name: "Contacts", url: "/contacts" },
  { name: "Chats", url: "/chats" },
];

export interface INavLink {
  name: string;
  url: string;
}

export interface IChats {
  name: string;
  url: string;
}

export interface IContacts {
  name: string;
  id: string;
}
