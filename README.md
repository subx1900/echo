# echo
Echo web server

## How to build

```sh
docker build .
```

## How to use

The resulting docker container listens on port 3000 for requests and will return for each request:
- Body
- Method
- Headers
- Path
