const express = require("express");
const mysql = require("mysql");
const app = express();
const port = 3000;

const connection = mysql.createConnection({
  host: "localhost",
  database: "nameList",
});

connection.connect((err) => {
  if (err) throw err;
  console.log("Connected to MySQL database!");
});

app.set("view engine", "ejs");

app.get("/", (req, res) => {
  connection.query("SELECT * FROM kiraKills", (err, results) => {
    if (err) throw err;
    res.render("index.ejs", { data: results });
  });
});

app.listen(port, () => {
  console.log("Server is running at http://localhost:3000");
});

