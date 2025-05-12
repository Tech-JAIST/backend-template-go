# 実装方法
## 前提知識
### ディレクトリ構造
ディレクトリ構造は以下のようになっています。
```sh
.
├── Dockerfile # GoアプリのDockerfile
├── compose.yaml # Goアプリ, DB, adminerのcompose.yml
├── internal
│   ├── config # 環境変数読み込み
│   ├── db
│   │   ├── migration
│   │   └── model
│   ├── domain # サーバーにおけるメインのデータ型
│   ├── handler # ルーティング
│   └── repository # DB操作
└── main.go
```

## APIの新規実装
順不同で
- `internal/handler/handler.go`にルーティングを追加する
- 対応する処理を`internal/handler/{group名}.go`に記述する
- (DB操作があれば)`internal/repository/{model名}.go`に記述し、`handler/{group名}.go`に呼び出し処理を記述する

と変更すると新たなAPIが実装できます。

## DBスキーマの変更
`db/model`に構造体の定義を記述し、サービス開始前(初期化できる)なら`migration/migrate.go`の`InitialTables`に全ての構造体を記述すると良いです。
そうでないなら`migration/migrations.go`にテーブルを追加する操作を記述し、`migration/migrate.go`の`Migrations`に追加すると良いです。

## 内部データ操作(DBを除く)
開発中、サーバー内においたJSONファイル等を読み書きしたくなることがあると思います。
その際、アーキテクチャに悩むこともあると思います。
その際の方針についてヒントを記します。

### 1. 読み込みのみしたい
設定と同値なので`internal/config`に記述すると良いと思います。

### 2. 読み書きをしたい
`internal/repository`に新たなファイルを作ると良いと思います。
`repository`はデータの保存操作を抽象化したファイル群をおくディレクトリであり、保存先は気にしません。
(ちなみに、MVVMやClean Architectureにrepositoryが登場するので、気になる方は調べてみてください。注意点として、これらで指しているrepositoryと本テンプレートのrepositoryは少し違います。)

## 外部データ操作
`internal/external`、`internal/infrastructure`などのディレクトリを作成するのが良いと思います。
