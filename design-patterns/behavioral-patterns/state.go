package behavioralpatterns

import "fmt"

// State
type DocState interface {
  state()
}
// Concrete State
type Draft struct {}

func (d *Draft) state() {
  fmt.Println("Draft State")
}

type Moderation struct {}

func (m *Moderation) state() {
  fmt.Println("Moderation state")
}

type Published struct {}

func (p *Published) state() {
  fmt.Println("Published state")
}

// Context
type stateContext struct {
  currentState DocState
}

func GetContext() *stateContext {
  return &stateContext{
    currentState: &Draft{},
  }
}

func (s *stateContext) SetState(state DocState) {
  s.currentState = state
}

func (s *stateContext) GetState() {
  s.currentState.state()
}
