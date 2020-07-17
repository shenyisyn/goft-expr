// Code generated from E:/video/github/goft-expr/expr/src/g4\BeanExpr.g4 by ANTLR 4.8. DO NOT EDIT.

package BeanExprLib // BeanExpr
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 63, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 3, 2, 3, 2, 5, 2, 19, 10, 2, 3, 3, 3, 3, 3, 3, 5, 3, 24, 10, 3,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 5, 4, 31, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	5, 4, 37, 10, 4, 3, 4, 3, 4, 5, 4, 41, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 46,
	10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 52, 10, 5, 7, 5, 54, 10, 5, 12, 5,
	14, 5, 57, 11, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8, 10,
	12, 2, 4, 4, 2, 14, 14, 16, 20, 3, 2, 6, 10, 2, 66, 2, 18, 3, 2, 2, 2,
	4, 20, 3, 2, 2, 2, 6, 40, 3, 2, 2, 2, 8, 45, 3, 2, 2, 2, 10, 58, 3, 2,
	2, 2, 12, 60, 3, 2, 2, 2, 14, 19, 5, 6, 4, 2, 15, 16, 5, 4, 3, 2, 16, 17,
	7, 2, 2, 3, 17, 19, 3, 2, 2, 2, 18, 14, 3, 2, 2, 2, 18, 15, 3, 2, 2, 2,
	19, 3, 3, 2, 2, 2, 20, 21, 7, 12, 2, 2, 21, 23, 7, 3, 2, 2, 22, 24, 5,
	8, 5, 2, 23, 22, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25,
	26, 7, 4, 2, 2, 26, 5, 3, 2, 2, 2, 27, 28, 7, 11, 2, 2, 28, 30, 7, 3, 2,
	2, 29, 31, 5, 8, 5, 2, 30, 29, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 32,
	3, 2, 2, 2, 32, 41, 7, 4, 2, 2, 33, 34, 7, 11, 2, 2, 34, 36, 7, 3, 2, 2,
	35, 37, 5, 8, 5, 2, 36, 35, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 38, 3,
	2, 2, 2, 38, 39, 7, 4, 2, 2, 39, 41, 5, 10, 6, 2, 40, 27, 3, 2, 2, 2, 40,
	33, 3, 2, 2, 2, 41, 7, 3, 2, 2, 2, 42, 46, 5, 12, 7, 2, 43, 46, 5, 6, 4,
	2, 44, 46, 5, 4, 3, 2, 45, 42, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 45, 44,
	3, 2, 2, 2, 46, 55, 3, 2, 2, 2, 47, 51, 7, 5, 2, 2, 48, 52, 5, 12, 7, 2,
	49, 52, 5, 6, 4, 2, 50, 52, 5, 4, 3, 2, 51, 48, 3, 2, 2, 2, 51, 49, 3,
	2, 2, 2, 51, 50, 3, 2, 2, 2, 52, 54, 3, 2, 2, 2, 53, 47, 3, 2, 2, 2, 54,
	57, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 9, 3, 2, 2,
	2, 57, 55, 3, 2, 2, 2, 58, 59, 9, 2, 2, 2, 59, 11, 3, 2, 2, 2, 60, 61,
	9, 3, 2, 2, 61, 13, 3, 2, 2, 2, 10, 18, 23, 30, 36, 40, 45, 51, 55,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "','", "", "", "", "", "'nil'", "", "", "'.'", "'=='",
	"'='", "'>'", "'<'", "'>='", "'<='", "'!='",
}
var symbolicNames = []string{
	"", "", "", "", "StringArg", "FloatArg", "IntArg", "BoolArg", "NilArg",
	"FuncName", "MethodName", "Dot", "EQUALS", "ASSIGN", "GT", "LT", "GTE",
	"LTE", "NOTEQUALS", "Float", "WHITESPACE",
}

