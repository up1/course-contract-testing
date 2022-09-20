const express = require("express");
const server = express();

server.use(express.json());
server.use((_, res, next) => {
  res.header("Content-Type", "application/json; charset=utf-8");
  next();
});

server.get("/users/:id", (req, res) => {
	res.json({
        user: "somkiat",
        name: "Somkiat Pui",
        role: "admin"
      });
});

module.exports = {
  server
};