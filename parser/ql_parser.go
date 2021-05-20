// Code generated from parser/QlParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // QlParser

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 34, 91, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 2, 5, 2, 32, 10, 2, 3, 2, 5, 2, 35, 10,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 43, 10, 3, 12, 3, 14, 3, 46,
	11, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 7, 5, 7, 60, 10, 7, 3, 8, 3, 8, 3, 8, 5, 8, 65, 10, 8, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 74, 10, 10, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 12, 3, 12, 5, 12, 82, 10, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3,
	14, 5, 14, 89, 10, 14, 3, 14, 2, 2, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 2, 2, 2, 85, 2, 28, 3, 2, 2, 2, 4, 38, 3, 2, 2, 2, 6, 47,
	3, 2, 2, 2, 8, 50, 3, 2, 2, 2, 10, 53, 3, 2, 2, 2, 12, 56, 3, 2, 2, 2,
	14, 61, 3, 2, 2, 2, 16, 66, 3, 2, 2, 2, 18, 68, 3, 2, 2, 2, 20, 77, 3,
	2, 2, 2, 22, 81, 3, 2, 2, 2, 24, 83, 3, 2, 2, 2, 26, 88, 3, 2, 2, 2, 28,
	29, 5, 4, 3, 2, 29, 31, 5, 6, 4, 2, 30, 32, 5, 8, 5, 2, 31, 30, 3, 2, 2,
	2, 31, 32, 3, 2, 2, 2, 32, 34, 3, 2, 2, 2, 33, 35, 5, 10, 6, 2, 34, 33,
	3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 37, 7, 2, 2, 3,
	37, 3, 3, 2, 2, 2, 38, 39, 7, 13, 2, 2, 39, 44, 5, 26, 14, 2, 40, 41, 7,
	3, 2, 2, 41, 43, 5, 26, 14, 2, 42, 40, 3, 2, 2, 2, 43, 46, 3, 2, 2, 2,
	44, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 5, 3, 2, 2, 2, 46, 44, 3, 2,
	2, 2, 47, 48, 7, 15, 2, 2, 48, 49, 7, 27, 2, 2, 49, 7, 3, 2, 2, 2, 50,
	51, 7, 16, 2, 2, 51, 52, 5, 22, 12, 2, 52, 9, 3, 2, 2, 2, 53, 54, 7, 17,
	2, 2, 54, 55, 5, 22, 12, 2, 55, 11, 3, 2, 2, 2, 56, 59, 5, 18, 10, 2, 57,
	58, 7, 14, 2, 2, 58, 60, 5, 16, 9, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2,
	2, 2, 60, 13, 3, 2, 2, 2, 61, 64, 7, 27, 2, 2, 62, 63, 7, 14, 2, 2, 63,
	65, 5, 16, 9, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 15, 3, 2,
	2, 2, 66, 67, 7, 27, 2, 2, 67, 17, 3, 2, 2, 2, 68, 69, 7, 5, 2, 2, 69,
	70, 5, 14, 8, 2, 70, 73, 7, 7, 2, 2, 71, 74, 5, 14, 8, 2, 72, 74, 7, 28,
	2, 2, 73, 71, 3, 2, 2, 2, 73, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 76,
	7, 6, 2, 2, 76, 19, 3, 2, 2, 2, 77, 78, 7, 27, 2, 2, 78, 21, 3, 2, 2, 2,
	79, 82, 7, 33, 2, 2, 80, 82, 5, 24, 13, 2, 81, 79, 3, 2, 2, 2, 81, 80,
	3, 2, 2, 2, 82, 23, 3, 2, 2, 2, 83, 84, 7, 28, 2, 2, 84, 85, 7, 18, 2,
	2, 85, 25, 3, 2, 2, 2, 86, 89, 5, 14, 8, 2, 87, 89, 5, 12, 7, 2, 88, 86,
	3, 2, 2, 2, 88, 87, 3, 2, 2, 2, 89, 27, 3, 2, 2, 2, 10, 31, 34, 44, 59,
	64, 73, 81, 88,
}
var literalNames = []string{
	"", "','", "'.'", "'('", "')'", "", "'+'", "'-'", "'/'", "'%'", "'*'",
	"", "", "", "", "", "", "'s'", "'min'", "'h'", "'d'", "'w'", "'m'", "'q'",
	"'y'",
}
var symbolicNames = []string{
	"", "COMMA", "DOT", "LR_BRACKET", "RR_BRACKET", "ARITHMETIC_OPERATOR",
	"PLUS", "MINUS", "DIV", "MOD", "MUL", "SHOW", "AS", "FROM", "SINCE", "UNTIL",
	"TIMESPAN_UNIT", "SECONDS", "MINUTES", "HOURS", "DAYS", "WEEKS", "MONTHS",
	"QUARTERS", "YEARS", "IDENTIFIER", "SIGNED_NUMERIC_LITERAL", "NUMERIC_LITERAL",
	"YEAR", "MONTH", "DAY", "DATE_LITERAL", "SPACES",
}

