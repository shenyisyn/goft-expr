// Code generated from E:/video/github/goft-expr/expr/src/g4\BeanExpr.g4 by ANTLR 4.8. DO NOT EDIT.

package BeanExprLib

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 98, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 42, 10, 6, 12, 6, 14, 6, 45, 11, 6, 3, 6,
	3, 6, 3, 7, 5, 7, 50, 10, 7, 3, 7, 3, 7, 3, 8, 5, 8, 55, 10, 8, 3, 8, 6,
	8, 58, 10, 8, 13, 8, 14, 8, 59, 3, 9, 3, 9, 7, 9, 64, 10, 9, 12, 9, 14,
	9, 67, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 6, 10, 73, 10, 10, 13, 10, 14,
	10, 74, 3, 11, 3, 11, 3, 12, 6, 12, 80, 10, 12, 13, 12, 14, 12, 81, 5,
	12, 84, 10, 12, 3, 12, 3, 12, 6, 12, 88, 10, 12, 13, 12, 14, 12, 89, 3,
	13, 6, 13, 93, 10, 13, 13, 13, 14, 13, 94, 3, 13, 3, 13, 2, 2, 14, 3, 3,
	5, 4, 7, 5, 9, 2, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
	13, 3, 2, 7, 3, 2, 50, 59, 4, 2, 41, 41, 94, 94, 4, 2, 67, 92, 99, 124,
	5, 2, 50, 59, 67, 92, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 108, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2,
	13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2,
	2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 3, 27, 3, 2, 2,
	2, 5, 29, 3, 2, 2, 2, 7, 31, 3, 2, 2, 2, 9, 33, 3, 2, 2, 2, 11, 35, 3,
	2, 2, 2, 13, 49, 3, 2, 2, 2, 15, 54, 3, 2, 2, 2, 17, 61, 3, 2, 2, 2, 19,
	68, 3, 2, 2, 2, 21, 76, 3, 2, 2, 2, 23, 83, 3, 2, 2, 2, 25, 92, 3, 2, 2,
	2, 27, 28, 7, 42, 2, 2, 28, 4, 3, 2, 2, 2, 29, 30, 7, 43, 2, 2, 30, 6,
	3, 2, 2, 2, 31, 32, 7, 46, 2, 2, 32, 8, 3, 2, 2, 2, 33, 34, 9, 2, 2, 2,
	34, 10, 3, 2, 2, 2, 35, 43, 7, 41, 2, 2, 36, 37, 7, 94, 2, 2, 37, 42, 11,
	2, 2, 2, 38, 39, 7, 41, 2, 2, 39, 42, 7, 41, 2, 2, 40, 42, 10, 3, 2, 2,
	41, 36, 3, 2, 2, 2, 41, 38, 3, 2, 2, 2, 41, 40, 3, 2, 2, 2, 42, 45, 3,
	2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 46, 3, 2, 2, 2, 45,
	43, 3, 2, 2, 2, 46, 47, 7, 41, 2, 2, 47, 12, 3, 2, 2, 2, 48, 50, 7, 47,
	2, 2, 49, 48, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 52,
	5, 23, 12, 2, 52, 14, 3, 2, 2, 2, 53, 55, 7, 47, 2, 2, 54, 53, 3, 2, 2,
	2, 54, 55, 3, 2, 2, 2, 55, 57, 3, 2, 2, 2, 56, 58, 5, 9, 5, 2, 57, 56,
	3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2,
	60, 16, 3, 2, 2, 2, 61, 65, 9, 4, 2, 2, 62, 64, 9, 5, 2, 2, 63, 62, 3,
	2, 2, 2, 64, 67, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66,
	18, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 68, 72, 5, 17, 9, 2, 69, 70, 5, 21,
	11, 2, 70, 71, 5, 17, 9, 2, 71, 73, 3, 2, 2, 2, 72, 69, 3, 2, 2, 2, 73,
	74, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 20, 3, 2, 2,
	2, 76, 77, 7, 48, 2, 2, 77, 22, 3, 2, 2, 2, 78, 80, 5, 9, 5, 2, 79, 78,
	3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2,
	82, 84, 3, 2, 2, 2, 83, 79, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85, 3,
	2, 2, 2, 85, 87, 7, 48, 2, 2, 86, 88, 5, 9, 5, 2, 87, 86, 3, 2, 2, 2, 88,
	89, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 24, 3, 2, 2,
	2, 91, 93, 9, 6, 2, 2, 92, 91, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 92,
	3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 97, 8, 13, 2, 2,
	97, 26, 3, 2, 2, 2, 14, 2, 41, 43, 49, 54, 59, 65, 74, 81, 83, 89, 94,
	3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "','", "", "", "", "", "", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "StringArg", "FloatArg", "IntArg", "FuncName", "MethodName",
	"Dot", "Float", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "DIGIT", "StringArg", "FloatArg", "IntArg", "FuncName",
	"MethodName", "Dot", "Float", "WHITESPACE",
}

type BeanExprLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewBeanExprLexer(input antlr.CharStream) *BeanExprLexer {

	l := new(BeanExprLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "BeanExpr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BeanExprLexer tokens.
const (
	BeanExprLexerT__0       = 1
	BeanExprLexerT__1       = 2
	BeanExprLexerT__2       = 3
	BeanExprLexerStringArg  = 4
	BeanExprLexerFloatArg   = 5
	BeanExprLexerIntArg     = 6
	BeanExprLexerFuncName   = 7
	BeanExprLexerMethodName = 8
	BeanExprLexerDot        = 9
	BeanExprLexerFloat      = 10
	BeanExprLexerWHITESPACE = 11
)
