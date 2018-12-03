const moment = require("moment-timezone");
const { timezone } = require("../config/config");
moment.tz.setDefault(timezone);

/**
 * アクセスログをテーブル形式で出力するmiddleware
 */
module.exports = function(req, res, next) {
  const table = {
    time: moment().format("YYYY/MM/DD HH:mm:ss"),
    url: req.originalUrl,
    type: req.method,
    ip: req.ip
  };
  console.table(table);
  next();
};
