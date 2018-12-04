const express = require("express");
const r = express.Router();
// r.get(route, r)

r.get("/api/sample", function(req, res) {
  return res.json({ sample: "sample" });
});

module.exports = r;
