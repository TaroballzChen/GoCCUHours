package MainMenu

import (
	"github.com/curtis992250/GoCCUHours/driver"
	"fmt"
	"github.com/curtis992250/GoCCUHours/menu"
	"strconv"
)

const (
	Parttime = "1"
	Pt_proj = "2"
)

func InitMainMenu() *mainMenu {
	m := new(mainMenu)
	m.Options = map[string]*menu.MenuOptions{
		"UserInfo":&menu.MenuOptions{
			ID:          1,
			OptName:     "UserInfo",
			Text:        "帳號、密碼、身分",
			CallSubMenu: MainMenuItemCallSubMenus["UserInfo"],
		},
		"Contents":&menu.MenuOptions{
			ID:          2,
			OptName:     "Contents",
			Text:        "工作項目",
			CallSubMenu: MainMenuItemCallSubMenus["Contents"],
		},
		"Options":&menu.MenuOptions{
			ID:          3,
			OptName:     "Options",
			Text:        "產生批號設置",
			CallSubMenu: MainMenuItemCallSubMenus["Options"],
		},
		"Hours":&menu.MenuOptions{
			ID:          4,
			OptName:     "Hours",
			Text:        "日期、時數設置",
			CallSubMenu: MainMenuItemCallSubMenus["Hours"],
		},
		"RunOptions":&menu.MenuOptions{
			ID:          5,
			OptName:     "RunOptions",
			Text:        "記錄儲存設置",
			CallSubMenu: MainMenuItemCallSubMenus["RunOptions"],
		},
	}


	m.AllowAction = map[string]menu.ActionID{
		fmt.Sprintf("1-%d",len(m.Options)):100,
		//"exit": menu.Exit,
		"run": menu.Run,
	}
	for i:=1;i<=len(m.Options);i++{
		m.AllowAction[strconv.Itoa(i)] = menu.Number
	}


	m.RunAction = func() error {

		writejsonflag := MainMenuItemCallSubMenus["RunOptions"].Options["WriteFile"].Value
		switch writejsonflag{
		case "Y","y","Yes","yes","YES","True","true","TRUE":
			filePath := MainMenuItemCallSubMenus["RunOptions"].Options["WriteFilePath"].Value
			if err := menu.WriteJsonFile(filePath,m.Options);err!=nil{
				return err
			}
		case "N","n","No","no","NO","false","False","FALSE":
		}

		return mmRunAction()
	}

	return m
}


func mmRunAction() error {
	chromeDrvOpt := driver.DefaultChromeDrvOpt()
	Chromedrv,err := chromeDrvOpt.NewDriver()
	if err != nil {
		return err
	}

	switch chosen:= MainMenuItemCallSubMenus["Hours"].Options["Action"].Value; chosen{
	case Parttime:
		if err := parttime(Chromedrv);err != nil {
			return err
		}
	case Pt_proj:
		if err := pt_proj(Chromedrv);err != nil {
			return err
		}

	}
	return nil


}