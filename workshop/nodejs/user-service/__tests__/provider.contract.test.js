const { Verifier } = require("@pact-foundation/pact");
const { fail } = require("assert");
const path = require("path");
const { server } = require("../user");

// Setup provider server to verify
const app = server.listen(4000);

test("validates the expectations of user-service", async() => {
  const opts = {
    logLevel: "INFO",
    providerBaseUrl: "http://localhost:4000",
    provider: "user-service",
    pactUrls: [
      path.resolve(__dirname, "../../pacts/login-service-user-service.json"),
    ],
  };
   await new Verifier(opts).verifyProvider()
    .then((output) => {
      console.log(output);
    })
	.catch(e => {
	  fail(e); 
	})
    .finally(() => {
      app.close();
    });
});

afterEach(async() => {
    await app.close();
});