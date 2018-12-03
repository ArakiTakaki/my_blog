const path = require('path');

module.exports = {
  extensions: [".js", ".jsx"],
  alias: {
    '~': path.resolve(__dirname, '../src/js/'),
    'sass': path.resolve(__dirname, '../src/sass')
  }
};
