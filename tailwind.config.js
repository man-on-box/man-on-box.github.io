module.exports = {
  content: ["./view/**/*.html"],
  darkMode: "selector",
  theme: {
    fontFamily: {
      sans: ["Roboto Slab", "sans-serif"],
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
