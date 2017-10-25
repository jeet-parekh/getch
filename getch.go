package getch

import (
	"github.com/jeet-parekh/coninlogger"
)

// Getch gets a character entry from the console. It returns an error if unsuccessful.
func Getch() (ch uint16, err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			ch = 0
			err = panicErr.(error)
		}
	}()

	ch = MustGetch()
	return
}

// MustGetch gets a character entry from the console. It panics if unsuccessful.
func MustGetch() (ch uint16) {
	cin := coninlogger.NewConsoleInputLogger(1)
	cinmsgs := cin.GetMessageChannel()
	cin.Start()

	for msg := range cinmsgs {
		if msg.UnicodeCharacter != 0 {
			ch = msg.UnicodeCharacter
			break
		}
	}
	cin.Stop()

	return ch
}
