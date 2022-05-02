package frog

import (
	"errors"
	"flag"
	"strconv"
)

type boolValue struct {
	ptr *bool
}

func (b *boolValue) String() string {
	return strconv.FormatBool(*b.ptr)
}

func (b boolValue) Set(val string) error {
	res, err := strconv.ParseBool(val)
	if err != nil {
		return errors.New("parse error")
	}
	b.ptr = &res
	return nil
}

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

// BoolVar defines a bool flag with specified name, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func BoolVar(p *bool, name string, usage string) {
	flag.CommandLine.Var(&boolValue{ptr: p}, name, usage)
}

// StringVar defines a string flag with specified name, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func StringVar(p *string, name string, usage string) {
	flag.CommandLine.Var(&stringValue{ptr: p}, name, usage)
}
