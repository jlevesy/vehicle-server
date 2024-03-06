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
