/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./internal/templates/**/*.{templ, html}",],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/typography'),
    require("daisyui")
  ],
}

