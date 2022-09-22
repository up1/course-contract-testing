const nock = require('nock')
const UserService = require("../gateway");

test("Success case with Get user by id = 1", async() => {

  nock('http://localhost:4000')
  .get('/users/1')
  .reply(200, {"name": "Somkiat Pui", "role": "admin", "user": "somkiat"})

  const service = new UserService("http://localhost:4000");
  const acutualResult = await service.getUserById(1);
  expect(acutualResult).toEqual({"name": "Somkiat Pui", "role": "admin", "user": "somkiat"});
});

test("Failure case with Get user by id = 2 => User not found", async() => {

  nock('http://localhost:4000')
  .get('/users/2')
  .reply(404, {})

  const service = new UserService("http://localhost:4000");
  const acutualResult = await service.getUserById(2);
  expect(acutualResult).toEqual({error: "User not found"});
});
