package menu

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func EncodeJson(data map[string]*MenuOptions)([]byte, error){
	b ,err := json.MarshalIndent(data,""," ")
	if err != nil {
		return nil,err
	}
	return b,nil
}

func WriteJsonFile(path string, data map[string]*MenuOptions)error{
	Write,err := EncodeJson(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path,Write,os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func LoadJsonFile(path string)(map[string]*MenuOptions, error){
	encode,err := ioutil.ReadFile(path)
	if err != nil {
		return nil,err
	}
	data := make(map[string]*MenuOptions)

	err = json.Unmarshal(encode,&data)
	if err != nil {
		return nil,err
	}
	return data,nil
}


