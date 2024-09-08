import { NavLink } from "react-router-dom";

import { NAVBAR } from "../../scripts/constants";
import "./Navbar.css";

const Navbar = () => {
  return (
    <nav className="navbar">
      {NAVBAR.map((link) => (
        <NavLink to={link.url} key={link.name}>
          {link.name}
        </NavLink>
      ))}
    </nav>
  );
};

export default Navbar;
