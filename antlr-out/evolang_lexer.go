// Code generated from ./EvoLang.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type EvoLangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var EvoLangLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func evolanglexerLexerInit() {
	staticData := &EvoLangLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
		"T__41", "T__42", "ID", "HTTP_METHOD", "STRING", "GO_CODE", "WS", "INT",
		"BOOLEAN", "OP",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 51, 406, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33,
		1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1,
		37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 5, 43, 327, 8, 43, 10, 43,
		12, 43, 330, 9, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 3,
		44, 339, 8, 44, 1, 45, 1, 45, 1, 45, 1, 45, 5, 45, 345, 8, 45, 10, 45,
		12, 45, 348, 9, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1,
		46, 5, 46, 358, 8, 46, 10, 46, 12, 46, 361, 9, 46, 1, 46, 1, 46, 1, 46,
		1, 46, 1, 46, 1, 46, 1, 47, 4, 47, 370, 8, 47, 11, 47, 12, 47, 371, 1,
		47, 1, 47, 1, 48, 4, 48, 377, 8, 48, 11, 48, 12, 48, 378, 1, 49, 1, 49,
		1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 3, 49, 390, 8, 49, 1,
		50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50,
		1, 50, 1, 50, 3, 50, 405, 8, 50, 1, 359, 0, 51, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
		65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41,
		83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50,
		101, 51, 1, 0, 6, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90,
		95, 95, 97, 122, 2, 0, 34, 34, 92, 92, 3, 0, 9, 10, 13, 13, 32, 32, 1,
		0, 48, 57, 5, 0, 42, 43, 45, 45, 47, 47, 60, 60, 62, 62, 419, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0,
		25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0,
		0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0,
		0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0,
		0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1,
		0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63,
		1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0,
		71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0,
		0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0,
		0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0,
		0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101,
		1, 0, 0, 0, 1, 103, 1, 0, 0, 0, 3, 109, 1, 0, 0, 0, 5, 111, 1, 0, 0, 0,
		7, 113, 1, 0, 0, 0, 9, 115, 1, 0, 0, 0, 11, 122, 1, 0, 0, 0, 13, 128, 1,
		0, 0, 0, 15, 130, 1, 0, 0, 0, 17, 132, 1, 0, 0, 0, 19, 137, 1, 0, 0, 0,
		21, 139, 1, 0, 0, 0, 23, 153, 1, 0, 0, 0, 25, 159, 1, 0, 0, 0, 27, 161,
		1, 0, 0, 0, 29, 163, 1, 0, 0, 0, 31, 165, 1, 0, 0, 0, 33, 174, 1, 0, 0,
		0, 35, 181, 1, 0, 0, 0, 37, 187, 1, 0, 0, 0, 39, 201, 1, 0, 0, 0, 41, 208,
		1, 0, 0, 0, 43, 228, 1, 0, 0, 0, 45, 239, 1, 0, 0, 0, 47, 243, 1, 0, 0,
		0, 49, 250, 1, 0, 0, 0, 51, 258, 1, 0, 0, 0, 53, 260, 1, 0, 0, 0, 55, 262,
		1, 0, 0, 0, 57, 274, 1, 0, 0, 0, 59, 277, 1, 0, 0, 0, 61, 279, 1, 0, 0,
		0, 63, 281, 1, 0, 0, 0, 65, 283, 1, 0, 0, 0, 67, 285, 1, 0, 0, 0, 69, 288,
		1, 0, 0, 0, 71, 291, 1, 0, 0, 0, 73, 294, 1, 0, 0, 0, 75, 297, 1, 0, 0,
		0, 77, 300, 1, 0, 0, 0, 79, 303, 1, 0, 0, 0, 81, 308, 1, 0, 0, 0, 83, 314,
		1, 0, 0, 0, 85, 316, 1, 0, 0, 0, 87, 324, 1, 0, 0, 0, 89, 338, 1, 0, 0,
		0, 91, 340, 1, 0, 0, 0, 93, 351, 1, 0, 0, 0, 95, 369, 1, 0, 0, 0, 97, 376,
		1, 0, 0, 0, 99, 389, 1, 0, 0, 0, 101, 404, 1, 0, 0, 0, 103, 104, 5, 109,
		0, 0, 104, 105, 5, 111, 0, 0, 105, 106, 5, 100, 0, 0, 106, 107, 5, 101,
		0, 0, 107, 108, 5, 108, 0, 0, 108, 2, 1, 0, 0, 0, 109, 110, 5, 123, 0,
		0, 110, 4, 1, 0, 0, 0, 111, 112, 5, 125, 0, 0, 112, 6, 1, 0, 0, 0, 113,
		114, 5, 58, 0, 0, 114, 8, 1, 0, 0, 0, 115, 116, 5, 101, 0, 0, 116, 117,
		5, 118, 0, 0, 117, 118, 5, 101, 0, 0, 118, 119, 5, 110, 0, 0, 119, 120,
		5, 116, 0, 0, 120, 121, 5, 115, 0, 0, 121, 10, 1, 0, 0, 0, 122, 123, 5,
		69, 0, 0, 123, 124, 5, 118, 0, 0, 124, 125, 5, 101, 0, 0, 125, 126, 5,
		110, 0, 0, 126, 127, 5, 116, 0, 0, 127, 12, 1, 0, 0, 0, 128, 129, 5, 60,
		0, 0, 129, 14, 1, 0, 0, 0, 130, 131, 5, 62, 0, 0, 131, 16, 1, 0, 0, 0,
		132, 133, 5, 101, 0, 0, 133, 134, 5, 110, 0, 0, 134, 135, 5, 117, 0, 0,
		135, 136, 5, 109, 0, 0, 136, 18, 1, 0, 0, 0, 137, 138, 5, 44, 0, 0, 138,
		20, 1, 0, 0, 0, 139, 140, 5, 97, 0, 0, 140, 141, 5, 99, 0, 0, 141, 142,
		5, 99, 0, 0, 142, 143, 5, 101, 0, 0, 143, 144, 5, 115, 0, 0, 144, 145,
		5, 115, 0, 0, 145, 146, 5, 67, 0, 0, 146, 147, 5, 111, 0, 0, 147, 148,
		5, 110, 0, 0, 148, 149, 5, 116, 0, 0, 149, 150, 5, 114, 0, 0, 150, 151,
		5, 111, 0, 0, 151, 152, 5, 108, 0, 0, 152, 22, 1, 0, 0, 0, 153, 154, 5,
		97, 0, 0, 154, 155, 5, 108, 0, 0, 155, 156, 5, 108, 0, 0, 156, 157, 5,
		111, 0, 0, 157, 158, 5, 119, 0, 0, 158, 24, 1, 0, 0, 0, 159, 160, 5, 46,
		0, 0, 160, 26, 1, 0, 0, 0, 161, 162, 5, 40, 0, 0, 162, 28, 1, 0, 0, 0,
		163, 164, 5, 41, 0, 0, 164, 30, 1, 0, 0, 0, 165, 166, 5, 101, 0, 0, 166,
		167, 5, 110, 0, 0, 167, 168, 5, 100, 0, 0, 168, 169, 5, 112, 0, 0, 169,
		170, 5, 111, 0, 0, 170, 171, 5, 105, 0, 0, 171, 172, 5, 110, 0, 0, 172,
		173, 5, 116, 0, 0, 173, 32, 1, 0, 0, 0, 174, 175, 5, 115, 0, 0, 175, 176,
		5, 101, 0, 0, 176, 177, 5, 114, 0, 0, 177, 178, 5, 118, 0, 0, 178, 179,
		5, 101, 0, 0, 179, 180, 5, 114, 0, 0, 180, 34, 1, 0, 0, 0, 181, 182, 5,
		98, 0, 0, 182, 183, 5, 111, 0, 0, 183, 184, 5, 100, 0, 0, 184, 185, 5,
		121, 0, 0, 185, 186, 5, 58, 0, 0, 186, 36, 1, 0, 0, 0, 187, 188, 5, 97,
		0, 0, 188, 189, 5, 117, 0, 0, 189, 190, 5, 116, 0, 0, 190, 191, 5, 104,
		0, 0, 191, 192, 5, 101, 0, 0, 192, 193, 5, 110, 0, 0, 193, 194, 5, 116,
		0, 0, 194, 195, 5, 105, 0, 0, 195, 196, 5, 99, 0, 0, 196, 197, 5, 97, 0,
		0, 197, 198, 5, 116, 0, 0, 198, 199, 5, 101, 0, 0, 199, 200, 5, 58, 0,
		0, 200, 38, 1, 0, 0, 0, 201, 202, 5, 114, 0, 0, 202, 203, 5, 101, 0, 0,
		203, 204, 5, 116, 0, 0, 204, 205, 5, 117, 0, 0, 205, 206, 5, 114, 0, 0,
		206, 207, 5, 110, 0, 0, 207, 40, 1, 0, 0, 0, 208, 209, 5, 60, 0, 0, 209,
		210, 5, 115, 0, 0, 210, 211, 5, 97, 0, 0, 211, 212, 5, 110, 0, 0, 212,
		213, 5, 100, 0, 0, 213, 214, 5, 98, 0, 0, 214, 215, 5, 111, 0, 0, 215,
		216, 5, 120, 0, 0, 216, 217, 5, 32, 0, 0, 217, 218, 5, 108, 0, 0, 218,
		219, 5, 97, 0, 0, 219, 220, 5, 110, 0, 0, 220, 221, 5, 103, 0, 0, 221,
		222, 5, 61, 0, 0, 222, 223, 5, 34, 0, 0, 223, 224, 5, 103, 0, 0, 224, 225,
		5, 111, 0, 0, 225, 226, 5, 34, 0, 0, 226, 227, 5, 62, 0, 0, 227, 42, 1,
		0, 0, 0, 228, 229, 5, 60, 0, 0, 229, 230, 5, 47, 0, 0, 230, 231, 5, 115,
		0, 0, 231, 232, 5, 97, 0, 0, 232, 233, 5, 110, 0, 0, 233, 234, 5, 100,
		0, 0, 234, 235, 5, 98, 0, 0, 235, 236, 5, 111, 0, 0, 236, 237, 5, 120,
		0, 0, 237, 238, 5, 62, 0, 0, 238, 44, 1, 0, 0, 0, 239, 240, 5, 73, 0, 0,
		240, 241, 5, 110, 0, 0, 241, 242, 5, 116, 0, 0, 242, 46, 1, 0, 0, 0, 243,
		244, 5, 83, 0, 0, 244, 245, 5, 116, 0, 0, 245, 246, 5, 114, 0, 0, 246,
		247, 5, 105, 0, 0, 247, 248, 5, 110, 0, 0, 248, 249, 5, 103, 0, 0, 249,
		48, 1, 0, 0, 0, 250, 251, 5, 66, 0, 0, 251, 252, 5, 111, 0, 0, 252, 253,
		5, 111, 0, 0, 253, 254, 5, 108, 0, 0, 254, 255, 5, 101, 0, 0, 255, 256,
		5, 97, 0, 0, 256, 257, 5, 110, 0, 0, 257, 50, 1, 0, 0, 0, 258, 259, 5,
		91, 0, 0, 259, 52, 1, 0, 0, 0, 260, 261, 5, 93, 0, 0, 261, 54, 1, 0, 0,
		0, 262, 263, 5, 64, 0, 0, 263, 264, 5, 114, 0, 0, 264, 265, 5, 101, 0,
		0, 265, 266, 5, 102, 0, 0, 266, 267, 5, 101, 0, 0, 267, 268, 5, 114, 0,
		0, 268, 269, 5, 101, 0, 0, 269, 270, 5, 110, 0, 0, 270, 271, 5, 99, 0,
		0, 271, 272, 5, 101, 0, 0, 272, 273, 5, 115, 0, 0, 273, 56, 1, 0, 0, 0,
		274, 275, 5, 105, 0, 0, 275, 276, 5, 102, 0, 0, 276, 58, 1, 0, 0, 0, 277,
		278, 5, 42, 0, 0, 278, 60, 1, 0, 0, 0, 279, 280, 5, 47, 0, 0, 280, 62,
		1, 0, 0, 0, 281, 282, 5, 43, 0, 0, 282, 64, 1, 0, 0, 0, 283, 284, 5, 45,
		0, 0, 284, 66, 1, 0, 0, 0, 285, 286, 5, 61, 0, 0, 286, 287, 5, 61, 0, 0,
		287, 68, 1, 0, 0, 0, 288, 289, 5, 33, 0, 0, 289, 290, 5, 61, 0, 0, 290,
		70, 1, 0, 0, 0, 291, 292, 5, 62, 0, 0, 292, 293, 5, 61, 0, 0, 293, 72,
		1, 0, 0, 0, 294, 295, 5, 60, 0, 0, 295, 296, 5, 61, 0, 0, 296, 74, 1, 0,
		0, 0, 297, 298, 5, 38, 0, 0, 298, 299, 5, 38, 0, 0, 299, 76, 1, 0, 0, 0,
		300, 301, 5, 124, 0, 0, 301, 302, 5, 124, 0, 0, 302, 78, 1, 0, 0, 0, 303,
		304, 5, 109, 0, 0, 304, 305, 5, 97, 0, 0, 305, 306, 5, 105, 0, 0, 306,
		307, 5, 110, 0, 0, 307, 80, 1, 0, 0, 0, 308, 309, 5, 112, 0, 0, 309, 310,
		5, 114, 0, 0, 310, 311, 5, 105, 0, 0, 311, 312, 5, 110, 0, 0, 312, 313,
		5, 116, 0, 0, 313, 82, 1, 0, 0, 0, 314, 315, 5, 59, 0, 0, 315, 84, 1, 0,
		0, 0, 316, 317, 5, 46, 0, 0, 317, 318, 5, 76, 0, 0, 318, 319, 5, 105, 0,
		0, 319, 320, 5, 115, 0, 0, 320, 321, 5, 116, 0, 0, 321, 322, 5, 101, 0,
		0, 322, 323, 5, 110, 0, 0, 323, 86, 1, 0, 0, 0, 324, 328, 7, 0, 0, 0, 325,
		327, 7, 1, 0, 0, 326, 325, 1, 0, 0, 0, 327, 330, 1, 0, 0, 0, 328, 326,
		1, 0, 0, 0, 328, 329, 1, 0, 0, 0, 329, 88, 1, 0, 0, 0, 330, 328, 1, 0,
		0, 0, 331, 332, 5, 71, 0, 0, 332, 333, 5, 69, 0, 0, 333, 339, 5, 84, 0,
		0, 334, 335, 5, 80, 0, 0, 335, 336, 5, 79, 0, 0, 336, 337, 5, 83, 0, 0,
		337, 339, 5, 84, 0, 0, 338, 331, 1, 0, 0, 0, 338, 334, 1, 0, 0, 0, 339,
		90, 1, 0, 0, 0, 340, 346, 5, 34, 0, 0, 341, 345, 8, 2, 0, 0, 342, 343,
		5, 92, 0, 0, 343, 345, 9, 0, 0, 0, 344, 341, 1, 0, 0, 0, 344, 342, 1, 0,
		0, 0, 345, 348, 1, 0, 0, 0, 346, 344, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0,
		347, 349, 1, 0, 0, 0, 348, 346, 1, 0, 0, 0, 349, 350, 5, 34, 0, 0, 350,
		92, 1, 0, 0, 0, 351, 352, 5, 60, 0, 0, 352, 353, 5, 103, 0, 0, 353, 354,
		5, 111, 0, 0, 354, 355, 5, 62, 0, 0, 355, 359, 1, 0, 0, 0, 356, 358, 9,
		0, 0, 0, 357, 356, 1, 0, 0, 0, 358, 361, 1, 0, 0, 0, 359, 360, 1, 0, 0,
		0, 359, 357, 1, 0, 0, 0, 360, 362, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 362,
		363, 5, 60, 0, 0, 363, 364, 5, 47, 0, 0, 364, 365, 5, 103, 0, 0, 365, 366,
		5, 111, 0, 0, 366, 367, 5, 62, 0, 0, 367, 94, 1, 0, 0, 0, 368, 370, 7,
		3, 0, 0, 369, 368, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371, 369, 1, 0, 0,
		0, 371, 372, 1, 0, 0, 0, 372, 373, 1, 0, 0, 0, 373, 374, 6, 47, 0, 0, 374,
		96, 1, 0, 0, 0, 375, 377, 7, 4, 0, 0, 376, 375, 1, 0, 0, 0, 377, 378, 1,
		0, 0, 0, 378, 376, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 98, 1, 0, 0,
		0, 380, 381, 5, 116, 0, 0, 381, 382, 5, 114, 0, 0, 382, 383, 5, 117, 0,
		0, 383, 390, 5, 101, 0, 0, 384, 385, 5, 102, 0, 0, 385, 386, 5, 97, 0,
		0, 386, 387, 5, 108, 0, 0, 387, 388, 5, 115, 0, 0, 388, 390, 5, 101, 0,
		0, 389, 380, 1, 0, 0, 0, 389, 384, 1, 0, 0, 0, 390, 100, 1, 0, 0, 0, 391,
		405, 7, 5, 0, 0, 392, 393, 5, 61, 0, 0, 393, 405, 5, 61, 0, 0, 394, 395,
		5, 33, 0, 0, 395, 405, 5, 61, 0, 0, 396, 397, 5, 62, 0, 0, 397, 405, 5,
		61, 0, 0, 398, 399, 5, 60, 0, 0, 399, 405, 5, 61, 0, 0, 400, 401, 5, 38,
		0, 0, 401, 405, 5, 38, 0, 0, 402, 403, 5, 124, 0, 0, 403, 405, 5, 124,
		0, 0, 404, 391, 1, 0, 0, 0, 404, 392, 1, 0, 0, 0, 404, 394, 1, 0, 0, 0,
		404, 396, 1, 0, 0, 0, 404, 398, 1, 0, 0, 0, 404, 400, 1, 0, 0, 0, 404,
		402, 1, 0, 0, 0, 405, 102, 1, 0, 0, 0, 10, 0, 328, 338, 344, 346, 359,
		371, 378, 389, 404, 1, 6, 0, 0,
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

