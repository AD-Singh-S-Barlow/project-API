const express = require("express");
const mysql = require("mysql");
const app = express();
const port = 3000;

const connection = mysql.createConnection({
  host: "localhost",
  user: "root",
  
  database: "nameList",
});

connection.connect((err) => {
  if (err) {
    console.error("Error connecting to MySQL database:", err);
    return;
  }
  console.log("Connected to MySQL database!");
});

app.set("view engine", "ejs");

app.get("/", (req, res) => {
  connection.query("SELECT * FROM kiraKills", (err, results) => {
    if (err) {
      console.error("Error executing MySQL query:", err);
      return;
    }
    res.render("index.ejs", { data: results });
  });
});

app.listen(port, () => {
  console.log("Server is running at http://localhost:3000");
});


