package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func CopyFile(source, dest string) bool {
	if source == "" || dest == "" {
		log.Println("source or dest is null")
		return false
	}

	source_open, err := os.Open(source)

	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer source_open.Close()

	dest_open, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 644)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	defer dest_open.Close()

	_, copy_err := io.Copy(dest_open, source_open)
	if copy_err != nil {
		log.Println(copy_err.Error())
		return false
	} else {
		return true
	}
}

func Merge(res1, res2 map[string](map[string]string)) map[string](map[string]string){
	for k,v := range res2{
		if _, ok := res1[k]; ok {
			if len(v["password"])>0 && len(res1[k]["password"])==0{
				res1[k]["password"] = v["password"]
			}
		}else{
			res1[k] = v
		}
	}

	return res1
}

func FormatOutput(total_res map[string](map[string]string), pwd_db string) error{
	for k,v := range total_res{
		fmt.Printf("====================\n")
		fmt.Printf("Url: %s\nUsername: %s\nPassword:%s\n\n", k, v["username"], v["password"])
	}

	fmt.Printf("\nTotal Auth: %d", len(total_res))
	err := os.Remove(pwd_db)
	if err != nil{
		return err
	}
	return nil
}
