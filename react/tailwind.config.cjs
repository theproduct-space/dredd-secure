/* eslint-env node */
/** @type {import('tailwindcss').Config} */

const pxToRem = (dest) => 1 / (16 / dest);

module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
    "./node_modules/@ignt/react-library/src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    boxShadow: {
      DEFAULT: "16px 32px 128px -8px rgba(0, 0, 0, 0.07)",
    },
    fontFamily: {
      revalia: ["Revalia", "cursive"],
      sans: ["Inter", "sans-serif"],
    },
    orangeOnHover: {
      DEFAULT: "hover:text-orange",
    },
    extend: {
      colors: {
        red: "red", // for tests, TODO remove
        transparent: "transparent",
        current: "currentColor",
        white: {
          200: "rgba(255,255,255,0.2)",
          500: "rgba(255,255,255,0.5)",
          800: "rgba(255,255,255,0.82)",
          1000: "#f9f9f9",
        },
        black: "#11100B",
        gray: "#241F16",
        coolSilver: "#F0FBFF",
        orange: "#FF8A00",
        yellow: "#FFC75A",
        // functional colors
        bg: "#11100B",
        buttonBg: "#32302B",
        title: "#ffffff",
        text: "#f9f9f9",
        inverse: "#000",
        muted: "rgba(0, 0, 0, 0.667)",
        inactive: "rgba(0, 0, 0, 0.33)",
        link: "#000",
        linkHover: "rgba(0, 0, 0, 0.667)",
        border: "rgba(0, 0, 0, 0.07)",
        checkbox: "#C4C4C4",
        radio: "#C4C4C4",
        warning: "#FC8C0B",
        error: "#D80228",
        negative: "#D80228",
        notification: "#FE475F",
      },
      maxWidth: {
        "app-max": "1600px",
      },
      boxShadow: {
        border: "0px 0.5px 0px #CBCBCB",
        border_double: "0px -0.5px 0px #CBCBCB, 0px 0.5px 0px #CBCBCB",
        max: "16px 32px 128px 8px rgba(0, 0, 0, 0.07)",
        dropdown: "-112px 73px 191px 59px rgba(0, 0, 0, 0.09)",
        select: "40px 64px 128px -8px rgba(0, 0, 0, 0.14)",
        outline: "inset 0 0 0 1px rgba(9, 78, 253, 1)",
        std: "3px 9px 32px -4px rgb(0 0 0 / 7%)",
      },
      dropShadow: {
        orange: [
          "-35px 35px 85px rgba(255, 199, 90, 0.75)",
          "35px 35px 25px rgba(255, 138, 0, 0.75)",
        ],
        orangeCenter: [
          "-0px 0px 45px rgba(255, 199, 90, 0.45)",
          "5px 5px 25px rgba(255, 138, 0, 0.45)",
        ],
        yellow: [
          "-35px 35px 100px rgba(255, 199, 90, 0.25)",
          "35px 35px 25px rgba(255, 138, 0, 0.25)",
        ],
        lightOrange: [
          "-35px 35px 35px rgba(255, 199, 90, 0.15)",
          "35px 35px 15px rgba(255, 138, 0, 0.05)",
        ],
      },
      fontSize: {
        h1: "4rem",
        h2: "2.75rem",
        p1: "1.5rem",
        xxs: "0.625rem",
      },
      inset: {
        center: "50%",
      },
      lineHeight: {
        title: "127%",
        text: "153.8%",
      },
      borderRadius: {
        "2sm": "8px",
        "3sm": "10px",
        "4sm": "12px",
        "5sm": "14px",
        "6sm": "16px",
      },
      // height: {
      //   screen: ["100vh /* fallback for Opera, IE and etc. */", "100dvh"],
      // },
    },
    letterSpacing: {
      tighter: "-.007em",
      tight: "-.02em",
      normal: "0",
    },

    screens: {
      // mobile first -> (min-width:xxx)
      xs: `${pxToRem(320)}rem`,
      sm: `${pxToRem(576)}rem`,
      md: `${pxToRem(768)}rem`,
      lg: `${pxToRem(1024)}rem`,
      llg: `${pxToRem(1200)}rem`,
      xl: `${pxToRem(1380)}rem`,
    },
  },
  plugins: [require("@headlessui/tailwindcss")],
};
