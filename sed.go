package streamutil

import (
	"bufio"
	"os"
)

// SedOptions describes options for Sed functions.
type SedOptions struct {
	Silent  bool
	Inplace bool
}

// SedInstruction specifies the parameters to the Sed functions.
// It holds both the address parameter that selects the lines to
// perform sed operations on as well as the operation itself.
type SedInstruction struct {
	Address string
	Pattern string
	Action  string
}

// Sed applies operations to the supplied files and returns a chan string.
func Sed(ins SedInstruction, opt SedOptions, out chan<- string, files ...string) {
	var patternSpace []string

	go func() {
		for _, file := range files {
			f, err := os.Open(file)
			defer f.Close()

			if err != nil {
			}

			s := bufio.NewScanner(f)
			for s.Scan() {
				patternSpace = append(patternSpace, s.Text())
				for _, l := range patternSpace {
					// Check address
					// Check pattern
					// Run command
					out <- l + "\n"
				}

			}

			if err := s.Err(); err != nil {
			}
		}
	}()
}

// SedPipe is similar to func Sed, except it receives input over a chan string,
// forming a I/O pipe of kinds.
// It allows chaining other functions Grep with Sed, for instance a Grep | Sed pipe.
func SedPipe(ins SedInstruction, opt SedOptions, in <-chan string) <-chan string {
	return make(<-chan string)
}

// SedRaw runs args in the standard command-line sed form against the supplied files.
func SedRaw(args string, file ...string) <-chan string {
	return make(<-chan string)
}

// SedRawPipe is similar to SedRaw with support for chaining other functions like Grep.
func SedRawPipe(args string, in <-chan string) <-chan string {
	return make(<-chan string)
}
