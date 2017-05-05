package streamutil

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type addressType int

const (
	invalidAddressType addressType = iota
	lineNumberAddressType
	noAddressType
	regexpAddressType

	print
	substitute
)

// SedOptions describes options for Sed functions.
type SedOptions struct {
	Silent  bool
	Inplace bool
}

// SedInstruction specifies the parameters to the Sed functions.
// It holds both the address parameter that selects the lines to
// perform sed operations on as well as the operation itself.
// Address at the moment supports only *regular expressions* and
// line numbers nothing as comprehensive like the sed tool.
type SedInstruction struct {
	Address string
	Pattern string
	Action  string
	SedOptions
}

// Sed implements a limited sed(1) funtion for filtering and transforming text.
func Sed(ins SedInstruction, files ...string) <-chan string {
	var addrRegexp *regexp.Regexp
	var addrType addressType
	var ln int

	lineNumber := 0
	out := make(chan string)

	// Address processing.
	addrType, err := getAddressType(ins)
	if err != nil {
		panic(err)
	}

	if addrType == regexpAddressType {
		a := strings.TrimLeft(ins.Address, "/")
		addrRegexp = regexp.MustCompile(strings.TrimRight(a, "/"))
	}

	if addrType == lineNumberAddressType {
		ln, err = strconv.Atoi(ins.Address)
		if err != nil {
			panic(err)
		}
	}

	// Action processing. Default action 'print'.

	go func() {
		// Open files for processing.
		for _, file := range files {
			f, err := os.Open(file)
			defer f.Close()

			if err != nil {
				if err != nil {
					panic(err)
				}
			}

			// Scan through file.
			s := bufio.NewScanner(f)
			for s.Scan() {
				l := s.Text()
				lineNumber++

				// Start cycle.
				// Read into pattern space.
				// TODO: Consider reading multiple lines into pattern space

				// SEDOPTSILENT
				if !ins.SedOptions.Silent {
					out <- l
				}

				patternSpace := []string{}
				patternSpace = append(patternSpace, l)

				// Operate on pattern space.
				// Rationale: on the off chance we have to deal with multiple lines
				// in pattern space.
				for _, l := range patternSpace {

					// TODO: Move check for address type out of for loop.
					// Possibly pass a function fn which does matches based
					// on address type.
					switch addrType {
					case regexpAddressType:
						if matched := addrRegexp.MatchString(l); matched {
							out <- l
						}
					case lineNumberAddressType:
						if ln == lineNumber {
							out <- l
						}
					case noAddressType:
						out <- l
					}
				}

				// clear pattern space.
				// Assumption: GC will take care of clearing pattern space.
			}

			if err := s.Err(); err != nil {
				if err != nil {
					panic(err)
				}
			}
		}
		// End of file processing.
		close(out)
	}()

	return out
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

func getAddressType(ins SedInstruction) (addressType, error) {
	if ins.Address == "" {
		return noAddressType, nil
	}

	// Detect regular expression address
	// TODO: Insist on supporting only a subset of RE syntax
	if matched, err := regexp.MatchString(`^\/.*\/$`, ins.Address); matched {
		if err != nil {
			return invalidAddressType, err
		}
		return regexpAddressType, nil
	}

	// Detect line number address
	_, err := strconv.Atoi(ins.Address)
	if err == nil {
		return lineNumberAddressType, nil
	}

	// TODO: Suspicious
	return invalidAddressType, nil
}
