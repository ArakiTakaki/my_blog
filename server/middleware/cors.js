module.exports = function(req, res, next) {
  res.setHeader("Access-Control-Allow-Origin", "http://localhost:3030");
  res.setHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE");
  res.setHeader(
    "Access-Control-Allow-Headers",
    "Origin, Authorization, Accept, Content-Type"
  );
  res.setHeader("Access-Control-Allow-Credentials", "true");

  next();
};
