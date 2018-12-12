package SQLHelper

import (
	."../G"
	"fmt"
)

func InitDBTable() {
	conn := Connection()
	defer conn.Close()
	rows, err := conn.Query("show DATABASES")
	CheckErr(err)
	fmt.Println(rows.Columns())
	/*
	sqlStr := fmt.Sprintf("show databases like '%s'", Dbname)
	rows, err := conn.Query(sqlStr)
	CheckErr(err)
	cols, err := rows.Columns()
	CheckErr(err)
	fmt.Println(cols)
	fmt.Println(len(cols))
//
	//sqlStr = fmt.Sprintf("create database if not exists %s", Dbname)
	//result, err := conn.Query(sqlStr)
	//CheckErr(err)
	//fmt.Println(result)
	*/
}
