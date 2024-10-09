/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        customBlue: "#7aa2f7",
        background: "#99000000",
        "background-alt": "#373B41",
        foreground: "#ffffff",
        primary: "#00aaff",
        secondary: "#4c7899",
        alert: "#A54242",
        disabled: "#707880",
        Transparent: "#FF0000",
        textColor: "#F0C674",
      },
    },
  },
  plugins: [],
};
