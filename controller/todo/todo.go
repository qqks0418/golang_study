package todo

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qqks0418/golang_study/entity"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func TodoApi() {

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/test_db?parseTime=true")

	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}

	boil.SetDB(db)
	//defer db.Close()
	//boil.DebugMode = true

	ctx := context.Background()
	//connectOnly()
	//sqlboilerInsert(ctx)
	sqlboilerSelect(ctx)
	//sqlboilerDelete(ctx)
	//sqlboilerUpdate(ctx)
}

// 更新
func sqlboilerUpdate(ctx context.Context) {

	t := entity.Todo{
		ID:   8,
		TodoName: "aaa",
	}

	// t.UpdateG(ctx, boil.Infer())
	// boil.Whitelist(model.TodoColumns.TodoName)
	if _, err := t.UpdateG(ctx, boil.Infer()); err != nil {
		fmt.Println("[NG] TODOテーブル更新")
		return
	}

	/*
	if _, err := t.UpsertG(ctx, boil.Infer(), boil.Infer()); err != nil {
		fmt.Println("[NG] TODOテーブル更新")
		return
	}
	*/
	//t.UpsertG(ctx, boil.Infer(), boil.Infer())

	/*
	upCount, err := t.UpsertG(ctx, boil.Infer(), boil.Infer())

	if err != nil {
		//loger.Error(err.Error())
		fmt.Println("UpsertNG")
		return
	}
	fmt.Println(upCount)
	*/
	
	fmt.Println("[OK] TODOテーブル更新")
}

// 削除
func sqlboilerDelete(ctx context.Context) {

	var t entity.Todo
	t.ID = 5

	if _, err := t.DeleteG(ctx); err != nil {
		fmt.Println("[NG] TODOテーブル削除")
		return
	}

	fmt.Println("[OK] TODOテーブル削除")
}

// 登録
func sqlboilerInsert(ctx context.Context) {

	// トランザクション?
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		//loger.Error(err.Error())
		fmt.Println("NG")
		return
	}

	var t entity.Todo
	t.TodoName = "川中島合戦"
	t.TodoStatus = 0

	// TODOテーブルに登録
	if err := t.Insert(ctx, tx, boil.Infer()); err != nil {
		fmt.Println("[NG] TODOテーブル登録")
		tx.Rollback()
		return
	}

	var u entity.User
	//u.ID = 1
	u.UserName = "武田信玄"
	u.UserJob = "甲斐の虎あいうえおかきくけこ"
	u.TodoID = "3"

	// userテーブルに登録
	if err := u.Insert(ctx, tx, boil.Infer()); err != nil {
		fmt.Println("[NG] userテーブル登録")
		tx.Rollback()
		return
	}

	// コミット
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		panic(err)
	}

	//_ = t.InsertG(ctx, boil.Infer())
	fmt.Println("[OK] 登録完了")
}

func sqlboilerSelect(ctx context.Context) {

	fmt.Println("===========")
	fmt.Println("ソート")
	fmt.Println("===========")
	todos, err1 := entity.Todos(qm.OrderBy("id desc")).AllG(ctx)

	if err1 != nil {
		// ここではエラーを返さない
		log.Fatal(err1)
	}

	for n, v := range todos {
		fmt.Println(n, v.TodoName)
	}
	fmt.Println("取得件数: ", len(todos))

	fmt.Println("")
	fmt.Println("===========")
	fmt.Println("クエリメンバーによるSELECT")
	fmt.Println("===========")
	todoLists, _ := entity.Todos(
		//qm.Where("todo_name = ?", "大阪夏の陣"),
		//qm.Or("todo_name like ?", `大阪%`),
		//qm.WhereIn("todo_name", []string{"大阪夏の陣", "朝鮮出兵"}),
		//qm.AndIn("todo_name", []string{"大阪夏の陣", "朝鮮出兵"}),
		qm.AndIn("todo_name in ?", "大阪夏の陣", "朝鮮出兵"),
		//qm.SQL("SELECT * from todo where id=10"),
	).AllG(ctx)

	for i, v := range todoLists {
		fmt.Println(i, v)
	}

	fmt.Println("")
	fmt.Println("===========")
	fmt.Println("結合")
	fmt.Println("===========")
	type userTodo struct {
		entity.User `boil:",bind"`
		entity.Todo `boil:",bind"`
	}

	var uts[] userTodo
	entity.Todos(
		qm.Select("user.*, todo.*"), 
		qm.LeftOuterJoin("user on user.todo_id = todo.id"),
	).BindG(ctx, &uts)

	for _, ut := range uts {
		fmt.Printf("ut = %+v\n", ut.UserName + "_" + ut.TodoName)
	}
}


/*
func connectOnly() {
	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/test_db")
	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}
	defer db.Close()

	// 実際に接続する
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("データベース接続完了!")
	}
}
*/