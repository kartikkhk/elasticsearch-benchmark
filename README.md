# Elasticsearch Benchmarks

### Setup
* To setup elasticsearch on your system run: `docker-compose up -d`
* To run the node server: `cd node` -> `npm ci` -> `npm start`
* To setup k6 go through: [setup-k6-docs](https://github.com/kartikkhk/elasticsearch-benchmarks/tree/master/docs)

### API Reference
[![Run in Postman](https://run.pstmn.io/button.svg)](https://www.getpostman.com/collections/72e3c06691e1b16ceb33)

##### Curl requests
* Write:

```zsh
curl --location --request POST 'http://localhost:3000/write' \
--header 'Content-Type: application/json' \
--data-raw '{
    "postId": 100,
    "id": 500,
    "name": "ex eaque eum natus",
    "email": "Emma@joanny.ca",
    "body": "perspiciatis quis doloremque veniam nisi eos velit sed id totam inventore voluptatem laborum et eveniet aut aut aut maxime quia temporibus ut omnis"
}'
```
* Read:

```zsh
curl --location --request GET 'http://localhost:3000/read?keyword=dignissimos'
```

### Loadtest
* Please install k6 before running loadtests on your local system. [setup-k6-docs](https://github.com/kartikkhk/elasticsearch-benchmarks/tree/master/docs)
* To run a loadtest for writes: `k6 run --vus 10 --duration 30s scripts/write-loadtest.js`
* To run a loadtest for reads: `k6 run --vus 10 --duration 30s scripts/read-loadtest.js`
* Please note that vus refers to **virtual users** and duration is the amount of time you want to run the load test for. 

