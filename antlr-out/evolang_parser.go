// Code generated from ./EvoLangParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // EvoLangParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type EvoLangParser struct {
	*antlr.BaseParser
}

var EvoLangParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func evolangparserParserInit() {
	staticData := &EvoLangParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'model'", "'events'", "'enum'", "'@references'", "'@linkedBy'",
		"'body'", "'authenticate'", "'return'", "'Int'", "'String'", "'Boolean'",
		"'if'", "'main'", "'print'", "'Listen'", "'query'", "'rule'", "'mutation'",
		"'input'", "'Function'", "'Middleware'", "'Input'", "'*'", "'/'", "'+'",
		"'-'", "'>'", "'<'", "'=='", "'!='", "'>='", "'<='", "'&&'", "'||'",
		"'in'", "'='", "'define'", "'lang'", "", "", "", "", "", "", "", "'{'",
		"'}'", "'['", "']'", "'('", "')'", "':'", "'#'", "'.'", "','", "'@'",
		"'|'", "'[sandbox'", "'[/sandbox]'",
	}
	staticData.SymbolicNames = []string{
		"", "MODEL_DEF", "EVENTS_DEF", "ENUM_DEF", "REFERENCES_DEF", "LINKED_BY_DEF",
		"BODY_DEF", "AUTHENTICATE_DEF", "RETURN_DEF", "TYPE_INT_DEF", "TYPE_STRING_DEF",
		"TYPE_BOOLEAN_DEF", "IF_DEF", "MAIN_DEF", "PRINT_DEF", "LISTEN_DEF",
		"QUERY_DEF", "RULE_DEF", "MUTATION_DEF", "INPUT_DEF", "OPERATION_FUNCTION_DEF",
		"OPERATION_MIDDLEWARE_DEF", "OPERATION_INPUT_DEF", "OP_MUL", "OP_DIV",
		"OP_SUM", "OP_SUB", "OP_GT", "OP_LT", "OP_EQEQ", "OP_NEQ", "OP_GTEQ",
		"OP_LTEQ", "OP_AND", "OP_OR", "OP_IN", "OP_ASSIGN", "DEFINE_DEF", "LANG_DEF",
		"ID", "STRING", "INT", "BOOLEAN", "WS", "COMMENT", "WHITESPACE", "OPEN_BRACKET",
		"CLOSE_BRACKET", "OPEN_S_BRACKET", "CLOSE_S_BRACKET", "OPEN_PAREN",
		"CLOSE_PAREN", "COLON", "HASH", "DOT", "COMMA", "AT", "PIPE", "OPEN_SANDBOX_MODE",
		"CLOSE_SANDBOX_MODE", "SANDBOX", "SANDBOX_CONTENT",
	}
	staticData.RuleNames = []string{
		"prog", "statement", "modelDef", "modelBody", "modelFieldDef", "ruleDef",
		"ruleBody", "ruleStatement", "ruleParameters", "functionBody", "assignmentStatement",
		"conditionalStatement", "condition", "returnStatement", "ruleConditionExpression",
		"equalityExpression", "inExpression", "ruleExpressionPrimary", "mutationDef",
		"mutationBody", "mutationStatement", "mutationParameters", "mutationInnerBlock",
		"middlewareBlock", "functionBlock", "inputDef", "inputBody", "inputFieldDef",
		"eventDef", "eventChannelHandler", "enumDef", "references", "linkedBy",
		"action", "paramList", "bodyDef", "authenticateDef", "returnDef", "typeDef",
		"arrayDef", "eventType", "expression", "methodCallExpression", "primaryExpression",
		"propertyAccess", "functionCall", "exprList", "mainFunction", "mainStatements",
		"mainStatement", "printStatement", "serverStatement", "sandboxBlock",
		"sandboxAttributes", "sandboxAttribute", "sandboxLangAttribute", "sandboxDefineAttribute",
		"sandboxDefineContent",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 61, 513, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		1, 0, 4, 0, 118, 8, 0, 11, 0, 12, 0, 119, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 3, 1, 130, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 4, 3, 140, 8, 3, 11, 3, 12, 3, 141, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 3, 4, 149, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 5, 6, 157, 8,
		6, 10, 6, 12, 6, 160, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 176, 8, 8, 10, 8, 12, 8, 179,
		9, 8, 1, 9, 1, 9, 1, 9, 5, 9, 184, 8, 9, 10, 9, 12, 9, 187, 9, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 14, 1, 14, 3, 14, 205, 8, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 19, 5, 19, 225, 8, 19, 10, 19, 12, 19, 228,
		9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 244, 8, 21, 10, 21, 12, 21, 247,
		9, 21, 1, 22, 1, 22, 5, 22, 251, 8, 22, 10, 22, 12, 22, 254, 9, 22, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 5, 26, 275, 8,
		26, 10, 26, 12, 26, 278, 9, 26, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 284,
		8, 27, 1, 28, 1, 28, 1, 28, 4, 28, 289, 8, 28, 11, 28, 12, 28, 290, 1,
		28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 313,
		8, 30, 10, 30, 12, 30, 316, 9, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 3,
		31, 323, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 5, 34, 340, 8, 34, 10,
		34, 12, 34, 343, 9, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 37,
		1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 359, 8, 38, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 5, 40, 368, 8, 40, 10, 40,
		12, 40, 371, 9, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 385, 8, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 5, 41, 399,
		8, 41, 10, 41, 12, 41, 402, 9, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3,
		42, 409, 8, 42, 1, 42, 4, 42, 412, 8, 42, 11, 42, 12, 42, 413, 1, 43, 1,
		43, 1, 43, 1, 43, 3, 43, 420, 8, 43, 1, 43, 1, 43, 1, 43, 1, 43, 3, 43,
		426, 8, 43, 1, 44, 1, 44, 1, 44, 4, 44, 431, 8, 44, 11, 44, 12, 44, 432,
		1, 45, 1, 45, 1, 45, 3, 45, 438, 8, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1,
		46, 5, 46, 445, 8, 46, 10, 46, 12, 46, 448, 9, 46, 1, 47, 1, 47, 1, 47,
		1, 47, 1, 47, 1, 48, 4, 48, 456, 8, 48, 11, 48, 12, 48, 457, 1, 49, 1,
		49, 3, 49, 462, 8, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51,
		1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1,
		53, 5, 53, 482, 8, 53, 10, 53, 12, 53, 485, 9, 53, 1, 54, 1, 54, 3, 54,
		489, 8, 54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56, 1, 56, 1,
		56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 5, 57, 508,
		8, 57, 10, 57, 12, 57, 511, 9, 57, 1, 57, 0, 1, 82, 58, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80,
		82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112,
		114, 0, 4, 1, 0, 23, 24, 1, 0, 25, 26, 1, 0, 27, 32, 1, 0, 33, 34, 509,
		0, 117, 1, 0, 0, 0, 2, 129, 1, 0, 0, 0, 4, 131, 1, 0, 0, 0, 6, 139, 1,
		0, 0, 0, 8, 143, 1, 0, 0, 0, 10, 150, 1, 0, 0, 0, 12, 158, 1, 0, 0, 0,
		14, 161, 1, 0, 0, 0, 16, 170, 1, 0, 0, 0, 18, 185, 1, 0, 0, 0, 20, 188,
		1, 0, 0, 0, 22, 192, 1, 0, 0, 0, 24, 198, 1, 0, 0, 0, 26, 200, 1, 0, 0,
		0, 28, 204, 1, 0, 0, 0, 30, 206, 1, 0, 0, 0, 32, 210, 1, 0, 0, 0, 34, 214,
		1, 0, 0, 0, 36, 218, 1, 0, 0, 0, 38, 226, 1, 0, 0, 0, 40, 229, 1, 0, 0,
		0, 42, 238, 1, 0, 0, 0, 44, 252, 1, 0, 0, 0, 46, 255, 1, 0, 0, 0, 48, 261,
		1, 0, 0, 0, 50, 267, 1, 0, 0, 0, 52, 276, 1, 0, 0, 0, 54, 279, 1, 0, 0,
		0, 56, 285, 1, 0, 0, 0, 58, 294, 1, 0, 0, 0, 60, 307, 1, 0, 0, 0, 62, 319,
		1, 0, 0, 0, 64, 328, 1, 0, 0, 0, 66, 334, 1, 0, 0, 0, 68, 336, 1, 0, 0,
		0, 70, 344, 1, 0, 0, 0, 72, 347, 1, 0, 0, 0, 74, 350, 1, 0, 0, 0, 76, 358,
		1, 0, 0, 0, 78, 360, 1, 0, 0, 0, 80, 364, 1, 0, 0, 0, 82, 384, 1, 0, 0,
		0, 84, 403, 1, 0, 0, 0, 86, 425, 1, 0, 0, 0, 88, 427, 1, 0, 0, 0, 90, 434,
		1, 0, 0, 0, 92, 441, 1, 0, 0, 0, 94, 449, 1, 0, 0, 0, 96, 455, 1, 0, 0,
		0, 98, 461, 1, 0, 0, 0, 100, 463, 1, 0, 0, 0, 102, 468, 1, 0, 0, 0, 104,
		475, 1, 0, 0, 0, 106, 483, 1, 0, 0, 0, 108, 488, 1, 0, 0, 0, 110, 490,
		1, 0, 0, 0, 112, 494, 1, 0, 0, 0, 114, 500, 1, 0, 0, 0, 116, 118, 3, 2,
		1, 0, 117, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0,
		119, 120, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 122, 5, 0, 0, 1, 122,
		1, 1, 0, 0, 0, 123, 130, 3, 4, 2, 0, 124, 130, 3, 60, 30, 0, 125, 130,
		3, 10, 5, 0, 126, 130, 3, 36, 18, 0, 127, 130, 3, 50, 25, 0, 128, 130,
		3, 94, 47, 0, 129, 123, 1, 0, 0, 0, 129, 124, 1, 0, 0, 0, 129, 125, 1,
		0, 0, 0, 129, 126, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 129, 128, 1, 0, 0,
		0, 130, 3, 1, 0, 0, 0, 131, 132, 5, 1, 0, 0, 132, 133, 5, 39, 0, 0, 133,
		134, 5, 46, 0, 0, 134, 135, 3, 6, 3, 0, 135, 136, 5, 47, 0, 0, 136, 5,
		1, 0, 0, 0, 137, 140, 3, 8, 4, 0, 138, 140, 3, 56, 28, 0, 139, 137, 1,
		0, 0, 0, 139, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 139, 1, 0, 0,
		0, 141, 142, 1, 0, 0, 0, 142, 7, 1, 0, 0, 0, 143, 144, 5, 39, 0, 0, 144,
		145, 5, 52, 0, 0, 145, 148, 3, 76, 38, 0, 146, 149, 3, 62, 31, 0, 147,
		149, 3, 64, 32, 0, 148, 146, 1, 0, 0, 0, 148, 147, 1, 0, 0, 0, 148, 149,
		1, 0, 0, 0, 149, 9, 1, 0, 0, 0, 150, 151, 5, 17, 0, 0, 151, 152, 5, 46,
		0, 0, 152, 153, 3, 12, 6, 0, 153, 154, 5, 47, 0, 0, 154, 11, 1, 0, 0, 0,
		155, 157, 3, 14, 7, 0, 156, 155, 1, 0, 0, 0, 157, 160, 1, 0, 0, 0, 158,
		156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 13, 1, 0, 0, 0, 160, 158, 1,
		0, 0, 0, 161, 162, 5, 39, 0, 0, 162, 163, 5, 52, 0, 0, 163, 164, 5, 50,
		0, 0, 164, 165, 3, 16, 8, 0, 165, 166, 5, 51, 0, 0, 166, 167, 5, 46, 0,
		0, 167, 168, 3, 18, 9, 0, 168, 169, 5, 47, 0, 0, 169, 15, 1, 0, 0, 0, 170,
		171, 5, 39, 0, 0, 171, 177, 5, 39, 0, 0, 172, 173, 5, 55, 0, 0, 173, 174,
		5, 39, 0, 0, 174, 176, 5, 39, 0, 0, 175, 172, 1, 0, 0, 0, 176, 179, 1,
		0, 0, 0, 177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 17, 1, 0, 0,
		0, 179, 177, 1, 0, 0, 0, 180, 184, 3, 20, 10, 0, 181, 184, 3, 22, 11, 0,
		182, 184, 3, 26, 13, 0, 183, 180, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183,
		182, 1, 0, 0, 0, 184, 187, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 186,
		1, 0, 0, 0, 186, 19, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188, 189, 5, 39,
		0, 0, 189, 190, 5, 36, 0, 0, 190, 191, 3, 82, 41, 0, 191, 21, 1, 0, 0,
		0, 192, 193, 5, 12, 0, 0, 193, 194, 3, 24, 12, 0, 194, 195, 5, 46, 0, 0,
		195, 196, 3, 26, 13, 0, 196, 197, 5, 47, 0, 0, 197, 23, 1, 0, 0, 0, 198,
		199, 3, 28, 14, 0, 199, 25, 1, 0, 0, 0, 200, 201, 3, 74, 37, 0, 201, 27,
		1, 0, 0, 0, 202, 205, 3, 30, 15, 0, 203, 205, 3, 32, 16, 0, 204, 202, 1,
		0, 0, 0, 204, 203, 1, 0, 0, 0, 205, 29, 1, 0, 0, 0, 206, 207, 3, 34, 17,
		0, 207, 208, 5, 29, 0, 0, 208, 209, 3, 34, 17, 0, 209, 31, 1, 0, 0, 0,
		210, 211, 3, 34, 17, 0, 211, 212, 5, 35, 0, 0, 212, 213, 3, 34, 17, 0,
		213, 33, 1, 0, 0, 0, 214, 215, 5, 39, 0, 0, 215, 216, 5, 54, 0, 0, 216,
		217, 5, 39, 0, 0, 217, 35, 1, 0, 0, 0, 218, 219, 5, 18, 0, 0, 219, 220,
		5, 46, 0, 0, 220, 221, 3, 38, 19, 0, 221, 222, 5, 47, 0, 0, 222, 37, 1,
		0, 0, 0, 223, 225, 3, 40, 20, 0, 224, 223, 1, 0, 0, 0, 225, 228, 1, 0,
		0, 0, 226, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 39, 1, 0, 0, 0,
		228, 226, 1, 0, 0, 0, 229, 230, 5, 39, 0, 0, 230, 231, 5, 52, 0, 0, 231,
		232, 5, 50, 0, 0, 232, 233, 3, 42, 21, 0, 233, 234, 5, 51, 0, 0, 234, 235,
		5, 46, 0, 0, 235, 236, 3, 44, 22, 0, 236, 237, 5, 47, 0, 0, 237, 41, 1,
		0, 0, 0, 238, 239, 5, 39, 0, 0, 239, 245, 5, 39, 0, 0, 240, 241, 5, 55,
		0, 0, 241, 242, 5, 39, 0, 0, 242, 244, 5, 39, 0, 0, 243, 240, 1, 0, 0,
		0, 244, 247, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246,
		43, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 248, 251, 3, 46, 23, 0, 249, 251,
		3, 48, 24, 0, 250, 248, 1, 0, 0, 0, 250, 249, 1, 0, 0, 0, 251, 254, 1,
		0, 0, 0, 252, 250, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 45, 1, 0, 0,
		0, 254, 252, 1, 0, 0, 0, 255, 256, 5, 21, 0, 0, 256, 257, 5, 52, 0, 0,
		257, 258, 5, 46, 0, 0, 258, 259, 3, 18, 9, 0, 259, 260, 5, 47, 0, 0, 260,
		47, 1, 0, 0, 0, 261, 262, 5, 20, 0, 0, 262, 263, 5, 52, 0, 0, 263, 264,
		5, 46, 0, 0, 264, 265, 5, 60, 0, 0, 265, 266, 5, 47, 0, 0, 266, 49, 1,
		0, 0, 0, 267, 268, 5, 19, 0, 0, 268, 269, 5, 39, 0, 0, 269, 270, 5, 46,
		0, 0, 270, 271, 3, 52, 26, 0, 271, 272, 5, 47, 0, 0, 272, 51, 1, 0, 0,
		0, 273, 275, 3, 54, 27, 0, 274, 273, 1, 0, 0, 0, 275, 278, 1, 0, 0, 0,
		276, 274, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 53, 1, 0, 0, 0, 278, 276,
		1, 0, 0, 0, 279, 280, 5, 39, 0, 0, 280, 281, 5, 52, 0, 0, 281, 283, 3,
		76, 38, 0, 282, 284, 3, 62, 31, 0, 283, 282, 1, 0, 0, 0, 283, 284, 1, 0,
		0, 0, 284, 55, 1, 0, 0, 0, 285, 286, 5, 2, 0, 0, 286, 288, 5, 46, 0, 0,
		287, 289, 3, 58, 29, 0, 288, 287, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290,
		288, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 293,
		5, 47, 0, 0, 293, 57, 1, 0, 0, 0, 294, 295, 5, 53, 0, 0, 295, 296, 5, 54,
		0, 0, 296, 297, 5, 39, 0, 0, 297, 298, 5, 52, 0, 0, 298, 299, 5, 50, 0,
		0, 299, 300, 5, 39, 0, 0, 300, 301, 5, 55, 0, 0, 301, 302, 5, 39, 0, 0,
		302, 303, 5, 51, 0, 0, 303, 304, 5, 46, 0, 0, 304, 305, 5, 60, 0, 0, 305,
		306, 5, 47, 0, 0, 306, 59, 1, 0, 0, 0, 307, 308, 5, 3, 0, 0, 308, 309,
		5, 39, 0, 0, 309, 310, 5, 46, 0, 0, 310, 314, 5, 39, 0, 0, 311, 313, 5,
		39, 0, 0, 312, 311, 1, 0, 0, 0, 313, 316, 1, 0, 0, 0, 314, 312, 1, 0, 0,
		0, 314, 315, 1, 0, 0, 0, 315, 317, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 317,
		318, 5, 47, 0, 0, 318, 61, 1, 0, 0, 0, 319, 320, 5, 4, 0, 0, 320, 322,
		5, 50, 0, 0, 321, 323, 5, 39, 0, 0, 322, 321, 1, 0, 0, 0, 322, 323, 1,
		0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 325, 5, 54, 0, 0, 325, 326, 5, 39,
		0, 0, 326, 327, 5, 51, 0, 0, 327, 63, 1, 0, 0, 0, 328, 329, 5, 5, 0, 0,
		329, 330, 5, 50, 0, 0, 330, 331, 5, 54, 0, 0, 331, 332, 5, 39, 0, 0, 332,
		333, 5, 51, 0, 0, 333, 65, 1, 0, 0, 0, 334, 335, 5, 39, 0, 0, 335, 67,
		1, 0, 0, 0, 336, 341, 5, 39, 0, 0, 337, 338, 5, 55, 0, 0, 338, 340, 5,
		39, 0, 0, 339, 337, 1, 0, 0, 0, 340, 343, 1, 0, 0, 0, 341, 339, 1, 0, 0,
		0, 341, 342, 1, 0, 0, 0, 342, 69, 1, 0, 0, 0, 343, 341, 1, 0, 0, 0, 344,
		345, 5, 6, 0, 0, 345, 346, 5, 39, 0, 0, 346, 71, 1, 0, 0, 0, 347, 348,
		5, 7, 0, 0, 348, 349, 5, 39, 0, 0, 349, 73, 1, 0, 0, 0, 350, 351, 5, 8,
		0, 0, 351, 352, 3, 82, 41, 0, 352, 75, 1, 0, 0, 0, 353, 359, 5, 9, 0, 0,
		354, 359, 5, 10, 0, 0, 355, 359, 5, 11, 0, 0, 356, 359, 3, 78, 39, 0, 357,
		359, 5, 39, 0, 0, 358, 353, 1, 0, 0, 0, 358, 354, 1, 0, 0, 0, 358, 355,
		1, 0, 0, 0, 358, 356, 1, 0, 0, 0, 358, 357, 1, 0, 0, 0, 359, 77, 1, 0,
		0, 0, 360, 361, 5, 48, 0, 0, 361, 362, 3, 76, 38, 0, 362, 363, 5, 49, 0,
		0, 363, 79, 1, 0, 0, 0, 364, 369, 5, 39, 0, 0, 365, 366, 5, 55, 0, 0, 366,
		368, 5, 39, 0, 0, 367, 365, 1, 0, 0, 0, 368, 371, 1, 0, 0, 0, 369, 367,
		1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 81, 1, 0, 0, 0, 371, 369, 1, 0,
		0, 0, 372, 373, 6, 41, -1, 0, 373, 385, 3, 84, 42, 0, 374, 375, 5, 50,
		0, 0, 375, 376, 3, 82, 41, 0, 376, 377, 5, 51, 0, 0, 377, 385, 1, 0, 0,
		0, 378, 385, 3, 88, 44, 0, 379, 385, 3, 90, 45, 0, 380, 385, 5, 39, 0,
		0, 381, 385, 5, 41, 0, 0, 382, 385, 5, 40, 0, 0, 383, 385, 5, 42, 0, 0,
		384, 372, 1, 0, 0, 0, 384, 374, 1, 0, 0, 0, 384, 378, 1, 0, 0, 0, 384,
		379, 1, 0, 0, 0, 384, 380, 1, 0, 0, 0, 384, 381, 1, 0, 0, 0, 384, 382,
		1, 0, 0, 0, 384, 383, 1, 0, 0, 0, 385, 400, 1, 0, 0, 0, 386, 387, 10, 11,
		0, 0, 387, 388, 7, 0, 0, 0, 388, 399, 3, 82, 41, 12, 389, 390, 10, 10,
		0, 0, 390, 391, 7, 1, 0, 0, 391, 399, 3, 82, 41, 11, 392, 393, 10, 9, 0,
		0, 393, 394, 7, 2, 0, 0, 394, 399, 3, 82, 41, 10, 395, 396, 10, 8, 0, 0,
		396, 397, 7, 3, 0, 0, 397, 399, 3, 82, 41, 9, 398, 386, 1, 0, 0, 0, 398,
		389, 1, 0, 0, 0, 398, 392, 1, 0, 0, 0, 398, 395, 1, 0, 0, 0, 399, 402,
		1, 0, 0, 0, 400, 398, 1, 0, 0, 0, 400, 401, 1, 0, 0, 0, 401, 83, 1, 0,
		0, 0, 402, 400, 1, 0, 0, 0, 403, 411, 3, 86, 43, 0, 404, 405, 5, 54, 0,
		0, 405, 406, 5, 39, 0, 0, 406, 408, 5, 50, 0, 0, 407, 409, 3, 92, 46, 0,
		408, 407, 1, 0, 0, 0, 408, 409, 1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 410,
		412, 5, 51, 0, 0, 411, 404, 1, 0, 0, 0, 412, 413, 1, 0, 0, 0, 413, 411,
		1, 0, 0, 0, 413, 414, 1, 0, 0, 0, 414, 85, 1, 0, 0, 0, 415, 426, 5, 39,
		0, 0, 416, 417, 5, 39, 0, 0, 417, 419, 5, 50, 0, 0, 418, 420, 3, 92, 46,
		0, 419, 418, 1, 0, 0, 0, 419, 420, 1, 0, 0, 0, 420, 421, 1, 0, 0, 0, 421,
		426, 5, 51, 0, 0, 422, 423, 5, 39, 0, 0, 423, 424, 5, 54, 0, 0, 424, 426,
		5, 39, 0, 0, 425, 415, 1, 0, 0, 0, 425, 416, 1, 0, 0, 0, 425, 422, 1, 0,
		0, 0, 426, 87, 1, 0, 0, 0, 427, 430, 5, 39, 0, 0, 428, 429, 5, 54, 0, 0,
		429, 431, 5, 39, 0, 0, 430, 428, 1, 0, 0, 0, 431, 432, 1, 0, 0, 0, 432,
		430, 1, 0, 0, 0, 432, 433, 1, 0, 0, 0, 433, 89, 1, 0, 0, 0, 434, 435, 5,
		39, 0, 0, 435, 437, 5, 50, 0, 0, 436, 438, 3, 92, 46, 0, 437, 436, 1, 0,
		0, 0, 437, 438, 1, 0, 0, 0, 438, 439, 1, 0, 0, 0, 439, 440, 5, 51, 0, 0,
		440, 91, 1, 0, 0, 0, 441, 446, 3, 82, 41, 0, 442, 443, 5, 55, 0, 0, 443,
		445, 3, 82, 41, 0, 444, 442, 1, 0, 0, 0, 445, 448, 1, 0, 0, 0, 446, 444,
		1, 0, 0, 0, 446, 447, 1, 0, 0, 0, 447, 93, 1, 0, 0, 0, 448, 446, 1, 0,
		0, 0, 449, 450, 5, 13, 0, 0, 450, 451, 5, 46, 0, 0, 451, 452, 3, 96, 48,
		0, 452, 453, 5, 47, 0, 0, 453, 95, 1, 0, 0, 0, 454, 456, 3, 98, 49, 0,
		455, 454, 1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457, 455, 1, 0, 0, 0, 457,
		458, 1, 0, 0, 0, 458, 97, 1, 0, 0, 0, 459, 462, 3, 100, 50, 0, 460, 462,
		3, 102, 51, 0, 461, 459, 1, 0, 0, 0, 461, 460, 1, 0, 0, 0, 462, 99, 1,
		0, 0, 0, 463, 464, 5, 14, 0, 0, 464, 465, 5, 50, 0, 0, 465, 466, 5, 40,
		0, 0, 466, 467, 5, 51, 0, 0, 467, 101, 1, 0, 0, 0, 468, 469, 5, 39, 0,
		0, 469, 470, 5, 54, 0, 0, 470, 471, 5, 15, 0, 0, 471, 472, 5, 50, 0, 0,
		472, 473, 5, 41, 0, 0, 473, 474, 5, 51, 0, 0, 474, 103, 1, 0, 0, 0, 475,
		476, 5, 58, 0, 0, 476, 477, 3, 106, 53, 0, 477, 478, 5, 61, 0, 0, 478,
		479, 5, 59, 0, 0, 479, 105, 1, 0, 0, 0, 480, 482, 3, 108, 54, 0, 481, 480,
		1, 0, 0, 0, 482, 485, 1, 0, 0, 0, 483, 481, 1, 0, 0, 0, 483, 484, 1, 0,
		0, 0, 484, 107, 1, 0, 0, 0, 485, 483, 1, 0, 0, 0, 486, 489, 3, 110, 55,
		0, 487, 489, 3, 112, 56, 0, 488, 486, 1, 0, 0, 0, 488, 487, 1, 0, 0, 0,
		489, 109, 1, 0, 0, 0, 490, 491, 5, 38, 0, 0, 491, 492, 5, 36, 0, 0, 492,
		493, 5, 40, 0, 0, 493, 111, 1, 0, 0, 0, 494, 495, 5, 37, 0, 0, 495, 496,
		5, 36, 0, 0, 496, 497, 5, 46, 0, 0, 497, 498, 3, 114, 57, 0, 498, 499,
		5, 47, 0, 0, 499, 113, 1, 0, 0, 0, 500, 501, 5, 39, 0, 0, 501, 502, 5,
		52, 0, 0, 502, 509, 5, 39, 0, 0, 503, 504, 5, 55, 0, 0, 504, 505, 5, 39,
		0, 0, 505, 506, 5, 52, 0, 0, 506, 508, 5, 39, 0, 0, 507, 503, 1, 0, 0,
		0, 508, 511, 1, 0, 0, 0, 509, 507, 1, 0, 0, 0, 509, 510, 1, 0, 0, 0, 510,
		115, 1, 0, 0, 0, 511, 509, 1, 0, 0, 0, 37, 119, 129, 139, 141, 148, 158,
		177, 183, 185, 204, 226, 245, 250, 252, 276, 283, 290, 314, 322, 341, 358,
		369, 384, 398, 400, 408, 413, 419, 425, 432, 437, 446, 457, 461, 483, 488,
		509,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// EvoLangParserInit initializes any static state used to implement EvoLangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEvoLangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EvoLangParserInit() {
	staticData := &EvoLangParserParserStaticData
	staticData.once.Do(evolangparserParserInit)
}

// NewEvoLangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEvoLangParser(input antlr.TokenStream) *EvoLangParser {
	EvoLangParserInit()
	this := new(EvoLangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &EvoLangParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "EvoLangParser.g4"

	return this
}

// EvoLangParser tokens.
const (
	EvoLangParserEOF                      = antlr.TokenEOF
	EvoLangParserMODEL_DEF                = 1
	EvoLangParserEVENTS_DEF               = 2
	EvoLangParserENUM_DEF                 = 3
	EvoLangParserREFERENCES_DEF           = 4
	EvoLangParserLINKED_BY_DEF            = 5
	EvoLangParserBODY_DEF                 = 6
	EvoLangParserAUTHENTICATE_DEF         = 7
	EvoLangParserRETURN_DEF               = 8
	EvoLangParserTYPE_INT_DEF             = 9
	EvoLangParserTYPE_STRING_DEF          = 10
	EvoLangParserTYPE_BOOLEAN_DEF         = 11
	EvoLangParserIF_DEF                   = 12
	EvoLangParserMAIN_DEF                 = 13
	EvoLangParserPRINT_DEF                = 14
	EvoLangParserLISTEN_DEF               = 15
	EvoLangParserQUERY_DEF                = 16
	EvoLangParserRULE_DEF                 = 17
	EvoLangParserMUTATION_DEF             = 18
	EvoLangParserINPUT_DEF                = 19
	EvoLangParserOPERATION_FUNCTION_DEF   = 20
	EvoLangParserOPERATION_MIDDLEWARE_DEF = 21
	EvoLangParserOPERATION_INPUT_DEF      = 22
	EvoLangParserOP_MUL                   = 23
	EvoLangParserOP_DIV                   = 24
	EvoLangParserOP_SUM                   = 25
	EvoLangParserOP_SUB                   = 26
	EvoLangParserOP_GT                    = 27
	EvoLangParserOP_LT                    = 28
	EvoLangParserOP_EQEQ                  = 29
	EvoLangParserOP_NEQ                   = 30
	EvoLangParserOP_GTEQ                  = 31
	EvoLangParserOP_LTEQ                  = 32
	EvoLangParserOP_AND                   = 33
	EvoLangParserOP_OR                    = 34
	EvoLangParserOP_IN                    = 35
	EvoLangParserOP_ASSIGN                = 36
	EvoLangParserDEFINE_DEF               = 37
	EvoLangParserLANG_DEF                 = 38
	EvoLangParserID                       = 39
	EvoLangParserSTRING                   = 40
	EvoLangParserINT                      = 41
	EvoLangParserBOOLEAN                  = 42
	EvoLangParserWS                       = 43
	EvoLangParserCOMMENT                  = 44
	EvoLangParserWHITESPACE               = 45
	EvoLangParserOPEN_BRACKET             = 46
	EvoLangParserCLOSE_BRACKET            = 47
	EvoLangParserOPEN_S_BRACKET           = 48
	EvoLangParserCLOSE_S_BRACKET          = 49
	EvoLangParserOPEN_PAREN               = 50
	EvoLangParserCLOSE_PAREN              = 51
	EvoLangParserCOLON                    = 52
	EvoLangParserHASH                     = 53
	EvoLangParserDOT                      = 54
	EvoLangParserCOMMA                    = 55
	EvoLangParserAT                       = 56
	EvoLangParserPIPE                     = 57
	EvoLangParserOPEN_SANDBOX_MODE        = 58
	EvoLangParserCLOSE_SANDBOX_MODE       = 59
	EvoLangParserSANDBOX                  = 60
	EvoLangParserSANDBOX_CONTENT          = 61
)

