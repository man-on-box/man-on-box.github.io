(function NavToggle() {
  const navBtn = document.getElementById("nav-btn");
  const navMenu = document.getElementById("nav-menu");
  const navLinks = navMenu.querySelectorAll("a");
  let expanded = false;

  navBtn.addEventListener("click", () => {
    toggleNav();
  });

  navLinks.forEach((link) => {
    link.addEventListener("click", () => {
      toggleNav();
    });
  });

  function toggleNav() {
    expanded = !expanded;
    navMenu.classList.toggle("hidden");
    navMenu.classList.toggle("flex");
    navMenu.setAttribute("aria-hidden", !expanded);
    navBtn.setAttribute("aria-expanded", expanded);

    navLinks.forEach((link) => {
      link.tabIndex = expanded ? 0 : -1;
    });
    if (expanded) navLinks[0].focus();
  }
})();

(function DarkToggle() {
  const root = document.documentElement;
  const toggleBtn = document.getElementById("dark-toggle");
  let isDark = root.classList.contains("dark");

  toggleBtn.addEventListener("click", () => {
    isDark = !isDark;
    root.classList.toggle("dark");
    localStorage.setItem("theme", isDark ? "dark" : "light");
  });
})();
