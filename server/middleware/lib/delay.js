require("colors");
module.exports = function(req, res, next) {
  const delay = 1000 * Math.random();
  console.log(`delay ${delay} ms`.green);
  setTimeout(() => {
    next();
  }, delay);
};
