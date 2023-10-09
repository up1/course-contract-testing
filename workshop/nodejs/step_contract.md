## 1. Create contract
```
$npm i
$npm run consumer-contract-test
```

## 2. Create Pact Broker
```
$docker compose up -d postgres
$docker compose up -d pact_broker
```

## 3. Publish contract(s) to Pact Broker
```
$docker compose up pact_broker_publish --remove-orphans
```

## 4. Verify contract(s) from Pact Broker
```
$npm run provider-run
$docker compose up pact_broker_verify --remove-orphans
```