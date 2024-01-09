// Code generated from ./EvoLang.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // EvoLang

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

var EvoLangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func evolangParserInit() {
	staticData := &EvoLangParserStaticData
	staticData.LiteralNames = []string{
		"", "'model'", "'{'", "'}'", "':'", "'events'", "'Event'", "'<'", "'>'",
		"'enum'", "','", "'accessControl'", "'allow'", "'.'", "'('", "')'",
		"'endpoint'", "'server'", "'body:'", "'authenticate:'", "'return'",
		"'<sandbox lang=\"go\">'", "'</sandbox>'", "'Int'", "'String'", "'Boolean'",
		"'['", "']'", "'@references'", "'if'", "'*'", "'/'", "'+'", "'-'", "'=='",
		"'!='", "'>='", "'<='", "'&&'", "'||'", "'main'", "'print'", "';'",
		"'.Listen'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "ID", "HTTP_METHOD", "STRING",
		"GO_CODE", "WS", "INT", "BOOLEAN", "OP",
	}
	staticData.RuleNames = []string{
		"prog", "statement", "modelDef", "modelBody", "fieldDef", "eventDef",
		"eventBody", "enumDef", "accessControlDef", "accessRule", "action",
		"paramList", "endpointDef", "endpointBody", "serverDef", "serverBody",
		"bodyDef", "authenticateDef", "returnDef", "sandbox", "goCode", "typeDef",
		"arrayDef", "refDef", "eventType", "ifCondition", "expression", "exprList",
		"mainFunction", "mainStatements", "mainStatement", "printStatement",
		"serverStatement",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 51, 295, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 1, 0, 4, 0, 68, 8, 0, 11, 0, 12, 0, 69, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 3, 1, 78, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 4, 3, 88, 8, 3, 11, 3, 12, 3, 89, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 4, 6, 111, 8, 6, 11, 6, 12, 6, 112, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 5, 7, 121, 8, 7, 10, 7, 12, 7, 124, 9, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 4, 8, 131, 8, 8, 11, 8, 12, 8, 132, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 5, 11, 151, 8, 11, 10, 11, 12, 11, 154, 9, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 5, 13, 165, 8, 13, 10, 13,
		12, 13, 168, 9, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 5,
		15, 177, 8, 15, 10, 15, 12, 15, 180, 9, 15, 1, 16, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 202, 8, 21, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 24, 5, 24, 218, 8, 24, 10, 24, 12, 24, 221, 9, 24, 1, 25, 1, 25,
		1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 234,
		8, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 241, 8, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		5, 26, 255, 8, 26, 10, 26, 12, 26, 258, 9, 26, 1, 27, 1, 27, 1, 27, 5,
		27, 263, 8, 27, 10, 27, 12, 27, 266, 9, 27, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 29, 4, 29, 274, 8, 29, 11, 29, 12, 29, 275, 1, 30, 1, 30, 3,
		30, 280, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 0, 1, 52, 33, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 0, 4, 1, 0, 30, 31, 1, 0, 32, 33,
		2, 0, 7, 8, 34, 37, 1, 0, 38, 39, 295, 0, 67, 1, 0, 0, 0, 2, 77, 1, 0,
		0, 0, 4, 79, 1, 0, 0, 0, 6, 87, 1, 0, 0, 0, 8, 91, 1, 0, 0, 0, 10, 95,
		1, 0, 0, 0, 12, 110, 1, 0, 0, 0, 14, 114, 1, 0, 0, 0, 16, 127, 1, 0, 0,
		0, 18, 136, 1, 0, 0, 0, 20, 145, 1, 0, 0, 0, 22, 147, 1, 0, 0, 0, 24, 155,
		1, 0, 0, 0, 26, 166, 1, 0, 0, 0, 28, 169, 1, 0, 0, 0, 30, 178, 1, 0, 0,
		0, 32, 181, 1, 0, 0, 0, 34, 184, 1, 0, 0, 0, 36, 187, 1, 0, 0, 0, 38, 190,
		1, 0, 0, 0, 40, 194, 1, 0, 0, 0, 42, 201, 1, 0, 0, 0, 44, 203, 1, 0, 0,
		0, 46, 207, 1, 0, 0, 0, 48, 214, 1, 0, 0, 0, 50, 222, 1, 0, 0, 0, 52, 240,
		1, 0, 0, 0, 54, 259, 1, 0, 0, 0, 56, 267, 1, 0, 0, 0, 58, 273, 1, 0, 0,
		0, 60, 279, 1, 0, 0, 0, 62, 281, 1, 0, 0, 0, 64, 287, 1, 0, 0, 0, 66, 68,
		3, 2, 1, 0, 67, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0,
		69, 70, 1, 0, 0, 0, 70, 1, 1, 0, 0, 0, 71, 78, 3, 4, 2, 0, 72, 78, 3, 14,
		7, 0, 73, 78, 3, 16, 8, 0, 74, 78, 3, 24, 12, 0, 75, 78, 3, 28, 14, 0,
		76, 78, 3, 56, 28, 0, 77, 71, 1, 0, 0, 0, 77, 72, 1, 0, 0, 0, 77, 73, 1,
		0, 0, 0, 77, 74, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 76, 1, 0, 0, 0, 78,
		3, 1, 0, 0, 0, 79, 80, 5, 1, 0, 0, 80, 81, 5, 44, 0, 0, 81, 82, 5, 2, 0,
		0, 82, 83, 3, 6, 3, 0, 83, 84, 5, 3, 0, 0, 84, 5, 1, 0, 0, 0, 85, 88, 3,
		8, 4, 0, 86, 88, 3, 10, 5, 0, 87, 85, 1, 0, 0, 0, 87, 86, 1, 0, 0, 0, 88,
		89, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 7, 1, 0, 0,
		0, 91, 92, 5, 44, 0, 0, 92, 93, 5, 4, 0, 0, 93, 94, 3, 42, 21, 0, 94, 9,
		1, 0, 0, 0, 95, 96, 5, 5, 0, 0, 96, 97, 5, 2, 0, 0, 97, 98, 3, 12, 6, 0,
		98, 99, 5, 3, 0, 0, 99, 11, 1, 0, 0, 0, 100, 101, 5, 44, 0, 0, 101, 102,
		5, 4, 0, 0, 102, 103, 5, 6, 0, 0, 103, 104, 5, 7, 0, 0, 104, 105, 3, 48,
		24, 0, 105, 106, 5, 8, 0, 0, 106, 107, 5, 2, 0, 0, 107, 108, 3, 38, 19,
		0, 108, 109, 5, 3, 0, 0, 109, 111, 1, 0, 0, 0, 110, 100, 1, 0, 0, 0, 111,
		112, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 13, 1,
		0, 0, 0, 114, 115, 5, 9, 0, 0, 115, 116, 5, 44, 0, 0, 116, 117, 5, 2, 0,
		0, 117, 122, 5, 44, 0, 0, 118, 119, 5, 10, 0, 0, 119, 121, 5, 44, 0, 0,
		120, 118, 1, 0, 0, 0, 121, 124, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122,
		123, 1, 0, 0, 0, 123, 125, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 125, 126,
		5, 3, 0, 0, 126, 15, 1, 0, 0, 0, 127, 128, 5, 11, 0, 0, 128, 130, 5, 2,
		0, 0, 129, 131, 3, 18, 9, 0, 130, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0,
		132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134,
		135, 5, 3, 0, 0, 135, 17, 1, 0, 0, 0, 136, 137, 5, 12, 0, 0, 137, 138,
		5, 44, 0, 0, 138, 139, 5, 13, 0, 0, 139, 140, 3, 20, 10, 0, 140, 141, 5,
		14, 0, 0, 141, 142, 3, 22, 11, 0, 142, 143, 5, 15, 0, 0, 143, 144, 3, 50,
		25, 0, 144, 19, 1, 0, 0, 0, 145, 146, 5, 44, 0, 0, 146, 21, 1, 0, 0, 0,
		147, 152, 5, 44, 0, 0, 148, 149, 5, 10, 0, 0, 149, 151, 5, 44, 0, 0, 150,
		148, 1, 0, 0, 0, 151, 154, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152, 153,
		1, 0, 0, 0, 153, 23, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 155, 156, 5, 16,
		0, 0, 156, 157, 5, 46, 0, 0, 157, 158, 5, 2, 0, 0, 158, 159, 3, 26, 13,
		0, 159, 160, 5, 3, 0, 0, 160, 25, 1, 0, 0, 0, 161, 165, 3, 32, 16, 0, 162,
		165, 3, 34, 17, 0, 163, 165, 3, 36, 18, 0, 164, 161, 1, 0, 0, 0, 164, 162,
		1, 0, 0, 0, 164, 163, 1, 0, 0, 0, 165, 168, 1, 0, 0, 0, 166, 164, 1, 0,
		0, 0, 166, 167, 1, 0, 0, 0, 167, 27, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0,
		169, 170, 5, 17, 0, 0, 170, 171, 5, 44, 0, 0, 171, 172, 5, 2, 0, 0, 172,
		173, 3, 30, 15, 0, 173, 174, 5, 3, 0, 0, 174, 29, 1, 0, 0, 0, 175, 177,
		3, 24, 12, 0, 176, 175, 1, 0, 0, 0, 177, 180, 1, 0, 0, 0, 178, 176, 1,
		0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 31, 1, 0, 0, 0, 180, 178, 1, 0, 0,
		0, 181, 182, 5, 18, 0, 0, 182, 183, 5, 44, 0, 0, 183, 33, 1, 0, 0, 0, 184,
		185, 5, 19, 0, 0, 185, 186, 5, 44, 0, 0, 186, 35, 1, 0, 0, 0, 187, 188,
		5, 20, 0, 0, 188, 189, 3, 52, 26, 0, 189, 37, 1, 0, 0, 0, 190, 191, 5,
		21, 0, 0, 191, 192, 3, 40, 20, 0, 192, 193, 5, 22, 0, 0, 193, 39, 1, 0,
		0, 0, 194, 195, 5, 47, 0, 0, 195, 41, 1, 0, 0, 0, 196, 202, 5, 23, 0, 0,
		197, 202, 5, 24, 0, 0, 198, 202, 5, 25, 0, 0, 199, 202, 3, 44, 22, 0, 200,
		202, 3, 46, 23, 0, 201, 196, 1, 0, 0, 0, 201, 197, 1, 0, 0, 0, 201, 198,
		1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 201, 200, 1, 0, 0, 0, 202, 43, 1, 0,
		0, 0, 203, 204, 5, 26, 0, 0, 204, 205, 3, 42, 21, 0, 205, 206, 5, 27, 0,
		0, 206, 45, 1, 0, 0, 0, 207, 208, 5, 28, 0, 0, 208, 209, 5, 14, 0, 0, 209,
		210, 5, 44, 0, 0, 210, 211, 5, 13, 0, 0, 211, 212, 5, 44, 0, 0, 212, 213,
		5, 15, 0, 0, 213, 47, 1, 0, 0, 0, 214, 219, 5, 44, 0, 0, 215, 216, 5, 10,
		0, 0, 216, 218, 5, 44, 0, 0, 217, 215, 1, 0, 0, 0, 218, 221, 1, 0, 0, 0,
		219, 217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 49, 1, 0, 0, 0, 221, 219,
		1, 0, 0, 0, 222, 223, 5, 29, 0, 0, 223, 224, 3, 52, 26, 0, 224, 51, 1,
		0, 0, 0, 225, 226, 6, 26, -1, 0, 226, 227, 5, 14, 0, 0, 227, 228, 3, 52,
		26, 0, 228, 229, 5, 15, 0, 0, 229, 241, 1, 0, 0, 0, 230, 231, 5, 44, 0,
		0, 231, 233, 5, 14, 0, 0, 232, 234, 3, 54, 27, 0, 233, 232, 1, 0, 0, 0,
		233, 234, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 241, 5, 15, 0, 0, 236,
		241, 5, 44, 0, 0, 237, 241, 5, 49, 0, 0, 238, 241, 5, 46, 0, 0, 239, 241,
		5, 50, 0, 0, 240, 225, 1, 0, 0, 0, 240, 230, 1, 0, 0, 0, 240, 236, 1, 0,
		0, 0, 240, 237, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 240, 239, 1, 0, 0, 0,
		241, 256, 1, 0, 0, 0, 242, 243, 10, 10, 0, 0, 243, 244, 7, 0, 0, 0, 244,
		255, 3, 52, 26, 11, 245, 246, 10, 9, 0, 0, 246, 247, 7, 1, 0, 0, 247, 255,
		3, 52, 26, 10, 248, 249, 10, 8, 0, 0, 249, 250, 7, 2, 0, 0, 250, 255, 3,
		52, 26, 9, 251, 252, 10, 7, 0, 0, 252, 253, 7, 3, 0, 0, 253, 255, 3, 52,
		26, 8, 254, 242, 1, 0, 0, 0, 254, 245, 1, 0, 0, 0, 254, 248, 1, 0, 0, 0,
		254, 251, 1, 0, 0, 0, 255, 258, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0, 256,
		257, 1, 0, 0, 0, 257, 53, 1, 0, 0, 0, 258, 256, 1, 0, 0, 0, 259, 264, 3,
		52, 26, 0, 260, 261, 5, 10, 0, 0, 261, 263, 3, 52, 26, 0, 262, 260, 1,
		0, 0, 0, 263, 266, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 264, 265, 1, 0, 0,
		0, 265, 55, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 267, 268, 5, 40, 0, 0, 268,
		269, 5, 2, 0, 0, 269, 270, 3, 58, 29, 0, 270, 271, 5, 3, 0, 0, 271, 57,
		1, 0, 0, 0, 272, 274, 3, 60, 30, 0, 273, 272, 1, 0, 0, 0, 274, 275, 1,
		0, 0, 0, 275, 273, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 59, 1, 0, 0,
		0, 277, 280, 3, 62, 31, 0, 278, 280, 3, 64, 32, 0, 279, 277, 1, 0, 0, 0,
		279, 278, 1, 0, 0, 0, 280, 61, 1, 0, 0, 0, 281, 282, 5, 41, 0, 0, 282,
		283, 5, 14, 0, 0, 283, 284, 5, 46, 0, 0, 284, 285, 5, 15, 0, 0, 285, 286,
		5, 42, 0, 0, 286, 63, 1, 0, 0, 0, 287, 288, 5, 44, 0, 0, 288, 289, 5, 43,
		0, 0, 289, 290, 5, 14, 0, 0, 290, 291, 5, 49, 0, 0, 291, 292, 5, 15, 0,
		0, 292, 293, 5, 42, 0, 0, 293, 65, 1, 0, 0, 0, 20, 69, 77, 87, 89, 112,
		122, 132, 152, 164, 166, 178, 201, 219, 233, 240, 254, 256, 264, 275, 279,
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
	staticData := &EvoLangParserStaticData
	staticData.once.Do(evolangParserInit)
}

// NewEvoLangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEvoLangParser(input antlr.TokenStream) *EvoLangParser {
	EvoLangParserInit()
	this := new(EvoLangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &EvoLangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "EvoLang.g4"

	return this
}

// EvoLangParser tokens.
const (
	EvoLangParserEOF         = antlr.TokenEOF
	EvoLangParserT__0        = 1
	EvoLangParserT__1        = 2
	EvoLangParserT__2        = 3
	EvoLangParserT__3        = 4
	EvoLangParserT__4        = 5
	EvoLangParserT__5        = 6
	EvoLangParserT__6        = 7
	EvoLangParserT__7        = 8
	EvoLangParserT__8        = 9
	EvoLangParserT__9        = 10
	EvoLangParserT__10       = 11
	EvoLangParserT__11       = 12
	EvoLangParserT__12       = 13
	EvoLangParserT__13       = 14
	EvoLangParserT__14       = 15
	EvoLangParserT__15       = 16
	EvoLangParserT__16       = 17
	EvoLangParserT__17       = 18
	EvoLangParserT__18       = 19
	EvoLangParserT__19       = 20
	EvoLangParserT__20       = 21
	EvoLangParserT__21       = 22
	EvoLangParserT__22       = 23
	EvoLangParserT__23       = 24
	EvoLangParserT__24       = 25
	EvoLangParserT__25       = 26
	EvoLangParserT__26       = 27
	EvoLangParserT__27       = 28
	EvoLangParserT__28       = 29
	EvoLangParserT__29       = 30
	EvoLangParserT__30       = 31
	EvoLangParserT__31       = 32
	EvoLangParserT__32       = 33
	EvoLangParserT__33       = 34
	EvoLangParserT__34       = 35
	EvoLangParserT__35       = 36
	EvoLangParserT__36       = 37
	EvoLangParserT__37       = 38
	EvoLangParserT__38       = 39
	EvoLangParserT__39       = 40
	EvoLangParserT__40       = 41
	EvoLangParserT__41       = 42
	EvoLangParserT__42       = 43
	EvoLangParserID          = 44
	EvoLangParserHTTP_METHOD = 45
	EvoLangParserSTRING      = 46
	EvoLangParserGO_CODE     = 47
	EvoLangParserWS          = 48
	EvoLangParserINT         = 49
	EvoLangParserBOOLEAN     = 50
	EvoLangParserOP          = 51
)

// EvoLangParser rules.
const (
	EvoLangParserRULE_prog             = 0
	EvoLangParserRULE_statement        = 1
	EvoLangParserRULE_modelDef         = 2
	EvoLangParserRULE_modelBody        = 3
	EvoLangParserRULE_fieldDef         = 4
	EvoLangParserRULE_eventDef         = 5
	EvoLangParserRULE_eventBody        = 6
	EvoLangParserRULE_enumDef          = 7
	EvoLangParserRULE_accessControlDef = 8
	EvoLangParserRULE_accessRule       = 9
	EvoLangParserRULE_action           = 10
	EvoLangParserRULE_paramList        = 11
	EvoLangParserRULE_endpointDef      = 12
	EvoLangParserRULE_endpointBody     = 13
	EvoLangParserRULE_serverDef        = 14
	EvoLangParserRULE_serverBody       = 15
	EvoLangParserRULE_bodyDef          = 16
	EvoLangParserRULE_authenticateDef  = 17
	EvoLangParserRULE_returnDef        = 18
	EvoLangParserRULE_sandbox          = 19
	EvoLangParserRULE_goCode           = 20
	EvoLangParserRULE_typeDef          = 21
	EvoLangParserRULE_arrayDef         = 22
	EvoLangParserRULE_refDef           = 23
	EvoLangParserRULE_eventType        = 24
	EvoLangParserRULE_ifCondition      = 25
	EvoLangParserRULE_expression       = 26
	EvoLangParserRULE_exprList         = 27
	EvoLangParserRULE_mainFunction     = 28
	EvoLangParserRULE_mainStatements   = 29
	EvoLangParserRULE_mainStatement    = 30
	EvoLangParserRULE_printStatement   = 31
	EvoLangParserRULE_serverStatement  = 32
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
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
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
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
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1099511826946) != 0) {
		{
			p.SetState(66)
			p.Statement()
		}

		p.SetState(69)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ModelDef() IModelDefContext
	EnumDef() IEnumDefContext
	AccessControlDef() IAccessControlDefContext
	EndpointDef() IEndpointDefContext
	ServerDef() IServerDefContext
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

func (s *StatementContext) AccessControlDef() IAccessControlDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessControlDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessControlDefContext)
}