var ruleNames = []string{
	"query", "show_stmt", "from_stmt", "since_stmt", "until_stmt", "operation_column",
	"result_column", "column_alias", "arithmetic_operation", "expr", "date_spec",
	"relative_date", "column_stmt",
}

type QlParser struct {
	*antlr.BaseParser
}

// NewQlParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *QlParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewQlParser(input antlr.TokenStream) *QlParser {
	this := new(QlParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "QlParser.g4"

	return this
}

// QlParser tokens.
const (
	QlParserEOF                    = antlr.TokenEOF
	QlParserCOMMA                  = 1
	QlParserDOT                    = 2
	QlParserLR_BRACKET             = 3
	QlParserRR_BRACKET             = 4
	QlParserARITHMETIC_OPERATOR    = 5
	QlParserPLUS                   = 6
	QlParserMINUS                  = 7
	QlParserDIV                    = 8
	QlParserMOD                    = 9
	QlParserMUL                    = 10
	QlParserSHOW                   = 11
	QlParserAS                     = 12
	QlParserFROM                   = 13
	QlParserSINCE                  = 14
	QlParserUNTIL                  = 15
	QlParserTIMESPAN_UNIT          = 16
	QlParserSECONDS                = 17
	QlParserMINUTES                = 18
	QlParserHOURS                  = 19
	QlParserDAYS                   = 20
	QlParserWEEKS                  = 21
	QlParserMONTHS                 = 22
	QlParserQUARTERS               = 23
	QlParserYEARS                  = 24
	QlParserIDENTIFIER             = 25
	QlParserSIGNED_NUMERIC_LITERAL = 26
	QlParserNUMERIC_LITERAL        = 27
	QlParserYEAR                   = 28
	QlParserMONTH                  = 29
	QlParserDAY                    = 30
	QlParserDATE_LITERAL           = 31
	QlParserSPACES                 = 32
)

// QlParser rules.
const (
	QlParserRULE_query                = 0
	QlParserRULE_show_stmt            = 1
	QlParserRULE_from_stmt            = 2
	QlParserRULE_since_stmt           = 3
	QlParserRULE_until_stmt           = 4
	QlParserRULE_operation_column     = 5
	QlParserRULE_result_column        = 6
	QlParserRULE_column_alias         = 7
	QlParserRULE_arithmetic_operation = 8
	QlParserRULE_expr                 = 9
	QlParserRULE_date_spec            = 10
	QlParserRULE_relative_date        = 11
	QlParserRULE_column_stmt          = 12
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QlParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QlParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) Show_stmt() IShow_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShow_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShow_stmtContext)
}

func (s *QueryContext) From_stmt() IFrom_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFrom_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFrom_stmtContext)
}

func (s *QueryContext) EOF() antlr.TerminalNode {
	return s.GetToken(QlParserEOF, 0)
}

func (s *QueryContext) Since_stmt() ISince_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISince_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISince_stmtContext)
}

func (s *QueryContext) Until_stmt() IUntil_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUntil_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUntil_stmtContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QlParserVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QlParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QlParserRULE_query)
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
		p.SetState(26)
		p.Show_stmt()
	}
	{
		p.SetState(27)
		p.From_stmt()
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QlParserSINCE {
		{
			p.SetState(28)
			p.Since_stmt()
		}

	}
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QlParserUNTIL {
		{
			p.SetState(31)
			p.Until_stmt()
		}

	}
	{
		p.SetState(34)
		p.Match(QlParserEOF)
	}

	return localctx
}

