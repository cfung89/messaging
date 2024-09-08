import { useState } from "react";
import { Outlet } from "react-router-dom";

import Navbar from "../components/Navbar/Navbar";
import ChatList from "../components/ChatList/ChatList";

import "./RootLayout.css";

const RootLayout = () => {
  const [sidebarOpen, setSidebarOpen] = useState<boolean>(true);

  return (
    <div
      className={`root-layout ${sidebarOpen ? "sidebar-expanded" : "sidebar-collapsed"}`}
    >
      <header>
        <Navbar />
      </header>
      <aside>
        <ChatList sidebarOpen={sidebarOpen} setSidebarOpen={setSidebarOpen} />
      </aside>

      <main>
        <Outlet />
      </main>
    </div>
  );
};

export default RootLayout;
