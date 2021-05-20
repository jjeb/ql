package main

import (
	"bytes"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ql/parser"
)

func main() {
	getCompletionItems("show event_count as example from visits")
}

func returnString(value string, length int) {}

//export getCompletionItems
func getCompletionItems(input string) {
 result := Parse(input)
 returnString(result.String(), len(result.String()))
}

func NewQlVisitorImpl() *QlVisitorImpl {
	return &QlVisitorImpl{
		ParseTreeVisitor: &parser.BaseQlParserVisitor{},
	}
}

// Parse parses the input query string into a Query AST.
func Parse(s string) Query {
	is := antlr.NewInputStream(s)
	lexer := parser.NewQlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewQlParser(stream)

	visitor := NewQlVisitorImpl()
	q := visitor.Visit(p.Query()).(Query)
	return q
}

/***********************************
	Visitor implementaion
************************************/
type QlVisitorImpl struct {
	antlr.ParseTreeVisitor
	query Query
}

type Query struct {
	Fields []Field
	Source string
	Since string
	Until  string
}

func (q *Query) String() string {
	var buf bytes.Buffer
	f := q.Fields
	fmt.Fprintf(&buf, "SHOW %s", f[0].String())

	for _, field := range f[1:] {
		fmt.Fprintf(&buf, ", %s", field.String())
	}

	fmt.Fprintf(&buf, " FROM %s", q.Source)

	if q.Since != "" {
		fmt.Fprintf(&buf, " SINCE %s", q.Since)
	}

	if q.Until != "" {
		fmt.Fprintf(&buf, " UNTIL %s", q.Until)
	}

	return buf.String()
}

type Field struct {
	Expr  VarRef
	Alias string
}

func (f *Field) String() string {
	s := f.Expr.String()
	if f.Alias == "" || f.Alias == s {
		return s
	}
	return fmt.Sprintf("%s AS %s", f.Expr.String(), f.Alias)
}

type VarRef struct {
	Identity string
}

func (v *VarRef) String() string {
	return v.Identity
}

func (v *QlVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.QueryContext:
		return val.Accept(v)
	}
	return nil
}

func (v *QlVisitorImpl) VisitQuery(ctx *parser.QueryContext) interface{} {
	v.query = Query{}
	v.query.Fields = ctx.Show_stmt().Accept(v).([]Field)
	v.query.Source = ctx.From_stmt().Accept(v).(string)
	if ctx.Since_stmt() != nil {
		v.query.Since = ctx.Since_stmt().Accept(v).(string)
	}
	if ctx.Until_stmt() != nil {
		v.query.Until = ctx.Until_stmt().Accept(v).(string)
	}
	return v.query
}

func (v *QlVisitorImpl) VisitShow_stmt(ctx *parser.Show_stmtContext) interface{} {
	var fields = make([]Field, 0)

	for _, column := range ctx.AllColumn_stmt() {
		field := column.Accept(v).(Field)
		fields = append(fields, field)
	}

	return fields
}

func (v *QlVisitorImpl) VisitFrom_stmt(ctx *parser.From_stmtContext) interface{} {
	return ctx.IDENTIFIER().GetText()
}

func (v *QlVisitorImpl) VisitResult_column(ctx *parser.Result_columnContext) interface{} {
	field := Field{}

	field.Expr = VarRef{Identity: ctx.IDENTIFIER().GetText()}

	if ctx.Column_alias() != nil {
		field.Alias = ctx.Column_alias().Accept(v).(string)
	}
	return field
}

func (v *QlVisitorImpl) VisitColumn_alias(ctx *parser.Column_aliasContext) interface{} {
	return ctx.IDENTIFIER().GetText()
}

func (v *QlVisitorImpl) VisitExpr(ctx *parser.ExprContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *QlVisitorImpl) VisitSince_stmt(ctx *parser.Since_stmtContext) interface{} {
	return ctx.Date_spec().Accept(v).(string)
}

func (v *QlVisitorImpl) VisitUntil_stmt(ctx *parser.Until_stmtContext) interface{} {
	return ctx.Date_spec().Accept(v).(string)
}

func (v *QlVisitorImpl) VisitRelative_date(ctx *parser.Relative_dateContext) interface{} {
	result := fmt.Sprintf("%s%s", ctx.SIGNED_NUMERIC_LITERAL().GetText(),  ctx.TIMESPAN_UNIT().GetText())
	return result
}

func (v *QlVisitorImpl) VisitOperation_column(ctx *parser.Operation_columnContext) interface{} {
	field := Field{}

	field.Expr = VarRef{Identity: ctx.Arithmetic_operation().GetText()}

	if ctx.Column_alias() != nil {
		field.Alias = ctx.Column_alias().Accept(v).(string)
	}
	return field
}

func (v *QlVisitorImpl) VisitArithmetic_operation(ctx *parser.Arithmetic_operationContext) interface{} {
	return ctx.GetText()
}

func (v *QlVisitorImpl) VisitDate_spec(ctx *parser.Date_specContext) interface{} {
	if ctx.Relative_date() != nil {
		return ctx.Relative_date().Accept(v).(string)
	}

	return ctx.DATE_LITERAL().GetText()
}

func (v *QlVisitorImpl) VisitColumn_stmt(ctx *parser.Column_stmtContext) interface{} {
	if ctx.Result_column() != nil {
		return ctx.Result_column().Accept(v).(Field)
	} else {
		return ctx.Operation_column().Accept(v).(Field)
	}
}

