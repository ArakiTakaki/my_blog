const sqlite3 = require("sqlite3");

module.exports.init = file => {
  return new sqlite3.Database(file);
};
