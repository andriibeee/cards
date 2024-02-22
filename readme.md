# Example of requests

```sh
curl http://localhost:3000/validate -v -X POST -d "{\"card_number\":\"4111111111111111\", \"exp_month\":\"12\", \"exp_year\": \"13\"}"
```

```sh
curl http://localhost:3000/validate -v -X POST -d "{\"card_number\":\"4111111111111111\", \"exp_month\":\"12\", \"exp_year\": \"28\"}"
```

```sh
curl http://localhost:3000/validate -v -X POST -d "{\"card_number\":\"1111111111111\", \"exp_month\":\"12\", \"exp_year\": \"11\"}"
```