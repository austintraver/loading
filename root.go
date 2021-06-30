// Package loading provides a very simple loading wheel to display on a TTY
// that displays a spinning animation as time passes
package loading

import (
	"fmt"
	"sync/atomic"
	"time"
)

// states are what characters to cycle through as a Spinner is running
var states = []rune{'⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏', '⠋', '⠙', '⠹'}

// Spinner main type
type Spinner struct {
	frames []rune
	index  int
	active uint64
	text   string
}

// New Spinner with args
func New(text string) *Spinner {
	s := &Spinner{
		text: ("\r\033[K") + text,
	}
	s.Set(states)
	return s
}

// Set frames to the given string which must not use spaces.
func (target *Spinner) Set(frames []rune) {
	target.frames = frames
}

// Start shows the spinner.
func (target *Spinner) Start() *Spinner {
	if atomic.LoadUint64(&target.active) > 0 {
		return target
	}
	atomic.StoreUint64(&target.active, 1)
	go func() {
		for atomic.LoadUint64(&target.active) > 0 {
			fmt.Printf(target.text, target.next())
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return target
}

// The clear function moves the cursor to the beginning of the current line
// and erases any existing text, in order to replace it with the next
// incremental stage of the Spinner
func clear() (n int, err error) {
	return fmt.Print("\r\033[K")
}

// The Stop function hides a spinner. The value returned by Stop is true
// if a running spinner was stopped, and false otherwise.
func (target *Spinner) Stop() (ok bool, err error) {
	x := atomic.SwapUint64(&target.active, 0)
	if x > 0 {
		_, err = clear()
		ok = true
	}
	return
}

// The next function advances the current char displayed by the
// spinner to the symbol directly after the current symbol.
func (target *Spinner) next() (char string) {
	char = string(target.frames[target.index%len(target.frames)])
	target.index++
	return
}
