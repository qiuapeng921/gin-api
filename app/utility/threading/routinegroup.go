package threading

import "sync"

// RoutineGroup goroutine group
type RoutineGroup struct {
	waitGroup sync.WaitGroup
}

// NewRoutineGroup new RoutineGroup
func NewRoutineGroup() *RoutineGroup {
	return new(RoutineGroup)
}

// Run Don't reference the variables from outside,
// because outside variables can be changed by other goroutines
func (g *RoutineGroup) Run(fn func()) {
	g.waitGroup.Add(1)

	go func() {
		defer g.waitGroup.Done()
		fn()
	}()
}

// RunSafe Don't reference the variables from outside,
// because outside variables can be changed by other goroutines
func (g *RoutineGroup) RunSafe(fn func()) {
	g.waitGroup.Add(1)

	GoSafe(func() error {
		defer g.waitGroup.Done()
		fn()
		return nil
	})
}

// Wait 等待组完成
func (g *RoutineGroup) Wait() {
	g.waitGroup.Wait()
}
