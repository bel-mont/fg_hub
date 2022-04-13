module.exports = {
  content: ["./src/**/*.{html,js}"],
  theme: {
    extend: {
      boxShadow: {
        mdr: '0 4px 15px -5px rgb(0 0 0 / 0.25)'
      },
      colors: {
        red: {
          "50": "#FFE5E5",
          "100": "#FFC7C7",
          "200": "#FF8F8F",
          "300": "#FF5757",
          "400": "#FF2424",
          "500": "#E90000",
          "600": "#BD0000",
          "700": "#8A0000",
          "800": "#5C0000",
          "900": "#2E0000"
        }
      }
    },
  },
  plugins: [],
}
