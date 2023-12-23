package statemachine

type State int

const (
	None State = iota
	SendBotInfoMessage
	WaitingForServiceName
	WaitingForPassword
	WaitingForPasswordType
)
