// Code generated from E:/video/github/goft-expr/expr/src/g4\BeanExpr.g4 by ANTLR 4.8. DO NOT EDIT.

package BeanExprLib // BeanExpr
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BeanExprListener is a complete expr for a parse tree produced by BeanExprParser.
type BeanExprListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterMethodCall is called when entering the methodCall production.
	EnterMethodCall(c *MethodCallContext)

	// EnterFuncCall is called when entering the FuncCall production.
	EnterFuncCall(c *FuncCallContext)

	// EnterFuncArgs is called when entering the FuncArgs production.
	EnterFuncArgs(c *FuncArgsContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitMethodCall is called when exiting the methodCall production.
	ExitMethodCall(c *MethodCallContext)

	// ExitFuncCall is called when exiting the FuncCall production.
	ExitFuncCall(c *FuncCallContext)

	// ExitFuncArgs is called when exiting the FuncArgs production.
	ExitFuncArgs(c *FuncArgsContext)
}
