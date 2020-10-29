# todo

My example of a todo app.

## Validate swagger file

```sh
swagger \
  validate \
  ./swagger.yml
```

## Generate code

```sh
swagger \
  generate \
  server \
  -A todo \
  -f ./swagger.yml
```

## Start the app

```sh
go run cmd/todo-server/main.go --port 8080
```
