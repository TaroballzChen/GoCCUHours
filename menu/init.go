package menu

import (
	"errors"
	"fmt"
)

func UserInfoInit()*Menu {
	menu := NewMenu()
	menu.Options = map[string]*MenuOptions{
		"UserName": &MenuOptions{
			ID:       1,
			OptName:  "UserName",
			Text:     "[必填] 帳號",
		},
		"PassWord": &MenuOptions{
			ID:       2,
			OptName:  "PassWord",
			Text:     "[必填] 密碼",
		},
		"JobType": &MenuOptions{
			ID:       3,
			OptName:  "JobType",
			Value:    "3",
			Text:     "[日誌必填] 1主持人 2專任 3兼任 4臨時工",
		},
	}

	menu.AllowAction = allowAct{
		"set":  Set,
		"exit": Exit,
	}
	return menu
}

func ContentsInit()*Menu {
	menu := NewMenu()
	menu.Options = make(map[string]*MenuOptions,5)
	menu.Options["work1"] = NewMenuOption(1,"work1","")
	menu.Root = menu.Options["work1"]
	menu.AllowAction = allowAct{
		"add":  Add,
		"rm":   Remove,
		"set":  Set,
		"exit": Exit,
	}
	return menu
}

func OptionsInit()*Menu {
	menu := NewMenu()
	menu.Options = map[string]*MenuOptions{
		"HourMoney":&MenuOptions{
			ID:          1,
			OptName:     "HourMoney",
			Value:       "150",
			Text:        "[時數必填] 時薪 (default:150)",
		},
		"Employer":&MenuOptions{
			ID:          2,
			OptName:     "Employer",
			Text:        "[時數必填] 參考departmentCode 格式: Z135奈米生物檢測科技研究中心 ",
		},
		"EmployerCode":&MenuOptions{
			ID:          3,
			OptName:     "EmployerCode",
			Text:        "[時數必填] 參考departmentCode 格式: Z135",
		},
		"sutype":&MenuOptions{
			ID:          4,
			OptName:     "sutype",
			Value:       "1",
			Text:        "[時數必填] 你的身分是：(1-6)，1為僑生、外籍生、一般學生 (其他需查詢)",
		},
		"iswork":&MenuOptions{
			ID:          5,
			OptName:     "iswork",
			Value:       "0",
			Text:        "[時數必填] 在校外有無專職工作：(0 or 1) (default: 0) ",
		},
		"emp_type":&MenuOptions{
			ID:          6,
			OptName:     "emp_type",
			Value:       "1",
			Text:        "[時數必填] 您的學習暨勞僱類型為：(1-3)，1為行政學習助理 (其他需查詢)",
		},
		"projNum":&MenuOptions{
			ID:          7,
			OptName:     "projNum",
			Text:        "[日誌必填] 計畫代碼 Ex: 105-00018",
		},
	}
	menu.AllowAction = allowAct{
		"set":  Set,
		"exit": Exit,
	}
	return menu
}

func HoursInit()*Menu {
	menu := NewMenu()
	menu.Options = map[string]*MenuOptions{
		"WorkHour":&MenuOptions{
			ID:          1,
			OptName:     "WorkHour",
			Text:        "[必填] 日誌：每天多少小時；時數：總時數",
		},
		"Year":&MenuOptions{
			ID:          2,
			OptName:     "Year",
			Text:        "[必填] 欲填寫的年份(民國)",
		},
		"Month":&MenuOptions{
			ID:          3,
			OptName:     "Month",
			Text:        "[必填] 欲填寫的月份",
		},
		"ExcludeDays":&MenuOptions{
			ID:          4,
			OptName:    "ExcludeDays",
			Value:      "0,6",
			Text:       "[選填] 欲扣除禮拜幾不填(0為日；6為六)",
		},
		"Action": &MenuOptions{
			ID:          5,
			OptName:    "Action",
			Value:      "",
			Text:      "[必填] 1).時數登錄 ； 2).日誌登錄",
		},
	}
	menu.AllowAction = allowAct{
		"run":  Run,
		"set":  Set,
		"exit": Exit,
	}

	menu.RunAction = func() error {
		if menu.Options["WorkHour"].Value == "" || menu.Options["Year"].Value == "" || menu.Options["Month"].Value == "" {
			return errors.New("WorkHour,Year,Month Value have to fill!!")
		}
		if _,ok := menu.Options["HourData"];ok {
			delete(menu.Options,"HourData")
		}


		hm,err := HourDataInit(menu.Options)
		if err != nil {
			return err
		}

		menu.Options["HourData"] = &MenuOptions{
			ID:          6,
			OptName:     "HourData",
			Text:        "時數資料",
			CallSubMenu: hm,
		}

		menu.AllowAction["6"] = HourData
		return nil
	}
	return menu
}


func RunOptionsInit()*Menu {
	menu := NewMenu()
	menu.Options = map[string]*MenuOptions{
		"WriteFile":&MenuOptions{
			ID:          1,
			OptName:     "WriteFile",
			Value:       "Y",
			Text:        "是否寫入記錄檔 .json 用於加載",
		},
		"WriteFilePath":&MenuOptions{
			ID:          2,
			OptName:     "WriteFilePath",
			Value:       "./Salary.json",
			Text:        "寫入json檔的路徑",
		},
	}

	menu.AllowAction = allowAct{
		"set":  Set,
		"exit": Exit,
	}
	return menu
}

func HourDataInit(opts map[string]*MenuOptions)(*Menu,error){

	intHoursParam,err := TimestringTransfertoInt(opts)
	if err != nil {
		return nil,err
	}
	ex := intHoursParam["ExcludeDays"].([]int)
	workDay := WorkDayList(intHoursParam["Year"].(int),intHoursParam["Month"].(int),ex...)

	var timefmtList  = []string{}


	switch intHoursParam["Action"] {
	case 1:
		hourZone, leftHour := intHoursParam["WorkHour"].(int)/4 ,intHoursParam["WorkHour"].(int)%4
		timefmtList = outCalcTimefmtList(hourZone,leftHour)
	default:
		for i:=0;i<len(workDay);i++{
			timefmtList = append(timefmtList,opts["WorkHour"].Value)
		}
	}


	menu:= NewMenu()
	menu.Options = make(map[string]*MenuOptions,20)
	day := fmt.Sprintf("d%s",workDay[0])
	menu.Options[day] = NewMenuOption(1,day,timefmtList[0])
	menu.Root = menu.Options[day]

	for i,v := range timefmtList {
		if i==0 {
			continue
		}
		day := fmt.Sprintf("d%s",workDay[i])
		if err := menu.OptionOperate("add",day,v);err != nil {
			return nil,	err
		}

	}

	menu.AllowAction = allowAct{
		"add":  Add,
		"rm":   Remove,
		"set":  Set,
		"exit": Exit,
	}

	return menu,nil

}