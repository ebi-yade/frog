package frog

import "flag"

// stringValue represents nil-able flag.Value of string
type stringValue struct {
	ptr *string
}

func (s *stringValue) Set(val string) error {
	s.ptr = &val
	return nil
}

func (s *stringValue) String() string {
	return *s.ptr
}

// StringVar defines a string flag with specified name, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func StringVar(p *string, name string, usage string) {
	flag.CommandLine.Var(&stringValue{ptr: p}, name, usage)
}
