import { useEffect, useState } from "react";

export const Navlist = () => {
  const [activeLink, setActiveLink] = useState("");

  useEffect(() => {
    const handleScroll = () => {
      const sections = [
        { id: "about", link: "about" },
        { id: "features", link: "features" },
        { id: "contact", link: "contact" },
      ];

      const scrollPosition = window.scrollY + window.innerHeight / 2;

      for (let section of sections) {
        const sectionElement = document.getElementById(section.id);
        if (sectionElement) {
          const sectionTop = sectionElement.offsetTop;
          const sectionHeight = sectionElement.offsetHeight;
          if (
            scrollPosition >= sectionTop &&
            scrollPosition < sectionTop + sectionHeight
          ) {
            setActiveLink(section.link);
            return;
          }
        }
      }

      setActiveLink("");
    };

    window.addEventListener("scroll", handleScroll);
    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);

  return (
    <div className="navlist">
      <nav className="nav">
        <a
          href="#about"
          className={`nav-link ${activeLink === "about" ? "active" : ""}`}
        >
          About us
        </a>
        <a
          href="#features"
          className={`nav-link ${activeLink === "features" ? "active" : ""}`}
        >
          Features
        </a>
        <a
          href="#contact"
          className={`nav-link ${activeLink === "contact" ? "active" : ""}`}
        >
          Contact us
        </a>
      </nav>
    </div>
  );
};
