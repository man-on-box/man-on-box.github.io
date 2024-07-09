module.exports = {
  content: ["./**/*.html"],
  theme: {
    fontFamily: {
      sans: ["Wittgenstein", "sans-serif"],
      serif: ["Wittgenstein", "serif"],
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
