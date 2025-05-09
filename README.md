# backend-template-go
JAIST Techサークルのバックエンドテンプレート(Go)です。

## Tasks

開発に用いるコマンド一覧

> [!TIP]
> `xc` を使うことでMDのコマンドを簡単に実行できます。
> 参考: [MarkdownベースのGo製タスクランナー「xc」のススメ](https://zenn.dev/trap/articles/af32614c07214d)
>
> ```bash
> go install github.com/joerdav/xc/cmd/xc@latest
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

### format

golangci-lintでformatを実行します。

```sh
golangci-lint fmt ./...
```

## References
参考にしたテンプレート: [ras0q/go-backend-template](https://github.com/ras0q/go-backend-template/tree/main)