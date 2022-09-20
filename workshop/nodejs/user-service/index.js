const { server } = require("./user.js");
const port = process.env.API_PORT || 4000;

server.listen(port, () => {
  console.log(`User Service listening on http://localhost:${port}`);
});