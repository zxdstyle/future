/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        './app/**/*.{js,ts,jsx,tsx}',
        './pages/**/*.{js,ts,jsx,tsx}',
        './components/**/*.{js,ts,jsx,tsx}',
        './src/**/*.{js,ts,jsx,tsx}',
    ],
    theme: {
        extend: {
            colors: {
                gray: { DEFAULT: 'gray' },
            },
            zIndex: {
                top: 9999,
            },
            scale: {
                200: '2',
            },
        },
    },
    plugins: [],
};
