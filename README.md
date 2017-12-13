# Dimple

This is an example faucet to demonstrate basic use on how to send value on ethereum

## Example docker-compose file

```version: '2'
services:
  dimple:
    build: .
    environment:
      - WS_URI=ws://localhost:8546
      - DB_DSN=host=postgres user=eb7f87edf260973 dbname=dimple sslmode=disable password=078f2f145dac4222
      - PRIVATE_KEY=f485e851f93defcd877e6212a4c51351fc49a335a7cd155f713a5647c6496a56
      - NETWORK_ID=4 #
    ports:
      - "0.0.0.0:8000:8090"
    links:
      - postgres

  postgres:
    image: postgres:9.5
    environment:
      - POSTGRES_USER=eb7f87edf260973
      - POSTGRES_PASSWORD=078f2f145dac4222
      - POSTGRES_DB=dimple
```
