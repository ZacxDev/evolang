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
		"'input'", "'@key'", "'Function'", "'Rules'", "'Input'", "'*'", "'/'",
		"'+'", "'-'", "'>'", "'<'", "'=='", "'!='", "'>='", "'<='", "'&&'",
		"'||'", "'in'", "", "", "", "", "", "", "", "'{'", "'}'", "'['", "']'",
		"'('", "')'", "':'", "'#'", "'.'", "','", "'@'", "'|'", "'[sandbox lang=\"go\"]'",
		"'[/sandbox]'",
	}
	staticData.SymbolicNames = []string{
		"", "MODEL_DEF", "EVENTS_DEF", "ENUM_DEF", "REFERENCES_DEF", "LINKED_BY_DEF",
		"BODY_DEF", "AUTHENTICATE_DEF", "RETURN_DEF", "TYPE_INT_DEF", "TYPE_STRING_DEF",
		"TYPE_BOOLEAN_DEF", "IF_DEF", "MAIN_DEF", "PRINT_DEF", "LISTEN_DEF",
		"QUERY_DEF", "RULE_DEF", "MUTATION_DEF", "INPUT_DEF", "KEY_DEF", "OPERATION_FUNCTION_DEF",
		"OPERATION_RULES_DEF", "OPERATION_INPUT_DEF", "OP_MUL", "OP_DIV", "OP_SUM",
		"OP_SUB", "OP_GT", "OP_LT", "OP_EQEQ", "OP_NEQ", "OP_GTEQ", "OP_LTEQ",
		"OP_AND", "OP_OR", "OP_IN", "ID", "STRING", "INT", "BOOLEAN", "WS",
		"COMMENT", "WHITESPACE", "OPEN_BRACKET", "CLOSE_BRACKET", "OPEN_S_BRACKET",
		"CLOSE_S_BRACKET", "OPEN_PAREN", "CLOSE_PAREN", "COLON", "HASH", "DOT",
		"COMMA", "AT", "PIPE", "OPEN_SANDBOX_MODE", "CLOSE_SANDBOX_MODE", "SANDBOX",
		"SANDBOX_CONTENT",
	}
	staticData.RuleNames = []string{
		"prog", "statement", "modelDef", "modelBody", "modelFieldDef", "ruleDef",
		"ruleBody", "ruleStatement", "ruleParameters", "functionBody", "conditionalStatement",
		"condition", "returnStatement", "ruleConditionExpression", "equalityExpression",
		"inExpression", "ruleExpressionPrimary", "mutationDef", "mutationBody",
		"mutationStatement", "mutationDetails", "mutationRulesApplication",
		"orExpression", "andExpression", "mutationRulesApplicationExpressionPrimary",
		"inputDef", "inputBody", "inputFieldDef", "eventDef", "eventChannelHandler",
		"enumDef", "references", "linkedBy", "action", "paramList", "bodyDef",
		"authenticateDef", "returnDef", "typeDef", "arrayDef", "eventType",
		"expression", "exprList", "mainFunction", "mainStatements", "mainStatement",
		"printStatement", "serverStatement",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 59, 419, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 1, 0, 4, 0, 98, 8, 0, 11, 0, 12, 0, 99, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 110, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 4, 3, 120, 8, 3, 11, 3, 12, 3, 121, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 3, 4, 129, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 5, 6,
		137, 8, 6, 10, 6, 12, 6, 140, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 156, 8, 8, 10, 8,
		12, 8, 159, 9, 8, 1, 9, 1, 9, 5, 9, 163, 8, 9, 10, 9, 12, 9, 166, 9, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		13, 1, 13, 3, 13, 180, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 18, 5, 18, 200, 8, 18, 10, 18, 12, 18, 203, 9, 18, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 22, 5, 22, 229, 8, 22, 10, 22, 12, 22, 232, 9, 22, 1, 23, 1, 23, 1,
		23, 5, 23, 237, 8, 23, 10, 23, 12, 23, 240, 9, 23, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 3, 24, 247, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 26, 5, 26, 256, 8, 26, 10, 26, 12, 26, 259, 9, 26, 1, 27, 1, 27,
		1, 27, 1, 27, 3, 27, 265, 8, 27, 1, 28, 1, 28, 1, 28, 4, 28, 270, 8, 28,
		11, 28, 12, 28, 271, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 5, 30, 294, 8, 30, 10, 30, 12, 30, 297, 9, 30, 1, 30, 1,
		30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 5, 34, 318, 8, 34, 10,
		34, 12, 34, 321, 9, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 37,
		1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 337, 8, 38, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 5, 40, 346, 8, 40, 10, 40,
		12, 40, 349, 9, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 3, 41, 359, 8, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 366, 8,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 41, 5, 41, 380, 8, 41, 10, 41, 12, 41, 383, 9, 41, 1, 42, 1,
		42, 1, 42, 5, 42, 388, 8, 42, 10, 42, 12, 42, 391, 9, 42, 1, 43, 1, 43,
		1, 43, 1, 43, 1, 43, 1, 44, 4, 44, 399, 8, 44, 11, 44, 12, 44, 400, 1,
		45, 1, 45, 3, 45, 405, 8, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 0, 1, 82, 48, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
		80, 82, 84, 86, 88, 90, 92, 94, 0, 4, 1, 0, 24, 25, 1, 0, 26, 27, 1, 0,
		28, 33, 1, 0, 34, 35, 412, 0, 97, 1, 0, 0, 0, 2, 109, 1, 0, 0, 0, 4, 111,
		1, 0, 0, 0, 6, 119, 1, 0, 0, 0, 8, 123, 1, 0, 0, 0, 10, 130, 1, 0, 0, 0,
		12, 138, 1, 0, 0, 0, 14, 141, 1, 0, 0, 0, 16, 150, 1, 0, 0, 0, 18, 164,
		1, 0, 0, 0, 20, 167, 1, 0, 0, 0, 22, 173, 1, 0, 0, 0, 24, 175, 1, 0, 0,
		0, 26, 179, 1, 0, 0, 0, 28, 181, 1, 0, 0, 0, 30, 185, 1, 0, 0, 0, 32, 189,
		1, 0, 0, 0, 34, 193, 1, 0, 0, 0, 36, 201, 1, 0, 0, 0, 38, 204, 1, 0, 0,
		0, 40, 209, 1, 0, 0, 0, 42, 223, 1, 0, 0, 0, 44, 225, 1, 0, 0, 0, 46, 233,
		1, 0, 0, 0, 48, 246, 1, 0, 0, 0, 50, 248, 1, 0, 0, 0, 52, 257, 1, 0, 0,
		0, 54, 260, 1, 0, 0, 0, 56, 266, 1, 0, 0, 0, 58, 275, 1, 0, 0, 0, 60, 288,
		1, 0, 0, 0, 62, 300, 1, 0, 0, 0, 64, 306, 1, 0, 0, 0, 66, 312, 1, 0, 0,
		0, 68, 314, 1, 0, 0, 0, 70, 322, 1, 0, 0, 0, 72, 325, 1, 0, 0, 0, 74, 328,
		1, 0, 0, 0, 76, 336, 1, 0, 0, 0, 78, 338, 1, 0, 0, 0, 80, 342, 1, 0, 0,
		0, 82, 365, 1, 0, 0, 0, 84, 384, 1, 0, 0, 0, 86, 392, 1, 0, 0, 0, 88, 398,
		1, 0, 0, 0, 90, 404, 1, 0, 0, 0, 92, 406, 1, 0, 0, 0, 94, 411, 1, 0, 0,
		0, 96, 98, 3, 2, 1, 0, 97, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 97,
		1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 5, 0,
		0, 1, 102, 1, 1, 0, 0, 0, 103, 110, 3, 4, 2, 0, 104, 110, 3, 60, 30, 0,
		105, 110, 3, 10, 5, 0, 106, 110, 3, 34, 17, 0, 107, 110, 3, 50, 25, 0,
		108, 110, 3, 86, 43, 0, 109, 103, 1, 0, 0, 0, 109, 104, 1, 0, 0, 0, 109,
		105, 1, 0, 0, 0, 109, 106, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 108,
		1, 0, 0, 0, 110, 3, 1, 0, 0, 0, 111, 112, 5, 1, 0, 0, 112, 113, 5, 37,
		0, 0, 113, 114, 5, 44, 0, 0, 114, 115, 3, 6, 3, 0, 115, 116, 5, 45, 0,
		0, 116, 5, 1, 0, 0, 0, 117, 120, 3, 8, 4, 0, 118, 120, 3, 56, 28, 0, 119,
		117, 1, 0, 0, 0, 119, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 119,
		1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 7, 1, 0, 0, 0, 123, 124, 5, 37,
		0, 0, 124, 125, 5, 50, 0, 0, 125, 128, 3, 76, 38, 0, 126, 129, 3, 62, 31,
		0, 127, 129, 3, 64, 32, 0, 128, 126, 1, 0, 0, 0, 128, 127, 1, 0, 0, 0,
		128, 129, 1, 0, 0, 0, 129, 9, 1, 0, 0, 0, 130, 131, 5, 17, 0, 0, 131, 132,
		5, 44, 0, 0, 132, 133, 3, 12, 6, 0, 133, 134, 5, 45, 0, 0, 134, 11, 1,
		0, 0, 0, 135, 137, 3, 14, 7, 0, 136, 135, 1, 0, 0, 0, 137, 140, 1, 0, 0,
		0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 13, 1, 0, 0, 0, 140,
		138, 1, 0, 0, 0, 141, 142, 5, 37, 0, 0, 142, 143, 5, 50, 0, 0, 143, 144,
		5, 48, 0, 0, 144, 145, 3, 16, 8, 0, 145, 146, 5, 49, 0, 0, 146, 147, 5,
		44, 0, 0, 147, 148, 3, 18, 9, 0, 148, 149, 5, 45, 0, 0, 149, 15, 1, 0,
		0, 0, 150, 151, 5, 37, 0, 0, 151, 157, 5, 37, 0, 0, 152, 153, 5, 53, 0,
		0, 153, 154, 5, 37, 0, 0, 154, 156, 5, 37, 0, 0, 155, 152, 1, 0, 0, 0,
		156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158,
		17, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 163, 3, 20, 10, 0, 161, 163,
		3, 24, 12, 0, 162, 160, 1, 0, 0, 0, 162, 161, 1, 0, 0, 0, 163, 166, 1,
		0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 19, 1, 0, 0,
		0, 166, 164, 1, 0, 0, 0, 167, 168, 5, 12, 0, 0, 168, 169, 3, 22, 11, 0,
		169, 170, 5, 44, 0, 0, 170, 171, 3, 24, 12, 0, 171, 172, 5, 45, 0, 0, 172,
		21, 1, 0, 0, 0, 173, 174, 3, 26, 13, 0, 174, 23, 1, 0, 0, 0, 175, 176,
		3, 74, 37, 0, 176, 25, 1, 0, 0, 0, 177, 180, 3, 28, 14, 0, 178, 180, 3,
		30, 15, 0, 179, 177, 1, 0, 0, 0, 179, 178, 1, 0, 0, 0, 180, 27, 1, 0, 0,
		0, 181, 182, 3, 32, 16, 0, 182, 183, 5, 30, 0, 0, 183, 184, 3, 32, 16,
		0, 184, 29, 1, 0, 0, 0, 185, 186, 3, 32, 16, 0, 186, 187, 5, 36, 0, 0,
		187, 188, 3, 32, 16, 0, 188, 31, 1, 0, 0, 0, 189, 190, 5, 37, 0, 0, 190,
		191, 5, 52, 0, 0, 191, 192, 5, 37, 0, 0, 192, 33, 1, 0, 0, 0, 193, 194,
		5, 18, 0, 0, 194, 195, 5, 44, 0, 0, 195, 196, 3, 36, 18, 0, 196, 197, 5,
		45, 0, 0, 197, 35, 1, 0, 0, 0, 198, 200, 3, 38, 19, 0, 199, 198, 1, 0,
		0, 0, 200, 203, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0,
		202, 37, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 204, 205, 5, 37, 0, 0, 205,
		206, 5, 44, 0, 0, 206, 207, 3, 40, 20, 0, 207, 208, 5, 45, 0, 0, 208, 39,
		1, 0, 0, 0, 209, 210, 5, 23, 0, 0, 210, 211, 5, 50, 0, 0, 211, 212, 5,
		37, 0, 0, 212, 213, 5, 22, 0, 0, 213, 214, 5, 50, 0, 0, 214, 215, 5, 48,
		0, 0, 215, 216, 3, 42, 21, 0, 216, 217, 5, 49, 0, 0, 217, 218, 5, 21, 0,
		0, 218, 219, 5, 50, 0, 0, 219, 220, 5, 44, 0, 0, 220, 221, 5, 58, 0, 0,
		221, 222, 5, 45, 0, 0, 222, 41, 1, 0, 0, 0, 223, 224, 3, 44, 22, 0, 224,
		43, 1, 0, 0, 0, 225, 230, 3, 46, 23, 0, 226, 227, 5, 35, 0, 0, 227, 229,
		3, 46, 23, 0, 228, 226, 1, 0, 0, 0, 229, 232, 1, 0, 0, 0, 230, 228, 1,
		0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 45, 1, 0, 0, 0, 232, 230, 1, 0, 0,
		0, 233, 238, 3, 48, 24, 0, 234, 235, 5, 34, 0, 0, 235, 237, 3, 48, 24,
		0, 236, 234, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 238,
		239, 1, 0, 0, 0, 239, 47, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 241, 247, 5,
		37, 0, 0, 242, 243, 5, 48, 0, 0, 243, 244, 3, 82, 41, 0, 244, 245, 5, 49,
		0, 0, 245, 247, 1, 0, 0, 0, 246, 241, 1, 0, 0, 0, 246, 242, 1, 0, 0, 0,
		247, 49, 1, 0, 0, 0, 248, 249, 5, 19, 0, 0, 249, 250, 5, 37, 0, 0, 250,
		251, 5, 44, 0, 0, 251, 252, 3, 52, 26, 0, 252, 253, 5, 45, 0, 0, 253, 51,
		1, 0, 0, 0, 254, 256, 3, 54, 27, 0, 255, 254, 1, 0, 0, 0, 256, 259, 1,
		0, 0, 0, 257, 255, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 53, 1, 0, 0,
		0, 259, 257, 1, 0, 0, 0, 260, 261, 5, 37, 0, 0, 261, 262, 5, 50, 0, 0,
		262, 264, 3, 76, 38, 0, 263, 265, 5, 20, 0, 0, 264, 263, 1, 0, 0, 0, 264,
		265, 1, 0, 0, 0, 265, 55, 1, 0, 0, 0, 266, 267, 5, 2, 0, 0, 267, 269, 5,
		44, 0, 0, 268, 270, 3, 58, 29, 0, 269, 268, 1, 0, 0, 0, 270, 271, 1, 0,
		0, 0, 271, 269, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 273, 1, 0, 0, 0,
		273, 274, 5, 45, 0, 0, 274, 57, 1, 0, 0, 0, 275, 276, 5, 51, 0, 0, 276,
		277, 5, 52, 0, 0, 277, 278, 5, 37, 0, 0, 278, 279, 5, 50, 0, 0, 279, 280,
		5, 48, 0, 0, 280, 281, 5, 37, 0, 0, 281, 282, 5, 53, 0, 0, 282, 283, 5,
		37, 0, 0, 283, 284, 5, 49, 0, 0, 284, 285, 5, 44, 0, 0, 285, 286, 5, 58,
		0, 0, 286, 287, 5, 45, 0, 0, 287, 59, 1, 0, 0, 0, 288, 289, 5, 3, 0, 0,
		289, 290, 5, 37, 0, 0, 290, 291, 5, 44, 0, 0, 291, 295, 5, 37, 0, 0, 292,
		294, 5, 37, 0, 0, 293, 292, 1, 0, 0, 0, 294, 297, 1, 0, 0, 0, 295, 293,
		1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 298, 1, 0, 0, 0, 297, 295, 1, 0,
		0, 0, 298, 299, 5, 45, 0, 0, 299, 61, 1, 0, 0, 0, 300, 301, 5, 4, 0, 0,
		301, 302, 5, 48, 0, 0, 302, 303, 5, 52, 0, 0, 303, 304, 5, 37, 0, 0, 304,
		305, 5, 49, 0, 0, 305, 63, 1, 0, 0, 0, 306, 307, 5, 5, 0, 0, 307, 308,
		5, 48, 0, 0, 308, 309, 5, 52, 0, 0, 309, 310, 5, 37, 0, 0, 310, 311, 5,
		49, 0, 0, 311, 65, 1, 0, 0, 0, 312, 313, 5, 37, 0, 0, 313, 67, 1, 0, 0,
		0, 314, 319, 5, 37, 0, 0, 315, 316, 5, 53, 0, 0, 316, 318, 5, 37, 0, 0,
		317, 315, 1, 0, 0, 0, 318, 321, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 319,
		320, 1, 0, 0, 0, 320, 69, 1, 0, 0, 0, 321, 319, 1, 0, 0, 0, 322, 323, 5,
		6, 0, 0, 323, 324, 5, 37, 0, 0, 324, 71, 1, 0, 0, 0, 325, 326, 5, 7, 0,
		0, 326, 327, 5, 37, 0, 0, 327, 73, 1, 0, 0, 0, 328, 329, 5, 8, 0, 0, 329,
		330, 3, 82, 41, 0, 330, 75, 1, 0, 0, 0, 331, 337, 5, 9, 0, 0, 332, 337,
		5, 10, 0, 0, 333, 337, 5, 11, 0, 0, 334, 337, 3, 78, 39, 0, 335, 337, 5,
		37, 0, 0, 336, 331, 1, 0, 0, 0, 336, 332, 1, 0, 0, 0, 336, 333, 1, 0, 0,
		0, 336, 334, 1, 0, 0, 0, 336, 335, 1, 0, 0, 0, 337, 77, 1, 0, 0, 0, 338,
		339, 5, 46, 0, 0, 339, 340, 3, 76, 38, 0, 340, 341, 5, 47, 0, 0, 341, 79,
		1, 0, 0, 0, 342, 347, 5, 37, 0, 0, 343, 344, 5, 53, 0, 0, 344, 346, 5,
		37, 0, 0, 345, 343, 1, 0, 0, 0, 346, 349, 1, 0, 0, 0, 347, 345, 1, 0, 0,
		0, 347, 348, 1, 0, 0, 0, 348, 81, 1, 0, 0, 0, 349, 347, 1, 0, 0, 0, 350,
		351, 6, 41, -1, 0, 351, 352, 5, 48, 0, 0, 352, 353, 3, 82, 41, 0, 353,
		354, 5, 49, 0, 0, 354, 366, 1, 0, 0, 0, 355, 356, 5, 37, 0, 0, 356, 358,
		5, 48, 0, 0, 357, 359, 3, 84, 42, 0, 358, 357, 1, 0, 0, 0, 358, 359, 1,
		0, 0, 0, 359, 360, 1, 0, 0, 0, 360, 366, 5, 49, 0, 0, 361, 366, 5, 37,
		0, 0, 362, 366, 5, 39, 0, 0, 363, 366, 5, 38, 0, 0, 364, 366, 5, 40, 0,
		0, 365, 350, 1, 0, 0, 0, 365, 355, 1, 0, 0, 0, 365, 361, 1, 0, 0, 0, 365,
		362, 1, 0, 0, 0, 365, 363, 1, 0, 0, 0, 365, 364, 1, 0, 0, 0, 366, 381,
		1, 0, 0, 0, 367, 368, 10, 10, 0, 0, 368, 369, 7, 0, 0, 0, 369, 380, 3,
		82, 41, 11, 370, 371, 10, 9, 0, 0, 371, 372, 7, 1, 0, 0, 372, 380, 3, 82,
		41, 10, 373, 374, 10, 8, 0, 0, 374, 375, 7, 2, 0, 0, 375, 380, 3, 82, 41,
		9, 376, 377, 10, 7, 0, 0, 377, 378, 7, 3, 0, 0, 378, 380, 3, 82, 41, 8,
		379, 367, 1, 0, 0, 0, 379, 370, 1, 0, 0, 0, 379, 373, 1, 0, 0, 0, 379,
		376, 1, 0, 0, 0, 380, 383, 1, 0, 0, 0, 381, 379, 1, 0, 0, 0, 381, 382,
		1, 0, 0, 0, 382, 83, 1, 0, 0, 0, 383, 381, 1, 0, 0, 0, 384, 389, 3, 82,
		41, 0, 385, 386, 5, 53, 0, 0, 386, 388, 3, 82, 41, 0, 387, 385, 1, 0, 0,
		0, 388, 391, 1, 0, 0, 0, 389, 387, 1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390,
		85, 1, 0, 0, 0, 391, 389, 1, 0, 0, 0, 392, 393, 5, 13, 0, 0, 393, 394,
		5, 44, 0, 0, 394, 395, 3, 88, 44, 0, 395, 396, 5, 45, 0, 0, 396, 87, 1,
		0, 0, 0, 397, 399, 3, 90, 45, 0, 398, 397, 1, 0, 0, 0, 399, 400, 1, 0,
		0, 0, 400, 398, 1, 0, 0, 0, 400, 401, 1, 0, 0, 0, 401, 89, 1, 0, 0, 0,
		402, 405, 3, 92, 46, 0, 403, 405, 3, 94, 47, 0, 404, 402, 1, 0, 0, 0, 404,
		403, 1, 0, 0, 0, 405, 91, 1, 0, 0, 0, 406, 407, 5, 14, 0, 0, 407, 408,
		5, 48, 0, 0, 408, 409, 5, 38, 0, 0, 409, 410, 5, 49, 0, 0, 410, 93, 1,
		0, 0, 0, 411, 412, 5, 37, 0, 0, 412, 413, 5, 52, 0, 0, 413, 414, 5, 15,
		0, 0, 414, 415, 5, 48, 0, 0, 415, 416, 5, 39, 0, 0, 416, 417, 5, 49, 0,
		0, 417, 95, 1, 0, 0, 0, 28, 99, 109, 119, 121, 128, 138, 157, 162, 164,
		179, 201, 230, 238, 246, 257, 264, 271, 295, 319, 336, 347, 358, 365, 379,
		381, 389, 400, 404,
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
	EvoLangParserEOF                    = antlr.TokenEOF
	EvoLangParserMODEL_DEF              = 1
	EvoLangParserEVENTS_DEF             = 2
	EvoLangParserENUM_DEF               = 3
	EvoLangParserREFERENCES_DEF         = 4
	EvoLangParserLINKED_BY_DEF          = 5
	EvoLangParserBODY_DEF               = 6
	EvoLangParserAUTHENTICATE_DEF       = 7
	EvoLangParserRETURN_DEF             = 8
	EvoLangParserTYPE_INT_DEF           = 9
	EvoLangParserTYPE_STRING_DEF        = 10
	EvoLangParserTYPE_BOOLEAN_DEF       = 11
	EvoLangParserIF_DEF                 = 12
	EvoLangParserMAIN_DEF               = 13
	EvoLangParserPRINT_DEF              = 14
	EvoLangParserLISTEN_DEF             = 15
	EvoLangParserQUERY_DEF              = 16
	EvoLangParserRULE_DEF               = 17
	EvoLangParserMUTATION_DEF           = 18
	EvoLangParserINPUT_DEF              = 19
	EvoLangParserKEY_DEF                = 20
	EvoLangParserOPERATION_FUNCTION_DEF = 21
	EvoLangParserOPERATION_RULES_DEF    = 22
	EvoLangParserOPERATION_INPUT_DEF    = 23
	EvoLangParserOP_MUL                 = 24
	EvoLangParserOP_DIV                 = 25
	EvoLangParserOP_SUM                 = 26
	EvoLangParserOP_SUB                 = 27
	EvoLangParserOP_GT                  = 28
	EvoLangParserOP_LT                  = 29
	EvoLangParserOP_EQEQ                = 30
	EvoLangParserOP_NEQ                 = 31
	EvoLangParserOP_GTEQ                = 32
	EvoLangParserOP_LTEQ                = 33
	EvoLangParserOP_AND                 = 34
	EvoLangParserOP_OR                  = 35
	EvoLangParserOP_IN                  = 36
	EvoLangParserID                     = 37
	EvoLangParserSTRING                 = 38
	EvoLangParserINT                    = 39
	EvoLangParserBOOLEAN                = 40
	EvoLangParserWS                     = 41
	EvoLangParserCOMMENT                = 42
	EvoLangParserWHITESPACE             = 43
	EvoLangParserOPEN_BRACKET           = 44
	EvoLangParserCLOSE_BRACKET          = 45
	EvoLangParserOPEN_S_BRACKET         = 46
	EvoLangParserCLOSE_S_BRACKET        = 47
	EvoLangParserOPEN_PAREN             = 48
	EvoLangParserCLOSE_PAREN            = 49
	EvoLangParserCOLON                  = 50
	EvoLangParserHASH                   = 51
	EvoLangParserDOT                    = 52
	EvoLangParserCOMMA                  = 53
	EvoLangParserAT                     = 54
	EvoLangParserPIPE                   = 55
	EvoLangParserOPEN_SANDBOX_MODE      = 56
	EvoLangParserCLOSE_SANDBOX_MODE     = 57
	EvoLangParserSANDBOX                = 58
	EvoLangParserSANDBOX_CONTENT        = 59
)

