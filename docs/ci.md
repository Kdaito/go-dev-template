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