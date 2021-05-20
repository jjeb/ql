// Code generated from parser/QlParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // QlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// QlParserListener is a complete listener for a parse tree produced by QlParser.
type QlParserListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterShow_stmt is called when entering the show_stmt production.
	EnterShow_stmt(c *Show_stmtContext)

	// EnterFrom_stmt is called when entering the from_stmt production.
	EnterFrom_stmt(c *From_stmtContext)

	// EnterSince_stmt is called when entering the since_stmt production.
	EnterSince_stmt(c *Since_stmtContext)

	// EnterUntil_stmt is called when entering the until_stmt production.
	EnterUntil_stmt(c *Until_stmtContext)

	// EnterOperation_column is called when entering the operation_column production.
	EnterOperation_column(c *Operation_columnContext)

	// EnterResult_column is called when entering the result_column production.
	EnterResult_column(c *Result_columnContext)

	// EnterColumn_alias is called when entering the column_alias production.
	EnterColumn_alias(c *Column_aliasContext)

	// EnterArithmetic_operation is called when entering the arithmetic_operation production.
	EnterArithmetic_operation(c *Arithmetic_operationContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterDate_spec is called when entering the date_spec production.
	EnterDate_spec(c *Date_specContext)

	// EnterRelative_date is called when entering the relative_date production.
	EnterRelative_date(c *Relative_dateContext)

	// EnterColumn_stmt is called when entering the column_stmt production.
	EnterColumn_stmt(c *Column_stmtContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitShow_stmt is called when exiting the show_stmt production.
	ExitShow_stmt(c *Show_stmtContext)

	// ExitFrom_stmt is called when exiting the from_stmt production.
	ExitFrom_stmt(c *From_stmtContext)

	// ExitSince_stmt is called when exiting the since_stmt production.
	ExitSince_stmt(c *Since_stmtContext)

	// ExitUntil_stmt is called when exiting the until_stmt production.
	ExitUntil_stmt(c *Until_stmtContext)

	// ExitOperation_column is called when exiting the operation_column production.
	ExitOperation_column(c *Operation_columnContext)

	// ExitResult_column is called when exiting the result_column production.
	ExitResult_column(c *Result_columnContext)

	// ExitColumn_alias is called when exiting the column_alias production.
	ExitColumn_alias(c *Column_aliasContext)

	// ExitArithmetic_operation is called when exiting the arithmetic_operation production.
	ExitArithmetic_operation(c *Arithmetic_operationContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitDate_spec is called when exiting the date_spec production.
	ExitDate_spec(c *Date_specContext)

	// ExitRelative_date is called when exiting the relative_date production.
	ExitRelative_date(c *Relative_dateContext)

	// ExitColumn_stmt is called when exiting the column_stmt production.
	ExitColumn_stmt(c *Column_stmtContext)
}