func (s *StatementContext) EndpointDef() IEndpointDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndpointDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEndpointDefContext)
}

func (s *StatementContext) ServerDef() IServerDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServerDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServerDefContext)
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
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EvoLangParserRULE_statement)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EvoLangParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.ModelDef()
		}

	case EvoLangParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.EnumDef()
		}

	case EvoLangParserT__10:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(73)
			p.AccessControlDef()
		}

	case EvoLangParserT__15:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(74)
			p.EndpointDef()
		}

	case EvoLangParserT__16:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(75)
			p.ServerDef()
		}

	case EvoLangParserT__39:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(76)
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
	ID() antlr.TerminalNode
	ModelBody() IModelBodyContext

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

func (s *ModelDefContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
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

func (s *ModelDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModelDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterModelDef(s)
	}
}

func (s *ModelDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitModelDef(s)
	}
}

func (s *ModelDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
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
		p.SetState(79)
		p.Match(EvoLangParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Match(EvoLangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.ModelBody()
	}
	{
		p.SetState(83)
		p.Match(EvoLangParserT__2)
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
	AllFieldDef() []IFieldDefContext
	FieldDef(i int) IFieldDefContext
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

func (s *ModelBodyContext) AllFieldDef() []IFieldDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldDefContext); ok {
			len++
		}
	}

	tst := make([]IFieldDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldDefContext); ok {
			tst[i] = t.(IFieldDefContext)
			i++
		}
	}

	return tst
}

