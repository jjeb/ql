// Code generated from parser/QlParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // QlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by QlParser.
type QlParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by QlParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by QlParser#show_stmt.
	VisitShow_stmt(ctx *Show_stmtContext) interface{}

	// Visit a parse tree produced by QlParser#from_stmt.
	VisitFrom_stmt(ctx *From_stmtContext) interface{}

	// Visit a parse tree produced by QlParser#since_stmt.
	VisitSince_stmt(ctx *Since_stmtContext) interface{}

	// Visit a parse tree produced by QlParser#until_stmt.
	VisitUntil_stmt(ctx *Until_stmtContext) interface{}

	// Visit a parse tree produced by QlParser#operation_column.
	VisitOperation_column(ctx *Operation_columnContext) interface{}

	// Visit a parse tree produced by QlParser#result_column.
	VisitResult_column(ctx *Result_columnContext) interface{}

	// Visit a parse tree produced by QlParser#column_alias.
	VisitColumn_alias(ctx *Column_aliasContext) interface{}

	// Visit a parse tree produced by QlParser#arithmetic_operation.
	VisitArithmetic_operation(ctx *Arithmetic_operationContext) interface{}

	// Visit a parse tree produced by QlParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by QlParser#date_spec.
	VisitDate_spec(ctx *Date_specContext) interface{}

	// Visit a parse tree produced by QlParser#relative_date.
	VisitRelative_date(ctx *Relative_dateContext) interface{}

	// Visit a parse tree produced by QlParser#column_stmt.
	VisitColumn_stmt(ctx *Column_stmtContext) interface{}
}
