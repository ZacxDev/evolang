parser grammar EvoLangParser;

options {
    tokenVocab=EvoLangLexer;
}

// Parser Rules
prog:   statement+ EOF ;
statement: modelDef | enumDef | ruleDef | mutationDef | inputDef | mainFunction;

modelDef: MODEL_DEF ID OPEN_BRACKET modelBody CLOSE_BRACKET ;
modelBody: (modelFieldDef | eventDef)+ ;

modelFieldDef
    : ID COLON typeDef (references | linkedBy)?
    ;

ruleDef: RULE_DEF OPEN_BRACKET ruleBody CLOSE_BRACKET ;
ruleBody: (ruleStatement)* ;
ruleStatement: ID COLON OPEN_PAREN ruleParameters CLOSE_PAREN OPEN_BRACKET functionBody CLOSE_BRACKET ;
ruleParameters: ID ID (COMMA ID ID)* ;
functionBody: (assignmentStatement | conditionalStatement | returnStatement)* ;
assignmentStatement: ID OP_ASSIGN expression ;

conditionalStatement: IF_DEF condition OPEN_BRACKET returnStatement CLOSE_BRACKET ;
condition: ruleConditionExpression ;
returnStatement: returnDef ;

ruleConditionExpression: equalityExpression | inExpression ;
equalityExpression: ruleExpressionPrimary OP_EQEQ ruleExpressionPrimary ;
inExpression: ruleExpressionPrimary OP_IN ruleExpressionPrimary ;
ruleExpressionPrimary: ID DOT ID ;

mutationDef: MUTATION_DEF OPEN_BRACKET mutationBody CLOSE_BRACKET ;
mutationBody: (mutationStatement)* ;
mutationStatement: ID COLON OPEN_PAREN mutationParameters CLOSE_PAREN OPEN_BRACKET mutationInnerBlock CLOSE_BRACKET ;
mutationParameters: ID ID (COMMA ID ID)* ;
mutationInnerBlock: (middlewareBlock | functionBlock)* ;

middlewareBlock: OPERATION_MIDDLEWARE_DEF COLON OPEN_BRACKET functionBody CLOSE_BRACKET ;
functionBlock: OPERATION_FUNCTION_DEF COLON OPEN_BRACKET SANDBOX CLOSE_BRACKET ;

inputDef: INPUT_DEF ID OPEN_BRACKET inputBody CLOSE_BRACKET ;
inputBody: (inputFieldDef)* ;
inputFieldDef: ID COLON typeDef references? ;

eventDef: EVENTS_DEF OPEN_BRACKET eventChannelHandler+ CLOSE_BRACKET ;
eventChannelHandler: HASH DOT ID COLON OPEN_PAREN ID COMMA ID CLOSE_PAREN OPEN_BRACKET SANDBOX CLOSE_BRACKET ;

enumDef: ENUM_DEF ID OPEN_BRACKET ID (ID)* CLOSE_BRACKET ;

// Custom annotations for relationships
references
    : REFERENCES_DEF OPEN_PAREN ID? DOT ID CLOSE_PAREN
    ;

linkedBy
    : LINKED_BY_DEF OPEN_PAREN DOT ID CLOSE_PAREN
    ;

action: ID ;

paramList: ID (COMMA ID)* ;

bodyDef: BODY_DEF ID ;
authenticateDef: AUTHENTICATE_DEF ID ;
returnDef: RETURN_DEF expression ;

typeDef: TYPE_INT_DEF | TYPE_STRING_DEF | TYPE_BOOLEAN_DEF | arrayDef | ID;
arrayDef: OPEN_S_BRACKET typeDef CLOSE_S_BRACKET ;
eventType: ID (COMMA ID)* ;
expression
   : methodCallExpression                                 # MethodCallExpr
   | left=expression op=(OP_MUL|OP_DIV) right=expression  # MulDivExpr
   | left=expression op=(OP_SUM|OP_SUB) right=expression  # AddSubExpr
   | left=expression op=(OP_GT|OP_LT|OP_EQEQ|OP_NEQ|OP_GTEQ|OP_LTEQ) right=expression # RelationalExpr
   | left=expression op=(OP_AND|OP_OR) right=expression # LogicalExpr
   | OPEN_PAREN inner=expression CLOSE_PAREN                # ParenExpr
   | propertyAccess                                         # PropertyAccessExpr
   | functionCall                                           # FunctionCallExpr
   | id=ID                                                  # IdExpr
   | intVal=INT                                             # IntLiteral
   | strVal=STRING                                          # StringLiteral
   | boolVal=BOOLEAN                                        # BooleanLiteral
   ;

methodCallExpression
   : primaryExpression (DOT ID OPEN_PAREN args=exprList? CLOSE_PAREN)+
   ;

primaryExpression
   : ID  # SimpleIdExpr
   | ID OPEN_PAREN args=exprList? CLOSE_PAREN  # PrimaryFunctionCallExpr
   | ID DOT ID  # PrimaryPropertyAccessExpr
   // ... add other primary expression types as needed
   ;

propertyAccess
   : ID (DOT ID)+
   ;

functionCall
   : ID OPEN_PAREN args=exprList? CLOSE_PAREN
   ;

exprList: expression (COMMA expression)* ;

mainFunction
  : MAIN_DEF OPEN_BRACKET mainStatements CLOSE_BRACKET
  ;

mainStatements
  : (mainStatement)+
  ;

mainStatement
  : printStatement
  | serverStatement
  ;

printStatement
  : PRINT_DEF OPEN_PAREN STRING CLOSE_PAREN
  ;

serverStatement
  : ID DOT LISTEN_DEF OPEN_PAREN INT CLOSE_PAREN
  ;

sandboxBlock
    : OPEN_SANDBOX_MODE sandboxAttributes SANDBOX_CONTENT CLOSE_SANDBOX_MODE
    ;

sandboxAttributes
    : (sandboxAttribute)*
    ;

sandboxAttribute
    : sandboxLangAttribute
    | sandboxDefineAttribute
    ;

sandboxLangAttribute
    : LANG_DEF OP_ASSIGN STRING
    ;

sandboxDefineAttribute
    : DEFINE_DEF OP_ASSIGN OPEN_BRACKET (sandboxDefineContent) CLOSE_BRACKET
    ;

sandboxDefineContent
    : ID COLON ID (COMMA ID COLON ID)*
    ;

