**go-gin-auth**

Golang + Gin + Auth(JWT)

## はじめに
当プロジェクトはGolangとGinを利用したログイン認証及びアクセス認可です。JWTを利用したシンプルな実装で構成しています。

## 構成
```
├── main.go
├── common
|    ├── connect
|    |    └── db.go
|    ├── sessions
|　　　　  └── sessions.go
|    └── response
|　　　　  └── response.go
├── config
|    ├── app.ini
|    └── config.go
└── routers
　　　├── routers.go
　　　└── api
　　　　　　├── api.go
　　　　　　├── auths
　　　　　　|　　　├── routes.go
　　　　　　|　　　├── app_auth_controllers.go 
　　　　　　|　　　├── model.go
　　　　　　|　　　├── middleware.go
　　　　　　|　　　├── utils.go
　　　　　　|　　　└── validators.go
　　　　　  └── books
                  ├── routes.go
                  ├── book_controllers.go 
                  ├── model.go
                  ├── middleware.go
                  └── validators.go
```

## ドキュメント

### データ定義
| テーブル名 | 概要 |
|-----------|------------|
| [users](https://github.com/chalmeal/go-gin-auth/blob/master/.doc/data/users.md) | ユーザーテーブル、認証情報の管理を含む|
| [books](https://github.com/chalmeal/go-gin-auth/blob/master/.doc/data/books.md) | 書籍マスタテーブル |

### 仕様
| 書名 | 概要 |
|-----------|------------|
| [認証認可](https://github.com/chalmeal/go-gin-auth/blob/master/.doc/method/%E8%AA%8D%E8%A8%BC%E8%AA%8D%E5%8F%AF.md) | ユーザー認証をする機能について |

## セットアップ
### DB
* DBの環境は以下を想定します。
  * MySQL
  * GORM
* create schemaのみ行う必要があります。[DDL](.db/setup/ddl-create-chema.sql)
* [app.ini](config/app.ini)に対してDB接続情報を定義してください。
* テーブルはGORMが提供するAutoMigrateを利用します。
  * 各テーブルは初回API実行時に生成されます。
  * テストレコードを追加したい場合は[dml-insert-into_users.sql](https://github.com/chalmeal/go-gin-auth/blob/master/.db/setup/dml-insert-into_users.sql)及び[dml-insert-into_books.sql](https://github.com/chalmeal/go-gin-auth/blob/master/.db/setup/dml-insert-into_books.sql)を参考に実行してください。
  * 本プロジェクトで定義されている詳細なデータ定義に関しては、各ドメインのdocを参照してください。
    * ユーザー：[users](https://github.com/chalmeal/go-gin-auth/blob/master/.doc/auth/App/App-Auth.md)
    * 書籍マスタ：[books](https://github.com/chalmeal/go-gin-auth/blob/master/.doc/books/books.md)
   
## アプリケーションスタート
アプリケーションのスタートはデバッガを推奨しています。

Run and Debugの`Run go-gin-auth`から実行してください。

## その他
**Qiita**
* [【Golang + Gin】JSON Web Token(JWT)でサクっとREST-API実装](https://qiita.com/chalmeal/items/740bf98c64a9a341da54#%E6%88%90%E6%9E%9C%E7%89%A9)
* [【Golang + Gin】JWT + sessionsでログインセッション管理]()