// EvoLangParser rules.
const (
	EvoLangParserRULE_prog                    = 0
	EvoLangParserRULE_statement               = 1
	EvoLangParserRULE_modelDef                = 2
	EvoLangParserRULE_modelBody               = 3
	EvoLangParserRULE_modelFieldDef           = 4
	EvoLangParserRULE_ruleDef                 = 5
	EvoLangParserRULE_ruleBody                = 6
	EvoLangParserRULE_ruleStatement           = 7
	EvoLangParserRULE_ruleParameters          = 8
	EvoLangParserRULE_functionBody            = 9
	EvoLangParserRULE_assignmentStatement     = 10
	EvoLangParserRULE_conditionalStatement    = 11
	EvoLangParserRULE_condition               = 12
	EvoLangParserRULE_returnStatement         = 13
	EvoLangParserRULE_ruleConditionExpression = 14
	EvoLangParserRULE_equalityExpression      = 15
	EvoLangParserRULE_inExpression            = 16
	EvoLangParserRULE_ruleExpressionPrimary   = 17
	EvoLangParserRULE_mutationDef             = 18
	EvoLangParserRULE_mutationBody            = 19
	EvoLangParserRULE_mutationStatement       = 20
	EvoLangParserRULE_mutationParameters      = 21
	EvoLangParserRULE_mutationInnerBlock      = 22
	EvoLangParserRULE_middlewareBlock         = 23
	EvoLangParserRULE_functionBlock           = 24
	EvoLangParserRULE_inputDef                = 25
	EvoLangParserRULE_inputBody               = 26
	EvoLangParserRULE_inputFieldDef           = 27
	EvoLangParserRULE_eventDef                = 28
	EvoLangParserRULE_eventChannelHandler     = 29
	EvoLangParserRULE_enumDef                 = 30
	EvoLangParserRULE_references              = 31
	EvoLangParserRULE_linkedBy                = 32
	EvoLangParserRULE_action                  = 33
	EvoLangParserRULE_paramList               = 34
	EvoLangParserRULE_bodyDef                 = 35
	EvoLangParserRULE_authenticateDef         = 36
	EvoLangParserRULE_returnDef               = 37
	EvoLangParserRULE_typeDef                 = 38
	EvoLangParserRULE_arrayDef                = 39
	EvoLangParserRULE_eventType               = 40
	EvoLangParserRULE_expression              = 41
	EvoLangParserRULE_methodCallExpression    = 42
	EvoLangParserRULE_primaryExpression       = 43
	EvoLangParserRULE_propertyAccess          = 44
	EvoLangParserRULE_functionCall            = 45
	EvoLangParserRULE_exprList                = 46
	EvoLangParserRULE_mainFunction            = 47
	EvoLangParserRULE_mainStatements          = 48
	EvoLangParserRULE_mainStatement           = 49
	EvoLangParserRULE_printStatement          = 50
	EvoLangParserRULE_serverStatement         = 51
	EvoLangParserRULE_sandboxBlock            = 52
	EvoLangParserRULE_sandboxAttributes       = 53
	EvoLangParserRULE_sandboxAttribute        = 54
	EvoLangParserRULE_sandboxLangAttribute    = 55
	EvoLangParserRULE_sandboxDefineAttribute  = 56
	EvoLangParserRULE_sandboxDefineContent    = 57
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserEOF, 0)
}

func (s *ProgContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EvoLangParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&925706) != 0) {
		{
			p.SetState(116)
			p.Statement()
		}

		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(121)
		p.Match(EvoLangParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ModelDef() IModelDefContext
	EnumDef() IEnumDefContext
	RuleDef() IRuleDefContext
	MutationDef() IMutationDefContext
	InputDef() IInputDefContext
	MainFunction() IMainFunctionContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ModelDef() IModelDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModelDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModelDefContext)
}

func (s *StatementContext) EnumDef() IEnumDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumDefContext)
}

func (s *StatementContext) RuleDef() IRuleDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleDefContext)
}

func (s *StatementContext) MutationDef() IMutationDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMutationDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMutationDefContext)
}

func (s *StatementContext) InputDef() IInputDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputDefContext)
}

func (s *StatementContext) MainFunction() IMainFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMainFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMainFunctionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EvoLangParserRULE_statement)
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EvoLangParserMODEL_DEF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.ModelDef()
		}

	case EvoLangParserENUM_DEF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.EnumDef()
		}

	case EvoLangParserRULE_DEF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			p.RuleDef()
		}

	case EvoLangParserMUTATION_DEF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(126)
			p.MutationDef()
		}

	case EvoLangParserINPUT_DEF:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(127)
			p.InputDef()
		}

	case EvoLangParserMAIN_DEF:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(128)
			p.MainFunction()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IModelDefContext is an interface to support dynamic dispatch.
type IModelDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MODEL_DEF() antlr.TerminalNode
	ID() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	ModelBody() IModelBodyContext
	CLOSE_BRACKET() antlr.TerminalNode

	// IsModelDefContext differentiates from other interfaces.
	IsModelDefContext()
}

type ModelDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelDefContext() *ModelDefContext {
	var p = new(ModelDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_modelDef
	return p
}

func InitEmptyModelDefContext(p *ModelDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_modelDef
}

func (*ModelDefContext) IsModelDefContext() {}

func NewModelDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelDefContext {
	var p = new(ModelDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_modelDef

	return p
}

func (s *ModelDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ModelDefContext) MODEL_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserMODEL_DEF, 0)
}

func (s *ModelDefContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *ModelDefContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_BRACKET, 0)
}

func (s *ModelDefContext) ModelBody() IModelBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModelBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModelBodyContext)
}

func (s *ModelDefContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_BRACKET, 0)
}

func (s *ModelDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModelDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterModelDef(s)
	}
}

func (s *ModelDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitModelDef(s)
	}
}

func (s *ModelDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitModelDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) ModelDef() (localctx IModelDefContext) {
	localctx = NewModelDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, EvoLangParserRULE_modelDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(EvoLangParserMODEL_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.ModelBody()
	}
	{
		p.SetState(135)
		p.Match(EvoLangParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IModelBodyContext is an interface to support dynamic dispatch.
type IModelBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllModelFieldDef() []IModelFieldDefContext
	ModelFieldDef(i int) IModelFieldDefContext
	AllEventDef() []IEventDefContext
	EventDef(i int) IEventDefContext

	// IsModelBodyContext differentiates from other interfaces.
	IsModelBodyContext()
}

type ModelBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelBodyContext() *ModelBodyContext {
	var p = new(ModelBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_modelBody
	return p
}

func InitEmptyModelBodyContext(p *ModelBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_modelBody
}

func (*ModelBodyContext) IsModelBodyContext() {}

func NewModelBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelBodyContext {
	var p = new(ModelBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_modelBody

	return p
}

func (s *ModelBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ModelBodyContext) AllModelFieldDef() []IModelFieldDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IModelFieldDefContext); ok {
			len++
		}
	}

	tst := make([]IModelFieldDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IModelFieldDefContext); ok {
			tst[i] = t.(IModelFieldDefContext)
			i++
		}
	}

	return tst
}

func (s *ModelBodyContext) ModelFieldDef(i int) IModelFieldDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModelFieldDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModelFieldDefContext)
}

func (s *ModelBodyContext) AllEventDef() []IEventDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEventDefContext); ok {
			len++
		}
	}

	tst := make([]IEventDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEventDefContext); ok {
			tst[i] = t.(IEventDefContext)
			i++
		}
	}

	return tst
}

func (s *ModelBodyContext) EventDef(i int) IEventDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEventDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEventDefContext)
}

func (s *ModelBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModelBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterModelBody(s)
	}
}

func (s *ModelBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitModelBody(s)
	}
}

func (s *ModelBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitModelBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) ModelBody() (localctx IModelBodyContext) {
	localctx = NewModelBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EvoLangParserRULE_modelBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EvoLangParserEVENTS_DEF || _la == EvoLangParserID {
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case EvoLangParserID:
			{
				p.SetState(137)
				p.ModelFieldDef()
			}

		case EvoLangParserEVENTS_DEF:
			{
				p.SetState(138)
				p.EventDef()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IModelFieldDefContext is an interface to support dynamic dispatch.
type IModelFieldDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	TypeDef() ITypeDefContext
	References() IReferencesContext
	LinkedBy() ILinkedByContext

	// IsModelFieldDefContext differentiates from other interfaces.
	IsModelFieldDefContext()
}

type ModelFieldDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelFieldDefContext() *ModelFieldDefContext {
	var p = new(ModelFieldDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_modelFieldDef
	return p
}

func InitEmptyModelFieldDefContext(p *ModelFieldDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_modelFieldDef
}

func (*ModelFieldDefContext) IsModelFieldDefContext() {}

func NewModelFieldDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelFieldDefContext {
	var p = new(ModelFieldDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_modelFieldDef

	return p
}

func (s *ModelFieldDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ModelFieldDefContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *ModelFieldDefContext) COLON() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCOLON, 0)
}

func (s *ModelFieldDefContext) TypeDef() ITypeDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDefContext)
}

func (s *ModelFieldDefContext) References() IReferencesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferencesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferencesContext)
}

func (s *ModelFieldDefContext) LinkedBy() ILinkedByContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILinkedByContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILinkedByContext)
}

func (s *ModelFieldDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelFieldDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModelFieldDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterModelFieldDef(s)
	}
}

func (s *ModelFieldDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitModelFieldDef(s)
	}
}

func (s *ModelFieldDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitModelFieldDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) ModelFieldDef() (localctx IModelFieldDefContext) {
	localctx = NewModelFieldDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EvoLangParserRULE_modelFieldDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(EvoLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.TypeDef()
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case EvoLangParserREFERENCES_DEF:
		{
			p.SetState(146)
			p.References()
		}

	case EvoLangParserLINKED_BY_DEF:
		{
			p.SetState(147)
			p.LinkedBy()
		}

	case EvoLangParserEVENTS_DEF, EvoLangParserID, EvoLangParserCLOSE_BRACKET:

	default:
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRuleDefContext is an interface to support dynamic dispatch.
type IRuleDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RULE_DEF() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	RuleBody() IRuleBodyContext
	CLOSE_BRACKET() antlr.TerminalNode

	// IsRuleDefContext differentiates from other interfaces.
	IsRuleDefContext()
}

type RuleDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleDefContext() *RuleDefContext {
	var p = new(RuleDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_ruleDef
	return p
}

func InitEmptyRuleDefContext(p *RuleDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_ruleDef
}

func (*RuleDefContext) IsRuleDefContext() {}

func NewRuleDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleDefContext {
	var p = new(RuleDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_ruleDef

	return p
}

func (s *RuleDefContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleDefContext) RULE_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserRULE_DEF, 0)
}

func (s *RuleDefContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_BRACKET, 0)
}

func (s *RuleDefContext) RuleBody() IRuleBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleBodyContext)
}

func (s *RuleDefContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_BRACKET, 0)
}

func (s *RuleDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterRuleDef(s)
	}
}

func (s *RuleDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitRuleDef(s)
	}
}

func (s *RuleDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitRuleDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) RuleDef() (localctx IRuleDefContext) {
	localctx = NewRuleDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EvoLangParserRULE_ruleDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(EvoLangParserRULE_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.RuleBody()
	}
	{
		p.SetState(153)
		p.Match(EvoLangParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRuleBodyContext is an interface to support dynamic dispatch.
type IRuleBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRuleStatement() []IRuleStatementContext
	RuleStatement(i int) IRuleStatementContext

	// IsRuleBodyContext differentiates from other interfaces.
	IsRuleBodyContext()
}

type RuleBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleBodyContext() *RuleBodyContext {
	var p = new(RuleBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_ruleBody
	return p
}

func InitEmptyRuleBodyContext(p *RuleBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_ruleBody
}

func (*RuleBodyContext) IsRuleBodyContext() {}

func NewRuleBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleBodyContext {
	var p = new(RuleBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_ruleBody

	return p
}

func (s *RuleBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleBodyContext) AllRuleStatement() []IRuleStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRuleStatementContext); ok {
			len++
		}
	}

	tst := make([]IRuleStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRuleStatementContext); ok {
			tst[i] = t.(IRuleStatementContext)
			i++
		}
	}

	return tst
}

func (s *RuleBodyContext) RuleStatement(i int) IRuleStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleStatementContext)
}

func (s *RuleBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterRuleBody(s)
	}
}

func (s *RuleBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitRuleBody(s)
	}
}

func (s *RuleBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitRuleBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) RuleBody() (localctx IRuleBodyContext) {
	localctx = NewRuleBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, EvoLangParserRULE_ruleBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserID {
		{
			p.SetState(155)
			p.RuleStatement()
		}

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRuleStatementContext is an interface to support dynamic dispatch.
type IRuleStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	OPEN_PAREN() antlr.TerminalNode
	RuleParameters() IRuleParametersContext
	CLOSE_PAREN() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	FunctionBody() IFunctionBodyContext
	CLOSE_BRACKET() antlr.TerminalNode

	// IsRuleStatementContext differentiates from other interfaces.
	IsRuleStatementContext()
}

type RuleStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleStatementContext() *RuleStatementContext {
	var p = new(RuleStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_ruleStatement
	return p
}

func InitEmptyRuleStatementContext(p *RuleStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_ruleStatement
}

func (*RuleStatementContext) IsRuleStatementContext() {}

func NewRuleStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleStatementContext {
	var p = new(RuleStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_ruleStatement

	return p
}

func (s *RuleStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *RuleStatementContext) COLON() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCOLON, 0)
}

func (s *RuleStatementContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_PAREN, 0)
}

func (s *RuleStatementContext) RuleParameters() IRuleParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleParametersContext)
}

func (s *RuleStatementContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_PAREN, 0)
}

func (s *RuleStatementContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_BRACKET, 0)
}

func (s *RuleStatementContext) FunctionBody() IFunctionBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *RuleStatementContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_BRACKET, 0)
}

func (s *RuleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterRuleStatement(s)
	}
}

func (s *RuleStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitRuleStatement(s)
	}
}

func (s *RuleStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitRuleStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) RuleStatement() (localctx IRuleStatementContext) {
	localctx = NewRuleStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, EvoLangParserRULE_ruleStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Match(EvoLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Match(EvoLangParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)
		p.RuleParameters()
	}
	{
		p.SetState(165)
		p.Match(EvoLangParserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.FunctionBody()
	}
	{
		p.SetState(168)
		p.Match(EvoLangParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRuleParametersContext is an interface to support dynamic dispatch.
type IRuleParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsRuleParametersContext differentiates from other interfaces.
	IsRuleParametersContext()
}

type RuleParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleParametersContext() *RuleParametersContext {
	var p = new(RuleParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_ruleParameters
	return p
}

func InitEmptyRuleParametersContext(p *RuleParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_ruleParameters
}

func (*RuleParametersContext) IsRuleParametersContext() {}

func NewRuleParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleParametersContext {
	var p = new(RuleParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_ruleParameters

	return p
}

func (s *RuleParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleParametersContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserID)
}

func (s *RuleParametersContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, i)
}

func (s *RuleParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserCOMMA)
}

func (s *RuleParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserCOMMA, i)
}

func (s *RuleParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterRuleParameters(s)
	}
}

func (s *RuleParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitRuleParameters(s)
	}
}

func (s *RuleParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitRuleParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) RuleParameters() (localctx IRuleParametersContext) {
	localctx = NewRuleParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, EvoLangParserRULE_ruleParameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserCOMMA {
		{
			p.SetState(172)
			p.Match(EvoLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionBodyContext is an interface to support dynamic dispatch.
type IFunctionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAssignmentStatement() []IAssignmentStatementContext
	AssignmentStatement(i int) IAssignmentStatementContext
	AllConditionalStatement() []IConditionalStatementContext
	ConditionalStatement(i int) IConditionalStatementContext
	AllReturnStatement() []IReturnStatementContext
	ReturnStatement(i int) IReturnStatementContext

	// IsFunctionBodyContext differentiates from other interfaces.
	IsFunctionBodyContext()
}

type FunctionBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionBodyContext() *FunctionBodyContext {
	var p = new(FunctionBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_functionBody
	return p
}

func InitEmptyFunctionBodyContext(p *FunctionBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_functionBody
}

func (*FunctionBodyContext) IsFunctionBodyContext() {}

func NewFunctionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBodyContext {
	var p = new(FunctionBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_functionBody

	return p
}

func (s *FunctionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionBodyContext) AllAssignmentStatement() []IAssignmentStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentStatementContext); ok {
			tst[i] = t.(IAssignmentStatementContext)
			i++
		}
	}

	return tst
}

func (s *FunctionBodyContext) AssignmentStatement(i int) IAssignmentStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *FunctionBodyContext) AllConditionalStatement() []IConditionalStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionalStatementContext); ok {
			len++
		}
	}

	tst := make([]IConditionalStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionalStatementContext); ok {
			tst[i] = t.(IConditionalStatementContext)
			i++
		}
	}

	return tst
}

func (s *FunctionBodyContext) ConditionalStatement(i int) IConditionalStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalStatementContext)
}

func (s *FunctionBodyContext) AllReturnStatement() []IReturnStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IReturnStatementContext); ok {
			len++
		}
	}

	tst := make([]IReturnStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IReturnStatementContext); ok {
			tst[i] = t.(IReturnStatementContext)
			i++
		}
	}

	return tst
}