// IShow_stmtContext is an interface to support dynamic dispatch.
type IShow_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShow_stmtContext differentiates from other interfaces.
	IsShow_stmtContext()
}

type Show_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShow_stmtContext() *Show_stmtContext {
	var p = new(Show_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QlParserRULE_show_stmt
	return p
}

func (*Show_stmtContext) IsShow_stmtContext() {}

func NewShow_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Show_stmtContext {
	var p = new(Show_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QlParserRULE_show_stmt

	return p
}

func (s *Show_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Show_stmtContext) SHOW() antlr.TerminalNode {
	return s.GetToken(QlParserSHOW, 0)
}

func (s *Show_stmtContext) AllColumn_stmt() []IColumn_stmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumn_stmtContext)(nil)).Elem())
	var tst = make([]IColumn_stmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumn_stmtContext)
		}
	}

	return tst
}

func (s *Show_stmtContext) Column_stmt(i int) IColumn_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_stmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumn_stmtContext)
}

func (s *Show_stmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QlParserCOMMA)
}

func (s *Show_stmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QlParserCOMMA, i)
}

func (s *Show_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Show_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Show_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.EnterShow_stmt(s)
	}
}

func (s *Show_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.ExitShow_stmt(s)
	}
}

func (s *Show_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QlParserVisitor:
		return t.VisitShow_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QlParser) Show_stmt() (localctx IShow_stmtContext) {
	localctx = NewShow_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, QlParserRULE_show_stmt)
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
		p.SetState(36)
		p.Match(QlParserSHOW)
	}
	{
		p.SetState(37)
		p.Column_stmt()
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QlParserCOMMA {
		{
			p.SetState(38)
			p.Match(QlParserCOMMA)
		}
		{
			p.SetState(39)
			p.Column_stmt()
		}

		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFrom_stmtContext is an interface to support dynamic dispatch.
type IFrom_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFrom_stmtContext differentiates from other interfaces.
	IsFrom_stmtContext()
}

type From_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrom_stmtContext() *From_stmtContext {
	var p = new(From_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QlParserRULE_from_stmt
	return p
}

func (*From_stmtContext) IsFrom_stmtContext() {}

func NewFrom_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *From_stmtContext {
	var p = new(From_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QlParserRULE_from_stmt

	return p
}

func (s *From_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *From_stmtContext) FROM() antlr.TerminalNode {
	return s.GetToken(QlParserFROM, 0)
}

func (s *From_stmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(QlParserIDENTIFIER, 0)
}

func (s *From_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *From_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *From_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.EnterFrom_stmt(s)
	}
}

func (s *From_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.ExitFrom_stmt(s)
	}
}

func (s *From_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QlParserVisitor:
		return t.VisitFrom_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QlParser) From_stmt() (localctx IFrom_stmtContext) {
	localctx = NewFrom_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QlParserRULE_from_stmt)

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
		p.SetState(45)
		p.Match(QlParserFROM)
	}
	{
		p.SetState(46)
		p.Match(QlParserIDENTIFIER)
	}

	return localctx
}

// ISince_stmtContext is an interface to support dynamic dispatch.
type ISince_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSince_stmtContext differentiates from other interfaces.
	IsSince_stmtContext()
}

type Since_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySince_stmtContext() *Since_stmtContext {
	var p = new(Since_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QlParserRULE_since_stmt
	return p
}

func (*Since_stmtContext) IsSince_stmtContext() {}

func NewSince_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Since_stmtContext {
	var p = new(Since_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QlParserRULE_since_stmt

	return p
}

func (s *Since_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Since_stmtContext) SINCE() antlr.TerminalNode {
	return s.GetToken(QlParserSINCE, 0)
}

func (s *Since_stmtContext) Date_spec() IDate_specContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDate_specContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDate_specContext)
}

func (s *Since_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Since_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Since_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.EnterSince_stmt(s)
	}
}

func (s *Since_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.ExitSince_stmt(s)
	}
}

func (s *Since_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QlParserVisitor:
		return t.VisitSince_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QlParser) Since_stmt() (localctx ISince_stmtContext) {
	localctx = NewSince_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, QlParserRULE_since_stmt)

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
		p.SetState(48)
		p.Match(QlParserSINCE)
	}
	{
		p.SetState(49)
		p.Date_spec()
	}

	return localctx
}

