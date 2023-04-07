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
サービス内でコーディングやプログラムのコンパイル、実行ができます。

### ターゲット
プログラマー、特に新しい技術を楽に習得したい人。

### 課題

プログラマーとして初学者を抜け出したが、今勉強していることもかつてProgateで学習したような手取り足取り教えてくれる学習コースが欲しい。

### 課題の解決策

Progateのような学習コースを作って公開できるようにした。これにより講座の通りに進めるだけで知識が見についていく。Progateによって舗装された道の延長を提供する。

### コース受講画面

![ProgrammingCourseMarket - Google Chrome 2023-04-07 16-36-50](https://user-images.githubusercontent.com/30449505/230566565-366cc75b-35f2-4325-b25c-fb77b5a6e766.gif)

### コース作成画面

![ProgrammingCourseMarket - Google Chrome 2023-04-07 16-04-13](https://user-images.githubusercontent.com/30449505/230563637-ae9a252b-f4b3-4a55-a44b-9a77b50f734b.gif)

### こだわったところ

プログラムを自動実行してテストケースの入力から出力を表示できるようにして、快適なユーザ体験を追求した。

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
  * Route53, CloudFront, S3, AWS Secrets Manager, Amazon ECR, Amazon ECS, AWS Fargate, RDS, Amazon CloudWatch, AWS Certificate Manager, VPC Endpoints, ALB
  * Terraform
* Docker & Docker Compose
* CI & CD (GitHub Actions)
* Bash Script
* Makefile

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
