(function NavToggle() {
  const navBtn = document.getElementById("nav-btn");
  const navMenu = document.getElementById("nav-menu");
  const navLinks = document.querySelectorAll("nav .nav-link");

  navBtn.addEventListener("click", () => {
    toggleNav();
  });

  navLinks.forEach((link) => {
    link.addEventListener("click", () => {
      toggleNav();
    });
  });

  function toggleNav() {
    navMenu.classList.toggle("hidden");
    navMenu.classList.toggle("flex");
  }
})();

(function DarkToggle() {
  const root = document.documentElement;
  let isDark = root.classList.contains("dark");

  const toggleBtn = document.getElementById("dark-toggle");
  toggleBtn.addEventListener("click", () => {
    isDark = !isDark;
    root.classList.toggle("dark");
    localStorage.setItem("theme", isDark ? "dark" : "light");
  });
})();
