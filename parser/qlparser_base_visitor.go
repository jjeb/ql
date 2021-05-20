// Code generated from parser/QlParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // QlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseQlParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseQlParserVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQlParserVisitor) VisitShow_stmt(ctx *Show_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQlParserVisitor) VisitFrom_stmt(ctx *From_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQlParserVisitor) VisitSince_stmt(ctx *Since_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQlParserVisitor) VisitUntil_stmt(ctx *Until_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQlParserVisitor) VisitOperation_column(ctx *Operation_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQlParserVisitor) VisitResult_column(ctx *Result_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQlParserVisitor) VisitColumn_alias(ctx *Column_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQlParserVisitor) VisitArithmetic_operation(ctx *Arithmetic_operationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQlParserVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQlParserVisitor) VisitDate_spec(ctx *Date_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQlParserVisitor) VisitRelative_date(ctx *Relative_dateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQlParserVisitor) VisitColumn_stmt(ctx *Column_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}
