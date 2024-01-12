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
functionBody: (conditionalStatement | returnStatement)* ;

conditionalStatement: IF_DEF condition OPEN_BRACKET returnStatement CLOSE_BRACKET ;
condition: ruleConditionExpression ;
returnStatement: returnDef ;

ruleConditionExpression: equalityExpression | inExpression ;
equalityExpression: ruleExpressionPrimary OP_EQEQ ruleExpressionPrimary ;
inExpression: ruleExpressionPrimary OP_IN ruleExpressionPrimary ;
ruleExpressionPrimary: ID DOT ID ;

mutationDef: MUTATION_DEF OPEN_BRACKET mutationBody CLOSE_BRACKET ;
mutationBody: (mutationStatement)* ;
mutationStatement: ID OPEN_BRACKET mutationDetails CLOSE_BRACKET ;

mutationDetails: OPERATION_INPUT_DEF COLON ID OPERATION_RULES_DEF COLON OPEN_PAREN mutationRulesApplication CLOSE_PAREN OPERATION_FUNCTION_DEF COLON OPEN_BRACKET SANDBOX CLOSE_BRACKET ;

mutationRulesApplication: orExpression ;
orExpression: andExpression (OP_OR andExpression)* ;
andExpression: mutationRulesApplicationExpressionPrimary (OP_AND mutationRulesApplicationExpressionPrimary)* ;
mutationRulesApplicationExpressionPrimary: ID | OPEN_PAREN expression CLOSE_PAREN ;

inputDef: INPUT_DEF ID OPEN_BRACKET inputBody CLOSE_BRACKET ;
inputBody: (inputFieldDef)* ;
inputFieldDef: ID COLON typeDef (KEY_DEF)? ;

eventDef: EVENTS_DEF OPEN_BRACKET eventChannelHandler+ CLOSE_BRACKET ;
eventChannelHandler: HASH DOT ID COLON OPEN_PAREN ID COMMA ID CLOSE_PAREN OPEN_BRACKET SANDBOX CLOSE_BRACKET ;

enumDef: ENUM_DEF ID OPEN_BRACKET ID (ID)* CLOSE_BRACKET ;

// Custom annotations for relationships
references
    : REFERENCES_DEF OPEN_PAREN DOT ID CLOSE_PAREN
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
   : left=expression op=(OP_MUL|OP_DIV) right=expression  # MulDivExpr
   | left=expression op=(OP_SUM|OP_SUB) right=expression  # AddSubExpr
   | left=expression op=(OP_GT|OP_LT|OP_EQEQ|OP_NEQ|OP_GTEQ|OP_LTEQ) right=expression # RelationalExpr
   | left=expression op=(OP_AND|OP_OR) right=expression # LogicalExpr
   | OPEN_PAREN inner=expression CLOSE_PAREN                # ParenExpr
   | functionId=ID OPEN_PAREN args=exprList? CLOSE_PAREN # FunctionCall
   | id=ID                             # IdExpr
   | intVal=INT                        # IntLiteral
   | strVal=STRING                     # StringLiteral
   | boolVal=BOOLEAN                   # BooleanLiteral
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
