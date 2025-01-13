# エラーハンドリングについて
APIのエラー発生時に使用するエラーオブジェクト、ハンドリングについて

## 使用するエラーオブジェクト
`internal/lib/errors`にカスタムエラーがあるのでそれを使用する。

`echo`にも`echo.*HTTPError`があるが、echoに依存しない層からもエラーオブジェクトを投げれるように、共通のカスタムエラーを用意

カスタムエラー: [errors.Error](../internal/lib/errors/errors.go)

## エラーハンドリング

[middleware.error_handler](../internal/application/middleware/error_handler.go)で実装

基本的にはカスタムエラーを返す想定だが、標準の`error`や`echo.*HTTPError`が返却された場合にも対応している。

### 標準のerrorを返却した場合
標準エラーを返却した場合、500レスポンスが返却され、APIログにエラー内容が出力される

レスポンス
```json
{
  "code": 500,
  "message": "internal server error"
}
```

APIのログ
```sh
2025-01-13 13:37:00 {"time":"2025-01-13T04:37:00.460143978Z","level":"ERROR","prefix":"echo","file":"error_handler.go","line":"22","message":"'errors.Error()'の実行内容"}
```

### カスタムエラーを返却した場合
カスタムエラーを返却した場合、カスタムエラー初期化時に指定したステータスコードでレスポンスが返却され、APIログに初期化時に指定したメッセージが出力される

例
```go
return nil, errors.Error(http.StatusNotFound, "user not found id = 5")
```
のようにカスタムエラーを返却した場合

レスポンス
```json
{
  "code": 404,
  "message": "Not found"
}
```

APIのログ
```sh
2025-01-13 13:37:00 {"time":"2025-01-13T04:37:00.460143978Z","level":"ERROR","prefix":"echo","file":"error_handler.go","line":"22","message":"[404] user not found id = 5"}
```

指定したコードに対して出力されるメッセージの対応は[こちら](https://go.dev/src/net/http/status.go)

### 注意点
また、カスタムエラー初期化時に第二引数に渡すメッセージはAPIのログのみに出力され、レスポンスには設定されない
