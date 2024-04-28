**users**

ユーザーテーブル。認証情報の管理を含む。

**テーブル：users**
| 物理名       |  主   | 型       | 必須  | 桁数  | 一意  | hash  |
| ------------ | :---: | -------- | :---: | :---: | :---: | :---: |
| id           |  〇   | int      |  〇   |       |  〇   |       |
| user_id      |       | varchar  |  〇   |  50   |  〇   |       |
| name         |       | varchar  |       |  50   |  〇   |       |
| password     |       | varchar  |  〇   |  500  |       |  〇   |
| access_token |       | varchar  |       |  500  |       |  〇   |
| created_at   |       | datetime |       |       |       |       |
| updated_at   |       | datetime |       |       |       |       |
| deleted_at   |       | datetime |       |       |       |       |
