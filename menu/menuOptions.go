package menu

type MenuOptions struct {
	ID          int
	OptName     string
	Value       string
	Text        string `json:"-"`
	CallSubMenu *Menu
	Prev        *MenuOptions `json:"_"`
	Next        *MenuOptions `json:"_"`
}

func NewMenuOption(ID int, Name, Value string) *MenuOptions {
	mo := new(MenuOptions)
	mo.ID = ID
	mo.OptName = Name
	mo.Value = Value
	return mo
}

func (rootmo *MenuOptions) TraverseBy(condition interface{}) (*MenuOptions, bool) {
	switch condition.(type) {
	case int:
		if rootmo.ID == condition.(int) {
			return rootmo, true
		}
	case string:
		if rootmo.OptName == condition.(string) {
			return rootmo, true
		}
	case *MenuOptions:
		if rootmo == condition.(*MenuOptions) {
			return rootmo,true
		}
	case nil:
		if rootmo.Next == nil {
			return rootmo, true
		}
	}

	if rootmo.Next == nil {
		return nil, false
	}
	rootmo = rootmo.Next
	return rootmo.TraverseBy(condition)
}

func (rootmo *MenuOptions) AddOption(ID int, Name, Value string) (*MenuOptions, bool) {
	newMO := NewMenuOption(ID, Name, Value)
	if last, ok := rootmo.TraverseBy(nil); ok {
		last.Next = newMO
		newMO.Prev = last
		return newMO, true
	}
	return nil, false
}

func (rootmo *MenuOptions) DelOptionBy(condition interface{}) bool {
	if search, ok := rootmo.TraverseBy(condition); ok {
		search.Prev.Next = search.Next
		search.Next.Prev = search.Prev
		rootmo.ReNumber(1)
		return true
	}
	return false
}

func (rootmo *MenuOptions) ReNumber(id int) {
	rootmo.ID = id
	if rootmo.Next == nil {
		return
	}
	rootmo.Next.ReNumber(id + 1)
}

func (mo *MenuOptions) SetValue(value string) {
	mo.Value = value
}
