# Vehicle Server

This servers communicates with a postgis database and allows to manage a collection of vehicles and find the closest one from a given position.

## How to use

### Running the Server

Prerequsites:

- go 1.22
- docker

```bash
# Starts the DB.
make dev_db

# Starts the server.
make dev

# When you're done, stop the db.
make stop_dev_db
```

### Running the Unit Test Suite

```bash
make unit_test
```

### Running the Integration Test Suite

```bash
make integration_test
```

### API examples.

#### Create a Vehicle

```bash
curl --header "Content-Type: application/json" --data '{"latitude": 3.32,"longitude": 4.323, "shortcode":"abed", "battery": 10}' localhost:8080/vehicles | jq .
```

#### Find Nearest Vehicles

```bash
curl localhost:8080/vehicles\?latitude=34.2\&longitude=23.4\&limit=10
```

#### Delete a Vehicle

```bash
curl --request delete localhost:8080/vehicles/${VEHICLE_ID}
```
