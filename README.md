# SyadoWCSトーナメント

シングルエリミネーショントーナメント

# 使用技術
- インフラ
  - AWS(EC2,RDS,VPC,Route53,ACM,ECR,ECS,Fargate) ※EC2はALB,SG,TGのために使用
  - Terraform(ALB,Route53,ECR,ECS,Fargate) ※本番環境に変更分反映のために必要なものだけ作成
- 開発環境
  - Docker
  - CircleCI ※テスト自動化まで
- フロントエンド
  - HTML/CSS
  - JavaScript
  - Vue.js(TypeScript)
  - Bootstrap
  - ESLint
- バックエンド
  - Golang
  - Echo

# インフラ構成図
![スクリーンショット 2021-08-15 16 19 31](https://user-images.githubusercontent.com/8272683/129470462-4dcdb8da-9a64-433e-bc44-a7419a443ed9.png)

# ER図

# 機能一覧
- 基本機能
  - ユーザ登録機能
  - ログイン機能
  - ログアウト機能
- サービス機能
  - 大会作成機能
  - 大会更新機能
  - 大会削除機能
