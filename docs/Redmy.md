# 学習まとめ　 ctr + k v

- 初期化

  ```
  go mod init github.com/qqks0418/golang_study
  ```

- GIN インストール (フレームワーク)　例：[参考](https://deku.posstree.com/golang/gin/start/)

  ```
  go get -u github.com/gin-gonic/gin
  ```

- sqlboiler インストール (ORM) ※ カラム名スネークケース

  ドライバー用

  ```
  go get -u github.com/go-sql-driver/mysql
  ```

  コマンド用

  ```
  go install github.com/volatiletech/sqlboiler/v4@latest
  ```

  ```
  go insなどll github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest

  go get github.com/volatiletech/sqlboiler/v4/boil

  sqlboiler --version
  ```

  モデル作成

  ```
  sqlboiler mysql
  sqlboiler mysql -t form --wipe
  sqlboiler mysql -t boil --wipe --add-global-variants  ※ AllGなどを使う場合
  (tomlファイルに書き込めば「add-global-variants」オプションなど付けなくてもよい)
  ```

  Docker(MySQL) 　例：[参考](https://zenn.dev/jojojo/articles/f1223bb06cf5be)
  SQLBoiler 　例：[参考](https://scrapbox.io/javememo/Go%E3%81%AB%E3%81%8A%E3%81%91%E3%82%8BORM%E3%81%A8%E3%80%81SQLBoiler%E5%85%A5%E9%96%80%E3%83%9E%E3%83%8B%E3%83%A5%E3%82%A2%E3%83%AB)

  ```
  whitelistにはコード生成対象のテーブルを明示的に記載
  デバッグ出力: boil.DebugMode = true
  ```

  SQL 文を自作する。結果を構造体に入れられる

  ```
  type Summary struct {
  	Max int     `boil:"max"`
  	Avg float32 `boil:"avg"`
  }
  func MakeSummary(ctx context.Context, db *sql.DB) (*Summary, error) {
  	var s *Summary
  	sql := "SELECT SUM(cnt) AS sum, AVG(cnt) AS avg FROM " +
  		" ( SELECT author_id, COUNT(*) AS cnt FROM articles GROUP BY author_id ) AS s"
  	err := queries.Raw(sql).Bind(ctx, db, &s) // ここ
  	if err != nil { return nil, err }
  	return s, nil
  }
  ```

  null パッケージ

  ```
  golangには値にnullが存在しないので、そのままではDBにnullを書き込めない
  nullがあり得るstring（not null制約の無いカラムとか）は、`null.String`型になる
  ```

**hoge**
~~テキスト~~

| 左  | 中央 | 右  |
| --- | ---- | --- |
| td  | td   | td  |
