# SyadoWCSトーナメント

シングルエリミネーショントーナメント

基本機能はこちら
https://syadowcs-tournament.com

![ホーム画面](https://user-images.githubusercontent.com/8272683/130260056-aee47f83-ab7f-429a-aa0a-d07201700eb8.png)

※サービス機能はこちら(APIのみ実装)
https://backend.syadowcs-tournament.com

![スクリーンショット 2021-08-23 2 54 36](https://user-images.githubusercontent.com/8272683/130365107-d7eb1d68-2b9f-4ddb-b5e9-e9c98db8922c.png)
![エントリー](https://user-images.githubusercontent.com/8272683/130261499-fca1a9e2-f91a-4c82-b521-7d5467875c02.png)

# 機能一覧
- 基本機能
  - ユーザ登録
  - ログイン
  - ログアウト
- サービス機能
  - 大会作成・更新・削除
    - 大会名、参加人数上限設定
  - エントリー者の追加・削除
    - 参加者名

# 使用技術
- フロントエンド
  - HTML/CSS
  - JavaScript
  - Vue.js(TypeScript)
  - Bootstrap
  - ESLint
- バックエンド
  - Go
  - Echo
  - gorm
- 開発環境
  - Docker
  - CircleCI ※テスト自動化まで
- インフラ
  - AWS(EC2,RDS,VPC,Route53,ACM,ECR,ECS,Fargate) ※EC2はALB関連設定,セキュリティグループ,ターゲットグループのために使用
  - Terraform(ALB,Route53,ECR,ECS,Fargate) ※本番環境に変更分反映のために必要なものだけ作成

# 動作方法

- 開発環境動作URL
  - フロントエンド:http://localhost:8080
  - バックエンド:http://localhost:80

環境変数の設定
.envrcに以下の内容の情報を記載
```
export NODE_ENV=development
export MYSQL_USER=xxxxx
export MYSQL_PASSWORD=xxxxx
export MYSQL_ENDPOINT=xxxxx
export MYSQL_DATABASE=xxxxx
```
Dockerコンテナビルド&起動
```
docker-compose build
docker-compose up -d
```

# インフラ構成図
![スクリーンショット 2021-08-15 16 19 31](https://user-images.githubusercontent.com/8272683/129470462-4dcdb8da-9a64-433e-bc44-a7419a443ed9.png)

# ER図
![スクリーンショット 2021-08-23 0 39 38](https://user-images.githubusercontent.com/8272683/130361245-64b40b9a-60b9-42ad-976c-d026f851ec37.png)
