package net

type EventHandler interface {
	RegisterEvent(event EventType, callback func(*Event))
	FireEvent(event *Event)
}

type EventHandlerImplement struct {
	eventHandlerMap map[EventType][]func(*Event)
}

func (Self *EventHandlerImplement) RegisterEvent(event EventType, callback func(*Event)) {
	if Self.eventHandlerMap == nil{
		Self.eventHandlerMap = make(map[EventType][]func(*Event))
	}
	if v, ok := Self.eventHandlerMap[event]; ok {
		v = append(v, callback)
	} else {
		v := []func(*Event){callback}
		Self.eventHandlerMap[event] = v
	}
}

func (Self *EventHandlerImplement) FireEvent(event *Event)  {
	if Self.eventHandlerMap == nil {
		return
	}

	if v, ok := Self.eventHandlerMap[event.Type]; ok {
		for i := 0; i < len(v); i++{
			v[i](event)
		}
	}
}
