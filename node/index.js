const app = require("express")();
const morgan = require("morgan");

const PORT = process.env.PORT || 3000;
app.use(morgan("dev"));

app.get("/ping", (_, res) => res.send("AoK"));

app.listen(PORT, () => {
  console.log(`listening on port ${PORT}`);
});
