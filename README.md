ローカルでlinterを動かしたい時

```sh
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.62.2 golangci-lint run -v
```
