## ディレクトリ構成

https://github.com/golang-standards/project-layout/tree/master

ここを参考に作ってみた。

この方針だとapiの実装は`internal/app`に作っていくことになりそう？

- assets
  - リポジトリ固有の画像等を格納
- cmd
  - `internal/app`に作成したappを実行する。
- internal
  - アプリの本体をここに作っていくイメージ
- scripts
  - スクリプト置いていく
  - 現在はlinterだけだが、マイグレーションやtestなどのスクリプトが今後増えてきた場合はここに置いていくイメージ
- docs
  - ドキュメント格納用