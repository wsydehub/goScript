grammar GoScript;

compilationUnit
    :   (functionDeclaration | (variableDeclaration ';'))* EOF
    ;

//Function
functionDeclaration
    :   'func' Identifier formalParameters (type_)* block
    ;

formalParameters
    :   '(' formalParameterDecl* ')'
    ;

formalParameterDecl
    :   type_ Identifier
    ;

//Statement | Block

block
    :   '{' blockStatement* '}'
    ;

blockStatement
    :   variableDeclaration ';'
    |   statement
    ;

//Variable

variableDeclaration
    : type_ variableDeclarators
    ;

variableDeclarators
    :   variableDeclarator (',' variableDeclarator)*
    ;

variableDeclarator
    :   Identifier ('=' variableInitializer)?
    ;

variableInitializer
    :   arrayInitializer
    |   mapInitializer
    |   expression
    ;

arrayInitializer
    :   '{' ( variableInitializer (',' variableInitializer)* )? '}'
    ;

mapInitializer
    :   '{' (expression ':' variableInitializer)* '}'
    ;

type_
    :   primitiveType Brackets*
    |   mapType Brackets*
    |   connectorType Brackets*
    |   dynamicType Brackets*
    ;

mapType
    :   'map' '<' primitiveType ',' type_ '>'
    ;

connectorType
    :   'connector' '<' Identifier '>'
    ;

dynamicType
    :   'dynamic'
    |   'error'
    ;

primitiveType
    :   'int'
    |   'uint'
    |   'float'
    |   'bool'
    |   'char'
    |   'string'
    ;

// Statement
statement
    :   block
    |   ifStatement
    |   forStatement
    |   returnStatement
    |   breakStatement
    |   continueStatement
    |   expressionStatement
    ;

ifStatement
    :   'if' expression statement ('else' statement)?
    ;

forStatement
    :   'for' '(' forControl ')' statement
    ;

forControl
    :   forInit? ';' expression? ';' forUpdate?
    ;

forInit
    :   variableDeclaration
    |   expressionList
    ;

forUpdate
    :   expressionList
    ;

returnStatement
    :   'return' expression? ';'
    ;

breakStatement
    :   'break' ';'
    ;

continueStatement
    :   'continue' ';'
    ;

expressionStatement
    :   expression ';'
    ;

//Expression

expression
    :   primary #PrimaryExpr
    |   expression '.' expression   #SelectorExpr
    |   expression '[' expression ']'   #IndexExpr
    |   expression '(' expressionList ')'  #CallExpr
    |   expression ('++' | '--')    #SelfAddExpr
    |   ('+'|'-'|'!') expression    #UnaryExpr
    |   'new' creator   #CreateExpr
    |   expression ('*'|'/'|'%') expression #MulExpr
    |   expression ('+'|'-') expression #AddExpr
    |   expression ('<=' | '>=' | '>' | '<' | '==' | '!=') expression   #ConditionalExpr
    |   expression '&&' expression  #   AndExpr
    |   expression '||' expression  #   OrExpr
    |   expression '?' expression ':' expression    #   TernaryExpr
    |<assoc=right>  expression (',' expression)* '=' expression (',' expression)*   # AssignExpr
    |<assoc=right>  expression (',' expression)* ':=' expression (',' expression)*  #   CreateAndAssignExpr
    ;

primary
    :   '(' expression ')'
    |   literal
    |   Identifier
    ;

literal
    :   integerLiteral
    |   FloatingPointLiteral
    |   CharacterLiteral
    |   StringLiteral
    |   BooleanLiteral
    |   NULL
    ;

integerLiteral
    :   HexLiteral
    |   OctalLiteral
    |   DecimalLiteral
    ;

BooleanLiteral
    :   'true'
    |   'false'
    ;

expressionList
    :   expression (',' expression)*
    ;

creator
    :   mapCreator
    |   arrayCreator
    |   connectorCreator
    |   primitiveCreator
    |   dynamicCreator
    ;

mapCreator
    :   mapType mapInitializer
    ;

arrayCreator
    :   creatorName Brackets+ arrayInitializer
    ;

creatorName
    :   primitiveType
    |   mapType
    |   connectorType
    |   dynamicType
    ;

connectorCreator
    :   connectorType '(' ')'
    ;

primitiveCreator
    :   primitiveType '(' expression? ')'
    ;

dynamicCreator
    :   dynamicType '(' expression? ')'
    ;

//Lexer
Identifier : [a-zA-Z_]([a-zA-Z0-9_])*;

Brackets    :   '['']';

NULL    :   'null';

HexLiteral : '0' ('x'|'X') HexDigit+ ;

DecimalLiteral : ('0' | '1'..'9' '0'..'9'*) ;

OctalLiteral : '0' ('0'..'7')+ ;

FloatingPointLiteral
    :   ('0'..'9')+ '.' ('0'..'9')* Exponent?
    |   ('0'..'9')+ Exponent
    |   ('0'..'9')+
    ;

Plus :   '+';

Minus :   '-';

Multiplication : '*';

Division    :   '/';

Percent :   '%';

Gt  : '>';

Get :   '>=';

Lt  : '<';

Let : '<=';

Assign  : '=';

Eq  : '==';

NEq :   '!=';

Not :   '!';

fragment
Exponent : ('e'|'E') ('+'|'-')? ('0'..'9')+ ;

CharacterLiteral
    :   '\'' ( EscapeSequence | ~('\''|'\\') ) '\''
    ;

StringLiteral
    :  '"' ( EscapeSequence | ~('\\'|'"') )* '"'
    ;

fragment
EscapeSequence
    :   '\\' ('b'|'t'|'n'|'f'|'r'|'"'|'\''|'\\')
    |   UnicodeEscape
    |   OctalEscape
    ;

fragment
OctalEscape
    :   '\\' ('0'..'3') ('0'..'7') ('0'..'7')
    |   '\\' ('0'..'7') ('0'..'7')
    |   '\\' ('0'..'7')
    ;

fragment
UnicodeEscape
    :   '\\' 'u' HexDigit HexDigit HexDigit HexDigit
    ;

fragment
HexDigit : ('0'..'9'|'a'..'f'|'A'..'F') ;

COMMENT
    :   '/*' .*? '*/'    -> channel(HIDDEN) // match anything between /* and */
    ;
WS  :   [ \r\t\u000C\n]+ -> channel(HIDDEN)
    ;

LINE_COMMENT
    : '//' ~[\r\n]* '\r'? '\n' -> channel(HIDDEN)
    ;