var ruleNames = []string{
	"start", "methodCall", "functionCall", "functionArgs", "comparisonOperator",
	"constant",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type BeanExprParser struct {
	*antlr.BaseParser
}

func NewBeanExprParser(input antlr.TokenStream) *BeanExprParser {
	this := new(BeanExprParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "BeanExpr.g4"

	return this
}

// BeanExprParser tokens.
const (
	BeanExprParserEOF        = antlr.TokenEOF
	BeanExprParserT__0       = 1
	BeanExprParserT__1       = 2
	BeanExprParserT__2       = 3
	BeanExprParserStringArg  = 4
	BeanExprParserFloatArg   = 5
	BeanExprParserIntArg     = 6
	BeanExprParserBoolArg    = 7
	BeanExprParserNilArg     = 8
	BeanExprParserFuncName   = 9
	BeanExprParserMethodName = 10
	BeanExprParserDot        = 11
	BeanExprParserEQUALS     = 12
	BeanExprParserASSIGN     = 13
	BeanExprParserGT         = 14
	BeanExprParserLT         = 15
	BeanExprParserGTE        = 16
	BeanExprParserLTE        = 17
	BeanExprParserNOTEQUALS  = 18
	BeanExprParserFloat      = 19
	BeanExprParserWHITESPACE = 20
)

// BeanExprParser rules.
const (
	BeanExprParserRULE_start              = 0
	BeanExprParserRULE_methodCall         = 1
	BeanExprParserRULE_functionCall       = 2
	BeanExprParserRULE_functionArgs       = 3
	BeanExprParserRULE_comparisonOperator = 4
	BeanExprParserRULE_constant           = 5
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanExprParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanExprParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *StartContext) MethodCall() IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(BeanExprParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeanExprListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeanExprListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *BeanExprParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BeanExprParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(16)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeanExprParserFuncName:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(12)
			p.FunctionCall()
		}

	case BeanExprParserMethodName:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(13)
			p.MethodCall()
		}
		{
			p.SetState(14)
			p.Match(BeanExprParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMethodCallContext is an interface to support dynamic dispatch.
type IMethodCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodCallContext differentiates from other interfaces.
	IsMethodCallContext()
}

type MethodCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodCallContext() *MethodCallContext {
	var p = new(MethodCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanExprParserRULE_methodCall
	return p
}

func (*MethodCallContext) IsMethodCallContext() {}

func NewMethodCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallContext {
	var p = new(MethodCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanExprParserRULE_methodCall

	return p
}

func (s *MethodCallContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallContext) MethodName() antlr.TerminalNode {
	return s.GetToken(BeanExprParserMethodName, 0)
}

func (s *MethodCallContext) FunctionArgs() IFunctionArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *MethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeanExprListener); ok {
		listenerT.EnterMethodCall(s)
	}
}

func (s *MethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeanExprListener); ok {
		listenerT.ExitMethodCall(s)
	}
}

func (p *BeanExprParser) MethodCall() (localctx IMethodCallContext) {
	localctx = NewMethodCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BeanExprParserRULE_methodCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Match(BeanExprParserMethodName)
	}
	{
		p.SetState(19)
		p.Match(BeanExprParserT__0)
	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeanExprParserStringArg)|(1<<BeanExprParserFloatArg)|(1<<BeanExprParserIntArg)|(1<<BeanExprParserBoolArg)|(1<<BeanExprParserNilArg)|(1<<BeanExprParserFuncName)|(1<<BeanExprParserMethodName))) != 0 {
		{
			p.SetState(20)
			p.FunctionArgs()
		}

	}
	{
		p.SetState(23)
		p.Match(BeanExprParserT__1)
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanExprParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanExprParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) CopyFrom(ctx *FunctionCallContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncCallContext struct {
	*FunctionCallContext
}

func NewFuncCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncCallContext {
	var p = new(FuncCallContext)

	p.FunctionCallContext = NewEmptyFunctionCallContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionCallContext))

	return p
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) FuncName() antlr.TerminalNode {
	return s.GetToken(BeanExprParserFuncName, 0)
}

