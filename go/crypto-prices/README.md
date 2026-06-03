# crypto-prices

## Grabbing crypto prices to practice Go

### requirements

- .env file with API_KEY set to your CoinMarketCap API key

### usage

#### local go

Ensure you have a CoinMarketCap API Key (free) added to your `.env` file, and just run `go run .` from this directory

#### local docker container

```
# build, push & run the docker image
docker build -t <dockerhub-user>/crypto-prices:latest .
docker push <dockerhub-user>/crypto-prices:latest
docker run  --env-file ${PWD}/.env -p 8080:8080 <dockerhub-user>/crypto-prices:latest
```
