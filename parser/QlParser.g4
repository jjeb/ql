parser grammar QlParser;

options {
	tokenVocab = QlLexer;
}

query: show_stmt from_stmt since_stmt? until_stmt? EOF;
show_stmt: SHOW column_stmt ( ',' column_stmt)*;
from_stmt: FROM IDENTIFIER;
since_stmt: SINCE date_spec;
until_stmt: UNTIL date_spec;
operation_column: arithmetic_operation ( AS column_alias)?;
result_column: IDENTIFIER ( AS column_alias)?;
column_alias: IDENTIFIER;
arithmetic_operation: '(' result_column ARITHMETIC_OPERATOR (result_column | SIGNED_NUMERIC_LITERAL) ')';
expr:	IDENTIFIER;
date_spec: DATE_LITERAL | relative_date;
relative_date: SIGNED_NUMERIC_LITERAL TIMESPAN_UNIT;
column_stmt: result_column | operation_column;