func (s *FuncCallContext) FunctionArgs() IFunctionArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *FuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeanExprListener); ok {
		listenerT.EnterFuncCall(s)
	}
}

func (s *FuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeanExprListener); ok {
		listenerT.ExitFuncCall(s)
	}
}

type FuncCallWithCompareContext struct {
	*FunctionCallContext
}

func NewFuncCallWithCompareContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncCallWithCompareContext {
	var p = new(FuncCallWithCompareContext)

	p.FunctionCallContext = NewEmptyFunctionCallContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionCallContext))

	return p
}

func (s *FuncCallWithCompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallWithCompareContext) FuncName() antlr.TerminalNode {
	return s.GetToken(BeanExprParserFuncName, 0)
}

func (s *FuncCallWithCompareContext) ComparisonOperator() IComparisonOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *FuncCallWithCompareContext) FunctionArgs() IFunctionArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *FuncCallWithCompareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeanExprListener); ok {
		listenerT.EnterFuncCallWithCompare(s)
	}
}

func (s *FuncCallWithCompareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeanExprListener); ok {
		listenerT.ExitFuncCallWithCompare(s)
	}
}

func (p *BeanExprParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BeanExprParserRULE_functionCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFuncCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(25)
			p.Match(BeanExprParserFuncName)
		}
		{
			p.SetState(26)
			p.Match(BeanExprParserT__0)
		}
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeanExprParserStringArg)|(1<<BeanExprParserFloatArg)|(1<<BeanExprParserIntArg)|(1<<BeanExprParserBoolArg)|(1<<BeanExprParserNilArg)|(1<<BeanExprParserFuncName)|(1<<BeanExprParserMethodName))) != 0 {
			{
				p.SetState(27)
				p.FunctionArgs()
			}

		}
		{
			p.SetState(30)
			p.Match(BeanExprParserT__1)
		}

	case 2:
		localctx = NewFuncCallWithCompareContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.Match(BeanExprParserFuncName)
		}
		{
			p.SetState(32)
			p.Match(BeanExprParserT__0)
		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeanExprParserStringArg)|(1<<BeanExprParserFloatArg)|(1<<BeanExprParserIntArg)|(1<<BeanExprParserBoolArg)|(1<<BeanExprParserNilArg)|(1<<BeanExprParserFuncName)|(1<<BeanExprParserMethodName))) != 0 {
			{
				p.SetState(33)
				p.FunctionArgs()
			}

		}
		{
			p.SetState(36)
			p.Match(BeanExprParserT__1)
		}
		{
			p.SetState(37)
			p.ComparisonOperator()
		}

	}

	return localctx
}

// IFunctionArgsContext is an interface to support dynamic dispatch.
type IFunctionArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionArgsContext differentiates from other interfaces.
	IsFunctionArgsContext()
}

type FunctionArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgsContext() *FunctionArgsContext {
	var p = new(FunctionArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanExprParserRULE_functionArgs
	return p
}

func (*FunctionArgsContext) IsFunctionArgsContext() {}

func NewFunctionArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgsContext {
	var p = new(FunctionArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanExprParserRULE_functionArgs

	return p
}

func (s *FunctionArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgsContext) CopyFrom(ctx *FunctionArgsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FunctionArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncArgsContext struct {
	*FunctionArgsContext
}

func NewFuncArgsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncArgsContext {
	var p = new(FuncArgsContext)

	p.FunctionArgsContext = NewEmptyFunctionArgsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionArgsContext))

	return p
}

func (s *FuncArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgsContext) AllConstant() []IConstantContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstantContext)(nil)).Elem())
	var tst = make([]IConstantContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstantContext)
		}
	}

	return tst
}

