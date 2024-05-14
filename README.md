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

## その他
**Qiita**
* [【Golang + Gin】JSON Web Token(JWT)でサクっとREST-API実装](https://qiita.com/chalmeal/items/740bf98c64a9a341da54#%E6%88%90%E6%9E%9C%E7%89%A9)
* [【Golang(Gin)】JWT + sessionsでログインセッション管理]()
