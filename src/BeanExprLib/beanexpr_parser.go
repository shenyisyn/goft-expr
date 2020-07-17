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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 39, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3, 2, 5,
	2, 15, 10, 2, 3, 3, 3, 3, 3, 3, 5, 3, 20, 10, 3, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 4, 5, 4, 27, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 7, 5, 34, 10, 5, 12,
	5, 14, 5, 37, 11, 5, 3, 5, 2, 2, 6, 2, 4, 6, 8, 2, 3, 3, 2, 6, 8, 2, 38,
	2, 14, 3, 2, 2, 2, 4, 16, 3, 2, 2, 2, 6, 23, 3, 2, 2, 2, 8, 30, 3, 2, 2,
	2, 10, 15, 5, 6, 4, 2, 11, 12, 5, 4, 3, 2, 12, 13, 7, 2, 2, 3, 13, 15,
	3, 2, 2, 2, 14, 10, 3, 2, 2, 2, 14, 11, 3, 2, 2, 2, 15, 3, 3, 2, 2, 2,
	16, 17, 7, 10, 2, 2, 17, 19, 7, 3, 2, 2, 18, 20, 5, 8, 5, 2, 19, 18, 3,
	2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 22, 7, 4, 2, 2, 22,
	5, 3, 2, 2, 2, 23, 24, 7, 9, 2, 2, 24, 26, 7, 3, 2, 2, 25, 27, 5, 8, 5,
	2, 26, 25, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 29,
	7, 4, 2, 2, 29, 7, 3, 2, 2, 2, 30, 35, 9, 2, 2, 2, 31, 32, 7, 5, 2, 2,
	32, 34, 9, 2, 2, 2, 33, 31, 3, 2, 2, 2, 34, 37, 3, 2, 2, 2, 35, 33, 3,
	2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 9, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 6,
	14, 19, 26, 35,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "','", "", "", "", "", "", "'.'",
}
var symbolicNames = []string{
	"", "", "", "", "StringArg", "FloatArg", "IntArg", "FuncName", "MethodName",
	"Dot", "Float", "WHITESPACE",
}

var ruleNames = []string{
	"start", "methodCall", "functionCall", "functionArgs",
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
	BeanExprParserFuncName   = 7
	BeanExprParserMethodName = 8
	BeanExprParserDot        = 9
	BeanExprParserFloat      = 10
	BeanExprParserWHITESPACE = 11
)

// BeanExprParser rules.
const (
	BeanExprParserRULE_start        = 0
	BeanExprParserRULE_methodCall   = 1
	BeanExprParserRULE_functionCall = 2
	BeanExprParserRULE_functionArgs = 3
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

	p.SetState(12)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeanExprParserFuncName:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(8)
			p.FunctionCall()
		}

	case BeanExprParserMethodName:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(9)
			p.MethodCall()
		}
		{
			p.SetState(10)
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
		p.SetState(14)
		p.Match(BeanExprParserMethodName)
	}
	{
		p.SetState(15)
		p.Match(BeanExprParserT__0)
	}
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeanExprParserStringArg)|(1<<BeanExprParserFloatArg)|(1<<BeanExprParserIntArg))) != 0 {
		{
			p.SetState(16)
			p.FunctionArgs()
		}

	}
	{
		p.SetState(19)
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

	localctx = NewFuncCallContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(21)
		p.Match(BeanExprParserFuncName)
	}
	{
		p.SetState(22)
		p.Match(BeanExprParserT__0)
	}
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeanExprParserStringArg)|(1<<BeanExprParserFloatArg)|(1<<BeanExprParserIntArg))) != 0 {
		{
			p.SetState(23)
			p.FunctionArgs()
		}

	}
	{
		p.SetState(26)
		p.Match(BeanExprParserT__1)
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

func (s *FuncArgsContext) AllFloatArg() []antlr.TerminalNode {
	return s.GetTokens(BeanExprParserFloatArg)
}

func (s *FuncArgsContext) FloatArg(i int) antlr.TerminalNode {
	return s.GetToken(BeanExprParserFloatArg, i)
}

func (s *FuncArgsContext) AllStringArg() []antlr.TerminalNode {
	return s.GetTokens(BeanExprParserStringArg)
}

func (s *FuncArgsContext) StringArg(i int) antlr.TerminalNode {
	return s.GetToken(BeanExprParserStringArg, i)
}

func (s *FuncArgsContext) AllIntArg() []antlr.TerminalNode {
	return s.GetTokens(BeanExprParserIntArg)
}

func (s *FuncArgsContext) IntArg(i int) antlr.TerminalNode {
	return s.GetToken(BeanExprParserIntArg, i)
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
	{
		p.SetState(28)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeanExprParserStringArg)|(1<<BeanExprParserFloatArg)|(1<<BeanExprParserIntArg))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BeanExprParserT__2 {
		{
			p.SetState(29)
			p.Match(BeanExprParserT__2)
		}
		{
			p.SetState(30)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeanExprParserStringArg)|(1<<BeanExprParserFloatArg)|(1<<BeanExprParserIntArg))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
