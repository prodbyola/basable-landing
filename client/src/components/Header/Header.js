import "./Header.scss";
import X from "./assets/X.svg";
import Logo from "./assets/logo.svg";
import React, { useState } from "react";
import { Navlist } from "../Nav/Navlist";
import Github from "./assets/Github.svg";
import { makeStyles } from "@mui/styles";
import Drawer from "@mui/material/Drawer";
import LinkedIn from "./assets/LinkedIn.svg";
import { HeaderButtons } from "../Nav/Buttons";
import HamburgerIcon from "./assets/Hamburger.png";
import { Divider, Link, List, ListItem } from "@mui/material";
import { LogoSvg } from "./assets";

const drawerWidth = 240;
const useStyles = makeStyles((theme) => {
  return {
    drawer: {
      width: drawerWidth,
    },
    drawerPaper: {
      width: drawerWidth,
    },
    headerButtons: {
      display: "flex",
      flexWrap: "wrap",
      gap: "32px",
      marginTop: "2rem",
      marginLeft: "auto",
      marginRight: "auto",
    },
    btn: {
      width: "fit-content",
      padding: "0.5rem 1rem",
      border: "none",
      borderRadius: "4px",
      marginTop: "2rem",
      marginInline: "20px",
      cursor: "pointer",
      background: "#4451ca",
      color: "#fff",
      fontFamily: "Lexend",
      fontSize: "16px",
      fontWeight: "400",
      lineHeight: "20px",
      transition: " background 0.3s",
      "&:hover": {
        background: "#0056b3",
      },
    },
    icon: {
      height: "32px",
      transition: "opacity 0.3s",

      "&:hover": {
        opacity: "0.7",
      },
    },
    active: {
      backgroundColor: "#f4f4f4",
    },
    logo: {
      maxWidth: "129.69px",
      marginTop: "24px",
      marginInline: "10px",
      marginBottom: "34px",
    },
  };
});

const Header = () => {
  const [activeLink, setActiveLink] = useState("about");
  const [menuOpen, setMenuOpen] = useState(false);
  const sections = [
    { id: "about", link: "about" },
    { id: "features", link: "features" },
    { id: "contact", link: "contact" },
  ];
  const classes = useStyles();

  return (
    <header className="header">
      <div className="container">
        <Drawer
          className={classes.drawer}
          open={menuOpen}
          onClose={() => setMenuOpen(!menuOpen)}
          anchor="right"
          classes={{ paper: classes.drawerPaper }}
        >
          <div className={classes.logo}>
            <LogoSvg />
          </div>

          <List>
            {sections.map((section) => (
              <Link
                onClick={() => {
                  setActiveLink(section.link);
                  console.log(activeLink);
                  setMenuOpen(!menuOpen);
                }}
                key={section.id}
                underline="none"
                className={`nav-link`}
                href={`#${section.id}`}
              >
                <ListItem
                  href={`#${section.id}`}
                  sx={{
                    paddingBottom: "24px",
                    paddingTop: "24px",
                  }}
                  className={
                    activeLink === section.link ? classes.active : null
                  }
                >
                  {section.id}
                </ListItem>
              </Link>
            ))}
          </List>
          <Divider />

          <button className={classes.btn}>Join now</button>
          <div className={classes.headerButtons}>
            <a
              href="https://www.linkedin.com"
              target="_blank"
              rel="noopener noreferrer"
            >
              <img src={LinkedIn} alt="LinkedIn" className={classes.icon} />
            </a>
            <a
              href="https://www.twitter.com"
              target="_blank"
              rel="noopener noreferrer"
            >
              <img src={X} alt="X" className={classes.icon} />
            </a>
            <a
              href="https://www.github.com"
              target="_blank"
              rel="noopener noreferrer"
            >
              <img src={Github} alt="GitHub" className={classes.icon} />
            </a>
          </div>
        </Drawer>
        <img src={Logo} alt="Logo" className="logo" />
        <div className="nav-m">
          <Navlist />
        </div>
        <div className="nav-m">
          <HeaderButtons />
        </div>

        <img
          src={HamburgerIcon}
          alt="Menu"
          className="hamburger-icon"
          onClick={() => setMenuOpen(!menuOpen)}
        />
      </div>
    </header>
  );
};

export default Header;