func (s *FunctionBodyContext) ReturnStatement(i int) IReturnStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *FunctionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterFunctionBody(s)
	}
}

func (s *FunctionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitFunctionBody(s)
	}
}

func (s *FunctionBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitFunctionBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) FunctionBody() (localctx IFunctionBodyContext) {
	localctx = NewFunctionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, EvoLangParserRULE_functionBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&549755818240) != 0 {
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case EvoLangParserID:
			{
				p.SetState(180)
				p.AssignmentStatement()
			}

		case EvoLangParserIF_DEF:
			{
				p.SetState(181)
				p.ConditionalStatement()
			}

		case EvoLangParserRETURN_DEF:
			{
				p.SetState(182)
				p.ReturnStatement()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentStatementContext is an interface to support dynamic dispatch.
type IAssignmentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	OP_ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssignmentStatementContext differentiates from other interfaces.
	IsAssignmentStatementContext()
}

type AssignmentStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentStatementContext() *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_assignmentStatement
	return p
}

func InitEmptyAssignmentStatementContext(p *AssignmentStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_assignmentStatement
}

func (*AssignmentStatementContext) IsAssignmentStatementContext() {}

func NewAssignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_assignmentStatement

	return p
}

func (s *AssignmentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *AssignmentStatementContext) OP_ASSIGN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_ASSIGN, 0)
}

func (s *AssignmentStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitAssignmentStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, EvoLangParserRULE_assignmentStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(189)
		p.Match(EvoLangParserOP_ASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(190)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionalStatementContext is an interface to support dynamic dispatch.
type IConditionalStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF_DEF() antlr.TerminalNode
	Condition() IConditionContext
	OPEN_BRACKET() antlr.TerminalNode
	ReturnStatement() IReturnStatementContext
	CLOSE_BRACKET() antlr.TerminalNode

	// IsConditionalStatementContext differentiates from other interfaces.
	IsConditionalStatementContext()
}

type ConditionalStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalStatementContext() *ConditionalStatementContext {
	var p = new(ConditionalStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_conditionalStatement
	return p
}

func InitEmptyConditionalStatementContext(p *ConditionalStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_conditionalStatement
}

func (*ConditionalStatementContext) IsConditionalStatementContext() {}

func NewConditionalStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalStatementContext {
	var p = new(ConditionalStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_conditionalStatement

	return p
}

func (s *ConditionalStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalStatementContext) IF_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserIF_DEF, 0)
}

func (s *ConditionalStatementContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ConditionalStatementContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_BRACKET, 0)
}

func (s *ConditionalStatementContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *ConditionalStatementContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_BRACKET, 0)
}

func (s *ConditionalStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterConditionalStatement(s)
	}
}

func (s *ConditionalStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitConditionalStatement(s)
	}
}

func (s *ConditionalStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitConditionalStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) ConditionalStatement() (localctx IConditionalStatementContext) {
	localctx = NewConditionalStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, EvoLangParserRULE_conditionalStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(EvoLangParserIF_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
		p.Condition()
	}
	{
		p.SetState(194)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)
		p.ReturnStatement()
	}
	{
		p.SetState(196)
		p.Match(EvoLangParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RuleConditionExpression() IRuleConditionExpressionContext

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) RuleConditionExpression() IRuleConditionExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleConditionExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleConditionExpressionContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, EvoLangParserRULE_condition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.RuleConditionExpression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ReturnDef() IReturnDefContext

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_returnStatement
	return p
}

func InitEmptyReturnStatementContext(p *ReturnStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_returnStatement
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) ReturnDef() IReturnDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnDefContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, EvoLangParserRULE_returnStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.ReturnDef()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRuleConditionExpressionContext is an interface to support dynamic dispatch.
type IRuleConditionExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EqualityExpression() IEqualityExpressionContext
	InExpression() IInExpressionContext

	// IsRuleConditionExpressionContext differentiates from other interfaces.
	IsRuleConditionExpressionContext()
}

type RuleConditionExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleConditionExpressionContext() *RuleConditionExpressionContext {
	var p = new(RuleConditionExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_ruleConditionExpression
	return p
}

func InitEmptyRuleConditionExpressionContext(p *RuleConditionExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_ruleConditionExpression
}

func (*RuleConditionExpressionContext) IsRuleConditionExpressionContext() {}

func NewRuleConditionExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleConditionExpressionContext {
	var p = new(RuleConditionExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_ruleConditionExpression

	return p
}

func (s *RuleConditionExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleConditionExpressionContext) EqualityExpression() IEqualityExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *RuleConditionExpressionContext) InExpression() IInExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInExpressionContext)
}

func (s *RuleConditionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleConditionExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleConditionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterRuleConditionExpression(s)
	}
}

func (s *RuleConditionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitRuleConditionExpression(s)
	}
}

func (s *RuleConditionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitRuleConditionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) RuleConditionExpression() (localctx IRuleConditionExpressionContext) {
	localctx = NewRuleConditionExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, EvoLangParserRULE_ruleConditionExpression)
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
			p.EqualityExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(203)
			p.InExpression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEqualityExpressionContext is an interface to support dynamic dispatch.
type IEqualityExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRuleExpressionPrimary() []IRuleExpressionPrimaryContext
	RuleExpressionPrimary(i int) IRuleExpressionPrimaryContext
	OP_EQEQ() antlr.TerminalNode

	// IsEqualityExpressionContext differentiates from other interfaces.
	IsEqualityExpressionContext()
}

type EqualityExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityExpressionContext() *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_equalityExpression
	return p
}

func InitEmptyEqualityExpressionContext(p *EqualityExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_equalityExpression
}

func (*EqualityExpressionContext) IsEqualityExpressionContext() {}

func NewEqualityExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_equalityExpression

	return p
}

func (s *EqualityExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExpressionContext) AllRuleExpressionPrimary() []IRuleExpressionPrimaryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRuleExpressionPrimaryContext); ok {
			len++
		}
	}

	tst := make([]IRuleExpressionPrimaryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRuleExpressionPrimaryContext); ok {
			tst[i] = t.(IRuleExpressionPrimaryContext)
			i++
		}
	}

	return tst
}

func (s *EqualityExpressionContext) RuleExpressionPrimary(i int) IRuleExpressionPrimaryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleExpressionPrimaryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleExpressionPrimaryContext)
}

func (s *EqualityExpressionContext) OP_EQEQ() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_EQEQ, 0)
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitEqualityExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) EqualityExpression() (localctx IEqualityExpressionContext) {
	localctx = NewEqualityExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, EvoLangParserRULE_equalityExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.RuleExpressionPrimary()
	}
	{
		p.SetState(207)
		p.Match(EvoLangParserOP_EQEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.RuleExpressionPrimary()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInExpressionContext is an interface to support dynamic dispatch.
type IInExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRuleExpressionPrimary() []IRuleExpressionPrimaryContext
	RuleExpressionPrimary(i int) IRuleExpressionPrimaryContext
	OP_IN() antlr.TerminalNode

	// IsInExpressionContext differentiates from other interfaces.
	IsInExpressionContext()
}

type InExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInExpressionContext() *InExpressionContext {
	var p = new(InExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_inExpression
	return p
}

func InitEmptyInExpressionContext(p *InExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_inExpression
}

func (*InExpressionContext) IsInExpressionContext() {}

func NewInExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InExpressionContext {
	var p = new(InExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_inExpression

	return p
}

func (s *InExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *InExpressionContext) AllRuleExpressionPrimary() []IRuleExpressionPrimaryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRuleExpressionPrimaryContext); ok {
			len++
		}
	}

	tst := make([]IRuleExpressionPrimaryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRuleExpressionPrimaryContext); ok {
			tst[i] = t.(IRuleExpressionPrimaryContext)
			i++
		}
	}

	return tst
}

func (s *InExpressionContext) RuleExpressionPrimary(i int) IRuleExpressionPrimaryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleExpressionPrimaryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleExpressionPrimaryContext)
}

func (s *InExpressionContext) OP_IN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_IN, 0)
}

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterInExpression(s)
	}
}

func (s *InExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitInExpression(s)
	}
}

func (s *InExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitInExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) InExpression() (localctx IInExpressionContext) {
	localctx = NewInExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, EvoLangParserRULE_inExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.RuleExpressionPrimary()
	}
	{
		p.SetState(211)
		p.Match(EvoLangParserOP_IN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(212)
		p.RuleExpressionPrimary()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRuleExpressionPrimaryContext is an interface to support dynamic dispatch.
type IRuleExpressionPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	DOT() antlr.TerminalNode

	// IsRuleExpressionPrimaryContext differentiates from other interfaces.
	IsRuleExpressionPrimaryContext()
}

type RuleExpressionPrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleExpressionPrimaryContext() *RuleExpressionPrimaryContext {
	var p = new(RuleExpressionPrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_ruleExpressionPrimary
	return p
}

func InitEmptyRuleExpressionPrimaryContext(p *RuleExpressionPrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_ruleExpressionPrimary
}

func (*RuleExpressionPrimaryContext) IsRuleExpressionPrimaryContext() {}

func NewRuleExpressionPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleExpressionPrimaryContext {
	var p = new(RuleExpressionPrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_ruleExpressionPrimary

	return p
}

func (s *RuleExpressionPrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleExpressionPrimaryContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserID)
}

func (s *RuleExpressionPrimaryContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, i)
}

func (s *RuleExpressionPrimaryContext) DOT() antlr.TerminalNode {
	return s.GetToken(EvoLangParserDOT, 0)
}

func (s *RuleExpressionPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleExpressionPrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleExpressionPrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterRuleExpressionPrimary(s)
	}
}

func (s *RuleExpressionPrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitRuleExpressionPrimary(s)
	}
}

func (s *RuleExpressionPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitRuleExpressionPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) RuleExpressionPrimary() (localctx IRuleExpressionPrimaryContext) {
	localctx = NewRuleExpressionPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, EvoLangParserRULE_ruleExpressionPrimary)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Match(EvoLangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMutationDefContext is an interface to support dynamic dispatch.
type IMutationDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUTATION_DEF() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	MutationBody() IMutationBodyContext
	CLOSE_BRACKET() antlr.TerminalNode

	// IsMutationDefContext differentiates from other interfaces.
	IsMutationDefContext()
}

type MutationDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMutationDefContext() *MutationDefContext {
	var p = new(MutationDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mutationDef
	return p
}

func InitEmptyMutationDefContext(p *MutationDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mutationDef
}

func (*MutationDefContext) IsMutationDefContext() {}

func NewMutationDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MutationDefContext {
	var p = new(MutationDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_mutationDef

	return p
}

func (s *MutationDefContext) GetParser() antlr.Parser { return s.parser }

func (s *MutationDefContext) MUTATION_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserMUTATION_DEF, 0)
}

func (s *MutationDefContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_BRACKET, 0)
}

func (s *MutationDefContext) MutationBody() IMutationBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMutationBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMutationBodyContext)
}

func (s *MutationDefContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_BRACKET, 0)
}

func (s *MutationDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutationDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MutationDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterMutationDef(s)
	}
}

func (s *MutationDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitMutationDef(s)
	}
}

func (s *MutationDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitMutationDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) MutationDef() (localctx IMutationDefContext) {
	localctx = NewMutationDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, EvoLangParserRULE_mutationDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(EvoLangParserMUTATION_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.MutationBody()
	}
	{
		p.SetState(221)
		p.Match(EvoLangParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMutationBodyContext is an interface to support dynamic dispatch.
type IMutationBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMutationStatement() []IMutationStatementContext
	MutationStatement(i int) IMutationStatementContext

	// IsMutationBodyContext differentiates from other interfaces.
	IsMutationBodyContext()
}

type MutationBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMutationBodyContext() *MutationBodyContext {
	var p = new(MutationBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mutationBody
	return p
}

func InitEmptyMutationBodyContext(p *MutationBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mutationBody
}

func (*MutationBodyContext) IsMutationBodyContext() {}

func NewMutationBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MutationBodyContext {
	var p = new(MutationBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_mutationBody

	return p
}

func (s *MutationBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MutationBodyContext) AllMutationStatement() []IMutationStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMutationStatementContext); ok {
			len++
		}
	}

	tst := make([]IMutationStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMutationStatementContext); ok {
			tst[i] = t.(IMutationStatementContext)
			i++
		}
	}

	return tst
}

func (s *MutationBodyContext) MutationStatement(i int) IMutationStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMutationStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMutationStatementContext)
}

func (s *MutationBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutationBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MutationBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterMutationBody(s)
	}
}

func (s *MutationBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitMutationBody(s)
	}
}

func (s *MutationBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitMutationBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) MutationBody() (localctx IMutationBodyContext) {
	localctx = NewMutationBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, EvoLangParserRULE_mutationBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserID {
		{
			p.SetState(223)
			p.MutationStatement()
		}

		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMutationStatementContext is an interface to support dynamic dispatch.
type IMutationStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	OPEN_PAREN() antlr.TerminalNode
	MutationParameters() IMutationParametersContext
	CLOSE_PAREN() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	MutationInnerBlock() IMutationInnerBlockContext
	CLOSE_BRACKET() antlr.TerminalNode

	// IsMutationStatementContext differentiates from other interfaces.
	IsMutationStatementContext()
}

type MutationStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMutationStatementContext() *MutationStatementContext {
	var p = new(MutationStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mutationStatement
	return p
}

func InitEmptyMutationStatementContext(p *MutationStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mutationStatement
}

func (*MutationStatementContext) IsMutationStatementContext() {}

func NewMutationStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MutationStatementContext {
	var p = new(MutationStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_mutationStatement

	return p
}

func (s *MutationStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *MutationStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *MutationStatementContext) COLON() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCOLON, 0)
}

func (s *MutationStatementContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_PAREN, 0)
}

func (s *MutationStatementContext) MutationParameters() IMutationParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMutationParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMutationParametersContext)
}

func (s *MutationStatementContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_PAREN, 0)
}

func (s *MutationStatementContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_BRACKET, 0)
}

func (s *MutationStatementContext) MutationInnerBlock() IMutationInnerBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMutationInnerBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMutationInnerBlockContext)
}

func (s *MutationStatementContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_BRACKET, 0)
}

func (s *MutationStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutationStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MutationStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterMutationStatement(s)
	}
}

func (s *MutationStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitMutationStatement(s)
	}
}

func (s *MutationStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitMutationStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) MutationStatement() (localctx IMutationStatementContext) {
	localctx = NewMutationStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, EvoLangParserRULE_mutationStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.Match(EvoLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
		p.Match(EvoLangParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.MutationParameters()
	}
	{
		p.SetState(233)
		p.Match(EvoLangParserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(234)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)
		p.MutationInnerBlock()
	}
	{
		p.SetState(236)
		p.Match(EvoLangParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMutationParametersContext is an interface to support dynamic dispatch.
type IMutationParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsMutationParametersContext differentiates from other interfaces.
	IsMutationParametersContext()
}

type MutationParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMutationParametersContext() *MutationParametersContext {
	var p = new(MutationParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mutationParameters
	return p
}

func InitEmptyMutationParametersContext(p *MutationParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mutationParameters
}

func (*MutationParametersContext) IsMutationParametersContext() {}

func NewMutationParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MutationParametersContext {
	var p = new(MutationParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_mutationParameters

	return p
}

func (s *MutationParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *MutationParametersContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserID)
}

func (s *MutationParametersContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, i)
}

func (s *MutationParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserCOMMA)
}

func (s *MutationParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserCOMMA, i)
}

func (s *MutationParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutationParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MutationParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterMutationParameters(s)
	}
}

func (s *MutationParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitMutationParameters(s)
	}
}

