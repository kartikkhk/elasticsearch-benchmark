const express = require("express");
const app = express();
const morgan = require("morgan");
const utils = require("./lib/utils");

const PORT = process.env.PORT || 3000;
app.use(morgan("dev"));
app.use(express.json());

app.get("/ping", (_, res) => res.send("AoK"));

/**
 * @description {write} post request
 */
app.post("/write", async (req, res) => {
  const data = req.body;
  try {
    const result = await utils.insert("comments", "comment", data);
    return res.status(201).json({ status: "success", statusCode: 201, result });
  } catch (err) {
    return res
      .status(500)
      .json({ error: err.message, statusCode: 500, status: "failure" });
  }
});

/**
 * @description {read} get request
 */
app.get("/read", async (req, res) => {
  const query = req.query;
  try {
    const result = await utils.search("comments", query);
    return res.status(200).json({ status: "success", statusCode: 200, result });
  } catch (err) {
    return res
      .status(500)
      .json({ error: err.message, statusCode: 500, status: "failure" });
  }
});

app.listen(PORT, () => {
  console.log(`listening on port ${PORT}`);
});
