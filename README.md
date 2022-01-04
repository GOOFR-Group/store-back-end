# Store Back-End

## Quickstart

### Build & Run (using docker-compose)

```bash
# build containers and deploy
docker-compose up --build
```

### Reset Database

```bash
# drops the volumes attached to the containers
docker-compose down -v
# brings up the containers
docker-compose up
```

## API

### Documentation

[http://localhost:8080/docs](http://localhost:8080/docs)

## Others

### Access database manually

```bash
psql -U goofr -d store_db
```
