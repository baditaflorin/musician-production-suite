/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{ts,tsx}"],
  theme: {
    extend: {
      colors: {
        ink: "#15171c",
        panel: "#f8f8f5",
        line: "#d8d7ce",
        signal: "#0d766e",
        warn: "#a35416"
      }
    }
  },
  plugins: []
};

