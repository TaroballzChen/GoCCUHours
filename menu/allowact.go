package menu

type ActionID = int

const (
	Set ActionID = iota
	Add
	Remove
	Run
	Number
	HourData
	Exit
)

type allowAct map[string]ActionID