// IUntil_stmtContext is an interface to support dynamic dispatch.
type IUntil_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUntil_stmtContext differentiates from other interfaces.
	IsUntil_stmtContext()
}

type Until_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUntil_stmtContext() *Until_stmtContext {
	var p = new(Until_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QlParserRULE_until_stmt
	return p
}

func (*Until_stmtContext) IsUntil_stmtContext() {}

func NewUntil_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Until_stmtContext {
	var p = new(Until_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QlParserRULE_until_stmt

	return p
}

func (s *Until_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Until_stmtContext) UNTIL() antlr.TerminalNode {
	return s.GetToken(QlParserUNTIL, 0)
}

func (s *Until_stmtContext) Date_spec() IDate_specContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDate_specContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDate_specContext)
}

func (s *Until_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Until_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Until_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.EnterUntil_stmt(s)
	}
}

func (s *Until_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.ExitUntil_stmt(s)
	}
}

func (s *Until_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QlParserVisitor:
		return t.VisitUntil_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QlParser) Until_stmt() (localctx IUntil_stmtContext) {
	localctx = NewUntil_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, QlParserRULE_until_stmt)

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
		p.SetState(51)
		p.Match(QlParserUNTIL)
	}
	{
		p.SetState(52)
		p.Date_spec()
	}

	return localctx
}

// IOperation_columnContext is an interface to support dynamic dispatch.
type IOperation_columnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperation_columnContext differentiates from other interfaces.
	IsOperation_columnContext()
}

type Operation_columnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperation_columnContext() *Operation_columnContext {
	var p = new(Operation_columnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QlParserRULE_operation_column
	return p
}

func (*Operation_columnContext) IsOperation_columnContext() {}

func NewOperation_columnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Operation_columnContext {
	var p = new(Operation_columnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QlParserRULE_operation_column

	return p
}

func (s *Operation_columnContext) GetParser() antlr.Parser { return s.parser }

func (s *Operation_columnContext) Arithmetic_operation() IArithmetic_operationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmetic_operationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithmetic_operationContext)
}

func (s *Operation_columnContext) AS() antlr.TerminalNode {
	return s.GetToken(QlParserAS, 0)
}

func (s *Operation_columnContext) Column_alias() IColumn_aliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_aliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumn_aliasContext)
}

func (s *Operation_columnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Operation_columnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Operation_columnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.EnterOperation_column(s)
	}
}

func (s *Operation_columnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.ExitOperation_column(s)
	}
}

func (s *Operation_columnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QlParserVisitor:
		return t.VisitOperation_column(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QlParser) Operation_column() (localctx IOperation_columnContext) {
	localctx = NewOperation_columnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, QlParserRULE_operation_column)
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
		p.SetState(54)
		p.Arithmetic_operation()
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QlParserAS {
		{
			p.SetState(55)
			p.Match(QlParserAS)
		}
		{
			p.SetState(56)
			p.Column_alias()
		}

	}

	return localctx
}

// IResult_columnContext is an interface to support dynamic dispatch.
type IResult_columnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResult_columnContext differentiates from other interfaces.
	IsResult_columnContext()
}

type Result_columnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResult_columnContext() *Result_columnContext {
	var p = new(Result_columnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QlParserRULE_result_column
	return p
}

func (*Result_columnContext) IsResult_columnContext() {}

func NewResult_columnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Result_columnContext {
	var p = new(Result_columnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QlParserRULE_result_column

	return p
}

func (s *Result_columnContext) GetParser() antlr.Parser { return s.parser }

func (s *Result_columnContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(QlParserIDENTIFIER, 0)
}

func (s *Result_columnContext) AS() antlr.TerminalNode {
	return s.GetToken(QlParserAS, 0)
}

func (s *Result_columnContext) Column_alias() IColumn_aliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_aliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumn_aliasContext)
}

func (s *Result_columnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Result_columnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Result_columnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.EnterResult_column(s)
	}
}

func (s *Result_columnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.ExitResult_column(s)
	}
}

