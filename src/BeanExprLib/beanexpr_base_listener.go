// Code generated from E:/video/github/goft-expr/expr/src/g4\BeanExpr.g4 by ANTLR 4.8. DO NOT EDIT.

package BeanExprLib // BeanExpr
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBeanExprListener is a complete listener for a parse tree produced by BeanExprParser.
type BaseBeanExprListener struct{}

var _ BeanExprListener = &BaseBeanExprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBeanExprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBeanExprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBeanExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBeanExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseBeanExprListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseBeanExprListener) ExitStart(ctx *StartContext) {}

// EnterMethodCall is called when production methodCall is entered.
func (s *BaseBeanExprListener) EnterMethodCall(ctx *MethodCallContext) {}

// ExitMethodCall is called when production methodCall is exited.
func (s *BaseBeanExprListener) ExitMethodCall(ctx *MethodCallContext) {}

// EnterFuncCall is called when production FuncCall is entered.
func (s *BaseBeanExprListener) EnterFuncCall(ctx *FuncCallContext) {}

// ExitFuncCall is called when production FuncCall is exited.
func (s *BaseBeanExprListener) ExitFuncCall(ctx *FuncCallContext) {}

// EnterFuncCallWithCompare is called when production FuncCallWithCompare is entered.
func (s *BaseBeanExprListener) EnterFuncCallWithCompare(ctx *FuncCallWithCompareContext) {}

// ExitFuncCallWithCompare is called when production FuncCallWithCompare is exited.
func (s *BaseBeanExprListener) ExitFuncCallWithCompare(ctx *FuncCallWithCompareContext) {}

// EnterFuncArgs is called when production FuncArgs is entered.
func (s *BaseBeanExprListener) EnterFuncArgs(ctx *FuncArgsContext) {}

// ExitFuncArgs is called when production FuncArgs is exited.
func (s *BaseBeanExprListener) ExitFuncArgs(ctx *FuncArgsContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseBeanExprListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseBeanExprListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseBeanExprListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseBeanExprListener) ExitConstant(ctx *ConstantContext) {}
