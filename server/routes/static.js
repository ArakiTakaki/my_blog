const express = require("express");
const r = express.Router();

const fallback = require("express-history-api-fallback");
r.use(express.static("public"));
r.use(fallback("index.html", { root: "public" }));

module.exports = r;
