Basic Go API server.

start with 
```go run . ```

Get initial books:
```curl localhost:8080/getbooks```

example Post, to add a new record:
```curl -X POST -H "Content-Type: application/json" -d '{"id": "1237", "title": "hacking kubernetes", "author": "AM", "quantity": 1}' localhost:8080/postbooks ```

Get list of books with our new one added from above, the new record should be added:
```curl localhost:8080/getbooks```

