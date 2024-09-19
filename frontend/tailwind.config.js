/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        background: "#282828",
        "background-light": "#3c3836",
        foreground: "#ebdbb2",
        comment: "#928374",
        cursor: "#fe8019",
        red: "#fb4934",
        green: "#b8bb26",
        yellow: "#fabd2f",
        blue: "#83a598",
        purple: "#d3869b",
        aqua: "#8ec07c",
        "bright-red": "#cc241d",
        "bright-green": "#98971a",
        "bright-yellow": "#d79921",
        "bright-blue": "#458588",
        "bright-purple": "#b16286",
        "bright-aqua": "#689d6a",
      },
      fontFamily: {
        jetbrains: ["JetBrainsMonoNerdFont", "monospace"],
      },
    },
  },
  plugins: [],
};