func (s *ModelBodyContext) FieldDef(i int) IFieldDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldDefContext); ok {
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

	return t.(IFieldDefContext)
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
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterModelBody(s)
	}
}

func (s *ModelBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitModelBody(s)
	}
}

func (s *ModelBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
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
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EvoLangParserT__4 || _la == EvoLangParserID {
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case EvoLangParserID:
			{
				p.SetState(85)
				p.FieldDef()
			}

		case EvoLangParserT__4:
			{
				p.SetState(86)
				p.EventDef()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(89)
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

// IFieldDefContext is an interface to support dynamic dispatch.
type IFieldDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	TypeDef() ITypeDefContext

	// IsFieldDefContext differentiates from other interfaces.
	IsFieldDefContext()
}

type FieldDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDefContext() *FieldDefContext {
	var p = new(FieldDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_fieldDef
	return p
}

func InitEmptyFieldDefContext(p *FieldDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_fieldDef
}

func (*FieldDefContext) IsFieldDefContext() {}

func NewFieldDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDefContext {
	var p = new(FieldDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_fieldDef

	return p
}

func (s *FieldDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDefContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *FieldDefContext) TypeDef() ITypeDefContext {
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

func (s *FieldDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterFieldDef(s)
	}
}

func (s *FieldDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitFieldDef(s)
	}
}

func (s *FieldDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitFieldDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) FieldDef() (localctx IFieldDefContext) {
	localctx = NewFieldDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EvoLangParserRULE_fieldDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
		p.Match(EvoLangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.TypeDef()
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
	EventBody() IEventBodyContext

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

func (s *EventDefContext) EventBody() IEventBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEventBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEventBodyContext)
}

func (s *EventDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EventDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterEventDef(s)
	}
}

func (s *EventDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitEventDef(s)
	}
}

func (s *EventDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitEventDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) EventDef() (localctx IEventDefContext) {
	localctx = NewEventDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EvoLangParserRULE_eventDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(EvoLangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.Match(EvoLangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)
		p.EventBody()
	}
	{
		p.SetState(98)
		p.Match(EvoLangParserT__2)
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

// IEventBodyContext is an interface to support dynamic dispatch.
type IEventBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllEventType() []IEventTypeContext
	EventType(i int) IEventTypeContext
	AllSandbox() []ISandboxContext
	Sandbox(i int) ISandboxContext

	// IsEventBodyContext differentiates from other interfaces.
	IsEventBodyContext()
}

type EventBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEventBodyContext() *EventBodyContext {
	var p = new(EventBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_eventBody
	return p
}

func InitEmptyEventBodyContext(p *EventBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_eventBody
}

func (*EventBodyContext) IsEventBodyContext() {}

func NewEventBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EventBodyContext {
	var p = new(EventBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_eventBody

	return p
}

func (s *EventBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EventBodyContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserID)
}

func (s *EventBodyContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, i)
}

func (s *EventBodyContext) AllEventType() []IEventTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEventTypeContext); ok {
			len++
		}
	}

	tst := make([]IEventTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEventTypeContext); ok {
			tst[i] = t.(IEventTypeContext)
			i++
		}
	}

	return tst
}

