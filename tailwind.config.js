module.exports = {
  content: ["./**/*.html"],
  darkMode: "selector",
  theme: {
    fontFamily: {
      sans: ["Roboto Slab", "sans-serif"],
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
