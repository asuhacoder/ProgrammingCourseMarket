# ProgrammingCourseMarket

![codeql-analysis.yml](https://github.com/Asuha-a/ProgrammingCourseMarket/actions/workflows/codeql-analysis.yml/badge.svg)
![deploy-react.yml](https://github.com/Asuha-a/ProgrammingCourseMarket/actions/workflows/deploy-react.yml/badge.svg)
![eslint.yml](https://github.com/Asuha-a/ProgrammingCourseMarket/actions/workflows/eslint.yml/badge.svg)
![go-build.yml](https://github.com/Asuha-a/ProgrammingCourseMarket/actions/workflows/go-build.yml/badge.svg)
![go-test.yml](https://github.com/Asuha-a/ProgrammingCourseMarket/actions/workflows/go-test.yml/badge.svg)

## URL
[https://skhole.club](https://skhole.club)  

## サービス内容
プログラミングのコースを公開、受講できるサービスです。  

### コース作成画面

### コース受講画面

### こだわったところ

プログラムを自動実行してテストケースの入力から出力を表示できるようにして、快適なユーザ体験を追求しました。

## Network Diagram
![network diagram](./docs/skhole.drawio.svg)

## 使用技術
* Backend
  * Go
  * Gin
  * GORM
  * GRPC
* Frontend
  * TypeScript
  * React (functional component) + Recoil
  * Webpack
  * Linaria (CSS in JS)
* AWS
  * Route53, CloudFront, S3, AWS Secrets Manager, Amazon ECR, Amazon ECS, AWS Fargate, RDS
  * Terraform
* Docker & Docker Compose
* CI & CD (GitHub Actions)
* Bash Script

## 機能一覧
* 認証/認可
  * ゲストログイン機能
* コースCRUD
* レッスンCRUD
  * ドラッグ & ドロップによるレッスンの順序変更機能
  * Markdownエディタ
  * テキストエディタ (27言語対応)
  * テストケースCRUD
  * テストケースの出力計算のためのプログラム自動実行機能
  * コンパイラ指定機能 (言語とバージョンを選択)
* レッスン受講機能
  * テストケース確認機能
  * MarkdownのHTMLへの変換
  * テキストエディタ (27言語対応)
  * プログラム実行機能
  * テストケースに基づくプログラムの正誤判定機能

## DB Diagram

![dbdiagram](./docs/dbdiagram.png)  