// EvoLangLexerInit initializes any static state used to implement EvoLangLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewEvoLangLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func EvoLangLexerInit() {
	staticData := &EvoLangLexerLexerStaticData
	staticData.once.Do(evolanglexerLexerInit)
}

// NewEvoLangLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewEvoLangLexer(input antlr.CharStream) *EvoLangLexer {
	EvoLangLexerInit()
	l := new(EvoLangLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &EvoLangLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "EvoLang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// EvoLangLexer tokens.
const (
	EvoLangLexerT__0        = 1
	EvoLangLexerT__1        = 2
	EvoLangLexerT__2        = 3
	EvoLangLexerT__3        = 4
	EvoLangLexerT__4        = 5
	EvoLangLexerT__5        = 6
	EvoLangLexerT__6        = 7
	EvoLangLexerT__7        = 8
	EvoLangLexerT__8        = 9
	EvoLangLexerT__9        = 10
	EvoLangLexerT__10       = 11
	EvoLangLexerT__11       = 12
	EvoLangLexerT__12       = 13
	EvoLangLexerT__13       = 14
	EvoLangLexerT__14       = 15
	EvoLangLexerT__15       = 16
	EvoLangLexerT__16       = 17
	EvoLangLexerT__17       = 18
	EvoLangLexerT__18       = 19
	EvoLangLexerT__19       = 20
	EvoLangLexerT__20       = 21
	EvoLangLexerT__21       = 22
	EvoLangLexerT__22       = 23
	EvoLangLexerT__23       = 24
	EvoLangLexerT__24       = 25
	EvoLangLexerT__25       = 26
	EvoLangLexerT__26       = 27
	EvoLangLexerT__27       = 28
	EvoLangLexerT__28       = 29
	EvoLangLexerT__29       = 30
	EvoLangLexerT__30       = 31
	EvoLangLexerT__31       = 32
	EvoLangLexerT__32       = 33
	EvoLangLexerT__33       = 34
	EvoLangLexerT__34       = 35
	EvoLangLexerT__35       = 36
	EvoLangLexerT__36       = 37
	EvoLangLexerT__37       = 38
	EvoLangLexerT__38       = 39
	EvoLangLexerT__39       = 40
	EvoLangLexerT__40       = 41
	EvoLangLexerT__41       = 42
	EvoLangLexerT__42       = 43
	EvoLangLexerID          = 44
	EvoLangLexerHTTP_METHOD = 45
	EvoLangLexerSTRING      = 46
	EvoLangLexerGO_CODE     = 47
	EvoLangLexerWS          = 48
	EvoLangLexerINT         = 49
	EvoLangLexerBOOLEAN     = 50
	EvoLangLexerOP          = 51
)