func (s *EventBodyContext) EventType(i int) IEventTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEventTypeContext); ok {
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

	return t.(IEventTypeContext)
}

func (s *EventBodyContext) AllSandbox() []ISandboxContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISandboxContext); ok {
			len++
		}
	}

	tst := make([]ISandboxContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISandboxContext); ok {
			tst[i] = t.(ISandboxContext)
			i++
		}
	}

	return tst
}

func (s *EventBodyContext) Sandbox(i int) ISandboxContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISandboxContext); ok {
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

	return t.(ISandboxContext)
}

func (s *EventBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EventBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterEventBody(s)
	}
}

func (s *EventBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitEventBody(s)
	}
}

func (s *EventBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitEventBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) EventBody() (localctx IEventBodyContext) {
	localctx = NewEventBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, EvoLangParserRULE_eventBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EvoLangParserID {
		{
			p.SetState(100)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.Match(EvoLangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.Match(EvoLangParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(EvoLangParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.EventType()
		}
		{
			p.SetState(105)
			p.Match(EvoLangParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(EvoLangParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Sandbox()
		}
		{
			p.SetState(108)
			p.Match(EvoLangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(112)
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

// IEnumDefContext is an interface to support dynamic dispatch.
type IEnumDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

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

func (s *EnumDefContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserID)
}

func (s *EnumDefContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, i)
}

func (s *EnumDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterEnumDef(s)
	}
}

func (s *EnumDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitEnumDef(s)
	}
}

func (s *EnumDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitEnumDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) EnumDef() (localctx IEnumDefContext) {
	localctx = NewEnumDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, EvoLangParserRULE_enumDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(EvoLangParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Match(EvoLangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserT__9 {
		{
			p.SetState(118)
			p.Match(EvoLangParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(125)
		p.Match(EvoLangParserT__2)
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

// IAccessControlDefContext is an interface to support dynamic dispatch.
type IAccessControlDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAccessRule() []IAccessRuleContext
	AccessRule(i int) IAccessRuleContext

	// IsAccessControlDefContext differentiates from other interfaces.
	IsAccessControlDefContext()
}

type AccessControlDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessControlDefContext() *AccessControlDefContext {
	var p = new(AccessControlDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_accessControlDef
	return p
}

func InitEmptyAccessControlDefContext(p *AccessControlDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_accessControlDef
}

func (*AccessControlDefContext) IsAccessControlDefContext() {}

func NewAccessControlDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessControlDefContext {
	var p = new(AccessControlDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_accessControlDef

	return p
}

func (s *AccessControlDefContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessControlDefContext) AllAccessRule() []IAccessRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAccessRuleContext); ok {
			len++
		}
	}

	tst := make([]IAccessRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAccessRuleContext); ok {
			tst[i] = t.(IAccessRuleContext)
			i++
		}
	}

	return tst
}

func (s *AccessControlDefContext) AccessRule(i int) IAccessRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessRuleContext); ok {
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

	return t.(IAccessRuleContext)
}

func (s *AccessControlDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessControlDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessControlDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterAccessControlDef(s)
	}
}

func (s *AccessControlDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitAccessControlDef(s)
	}
}

func (s *AccessControlDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitAccessControlDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) AccessControlDef() (localctx IAccessControlDefContext) {
	localctx = NewAccessControlDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, EvoLangParserRULE_accessControlDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(EvoLangParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Match(EvoLangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EvoLangParserT__11 {
		{
			p.SetState(129)
			p.AccessRule()
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(134)
		p.Match(EvoLangParserT__2)
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

// IAccessRuleContext is an interface to support dynamic dispatch.
type IAccessRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Action_() IActionContext
	ParamList() IParamListContext
	IfCondition() IIfConditionContext

	// IsAccessRuleContext differentiates from other interfaces.
	IsAccessRuleContext()
}

type AccessRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessRuleContext() *AccessRuleContext {
	var p = new(AccessRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_accessRule
	return p
}

func InitEmptyAccessRuleContext(p *AccessRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_accessRule
}

func (*AccessRuleContext) IsAccessRuleContext() {}

func NewAccessRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessRuleContext {
	var p = new(AccessRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_accessRule

	return p
}

func (s *AccessRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessRuleContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *AccessRuleContext) Action_() IActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionContext)
}

func (s *AccessRuleContext) ParamList() IParamListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *AccessRuleContext) IfCondition() IIfConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfConditionContext)
}

func (s *AccessRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterAccessRule(s)
	}
}

func (s *AccessRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitAccessRule(s)
	}
}

func (s *AccessRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitAccessRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) AccessRule() (localctx IAccessRuleContext) {
	localctx = NewAccessRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, EvoLangParserRULE_accessRule)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(EvoLangParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Match(EvoLangParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Action_()
	}
	{
		p.SetState(140)
		p.Match(EvoLangParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.ParamList()
	}
	{
		p.SetState(142)
		p.Match(EvoLangParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.IfCondition()
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
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterAction(s)
	}
}

func (s *ActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitAction(s)
	}
}

func (s *ActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) Action_() (localctx IActionContext) {
	localctx = NewActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, EvoLangParserRULE_action)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
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

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (s *ParamListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitParamList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) ParamList() (localctx IParamListContext) {
	localctx = NewParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, EvoLangParserRULE_paramList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserT__9 {
		{
			p.SetState(148)
			p.Match(EvoLangParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Match(EvoLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(154)
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

// IEndpointDefContext is an interface to support dynamic dispatch.
type IEndpointDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	EndpointBody() IEndpointBodyContext

	// IsEndpointDefContext differentiates from other interfaces.
	IsEndpointDefContext()
}

type EndpointDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndpointDefContext() *EndpointDefContext {
	var p = new(EndpointDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_endpointDef
	return p
}

func InitEmptyEndpointDefContext(p *EndpointDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_endpointDef
}

func (*EndpointDefContext) IsEndpointDefContext() {}

func NewEndpointDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndpointDefContext {
	var p = new(EndpointDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_endpointDef

	return p
}

func (s *EndpointDefContext) GetParser() antlr.Parser { return s.parser }

func (s *EndpointDefContext) STRING() antlr.TerminalNode {
	return s.GetToken(EvoLangParserSTRING, 0)
}

func (s *EndpointDefContext) EndpointBody() IEndpointBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndpointBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEndpointBodyContext)
}

func (s *EndpointDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndpointDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndpointDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterEndpointDef(s)
	}
}

func (s *EndpointDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitEndpointDef(s)
	}
}

func (s *EndpointDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitEndpointDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) EndpointDef() (localctx IEndpointDefContext) {
	localctx = NewEndpointDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, EvoLangParserRULE_endpointDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(EvoLangParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(EvoLangParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Match(EvoLangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.EndpointBody()
	}
	{
		p.SetState(159)
		p.Match(EvoLangParserT__2)
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

// IEndpointBodyContext is an interface to support dynamic dispatch.
type IEndpointBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBodyDef() []IBodyDefContext
	BodyDef(i int) IBodyDefContext
	AllAuthenticateDef() []IAuthenticateDefContext
	AuthenticateDef(i int) IAuthenticateDefContext
	AllReturnDef() []IReturnDefContext
	ReturnDef(i int) IReturnDefContext

	// IsEndpointBodyContext differentiates from other interfaces.
	IsEndpointBodyContext()
}

type EndpointBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndpointBodyContext() *EndpointBodyContext {
	var p = new(EndpointBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_endpointBody
	return p
}

func InitEmptyEndpointBodyContext(p *EndpointBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_endpointBody
}

func (*EndpointBodyContext) IsEndpointBodyContext() {}

func NewEndpointBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndpointBodyContext {
	var p = new(EndpointBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_endpointBody

	return p
}

func (s *EndpointBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EndpointBodyContext) AllBodyDef() []IBodyDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBodyDefContext); ok {
			len++
		}
	}

	tst := make([]IBodyDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBodyDefContext); ok {
			tst[i] = t.(IBodyDefContext)
			i++
		}
	}

	return tst
}

func (s *EndpointBodyContext) BodyDef(i int) IBodyDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyDefContext); ok {
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

	return t.(IBodyDefContext)
}

func (s *EndpointBodyContext) AllAuthenticateDef() []IAuthenticateDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAuthenticateDefContext); ok {
			len++
		}
	}

	tst := make([]IAuthenticateDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAuthenticateDefContext); ok {
			tst[i] = t.(IAuthenticateDefContext)
			i++
		}
	}

	return tst
}

func (s *EndpointBodyContext) AuthenticateDef(i int) IAuthenticateDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuthenticateDefContext); ok {
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

	return t.(IAuthenticateDefContext)
}

func (s *EndpointBodyContext) AllReturnDef() []IReturnDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IReturnDefContext); ok {
			len++
		}
	}

	tst := make([]IReturnDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IReturnDefContext); ok {
			tst[i] = t.(IReturnDefContext)
			i++
		}
	}

	return tst
}

