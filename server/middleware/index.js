const express = require("express");
const m = express.Router();

/**
 * Depelopmentのみに適応させるMddleware
 */
if (process.env.MODE === "development") {
  m.use("/api", require("./cors"));
  m.use(require("./delay"));
  m.use(require("./log"));
  m.use(require("./sessionLog"));
}

// 本番環境にのみ設定するもの
// if (process.env.MODE === "PRODCTION") {
// }

/**
 * passportの設定
 */
const passport = require("../middleware/passport");
m.use(passport.middleware);
m.use(passport.session);

module.exports = m;
