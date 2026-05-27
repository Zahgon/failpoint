// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Copyright 2016 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package failpoint

import (
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// terms encodes the state for a failpoint term string (see fail(9) for examples)
// <fp> :: <term> ( "->" <term> )*
type terms struct {
	// chain is a slice of all the terms from desc
	chain []*term
	// desc is the full term given for the failpoint
	desc string
	// mu protects the state of the terms chain
	mu sync.Mutex
}

// term is an executable unit of the failpoint terms chain
type term struct {
	desc string

	mods mod
	act  actFunc
	val  interface{}

	parent *terms
	fp     *Failpoint
}

type mod interface {
	allow() bool
}

type modCount struct{ c int }

func (mc *modCount) allow() bool { _ = "STUB: not implemented"; return false }

type modProb struct{ p float64 }

func (mp *modProb) allow() bool { _ = "STUB: not implemented"; return false }

type modList struct{ l []mod }

func (ml *modList) allow() bool { _ = "STUB: not implemented"; return false }

func newTerms(desc string, fp *Failpoint) (*terms, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *terms) String() string { _ = "STUB: not implemented"; return "" }

func (t *terms) eval() (Value, error) { _ = "STUB: not implemented"; return *new(Value), nil }

// split terms from a -> b -> ... into [a, b, ...]
func parse(desc string, fp *Failpoint) (chain []*term, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// <term> :: <mod> <act> [ "(" <val> ")" ]
func parseTerm(desc string, fp *Failpoint) *term { _ = "STUB: not implemented"; return nil }

// <mod> :: ((<float>|<int> "%")|(<int> "*" ))*
func parseMod(desc string) (ret string, mods []mod) { _ = "STUB: not implemented"; return "", nil }

// parseIntFloat parses an int or float from a string and returns the string
// it parsed it from (unlike scanf).
func parseIntFloat(desc string) (string, interface{}) {
	_ = "STUB: not implemented"
	// parse for ints
	return "", nil
}

// parse for floats

// parseAct parses an action
// <act> :: "off" | "return" | "sleep" | "panic" | "break" | "print" | "pause"
func parseAct(desc string) (string, actFunc) { _ = "STUB: not implemented"; return "", *new(actFunc) }

// <val> :: <int> | <string> | <bool> | <nothing>
func parseVal(desc string) (string, interface{}) {
	_ = "STUB: not implemented"
	// return => struct{}
	return "", nil
}

// malformed

// return() => struct{}

// return("s") => string

// return(1) => int

// return(true) => bool

// unknown type; malformed input?

type actFunc func(*term) (interface{}, error)

var actMap = map[string]actFunc{
	"off":    actOff,
	"return": actReturn,
	"sleep":  actSleep,
	"panic":  actPanic,
	"break":  actBreak,
	"print":  actPrint,
	"pause":  actPause,
}

func (t *term) do() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func actOff(t *term) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func actReturn(t *term) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func actSleep(t *term) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func actPause(t *term) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func actPanic(t *term) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func actBreak(t *term) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// wait for gdb prompt
// XXX: tried doing this by piping stdout here and waiting on "(gdb) "
// but the the output won't appear since the process is STOPed and
// can't copy it back to the actual stdout

// don't zombie gdb

func actPrint(t *term) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }
