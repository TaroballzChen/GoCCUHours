package MainMenu

import (
	"github.com/curtis992250/GoCCUHours/driver"
	"errors"
	"fmt"
	"github.com/tebeka/selenium"
	"strings"
)

func parttime(d driver.Driver)error{
	if _,ok := MainMenuItemCallSubMenus["Hours"].Options["HourData"];!ok{
		err := MainMenuItemCallSubMenus["Hours"].RunAction()

		if err !=nil {
			return err
		}
	}

	if ok := driver.IsGetPageSucc(d,"https://miswww1.ccu.edu.tw/parttime/index.php","學習暨勞僱時數登錄系統"); !ok{
		return errors.New("Get page failed")
	}

	if err := LoginSystem(d,true);err!= nil {
		return err
	}

	if a,err := d.AlertText(); a!= ""{
		if err = d.AcceptAlert();err != nil {
			return err
		}
	}

	if ok,err := IsLogin(d,"https://miswww1.ccu.edu.tw/parttime/frame_stu.php?type=0");!ok{
		return err
	}
	currentwindow,err := d.CurrentWindowHandle()
	if err != nil {
		return err
	}

	var Employer = MainMenuItemCallSubMenus["Options"].Options["Employer"].Value
	var yy = MainMenuItemCallSubMenus["Hours"].Options["Year"].Value
	var mm = MainMenuItemCallSubMenus["Hours"].Options["Month"].Value
	var timeroot =  MainMenuItemCallSubMenus["Hours"].Options["HourData"].CallSubMenu.Root
	var workroot = MainMenuItemCallSubMenus["Contents"].Root
	var BackupWR = workroot
	for timeroot != nil {

		timeZone :=  strings.Split(timeroot.Value,",")

		for _,v := range timeZone {
			v = strings.TrimSpace(v)
			if v == "" {
				continue
			}

			if workroot == nil {
				workroot = BackupWR
			}

			if err := d.Get("https://miswww1.ccu.edu.tw/parttime/main2.php"); err != nil {
				return err
			}

			InputCommonWorkInfo(d, yy, mm, workroot, timeroot)

			_ = driver.WebElemAction(d, driver.Click, selenium.ByXPATH, fmt.Sprintf(`//select[@name='type']/option[@value="%s"]`, Employer))

			hourZone := strings.Split(v,"-")

			_ = driver.WebElemAction(d, driver.SendKey,selenium.ByName,"shour",strings.TrimSpace(hourZone[0]))

			_ = driver.WebElemAction(d, driver.SendKey,selenium.ByName,"ehour",strings.TrimSpace(hourZone[1]))

			_ = driver.WebElemAction(d, driver.Click, selenium.ByXPATH, "/html/body/form/center/input[2]")

			//switch window for fast fill
			err := d.SwitchWindow(currentwindow)
			if err != nil {
				return err
			}

			workroot = workroot.Next
		}

		timeroot = timeroot.Next
	}


	if err := d.Get("https://miswww1.ccu.edu.tw/parttime/control2.php");err != nil {
		return err
	}
	if err := d.Get("https://miswww1.ccu.edu.tw/parttime/main2.php");err != nil {
		return err
	}
	if err := driver.WebElemAction(d,driver.Click,selenium.ByXPATH,"/html/body/form/center/input[3]");err!= nil{
		return err
	}
	//fill success

	if err := d.Get("https://miswww1.ccu.edu.tw/parttime/print_sel.php");err != nil {
		return err
	}

	EmployerCode := MainMenuItemCallSubMenus["Options"].Options["EmployerCode"].Value
	optionXpath := fmt.Sprintf(`//select[@name="unit_cd1"]/option[@value="%s"]`,EmployerCode)
	if err := driver.WebElemAction(d,driver.Click,selenium.ByXPATH,optionXpath);err!= nil{
		return err
	}

	if err := driver.WebElemAction(d,driver.Click,selenium.ByXPATH,"/html/body/form/center/table/tbody/tr/td/input[2]");err != nil {
		return err
	}

	//findHoursSuccess

	if err := driver.WebElemAction(d,driver.Click,selenium.ByXPATH,"/html/body/center/form/table[1]/tbody/tr[1]/th[1]/input");err != nil {
		return err
	}

	HourMoney := MainMenuItemCallSubMenus["Options"].Options["HourMoney"].Value
	if err := driver.WebElemAction(d,driver.SendKey,selenium.ByName,"hour_money",HourMoney);err != nil{
		return err
	}

	sutype := MainMenuItemCallSubMenus["Options"].Options["sutype"].Value
	sutypeXpath := fmt.Sprintf(`//*[@id="sutype" and @value="%s"]`,sutype)
	if err := driver.WebElemAction(d,driver.Click,selenium.ByXPATH,sutypeXpath);err!= nil{
		return err
	}

	iswork := MainMenuItemCallSubMenus["Options"].Options["iswork"].Value
	isworkXpath := fmt.Sprintf(`//*[@id="iswork" and @value="%s"]`,iswork)
	if err := driver.WebElemAction(d,driver.Click,selenium.ByXPATH,isworkXpath);err!= nil{
		return err
	}

	emp_type:= MainMenuItemCallSubMenus["Options"].Options["emp_type"].Value
	emp_typeXpath := fmt.Sprintf(`//*[@id="emp_type" and @value="%s"]`,emp_type)
	if err := driver.WebElemAction(d,driver.Click,selenium.ByXPATH,emp_typeXpath);err!= nil{
		return err
	}

	//agreethis
	if err := driver.WebElemAction(d,driver.Click,selenium.ByXPATH,`//*[@id="agreethis" and @value="1"]`);err!= nil{
		return err
	}

	//submit
	//if err := driver.WebElemAction(d,driver.Click,selenium.ByXPATH,`//*[@id="go_check"]`);err!= nil{
	//	return err
	//}

	return nil
}
