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

## CI

とりあえず Github action で CI を実行

- build
  - go モジュールのインポート
  - ビルド
- unit test
  - `go test`で実行
- lint
  - (golangci-lint)[https://golangci-lint.run]で実行（とりあえずデフォルト設定）

### 補足

- 最初に build を実行し、モジュールのインポート & 結果をキャッシュに
- build の後、unit test と lint が並列で走る

## ディレクトリ構成

https://github.com/golang-standards/project-layout/tree/master

ここを参考に作ってみた。

この方針だとapiの実装は`internal/app`に作っていくことになりそう？

- cmd
  - `internal/app`に作成したappを実行する。
- internal
  - アプリの本体をここに作っていくイメージ
- scripts
  - スクリプト置いていく
  - 現在はlinterだけだが、マイグレーションやtestなどのスクリプトが今後増えてきた場合はここに置いていくイメージ