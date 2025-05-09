# Docker
## インストール
### macOS
[OrbStack](https://orbstack.dev/)をインストールしてください。

### Windows
[Docker Desktop](https://www.docker.com/ja-jp/products/docker-desktop/)をインストールしてください。
[参考記事: Windows 11にDocker Desktopを入れる手順（令和5年最新版）](https://qiita.com/zembutsu/items/a98f6f25ef47c04893b3)

### Linux
[ドキュメント](https://docs.docker.jp/engine/installation/linux/index.html)を参考にDocker Engineを直接インストールしてください。


以下小話
## Dockerとは
[Docker](https://www.docker.com/)はコンテナ仮想化と呼ばれる仮想化技術を実現するアプリケーションです。
必要なファイルや依存を全てパッケージ化し、仮想環境を作成するため、既存の環境に影響を及ぼさず複数人で環境を共有できるメリットが存在します。

## Docker DesktopやOrbStackをインストールする理由
Docker EngineはLinuxカーネルの元に成り立っているため、Windows, macOSではLinuxをVMなどで立ち上げる必要があります。(メモリをたくさん消費する理由がこれ)
Docker Desktop, OrbStackなどを用いてLinux VMを立ち上げ、CLIからはVMを介してDocker Engineを使う方法だったり、WSLでLinux OSを立ち上げて、その中でDocker Engineをインストールして直接使う方法だったりがあります。
WSLの設定は面倒くさい(個人的意見)ので一つのソフトウェアで完結するDocker Desktop、OrbStackをインストールすることで解決しています。
[参考記事: Windows PC上でDocker Desktopを利用してLinuxの開発環境を構築するまでの仕組み](https://zenn.dev/skrikzts/articles/55f7744ee82aa9)