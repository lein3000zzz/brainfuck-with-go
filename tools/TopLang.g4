grammar TopLang;

program: statement* EOF
;

statement: varDecl
    | assignment
    | printStmt
    | printNumberStmt
    | printStringStmt
    | readStmt
    | ifStmt
    | whileStmt
    | block
;

varDecl: VAR ID ASSIGN expr SEMI
;

assignment: ID ASSIGN expr SEMI
;

printStmt: PRINT LPAREN expr RPAREN SEMI
;

printNumberStmt: PRINTNUMBER LPAREN expr RPAREN SEMI
;

printStringStmt: PRINTSTRING LPAREN STRING RPAREN SEMI
;

readStmt: READ LPAREN ID RPAREN SEMI
;

ifStmt: IF LPAREN expr RPAREN block (ELSE block)?
;

whileStmt: WHILE LPAREN expr RPAREN block
;

block: LBRACE statement* RBRACE
;

// expressions (the higher = the more imporant it is)

expr: left=expr op=(EQ | NEQ | LT | GT | LTE | GTE) right=expr      # CompExpr
    | left=expr op=(PLUS | MINUS) right=expr                        # AddExpr
    | left=expr op=(MUL | DIV | MOD) right=expr                     # MulExpr
    | MINUS expr                                                    # UnaryMinusExpr
    | LPAREN expr RPAREN                                            # ParenExpr
    | INT                                                           # IntLiteral
    | ID                                                            # IdExpr
;

// Keywords (must appear before ID rule not to mix them up or sth)
VAR         : 'var';
PRINT       : 'print';
PRINTNUMBER : 'printNumber';
PRINTSTRING : 'printString';
READ        : 'read';
IF      : 'if';
ELSE    : 'else';
WHILE   : 'while';

// Comparison operators (multi-char first not to break anything)
EQ      : '==';
NEQ     : '!=';
LTE     : '<=';
GTE     : '>=';
LT      : '<';
GT      : '>';

// Arithmetics
PLUS    : '+';
MINUS   : '-';
MUL     : '*';
DIV     : '/';
MOD     : '%';

ASSIGN  : '=';

LPAREN  : '(';
RPAREN  : ')';
LBRACE  : '{';
RBRACE  : '}';
SEMI    : ';';

// Literals and identifiers
INT     : [0-9]+;
STRING  : '"' ( '\\' . | ~["\\\r\n] )* '"';
ID      : [a-zA-Z_][a-zA-Z_0-9]*;

// Skip whitespace and comments
WS      : [ \t\r\n]+ -> skip;
COMMENT : '//' ~[\r\n]* -> skip;
