# backend-template-go
JAIST Techサークルのバックエンドテンプレート(Go)です。

## プロジェクト作成方法
### CLI
```sh
go run golang.org/x/tools/cmd/gonew@latest github.com/Tech-JAIST/backend-template-go {{ project_name }}
```
もしくは
```sh
go install golang.org/x/tools/cmd/gonew@latest
gonew github.com/Tech-JAIST/backend-template-go {{ project_name }}
```

### GUI
右上の `Use this template` からレポジトリを作成できます。
作成したのち、`go.mod`と関連するファイルのモジュール名を作成したプロジェクト名に変更してください。
ex: `github.com/Tech-JAIST/hackathon-project1`

## Tools
### VSCode
エディタはVSCode、もしくはそのフォークであるCursorを推奨しています。
拡張機能として、`Go`をインストールしてください。

### Docker
[公式ページ](https://www.docker.com/)
仮想環境を立ち上げるために必要です。
インストール方法は[こちら](docs/docker.md)を参考にしてください。

### Go
実行環境はDockerにて提供されるため必須ではありませんが、xc,golangci-lintを使う場合は必要になります。
Goの実行環境を[公式サイト](https://go.dev/doc/install)からインストールしてください。
インストールしたのち、`go/bin`を環境変数に設定してください。

### golangci-lint
linter, formatterを動作させるために必要です。
```
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
```

### xc
Go製のタスクランナーです。
Tasksに記述されたタスクを実行するために必要です。
```
go install github.com/joerdav/xc/cmd/xc@latest
```

## Tasks

開発に用いるコマンド一覧

> [!TIP]
> `xc` を使うことでMDのコマンドを簡単に実行できます。
> 参考: [MarkdownベースのGo製タスクランナー「xc」のススメ](https://zenn.dev/trap/articles/af32614c07214d)
>
> ```sh
> go install github.com/joerdav/xc/cmd/xc@latest
> ```
>
> 実行例
> ```sh
> xc dev
> ```

### build

アプリをビルドします。

```sh
go mod download
go build -o ./$(basename $PWD)
```

### dev

ホットリロードの開発環境を構築します。

```sh
docker compose watch
```

API、DB、DB管理画面が起動します。
各コンテナが起動したら、以下のURLにアクセスすることができます。
Compose Watchにより、ソースコードの変更を検知して自動で再起動します。

- <http://localhost:8080/> (API)
- <http://localhost:8081/> (DBの管理画面)

### lint

golangci-lintを実行します。

```sh
golangci-lint run --fix ./...
```

### lint_with_docker

golangci-lintのdockerイメージをpullしてlinterを実行します。
```sh
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v2.1.6-alpine golangci-lint run --fix ./...
```

### format

golangci-lintでformatを実行します。

```sh
golangci-lint fmt ./...
```

### format_with_docker

golangci-lintのdockerイメージをpullしてformatを実行します。

```sh
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v2.1.6-alpine golangci-lint fmt ./...
```

## References
参考にしたテンプレート: [ras0q/go-backend-template](https://github.com/ras0q/go-backend-template/tree/main)