func (s *EndpointBodyContext) ReturnDef(i int) IReturnDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnDefContext); ok {
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

	return t.(IReturnDefContext)
}

func (s *EndpointBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndpointBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndpointBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterEndpointBody(s)
	}
}

func (s *EndpointBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitEndpointBody(s)
	}
}

func (s *EndpointBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitEndpointBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) EndpointBody() (localctx IEndpointBodyContext) {
	localctx = NewEndpointBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, EvoLangParserRULE_endpointBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1835008) != 0 {
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case EvoLangParserT__17:
			{
				p.SetState(161)
				p.BodyDef()
			}

		case EvoLangParserT__18:
			{
				p.SetState(162)
				p.AuthenticateDef()
			}

		case EvoLangParserT__19:
			{
				p.SetState(163)
				p.ReturnDef()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(168)
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

// IServerDefContext is an interface to support dynamic dispatch.
type IServerDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ServerBody() IServerBodyContext

	// IsServerDefContext differentiates from other interfaces.
	IsServerDefContext()
}

type ServerDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServerDefContext() *ServerDefContext {
	var p = new(ServerDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_serverDef
	return p
}

func InitEmptyServerDefContext(p *ServerDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_serverDef
}

func (*ServerDefContext) IsServerDefContext() {}

func NewServerDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServerDefContext {
	var p = new(ServerDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_serverDef

	return p
}

func (s *ServerDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ServerDefContext) ID() antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, 0)
}

func (s *ServerDefContext) ServerBody() IServerBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServerBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServerBodyContext)
}

