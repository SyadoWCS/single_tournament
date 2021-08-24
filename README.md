# SyadoWCSトーナメント

- コンセプト
大会を運営する上で配信のオーバレイの準備や管理が面倒だったという経験があったためこの課題を解決できるものを作ってみたい

最終的に目指しているゴール
- シングルエリミネーショントーナメント機能
- 配信したい試合の組み合わせをクリックすると大会の試合配信用に使用するオーバレイテンプレート
 - 対戦中の選手名
 - スコアの表示(スコアは手動入力になるけど入力して更新ボタンを押すとオーバレイに自動反映できるものにしたい)

基本機能はこちら
https://syadowcs-tournament.com

https://user-images.githubusercontent.com/8272683/130488170-85250c58-bdbe-4437-b341-64b54a375336.mov

サービス機能はこちら(APIのみ実装)
https://backend.syadowcs-tournament.com/tournament/list

https://user-images.githubusercontent.com/8272683/130489439-75568265-c54e-499a-87f1-123f82159045.mov

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

# テーブル構成図
![スクリーンショット 2021-08-23 3 20 14](https://user-images.githubusercontent.com/8272683/130365785-cb11fc5f-a158-4247-b850-e8954fa71cdd.png)
