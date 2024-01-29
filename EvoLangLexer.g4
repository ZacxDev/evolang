lexer grammar EvoLangLexer;

MODEL_DEF: 'model' ;
EVENTS_DEF: 'events' ;
ENUM_DEF: 'enum' ;
REFERENCES_DEF: '@references' ;
LINKED_BY_DEF: '@linkedBy' ;
BODY_DEF: 'body' ;
AUTHENTICATE_DEF: 'authenticate' ;
RETURN_DEF: 'return' ;
TYPE_INT_DEF: 'Int' ;
TYPE_STRING_DEF: 'String' ;
TYPE_BOOLEAN_DEF: 'Boolean' ;
IF_DEF: 'if' ;
MAIN_DEF: 'main' ;
PRINT_DEF: 'print' ;
LISTEN_DEF: 'Listen' ;

QUERY_DEF: 'query';

RULE_DEF: 'rule' ;
MUTATION_DEF: 'mutation' ;
INPUT_DEF: 'input' ;

OPERATION_FUNCTION_DEF: 'Function' ;
OPERATION_MIDDLEWARE_DEF: 'Middleware' ;
OPERATION_INPUT_DEF: 'Input' ;

OP_MUL: '*' ;
OP_DIV: '/' ;
OP_SUM: '+' ;
OP_SUB: '-' ;
OP_GT: '>' ;
OP_LT: '<' ;
OP_EQEQ: '==' ;
OP_NEQ: '!=' ;
OP_GTEQ: '>=' ;
OP_LTEQ: '<=' ;
OP_AND: '&&' ;
OP_OR: '||' ;
OP_IN: 'in' ;
OP_ASSIGN: '=' ;

DEFINE_DEF: 'define' ;
LANG_DEF: 'lang' ;

// Basic Lexer Rules
ID: [a-zA-Z_][a-zA-Z0-9_]* ;
STRING: '"' ( ~["\\] | '\\' . )* '"' ;
INT: [0-9]+ ;
BOOLEAN: 'true' | 'false' ;
WS: [ \t\r\n]+ -> skip ;
COMMENT: '//' .*? '\n' -> skip ;
WHITESPACE: [ \t\r\n]+ -> skip ;

OPEN_BRACKET: '{' ;
CLOSE_BRACKET: '}' ;
OPEN_S_BRACKET: '[' ;
CLOSE_S_BRACKET: ']' ;
OPEN_PAREN: '(' ;
CLOSE_PAREN: ')' ;
COLON: ':' ;
HASH: '#' ;
DOT: '.' ;
COMMA: ',' ;
AT: '@' ;
PIPE: '|' ;

// Define tokens for sandbox delimiters
OPEN_SANDBOX_MODE: '[sandbox' -> pushMode(SANDBOX_MODE);
CLOSE_SANDBOX_MODE: '[/sandbox]' -> popMode;
SANDBOX: OPEN_SANDBOX_MODE SANDBOX_CONTENT CLOSE_SANDBOX_MODE ; // The complete sandbox block

mode SANDBOX_MODE;
    // Define how to handle the content within the sandbox
    SANDBOX_CONTENT: ~[\\[]+; // Adjust based on your content needs

