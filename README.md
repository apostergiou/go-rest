# rest api in Golang

Start the server:

```shell
$ go build
$ ./go-rest
```

Usage:

```shell
# Index
$ curl -i http://localhost:8000/items

# Show
$ curl -i http://localhost:8000/items/1

# Delete
curl -i -X DELETE http://localhost:8000/items/1
```
