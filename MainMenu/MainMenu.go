package MainMenu

import (
	"fmt"
	"menu"
	"os"
	"time"
)

type mainMenu struct {
	menu.Menu
}

var MainMenuItemCallSubMenus = menu.MenuItemCallSubMenus{
	"UserInfo": menu.UserInfoInit(),
	"Contents": menu.ContentsInit(),
	"Options":  menu.OptionsInit(),
	"Hours": menu.HoursInit(),
	"RunOptions": menu.RunOptionsInit(),
}

func (mm *mainMenu) Show()error{

	if ok := mm.isLoad();ok{
		if err := mm.Load();err != nil{
			return err
		}
	}

	for  {
		err := mm.Menu.Show()

		if err != nil {
			fmt.Println("Error:",err)
			time.Sleep(2*time.Second)
			continue
		}

	}
}

func (mm *mainMenu) isLoad ()bool{
	fmt.Println("載入先前記錄檔嗎[Y/n]？")
	loadFlag,err := mm.GetUserInput()
	if err!=nil{
		return false
	}
	switch loadFlag[0]{
	case "N","n","No","no","NO","false","False","FALSE":
		return false
	case "Y","y","Yes","yes","YES","True","true","TRUE":
		fallthrough
	default:
		return true
	}
}

func(mm *mainMenu)Load()error{
	fmt.Println(`json檔路徑 (default: "./Salary.json")`)
	jsonPath,err := mm.GetUserInput()
	if err!=nil {
		return err
	}

	path := ""

	switch jsonPath[0]{
	case"":
		path = "./Salary.json"
	default:
		path = jsonPath[0]
	}


	if _,err := os.Stat(path);os.IsNotExist(err){
		return err
	}

	data,err := menu.LoadJsonFile(path)

	if err!=nil{
		return err
	}

	if err := mm.updateData(data);err != nil{
		return err
	}

	return nil
}

func (mm *mainMenu)updateData(data map[string]*menu.MenuOptions)error{
	for k,v := range data{
		if k == "Contents" {
			for k_con,v_con := range v.CallSubMenu.Options{
				if k_con == "work1" {
					if err := mm.Options["Contents"].CallSubMenu.OptionOperate("set","work1",v_con.Value);err != nil{
						return err
					}
				continue
				}
				if err := mm.Options["Contents"].CallSubMenu.OptionOperate("add",k_con,v_con.Value);err !=nil{
					return err
				}
			}

			if mm.Options["Contents"].CallSubMenu.Options["work1"].Value == ""{
				if err := mm.Options["Contents"].CallSubMenu.OptionOperate("rm","work1","");err != nil{
					return err
				}
			}

			continue
		}
		for k2,v2 := range v.CallSubMenu.Options{
			if k2 == "Month"|| k2 == "HourData"{
				continue
			}
			mm.Options[k].CallSubMenu.Options[k2].Value = v2.Value
		}
	}



	return nil
}