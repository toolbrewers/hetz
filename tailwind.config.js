/** @type {import('tailwindcss').Config} */

module.exports = {
  prefix: 'tw-',
  content: ["./app/views/**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
};
