# Vehicle Server

This servers communicates with a postgis database and allows to manage a collection of vehicles and find the closest one from a given position.

## How to use

### Running the Server

```bash
go run ./cmd/server -listen-address :8080 -database-url ${POSTGRES_CONN_STRING}
```

### Running the Integration Test Suite

```bash
go test -v -count=1  ./app
```
