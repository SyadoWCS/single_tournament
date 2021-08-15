# SyadoWCSトーナメント

シングルエリミネーショントーナメント

基本機能はこちら
https://syadowcs-tournament.com

![home](https://user-images.githubusercontent.com/8272683/129481464-a7c1ab38-aa16-469c-b08d-13d3aaece422.png)
![register](https://user-images.githubusercontent.com/8272683/129481466-70b8f803-bc38-4d4b-a27e-c404c9b78120.png)
![login](https://user-images.githubusercontent.com/8272683/129481470-03723d7b-32fe-4a69-a4fc-aff161fb0d0a.png)
![after_login](https://user-images.githubusercontent.com/8272683/129481472-3bd6bcc5-3c68-4db6-9672-da9d30614c60.png)

※サービス機能はこちら(APIのみ実装)
https://backend.syadowcs-tournament.com

![tournament_list](https://user-images.githubusercontent.com/8272683/129481515-fdd439b1-d83e-40cc-ac15-d96d0d0a38bb.png)
![tournament_create](https://user-images.githubusercontent.com/8272683/129481519-58c7dd61-afba-42c5-8346-8b50b6a89b3e.png)
![tournament_update](https://user-images.githubusercontent.com/8272683/129481521-bae2e2c7-6b7a-4489-9429-58113d479572.png)

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
