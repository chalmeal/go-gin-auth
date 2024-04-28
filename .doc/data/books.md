**books**

書籍マスタテーブル。

**テーブル：books**
| 物理名 | 主 | 型 | 必須 | 桁数 | 一意 |
|-----------|:-----------:|------------|:------------:|:------------:|:------------:|
| id | 〇 | int | 〇 |  | 〇 |
| book_id |  | varchar | 〇 | 50 | 〇 |
| name |  | varchar | 〇 | 50 |  |
| author |  | varchar |  | 50 |  |
| created_at |  | datetime |  |  |  |
| updated_at |  | datetime |  |  |  |
| deleted_at |  | datetime |  |  |  |
