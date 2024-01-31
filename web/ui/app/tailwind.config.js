module.exports = {
  content: [
    './app/index.html',
    './app/templates/**/*.hbs',
    './app/components/**/*.hbs',
    './app/components/**/*.js',
    './app/modifiers/**/*.js',
  ],
  plugins: [require('@tailwindcss/forms')],
  darkMode: 'class',
};