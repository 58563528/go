// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

type token uint

//go:generate stringer -type token -linecomment tokens.go

const (
	_    token = iota
	_EOF       // EOF

	// names and literals
	_Name    // name
	_Literal // literal

	// operators and operations
	// _Operator is excluding '*' (_Star)
	_Operator // op
	_AssignOp // op=
	_IncOp    // opop
	_Assign   // =
	_Define   // :=
	_Arrow    // <-
	_Star     // *

	// delimiters
	_Lparen    // (
	_Lbrack    // [
	_Lbrace    // {
	_Rparen    // )
	_Rbrack    // ]
	_Rbrace    // }
	_Comma     // ,
	_Semi      // ;
	_Colon     // :
	_Dot       // .
	_DotDotDot // ...

	// keywords
	_Break       // 跳出
	_Case        // 为
	_Chan        // 管道
	_Const       // 常量
	_Continue    // 继续
	_Default     // 为其他
	_Defer       // 推迟
	_Else        // 否则
	_Fallthrough // 贯穿
	_For         // 循环
	_Func        // 函数
	_Go          // 异步
	_Goto        // 跳转
	_If          // 如果
	_Import      // 导入
	_Interface   // 接口
	_Map         // 映射表
	_Package     // 包
	_Range       // 范围
	_Return      // 返回
	_Select      // 选择
	_Struct      // 类
	_Switch      // 假如
	_Type        // 类型
	_Var         // 变量

	// empty line comment to exclude it from .String
	tokenCount //
)

const (
	// for BranchStmt
	Break       = _Break
	Continue    = _Continue
	Fallthrough = _Fallthrough
	Goto        = _Goto

	// for CallStmt
	Go    = _Go
	Defer = _Defer
)

// Make sure we have at most 64 tokens so we can use them in a set.
const _ uint64 = 1 << (tokenCount - 1)

// contains reports whether tok is in tokset.
func contains(tokset uint64, tok token) bool {
	return tokset&(1<<tok) != 0
}

type LitKind uint8

// TODO(gri) With the 'i' (imaginary) suffix now permitted on integer
//           and floating-point numbers, having a single ImagLit does
//           not represent the literal kind well anymore. Remove it?
const (
	IntLit LitKind = iota
	FloatLit
	ImagLit
	RuneLit
	StringLit
)

type Operator uint

//go:generate stringer -type Operator -linecomment tokens.go

const (
	_ Operator = iota

	// Def is the : in :=
	Def   // :
	Not   // !
	Recv  // <-
	Tilde // ~

	// precOrOr
	OrOr // ||

	// precAndAnd
	AndAnd // &&

	// precCmp
	Eql // ==
	Neq // !=
	Lss // <
	Leq // <=
	Gtr // >
	Geq // >=

	// precAdd
	Add // +
	Sub // -
	Or  // |
	Xor // ^

	// precMul
	Mul    // *
	Div    // /
	Rem    // %
	And    // &
	AndNot // &^
	Shl    // <<
	Shr    // >>
)

// Operator precedences
const (
	_ = iota
	precOrOr
	precAndAnd
	precCmp
	precAdd
	precMul
)
