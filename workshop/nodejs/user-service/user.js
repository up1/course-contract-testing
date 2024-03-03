const express = require("express");
const server = express();

server.use(express.json());
server.use((_, res, next) => {
  res.header("Content-Type", "application/json; charset=utf-8");
  next();
});

server.get("/users/:id", (req, res) => {
	if(req.params.id == 1) {
    res.json({
          user: "somkiat",
          name: "Somkiat Pui",
          role: "admin"
        });
  } else {
    res.status(404).send('User not found')
  }
});

module.exports = {
  server
};