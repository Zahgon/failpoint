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
)

type exprRewriter func(rewriter *Rewriter, call *ast.CallExpr) (rewritten bool, result ast.Stmt, err error)

var exprRewriters = map[string]exprRewriter{
	"Inject":        (*Rewriter).rewriteInject,
	"InjectContext": (*Rewriter).rewriteInjectContext,
	"InjectCall":    (*Rewriter).rewriteInjectCall,
	"Break":         (*Rewriter).rewriteBreak,
	"Continue":      (*Rewriter).rewriteContinue,
	"Label":         (*Rewriter).rewriteLabel,
	"Goto":          (*Rewriter).rewriteGoto,
	"Fallthrough":   (*Rewriter).rewriteFallthrough,
	"Return":        (*Rewriter).rewriteReturn,
}

func (r *Rewriter) rewriteInject(call *ast.CallExpr) (bool, ast.Stmt, error) {
	_ = "STUB: not implemented"
	return false, *new(ast.Stmt), nil
}

// First argument need not to be a string literal, any string type stuff is ok.
// Type safe is convinced by compiler.

// failpoint.Inject("failpoint-name", nil)

// failpoint.Inject("failpoint-name", func(){...})
// failpoint.Inject("failpoint-name", func(val failpoint.Value){...})

// closure signature:
// func(val failpoint.Value) {...}
// func() {...}

func (r *Rewriter) rewriteInjectContext(call *ast.CallExpr) (bool, ast.Stmt, error) {
	_ = "STUB: not implemented"
	return false, *new(ast.Stmt), nil
}

// Second argument need not to be a identifier, any context type token (e.g. selector) is OK.
// Type safe is convinced by compiler.

// Second argument need not to be a string literal, any string type stuff is ok.
// Type safe is convinced by compiler.

// failpoint.InjectContext("failpoint-name", ctx, nil)

// failpoint.InjectContext("failpoint-name", ctx, func(){...})
// failpoint.InjectContext("failpoint-name", ctx, func(val failpoint.Value){...})

// closure signature:
// func(val failpoint.Value) {...}
// func() {...}

func (r *Rewriter) rewriteInjectCall(call *ast.CallExpr) (bool, ast.Stmt, error) {
	_ = "STUB: not implemented"
	return false, *new(ast.Stmt), nil
}

// First argument need not to be a string literal, any string type stuff is ok.
// Type safe is convinced by compiler.

// failpoint.InjectCall("name", a, b, c)
//    |
//    v
// failpoint.Call(_curpkg_("name"), a, b, c)

func (r *Rewriter) rewriteBreak(call *ast.CallExpr) (bool, ast.Stmt, error) {
	_ = "STUB: not implemented"
	return false, *new(ast.Stmt), nil
}

func (r *Rewriter) rewriteContinue(call *ast.CallExpr) (bool, ast.Stmt, error) {
	_ = "STUB: not implemented"
	return false, *new(ast.Stmt), nil
}

func (r *Rewriter) rewriteLabel(call *ast.CallExpr) (bool, ast.Stmt, error) {
	_ = "STUB: not implemented"
	return false, *new(ast.Stmt), nil
}

// It's a trick here

func (r *Rewriter) rewriteGoto(call *ast.CallExpr) (bool, ast.Stmt, error) {
	_ = "STUB: not implemented"
	return false, *new(ast.Stmt), nil
}

func (r *Rewriter) rewriteFallthrough(call *ast.CallExpr) (bool, ast.Stmt, error) {
	_ = "STUB: not implemented"
	return false, *new(ast.Stmt), nil
}

func (r *Rewriter) rewriteReturn(call *ast.CallExpr) (bool, ast.Stmt, error) {
	_ = "STUB: not implemented"
	return false, *new(ast.Stmt), nil
}
