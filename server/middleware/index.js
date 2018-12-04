const express = require("express");
const m = express.Router();
const morgan = require("morgan");
const bodyParser = require("body-parser");

m.use(bodyParser.urlencoded({ extended: true }));
m.use(bodyParser.json());

/**
 * Depelopmentのみに適応させるMddleware
 */
if (process.env.MODE === "development") {
  m.use("/api", require("./lib/cors"));
  m.use(require("./lib/delay"));
  m.use(morgan("short"));
}

// 本番環境にのみ設定するもの
// if (process.env.MODE === "PRODCTION") {
// }

module.exports = m;
