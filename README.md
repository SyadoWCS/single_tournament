# SyadoWCSトーナメント

シングルエリミネーショントーナメント

基本機能はこちら
https://syadowcs-tournament.com

![スクリーンショット 2021-08-15 22 55 10](https://user-images.githubusercontent.com/8272683/129480913-47d6ae2a-d6a7-4a31-93ba-b705049e7a31.png)
![スクリーンショット 2021-08-15 22 55 20](https://user-images.githubusercontent.com/8272683/129480925-690b7015-a59f-4b6e-82f5-af0443b5d0a0.png)
![スクリーンショット 2021-08-15 22 55 29](https://user-images.githubusercontent.com/8272683/129481010-9ef79d04-f779-4e64-bb63-bbe079623941.png)
![スクリーンショット 2021-08-15 22 56 10](https://user-images.githubusercontent.com/8272683/129481018-3c28df49-5dac-4e1d-a565-7d44e71c2576.png)

※サービス機能はこちら(APIのみ実装)
https://backend.syadowcs-tournament.com

![スクリーンショット 2021-08-15 23 02 59](https://user-images.githubusercontent.com/8272683/129481118-df8c3235-abcc-4fe8-b97f-502933b9b2df.png)
![スクリーンショット 2021-08-15 23 03 15](https://user-images.githubusercontent.com/8272683/129481138-003d3a8f-3831-4731-bcf0-ac1e55c5cfbf.png)
![スクリーンショット 2021-08-15 23 03 34](https://user-images.githubusercontent.com/8272683/129481328-5a55a209-d779-44b1-84f2-5b35a851ce63.png)

# 機能一覧
- 基本機能
  - ユーザ登録機能
  - ログイン機能
  - ログアウト機能
- サービス機能
  - 大会作成機能
  - 大会更新機能
  - 大会削除機能

# 使用技術
- フロントエンド
  - HTML/CSS
  - JavaScript
  - Vue.js(TypeScript)
  - Bootstrap
  - ESLint
- バックエンド
  - Golang
  - Echo
- 開発環境
  - Docker
  - CircleCI ※テスト自動化まで
- インフラ
  - AWS(EC2,RDS,VPC,Route53,ACM,ECR,ECS,Fargate) ※EC2はALB,SG,TGのために使用
  - Terraform(ALB,Route53,ECR,ECS,Fargate) ※本番環境に変更分反映のために必要なものだけ作成

# インフラ構成図
![スクリーンショット 2021-08-15 16 19 31](https://user-images.githubusercontent.com/8272683/129470462-4dcdb8da-9a64-433e-bc44-a7419a443ed9.png)

# ER図

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
