const { Scanner, Utils } = require("spherov2.js");

const run = async () => {
  const sphero = await Scanner.findSpheroMini();

  if (!sphero) return console.log("sphero mini not available!");

  // appendContent

};

run();
