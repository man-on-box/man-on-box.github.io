(function NavToggle() {
  const navBtn = document.getElementById("nav-btn");
  const navMenu = document.getElementById("nav-menu");
  const navLinks = document.querySelectorAll(".nav-link");

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
  const toggleBtn = document.getElementById("dark-toggle");
  const key = "color-scheme";
  const prefersDark =
    window &&
    window.matchMedia &&
    window.matchMedia("(prefers-color-scheme: dark)").matches;
  const userPreference = localStorage.getItem(key);
  let isDark = userPreference !== "light" && prefersDark;
  if (isDark) {
    root.classList.add("dark");
  }

  toggleBtn.addEventListener("click", () => {
    isDark = !isDark;
    root.classList.toggle("dark");
    localStorage.setItem(key, isDark ? "dark" : "light");
  });
})();
