package MainMenu

import (
	"github.com/curtis992250/GoCCUHours/driver"
	"errors"
	"fmt"
	"github.com/tebeka/selenium"
)

func pt_proj(d driver.Driver)error{
	if _,ok := MainMenuItemCallSubMenus["Hours"].Options["HourData"];!ok{
		err := MainMenuItemCallSubMenus["Hours"].RunAction()

		if err !=nil {
			return err
		}
	}

	if ok := driver.IsGetPageSucc(d,"https://miswww1.ccu.edu.tw/pt_proj/index.php","兼任助理、臨時工工作日誌登錄系統"); !ok{
		return errors.New("Get page failed")
	}
	if err := LoginSystem(d,false);err!= nil {
		return err
	}

	if a,err := d.AlertText(); a!= ""{
		if err = d.AcceptAlert();err != nil {
			return err
		}
	}

	if ok,err := IsLogin(d,"https://miswww1.ccu.edu.tw/pt_proj/frame_stu.php");!ok{
		return err
	}

	var projNum = MainMenuItemCallSubMenus["Options"].Options["projNum"].Value
	var yy = MainMenuItemCallSubMenus["Hours"].Options["Year"].Value
	var mm = MainMenuItemCallSubMenus["Hours"].Options["Month"].Value
	//var hrs = MainMenuItemCallSubMenus["Hours"].Options["WorkHour"].Value
	var timeroot =  MainMenuItemCallSubMenus["Hours"].Options["HourData"].CallSubMenu.Root
	var workroot = MainMenuItemCallSubMenus["Contents"].Root
	var BackupWR = workroot
	for timeroot != nil {
		if workroot == nil {
			workroot = BackupWR
		}
		if err := d.Get("https://miswww1.ccu.edu.tw/pt_proj/main2.php"); err != nil {
			return err
		}
		InputCommonWorkInfo(d,yy,mm,workroot,timeroot)
		_ = driver.WebElemAction(d, driver.Click, selenium.ByXPATH, fmt.Sprintf(`//select[@name='type']/option[@value="%s"]`, projNum))
		_ = driver.WebElemAction(d, driver.SendKey,selenium.ByName,"hrs",timeroot.Value)
		_ = driver.WebElemAction(d, driver.Click, selenium.ByXPATH, "/html/body/form/center/input[1]")
		workroot = workroot.Next
		timeroot = timeroot.Next
	}
	if err := d.Get("https://miswww1.ccu.edu.tw/pt_proj/main2.php");err != nil {
		return err
	}
	if err := driver.WebElemAction(d,driver.Click,selenium.ByXPATH,"/html/body/form/center/input[2]");err!= nil{
		return err
	}
	if err := driver.WebElemAction(d,driver.Click,selenium.ByXPATH,"/html/body/center[2]/input");err!= nil{
		return err
	}
	//popup window
	if a,err := d.AlertText(); a!= ""{
		if err = d.AcceptAlert();err != nil {
			return err
		}
	}

	//produce batch number
	if err := d.Get("https://miswww1.ccu.edu.tw/pt_proj/print_sel.php");err != nil {
		return err
	}

	optionXpath := fmt.Sprintf(`//select[@name="unit_cd1"]/option[@value="%s"]`,projNum)
	if err := driver.WebElemAction(d,driver.Click,selenium.ByXPATH,optionXpath);err!= nil{
		return err
	}

	if err := driver.WebElemAction(d,driver.Click,selenium.ByXPATH,"/html/body/form/center/input[1]");err!= nil{
		return err
	}

	if err := driver.WebElemAction(d,driver.Click,selenium.ByXPATH,"/html/body/center/form/table/tbody/tr[1]/th[1]/input");err!= nil{
		return err
	}

	//submit
	//if err := driver.WebElemAction(d,driver.Click,selenium.ByXPATH,"/html/body/center/form/input[1]");err!= nil{
	//	return err
	//}

	return nil
}