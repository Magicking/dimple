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
      - RECAPTCHA_PRIVATE_KEY=
      - CHAIN_ID=4 #
    links:
      - postgres

  caddy-rev:
    image: abiosoft/caddy
    ports:
      - "0.0.0.0:80:80"
    volumes:
      - "./Caddyfile:/etc/Caddyfile:ro"
      - "./front:/var/www:ro"
    depends_on:
      - dimple

  postgres:
    image: postgres:9.5
    environment:
      - POSTGRES_USER=eb7f87edf260973
      - POSTGRES_PASSWORD=078f2f145dac4222
      - POSTGRES_DB=dimple
```
