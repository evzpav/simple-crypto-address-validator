# Simple Crypto Address Validator

Simplistic cryptocurrency address validator that checks length and first chars of wallet address.
Usage: get request to endpoint to verify if address is valid or not after user fills input.
It is very simple but saves time not waiting for node to reply for failure.


## Run locally :
```bash
go run main.go
# Server will be running on http://localhost:8888

curl localhost:8888/validate/btc/1CFNjwLjZdSKB8nZopxhLaR8vvqaQKD3Bi
curl http://addressvalidator.evzpav.com/validate/btc/1CFNjwLjZdSKB8nZopxhLaR8vvqaQKD3Bi
```

### Success return:
```
{
    "ok":true, //if no error is true
    "crypto":"btc", //same as received
    "address":"1CFNjwLjZdSKB8nZopxhLaR8vvqaQKD3Bi", //same as received
    "valid":true //same as received
}
```

### Error return:
```
{
    "ok":false,
    "crypto":"bt",
    "address":"1CFNjwLjZdSKB8nZopxhLaR8vvqaQKD3Bi",
    "valid":false,
    "error":"Cryptocurrency not available: bt"
}
```

### Run unit tests:
```
go test
```
