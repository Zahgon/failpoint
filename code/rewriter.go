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

package code

import (
	"go/ast"
	"go/token"
	"io"
)

const (
	packagePath     = "github.com/pingcap/failpoint"
	packageName     = "failpoint"
	evalFunction    = "Eval"
	callFunction    = "Call"
	evalCtxFunction = "EvalContext"
	ExtendPkgName   = "_curpkg_"
	// It is an indicator to indicate the label is converted from `failpoint.Label("...")`
	// We use an illegal suffix to avoid conflict with the user's code
	// So `failpoint.Label("label1")` will be converted to `label1-tmp-marker:` in expression
	// rewrite and be converted to the legal form in label statement organization.
	labelSuffix = "-tmp-marker"
)

// Rewriter represents a rewriting tool for converting the failpoint marker functions to
// corresponding statements in Golang. It will traverse the specified path and filter
// out files which do not have failpoint injection sites, and rewrite the remain files.
type Rewriter struct {
	rewriteDir      string
	currentPath     string
	currentFile     *ast.File
	currsetFset     *token.FileSet
	failpointName   string
	allowNotChecked bool
	rewritten       bool

	output io.Writer
}

// NewRewriter returns a non-nil rewriter which is used to rewrite the specified path
func NewRewriter(path string) *Rewriter { _ = "STUB: not implemented"; return nil }

// SetOutput sets a writer and the rewrite results will write to the writer instead of generate a stash file
func (r *Rewriter) SetOutput(out io.Writer) {
	_ = "STUB: not implemented"

	// SetAllowNotChecked sets whether the rewriter allows the file which does not import failpoint package.
	return
}

func (r *Rewriter) SetAllowNotChecked(b bool) { _ = "STUB: not implemented"; return }

// GetRewritten returns whether the rewriter has rewritten the file in a RewriteFile call.
func (r *Rewriter) GetRewritten() bool {
	_ = "STUB: not implemented"

	// GetCurrentFile returns the current file which is being rewritten
	return false
}

func (r *Rewriter) GetCurrentFile() *ast.File { _ = "STUB: not implemented"; return nil }

func (r *Rewriter) pos(pos token.Pos) string { _ = "STUB: not implemented"; return "" }

func (r *Rewriter) rewriteFuncLit(fn *ast.FuncLit) error { _ = "STUB: not implemented"; return nil }

func (r *Rewriter) rewriteAssign(v *ast.AssignStmt) error {
	_ = "STUB: not implemented"
	// fn1, fn2, fn3, ... := func(){...}, func(){...}, func(){...}, ...
	//
	//	x, fn := 100, func() {
	//	    failpoint.Marker(fpname, func() {
	//	        ...
	//	    })
	//	}
	//
	//	ch := <-func() chan interface{} {
	//	    failpoint.Marker(fpname, func() {
	//	        ...
	//	    })
	//	}
	return nil
}

// rewriteInitStmt rewrites non-nil initialization statement
func (r *Rewriter) rewriteInitStmt(v ast.Stmt) error { _ = "STUB: not implemented"; return nil }

func (r *Rewriter) rewriteIfStmt(v *ast.IfStmt) error {
	_ = "STUB: not implemented"
	// if a, b := func() {...}, func() int {...}(); cond {...}
	// if func() {...}(); cond {...}
	return nil
}

func (r *Rewriter) rewriteExpr(expr ast.Expr) error { _ = "STUB: not implemented"; return nil }

// expressions that can not inject failpoint

// func()[]int {}()[func()int{}()]

// array[low:high:max]
// => func()[]int {}()[func()int{}():func()int{}():func()int{}()]

// return func(){...},

// []int{func() int {...}()}

// return func() int {...}()

// return fn(func() int{...})

// *func() *T{}()

// !func() {...}()

// a && func() bool {...} ()
// func() bool {...} () && a
// func() bool {...} () && func() bool {...} () && a
// func() bool {...} () && a && func() bool {...} () && a

// (func() {...}())

// (func() {...}()).(type)

// Key: (func() {...}())

func (r *Rewriter) rewriteExprs(exprs []ast.Expr) error { _ = "STUB: not implemented"; return nil }

func (r *Rewriter) rewriteStmts(stmts []ast.Stmt) error { _ = "STUB: not implemented"; return nil }

// var fn1, fn2, fn3, ... = func(){...}, func(){...}, func(){...}, ...
// var x, fn = 100, func() {
//     failpoint.Marker(fpname, func() {
//         ...
//     })
// }

// failpoint.Marker("failpoint.name", func(context.Context, *failpoint.Arg)) {...}
// failpoint.Break()
// failpoint.Break("label")
// failpoint.Continue()
// failpoint.Fallthrough()
// failpoint.Continue("label")
// failpoint.Goto("label")
// failpoint.Label("label")

// x := (func() {...} ())

// go func() {...}()
// go func(fn) {...}(func(){...})

// defer func() {...}()
// defer func(fn) {...}(func(){...})

// return func() {...}()
// return func(fn) {...}(func(){...})

// {
//     func() {...}()
// }

// if func() {...}() {...}

// case func() int {...}() > 100 && func () bool {...}()

// case func() int {...}() > 100 && func () bool {...}():
//     fn := func(){...}
//     fn()

// switch x := func() {...}(); {...}

// switch (func() {...}()) {...}

// switch x {
// case 1:
// 	func() {...}()
// }

// select {
// case ch := <-func() chan bool {...}():
// case <- fromCh:
// case toCh <- x:
// case <- func() chan bool {...}():
// default:
// }

// for i := func() int {...}(); i < func() int {...}(); i += func() int {...}() {...}
// for iter.Begin(); !iter.End(); iter.Next() {...}

// for i := range func() {...}() {...}

// 	switch x := (func () {...}()).(type) {...}

// 	switch (func () {...}()).(type) {...}

// 	ch <- func () {...}()

// Label:
//     func () {...}()

// func() *FooType {...}().Field++
// func() *FooType {...}().Field--

// ignore keyword token (BREAK, CONTINUE, GOTO, FALLTHROUGH)

// Label statement must ahead of for loop

func (r *Rewriter) rewriteFuncDecl(fn *ast.FuncDecl) error { _ = "STUB: not implemented"; return nil }

// RewriteFile rewrites a single file
func (r *Rewriter) RewriteFile(path string) (err error) {
	_ = "STUB: not implemented"
	// Reset state up front so previous-file result does not leak into
	// files that have nothing to rewrite (for example, doc.go).
	return nil
}

// Generate binding code

// Backup origin file and replace content

// Rewrite does the rewrite action for specified path. It contains the main steps:
//
// 1. Filter out failpoint binding files and files that have no suffix `.go`
// 2. Filter out files which have not imported failpoint package (implying no failpoints)
// 3. Parse file to `ast.File` and rewrite the AST
// 4. Create failpoint binding file (which contains `_curpkg_` function) if it does not exist
// 5. Rename original file to `original-file-name + __failpoint_stash__`
// 6. Replace original file content base on the new AST
func (r *Rewriter) Rewrite() error { _ = "STUB: not implemented"; return nil }

// Will rewrite a file only if the file has imported "github.com/pingcap/failpoint"

// import path maybe in the form of:
//
// 1. normal import
//    - "github.com/pingcap/failpoint"
//    - `github.com/pingcap/failpoint`
// 2. ignore import
//    - _ "github.com/pingcap/failpoint"
//    - _ `github.com/pingcap/failpoint`
// 3. alias import
//    - alias "github.com/pingcap/failpoint"
//    - alias `github.com/pingcap/failpoint`
// we should trim '"' or '`' before compare it.
