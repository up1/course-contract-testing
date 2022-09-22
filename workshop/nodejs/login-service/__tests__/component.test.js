const UserService = require("../gateway");

test("Success case with Get user by id = 1", async() => {
  const service = new UserService("http://localhost:4000");
  const acutualResult = await service.getUserById(1);
  expect(acutualResult).toEqual({"name": "Somkiat Pui", "role": "admin", "user": "somkiat"});
});
