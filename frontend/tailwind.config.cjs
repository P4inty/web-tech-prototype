/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./index.html', './src/**/*.{svelte,js,ts}'],
  theme: {
    extend: {
      colors: {
        primary: "#ff6060",
        "primary-dark": "#ff606060",
        secondary: "#5ef5ff"
      }
    },
    fontFamily: {
      sans: ["Barlow", "sans-serif"]
    }
  },
  plugins: [],
}
