# RabbitMQRepo


```
    docker-compose up
    go run worker/worker.go
```

For sending requests to the cache server, use the following approach:

1. use this url: http://localhost:8081/publish [POST request]
2. with the following json:

```
{
	"requestType": 0,
	"quoteObj": {
		"stock": "S",
		"price": 7.5,
		"userid": "test123"
	}
}
