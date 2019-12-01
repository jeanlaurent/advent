const readLine = require("readline");
const fs = require("fs");

module.exports = {
  getInputRL: function(day) {
    const rl = readLine.createInterface({
      input: fs.createReadStream(day),
      crlfDelay: Infinity
    });
    return rl;
  }
};