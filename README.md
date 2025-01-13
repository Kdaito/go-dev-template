## 環境構築

docker

```sh
docker -v

Docker version 27.4.0, build bde2b89
```

docker-compose

```sh
docker-compose -v

Docker Compose version 2.32.1
```

起動

```sh
docker-compose up -d
```

linter 

```sh
./scripts/lint.sh
```

ローカルでのtest, build
```sh
同様にDockerコンテナをたてて行う？
```

lintもdockerイメージ上で行えるように実装
(開発者はdockerとdocker-composeさえあれば開発を開始できるようにした方が良いのかなと思いこうしました。)


## その他のドキュメント
- [`CIについて`](docs/ci.md)
- [`ディレクトリ構成`](docs/directory.md)
- [`アーキテクチャ`](docs/architecture.md)
- [`エラーハンドリング`](docs/error.md)
