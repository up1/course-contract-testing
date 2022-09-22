const UserService = require("../gateway");
const path = require("path");
const {
  PactV3,
  MatchersV3,
  SpecificationVersion,
} = require("@pact-foundation/pact");
const { like } = MatchersV3;

const provider = new PactV3({
  consumer: "login-service",
  provider: "user-service",
  log: path.resolve(process.cwd(), "logs", "pact.log"),
  logLevel: "warn",
  dir: path.resolve(process.cwd(), "pacts"),
  spec: SpecificationVersion.SPECIFICATION_VERSION_V2,
});

test("Success to get user detail", async () => {
  // set up Pact interactions
  await provider.addInteraction({
    states: [{ description: "user exist" }],
    uponReceiving: "get user by id",
    withRequest: {
      method: "GET",
      path: "/users/1",
    },
    willRespondWith: {
      status: 200,
      headers: {
        "Content-Type": "application/json; charset=utf-8",
      },
      body: like({
        user: "somkiat",
        name: "Somkiat Pui",
        role: "admin",
      }),
    },
  });

  await provider.executeTest(async (mockService) => {
    const service = new UserService(mockService.url);
    // make request to Pact mock server
    const user = await service.getUserById(1);
    expect(user).toStrictEqual(
      {
        user: "somkiat",
        name: "Somkiat Pui",
        role: "admin",
      }
    );
  });
}); 

test("Failure :: User not found", async () => {
  // set up Pact interactions
  await provider.addInteraction({
    states: [{ description: "user not found" }],
    uponReceiving: "get user by id = 2",
    withRequest: {
      method: "GET",
      path: "/users/2",
    },
    willRespondWith: {
      status: 404
    },
  });

  await provider.executeTest(async (mockService) => {
    const service = new UserService(mockService.url);
    // make request to Pact mock server
    const response = await service.getUserById(2);
    expect(response).toStrictEqual(
      {
        error: "User not found"
      }
    );
  });
}); 