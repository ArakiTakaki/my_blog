const crypto = require("crypto");

/**
 * @param {string} data ハッシュ化する文字列
 * @param {string} salt 付与するソルト
 */
const mb5Hash = (data, salt) => {
  const mb5 = crypto.createHash("md5");
  mb5.update(data + salt);
  return mb5.digest("hex");
};

module.exports = {
  mb5Hash
};
