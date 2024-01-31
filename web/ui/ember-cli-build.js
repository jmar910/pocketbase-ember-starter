'use strict';

const EmberApp = require('ember-cli/lib/broccoli/ember-app');
const autoprefixer = require('autoprefixer');

const isProduction = EmberApp.env() === 'production';

module.exports = function (defaults) {
  const app = new EmberApp(defaults, {
    minifyCSS: {
      enabled: false,
    },
    postcssOptions: {
      compile: {
        cacheInclude: [/.*\.(css|scss|hbs)$/, /.tailwind\.config\.js$/],
        plugins: [
          {
            module: require('postcss-import'),
            options: {
              path: ['node_modules'],
            },
          },
          require('tailwindcss')('./app/tailwind.config.js'),
          autoprefixer,
          ...(isProduction ? [require('cssnano')({ preset: 'default' })] : []),
        ],
      },
    },
  });

  return app.toTree();
};
