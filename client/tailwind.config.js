/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: "class",
  content: [
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        "light-theme": "#eef0f4"
      },
      fontFamily: {
        "display": ["Oswald", "sans-serif"],
        "body": ["Poppins", "sans-serif"],
        "post": ["Nunito", "sans-serif"]
      }
    },
  },
  plugins: [],
}
