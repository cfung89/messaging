import {
  createBrowserRouter,
  createRoutesFromElements,
  Route,
  RouterProvider,
} from "react-router-dom";

// Context
import { ChatsContextProvider } from "./context/ChatsContext";
import { ContactsContextProvider } from "./context/ContactsContext";
import { WebSocketContextProvider } from "./context/WebSocketContext";

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

const router = createBrowserRouter(
  createRoutesFromElements(
    <Route path="/" element={<RootLayout />}>
      <Route index element={<Home />} />
      <Route path="contacts" element={<Contacts />} />
      <Route path="chats" element={<ChatTable />}>
        <Route path=":chatID" element={<Chat />} />
      </Route>
      <Route path="/notfound" element={<NotFound />} />
      <Route path="*" element={<NotFound />} />
    </Route>,
  ),
);

const App = () => {
  // Initialize variables
  const chatList: Array<IChats> = [
    // { name: "Unnamed Chat", url: `/chats/${self.crypto.randomUUID()}` },
  ];

  const contactsList: Array<IContacts> = [
    { name: "Contact 1", id: "/" },
    { name: "Contact 2", id: "/contacts" },
    { name: "Contact 3", id: "/chats" },
  ];

  return (
    <WebSocketContextProvider>
      <ContactsContextProvider contacts={contactsList}>
        <ChatsContextProvider chats={chatList}>
          <RouterProvider router={router} />
        </ChatsContextProvider>
      </ContactsContextProvider>
    </WebSocketContextProvider>
  );
};

export default App;