func (s *Result_columnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QlParserVisitor:
		return t.VisitResult_column(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QlParser) Result_column() (localctx IResult_columnContext) {
	localctx = NewResult_columnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, QlParserRULE_result_column)
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
		p.SetState(59)
		p.Match(QlParserIDENTIFIER)
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QlParserAS {
		{
			p.SetState(60)
			p.Match(QlParserAS)
		}
		{
			p.SetState(61)
			p.Column_alias()
		}

	}

	return localctx
}

// IColumn_aliasContext is an interface to support dynamic dispatch.
type IColumn_aliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumn_aliasContext differentiates from other interfaces.
	IsColumn_aliasContext()
}

type Column_aliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_aliasContext() *Column_aliasContext {
	var p = new(Column_aliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QlParserRULE_column_alias
	return p
}

func (*Column_aliasContext) IsColumn_aliasContext() {}

func NewColumn_aliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_aliasContext {
	var p = new(Column_aliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QlParserRULE_column_alias

	return p
}

func (s *Column_aliasContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_aliasContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(QlParserIDENTIFIER, 0)
}

func (s *Column_aliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_aliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_aliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.EnterColumn_alias(s)
	}
}

func (s *Column_aliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.ExitColumn_alias(s)
	}
}

func (s *Column_aliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QlParserVisitor:
		return t.VisitColumn_alias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QlParser) Column_alias() (localctx IColumn_aliasContext) {
	localctx = NewColumn_aliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, QlParserRULE_column_alias)

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
		p.SetState(64)
		p.Match(QlParserIDENTIFIER)
	}

	return localctx
}

// IArithmetic_operationContext is an interface to support dynamic dispatch.
type IArithmetic_operationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArithmetic_operationContext differentiates from other interfaces.
	IsArithmetic_operationContext()
}

type Arithmetic_operationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithmetic_operationContext() *Arithmetic_operationContext {
	var p = new(Arithmetic_operationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QlParserRULE_arithmetic_operation
	return p
}

func (*Arithmetic_operationContext) IsArithmetic_operationContext() {}

func NewArithmetic_operationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arithmetic_operationContext {
	var p = new(Arithmetic_operationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QlParserRULE_arithmetic_operation

	return p
}

func (s *Arithmetic_operationContext) GetParser() antlr.Parser { return s.parser }

func (s *Arithmetic_operationContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(QlParserLR_BRACKET, 0)
}

func (s *Arithmetic_operationContext) AllResult_column() []IResult_columnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IResult_columnContext)(nil)).Elem())
	var tst = make([]IResult_columnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IResult_columnContext)
		}
	}

	return tst
}

func (s *Arithmetic_operationContext) Result_column(i int) IResult_columnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResult_columnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IResult_columnContext)
}

func (s *Arithmetic_operationContext) ARITHMETIC_OPERATOR() antlr.TerminalNode {
	return s.GetToken(QlParserARITHMETIC_OPERATOR, 0)
}

func (s *Arithmetic_operationContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(QlParserRR_BRACKET, 0)
}

func (s *Arithmetic_operationContext) SIGNED_NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(QlParserSIGNED_NUMERIC_LITERAL, 0)
}

func (s *Arithmetic_operationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arithmetic_operationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arithmetic_operationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.EnterArithmetic_operation(s)
	}
}

func (s *Arithmetic_operationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.ExitArithmetic_operation(s)
	}
}

func (s *Arithmetic_operationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QlParserVisitor:
		return t.VisitArithmetic_operation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QlParser) Arithmetic_operation() (localctx IArithmetic_operationContext) {
	localctx = NewArithmetic_operationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, QlParserRULE_arithmetic_operation)

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
		p.SetState(66)
		p.Match(QlParserLR_BRACKET)
	}
	{
		p.SetState(67)
		p.Result_column()
	}
	{
		p.SetState(68)
		p.Match(QlParserARITHMETIC_OPERATOR)
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QlParserIDENTIFIER:
		{
			p.SetState(69)
			p.Result_column()
		}

	case QlParserSIGNED_NUMERIC_LITERAL:
		{
			p.SetState(70)
			p.Match(QlParserSIGNED_NUMERIC_LITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(73)
		p.Match(QlParserRR_BRACKET)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QlParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QlParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(QlParserIDENTIFIER, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QlParserVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QlParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, QlParserRULE_expr)

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
		p.SetState(75)
		p.Match(QlParserIDENTIFIER)
	}

	return localctx
}

// IDate_specContext is an interface to support dynamic dispatch.
type IDate_specContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDate_specContext differentiates from other interfaces.
	IsDate_specContext()
}

