const webpack = require("webpack");

module.exports = {
  entry: "./entry.point",
  output: {
    filename: "dev.inc.js",
    libraryTarget: "this",
  }
};
