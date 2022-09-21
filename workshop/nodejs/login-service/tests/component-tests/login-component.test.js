const UserService = require("../../gateway");
const nock = require("nock");

const URL = "http://127.0.0.1:4000";

test("Success to get user by id from user server", async () => {
  const mockedUser = {
    user: "somkiat",
    name: "Somkiat Pui",
    role: "admin",
  };
  nock(URL).get("/users/1").reply(200, mockedUser);
  const service = new UserService(URL);
  const response = await service.getUserById(1);
  expect(response).toEqual(mockedUser);
});

test("Failure to get user by id from user server :: user not found", async () => {
  const expectedError = {code: 404};
  nock(URL).get("/users/2").reply(404);
  const service = new UserService(URL);
  const response = await service.getUserById(2);
  expect(response).toEqual(expectedError);
});
