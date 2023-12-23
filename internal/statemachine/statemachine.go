package statemachine

type StateMachine struct {
	CurrentState State
	Transaction  map[State]map[Event]State
}

func New() *StateMachine {
	return &StateMachine{
		CurrentState: None,
		Transaction:  make(map[State]map[Event]State),
	}
}

func (sm *StateMachine) Config() {
	// Стандартные команды
	sm.setTransaction(None, map[Event]State{
		Start:      SendBotInfoMessage,
		NewService: WaitingForServiceName,
	})

	sm.setTransaction(SendBotInfoMessage, nil)

	// Команды NewService
	// Ждет названия сервиса
	sm.setTransaction(WaitingForServiceName, map[Event]State{
		NewNameService: WaitingForPassword,
	})
	// Ждет пароля
	sm.setTransaction(WaitingForPassword, map[Event]State{
		NewNameService: WaitingForPassword,
	})
}

func (sm *StateMachine) setTransaction(currentState State, nextStates map[Event]State) {
	sm.Transaction[currentState] = nextStates
}
