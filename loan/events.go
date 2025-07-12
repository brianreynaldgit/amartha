// events.go
// Simulates an asynchronous message queue (e.g., Kafka or RabbitMQ).
// Logs emitted events in a background goroutine with thread-safe access.

package loan

import (
	"fmt"
	"sync"
)

type Event struct {
	Type string
	Data any
}

type EventBus struct {
	mu   sync.Mutex
	logs []Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		logs: make([]Event, 0),
	}
}

func (e *EventBus) EmitAsync(eventType string, data any) {
	go func() {
		e.mu.Lock()
		defer e.mu.Unlock()

		event := Event{Type: eventType, Data: data}
		e.logs = append(e.logs, event)
		fmt.Printf("[Async Event] %s — %+v\n", eventType, data)
	}()
}
