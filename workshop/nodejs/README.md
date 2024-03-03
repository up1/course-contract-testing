# Workshop Contract testing with Pact + NodeJS
* Login service
* User service
* [Pact :: mactching rules](https://docs.pact.io/implementation_guides/javascript/docs/matching#v3-matching-rules)

## Step to run provider :: User service (port = 4000)

Run user service
```
$npm install
$npm run provider-run
```

Test url = http://localhost:4000/users/1


## Step to run consumer :: Login service

Run login service
```
$npm install
$npm run consumer-run
```

Ready to testing ...


## Step to run Contract testing
```
// Run contract testing to create a contract
$npm run consumer-contract-test  

// Upload contract to broker
$docker compose up pact_broker_publish --remove-orphans

// Verify
$docker compose up pact_broker_verify --remove-orphans

// Start provider server
$npm run provider-run

// Verify again
$docker compose up pact_broker_verify --remove-orphans

```

## Testing
* [Component testing](https://github.com/up1/course-contract-testing/wiki/NodeJS#step-2--component-testing)
* [Contract testing](https://github.com/up1/course-contract-testing/wiki/NodeJS#step-3--contract-testing-with-pact-in-consumer-side)

