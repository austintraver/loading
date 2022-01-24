// Package loading provides a very simple loading wheel to display on a TTY
// that displays a spinning animation as time passes
package loading

import (
	"fmt"
	"time"

	"github.com/gosuri/uilive"
)

// Spinner main type
type Spinner struct {
	interval time.Duration
	writer   *uilive.Writer
}

// New Spinner with args
func New(interval time.Duration) *Spinner {
	return &Spinner{
		interval: interval,
		writer:   uilive.New(),
	}
}

func (s *Spinner) Loop(duration time.Duration) {
	deadline := time.Now().Add(duration)
	// states are what characters to cycle through as a Spinner is running
	states := []rune{'⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏', '⠋', '⠙', '⠹'}
	for _, c := range states {
		if time.Now().After(deadline) {
			break
		}
		_, err := fmt.Fprintf(s.writer, "Loading... %c\n", c)
		if err != nil {
			panic(err)
		}
		time.Sleep(s.interval)
	}
}

// The Stop function hides a spinner. The value returned by Stop is true
// if a running spinner was stopped, and false otherwise.
func (s *Spinner) Stop() {
	s.writer.Stop()
	return
}