func (s *MutationParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitMutationParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) MutationParameters() (localctx IMutationParametersContext) {
	localctx = NewMutationParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, EvoLangParserRULE_mutationParameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserCOMMA {
		{
			p.SetState(240)
			p.Match(EvoLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMutationInnerBlockContext is an interface to support dynamic dispatch.
type IMutationInnerBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMiddlewareBlock() []IMiddlewareBlockContext
	MiddlewareBlock(i int) IMiddlewareBlockContext
	AllFunctionBlock() []IFunctionBlockContext
	FunctionBlock(i int) IFunctionBlockContext

	// IsMutationInnerBlockContext differentiates from other interfaces.
	IsMutationInnerBlockContext()
}

type MutationInnerBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMutationInnerBlockContext() *MutationInnerBlockContext {
	var p = new(MutationInnerBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mutationInnerBlock
	return p
}

func InitEmptyMutationInnerBlockContext(p *MutationInnerBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mutationInnerBlock
}

func (*MutationInnerBlockContext) IsMutationInnerBlockContext() {}

func NewMutationInnerBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MutationInnerBlockContext {
	var p = new(MutationInnerBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_mutationInnerBlock

	return p
}

func (s *MutationInnerBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *MutationInnerBlockContext) AllMiddlewareBlock() []IMiddlewareBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMiddlewareBlockContext); ok {
			len++
		}
	}

	tst := make([]IMiddlewareBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMiddlewareBlockContext); ok {
			tst[i] = t.(IMiddlewareBlockContext)
			i++
		}
	}

	return tst
}

func (s *MutationInnerBlockContext) MiddlewareBlock(i int) IMiddlewareBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMiddlewareBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMiddlewareBlockContext)
}

func (s *MutationInnerBlockContext) AllFunctionBlock() []IFunctionBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionBlockContext); ok {
			len++
		}
	}

	tst := make([]IFunctionBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionBlockContext); ok {
			tst[i] = t.(IFunctionBlockContext)
			i++
		}
	}

	return tst
}

func (s *MutationInnerBlockContext) FunctionBlock(i int) IFunctionBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionBlockContext)
}

func (s *MutationInnerBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutationInnerBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MutationInnerBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterMutationInnerBlock(s)
	}
}

func (s *MutationInnerBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitMutationInnerBlock(s)
	}
}

func (s *MutationInnerBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitMutationInnerBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) MutationInnerBlock() (localctx IMutationInnerBlockContext) {
	localctx = NewMutationInnerBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, EvoLangParserRULE_mutationInnerBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserOPERATION_FUNCTION_DEF || _la == EvoLangParserOPERATION_MIDDLEWARE_DEF {
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case EvoLangParserOPERATION_MIDDLEWARE_DEF:
			{
				p.SetState(248)
				p.MiddlewareBlock()
			}

		case EvoLangParserOPERATION_FUNCTION_DEF:
			{
				p.SetState(249)
				p.FunctionBlock()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMiddlewareBlockContext is an interface to support dynamic dispatch.
type IMiddlewareBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPERATION_MIDDLEWARE_DEF() antlr.TerminalNode
	COLON() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	FunctionBody() IFunctionBodyContext
	CLOSE_BRACKET() antlr.TerminalNode

	// IsMiddlewareBlockContext differentiates from other interfaces.
	IsMiddlewareBlockContext()
}

type MiddlewareBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMiddlewareBlockContext() *MiddlewareBlockContext {
	var p = new(MiddlewareBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_middlewareBlock
	return p
}

func InitEmptyMiddlewareBlockContext(p *MiddlewareBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_middlewareBlock
}

func (*MiddlewareBlockContext) IsMiddlewareBlockContext() {}

func NewMiddlewareBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MiddlewareBlockContext {
	var p = new(MiddlewareBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_middlewareBlock

	return p
}

func (s *MiddlewareBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *MiddlewareBlockContext) OPERATION_MIDDLEWARE_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPERATION_MIDDLEWARE_DEF, 0)
}

func (s *MiddlewareBlockContext) COLON() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCOLON, 0)
}

func (s *MiddlewareBlockContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_BRACKET, 0)
}

func (s *MiddlewareBlockContext) FunctionBody() IFunctionBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *MiddlewareBlockContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_BRACKET, 0)
}

func (s *MiddlewareBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MiddlewareBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MiddlewareBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterMiddlewareBlock(s)
	}
}

func (s *MiddlewareBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitMiddlewareBlock(s)
	}
}

func (s *MiddlewareBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitMiddlewareBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) MiddlewareBlock() (localctx IMiddlewareBlockContext) {
	localctx = NewMiddlewareBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, EvoLangParserRULE_middlewareBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.Match(EvoLangParserOPERATION_MIDDLEWARE_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(256)
		p.Match(EvoLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.FunctionBody()
	}
	{
		p.SetState(259)
		p.Match(EvoLangParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionBlockContext is an interface to support dynamic dispatch.
type IFunctionBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPERATION_FUNCTION_DEF() antlr.TerminalNode
	COLON() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	SANDBOX() antlr.TerminalNode
	CLOSE_BRACKET() antlr.TerminalNode

	// IsFunctionBlockContext differentiates from other interfaces.
	IsFunctionBlockContext()
}

type FunctionBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionBlockContext() *FunctionBlockContext {
	var p = new(FunctionBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_functionBlock
	return p
}

func InitEmptyFunctionBlockContext(p *FunctionBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_functionBlock
}

func (*FunctionBlockContext) IsFunctionBlockContext() {}

func NewFunctionBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBlockContext {
	var p = new(FunctionBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_functionBlock

	return p
}

func (s *FunctionBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionBlockContext) OPERATION_FUNCTION_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPERATION_FUNCTION_DEF, 0)
}

func (s *FunctionBlockContext) COLON() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCOLON, 0)
}

func (s *FunctionBlockContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_BRACKET, 0)
}

func (s *FunctionBlockContext) SANDBOX() antlr.TerminalNode {
	return s.GetToken(EvoLangParserSANDBOX, 0)
}

func (s *FunctionBlockContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_BRACKET, 0)
}

func (s *FunctionBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterFunctionBlock(s)
	}
}

func (s *FunctionBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitFunctionBlock(s)
	}
}

func (s *FunctionBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitFunctionBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) FunctionBlock() (localctx IFunctionBlockContext) {
	localctx = NewFunctionBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, EvoLangParserRULE_functionBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Match(EvoLangParserOPERATION_FUNCTION_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.Match(EvoLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(263)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)
		p.Match(EvoLangParserSANDBOX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(265)
		p.Match(EvoLangParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInputDefContext is an interface to support dynamic dispatch.
type IInputDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INPUT_DEF() antlr.TerminalNode
	ID() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	InputBody() IInputBodyContext
	CLOSE_BRACKET() antlr.TerminalNode

	// IsInputDefContext differentiates from other interfaces.
	IsInputDefContext()
}

type InputDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputDefContext() *InputDefContext {
	var p = new(InputDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_inputDef
	return p
}

func InitEmptyInputDefContext(p *InputDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_inputDef
}

func (*InputDefContext) IsInputDefContext() {}

func NewInputDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputDefContext {
	var p = new(InputDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_inputDef

	return p
}

func (s *InputDefContext) GetParser() antlr.Parser { return s.parser }

func (s *InputDefContext) INPUT_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserINPUT_DEF, 0)
}

func (s *InputDefContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *InputDefContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_BRACKET, 0)
}

func (s *InputDefContext) InputBody() IInputBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputBodyContext)
}

func (s *InputDefContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_BRACKET, 0)
}

func (s *InputDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterInputDef(s)
	}
}

func (s *InputDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitInputDef(s)
	}
}

func (s *InputDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitInputDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) InputDef() (localctx IInputDefContext) {
	localctx = NewInputDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, EvoLangParserRULE_inputDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Match(EvoLangParserINPUT_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(270)
		p.InputBody()
	}
	{
		p.SetState(271)
		p.Match(EvoLangParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInputBodyContext is an interface to support dynamic dispatch.
type IInputBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInputFieldDef() []IInputFieldDefContext
	InputFieldDef(i int) IInputFieldDefContext

	// IsInputBodyContext differentiates from other interfaces.
	IsInputBodyContext()
}

type InputBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputBodyContext() *InputBodyContext {
	var p = new(InputBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_inputBody
	return p
}

func InitEmptyInputBodyContext(p *InputBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_inputBody
}

func (*InputBodyContext) IsInputBodyContext() {}

func NewInputBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputBodyContext {
	var p = new(InputBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_inputBody

	return p
}

func (s *InputBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *InputBodyContext) AllInputFieldDef() []IInputFieldDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInputFieldDefContext); ok {
			len++
		}
	}

	tst := make([]IInputFieldDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInputFieldDefContext); ok {
			tst[i] = t.(IInputFieldDefContext)
			i++
		}
	}

	return tst
}

func (s *InputBodyContext) InputFieldDef(i int) IInputFieldDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputFieldDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputFieldDefContext)
}

func (s *InputBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterInputBody(s)
	}
}

func (s *InputBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitInputBody(s)
	}
}

func (s *InputBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitInputBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) InputBody() (localctx IInputBodyContext) {
	localctx = NewInputBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, EvoLangParserRULE_inputBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserID {
		{
			p.SetState(273)
			p.InputFieldDef()
		}

		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInputFieldDefContext is an interface to support dynamic dispatch.
type IInputFieldDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	TypeDef() ITypeDefContext
	References() IReferencesContext

	// IsInputFieldDefContext differentiates from other interfaces.
	IsInputFieldDefContext()
}

type InputFieldDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputFieldDefContext() *InputFieldDefContext {
	var p = new(InputFieldDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_inputFieldDef
	return p
}

func InitEmptyInputFieldDefContext(p *InputFieldDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_inputFieldDef
}

func (*InputFieldDefContext) IsInputFieldDefContext() {}

func NewInputFieldDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputFieldDefContext {
	var p = new(InputFieldDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_inputFieldDef

	return p
}

func (s *InputFieldDefContext) GetParser() antlr.Parser { return s.parser }

func (s *InputFieldDefContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *InputFieldDefContext) COLON() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCOLON, 0)
}

func (s *InputFieldDefContext) TypeDef() ITypeDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDefContext)
}

func (s *InputFieldDefContext) References() IReferencesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferencesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferencesContext)
}

func (s *InputFieldDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputFieldDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputFieldDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterInputFieldDef(s)
	}
}

func (s *InputFieldDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitInputFieldDef(s)
	}
}

func (s *InputFieldDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitInputFieldDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) InputFieldDef() (localctx IInputFieldDefContext) {
	localctx = NewInputFieldDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, EvoLangParserRULE_inputFieldDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(280)
		p.Match(EvoLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(281)
		p.TypeDef()
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EvoLangParserREFERENCES_DEF {
		{
			p.SetState(282)
			p.References()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEventDefContext is an interface to support dynamic dispatch.
type IEventDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EVENTS_DEF() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	CLOSE_BRACKET() antlr.TerminalNode
	AllEventChannelHandler() []IEventChannelHandlerContext
	EventChannelHandler(i int) IEventChannelHandlerContext

	// IsEventDefContext differentiates from other interfaces.
	IsEventDefContext()
}

type EventDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEventDefContext() *EventDefContext {
	var p = new(EventDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_eventDef
	return p
}

func InitEmptyEventDefContext(p *EventDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_eventDef
}

func (*EventDefContext) IsEventDefContext() {}

func NewEventDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EventDefContext {
	var p = new(EventDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_eventDef

	return p
}

func (s *EventDefContext) GetParser() antlr.Parser { return s.parser }

func (s *EventDefContext) EVENTS_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserEVENTS_DEF, 0)
}

func (s *EventDefContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_BRACKET, 0)
}

func (s *EventDefContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_BRACKET, 0)
}

func (s *EventDefContext) AllEventChannelHandler() []IEventChannelHandlerContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEventChannelHandlerContext); ok {
			len++
		}
	}

	tst := make([]IEventChannelHandlerContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEventChannelHandlerContext); ok {
			tst[i] = t.(IEventChannelHandlerContext)
			i++
		}
	}

	return tst
}

func (s *EventDefContext) EventChannelHandler(i int) IEventChannelHandlerContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEventChannelHandlerContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEventChannelHandlerContext)
}

func (s *EventDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EventDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterEventDef(s)
	}
}

func (s *EventDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitEventDef(s)
	}
}

func (s *EventDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitEventDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) EventDef() (localctx IEventDefContext) {
	localctx = NewEventDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, EvoLangParserRULE_eventDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.Match(EvoLangParserEVENTS_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EvoLangParserHASH {
		{
			p.SetState(287)
			p.EventChannelHandler()
		}

		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(292)
		p.Match(EvoLangParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEventChannelHandlerContext is an interface to support dynamic dispatch.
type IEventChannelHandlerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HASH() antlr.TerminalNode
	DOT() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	COLON() antlr.TerminalNode
	OPEN_PAREN() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	CLOSE_PAREN() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	SANDBOX() antlr.TerminalNode
	CLOSE_BRACKET() antlr.TerminalNode

	// IsEventChannelHandlerContext differentiates from other interfaces.
	IsEventChannelHandlerContext()
}

type EventChannelHandlerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEventChannelHandlerContext() *EventChannelHandlerContext {
	var p = new(EventChannelHandlerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_eventChannelHandler
	return p
}

func InitEmptyEventChannelHandlerContext(p *EventChannelHandlerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_eventChannelHandler
}

func (*EventChannelHandlerContext) IsEventChannelHandlerContext() {}

func NewEventChannelHandlerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EventChannelHandlerContext {
	var p = new(EventChannelHandlerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_eventChannelHandler

	return p
}

func (s *EventChannelHandlerContext) GetParser() antlr.Parser { return s.parser }

func (s *EventChannelHandlerContext) HASH() antlr.TerminalNode {
	return s.GetToken(EvoLangParserHASH, 0)
}

func (s *EventChannelHandlerContext) DOT() antlr.TerminalNode {
	return s.GetToken(EvoLangParserDOT, 0)
}

func (s *EventChannelHandlerContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserID)
}

func (s *EventChannelHandlerContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, i)
}

func (s *EventChannelHandlerContext) COLON() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCOLON, 0)
}

func (s *EventChannelHandlerContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_PAREN, 0)
}

func (s *EventChannelHandlerContext) COMMA() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCOMMA, 0)
}

func (s *EventChannelHandlerContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_PAREN, 0)
}

func (s *EventChannelHandlerContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_BRACKET, 0)
}

func (s *EventChannelHandlerContext) SANDBOX() antlr.TerminalNode {
	return s.GetToken(EvoLangParserSANDBOX, 0)
}

func (s *EventChannelHandlerContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_BRACKET, 0)
}

func (s *EventChannelHandlerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventChannelHandlerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EventChannelHandlerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterEventChannelHandler(s)
	}
}

func (s *EventChannelHandlerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitEventChannelHandler(s)
	}
}

func (s *EventChannelHandlerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitEventChannelHandler(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) EventChannelHandler() (localctx IEventChannelHandlerContext) {
	localctx = NewEventChannelHandlerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, EvoLangParserRULE_eventChannelHandler)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)
		p.Match(EvoLangParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(295)
		p.Match(EvoLangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(296)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)
		p.Match(EvoLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(298)
		p.Match(EvoLangParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(299)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(300)
		p.Match(EvoLangParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(301)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(302)
		p.Match(EvoLangParserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(304)
		p.Match(EvoLangParserSANDBOX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(305)
		p.Match(EvoLangParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumDefContext is an interface to support dynamic dispatch.
type IEnumDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENUM_DEF() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	CLOSE_BRACKET() antlr.TerminalNode

	// IsEnumDefContext differentiates from other interfaces.
	IsEnumDefContext()
}

type EnumDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDefContext() *EnumDefContext {
	var p = new(EnumDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_enumDef
	return p
}

func InitEmptyEnumDefContext(p *EnumDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_enumDef
}

func (*EnumDefContext) IsEnumDefContext() {}

func NewEnumDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDefContext {
	var p = new(EnumDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_enumDef

	return p
}

func (s *EnumDefContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDefContext) ENUM_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserENUM_DEF, 0)
}

func (s *EnumDefContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserID)
}

func (s *EnumDefContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, i)
}

func (s *EnumDefContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_BRACKET, 0)
}

func (s *EnumDefContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_BRACKET, 0)
}

func (s *EnumDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterEnumDef(s)
	}
}

func (s *EnumDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitEnumDef(s)
	}
}

