**db-setup**

### DBのセットアップ方法
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
