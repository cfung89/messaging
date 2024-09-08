import {
  createBrowserRouter,
  createRoutesFromElements,
  Route,
  RouterProvider,
} from "react-router-dom";

// Context
import { ChatsContextProvider } from "./context/ChatsContext";
import { ContactsContextProvider } from "./context/ContactsContext";

// Layouts
import RootLayout from "./layouts/RootLayout";

// Components
import Home from "./components/Home/Home";
import Contacts from "./components/Contacts/Contacts";
import ChatTable from "./components/ChatTable/ChatTable";
import Chat from "./components/Chat/Chat";
import NotFound from "./components/NotFound/NotFound";

// Constants
import { IContacts, IChats } from "./scripts/constants";

// Initialize variables
const chatList: Array<IChats> = [
  { name: "Home", url: "/" },
  { name: "Contacts", url: "/contacts" },
  { name: "Chats", url: "/chats" },
];

const contactsList: Array<IContacts> = [
  { name: "Home", id: "/" },
  { name: "Contacts", id: "/contacts" },
  { name: "Chats", id: "/chats" },
];

const router = createBrowserRouter(
  createRoutesFromElements(
    <Route path="/" element={<RootLayout />}>
      <Route index element={<Home />} />
      <Route path="contacts" element={<Contacts />} />
      <Route path="chats" element={<ChatTable />}>
        <Route path=":chatID" element={<Chat />} />
      </Route>
      <Route path="*" element={<NotFound />} />
    </Route>,
  ),
);

const App = () => {
  return (
    <ContactsContextProvider contacts={contactsList}>
      <ChatsContextProvider chats={chatList}>
        <RouterProvider router={router} />
      </ChatsContextProvider>
    </ContactsContextProvider>
  );
};

export default App;
