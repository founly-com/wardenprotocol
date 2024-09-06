import type { Config } from "tailwindcss";

const config = {
	darkMode: ["class"],
	content: [
		"./pages/**/*.{ts,tsx}",
		"./components/**/*.{ts,tsx}",
		"./app/**/*.{ts,tsx}",
		"./src/**/*.{ts,tsx}",
	],
	prefix: "",
	theme: {
		fontFamily: {
			display: "Syne, Arial, sans-serif",
			sans: "Syne, Arial, sans-serif",
			mono: "Space Mono, monospace",
			spacemono: "Space Mono, monospace",
		},
		container: {
			center: true,
			padding: "2rem",
			screens: {
				"2xl": "1400px",
			},
		},
		extend: {
			colors: {
				"pixel-pink": "#CAFF94",
				"fill-accent-primary": "var(--fill-accent-primary)",
				"fill-accent-secondary": "var(--fill-accent-secondary)",
				"fill-primary": "var(--fill-primary)",
				"fill-gray": "#9EA4AE",
				"fill-quaternary": "var(--fill-quaternary)",
				"fill-elevated": "var(--fill-elevated)",
				"fill-accent-hover": "#BDFF00",
				"bg-elevated": "var(--bg-elevated)",
				"border-edge": "var(--border-edge)",
				"secondary-text": "rgba(229,238,255,0.60)",
				"label-primary": "var(--label-primary)",
				"label-on-light": "var(--label-on-light)",
				"label-tertiary": "var(--label-tertiary)",
				"border-quaternary": "var(--border-quaternary)",
				"border-accent": "var(--border-accent)",
				"label-accent": "var(--label-accent)",
				"label-secondary": "var(--label-secondary)",
				"label-invert": "var(--label-invert)",
				"border-primary": "var(--border-primary)",
				tertiary: "#141414",
				positive: "#48B037",
				"positive-secondary": "rgba(72, 176, 55, 0.15)",
				"fill-negative-secondary": "var(--fill-negative-secondary)",
				negative: "#E54545",
				"negative-secondary": "rgba(229,69,69,0.15)",
				"bg-negative": "rgba(229,69,69,0.15)",
				overlay: "var(--overlay)",
				"hover-bg": "rgba(255,174,238,0.15)",
				"overlay-secondary": "rgba(21,21,21,0.40)",
				"secondary-bg": "var(--secondary-bg)",
				"tertiary-bg": "#482E42",
				lightgray: "#232527",
				checkbox: "rgba(229,238,255,0.60)",
				"border-secondary": "var(--border-secondary)",
				orange: "#E57F45",
				"orange-secondary": "rgba(229, 127, 69, 0.15)",
				"staking-bg": "var(--staking-bg)",
				border: "hsl(var(--border))",
				input: "hsl(var(--input))",
				ring: "hsl(var(--ring))",
				background: "var(--background)",
				foreground: "hsl(var(--foreground))",
				primary: {
					DEFAULT: "hsl(var(--primary))",
					foreground: "hsl(var(--primary-foreground))",
				},
				secondary: {
					DEFAULT: "hsl(var(--secondary))",
					foreground: "hsl(var(--secondary-foreground))",
				},
				destructive: {
					DEFAULT: "hsl(var(--destructive))",
					foreground: "hsl(var(--destructive-foreground))",
				},
				muted: {
					DEFAULT: "hsl(var(--muted))",
					foreground: "hsl(var(--muted-foreground))",
				},
				accent: {
					DEFAULT: "hsl(var(--accent))",
					foreground: "hsl(var(--accent-foreground))",
				},
				popover: {
					DEFAULT: "hsl(var(--popover))",
					foreground: "hsl(var(--popover-foreground))",
				},
				card: {
					DEFAULT: "hsl(var(--card))",
					foreground: "hsl(var(--card-foreground))",
				},
			},
			boxShadow: {
				hoverGlow: "0px 0px 25px 0px #F186DB",
			},
			borderRadius: {
				lg: "var(--radius)",
				md: "calc(var(--radius) - 2px)",
				sm: "calc(var(--radius) - 4px)",
			},
			keyframes: {
				"accordion-down": {
					from: { height: "0" },
					to: { height: "var(--radix-accordion-content-height)" },
				},
				"accordion-up": {
					from: { height: "var(--radix-accordion-content-height)" },
					to: { height: "0" },
				},
			},
			animation: {
				"accordion-down": "accordion-down 0.2s ease-out",
				"accordion-up": "accordion-up 0.2s ease-out",
			},
		},
	},
	plugins: [require("tailwindcss-animate")],
} satisfies Config;

export default config;