func (s *FuncArgsContext) Constant(i int) IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FuncArgsContext) AllFunctionCall() []IFunctionCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem())
	var tst = make([]IFunctionCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionCallContext)
		}
	}

	return tst
}

func (s *FuncArgsContext) FunctionCall(i int) IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FuncArgsContext) AllMethodCall() []IMethodCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodCallContext)(nil)).Elem())
	var tst = make([]IMethodCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodCallContext)
		}
	}

	return tst
}

func (s *FuncArgsContext) MethodCall(i int) IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *FuncArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeanExprListener); ok {
		listenerT.EnterFuncArgs(s)
	}
}

func (s *FuncArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeanExprListener); ok {
		listenerT.ExitFuncArgs(s)
	}
}

func (p *BeanExprParser) FunctionArgs() (localctx IFunctionArgsContext) {
	localctx = NewFunctionArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BeanExprParserRULE_functionArgs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewFuncArgsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(43)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeanExprParserStringArg, BeanExprParserFloatArg, BeanExprParserIntArg, BeanExprParserBoolArg, BeanExprParserNilArg:
		{
			p.SetState(40)
			p.Constant()
		}

	case BeanExprParserFuncName:
		{
			p.SetState(41)
			p.FunctionCall()
		}

	case BeanExprParserMethodName:
		{
			p.SetState(42)
			p.MethodCall()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BeanExprParserT__2 {
		{
			p.SetState(45)
			p.Match(BeanExprParserT__2)
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case BeanExprParserStringArg, BeanExprParserFloatArg, BeanExprParserIntArg, BeanExprParserBoolArg, BeanExprParserNilArg:
			{
				p.SetState(46)
				p.Constant()
			}

		case BeanExprParserFuncName:
			{
				p.SetState(47)
				p.FunctionCall()
			}

		case BeanExprParserMethodName:
			{
				p.SetState(48)
				p.MethodCall()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanExprParserRULE_comparisonOperator
	return p
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanExprParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(BeanExprParserGT, 0)
}

func (s *ComparisonOperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(BeanExprParserLT, 0)
}

func (s *ComparisonOperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(BeanExprParserGTE, 0)
}

func (s *ComparisonOperatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(BeanExprParserLTE, 0)
}

func (s *ComparisonOperatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(BeanExprParserEQUALS, 0)
}

func (s *ComparisonOperatorContext) NOTEQUALS() antlr.TerminalNode {
	return s.GetToken(BeanExprParserNOTEQUALS, 0)
}

func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeanExprListener); ok {
		listenerT.EnterComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeanExprListener); ok {
		listenerT.ExitComparisonOperator(s)
	}
}

func (p *BeanExprParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BeanExprParserRULE_comparisonOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeanExprParserEQUALS)|(1<<BeanExprParserGT)|(1<<BeanExprParserLT)|(1<<BeanExprParserGTE)|(1<<BeanExprParserLTE)|(1<<BeanExprParserNOTEQUALS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanExprParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanExprParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) StringArg() antlr.TerminalNode {
	return s.GetToken(BeanExprParserStringArg, 0)
}

func (s *ConstantContext) FloatArg() antlr.TerminalNode {
	return s.GetToken(BeanExprParserFloatArg, 0)
}

func (s *ConstantContext) IntArg() antlr.TerminalNode {
	return s.GetToken(BeanExprParserIntArg, 0)
}

func (s *ConstantContext) BoolArg() antlr.TerminalNode {
	return s.GetToken(BeanExprParserBoolArg, 0)
}

func (s *ConstantContext) NilArg() antlr.TerminalNode {
	return s.GetToken(BeanExprParserNilArg, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeanExprListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeanExprListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *BeanExprParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BeanExprParserRULE_constant)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeanExprParserStringArg)|(1<<BeanExprParserFloatArg)|(1<<BeanExprParserIntArg)|(1<<BeanExprParserBoolArg)|(1<<BeanExprParserNilArg))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