// EvoLangParser rules.
const (
	EvoLangParserRULE_prog                                      = 0
	EvoLangParserRULE_statement                                 = 1
	EvoLangParserRULE_modelDef                                  = 2
	EvoLangParserRULE_modelBody                                 = 3
	EvoLangParserRULE_modelFieldDef                             = 4
	EvoLangParserRULE_ruleDef                                   = 5
	EvoLangParserRULE_ruleBody                                  = 6
	EvoLangParserRULE_ruleStatement                             = 7
	EvoLangParserRULE_ruleParameters                            = 8
	EvoLangParserRULE_functionBody                              = 9
	EvoLangParserRULE_conditionalStatement                      = 10
	EvoLangParserRULE_condition                                 = 11
	EvoLangParserRULE_returnStatement                           = 12
	EvoLangParserRULE_ruleConditionExpression                   = 13
	EvoLangParserRULE_equalityExpression                        = 14
	EvoLangParserRULE_inExpression                              = 15
	EvoLangParserRULE_ruleExpressionPrimary                     = 16
	EvoLangParserRULE_mutationDef                               = 17
	EvoLangParserRULE_mutationBody                              = 18
	EvoLangParserRULE_mutationStatement                         = 19
	EvoLangParserRULE_mutationDetails                           = 20
	EvoLangParserRULE_mutationRulesApplication                  = 21
	EvoLangParserRULE_orExpression                              = 22
	EvoLangParserRULE_andExpression                             = 23
	EvoLangParserRULE_mutationRulesApplicationExpressionPrimary = 24
	EvoLangParserRULE_inputDef                                  = 25
	EvoLangParserRULE_inputBody                                 = 26
	EvoLangParserRULE_inputFieldDef                             = 27
	EvoLangParserRULE_eventDef                                  = 28
	EvoLangParserRULE_eventChannelHandler                       = 29
	EvoLangParserRULE_enumDef                                   = 30
	EvoLangParserRULE_references                                = 31
	EvoLangParserRULE_linkedBy                                  = 32
	EvoLangParserRULE_action                                    = 33
	EvoLangParserRULE_paramList                                 = 34
	EvoLangParserRULE_bodyDef                                   = 35
	EvoLangParserRULE_authenticateDef                           = 36
	EvoLangParserRULE_returnDef                                 = 37
	EvoLangParserRULE_typeDef                                   = 38
	EvoLangParserRULE_arrayDef                                  = 39
	EvoLangParserRULE_eventType                                 = 40
	EvoLangParserRULE_expression                                = 41
	EvoLangParserRULE_exprList                                  = 42
	EvoLangParserRULE_mainFunction                              = 43
	EvoLangParserRULE_mainStatements                            = 44
	EvoLangParserRULE_mainStatement                             = 45
	EvoLangParserRULE_printStatement                            = 46
	EvoLangParserRULE_serverStatement                           = 47
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
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&925706) != 0) {
		{
			p.SetState(96)
			p.Statement()
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(101)
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
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EvoLangParserMODEL_DEF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.ModelDef()
		}

	case EvoLangParserENUM_DEF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.EnumDef()
		}

	case EvoLangParserRULE_DEF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.RuleDef()
		}

	case EvoLangParserMUTATION_DEF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(106)
			p.MutationDef()
		}

	case EvoLangParserINPUT_DEF:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(107)
			p.InputDef()
		}

	case EvoLangParserMAIN_DEF:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(108)
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
		p.SetState(111)
		p.Match(EvoLangParserMODEL_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.ModelBody()
	}
	{
		p.SetState(115)
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
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EvoLangParserEVENTS_DEF || _la == EvoLangParserID {
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case EvoLangParserID:
			{
				p.SetState(117)
				p.ModelFieldDef()
			}

		case EvoLangParserEVENTS_DEF:
			{
				p.SetState(118)
				p.EventDef()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(121)
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
		p.SetState(123)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(EvoLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.TypeDef()
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case EvoLangParserREFERENCES_DEF:
		{
			p.SetState(126)
			p.References()
		}

	case EvoLangParserLINKED_BY_DEF:
		{
			p.SetState(127)
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
		p.SetState(130)
		p.Match(EvoLangParserRULE_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.RuleBody()
	}
	{
		p.SetState(133)
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
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserID {
		{
			p.SetState(135)
			p.RuleStatement()
		}

		p.SetState(140)
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
		p.SetState(141)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.Match(EvoLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Match(EvoLangParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.RuleParameters()
	}
	{
		p.SetState(145)
		p.Match(EvoLangParserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.FunctionBody()
	}
	{
		p.SetState(148)
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
		p.SetState(150)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserCOMMA {
		{
			p.SetState(152)
			p.Match(EvoLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(159)
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
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserRETURN_DEF || _la == EvoLangParserIF_DEF {
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case EvoLangParserIF_DEF:
			{
				p.SetState(160)
				p.ConditionalStatement()
			}

		case EvoLangParserRETURN_DEF:
			{
				p.SetState(161)
				p.ReturnStatement()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(166)
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
	p.EnterRule(localctx, 20, EvoLangParserRULE_conditionalStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(EvoLangParserIF_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Condition()
	}
	{
		p.SetState(169)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.ReturnStatement()
	}
	{
		p.SetState(171)
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
	p.EnterRule(localctx, 22, EvoLangParserRULE_condition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
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
	p.EnterRule(localctx, 24, EvoLangParserRULE_returnStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
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
	p.EnterRule(localctx, 26, EvoLangParserRULE_ruleConditionExpression)
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(177)
			p.EqualityExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(178)
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
	p.EnterRule(localctx, 28, EvoLangParserRULE_equalityExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.RuleExpressionPrimary()
	}
	{
		p.SetState(182)
		p.Match(EvoLangParserOP_EQEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(183)
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
	p.EnterRule(localctx, 30, EvoLangParserRULE_inExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.RuleExpressionPrimary()
	}
	{
		p.SetState(186)
		p.Match(EvoLangParserOP_IN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
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
	p.EnterRule(localctx, 32, EvoLangParserRULE_ruleExpressionPrimary)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(190)
		p.Match(EvoLangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)
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
	p.EnterRule(localctx, 34, EvoLangParserRULE_mutationDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(EvoLangParserMUTATION_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
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
		p.MutationBody()
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
	p.EnterRule(localctx, 36, EvoLangParserRULE_mutationBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserID {
		{
			p.SetState(198)
			p.MutationStatement()
		}

		p.SetState(203)
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
	OPEN_BRACKET() antlr.TerminalNode
	MutationDetails() IMutationDetailsContext
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

func (s *MutationStatementContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_BRACKET, 0)
}

func (s *MutationStatementContext) MutationDetails() IMutationDetailsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMutationDetailsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMutationDetailsContext)
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
	p.EnterRule(localctx, 38, EvoLangParserRULE_mutationStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.MutationDetails()
	}
	{
		p.SetState(207)
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

// IMutationDetailsContext is an interface to support dynamic dispatch.
type IMutationDetailsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPERATION_INPUT_DEF() antlr.TerminalNode
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	ID() antlr.TerminalNode
	OPERATION_RULES_DEF() antlr.TerminalNode
	OPEN_PAREN() antlr.TerminalNode
	MutationRulesApplication() IMutationRulesApplicationContext
	CLOSE_PAREN() antlr.TerminalNode
	OPERATION_FUNCTION_DEF() antlr.TerminalNode
	OPEN_BRACKET() antlr.TerminalNode
	SANDBOX() antlr.TerminalNode
	CLOSE_BRACKET() antlr.TerminalNode

	// IsMutationDetailsContext differentiates from other interfaces.
	IsMutationDetailsContext()
}

type MutationDetailsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMutationDetailsContext() *MutationDetailsContext {
	var p = new(MutationDetailsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mutationDetails
	return p
}

func InitEmptyMutationDetailsContext(p *MutationDetailsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mutationDetails
}

func (*MutationDetailsContext) IsMutationDetailsContext() {}

func NewMutationDetailsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MutationDetailsContext {
	var p = new(MutationDetailsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_mutationDetails

	return p
}

func (s *MutationDetailsContext) GetParser() antlr.Parser { return s.parser }

func (s *MutationDetailsContext) OPERATION_INPUT_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPERATION_INPUT_DEF, 0)
}

func (s *MutationDetailsContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserCOLON)
}

func (s *MutationDetailsContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserCOLON, i)
}

func (s *MutationDetailsContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *MutationDetailsContext) OPERATION_RULES_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPERATION_RULES_DEF, 0)
}

func (s *MutationDetailsContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_PAREN, 0)
}

func (s *MutationDetailsContext) MutationRulesApplication() IMutationRulesApplicationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMutationRulesApplicationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMutationRulesApplicationContext)
}

func (s *MutationDetailsContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_PAREN, 0)
}

func (s *MutationDetailsContext) OPERATION_FUNCTION_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPERATION_FUNCTION_DEF, 0)
}

func (s *MutationDetailsContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_BRACKET, 0)
}

func (s *MutationDetailsContext) SANDBOX() antlr.TerminalNode {
	return s.GetToken(EvoLangParserSANDBOX, 0)
}

func (s *MutationDetailsContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_BRACKET, 0)
}

func (s *MutationDetailsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutationDetailsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MutationDetailsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterMutationDetails(s)
	}
}

func (s *MutationDetailsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitMutationDetails(s)
	}
}

func (s *MutationDetailsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitMutationDetails(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) MutationDetails() (localctx IMutationDetailsContext) {
	localctx = NewMutationDetailsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, EvoLangParserRULE_mutationDetails)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(EvoLangParserOPERATION_INPUT_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		p.Match(EvoLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(212)
		p.Match(EvoLangParserOPERATION_RULES_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.Match(EvoLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Match(EvoLangParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.MutationRulesApplication()
	}
	{
		p.SetState(216)
		p.Match(EvoLangParserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.Match(EvoLangParserOPERATION_FUNCTION_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)
		p.Match(EvoLangParserCOLON)
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
		p.Match(EvoLangParserSANDBOX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
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

// IMutationRulesApplicationContext is an interface to support dynamic dispatch.
type IMutationRulesApplicationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OrExpression() IOrExpressionContext

	// IsMutationRulesApplicationContext differentiates from other interfaces.
	IsMutationRulesApplicationContext()
}

type MutationRulesApplicationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMutationRulesApplicationContext() *MutationRulesApplicationContext {
	var p = new(MutationRulesApplicationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mutationRulesApplication
	return p
}

func InitEmptyMutationRulesApplicationContext(p *MutationRulesApplicationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mutationRulesApplication
}

func (*MutationRulesApplicationContext) IsMutationRulesApplicationContext() {}

func NewMutationRulesApplicationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MutationRulesApplicationContext {
	var p = new(MutationRulesApplicationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_mutationRulesApplication

	return p
}

func (s *MutationRulesApplicationContext) GetParser() antlr.Parser { return s.parser }

func (s *MutationRulesApplicationContext) OrExpression() IOrExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *MutationRulesApplicationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutationRulesApplicationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MutationRulesApplicationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterMutationRulesApplication(s)
	}
}

func (s *MutationRulesApplicationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitMutationRulesApplication(s)
	}
}

func (s *MutationRulesApplicationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitMutationRulesApplication(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) MutationRulesApplication() (localctx IMutationRulesApplicationContext) {
	localctx = NewMutationRulesApplicationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, EvoLangParserRULE_mutationRulesApplication)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.OrExpression()
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

// IOrExpressionContext is an interface to support dynamic dispatch.
type IOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAndExpression() []IAndExpressionContext
	AndExpression(i int) IAndExpressionContext
	AllOP_OR() []antlr.TerminalNode
	OP_OR(i int) antlr.TerminalNode

	// IsOrExpressionContext differentiates from other interfaces.
	IsOrExpressionContext()
}

type OrExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrExpressionContext() *OrExpressionContext {
	var p = new(OrExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_orExpression
	return p
}

func InitEmptyOrExpressionContext(p *OrExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_orExpression
}

func (*OrExpressionContext) IsOrExpressionContext() {}

func NewOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExpressionContext {
	var p = new(OrExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_orExpression

	return p
}

func (s *OrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrExpressionContext) AllAndExpression() []IAndExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAndExpressionContext); ok {
			len++
		}
	}

	tst := make([]IAndExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAndExpressionContext); ok {
			tst[i] = t.(IAndExpressionContext)
			i++
		}
	}

	return tst
}

func (s *OrExpressionContext) AndExpression(i int) IAndExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndExpressionContext); ok {
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

	return t.(IAndExpressionContext)
}

func (s *OrExpressionContext) AllOP_OR() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserOP_OR)
}

func (s *OrExpressionContext) OP_OR(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_OR, i)
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterOrExpression(s)
	}
}

func (s *OrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitOrExpression(s)
	}
}

func (s *OrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) OrExpression() (localctx IOrExpressionContext) {
	localctx = NewOrExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, EvoLangParserRULE_orExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.AndExpression()
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserOP_OR {
		{
			p.SetState(226)
			p.Match(EvoLangParserOP_OR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.AndExpression()
		}

		p.SetState(232)
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

// IAndExpressionContext is an interface to support dynamic dispatch.
type IAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMutationRulesApplicationExpressionPrimary() []IMutationRulesApplicationExpressionPrimaryContext
	MutationRulesApplicationExpressionPrimary(i int) IMutationRulesApplicationExpressionPrimaryContext
	AllOP_AND() []antlr.TerminalNode
	OP_AND(i int) antlr.TerminalNode

	// IsAndExpressionContext differentiates from other interfaces.
	IsAndExpressionContext()
}

type AndExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExpressionContext() *AndExpressionContext {
	var p = new(AndExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_andExpression
	return p
}

func InitEmptyAndExpressionContext(p *AndExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_andExpression
}

func (*AndExpressionContext) IsAndExpressionContext() {}

func NewAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExpressionContext {
	var p = new(AndExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_andExpression

	return p
}

func (s *AndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExpressionContext) AllMutationRulesApplicationExpressionPrimary() []IMutationRulesApplicationExpressionPrimaryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMutationRulesApplicationExpressionPrimaryContext); ok {
			len++
		}
	}

	tst := make([]IMutationRulesApplicationExpressionPrimaryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMutationRulesApplicationExpressionPrimaryContext); ok {
			tst[i] = t.(IMutationRulesApplicationExpressionPrimaryContext)
			i++
		}
	}

	return tst
}

func (s *AndExpressionContext) MutationRulesApplicationExpressionPrimary(i int) IMutationRulesApplicationExpressionPrimaryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMutationRulesApplicationExpressionPrimaryContext); ok {
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

	return t.(IMutationRulesApplicationExpressionPrimaryContext)
}

func (s *AndExpressionContext) AllOP_AND() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserOP_AND)
}

func (s *AndExpressionContext) OP_AND(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserOP_AND, i)
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterAndExpression(s)
	}
}

func (s *AndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitAndExpression(s)
	}
}

func (s *AndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) AndExpression() (localctx IAndExpressionContext) {
	localctx = NewAndExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, EvoLangParserRULE_andExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.MutationRulesApplicationExpressionPrimary()
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserOP_AND {
		{
			p.SetState(234)
			p.Match(EvoLangParserOP_AND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)
			p.MutationRulesApplicationExpressionPrimary()
		}

		p.SetState(240)
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

// IMutationRulesApplicationExpressionPrimaryContext is an interface to support dynamic dispatch.
type IMutationRulesApplicationExpressionPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	OPEN_PAREN() antlr.TerminalNode
	Expression() IExpressionContext
	CLOSE_PAREN() antlr.TerminalNode

	// IsMutationRulesApplicationExpressionPrimaryContext differentiates from other interfaces.
	IsMutationRulesApplicationExpressionPrimaryContext()
}

type MutationRulesApplicationExpressionPrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMutationRulesApplicationExpressionPrimaryContext() *MutationRulesApplicationExpressionPrimaryContext {
	var p = new(MutationRulesApplicationExpressionPrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mutationRulesApplicationExpressionPrimary
	return p
}

func InitEmptyMutationRulesApplicationExpressionPrimaryContext(p *MutationRulesApplicationExpressionPrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_mutationRulesApplicationExpressionPrimary
}

func (*MutationRulesApplicationExpressionPrimaryContext) IsMutationRulesApplicationExpressionPrimaryContext() {
}

func NewMutationRulesApplicationExpressionPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MutationRulesApplicationExpressionPrimaryContext {
	var p = new(MutationRulesApplicationExpressionPrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_mutationRulesApplicationExpressionPrimary

	return p
}

func (s *MutationRulesApplicationExpressionPrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *MutationRulesApplicationExpressionPrimaryContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *MutationRulesApplicationExpressionPrimaryContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_PAREN, 0)
}

func (s *MutationRulesApplicationExpressionPrimaryContext) Expression() IExpressionContext {
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

func (s *MutationRulesApplicationExpressionPrimaryContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_PAREN, 0)
}

func (s *MutationRulesApplicationExpressionPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutationRulesApplicationExpressionPrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MutationRulesApplicationExpressionPrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.EnterMutationRulesApplicationExpressionPrimary(s)
	}
}

func (s *MutationRulesApplicationExpressionPrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangParserListener); ok {
		listenerT.ExitMutationRulesApplicationExpressionPrimary(s)
	}
}

func (s *MutationRulesApplicationExpressionPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangParserVisitor:
		return t.VisitMutationRulesApplicationExpressionPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) MutationRulesApplicationExpressionPrimary() (localctx IMutationRulesApplicationExpressionPrimaryContext) {
	localctx = NewMutationRulesApplicationExpressionPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, EvoLangParserRULE_mutationRulesApplicationExpressionPrimary)
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EvoLangParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvoLangParserOPEN_PAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)
			p.Match(EvoLangParserOPEN_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.expression(0)
		}
		{
			p.SetState(244)
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
		p.SetState(248)
		p.Match(EvoLangParserINPUT_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(251)
		p.InputBody()
	}
	{
		p.SetState(252)
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
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserID {
		{
			p.SetState(254)
			p.InputFieldDef()
		}

		p.SetState(259)
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
	KEY_DEF() antlr.TerminalNode

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

func (s *InputFieldDefContext) KEY_DEF() antlr.TerminalNode {
	return s.GetToken(EvoLangParserKEY_DEF, 0)
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
		p.SetState(260)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Match(EvoLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.TypeDef()
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EvoLangParserKEY_DEF {
		{
			p.SetState(263)
			p.Match(EvoLangParserKEY_DEF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
		p.SetState(266)
		p.Match(EvoLangParserEVENTS_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EvoLangParserHASH {
		{
			p.SetState(268)
			p.EventChannelHandler()
		}

		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(273)
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
		p.SetState(275)
		p.Match(EvoLangParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(276)
		p.Match(EvoLangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(277)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
		p.Match(EvoLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(279)
		p.Match(EvoLangParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(280)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(281)
		p.Match(EvoLangParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(282)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(283)
		p.Match(EvoLangParserCLOSE_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(284)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(285)
		p.Match(EvoLangParserSANDBOX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)
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
		p.SetState(288)
		p.Match(EvoLangParserENUM_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(289)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(290)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserID {
		{
			p.SetState(292)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(298)
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
	ID() antlr.TerminalNode
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

func (s *ReferencesContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.Match(EvoLangParserREFERENCES_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(301)
		p.Match(EvoLangParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(302)
		p.Match(EvoLangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(304)
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
		p.SetState(306)
		p.Match(EvoLangParserLINKED_BY_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.Match(EvoLangParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(308)
		p.Match(EvoLangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(309)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(310)
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
		p.SetState(312)
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
		p.SetState(314)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserCOMMA {
		{
			p.SetState(315)
			p.Match(EvoLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(321)
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
		p.SetState(322)
		p.Match(EvoLangParserBODY_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)
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
		p.SetState(325)
		p.Match(EvoLangParserAUTHENTICATE_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(326)
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
		p.SetState(328)
		p.Match(EvoLangParserRETURN_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(329)
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
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EvoLangParserTYPE_INT_DEF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(331)
			p.Match(EvoLangParserTYPE_INT_DEF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvoLangParserTYPE_STRING_DEF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(332)
			p.Match(EvoLangParserTYPE_STRING_DEF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvoLangParserTYPE_BOOLEAN_DEF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(333)
			p.Match(EvoLangParserTYPE_BOOLEAN_DEF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvoLangParserOPEN_S_BRACKET:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(334)
			p.ArrayDef()
		}

	case EvoLangParserID:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(335)
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
		p.SetState(338)
		p.Match(EvoLangParserOPEN_S_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(339)
		p.TypeDef()
	}
	{
		p.SetState(340)
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
		p.SetState(342)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserCOMMA {
		{
			p.SetState(343)
			p.Match(EvoLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(344)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(349)
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

type FunctionCallContext struct {
	ExpressionContext
	functionId antlr.Token
	args       IExprListContext
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionCallContext) GetFunctionId() antlr.Token { return s.functionId }

func (s *FunctionCallContext) SetFunctionId(v antlr.Token) { s.functionId = v }

func (s *FunctionCallContext) GetArgs() IExprListContext { return s.args }

func (s *FunctionCallContext) SetArgs(v IExprListContext) { s.args = v }

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserOPEN_PAREN, 0)
}

func (s *FunctionCallContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(EvoLangParserCLOSE_PAREN, 0)
}

func (s *FunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
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
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(351)
			p.Match(EvoLangParserOPEN_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(352)

			var _x = p.expression(0)

			localctx.(*ParenExprContext).inner = _x
		}
		{
			p.SetState(353)
			p.Match(EvoLangParserCLOSE_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFunctionCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(355)

			var _m = p.Match(EvoLangParserID)

			localctx.(*FunctionCallContext).functionId = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Match(EvoLangParserOPEN_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&283536561012736) != 0 {
			{
				p.SetState(357)

				var _x = p.ExprList()

				localctx.(*FunctionCallContext).args = _x
			}

		}
		{
			p.SetState(360)
			p.Match(EvoLangParserCLOSE_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(361)

			var _m = p.Match(EvoLangParserID)

			localctx.(*IdExprContext).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewIntLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(362)

			var _m = p.Match(EvoLangParserINT)

			localctx.(*IntLiteralContext).intVal = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewStringLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(363)

			var _m = p.Match(EvoLangParserSTRING)

			localctx.(*StringLiteralContext).strVal = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewBooleanLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(364)

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
	p.SetState(381)
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
			p.SetState(379)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*MulDivExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvoLangParserRULE_expression)
				p.SetState(367)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(368)

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
					p.SetState(369)

					var _x = p.expression(11)

					localctx.(*MulDivExprContext).right = _x
				}

			case 2:
				localctx = NewAddSubExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*AddSubExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvoLangParserRULE_expression)
				p.SetState(370)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(371)

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
					p.SetState(372)

					var _x = p.expression(10)

					localctx.(*AddSubExprContext).right = _x
				}

			case 3:
				localctx = NewRelationalExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*RelationalExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvoLangParserRULE_expression)
				p.SetState(373)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(374)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16911433728) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(375)

					var _x = p.expression(9)

					localctx.(*RelationalExprContext).right = _x
				}

			case 4:
				localctx = NewLogicalExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*LogicalExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvoLangParserRULE_expression)
				p.SetState(376)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(377)

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
					p.SetState(378)

					var _x = p.expression(8)

					localctx.(*LogicalExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(383)
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
	p.EnterRule(localctx, 84, EvoLangParserRULE_exprList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(384)
		p.expression(0)
	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserCOMMA {
		{
			p.SetState(385)
			p.Match(EvoLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
			p.expression(0)
		}

		p.SetState(391)
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
	p.EnterRule(localctx, 86, EvoLangParserRULE_mainFunction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(392)
		p.Match(EvoLangParserMAIN_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(393)
		p.Match(EvoLangParserOPEN_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(394)
		p.MainStatements()
	}
	{
		p.SetState(395)
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
	p.EnterRule(localctx, 88, EvoLangParserRULE_mainStatements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EvoLangParserPRINT_DEF || _la == EvoLangParserID {
		{
			p.SetState(397)
			p.MainStatement()
		}

		p.SetState(400)
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
	p.EnterRule(localctx, 90, EvoLangParserRULE_mainStatement)
	p.SetState(404)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EvoLangParserPRINT_DEF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(402)
			p.PrintStatement()
		}

	case EvoLangParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(403)
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
	p.EnterRule(localctx, 92, EvoLangParserRULE_printStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		p.Match(EvoLangParserPRINT_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(407)
		p.Match(EvoLangParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(408)
		p.Match(EvoLangParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(409)
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
	p.EnterRule(localctx, 94, EvoLangParserRULE_serverStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(411)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(412)
		p.Match(EvoLangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(413)
		p.Match(EvoLangParserLISTEN_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(414)
		p.Match(EvoLangParserOPEN_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(415)
		p.Match(EvoLangParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(416)
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
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
