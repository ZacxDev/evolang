grammar EvoLang;

// Parser Rules
prog:   statement+ ;
statement: modelDef | enumDef | accessControlDef | endpointDef | serverDef | mainFunction;

modelDef: 'model' ID '{' modelBody '}' ;
modelBody: (fieldDef | eventDef)+ ;

fieldDef: ID ':' typeDef ;
eventDef: 'events' '{' eventBody '}' ;
eventBody: (ID ':' 'Event' '<' eventType '>' '{' sandbox '}' )+ ;

enumDef: 'enum' ID '{' ID (',' ID)* '}' ;

accessControlDef: 'accessControl' '{' accessRule+ '}' ;
accessRule: 'allow' ID '.' action '(' paramList ')' ifCondition ;

action: ID ;

paramList: ID (',' ID)* ;

endpointDef: 'endpoint' STRING '{' endpointBody '}' ;
endpointBody: (bodyDef | authenticateDef | returnDef)* ;

serverDef
  : 'server' ID '{' serverBody '}'
  ;

serverBody
  : (endpointDef)*
  ;

bodyDef: 'body:' ID ;
authenticateDef: 'authenticate:' ID ;
returnDef: 'return' expression ;

sandbox: '<sandbox lang="go">' goCode '</sandbox>' ;
goCode: GO_CODE ;

// Lexer Rules
ID: [a-zA-Z_][a-zA-Z0-9_]* ;
HTTP_METHOD: 'GET' | 'POST' ;
STRING: '"' ( ~["\\] | '\\' . )* '"' ;
GO_CODE: '<go>' .*? '</go>' ;
WS: [ \t\r\n]+ -> skip ;

typeDef: 'Int' | 'String' | 'Boolean' | arrayDef | refDef ;
arrayDef: '[' typeDef ']' ;
refDef: '@references' '(' ID '.' ID ')' ;
eventType: ID (',' ID)* ;
ifCondition: 'if' expression ;
expression
   : left=expression op=('*'|'/') right=expression  # MulDivExpr
   | left=expression op=('+'|'-') right=expression  # AddSubExpr
   | left=expression op=('>'|'<'|'=='|'!='|'>='|'<=') right=expression # RelationalExpr
   | left=expression op=('&&'|'||') right=expression # LogicalExpr
   | '(' inner=expression ')'                # ParenExpr
   | functionId=ID '(' args=exprList? ')' # FunctionCall
   | id=ID                             # IdExpr
   | intVal=INT                        # IntLiteral
   | strVal=STRING                     # StringLiteral
   | boolVal=BOOLEAN                   # BooleanLiteral
   ;

exprList: expression (',' expression)* ;

// Additional Lexer Rules
INT: [0-9]+ ;
BOOLEAN: 'true' | 'false' ;
OP: '*' | '/' | '+' | '-' | '>' | '<' | '==' | '!=' | '>=' | '<=' | '&&' | '||' ;

mainFunction
  : 'main' '{' mainStatements '}'
  ;

mainStatements
  : (mainStatement)+
  ;

mainStatement
  : printStatement
  | serverStatement
  ;

printStatement
  : 'print' '(' STRING ')' ';'
  ;

serverStatement
  : ID '.Listen' '(' INT ')' ';'
  ;
