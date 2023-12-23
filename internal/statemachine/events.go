package statemachine

type Event int

const (
	Start Event = iota
	NewService
	NewNameService
	GeneratePassword
	CustomMessage
	UnknownCommand
)
