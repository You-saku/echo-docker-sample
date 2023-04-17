# echo-docker-sample

This repository is Test app. I practice golang.

## requirement
 * git install
 * docker install
 * [golang-migrate(v4.15.2)](https://github.com/golang-migrate/migrate) install


## setup
```
1. git clone git@github.com:You-saku/echo-docker-sample.git
3. cd echo-docker
4. cp server/.env.example server/.env
5. make setup
6. make start
7. open http://localhost:80 (another terminal)
```

## Official Doc
 * [echo](https://echo.labstack.com/)

## Todo
 - いろんなアーキテクチャに挑戦

## 個人メモ

### APIの動作確認
 * postman-schemaディレクトリにpostmanでAPI叩けるようにするファイルを用意
 * importして使ってください(import方法は調べよう)

### フロントとの連携
 * localhost:3001のCORS設定をしてあります

### ファイル分割
 * [ファイル分割したときに参考になる](https://qiita.com/fetaro/items/31b02b940ce9ec579baf#%E3%83%A2%E3%82%B8%E3%83%A5%E3%83%BC%E3%83%AB%E3%83%A2%E3%83%BC%E3%83%89%E3%81%A7%E3%81%AE%E5%86%85%E9%83%A8%E3%83%91%E3%83%83%E3%82%B1%E3%83%BC%E3%82%B8%E3%81%AEimport)

### gorm
 * [公式Doc](https://gorm.io/ja_JP/docs/)
 * リレーション
    * [1対多](https://gorm.io/ja_JP/docs/has_many.html) 

### validate
 * 参考
   * [公式Doc](https://echo.labstack.com/guide/request/)
   * [ライブラリの公式Doc](https://github.com/go-playground/validator)
   * https://qiita.com/nanamen/items/c824a2c8f2e1767f90f8
   
 * やってみたい
   * https://codehex.hateblo.jp/entry/echo-context

### migration(golang-migrate)
 * 詳しくは[公式Docへ](https://github.com/golang-migrate/migrate)
 * makeコマンドでmigrationファイル作成(新規テーブル作成時に使えます)
```例)
make migration name=users // 複数形にすること
```

### テスト
#### コマンド
  * go test　　　 　　：テスト実行
  * go test -v　　　　：テスト実行（詳細な実行結果出力）
  * go test -cover　　：テスト実行＋コードカバレッジ
  * go test -cover -v ：テスト実行＋コードカバレッジ（詳細な実行結果出力）

このリポジトリではディレクトリ構成の関係でコマンドは以下の通り
```
docker-compose exec go go test ./test
```
