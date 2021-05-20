lexer grammar QlLexer;

COMMA: ',';
DOT: '.';
LR_BRACKET: '(';
RR_BRACKET: ')';

ARITHMETIC_OPERATOR:
  PLUS | MINUS | DIV | MOD | MUL;

PLUS: '+';
MINUS: '-';
DIV: '/';
MOD: '%';
MUL: '*';


SHOW: S H O W;
AS: A S;
FROM: F R O M;
SINCE: S I N C E;
UNTIL: U N T I L;

TIMESPAN_UNIT:
  SECONDS | MINUTES | HOURS | DAYS | WEEKS | MONTHS | QUARTERS | YEARS;

SECONDS: 's';
MINUTES: 'min';
HOURS: 'h';
DAYS: 'd';
WEEKS: 'w';
MONTHS: 'm';
QUARTERS: 'q';
YEARS: 'y';

IDENTIFIER: [a-zA-Z_] [a-zA-Z_0-9]*;

SIGNED_NUMERIC_LITERAL: ( PLUS | MINUS)? NUMERIC_LITERAL;

NUMERIC_LITERAL:
  ((DIGIT+ ('.' DIGIT*)?) | ('.' DIGIT+)) (E [-+]? DIGIT+)?;

YEAR: DIGIT DIGIT DIGIT DIGIT;
MONTH: DIGIT DIGIT;
DAY: DIGIT DIGIT;

DATE_LITERAL: YEAR '-' MONTH '-' DAY;

SPACES: [ \u000B\t\r\n] -> channel(HIDDEN);

fragment DIGIT: [0-9];

fragment A: [aA];
fragment B: [bB];
fragment C: [cC];
fragment D: [dD];
fragment E: [eE];
fragment F: [fF];
fragment G: [gG];
fragment H: [hH];
fragment I: [iI];
fragment J: [jJ];
fragment K: [kK];
fragment L: [lL];
fragment M: [mM];
fragment N: [nN];
fragment O: [oO];
fragment P: [pP];
fragment Q: [qQ];
fragment R: [rR];
fragment S: [sS];
fragment T: [tT];
fragment U: [uU];
fragment V: [vV];
fragment W: [wW];
fragment X: [xX];
fragment Y: [yY];
fragment Z: [zZ];
