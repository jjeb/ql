// Code generated from parser/QlParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // QlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseQlParserListener is a complete listener for a parse tree produced by QlParser.
type BaseQlParserListener struct{}

var _ QlParserListener = &BaseQlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseQlParserListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseQlParserListener) ExitQuery(ctx *QueryContext) {}

// EnterShow_stmt is called when production show_stmt is entered.
func (s *BaseQlParserListener) EnterShow_stmt(ctx *Show_stmtContext) {}

// ExitShow_stmt is called when production show_stmt is exited.
func (s *BaseQlParserListener) ExitShow_stmt(ctx *Show_stmtContext) {}

// EnterFrom_stmt is called when production from_stmt is entered.
func (s *BaseQlParserListener) EnterFrom_stmt(ctx *From_stmtContext) {}

// ExitFrom_stmt is called when production from_stmt is exited.
func (s *BaseQlParserListener) ExitFrom_stmt(ctx *From_stmtContext) {}

// EnterSince_stmt is called when production since_stmt is entered.
func (s *BaseQlParserListener) EnterSince_stmt(ctx *Since_stmtContext) {}

// ExitSince_stmt is called when production since_stmt is exited.
func (s *BaseQlParserListener) ExitSince_stmt(ctx *Since_stmtContext) {}

// EnterUntil_stmt is called when production until_stmt is entered.
func (s *BaseQlParserListener) EnterUntil_stmt(ctx *Until_stmtContext) {}

// ExitUntil_stmt is called when production until_stmt is exited.
func (s *BaseQlParserListener) ExitUntil_stmt(ctx *Until_stmtContext) {}

// EnterOperation_column is called when production operation_column is entered.
func (s *BaseQlParserListener) EnterOperation_column(ctx *Operation_columnContext) {}

// ExitOperation_column is called when production operation_column is exited.
func (s *BaseQlParserListener) ExitOperation_column(ctx *Operation_columnContext) {}

// EnterResult_column is called when production result_column is entered.
func (s *BaseQlParserListener) EnterResult_column(ctx *Result_columnContext) {}

// ExitResult_column is called when production result_column is exited.
func (s *BaseQlParserListener) ExitResult_column(ctx *Result_columnContext) {}

// EnterColumn_alias is called when production column_alias is entered.
func (s *BaseQlParserListener) EnterColumn_alias(ctx *Column_aliasContext) {}

// ExitColumn_alias is called when production column_alias is exited.
func (s *BaseQlParserListener) ExitColumn_alias(ctx *Column_aliasContext) {}

// EnterArithmetic_operation is called when production arithmetic_operation is entered.
func (s *BaseQlParserListener) EnterArithmetic_operation(ctx *Arithmetic_operationContext) {}

// ExitArithmetic_operation is called when production arithmetic_operation is exited.
func (s *BaseQlParserListener) ExitArithmetic_operation(ctx *Arithmetic_operationContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseQlParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseQlParserListener) ExitExpr(ctx *ExprContext) {}

// EnterDate_spec is called when production date_spec is entered.
func (s *BaseQlParserListener) EnterDate_spec(ctx *Date_specContext) {}

// ExitDate_spec is called when production date_spec is exited.
func (s *BaseQlParserListener) ExitDate_spec(ctx *Date_specContext) {}

// EnterRelative_date is called when production relative_date is entered.
func (s *BaseQlParserListener) EnterRelative_date(ctx *Relative_dateContext) {}

// ExitRelative_date is called when production relative_date is exited.
func (s *BaseQlParserListener) ExitRelative_date(ctx *Relative_dateContext) {}

// EnterColumn_stmt is called when production column_stmt is entered.
func (s *BaseQlParserListener) EnterColumn_stmt(ctx *Column_stmtContext) {}

// ExitColumn_stmt is called when production column_stmt is exited.
func (s *BaseQlParserListener) ExitColumn_stmt(ctx *Column_stmtContext) {}
