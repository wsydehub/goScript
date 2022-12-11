grammar GoScript;

compilationUnit
    :   (functionDeclaration | (variableDeclaration ';'))* EOF
    ;

//Function
functionDeclaration
    :   'func' Identifier formalParameters type_ block
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
    :   '{' ( arrayInitializerElement (',' arrayInitializerElement)* )? '}'
    ;

arrayInitializerElement
    :   mapInitializer
    |   expression
    ;

mapInitializer
    :   '{' (expression ':' variableInitializer)* '}'
    ;

type_
    :   primitiveType Brackets?
    |   mapType Brackets?
    |   connectorType Brackets?
    |   dynamicType Brackets?
    ;

mapType
    :   'map' '<' (primitiveType | connectorType) ',' type_ '>'
    ;

connectorType
    :   'connector' '<' Identifier '>'
    ;

dynamicType
    :   'dynamic'
    ;

primitiveType
    :   'int'
    |   'float'
    |   'bool'
    |   'char'
    |   'string'
    ;

// Statement
statement
    : block #   CommonBlockStatement
    |   'if' expression statement ('else' statement)?   #   IfStatement
    |   'for' '(' forControl ')' statement  #   ForStatement
    |   'return' expression? ';'    #   ReturnStatement
    |   'break' ';' #   BreakStatement
    |   'continue' ';'  #   ContinueStatement
    |   expression ';'  #   ExprStatement
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

//Expression

expression
    :   primary # PrimaryExpr
    |   expression '.' expression   #FetchExpr
    |   expression '[' expression ']'   # IndexExpr
    |   expression '(' expressionList? ')'  # CallExpr
    |   ('+'|'-') expression    # NegativeExpr
    |   '!' expression  # NotExpr
    |   'new' creator   # CreatorExpr
    |   expression ('*'|'/'|'%') expression #   MulDivExpr
    |   expression ('+'|'-') expression #   AddSubExpr
    |   expression ('<' '=' | '>' '=' | '>' | '<') expression   #   CompareExpr
    |   expression ('==' | '!=') expression #   EqNEqExpr
    |   expression '&&' expression  #   AndExpr
    |   expression '||' expression  #   OrExpr
    |   expression '?' expression ':' expression    #   TernaryExpr
    |<assoc=right>  expression '=' expression   # AssignExpr
    |<assoc=right>  Identifier ':=' expression  #   CreatorAndAssignExpr
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
    :   mapType mapInitializer # MapCreator
    |   creatorName '[' ']' arrayInitializer # ArrayCreator
    |   connectorType '(' ')' # MeegoCreator
    |   primitiveType '(' expression? ')' # PrimitiveCreator
    |   dynamicType '(' expression? ')' #DynamicCreator
    ;

creatorName
    :   primitiveType
    |   mapType
    |   connectorType
    |   dynamicType
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

Lt  : '<';

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