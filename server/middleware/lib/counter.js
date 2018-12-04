const moment = require("moment-timezone");
const { timezone } = require("../../config/config");
moment.tz.setDefault(timezone);

let time = null;

/**
 * アクセスログをテーブル形式で出力するmiddleware
 */
module.exports.before = function(req, res, next) {
  time = Date.now();
  next();
};

module.exports.after = (req, res, next) => {
  console.log(Date.now() - time);
  next();
};
