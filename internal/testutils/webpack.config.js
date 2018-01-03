const webpack = require("webpack");

module.exports = {
  entry: "./entry.point",
  output: {
    filename: "testutils.inc.js",
    libraryTarget: "this",
  }
};
