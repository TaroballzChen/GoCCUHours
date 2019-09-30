package menu

type UserInput struct {
	Action string
	Receptor string
	Value string
}

func NewUserInput(action,receptor,value string)*UserInput {
	return &UserInput{
		Action:   action,
		Receptor: receptor,
		Value:    value,
	}
}