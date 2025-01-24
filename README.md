# gin-gonic.com-docs-examples-querystring-param

- Query string parameters

- Reference: https://gin-gonic.com/docs/examples/querystring-param/

## gvm

```sh
gvm install go1.23.5
gvm use go1.23.5
```

## go get

```sh
go get .
```

## go run

```sh
go run .
```

## cURL

- GET /welcome

```sh
curl --location --request GET 'localhost:8080/welcome?firstname=Hello&lastname=World' \
--form 'message="MESSAGE"'
```
