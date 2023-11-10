# misskey-abuse-user-report-notifier
Misskeyサーバーに対する通報をDiscordサーバーに流すプログラム。  
プログラム自体は1回実行すると落ちるので、crontab等を使って定期実行する想定。

## 環境

### 開発時の設定

<https://pre-commit.com/> の手順に従って `pre-commit` をインストールする。  
これにより、[.pre-commit-config.yaml](.pre-commit-config.yaml)の設定に基づいて、コミット時にクレデンシャルが含まれていないかの検査が行われるようになる。

### 立ち上げ

#### 共通

1. `cp .env.example .env` を実行して `.env` を作成します。
2. `.env` 内のTODOコメントに従って設定します。

#### 開発環境

```sh
export TAG_NAME=$(git symbolic-ref --short HEAD | sed -e "s:/:-:g" | sed -e "s/^main$/latest/g")
docker compose -f docker-compose.yml -f dev.base.docker-compose.yml -f dev.docker-compose.yml build
docker compose -f docker-compose.yml -f dev.base.docker-compose.yml -f dev.docker-compose.yml watch
```

#### 本番環境
PostgreSQLのDBを別途用意した上で以下を実行します。

```sh
export TAG_NAME=$(git symbolic-ref --short HEAD | sed -e "s:/:-:g" | sed -e "s/^main$/latest/g")
docker compose up --build
```
