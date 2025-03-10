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

(function registerFadeInElements() {
  const elements = document.querySelectorAll("[data-fade-in]");
  if (elements.length === 0) return;
  const staggerDelay = 150;
  const duration = 500;
  const intersectionObserver = new IntersectionObserver(
    (entries, observer) => {
      entries.forEach((entry, i) => animateEntry(entry, i, observer));
    },
    {
      root: null,
      threshold: 0.1,
    },
  );
  elements.forEach((element) => {
    element.classList.add("transition");
    element.classList.add("duration-500");
    intersectionObserver.observe(element);
  });
  const animateEntry = (entry, index, observer) => {
    const element = entry.target;
    if (!entry.isIntersecting) {
      element.classList.add("opacity-0");
      element.classList.add("translate-y-10");
      return;
    }
    observer.unobserve(element);
    setTimeout(() => {
      element.classList.remove("opacity-0");
      element.classList.remove("translate-y-10");
      removeTransitionAfterAnimated(element);
    }, index * staggerDelay);
  };
  const removeTransitionAfterAnimated = (element) => {
    setTimeout(() => {
      element.classList.remove("transition");
      element.classList.remove("duration-500");
    }, duration);
  };
})();