func (s *EnumDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitEnumDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) EnumDef() (localctx IEnumDefContext) {
	localctx = NewEnumDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, EvoLangParserRULE_enumDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Match(EvoLangParserENUM_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(308)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(309)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(310)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserID {
		{
			p.SetState(311)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(317)
		p.Match(EvoLangParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReferencesContext is an interface to support dynamic dispatch.
type IReferencesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REFERENCES_DEF() antlr.TerminalNode
	OPEN_PAREN() antlr.TerminalNode
	DOT() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	CLOSE_PAREN() antlr.TerminalNode

	// IsReferencesContext differentiates from other interfaces.
	IsReferencesContext()
}

type ReferencesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferencesContext() *ReferencesContext {
	var p = new(ReferencesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_references
	return p
}

func InitEmptyReferencesContext(p *ReferencesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_references
}

func (*ReferencesContext) IsReferencesContext() {}

func NewReferencesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferencesContext {
	var p = new(ReferencesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_references

	return p
}

func (s *ReferencesContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferencesContext) REFERENCES_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserREFERENCES_DEF, 0)
}

func (s *ReferencesContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_PAREN, 0)
}

func (s *ReferencesContext) DOT() antlr.TerminalNode {
	return s.GetToken(EvoLangParserDOT, 0)
}

func (s *ReferencesContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserID)
}

func (s *ReferencesContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, i)
}

func (s *ReferencesContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_PAREN, 0)
}

func (s *ReferencesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferencesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferencesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterReferences(s)
	}
}

func (s *ReferencesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitReferences(s)
	}
}

func (s *ReferencesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitReferences(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) References() (localctx IReferencesContext) {
	localctx = NewReferencesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, EvoLangParserRULE_references)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(319)
		p.Match(EvoLangParserREFERENCES_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(320)
		p.Match(EvoLangParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EvoLangParserID {
		{
			p.SetState(321)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(324)
		p.Match(EvoLangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(325)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(326)
		p.Match(EvoLangParserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILinkedByContext is an interface to support dynamic dispatch.
type ILinkedByContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LINKED_BY_DEF() antlr.TerminalNode
	OPEN_PAREN() antlr.TerminalNode
	DOT() antlr.TerminalNode
	ID() antlr.TerminalNode
	CLOSE_PAREN() antlr.TerminalNode

	// IsLinkedByContext differentiates from other interfaces.
	IsLinkedByContext()
}

type LinkedByContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLinkedByContext() *LinkedByContext {
	var p = new(LinkedByContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_linkedBy
	return p
}

func InitEmptyLinkedByContext(p *LinkedByContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_linkedBy
}

func (*LinkedByContext) IsLinkedByContext() {}

func NewLinkedByContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LinkedByContext {
	var p = new(LinkedByContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_linkedBy

	return p
}

func (s *LinkedByContext) GetParser() antlr.Parser { return s.parser }

func (s *LinkedByContext) LINKED_BY_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserLINKED_BY_DEF, 0)
}

func (s *LinkedByContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_PAREN, 0)
}

func (s *LinkedByContext) DOT() antlr.TerminalNode {
	return s.GetToken(EvoLangParserDOT, 0)
}

func (s *LinkedByContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *LinkedByContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_PAREN, 0)
}

func (s *LinkedByContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LinkedByContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LinkedByContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterLinkedBy(s)
	}
}

func (s *LinkedByContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitLinkedBy(s)
	}
}

func (s *LinkedByContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitLinkedBy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) LinkedBy() (localctx ILinkedByContext) {
	localctx = NewLinkedByContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, EvoLangParserRULE_linkedBy)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Match(EvoLangParserLINKED_BY_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(329)
		p.Match(EvoLangParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(330)
		p.Match(EvoLangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(331)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(332)
		p.Match(EvoLangParserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IActionContext is an interface to support dynamic dispatch.
type IActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsActionContext differentiates from other interfaces.
	IsActionContext()
}

type ActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionContext() *ActionContext {
	var p = new(ActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_action
	return p
}

func InitEmptyActionContext(p *ActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_action
}

func (*ActionContext) IsActionContext() {}

func NewActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionContext {
	var p = new(ActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_action

	return p
}

func (s *ActionContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *ActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterAction(s)
	}
}

func (s *ActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitAction(s)
	}
}

func (s *ActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) Action_() (localctx IActionContext) {
	localctx = NewActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, EvoLangParserRULE_action)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParamListContext differentiates from other interfaces.
	IsParamListContext()
}

type ParamListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListContext() *ParamListContext {
	var p = new(ParamListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_paramList
	return p
}

func InitEmptyParamListContext(p *ParamListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_paramList
}

func (*ParamListContext) IsParamListContext() {}

func NewParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListContext {
	var p = new(ParamListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_paramList

	return p
}

func (s *ParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserID)
}

func (s *ParamListContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, i)
}

func (s *ParamListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserCOMMA)
}

func (s *ParamListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserCOMMA, i)
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (s *ParamListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitParamList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) ParamList() (localctx IParamListContext) {
	localctx = NewParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, EvoLangParserRULE_paramList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(336)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserCOMMA {
		{
			p.SetState(337)
			p.Match(EvoLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBodyDefContext is an interface to support dynamic dispatch.
type IBodyDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BODY_DEF() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsBodyDefContext differentiates from other interfaces.
	IsBodyDefContext()
}

type BodyDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyDefContext() *BodyDefContext {
	var p = new(BodyDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_bodyDef
	return p
}

func InitEmptyBodyDefContext(p *BodyDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_bodyDef
}

func (*BodyDefContext) IsBodyDefContext() {}

func NewBodyDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyDefContext {
	var p = new(BodyDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_bodyDef

	return p
}

func (s *BodyDefContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyDefContext) BODY_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserBODY_DEF, 0)
}

func (s *BodyDefContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *BodyDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterBodyDef(s)
	}
}

func (s *BodyDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitBodyDef(s)
	}
}

func (s *BodyDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitBodyDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) BodyDef() (localctx IBodyDefContext) {
	localctx = NewBodyDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, EvoLangParserRULE_bodyDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.Match(EvoLangParserBODY_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(345)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAuthenticateDefContext is an interface to support dynamic dispatch.
type IAuthenticateDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AUTHENTICATE_DEF() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsAuthenticateDefContext differentiates from other interfaces.
	IsAuthenticateDefContext()
}

type AuthenticateDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuthenticateDefContext() *AuthenticateDefContext {
	var p = new(AuthenticateDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_authenticateDef
	return p
}

func InitEmptyAuthenticateDefContext(p *AuthenticateDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_authenticateDef
}

func (*AuthenticateDefContext) IsAuthenticateDefContext() {}

func NewAuthenticateDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AuthenticateDefContext {
	var p = new(AuthenticateDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_authenticateDef

	return p
}

func (s *AuthenticateDefContext) GetParser() antlr.Parser { return s.parser }

func (s *AuthenticateDefContext) AUTHENTICATE_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserAUTHENTICATE_DEF, 0)
}

func (s *AuthenticateDefContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *AuthenticateDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuthenticateDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AuthenticateDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterAuthenticateDef(s)
	}
}

func (s *AuthenticateDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitAuthenticateDef(s)
	}
}

func (s *AuthenticateDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitAuthenticateDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) AuthenticateDef() (localctx IAuthenticateDefContext) {
	localctx = NewAuthenticateDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, EvoLangParserRULE_authenticateDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.Match(EvoLangParserAUTHENTICATE_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(348)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnDefContext is an interface to support dynamic dispatch.
type IReturnDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN_DEF() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturnDefContext differentiates from other interfaces.
	IsReturnDefContext()
}

type ReturnDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnDefContext() *ReturnDefContext {
	var p = new(ReturnDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_returnDef
	return p
}

func InitEmptyReturnDefContext(p *ReturnDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_returnDef
}

func (*ReturnDefContext) IsReturnDefContext() {}

func NewReturnDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnDefContext {
	var p = new(ReturnDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_returnDef

	return p
}

func (s *ReturnDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnDefContext) RETURN_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserRETURN_DEF, 0)
}

func (s *ReturnDefContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterReturnDef(s)
	}
}

func (s *ReturnDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitReturnDef(s)
	}
}

func (s *ReturnDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitReturnDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) ReturnDef() (localctx IReturnDefContext) {
	localctx = NewReturnDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, EvoLangParserRULE_returnDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)
		p.Match(EvoLangParserRETURN_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(351)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeDefContext is an interface to support dynamic dispatch.
type ITypeDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE_INT_DEF() antlr.TerminalNode
	TYPE_STRING_DEF() antlr.TerminalNode
	TYPE_BOOLEAN_DEF() antlr.TerminalNode
	ArrayDef() IArrayDefContext
	ID() antlr.TerminalNode

	// IsTypeDefContext differentiates from other interfaces.
	IsTypeDefContext()
}

type TypeDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefContext() *TypeDefContext {
	var p = new(TypeDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_typeDef
	return p
}

func InitEmptyTypeDefContext(p *TypeDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_typeDef
}

func (*TypeDefContext) IsTypeDefContext() {}

func NewTypeDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefContext {
	var p = new(TypeDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_typeDef

	return p
}

func (s *TypeDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefContext) TYPE_INT_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserTYPE_INT_DEF, 0)
}

func (s *TypeDefContext) TYPE_STRING_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserTYPE_STRING_DEF, 0)
}

func (s *TypeDefContext) TYPE_BOOLEAN_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserTYPE_BOOLEAN_DEF, 0)
}

func (s *TypeDefContext) ArrayDef() IArrayDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayDefContext)
}

func (s *TypeDefContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *TypeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterTypeDef(s)
	}
}

func (s *TypeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitTypeDef(s)
	}
}

func (s *TypeDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitTypeDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) TypeDef() (localctx ITypeDefContext) {
	localctx = NewTypeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, EvoLangParserRULE_typeDef)
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EvoLangParserTYPE_INT_DEF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(353)
			p.Match(EvoLangParserTYPE_INT_DEF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvoLangParserTYPE_STRING_DEF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(354)
			p.Match(EvoLangParserTYPE_STRING_DEF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvoLangParserTYPE_BOOLEAN_DEF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(355)
			p.Match(EvoLangParserTYPE_BOOLEAN_DEF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvoLangParserOPEN_S_BRACKET:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(356)
			p.ArrayDef()
		}

	case EvoLangParserID:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(357)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayDefContext is an interface to support dynamic dispatch.
type IArrayDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_S_BRACKET() antlr.TerminalNode
	TypeDef() ITypeDefContext
	CLOSE_S_BRACKET() antlr.TerminalNode

	// IsArrayDefContext differentiates from other interfaces.
	IsArrayDefContext()
}

type ArrayDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayDefContext() *ArrayDefContext {
	var p = new(ArrayDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_arrayDef
	return p
}

func InitEmptyArrayDefContext(p *ArrayDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_arrayDef
}

func (*ArrayDefContext) IsArrayDefContext() {}

func NewArrayDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayDefContext {
	var p = new(ArrayDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_arrayDef

	return p
}

func (s *ArrayDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayDefContext) OPEN_S_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_S_BRACKET, 0)
}

func (s *ArrayDefContext) TypeDef() ITypeDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDefContext)
}

func (s *ArrayDefContext) CLOSE_S_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_S_BRACKET, 0)
}

func (s *ArrayDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterArrayDef(s)
	}
}

func (s *ArrayDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitArrayDef(s)
	}
}

func (s *ArrayDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitArrayDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) ArrayDef() (localctx IArrayDefContext) {
	localctx = NewArrayDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, EvoLangParserRULE_arrayDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)
		p.Match(EvoLangParserOPEN_S_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(361)
		p.TypeDef()
	}
	{
		p.SetState(362)
		p.Match(EvoLangParserCLOSE_S_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEventTypeContext is an interface to support dynamic dispatch.
type IEventTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsEventTypeContext differentiates from other interfaces.
	IsEventTypeContext()
}

type EventTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEventTypeContext() *EventTypeContext {
	var p = new(EventTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_eventType
	return p
}

func InitEmptyEventTypeContext(p *EventTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_eventType
}

func (*EventTypeContext) IsEventTypeContext() {}

func NewEventTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EventTypeContext {
	var p = new(EventTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_eventType

	return p
}

func (s *EventTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *EventTypeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserID)
}

func (s *EventTypeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, i)
}

func (s *EventTypeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserCOMMA)
}

func (s *EventTypeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserCOMMA, i)
}

func (s *EventTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EventTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterEventType(s)
	}
}

func (s *EventTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitEventType(s)
	}
}

func (s *EventTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitEventType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) EventType() (localctx IEventTypeContext) {
	localctx = NewEventTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, EvoLangParserRULE_eventType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserCOMMA {
		{
			p.SetState(365)
			p.Match(EvoLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(366)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionCallExprContext struct {
	ExpressionContext
}

func NewFunctionCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExprContext {
	var p = new(FunctionCallExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExprContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionCallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterFunctionCallExpr(s)
	}
}

func (s *FunctionCallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitFunctionCallExpr(s)
	}
}

func (s *FunctionCallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitFunctionCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivExprContext struct {
	ExpressionContext
	left  IExpressionContext
	op    antlr.Token
	right IExpressionContext
}

func NewMulDivExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivExprContext {
	var p = new(MulDivExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MulDivExprContext) GetOp() antlr.Token { return s.op }

func (s *MulDivExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivExprContext) GetLeft() IExpressionContext { return s.left }

func (s *MulDivExprContext) GetRight() IExpressionContext { return s.right }

func (s *MulDivExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *MulDivExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *MulDivExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MulDivExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulDivExprContext) OP_MUL() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_MUL, 0)
}

func (s *MulDivExprContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_DIV, 0)
}

func (s *MulDivExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterMulDivExpr(s)
	}
}

func (s *MulDivExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitMulDivExpr(s)
	}
}

func (s *MulDivExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitMulDivExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExprContext struct {
	ExpressionContext
	id antlr.Token
}

func NewIdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExprContext {
	var p = new(IdExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IdExprContext) GetId() antlr.Token { return s.id }

func (s *IdExprContext) SetId(v antlr.Token) { s.id = v }

func (s *IdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExprContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *IdExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterIdExpr(s)
	}
}

func (s *IdExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitIdExpr(s)
	}
}

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringLiteralContext struct {
	ExpressionContext
	strVal antlr.Token
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *StringLiteralContext) GetStrVal() antlr.Token { return s.strVal }

func (s *StringLiteralContext) SetStrVal(v antlr.Token) { s.strVal = v }

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(EvoLangParserSTRING, 0)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanLiteralContext struct {
	ExpressionContext
	boolVal antlr.Token
}

func NewBooleanLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BooleanLiteralContext) GetBoolVal() antlr.Token { return s.boolVal }

func (s *BooleanLiteralContext) SetBoolVal(v antlr.Token) { s.boolVal = v }

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserBOOLEAN, 0)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntLiteralContext struct {
	ExpressionContext
	intVal antlr.Token
}

func NewIntLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLiteralContext {
	var p = new(IntLiteralContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IntLiteralContext) GetIntVal() antlr.Token { return s.intVal }

func (s *IntLiteralContext) SetIntVal(v antlr.Token) { s.intVal = v }

func (s *IntLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(EvoLangParserINT, 0)
}

func (s *IntLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterIntLiteral(s)
	}
}

func (s *IntLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitIntLiteral(s)
	}
}

func (s *IntLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitIntLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type RelationalExprContext struct {
	ExpressionContext
	left  IExpressionContext
	op    antlr.Token
	right IExpressionContext
}

func NewRelationalExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalExprContext {
	var p = new(RelationalExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RelationalExprContext) GetOp() antlr.Token { return s.op }

func (s *RelationalExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationalExprContext) GetLeft() IExpressionContext { return s.left }

func (s *RelationalExprContext) GetRight() IExpressionContext { return s.right }

func (s *RelationalExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *RelationalExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *RelationalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RelationalExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RelationalExprContext) OP_GT() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_GT, 0)
}

func (s *RelationalExprContext) OP_LT() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_LT, 0)
}

func (s *RelationalExprContext) OP_EQEQ() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_EQEQ, 0)
}

func (s *RelationalExprContext) OP_NEQ() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_NEQ, 0)
}

func (s *RelationalExprContext) OP_GTEQ() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_GTEQ, 0)
}

func (s *RelationalExprContext) OP_LTEQ() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_LTEQ, 0)
}

func (s *RelationalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterRelationalExpr(s)
	}
}

func (s *RelationalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitRelationalExpr(s)
	}
}

func (s *RelationalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitRelationalExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExprContext struct {
	ExpressionContext
	inner IExpressionContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParenExprContext) GetInner() IExpressionContext { return s.inner }

func (s *ParenExprContext) SetInner(v IExpressionContext) { s.inner = v }

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_PAREN, 0)
}

func (s *ParenExprContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_PAREN, 0)
}

func (s *ParenExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PropertyAccessExprContext struct {
	ExpressionContext
}

func NewPropertyAccessExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertyAccessExprContext {
	var p = new(PropertyAccessExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PropertyAccessExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyAccessExprContext) PropertyAccess() IPropertyAccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyAccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyAccessContext)
}

