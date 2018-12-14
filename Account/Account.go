package Account

import (
	. "../G"
	. "../SQLHelper"
	"fmt"
)

func CreateAccount(username, password string) byte {
	con := Connection()

	//检测账号是否存在
	stmt, err := con.Prepare(`select * from account where username=? `)
	CheckErr(err)
	rows, err := stmt.Query(username)
	CheckErr(err)
	if rows.Next() {
		fmt.Printf("%s插入失败,已经存在\n", username)
		return createAccount_haved
	}

	//检测账号密码 是否符合格式
	//len(username)<16

	stmt, err = con.Prepare(`insert into account set username=? ,password=? `)
	CheckErr(err)
	result, err := stmt.Exec(username, password)
	CheckErr(err)
	affect, err := result.RowsAffected()
	CheckErr(err)
	if affect > 0 {
		fmt.Printf("%s插入成功\n", username)
	}

	return createAccount_succeed
}

func ChangePassword(username, password string) byte {
	con := Connection()

	stmt, err := con.Prepare(`select * from account where username=? `)
	CheckErr(err)
	rows, err := stmt.Query(username)
	CheckErr(err)
	if rows.Next() {
		fmt.Printf("%s插入失败,已经存在\n", username)
		return 2
	}
	stmt, err = con.Prepare(`insert into account set username=? ,password=? `)
	CheckErr(err)
	result, err := stmt.Exec(username, password)
	CheckErr(err)
	affect, err := result.RowsAffected()
	CheckErr(err)
	if affect > 0 {
		fmt.Printf("%s插入成功\n", username)
	}

	return 1
}
