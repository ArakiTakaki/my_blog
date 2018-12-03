require("colors");
module.exports = function(req, res, next) {
  const delay = 1000 * Math.random();
  console.log(`delay ${delay}`.green);
  setTimeout(() => {
    next();
  }, delay);
};
