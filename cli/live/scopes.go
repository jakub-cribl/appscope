package live

import (
	"errors"

	"github.com/criblio/scope/libscope"
)

// Scope is an object model that represents an instance of a scoped process
type Scope struct {
	Conn         UnixConnection  `json:"-"`
	ProcessStart libscope.Header `json:"processStart"`
}

// Create a Scope
func NewScope() *Scope {
	scope := &Scope{}
	return scope
}

// Update a Scope
func (s *Scope) Update(obj interface{}) error {
	switch o := obj.(type) {
	// Update UnixConnection
	case UnixConnection:
		s.Conn = o
	// Update Header
	case libscope.Header:
		s.ProcessStart = o
	default:
		return errors.New("Invalid type passed into Update")
	}

	return nil
}