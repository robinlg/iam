package term

import "github.com/moby/term"

// TerminalSize represents the width and height of a terminal.
type TerminalSize struct {
	Width  uint16
	Height uint16
}

// GetSize returns the current size of the terminal associated with fd.
func GetSize(fd uintptr) *TerminalSize {
	winsize, err := term.GetWinsize(fd)
	if err != nil {
		// runtime.HandleError(fmt.Errorf("unable to get terminal size: %v", err))
		return nil
	}

	return &TerminalSize{Width: winsize.Width, Height: winsize.Height}
}