func (s *PropertyAccessExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterPropertyAccessExpr(s)
	}
}

func (s *PropertyAccessExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitPropertyAccessExpr(s)
	}
}

func (s *PropertyAccessExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitPropertyAccessExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubExprContext struct {
	ExpressionContext
	left  IExpressionContext
	op    antlr.Token
	right IExpressionContext
}

func NewAddSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubExprContext {
	var p = new(AddSubExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddSubExprContext) GetOp() antlr.Token { return s.op }

func (s *AddSubExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubExprContext) GetLeft() IExpressionContext { return s.left }

func (s *AddSubExprContext) GetRight() IExpressionContext { return s.right }

func (s *AddSubExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *AddSubExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *AddSubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddSubExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddSubExprContext) OP_SUM() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_SUM, 0)
}

func (s *AddSubExprContext) OP_SUB() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_SUB, 0)
}

func (s *AddSubExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterAddSubExpr(s)
	}
}

func (s *AddSubExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitAddSubExpr(s)
	}
}

func (s *AddSubExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitAddSubExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalExprContext struct {
	ExpressionContext
	left  IExpressionContext
	op    antlr.Token
	right IExpressionContext
}

func NewLogicalExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalExprContext {
	var p = new(LogicalExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalExprContext) GetOp() antlr.Token { return s.op }

func (s *LogicalExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicalExprContext) GetLeft() IExpressionContext { return s.left }

func (s *LogicalExprContext) GetRight() IExpressionContext { return s.right }

func (s *LogicalExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *LogicalExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *LogicalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicalExprContext) OP_AND() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_AND, 0)
}

func (s *LogicalExprContext) OP_OR() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_OR, 0)
}

func (s *LogicalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterLogicalExpr(s)
	}
}

func (s *LogicalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitLogicalExpr(s)
	}
}

func (s *LogicalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitLogicalExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MethodCallExprContext struct {
	ExpressionContext
}

func NewMethodCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MethodCallExprContext {
	var p = new(MethodCallExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MethodCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallExprContext) MethodCallExpression() IMethodCallExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodCallExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodCallExpressionContext)
}

func (s *MethodCallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterMethodCallExpr(s)
	}
}

func (s *MethodCallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitMethodCallExpr(s)
	}
}

func (s *MethodCallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitMethodCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *EvoLangParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 82
	p.EnterRecursionRule(localctx, 82, EvoLangParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMethodCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(373)
			p.MethodCallExpression()
		}

	case 2:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(374)
			p.Match(EvoLangParserOPEN_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)

			var _x = p.expression(0)

			localctx.(*ParenExprContext).inner = _x
		}
		{
			p.SetState(376)
			p.Match(EvoLangParserCLOSE_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewPropertyAccessExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(378)
			p.PropertyAccess()
		}

	case 4:
		localctx = NewFunctionCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(379)
			p.FunctionCall()
		}

	case 5:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(380)

			var _m = p.Match(EvoLangParserID)

			localctx.(*IdExprContext).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewIntLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(381)

			var _m = p.Match(EvoLangParserINT)

			localctx.(*IntLiteralContext).intVal = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewStringLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(382)

			var _m = p.Match(EvoLangParserSTRING)

			localctx.(*StringLiteralContext).strVal = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewBooleanLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(383)

			var _m = p.Match(EvoLangParserBOOLEAN)

			localctx.(*BooleanLiteralContext).boolVal = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(398)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*MulDivExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvoLangParserRULE_expression)
				p.SetState(386)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(387)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == EvoLangParserOP_MUL || _la == EvoLangParserOP_DIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(388)

					var _x = p.expression(12)

					localctx.(*MulDivExprContext).right = _x
				}

			case 2:
				localctx = NewAddSubExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*AddSubExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvoLangParserRULE_expression)
				p.SetState(389)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(390)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == EvoLangParserOP_SUM || _la == EvoLangParserOP_SUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(391)

					var _x = p.expression(11)

					localctx.(*AddSubExprContext).right = _x
				}

			case 3:
				localctx = NewRelationalExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*RelationalExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvoLangParserRULE_expression)
				p.SetState(392)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(393)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8455716864) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(394)

					var _x = p.expression(10)

					localctx.(*RelationalExprContext).right = _x
				}

			case 4:
				localctx = NewLogicalExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*LogicalExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvoLangParserRULE_expression)
				p.SetState(395)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(396)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*LogicalExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == EvoLangParserOP_AND || _la == EvoLangParserOP_OR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*LogicalExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(397)

					var _x = p.expression(9)

					localctx.(*LogicalExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(402)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethodCallExpressionContext is an interface to support dynamic dispatch.
type IMethodCallExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetArgs returns the args rule contexts.
	GetArgs() IExprListContext

	// SetArgs sets the args rule contexts.
	SetArgs(IExprListContext)

	// Getter signatures
	PrimaryExpression() IPrimaryExpressionContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllOPEN_PAREN() []antlr.TerminalNode
	OPEN_PAREN(i int) antlr.TerminalNode
	AllCLOSE_PAREN() []antlr.TerminalNode
	CLOSE_PAREN(i int) antlr.TerminalNode
	AllExprList() []IExprListContext
	ExprList(i int) IExprListContext

	// IsMethodCallExpressionContext differentiates from other interfaces.
	IsMethodCallExpressionContext()
}

type MethodCallExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	args   IExprListContext
}

func NewEmptyMethodCallExpressionContext() *MethodCallExpressionContext {
	var p = new(MethodCallExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_methodCallExpression
	return p
}

func InitEmptyMethodCallExpressionContext(p *MethodCallExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_methodCallExpression
}

func (*MethodCallExpressionContext) IsMethodCallExpressionContext() {}

func NewMethodCallExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallExpressionContext {
	var p = new(MethodCallExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_methodCallExpression

	return p
}

func (s *MethodCallExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallExpressionContext) GetArgs() IExprListContext { return s.args }

func (s *MethodCallExpressionContext) SetArgs(v IExprListContext) { s.args = v }

func (s *MethodCallExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *MethodCallExpressionContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserDOT)
}

func (s *MethodCallExpressionContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserDOT, i)
}

func (s *MethodCallExpressionContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserID)
}

func (s *MethodCallExpressionContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, i)
}

func (s *MethodCallExpressionContext) AllOPEN_PAREN() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserOPEN_PAREN)
}

func (s *MethodCallExpressionContext) OPEN_PAREN(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_PAREN, i)
}

func (s *MethodCallExpressionContext) AllCLOSE_PAREN() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserCLOSE_PAREN)
}

func (s *MethodCallExpressionContext) CLOSE_PAREN(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_PAREN, i)
}

func (s *MethodCallExpressionContext) AllExprList() []IExprListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprListContext); ok {
			len++
		}
	}

	tst := make([]IExprListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprListContext); ok {
			tst[i] = t.(IExprListContext)
			i++
		}
	}

	return tst
}

func (s *MethodCallExpressionContext) ExprList(i int) IExprListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *MethodCallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterMethodCallExpression(s)
	}
}

func (s *MethodCallExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitMethodCallExpression(s)
	}
}

func (s *MethodCallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitMethodCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) MethodCallExpression() (localctx IMethodCallExpressionContext) {
	localctx = NewMethodCallExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, EvoLangParserRULE_methodCallExpression)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(403)
		p.PrimaryExpression()
	}
	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(404)
				p.Match(EvoLangParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(405)
				p.Match(EvoLangParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(406)
				p.Match(EvoLangParserOPEN_PAREN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(408)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1134146244050944) != 0 {
				{
					p.SetState(407)

					var _x = p.ExprList()

					localctx.(*MethodCallExpressionContext).args = _x
				}

			}
			{
				p.SetState(410)
				p.Match(EvoLangParserCLOSE_PAREN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(413)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) CopyAll(ctx *PrimaryExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrimaryFunctionCallExprContext struct {
	PrimaryExpressionContext
	args IExprListContext
}

func NewPrimaryFunctionCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryFunctionCallExprContext {
	var p = new(PrimaryFunctionCallExprContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryFunctionCallExprContext) GetArgs() IExprListContext { return s.args }

func (s *PrimaryFunctionCallExprContext) SetArgs(v IExprListContext) { s.args = v }

func (s *PrimaryFunctionCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryFunctionCallExprContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *PrimaryFunctionCallExprContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_PAREN, 0)
}

func (s *PrimaryFunctionCallExprContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_PAREN, 0)
}

func (s *PrimaryFunctionCallExprContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *PrimaryFunctionCallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterPrimaryFunctionCallExpr(s)
	}
}

func (s *PrimaryFunctionCallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitPrimaryFunctionCallExpr(s)
	}
}

func (s *PrimaryFunctionCallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitPrimaryFunctionCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryPropertyAccessExprContext struct {
	PrimaryExpressionContext
}

func NewPrimaryPropertyAccessExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryPropertyAccessExprContext {
	var p = new(PrimaryPropertyAccessExprContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryPropertyAccessExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryPropertyAccessExprContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserID)
}

func (s *PrimaryPropertyAccessExprContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, i)
}

func (s *PrimaryPropertyAccessExprContext) DOT() antlr.TerminalNode {
	return s.GetToken(EvoLangParserDOT, 0)
}

func (s *PrimaryPropertyAccessExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterPrimaryPropertyAccessExpr(s)
	}
}

func (s *PrimaryPropertyAccessExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitPrimaryPropertyAccessExpr(s)
	}
}

func (s *PrimaryPropertyAccessExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitPrimaryPropertyAccessExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SimpleIdExprContext struct {
	PrimaryExpressionContext
}

func NewSimpleIdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleIdExprContext {
	var p = new(SimpleIdExprContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *SimpleIdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleIdExprContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *SimpleIdExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterSimpleIdExpr(s)
	}
}

func (s *SimpleIdExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitSimpleIdExpr(s)
	}
}

func (s *SimpleIdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitSimpleIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, EvoLangParserRULE_primaryExpression)
	var _la int

	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSimpleIdExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(415)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewPrimaryFunctionCallExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(416)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.Match(EvoLangParserOPEN_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(419)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1134146244050944) != 0 {
			{
				p.SetState(418)

				var _x = p.ExprList()

				localctx.(*PrimaryFunctionCallExprContext).args = _x
			}

		}
		{
			p.SetState(421)
			p.Match(EvoLangParserCLOSE_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewPrimaryPropertyAccessExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(422)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)
			p.Match(EvoLangParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(424)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPropertyAccessContext is an interface to support dynamic dispatch.
type IPropertyAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsPropertyAccessContext differentiates from other interfaces.
	IsPropertyAccessContext()
}

type PropertyAccessContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyAccessContext() *PropertyAccessContext {
	var p = new(PropertyAccessContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_propertyAccess
	return p
}

func InitEmptyPropertyAccessContext(p *PropertyAccessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_propertyAccess
}

func (*PropertyAccessContext) IsPropertyAccessContext() {}

func NewPropertyAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyAccessContext {
	var p = new(PropertyAccessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_propertyAccess

	return p
}

func (s *PropertyAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyAccessContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserID)
}

func (s *PropertyAccessContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, i)
}

func (s *PropertyAccessContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserDOT)
}

func (s *PropertyAccessContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserDOT, i)
}

func (s *PropertyAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterPropertyAccess(s)
	}
}

func (s *PropertyAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitPropertyAccess(s)
	}
}

func (s *PropertyAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitPropertyAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) PropertyAccess() (localctx IPropertyAccessContext) {
	localctx = NewPropertyAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, EvoLangParserRULE_propertyAccess)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(427)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(428)
				p.Match(EvoLangParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(429)
				p.Match(EvoLangParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetArgs returns the args rule contexts.
	GetArgs() IExprListContext

	// SetArgs sets the args rule contexts.
	SetArgs(IExprListContext)

	// Getter signatures
	ID() antlr.TerminalNode
	OPEN_PAREN() antlr.TerminalNode
	CLOSE_PAREN() antlr.TerminalNode
	ExprList() IExprListContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	args   IExprListContext
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) GetArgs() IExprListContext { return s.args }

func (s *FunctionCallContext) SetArgs(v IExprListContext) { s.args = v }

func (s *FunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *FunctionCallContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_PAREN, 0)
}

func (s *FunctionCallContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_PAREN, 0)
}

func (s *FunctionCallContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, EvoLangParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(434)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(435)
		p.Match(EvoLangParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1134146244050944) != 0 {
		{
			p.SetState(436)

			var _x = p.ExprList()

			localctx.(*FunctionCallContext).args = _x
		}

	}
	{
		p.SetState(439)
		p.Match(EvoLangParserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprListContext is an interface to support dynamic dispatch.
type IExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExprListContext differentiates from other interfaces.
	IsExprListContext()
}

type ExprListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprListContext() *ExprListContext {
	var p = new(ExprListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_exprList
	return p
}

func InitEmptyExprListContext(p *ExprListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_exprList
}

func (*ExprListContext) IsExprListContext() {}

func NewExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprListContext {
	var p = new(ExprListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_exprList

	return p
}

func (s *ExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExprListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExprListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserCOMMA)
}

func (s *ExprListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserCOMMA, i)
}

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterExprList(s)
	}
}

func (s *ExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitExprList(s)
	}
}

func (s *ExprListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitExprList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) ExprList() (localctx IExprListContext) {
	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, EvoLangParserRULE_exprList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(441)
		p.expression(0)
	}
	p.SetState(446)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserCOMMA {
		{
			p.SetState(442)
			p.Match(EvoLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(443)
			p.expression(0)
		}

		p.SetState(448)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMainFunctionContext is an interface to support dynamic dispatch.
type IMainFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAIN_DEF() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	MainStatements() IMainStatementsContext
	CLOSE_BRACKET() antlr.TerminalNode

	// IsMainFunctionContext differentiates from other interfaces.
	IsMainFunctionContext()
}

type MainFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainFunctionContext() *MainFunctionContext {
	var p = new(MainFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mainFunction
	return p
}

func InitEmptyMainFunctionContext(p *MainFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mainFunction
}

func (*MainFunctionContext) IsMainFunctionContext() {}

func NewMainFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainFunctionContext {
	var p = new(MainFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_mainFunction

	return p
}

func (s *MainFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *MainFunctionContext) MAIN_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserMAIN_DEF, 0)
}

func (s *MainFunctionContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_BRACKET, 0)
}

func (s *MainFunctionContext) MainStatements() IMainStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMainStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMainStatementsContext)
}

func (s *MainFunctionContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_BRACKET, 0)
}

func (s *MainFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterMainFunction(s)
	}
}

func (s *MainFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitMainFunction(s)
	}
}

func (s *MainFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitMainFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) MainFunction() (localctx IMainFunctionContext) {
	localctx = NewMainFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, EvoLangParserRULE_mainFunction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(449)
		p.Match(EvoLangParserMAIN_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(450)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(451)
		p.MainStatements()
	}
	{
		p.SetState(452)
		p.Match(EvoLangParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMainStatementsContext is an interface to support dynamic dispatch.
type IMainStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMainStatement() []IMainStatementContext
	MainStatement(i int) IMainStatementContext

	// IsMainStatementsContext differentiates from other interfaces.
	IsMainStatementsContext()
}

type MainStatementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainStatementsContext() *MainStatementsContext {
	var p = new(MainStatementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mainStatements
	return p
}

func InitEmptyMainStatementsContext(p *MainStatementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mainStatements
}

func (*MainStatementsContext) IsMainStatementsContext() {}

func NewMainStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainStatementsContext {
	var p = new(MainStatementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_mainStatements

	return p
}

func (s *MainStatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *MainStatementsContext) AllMainStatement() []IMainStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMainStatementContext); ok {
			len++
		}
	}

	tst := make([]IMainStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMainStatementContext); ok {
			tst[i] = t.(IMainStatementContext)
			i++
		}
	}

	return tst
}

func (s *MainStatementsContext) MainStatement(i int) IMainStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMainStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMainStatementContext)
}

func (s *MainStatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainStatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainStatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterMainStatements(s)
	}
}

func (s *MainStatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitMainStatements(s)
	}
}

