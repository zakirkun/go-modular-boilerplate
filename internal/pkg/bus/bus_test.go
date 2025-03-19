package bus

import "testing"

type testHandler struct {
	called bool
}

func (h *testHandler) Handle(event Event) {
	h.called = true
}

func TestEventBus(t *testing.T) {
	bus := NewEventBus()

	handler := &testHandler{}
	bus.Subscribe("test", handler)

	event := Event{Type: "test", Payload: "Hello, world!"}
	bus.Publish(event)

	bus.wg.Wait()

	if !handler.called {
		t.Errorf("Handler was not called")
	}

	t.Log("EventBus test passed")
}
