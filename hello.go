package main

import (
	"database/sql"
	"fmt"

	_"github.com/go-sql-driver/mysql"
)

func main() {

	var err error

	dbw := DbWorker{
		Dsn:"root:S12p_w99Q@tcp(139.198.5.192:3306)/xiangxin0828?charset=utf8",
	}

	dbw.Db, err = sql.Open("mysql", dbw.Dsn)
	if err != nil {
		fmt.Printf("error")
	}

	dbw.queryData(32001)
}

type DbWorker struct {
	Dsn	string
	Db	*sql.DB
	Push push
}

type push struct {
	Id int
	Type int
	Status int
	BoxId int
	TimeStamp int
	Time string
	CreateUserId int
	CreateTime string
	UpdateUserId int
	UpdateTime string
}


func (dbw *DbWorker) insertData() {
	//stmt, _ = dbw.Db.Prepare("INSERT INTO ")
}

func (dbw *DbWorker) queryData(id int) {

	stmt,_ := dbw.Db.Prepare("SELECT * FROM `push` AS `p` WHERE `p`.`id` = ? LIMIT 1")
	defer stmt.Close()

	dbw.Push = push{}

	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Printf("SELECT ERROR %s\n", err.Error())
		return
	}

	for rows.Next() {
		err = rows.Scan(&dbw.Push.Id, &dbw.Push.Type, &dbw.Push.Status, &dbw.Push.BoxId, &dbw.Push.TimeStamp, &dbw.Push.Time, &dbw.Push.CreateUserId, &dbw.Push.CreateTime, &dbw.Push.UpdateUserId, &dbw.Push.UpdateTime)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}

		fmt.Println("get data, id:", dbw.Push.Id, ", box_id:", dbw.Push.BoxId, "\n")
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func (dbw *DbWorker) changeStatus(id int, status int) {
	stmt, _ := dbw.Db.Prepare("UPDATE `push` SET `status` = ? WHERE `id` = ?")
	defer stmt.Close()

	ret, err := stmt.Exec(status, id)
	if err != nil {
		fmt.Printf("Update data error: %v\n", err.Error())
		return
	}

	if affected, err := ret.RowsAffected(); err == nil {
		fmt.Println("RowsAffected:", affected)
	}
}
