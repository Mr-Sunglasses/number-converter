# Roman To Decimal Converter

A simple REST API to Convert Roman Numeral to an Decimal Value

[![Go](https://github.com/Mr-Sunglasses/number-converter/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/Mr-Sunglasses/number-converter/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/Mr-Sunglasses/number-converter/branch/main/graph/badge.svg?token=MK95H1VJVL)](https://codecov.io/gh/Mr-Sunglasses/number-converter)

---

## End Points

1. **POST**: `127.0.0.1:8080/roman`

   Post Body

   ```
   {"roman": "MCMXCIV"}
   ```

   Response

   ```
   {"integer":1994,"roman":"MCMXCIV"}
   ```

## Installation

1. Clone the repository

   ```bash
   git clone https://github.com/Mr-Sunglasses/number-converter
   cd roman-to-no-api
   ```

2. Install dependencies

   ```bash
   go get ./...
   ```

3. Run Tests

   ```bash
   go test ./... -v -coverpkg=./... -coverprofile=coverage.out
   go tool cover -html=coverage.out
   ```

4. Run Gin Development server
   ```bash
   go run main.go
   ```

## License

Copyright © 2023 [Kanishk Pachauri](https://github.com/Mr-Sunglasses).<br />
This project is [MIT](https://github.com/Mr-Sunglasses/number-converter/blob/master/LICENSE) licensed.