func (s *ServerDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServerDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServerDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterServerDef(s)
	}
}

func (s *ServerDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitServerDef(s)
	}
}

func (s *ServerDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitServerDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) ServerDef() (localctx IServerDefContext) {
	localctx = NewServerDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, EvoLangParserRULE_serverDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(EvoLangParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
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
		p.Match(EvoLangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.ServerBody()
	}
	{
		p.SetState(173)
		p.Match(EvoLangParserT__2)
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

// IServerBodyContext is an interface to support dynamic dispatch.
type IServerBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEndpointDef() []IEndpointDefContext
	EndpointDef(i int) IEndpointDefContext

	// IsServerBodyContext differentiates from other interfaces.
	IsServerBodyContext()
}

type ServerBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServerBodyContext() *ServerBodyContext {
	var p = new(ServerBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_serverBody
	return p
}

func InitEmptyServerBodyContext(p *ServerBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_serverBody
}

func (*ServerBodyContext) IsServerBodyContext() {}

func NewServerBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServerBodyContext {
	var p = new(ServerBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_serverBody

	return p
}

func (s *ServerBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ServerBodyContext) AllEndpointDef() []IEndpointDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEndpointDefContext); ok {
			len++
		}
	}

	tst := make([]IEndpointDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEndpointDefContext); ok {
			tst[i] = t.(IEndpointDefContext)
			i++
		}
	}

	return tst
}

func (s *ServerBodyContext) EndpointDef(i int) IEndpointDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndpointDefContext); ok {
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

	return t.(IEndpointDefContext)
}

func (s *ServerBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServerBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServerBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterServerBody(s)
	}
}

func (s *ServerBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitServerBody(s)
	}
}

func (s *ServerBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitServerBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) ServerBody() (localctx IServerBodyContext) {
	localctx = NewServerBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, EvoLangParserRULE_serverBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserT__15 {
		{
			p.SetState(175)
			p.EndpointDef()
		}

		p.SetState(180)
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
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterBodyDef(s)
	}
}

func (s *BodyDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitBodyDef(s)
	}
}

func (s *BodyDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitBodyDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) BodyDef() (localctx IBodyDefContext) {
	localctx = NewBodyDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, EvoLangParserRULE_bodyDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(EvoLangParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
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
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterAuthenticateDef(s)
	}
}

func (s *AuthenticateDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitAuthenticateDef(s)
	}
}

func (s *AuthenticateDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitAuthenticateDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) AuthenticateDef() (localctx IAuthenticateDefContext) {
	localctx = NewAuthenticateDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, EvoLangParserRULE_authenticateDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(EvoLangParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
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
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterReturnDef(s)
	}
}

func (s *ReturnDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitReturnDef(s)
	}
}

