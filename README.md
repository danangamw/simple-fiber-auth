# Project test-coding

Simple REST API with Fiber (Golang)

## MakeFile

Run build make command with tests

```bash
make all
```

Build the application

```bash
make build
```

Run the application

```bash
make run
```

Live reload the application:

```bash
make watch
```

Run the test suite:

```bash
make test
```

Clean up binary from the last build:

```bash
make clean
```

## Example Request

Public Endpoint:

```bash
curl http://localhost:8080/public
```

```bash
{
 "message": "This is a public endpoint"
}
```

Protected Endpoint:

```bash
curl http://localhost:8080/protected -H "Authorization: Bearer mySecretTok3n"
```

response:

```bash
{
 	"message": "This is a protected endpoint.",
	"userID":  123,
}
```

request:

```bash
curl http://localhost:8080/protected
```

response:

```bash
{
    "error":"Missing Authorization header"
}
```

request:

```bash
curl http://localhost:8080/protected -H "Authorization: Bearer myS"
```

response:

```bash
{
 	"message":"Invalid token"
}
```