func (s *MainStatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitMainStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) MainStatements() (localctx IMainStatementsContext) {
	localctx = NewMainStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, EvoLangParserRULE_mainStatements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(455)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EvoLangParserPRINT_DEF || _la == EvoLangParserID {
		{
			p.SetState(454)
			p.MainStatement()
		}

		p.SetState(457)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMainStatementContext is an interface to support dynamic dispatch.
type IMainStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrintStatement() IPrintStatementContext
	ServerStatement() IServerStatementContext

	// IsMainStatementContext differentiates from other interfaces.
	IsMainStatementContext()
}

type MainStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainStatementContext() *MainStatementContext {
	var p = new(MainStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mainStatement
	return p
}

func InitEmptyMainStatementContext(p *MainStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mainStatement
}

func (*MainStatementContext) IsMainStatementContext() {}

func NewMainStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainStatementContext {
	var p = new(MainStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_mainStatement

	return p
}

func (s *MainStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *MainStatementContext) PrintStatement() IPrintStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
}

func (s *MainStatementContext) ServerStatement() IServerStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServerStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServerStatementContext)
}

func (s *MainStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterMainStatement(s)
	}
}

func (s *MainStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitMainStatement(s)
	}
}

func (s *MainStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitMainStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) MainStatement() (localctx IMainStatementContext) {
	localctx = NewMainStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, EvoLangParserRULE_mainStatement)
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EvoLangParserPRINT_DEF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(459)
			p.PrintStatement()
		}

	case EvoLangParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(460)
			p.ServerStatement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintStatementContext is an interface to support dynamic dispatch.
type IPrintStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT_DEF() antlr.TerminalNode
	OPEN_PAREN() antlr.TerminalNode
	STRING() antlr.TerminalNode
	CLOSE_PAREN() antlr.TerminalNode

	// IsPrintStatementContext differentiates from other interfaces.
	IsPrintStatementContext()
}

type PrintStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStatementContext() *PrintStatementContext {
	var p = new(PrintStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_printStatement
	return p
}

func InitEmptyPrintStatementContext(p *PrintStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_printStatement
}

func (*PrintStatementContext) IsPrintStatementContext() {}

func NewPrintStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStatementContext {
	var p = new(PrintStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_printStatement

	return p
}

func (s *PrintStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStatementContext) PRINT_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserPRINT_DEF, 0)
}

func (s *PrintStatementContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_PAREN, 0)
}

func (s *PrintStatementContext) STRING() antlr.TerminalNode {
	return s.GetToken(EvoLangParserSTRING, 0)
}

func (s *PrintStatementContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_PAREN, 0)
}

func (s *PrintStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterPrintStatement(s)
	}
}

func (s *PrintStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitPrintStatement(s)
	}
}

func (s *PrintStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitPrintStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) PrintStatement() (localctx IPrintStatementContext) {
	localctx = NewPrintStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, EvoLangParserRULE_printStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(463)
		p.Match(EvoLangParserPRINT_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(464)
		p.Match(EvoLangParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(465)
		p.Match(EvoLangParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(466)
		p.Match(EvoLangParserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IServerStatementContext is an interface to support dynamic dispatch.
type IServerStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	DOT() antlr.TerminalNode
	LISTEN_DEF() antlr.TerminalNode
	OPEN_PAREN() antlr.TerminalNode
	INT() antlr.TerminalNode
	CLOSE_PAREN() antlr.TerminalNode

	// IsServerStatementContext differentiates from other interfaces.
	IsServerStatementContext()
}

type ServerStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServerStatementContext() *ServerStatementContext {
	var p = new(ServerStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_serverStatement
	return p
}

func InitEmptyServerStatementContext(p *ServerStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_serverStatement
}

func (*ServerStatementContext) IsServerStatementContext() {}

func NewServerStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServerStatementContext {
	var p = new(ServerStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_serverStatement

	return p
}

func (s *ServerStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ServerStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *ServerStatementContext) DOT() antlr.TerminalNode {
	return s.GetToken(EvoLangParserDOT, 0)
}

func (s *ServerStatementContext) LISTEN_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserLISTEN_DEF, 0)
}

func (s *ServerStatementContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_PAREN, 0)
}

func (s *ServerStatementContext) INT() antlr.TerminalNode {
	return s.GetToken(EvoLangParserINT, 0)
}

func (s *ServerStatementContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_PAREN, 0)
}

func (s *ServerStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServerStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServerStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterServerStatement(s)
	}
}

func (s *ServerStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitServerStatement(s)
	}
}

func (s *ServerStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitServerStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) ServerStatement() (localctx IServerStatementContext) {
	localctx = NewServerStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, EvoLangParserRULE_serverStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(468)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(469)
		p.Match(EvoLangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(470)
		p.Match(EvoLangParserLISTEN_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(471)
		p.Match(EvoLangParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(472)
		p.Match(EvoLangParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(473)
		p.Match(EvoLangParserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISandboxBlockContext is an interface to support dynamic dispatch.
type ISandboxBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_SANDBOX_MODE() antlr.TerminalNode
	SandboxAttributes() ISandboxAttributesContext
	SANDBOX_CONTENT() antlr.TerminalNode
	CLOSE_SANDBOX_MODE() antlr.TerminalNode

	// IsSandboxBlockContext differentiates from other interfaces.
	IsSandboxBlockContext()
}

type SandboxBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySandboxBlockContext() *SandboxBlockContext {
	var p = new(SandboxBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_sandboxBlock
	return p
}

func InitEmptySandboxBlockContext(p *SandboxBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_sandboxBlock
}

func (*SandboxBlockContext) IsSandboxBlockContext() {}

func NewSandboxBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SandboxBlockContext {
	var p = new(SandboxBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_sandboxBlock

	return p
}

func (s *SandboxBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *SandboxBlockContext) OPEN_SANDBOX_MODE() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_SANDBOX_MODE, 0)
}

func (s *SandboxBlockContext) SandboxAttributes() ISandboxAttributesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISandboxAttributesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISandboxAttributesContext)
}

func (s *SandboxBlockContext) SANDBOX_CONTENT() antlr.TerminalNode {
	return s.GetToken(EvoLangParserSANDBOX_CONTENT, 0)
}

func (s *SandboxBlockContext) CLOSE_SANDBOX_MODE() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_SANDBOX_MODE, 0)
}

func (s *SandboxBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SandboxBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SandboxBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterSandboxBlock(s)
	}
}

func (s *SandboxBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitSandboxBlock(s)
	}
}

func (s *SandboxBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitSandboxBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) SandboxBlock() (localctx ISandboxBlockContext) {
	localctx = NewSandboxBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, EvoLangParserRULE_sandboxBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(475)
		p.Match(EvoLangParserOPEN_SANDBOX_MODE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(476)
		p.SandboxAttributes()
	}
	{
		p.SetState(477)
		p.Match(EvoLangParserSANDBOX_CONTENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(478)
		p.Match(EvoLangParserCLOSE_SANDBOX_MODE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISandboxAttributesContext is an interface to support dynamic dispatch.
type ISandboxAttributesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSandboxAttribute() []ISandboxAttributeContext
	SandboxAttribute(i int) ISandboxAttributeContext

	// IsSandboxAttributesContext differentiates from other interfaces.
	IsSandboxAttributesContext()
}

type SandboxAttributesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySandboxAttributesContext() *SandboxAttributesContext {
	var p = new(SandboxAttributesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_sandboxAttributes
	return p
}

func InitEmptySandboxAttributesContext(p *SandboxAttributesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_sandboxAttributes
}

func (*SandboxAttributesContext) IsSandboxAttributesContext() {}

func NewSandboxAttributesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SandboxAttributesContext {
	var p = new(SandboxAttributesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_sandboxAttributes

	return p
}

func (s *SandboxAttributesContext) GetParser() antlr.Parser { return s.parser }

func (s *SandboxAttributesContext) AllSandboxAttribute() []ISandboxAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISandboxAttributeContext); ok {
			len++
		}
	}

	tst := make([]ISandboxAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISandboxAttributeContext); ok {
			tst[i] = t.(ISandboxAttributeContext)
			i++
		}
	}

	return tst
}

func (s *SandboxAttributesContext) SandboxAttribute(i int) ISandboxAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISandboxAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISandboxAttributeContext)
}

func (s *SandboxAttributesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SandboxAttributesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SandboxAttributesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterSandboxAttributes(s)
	}
}

func (s *SandboxAttributesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitSandboxAttributes(s)
	}
}

func (s *SandboxAttributesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitSandboxAttributes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) SandboxAttributes() (localctx ISandboxAttributesContext) {
	localctx = NewSandboxAttributesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, EvoLangParserRULE_sandboxAttributes)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(483)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserDEFINE_DEF || _la == EvoLangParserLANG_DEF {
		{
			p.SetState(480)
			p.SandboxAttribute()
		}

		p.SetState(485)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISandboxAttributeContext is an interface to support dynamic dispatch.
type ISandboxAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SandboxLangAttribute() ISandboxLangAttributeContext
	SandboxDefineAttribute() ISandboxDefineAttributeContext

	// IsSandboxAttributeContext differentiates from other interfaces.
	IsSandboxAttributeContext()
}

type SandboxAttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySandboxAttributeContext() *SandboxAttributeContext {
	var p = new(SandboxAttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_sandboxAttribute
	return p
}

func InitEmptySandboxAttributeContext(p *SandboxAttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_sandboxAttribute
}

func (*SandboxAttributeContext) IsSandboxAttributeContext() {}

func NewSandboxAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SandboxAttributeContext {
	var p = new(SandboxAttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_sandboxAttribute

	return p
}

func (s *SandboxAttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *SandboxAttributeContext) SandboxLangAttribute() ISandboxLangAttributeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISandboxLangAttributeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISandboxLangAttributeContext)
}

func (s *SandboxAttributeContext) SandboxDefineAttribute() ISandboxDefineAttributeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISandboxDefineAttributeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISandboxDefineAttributeContext)
}

func (s *SandboxAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SandboxAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SandboxAttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterSandboxAttribute(s)
	}
}

func (s *SandboxAttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitSandboxAttribute(s)
	}
}

func (s *SandboxAttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitSandboxAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) SandboxAttribute() (localctx ISandboxAttributeContext) {
	localctx = NewSandboxAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, EvoLangParserRULE_sandboxAttribute)
	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EvoLangParserLANG_DEF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(486)
			p.SandboxLangAttribute()
		}

	case EvoLangParserDEFINE_DEF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(487)
			p.SandboxDefineAttribute()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISandboxLangAttributeContext is an interface to support dynamic dispatch.
type ISandboxLangAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LANG_DEF() antlr.TerminalNode
	OP_ASSIGN() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsSandboxLangAttributeContext differentiates from other interfaces.
	IsSandboxLangAttributeContext()
}

type SandboxLangAttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySandboxLangAttributeContext() *SandboxLangAttributeContext {
	var p = new(SandboxLangAttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_sandboxLangAttribute
	return p
}

func InitEmptySandboxLangAttributeContext(p *SandboxLangAttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_sandboxLangAttribute
}

func (*SandboxLangAttributeContext) IsSandboxLangAttributeContext() {}

func NewSandboxLangAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SandboxLangAttributeContext {
	var p = new(SandboxLangAttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_sandboxLangAttribute

	return p
}

func (s *SandboxLangAttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *SandboxLangAttributeContext) LANG_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserLANG_DEF, 0)
}

func (s *SandboxLangAttributeContext) OP_ASSIGN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_ASSIGN, 0)
}

func (s *SandboxLangAttributeContext) STRING() antlr.TerminalNode {
	return s.GetToken(EvoLangParserSTRING, 0)
}

func (s *SandboxLangAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SandboxLangAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SandboxLangAttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterSandboxLangAttribute(s)
	}
}

func (s *SandboxLangAttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitSandboxLangAttribute(s)
	}
}

func (s *SandboxLangAttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitSandboxLangAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) SandboxLangAttribute() (localctx ISandboxLangAttributeContext) {
	localctx = NewSandboxLangAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, EvoLangParserRULE_sandboxLangAttribute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(490)
		p.Match(EvoLangParserLANG_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(491)
		p.Match(EvoLangParserOP_ASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(492)
		p.Match(EvoLangParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISandboxDefineAttributeContext is an interface to support dynamic dispatch.
type ISandboxDefineAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFINE_DEF() antlr.TerminalNode
	OP_ASSIGN() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	CLOSE_BRACKET() antlr.TerminalNode
	SandboxDefineContent() ISandboxDefineContentContext

	// IsSandboxDefineAttributeContext differentiates from other interfaces.
	IsSandboxDefineAttributeContext()
}

type SandboxDefineAttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySandboxDefineAttributeContext() *SandboxDefineAttributeContext {
	var p = new(SandboxDefineAttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_sandboxDefineAttribute
	return p
}

func InitEmptySandboxDefineAttributeContext(p *SandboxDefineAttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_sandboxDefineAttribute
}

func (*SandboxDefineAttributeContext) IsSandboxDefineAttributeContext() {}

func NewSandboxDefineAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SandboxDefineAttributeContext {
	var p = new(SandboxDefineAttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_sandboxDefineAttribute

	return p
}

func (s *SandboxDefineAttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *SandboxDefineAttributeContext) DEFINE_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserDEFINE_DEF, 0)
}

func (s *SandboxDefineAttributeContext) OP_ASSIGN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_ASSIGN, 0)
}

func (s *SandboxDefineAttributeContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_BRACKET, 0)
}

func (s *SandboxDefineAttributeContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_BRACKET, 0)
}

func (s *SandboxDefineAttributeContext) SandboxDefineContent() ISandboxDefineContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISandboxDefineContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISandboxDefineContentContext)
}

func (s *SandboxDefineAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SandboxDefineAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SandboxDefineAttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterSandboxDefineAttribute(s)
	}
}

func (s *SandboxDefineAttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitSandboxDefineAttribute(s)
	}
}

func (s *SandboxDefineAttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitSandboxDefineAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) SandboxDefineAttribute() (localctx ISandboxDefineAttributeContext) {
	localctx = NewSandboxDefineAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, EvoLangParserRULE_sandboxDefineAttribute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(494)
		p.Match(EvoLangParserDEFINE_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(495)
		p.Match(EvoLangParserOP_ASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(496)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(497)
		p.SandboxDefineContent()
	}

	{
		p.SetState(498)
		p.Match(EvoLangParserCLOSE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISandboxDefineContentContext is an interface to support dynamic dispatch.
type ISandboxDefineContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsSandboxDefineContentContext differentiates from other interfaces.
	IsSandboxDefineContentContext()
}

type SandboxDefineContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySandboxDefineContentContext() *SandboxDefineContentContext {
	var p = new(SandboxDefineContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_sandboxDefineContent
	return p
}

func InitEmptySandboxDefineContentContext(p *SandboxDefineContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_sandboxDefineContent
}

func (*SandboxDefineContentContext) IsSandboxDefineContentContext() {}

func NewSandboxDefineContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SandboxDefineContentContext {
	var p = new(SandboxDefineContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_sandboxDefineContent

	return p
}

func (s *SandboxDefineContentContext) GetParser() antlr.Parser { return s.parser }

func (s *SandboxDefineContentContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserID)
}

func (s *SandboxDefineContentContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, i)
}

func (s *SandboxDefineContentContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserCOLON)
}

func (s *SandboxDefineContentContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserCOLON, i)
}

func (s *SandboxDefineContentContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserCOMMA)
}

func (s *SandboxDefineContentContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserCOMMA, i)
}

func (s *SandboxDefineContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SandboxDefineContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SandboxDefineContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterSandboxDefineContent(s)
	}
}

func (s *SandboxDefineContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitSandboxDefineContent(s)
	}
}

func (s *SandboxDefineContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitSandboxDefineContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) SandboxDefineContent() (localctx ISandboxDefineContentContext) {
	localctx = NewSandboxDefineContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, EvoLangParserRULE_sandboxDefineContent)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(500)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(501)
		p.Match(EvoLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(502)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserCOMMA {
		{
			p.SetState(503)
			p.Match(EvoLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(504)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(505)
			p.Match(EvoLangParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(506)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(511)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *EvoLangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 41:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *EvoLangParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
