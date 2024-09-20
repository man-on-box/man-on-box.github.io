module.exports = {
  content: ["./view/**/*.html", "./data/**/*.go"],
  darkMode: "selector",
  theme: {
    fontFamily: {
      sans: ["Roboto Slab", "sans-serif"],
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
