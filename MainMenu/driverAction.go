package MainMenu

import (
	"github.com/curtis992250/GoCCUHours/driver"
	"errors"
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/curtis992250/GoCCUHours/menu"
)

func LoginSystem(d driver.Driver,part_time bool)error{
	userinfo := MainMenuItemCallSubMenus["UserInfo"]
	username := userinfo.Options["UserName"].Value
	password := userinfo.Options["PassWord"].Value
	job := userinfo.Options["JobType"].Value

	if err := driver.WebElemAction(d,driver.SendKey,selenium.ByName,"staff_cd",username);err != nil{
		return err
	}

	if err := driver.WebElemAction(d,driver.SendKey,selenium.ByName,"passwd",password);err != nil{
		return err
	}
	switch part_time {
	case false:
		xpath := fmt.Sprintf("//select[@name='proj_type']/option[@value=%s]",job)
		if err := driver.WebElemAction(d,driver.Click,selenium.ByXPATH,xpath);err != nil{
			return err
		}
		fallthrough
	case true:
		if err := driver.WebElemAction(d,driver.Click,selenium.ByXPATH,"/html/body/center/form/input[1]");err != nil{
			return err
		}
	}
	return nil
	}



func IsLogin(d driver.Driver, succUrl string)(bool, error) {
	currentUrl,err := d.CurrentURL()
	if err != nil||currentUrl!=succUrl {
		return false,errors.New("Login Failed!!")
	}
	fmt.Println("Login Success!!")
	return true,nil
}

func InputCommonWorkInfo(d driver.Driver,yy,mm string, workinfo,timeinfo *menu.MenuOptions){
	_ = driver.WebElemAction(d,driver.SendKey,selenium.ByName,"yy",yy)
	_ = driver.WebElemAction(d,driver.SendKey,selenium.ByName,"mm",mm)
	_ = driver.WebElemAction(d,driver.SendKey,selenium.ByName,"dd",timeinfo.OptName[1:])
	_ = driver.WebElemAction(d,driver.SendKey,selenium.ByName,"workin",workinfo.Value)
}
