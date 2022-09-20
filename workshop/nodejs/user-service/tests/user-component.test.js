

const request = require("supertest");
const { server } = require("../user");

const expectedUser = {
    user: "somkiat",
    name: "Somkiat Pui",
    role: "admin",
  };

test("should return 200", async () => {
  await request(server)
    .get("/users/1")
    .expect("Content-Type", /json/)
    .expect(200)
    .then((response) => {
      expect(response.body).toStrictEqual(expectedUser);
    });
});