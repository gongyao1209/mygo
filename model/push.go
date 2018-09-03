package model

import (
	"database/sql"
	"fmt"
)

type Push struct {
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

func GetPush(Db *sql.DB, id int) Push {
	var push = Push{}

	stmt,_ := Db.Prepare("SELECT * FROM `push` AS `p` WHERE `p`.`id` = ? LIMIT 1")
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Printf("SELECT ERROR %s\n", err.Error())
		return push
	}

	for rows.Next() {
		err = rows.Scan(&push.Id, &push.Type, &push.Status, &push.BoxId, &push.TimeStamp, &push.Time, &push.CreateUserId, &push.CreateTime, &push.UpdateUserId, &push.UpdateTime)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}

		fmt.Println("get data, id:", push.Id, ", box_id:", push.BoxId, "\n")
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}

	return push
}