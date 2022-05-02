# go_fyne_app

- VSCode で作業する際、go の import で指定したパッケージの指定が消える現象が起きるため、以下の設定をする
  - VSCode の拡張機能から[Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)をインストールする
  - `ファイル > ユーザ設定 > 設定` を開き、go の Format Tool を検索する
  - `default` になっている場合、 `gofmt` へ変更する
- モジュール対応モードで動かすことを前提にしている
  - [Go のモジュール管理【バージョン 1.17 改訂版】](https://zenn.dev/spiegel/articles/20210223-go-module-aware-mode)
  - `go env -w GO111MODULE=on`
- `go mod init go_fyne_app` で Go モジュールモードの前提
- go の依存関係解決
  - [Go 言語の依存モジュール管理ツール Modules の使い方 - Qiita](https://qiita.com/uchiko/items/64fb3020dd64cf211d4e)

---

## 構成

|     |        |
| :-: | :----: |
| Go  | 1.18.1 |

---

## コマンド

### インストール

```bash
go get fyne.io/fyne
go get
```

- インストールしたら、 `GOPATH` 配下の `pkg/mod/fyne.io` が生成される
  - 中身は `fyne` パッケージ
- はじめてインストールすると、 `go.sum` というファイルがプロジェクトに追加される
- `go get` をすることで、依存モジュールの追加をする

### ビルド

```bash
go build -o build goファイルパス
```

- `build` に指定した go ファイルのビルドを書き出す
