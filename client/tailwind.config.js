/** @type {import('tailwindcss').Config} */

module.exports = {
  prefix: "tw-",
  content: ["./internal/views/**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
};