func (s *ReturnDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitReturnDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) ReturnDef() (localctx IReturnDefContext) {
	localctx = NewReturnDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, EvoLangParserRULE_returnDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Match(EvoLangParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
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

// ISandboxContext is an interface to support dynamic dispatch.
type ISandboxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GoCode() IGoCodeContext

	// IsSandboxContext differentiates from other interfaces.
	IsSandboxContext()
}

type SandboxContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySandboxContext() *SandboxContext {
	var p = new(SandboxContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_sandbox
	return p
}

func InitEmptySandboxContext(p *SandboxContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_sandbox
}

func (*SandboxContext) IsSandboxContext() {}

func NewSandboxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SandboxContext {
	var p = new(SandboxContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_sandbox

	return p
}

func (s *SandboxContext) GetParser() antlr.Parser { return s.parser }

func (s *SandboxContext) GoCode() IGoCodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGoCodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGoCodeContext)
}

func (s *SandboxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SandboxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SandboxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterSandbox(s)
	}
}

func (s *SandboxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitSandbox(s)
	}
}

func (s *SandboxContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitSandbox(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) Sandbox() (localctx ISandboxContext) {
	localctx = NewSandboxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, EvoLangParserRULE_sandbox)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(EvoLangParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)
		p.GoCode()
	}
	{
		p.SetState(192)
		p.Match(EvoLangParserT__21)
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

// IGoCodeContext is an interface to support dynamic dispatch.
type IGoCodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GO_CODE() antlr.TerminalNode

	// IsGoCodeContext differentiates from other interfaces.
	IsGoCodeContext()
}

type GoCodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGoCodeContext() *GoCodeContext {
	var p = new(GoCodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_goCode
	return p
}

func InitEmptyGoCodeContext(p *GoCodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_goCode
}

func (*GoCodeContext) IsGoCodeContext() {}

func NewGoCodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoCodeContext {
	var p = new(GoCodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_goCode

	return p
}

func (s *GoCodeContext) GetParser() antlr.Parser { return s.parser }

func (s *GoCodeContext) GO_CODE() antlr.TerminalNode {
	return s.GetToken(EvoLangParserGO_CODE, 0)
}

func (s *GoCodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoCodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GoCodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterGoCode(s)
	}
}

func (s *GoCodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitGoCode(s)
	}
}

func (s *GoCodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitGoCode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) GoCode() (localctx IGoCodeContext) {
	localctx = NewGoCodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, EvoLangParserRULE_goCode)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.Match(EvoLangParserGO_CODE)
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

// ITypeDefContext is an interface to support dynamic dispatch.
type ITypeDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ArrayDef() IArrayDefContext
	RefDef() IRefDefContext

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

func (s *TypeDefContext) RefDef() IRefDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRefDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRefDefContext)
}

func (s *TypeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterTypeDef(s)
	}
}

func (s *TypeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitTypeDef(s)
	}
}

func (s *TypeDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitTypeDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) TypeDef() (localctx ITypeDefContext) {
	localctx = NewTypeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, EvoLangParserRULE_typeDef)
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EvoLangParserT__22:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(196)
			p.Match(EvoLangParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvoLangParserT__23:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(197)
			p.Match(EvoLangParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvoLangParserT__24:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(198)
			p.Match(EvoLangParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvoLangParserT__25:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(199)
			p.ArrayDef()
		}

	case EvoLangParserT__27:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(200)
			p.RefDef()
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
	TypeDef() ITypeDefContext

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

func (s *ArrayDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterArrayDef(s)
	}
}

func (s *ArrayDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitArrayDef(s)
	}
}

func (s *ArrayDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitArrayDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) ArrayDef() (localctx IArrayDefContext) {
	localctx = NewArrayDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, EvoLangParserRULE_arrayDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Match(EvoLangParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.TypeDef()
	}
	{
		p.SetState(205)
		p.Match(EvoLangParserT__26)
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

// IRefDefContext is an interface to support dynamic dispatch.
type IRefDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsRefDefContext differentiates from other interfaces.
	IsRefDefContext()
}

type RefDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRefDefContext() *RefDefContext {
	var p = new(RefDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_refDef
	return p
}

func InitEmptyRefDefContext(p *RefDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_refDef
}

func (*RefDefContext) IsRefDefContext() {}

func NewRefDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RefDefContext {
	var p = new(RefDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_refDef

	return p
}

func (s *RefDefContext) GetParser() antlr.Parser { return s.parser }

func (s *RefDefContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(EvoLangParserID)
}

func (s *RefDefContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(EvoLangParserID, i)
}

func (s *RefDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RefDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RefDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterRefDef(s)
	}
}

func (s *RefDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitRefDef(s)
	}
}

func (s *RefDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitRefDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) RefDef() (localctx IRefDefContext) {
	localctx = NewRefDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, EvoLangParserRULE_refDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(EvoLangParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Match(EvoLangParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		p.Match(EvoLangParserT__12)
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
		p.Match(EvoLangParserT__14)
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

func (s *EventTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EventTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterEventType(s)
	}
}

func (s *EventTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitEventType(s)
	}
}

func (s *EventTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitEventType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) EventType() (localctx IEventTypeContext) {
	localctx = NewEventTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, EvoLangParserRULE_eventType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserT__9 {
		{
			p.SetState(215)
			p.Match(EvoLangParserT__9)
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

		p.SetState(221)
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

// IIfConditionContext is an interface to support dynamic dispatch.
type IIfConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsIfConditionContext differentiates from other interfaces.
	IsIfConditionContext()
}

type IfConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfConditionContext() *IfConditionContext {
	var p = new(IfConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_ifCondition
	return p
}

func InitEmptyIfConditionContext(p *IfConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvoLangParserRULE_ifCondition
}

func (*IfConditionContext) IsIfConditionContext() {}

func NewIfConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfConditionContext {
	var p = new(IfConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvoLangParserRULE_ifCondition

	return p
}

func (s *IfConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *IfConditionContext) Expression() IExpressionContext {
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

func (s *IfConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterIfCondition(s)
	}
}

func (s *IfConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitIfCondition(s)
	}
}

func (s *IfConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitIfCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) IfCondition() (localctx IIfConditionContext) {
	localctx = NewIfConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, EvoLangParserRULE_ifCondition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Match(EvoLangParserT__28)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
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

func (s *MulDivExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterMulDivExpr(s)
	}
}

func (s *MulDivExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitMulDivExpr(s)
	}
}

func (s *MulDivExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
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
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterIdExpr(s)
	}
}

func (s *IdExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitIdExpr(s)
	}
}

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
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
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
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
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
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
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterIntLiteral(s)
	}
}

