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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 22, 149,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 7, 6, 60, 10, 6, 12, 6, 14, 6, 63, 11, 6, 3, 6, 3, 6, 3, 7, 5, 7,
	68, 10, 7, 3, 7, 3, 7, 3, 8, 5, 8, 73, 10, 8, 3, 8, 6, 8, 76, 10, 8, 13,
	8, 14, 8, 77, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5,
	9, 89, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 7, 11, 97, 10,
	11, 12, 11, 14, 11, 100, 11, 11, 3, 12, 3, 12, 3, 12, 3, 12, 6, 12, 106,
	10, 12, 13, 12, 14, 12, 107, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15,
	3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3,
	19, 3, 20, 3, 20, 3, 20, 3, 21, 6, 21, 131, 10, 21, 13, 21, 14, 21, 132,
	5, 21, 135, 10, 21, 3, 21, 3, 21, 6, 21, 139, 10, 21, 13, 21, 14, 21, 140,
	3, 22, 6, 22, 144, 10, 22, 13, 22, 14, 22, 145, 3, 22, 3, 22, 2, 2, 23,
	3, 3, 5, 4, 7, 5, 9, 2, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23,
	12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41,
	21, 43, 22, 3, 2, 7, 3, 2, 50, 59, 4, 2, 41, 41, 94, 94, 4, 2, 67, 92,
	99, 124, 5, 2, 50, 59, 67, 92, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2,
	160, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3,
	2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27,
	3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2,
	35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2,
	2, 43, 3, 2, 2, 2, 3, 45, 3, 2, 2, 2, 5, 47, 3, 2, 2, 2, 7, 49, 3, 2, 2,
	2, 9, 51, 3, 2, 2, 2, 11, 53, 3, 2, 2, 2, 13, 67, 3, 2, 2, 2, 15, 72, 3,
	2, 2, 2, 17, 88, 3, 2, 2, 2, 19, 90, 3, 2, 2, 2, 21, 94, 3, 2, 2, 2, 23,
	101, 3, 2, 2, 2, 25, 109, 3, 2, 2, 2, 27, 111, 3, 2, 2, 2, 29, 114, 3,
	2, 2, 2, 31, 116, 3, 2, 2, 2, 33, 118, 3, 2, 2, 2, 35, 120, 3, 2, 2, 2,
	37, 123, 3, 2, 2, 2, 39, 126, 3, 2, 2, 2, 41, 134, 3, 2, 2, 2, 43, 143,
	3, 2, 2, 2, 45, 46, 7, 42, 2, 2, 46, 4, 3, 2, 2, 2, 47, 48, 7, 43, 2, 2,
	48, 6, 3, 2, 2, 2, 49, 50, 7, 46, 2, 2, 50, 8, 3, 2, 2, 2, 51, 52, 9, 2,
	2, 2, 52, 10, 3, 2, 2, 2, 53, 61, 7, 41, 2, 2, 54, 55, 7, 94, 2, 2, 55,
	60, 11, 2, 2, 2, 56, 57, 7, 41, 2, 2, 57, 60, 7, 41, 2, 2, 58, 60, 10,
	3, 2, 2, 59, 54, 3, 2, 2, 2, 59, 56, 3, 2, 2, 2, 59, 58, 3, 2, 2, 2, 60,
	63, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 64, 3, 2, 2,
	2, 63, 61, 3, 2, 2, 2, 64, 65, 7, 41, 2, 2, 65, 12, 3, 2, 2, 2, 66, 68,
	7, 47, 2, 2, 67, 66, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2,
	69, 70, 5, 41, 21, 2, 70, 14, 3, 2, 2, 2, 71, 73, 7, 47, 2, 2, 72, 71,
	3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 75, 3, 2, 2, 2, 74, 76, 5, 9, 5, 2,
	75, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3,
	2, 2, 2, 78, 16, 3, 2, 2, 2, 79, 80, 7, 118, 2, 2, 80, 81, 7, 116, 2, 2,
	81, 82, 7, 119, 2, 2, 82, 89, 7, 103, 2, 2, 83, 84, 7, 104, 2, 2, 84, 85,
	7, 99, 2, 2, 85, 86, 7, 110, 2, 2, 86, 87, 7, 117, 2, 2, 87, 89, 7, 103,
	2, 2, 88, 79, 3, 2, 2, 2, 88, 83, 3, 2, 2, 2, 89, 18, 3, 2, 2, 2, 90, 91,
	7, 112, 2, 2, 91, 92, 7, 107, 2, 2, 92, 93, 7, 110, 2, 2, 93, 20, 3, 2,
	2, 2, 94, 98, 9, 4, 2, 2, 95, 97, 9, 5, 2, 2, 96, 95, 3, 2, 2, 2, 97, 100,
	3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 22, 3, 2, 2, 2,
	100, 98, 3, 2, 2, 2, 101, 105, 5, 21, 11, 2, 102, 103, 5, 25, 13, 2, 103,
	104, 5, 21, 11, 2, 104, 106, 3, 2, 2, 2, 105, 102, 3, 2, 2, 2, 106, 107,
	3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 24, 3, 2,
	2, 2, 109, 110, 7, 48, 2, 2, 110, 26, 3, 2, 2, 2, 111, 112, 7, 63, 2, 2,
	112, 113, 7, 63, 2, 2, 113, 28, 3, 2, 2, 2, 114, 115, 7, 63, 2, 2, 115,
	30, 3, 2, 2, 2, 116, 117, 7, 64, 2, 2, 117, 32, 3, 2, 2, 2, 118, 119, 7,
	62, 2, 2, 119, 34, 3, 2, 2, 2, 120, 121, 7, 64, 2, 2, 121, 122, 7, 63,
	2, 2, 122, 36, 3, 2, 2, 2, 123, 124, 7, 62, 2, 2, 124, 125, 7, 63, 2, 2,
	125, 38, 3, 2, 2, 2, 126, 127, 7, 35, 2, 2, 127, 128, 7, 63, 2, 2, 128,
	40, 3, 2, 2, 2, 129, 131, 5, 9, 5, 2, 130, 129, 3, 2, 2, 2, 131, 132, 3,
	2, 2, 2, 132, 130, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 135, 3, 2, 2,
	2, 134, 130, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136,
	138, 7, 48, 2, 2, 137, 139, 5, 9, 5, 2, 138, 137, 3, 2, 2, 2, 139, 140,
	3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 42, 3, 2,
	2, 2, 142, 144, 9, 6, 2, 2, 143, 142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2,
	145, 143, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147,
	148, 8, 22, 2, 2, 148, 44, 3, 2, 2, 2, 15, 2, 59, 61, 67, 72, 77, 88, 98,
	107, 132, 134, 140, 145, 3, 8, 2, 2,
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
	"", "'('", "')'", "','", "", "", "", "", "'nil'", "", "", "'.'", "'=='",
	"'='", "'>'", "'<'", "'>='", "'<='", "'!='",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "StringArg", "FloatArg", "IntArg", "BoolArg", "NilArg",
	"FuncName", "MethodName", "Dot", "EQUALS", "ASSIGN", "GT", "LT", "GTE",
	"LTE", "NOTEQUALS", "Float", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "DIGIT", "StringArg", "FloatArg", "IntArg", "BoolArg",
	"NilArg", "FuncName", "MethodName", "Dot", "EQUALS", "ASSIGN", "GT", "LT",
	"GTE", "LTE", "NOTEQUALS", "Float", "WHITESPACE",
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
	BeanExprLexerBoolArg    = 7
	BeanExprLexerNilArg     = 8
	BeanExprLexerFuncName   = 9
	BeanExprLexerMethodName = 10
	BeanExprLexerDot        = 11
	BeanExprLexerEQUALS     = 12
	BeanExprLexerASSIGN     = 13
	BeanExprLexerGT         = 14
	BeanExprLexerLT         = 15
	BeanExprLexerGTE        = 16
	BeanExprLexerLTE        = 17
	BeanExprLexerNOTEQUALS  = 18
	BeanExprLexerFloat      = 19
	BeanExprLexerWHITESPACE = 20
)