type Date_specContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDate_specContext() *Date_specContext {
	var p = new(Date_specContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QlParserRULE_date_spec
	return p
}

func (*Date_specContext) IsDate_specContext() {}

func NewDate_specContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Date_specContext {
	var p = new(Date_specContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QlParserRULE_date_spec

	return p
}

func (s *Date_specContext) GetParser() antlr.Parser { return s.parser }

func (s *Date_specContext) DATE_LITERAL() antlr.TerminalNode {
	return s.GetToken(QlParserDATE_LITERAL, 0)
}

func (s *Date_specContext) Relative_date() IRelative_dateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelative_dateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelative_dateContext)
}

func (s *Date_specContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Date_specContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Date_specContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.EnterDate_spec(s)
	}
}

func (s *Date_specContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.ExitDate_spec(s)
	}
}

func (s *Date_specContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QlParserVisitor:
		return t.VisitDate_spec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QlParser) Date_spec() (localctx IDate_specContext) {
	localctx = NewDate_specContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, QlParserRULE_date_spec)

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

	p.SetState(79)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QlParserDATE_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.Match(QlParserDATE_LITERAL)
		}

	case QlParserSIGNED_NUMERIC_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(78)
			p.Relative_date()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRelative_dateContext is an interface to support dynamic dispatch.
type IRelative_dateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelative_dateContext differentiates from other interfaces.
	IsRelative_dateContext()
}

type Relative_dateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelative_dateContext() *Relative_dateContext {
	var p = new(Relative_dateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QlParserRULE_relative_date
	return p
}

func (*Relative_dateContext) IsRelative_dateContext() {}

func NewRelative_dateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Relative_dateContext {
	var p = new(Relative_dateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QlParserRULE_relative_date

	return p
}

func (s *Relative_dateContext) GetParser() antlr.Parser { return s.parser }

func (s *Relative_dateContext) SIGNED_NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(QlParserSIGNED_NUMERIC_LITERAL, 0)
}

func (s *Relative_dateContext) TIMESPAN_UNIT() antlr.TerminalNode {
	return s.GetToken(QlParserTIMESPAN_UNIT, 0)
}

func (s *Relative_dateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Relative_dateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Relative_dateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.EnterRelative_date(s)
	}
}

func (s *Relative_dateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.ExitRelative_date(s)
	}
}

func (s *Relative_dateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QlParserVisitor:
		return t.VisitRelative_date(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QlParser) Relative_date() (localctx IRelative_dateContext) {
	localctx = NewRelative_dateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, QlParserRULE_relative_date)

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
		p.SetState(81)
		p.Match(QlParserSIGNED_NUMERIC_LITERAL)
	}
	{
		p.SetState(82)
		p.Match(QlParserTIMESPAN_UNIT)
	}

	return localctx
}

// IColumn_stmtContext is an interface to support dynamic dispatch.
type IColumn_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumn_stmtContext differentiates from other interfaces.
	IsColumn_stmtContext()
}

type Column_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_stmtContext() *Column_stmtContext {
	var p = new(Column_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QlParserRULE_column_stmt
	return p
}

func (*Column_stmtContext) IsColumn_stmtContext() {}

func NewColumn_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_stmtContext {
	var p = new(Column_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QlParserRULE_column_stmt

	return p
}

func (s *Column_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_stmtContext) Result_column() IResult_columnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResult_columnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IResult_columnContext)
}

func (s *Column_stmtContext) Operation_column() IOperation_columnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperation_columnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperation_columnContext)
}

func (s *Column_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.EnterColumn_stmt(s)
	}
}

func (s *Column_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QlParserListener); ok {
		listenerT.ExitColumn_stmt(s)
	}
}

func (s *Column_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QlParserVisitor:
		return t.VisitColumn_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QlParser) Column_stmt() (localctx IColumn_stmtContext) {
	localctx = NewColumn_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, QlParserRULE_column_stmt)

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

	p.SetState(86)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Result_column()
		}

	case QlParserLR_BRACKET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Operation_column()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
