package driver

import (
	"fmt"
	"github.com/tebeka/selenium"
)

type Driver selenium.WebDriver

type elemAction int

const (
	SendKey elemAction = iota
	Click
)

func IsGetPageSucc(d Driver, url,conditional string)bool{
	for i:=0;i<3;i++ {
		if err := d.Get(url);err != nil {
			fmt.Printf("Get url: %s failed!! retrying...\n",url)
			continue
		}

		if title, _  := d.Title(); title == conditional{
			return true
		}

		if currentUrl, _ := d.CurrentURL(); currentUrl == conditional {
			return true
		}
		break
	}
	return false
}

func WebElemAction(d Driver,Action elemAction ,parm ...string)error{
	elem,err := d.FindElement(parm[0],parm[1])

	if err != nil {
		return fmt.Errorf("Not Find Element %s by %s:%v",parm[1],parm[0],err)
	}
	switch Action{
	case SendKey:
		if err := elem.SendKeys(fmt.Sprintf("%s%s%s",selenium.BackspaceKey,selenium.BackspaceKey,selenium.BackspaceKey));err!= nil{
			return err
		}
		if err := elem.SendKeys(parm[2]); err != nil {
			return err
		}
		return nil
	case Click:
		if err := elem.Click(); err != nil {
			return err
		}
		return nil


	}
	panic("Action Not Found !!")
}



