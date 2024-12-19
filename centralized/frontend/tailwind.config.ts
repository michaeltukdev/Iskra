import type { Config } from "tailwindcss";

export default {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        background: "var(--background)",
        foreground: "var(--foreground)",
        'foreground-border': "var(--foreground-border)",
        input: {
          DEFAULT: "var(--input)",
        },
        'input-border': {
          DEFAULT: "var(--input-border)",
        },
        text: {
          light: "var(--text-light)",
          medium: "var(--text-medium)",
          dark: "var(--text-dark)",
          spare: "var(--text-spare)",
        },
        primary: {
          DEFAULT: "var(--primary)",
          hover: "var(--primary-hover)",
        }
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
} satisfies Config;
