package menu

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"unicode"
)

type MenuItemCallSubMenus map[string]*Menu

type Menu struct {
	Options     map[string]*MenuOptions
	AllowAction allowAct     `json:"_"`
	Root        *MenuOptions `json:"_"`
	RunAction   func()error  `json:"-"`
}

func NewMenu() *Menu {
	return &Menu{}
}

func (m *Menu) Show()error{
	for {
		Clear()
		color.Set(color.FgHiMagenta,color.Bold)
		fmt.Println(ui)
		color.Unset()
		m.ListOpt()
		m.ListAllowAction()
		s,err := m.GetUserInput()
		if err != nil {
			return err
		}

		if ok,err := m.CheckAction(s[0]);!ok {
			return err
		}

		u := NewUserInput(s[0],s[1],s[2])

		if err := m.OptionOperate(u.Action,u.Receptor,u.Value); err != nil {
			return err
		}
	}
}

func (m *Menu) ListOpt() {
	color.Set(color.FgHiCyan,color.Bold)
	defer color.Unset()
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 20,0,0,' ',0)
	for i:=1;i<=len(m.Options);i++{
		v,_ := m.GetReceptorByID(i)
		_, _ = fmt.Fprintf(w, "%d). %s\t%s\t%s\t\n", v.ID, v.OptName, v.Value, v.Text)
		_ = w.Flush()
	}
	fmt.Printf("--------------------------------------------------------------------------------\n")
}

func (m *Menu) ListAllowAction() {
	color.Set(color.FgHiCyan)
	defer color.Unset()
	fmt.Println("Allow Action:")
	for k,v := range m.AllowAction {
		if v == Number {
			continue
		}
		fmt.Printf("\t\t%s\n",k)
	}
	fmt.Printf("--------------------------------------------------------------------------------\n")
}

func (m *Menu) GetUserInput()([]string,error){
	var Userin string
	fmt.Printf(">>>")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		Userin = scanner.Text()
	}
	switch result := strings.SplitN(Userin, " ",3); len(result) {
	case 1:
		result = append(result,"")
		fallthrough
	case 2:
		result = append(result,"")
		fallthrough
	case 3:
		return result,nil
	default:
		return nil, errors.New("Invalid operation !!")
	}

}

func (m *Menu)AppendOption(optName string,options *MenuOptions) {
	m.Options[optName] = options
}

func (m *Menu)DeleteOption(key string){
	if _,ok := m.Options[key] ;ok {
		delete(m.Options,key)
		return
	}
	panic("del op failed")
}

func (m *Menu) CheckAction(op string)(bool,error) {
	if m.CheckExit(op) {
		return false, nil
	}
	if _,ok := m.AllowAction[op];ok {
		return true,nil
	}
	return false,fmt.Errorf("%q is not allow action",op)
}

func (m *Menu) CheckExit(op string)bool{
	var exitop = []string{"Exit","exit","quit","Quit"}
	for _,v :=  range exitop {
		if v == op {
			return true
		}
	}
	return false
}

func (m *Menu) CheckReceptorExists(Receptor string)(*MenuOptions,error){

	switch unicode.IsDigit([]rune(Receptor)[0]) {
	case true:
		ID,_ := strconv.Atoi(Receptor)
		if ok,err := m.CheckOptValid(ID);!ok {
			return nil,err
		}
		return m.GetReceptorByID(ID)

	case false:
		if v, ok := m.Options[Receptor];ok {
			return v,nil
		}
		fallthrough
	default:
		return nil,fmt.Errorf("Receptor %q do not be operated",Receptor)
	}
}

func (m *Menu)GetReceptorByID(ID int)(*MenuOptions,error) {
	for _,v := range(m.Options) {
		if ID == v.ID {
			return v, nil
		}
	}
	return nil,fmt.Errorf("ID %q receptor isn't exists")
}


func (m *Menu)CheckOptValid(s int)(bool,error){
	if s > len(m.Options) {
		return false,fmt.Errorf("%d option is not in list",s)
	}
	return true,nil
}

func (m *Menu)OptionOperate(action,receptor,value string)error{
	switch action{
	case "1","2","3","4","5","6","7","8","9","10","11","12","13","14","15","16","17":
		r,err := m.CheckReceptorExists(action)
		if err != nil {
			return err
		}
		if err := r.CallSubMenu.Show();err != nil {
			return err
		}
		return nil


	case "add":
		if r, _:= m.CheckReceptorExists(receptor); r != nil {
			return fmt.Errorf("%q is exists, do not be add",receptor)
		}
		if NewOp,ok := m.Root.AddOption(len(m.Options)+1,receptor,value);ok{
			m.AppendOption(receptor,NewOp)
			return nil
		}
		panic("No Last Node!!")

	case "set":
		r,err := m.CheckReceptorExists(receptor)
		if err != nil {
			return err
		}
		r.SetValue(value)
		return nil

	case "rm":
		r,err := m.CheckReceptorExists(receptor)
		if err != nil {
			return err
		}
		switch{
		case r.ID ==1 && r.Next == nil:
			return fmt.Errorf("%q must exist",r.OptName)
		case r.ID ==1 && r.Next != nil:
			m.DeleteOption(r.OptName)
			m.Root = r.Next
			m.Root.ReNumber(1)
			return nil
		default:
			m.DeleteOption(r.OptName)
			if ok := m.Root.DelOptionBy(r);!ok{
				panic("rm Action Failed")
			}
			return nil
		}

	case "run":
		return m.RunAction()

	default:
		return fmt.Errorf("%q IS NOT ALLOW ACTION",action)
	}
}

