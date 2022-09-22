const supertest = require("supertest");
const { server } = require("../user");

test("Success case with Get user by id = 1", async () => {
  await supertest(server)
    .get("/users/1")
    .expect(200)
    .then((response) => {
      expect(response.body).toEqual({
        user: "somkiat",
        name: "Somkiat Pui",
        role: "admin",
      });
    });
});