func (s *IntLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitIntLiteral(s)
	}
}

func (s *IntLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
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
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
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

func (s *RelationalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterRelationalExpr(s)
	}
}

func (s *RelationalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitRelationalExpr(s)
	}
}

func (s *RelationalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
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
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
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

func (s *AddSubExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterAddSubExpr(s)
	}
}

func (s *AddSubExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitAddSubExpr(s)
	}
}

func (s *AddSubExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
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

func (s *LogicalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterLogicalExpr(s)
	}
}

func (s *LogicalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitLogicalExpr(s)
	}
}

func (s *LogicalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
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
	_startState := 52
	p.EnterRecursionRule(localctx, 52, EvoLangParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(226)
			p.Match(EvoLangParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)

			var _x = p.expression(0)

			localctx.(*ParenExprContext).inner = _x
		}
		{
			p.SetState(228)
			p.Match(EvoLangParserT__14)
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
			p.SetState(230)

			var _m = p.Match(EvoLangParserID)

			localctx.(*FunctionCallContext).functionId = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Match(EvoLangParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1776810790502400) != 0 {
			{
				p.SetState(232)

				var _x = p.ExprList()

				localctx.(*FunctionCallContext).args = _x
			}

		}
		{
			p.SetState(235)
			p.Match(EvoLangParserT__14)
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
			p.SetState(236)

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
			p.SetState(237)

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
			p.SetState(238)

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
			p.SetState(239)

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
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(254)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*MulDivExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvoLangParserRULE_expression)
				p.SetState(242)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(243)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == EvoLangParserT__29 || _la == EvoLangParserT__30) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(244)

					var _x = p.expression(11)

					localctx.(*MulDivExprContext).right = _x
				}

			case 2:
				localctx = NewAddSubExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*AddSubExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvoLangParserRULE_expression)
				p.SetState(245)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(246)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == EvoLangParserT__31 || _la == EvoLangParserT__32) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(247)

					var _x = p.expression(10)

					localctx.(*AddSubExprContext).right = _x
				}

			case 3:
				localctx = NewRelationalExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*RelationalExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvoLangParserRULE_expression)
				p.SetState(248)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(249)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&257698038144) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(250)

					var _x = p.expression(9)

					localctx.(*RelationalExprContext).right = _x
				}

			case 4:
				localctx = NewLogicalExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*LogicalExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvoLangParserRULE_expression)
				p.SetState(251)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(252)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*LogicalExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == EvoLangParserT__37 || _la == EvoLangParserT__38) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*LogicalExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(253)

					var _x = p.expression(8)

					localctx.(*LogicalExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
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

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterExprList(s)
	}
}

func (s *ExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitExprList(s)
	}
}

func (s *ExprListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitExprList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) ExprList() (localctx IExprListContext) {
	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, EvoLangParserRULE_exprList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.expression(0)
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvoLangParserT__9 {
		{
			p.SetState(260)
			p.Match(EvoLangParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)
			p.expression(0)
		}

		p.SetState(266)
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
	MainStatements() IMainStatementsContext

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

func (s *MainFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterMainFunction(s)
	}
}

func (s *MainFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitMainFunction(s)
	}
}

func (s *MainFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitMainFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) MainFunction() (localctx IMainFunctionContext) {
	localctx = NewMainFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, EvoLangParserRULE_mainFunction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Match(EvoLangParserT__39)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.Match(EvoLangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.MainStatements()
	}
	{
		p.SetState(270)
		p.Match(EvoLangParserT__2)
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
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterMainStatements(s)
	}
}

func (s *MainStatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitMainStatements(s)
	}
}

func (s *MainStatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitMainStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) MainStatements() (localctx IMainStatementsContext) {
	localctx = NewMainStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, EvoLangParserRULE_mainStatements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EvoLangParserT__40 || _la == EvoLangParserID {
		{
			p.SetState(272)
			p.MainStatement()
		}

		p.SetState(275)
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
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterMainStatement(s)
	}
}

func (s *MainStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitMainStatement(s)
	}
}

func (s *MainStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitMainStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) MainStatement() (localctx IMainStatementContext) {
	localctx = NewMainStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, EvoLangParserRULE_mainStatement)
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EvoLangParserT__40:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(277)
			p.PrintStatement()
		}

	case EvoLangParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(278)
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
	STRING() antlr.TerminalNode

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

func (s *PrintStatementContext) STRING() antlr.TerminalNode {
	return s.GetToken(EvoLangParserSTRING, 0)
}

func (s *PrintStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterPrintStatement(s)
	}
}

func (s *PrintStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitPrintStatement(s)
	}
}

func (s *PrintStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitPrintStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) PrintStatement() (localctx IPrintStatementContext) {
	localctx = NewPrintStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, EvoLangParserRULE_printStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(EvoLangParserT__40)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(282)
		p.Match(EvoLangParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(283)
		p.Match(EvoLangParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(284)
		p.Match(EvoLangParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(285)
		p.Match(EvoLangParserT__41)
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
	INT() antlr.TerminalNode

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

func (s *ServerStatementContext) INT() antlr.TerminalNode {
	return s.GetToken(EvoLangParserINT, 0)
}

func (s *ServerStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServerStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServerStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.EnterServerStatement(s)
	}
}

func (s *ServerStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvoLangListener); ok {
		listenerT.ExitServerStatement(s)
	}
}

func (s *ServerStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvoLangVisitor:
		return t.VisitServerStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvoLangParser) ServerStatement() (localctx IServerStatementContext) {
	localctx = NewServerStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, EvoLangParserRULE_serverStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		p.Match(EvoLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(288)
		p.Match(EvoLangParserT__42)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(289)
		p.Match(EvoLangParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(290)
		p.Match(EvoLangParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.Match(EvoLangParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
		p.Match(EvoLangParserT__41)
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
	case 26:
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
