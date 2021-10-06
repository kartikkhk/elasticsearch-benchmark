const elasticsearch = require("elasticsearch");
const client = new elasticsearch.Client({
  host: process.env.HOST || "localhost:9200",
  log: "error",
  apiVersion: "7.x",
});

module.exports = client;
