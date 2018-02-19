grammar Saucy;

/*
 * parser rules
 */


file_input
: ( NEWLINE | stmt )* EOF
;

stmt
:   simple_stmt
;

simple_stmt
:   small_stmt ( ';' small_stmt )* ';'? NEWLINE
;

small_stmt
:   import_stmt
;

import_stmt
:   IMPORT import_name
;

import_name
:   dotted_name
;

dotted_name
:   NAME ( '.' NAME )*
;


/*
 * lexer rules
 */

IMPORT
:   'import';

NAME
: [a-zA-Z0-9_]+
;

NEWLINE
:   '\r'? '\n'
;

SKIP_
:   ( SPACES | COMMENT | LINE_JOINING | SHEBANG) -> skip
;


/*
 * fragments
 */

fragment SPACES
:   [ \t]+
;

fragment COMMENT
:   '//' ~[\r\n\f]*     // inline comments
|   '/*' .*? '*/'       // block comments
;

fragment LINE_JOINING
: '\\' SPACES? ( '\r'? '\n' | '\r' | '\f' )
;

fragment SHEBANG
:   '#!' ~[\r\n\f]*     // '#!' followed by anything to end of line
;

