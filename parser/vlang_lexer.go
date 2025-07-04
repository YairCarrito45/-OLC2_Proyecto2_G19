// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler

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

type VlangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var VlangLexerLexerStaticData struct {
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

func vlanglexerLexerInit() {
	staticData := &VlangLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'int'", "'float64'", "'string'", "'bool'", "'char'", "'mut'", "'struct'",
		"'fn'", "'println'", "'if'", "'else'", "'for'", "'while'", "'switch'",
		"'case'", "'default'", "'break'", "'continue'", "'return'", "':'", "'len'",
		"'cap'", "'indexOf'", "'join'", "'append'", "'in'", "", "", "", "",
		"", "'nil'", "", "'+'", "'-'", "'*'", "'/'", "'%'", "'!'", "'&&'", "'||'",
		"'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'='", "':='", "'+='",
		"'-='", "'++'", "'--'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'.'",
		"','", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "MUT", "STRUCT", "FUNC", "PRINT", "IF", "ELSE",
		"FOR", "WHILE", "SWITCH", "CASE", "DEFAULT", "BREAK", "CONTINUE", "RETURN",
		"COLON", "LEN", "CAP", "INDEXOF", "JOIN", "APPEND", "IN", "BOOLEANO",
		"ENTERO", "DECIMAL", "CADENA", "CARACTER", "NIL", "ID", "PLUS", "MINUS",
		"MUL", "DIV", "MOD", "NOT", "AND", "OR", "EQ", "NEQ", "LT", "LE", "GT",
		"GE", "ASSIGN", "ASSIGN_DECL", "PLUSEQ", "MINUSEQ", "INC", "DEC", "LPAREN",
		"RPAREN", "LBRACE", "RBRACE", "LSQUARE", "RSQUARE", "DOT", "COMMA",
		"SEMICOLON", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "MUT", "STRUCT", "FUNC", "PRINT",
		"IF", "ELSE", "FOR", "WHILE", "SWITCH", "CASE", "DEFAULT", "BREAK",
		"CONTINUE", "RETURN", "COLON", "LEN", "CAP", "INDEXOF", "JOIN", "APPEND",
		"IN", "BOOLEANO", "ENTERO", "DECIMAL", "CADENA", "CARACTER", "NIL",
		"ID", "PLUS", "MINUS", "MUL", "DIV", "MOD", "NOT", "AND", "OR", "EQ",
		"NEQ", "LT", "LE", "GT", "GE", "ASSIGN", "ASSIGN_DECL", "PLUSEQ", "MINUSEQ",
		"INC", "DEC", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LSQUARE", "RSQUARE",
		"DOT", "COMMA", "SEMICOLON", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 65, 429, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 285, 8, 26, 1, 27, 4, 27, 288, 8, 27,
		11, 27, 12, 27, 289, 1, 28, 4, 28, 293, 8, 28, 11, 28, 12, 28, 294, 1,
		28, 1, 28, 4, 28, 299, 8, 28, 11, 28, 12, 28, 300, 1, 29, 1, 29, 1, 29,
		1, 29, 5, 29, 307, 8, 29, 10, 29, 12, 29, 310, 9, 29, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 5, 32,
		324, 8, 32, 10, 32, 12, 32, 327, 9, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39,
		1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1,
		43, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47,
		1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 51, 1,
		51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55,
		1, 56, 1, 56, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1,
		61, 1, 61, 1, 62, 4, 62, 399, 8, 62, 11, 62, 12, 62, 400, 1, 62, 1, 62,
		1, 63, 1, 63, 1, 63, 1, 63, 5, 63, 409, 8, 63, 10, 63, 12, 63, 412, 9,
		63, 1, 63, 1, 63, 1, 64, 1, 64, 1, 64, 1, 64, 5, 64, 420, 8, 64, 10, 64,
		12, 64, 423, 9, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 421, 0, 65, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38,
		77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47,
		95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111,
		56, 113, 57, 115, 58, 117, 59, 119, 60, 121, 61, 123, 62, 125, 63, 127,
		64, 129, 65, 1, 0, 6, 1, 0, 48, 57, 2, 0, 34, 34, 92, 92, 3, 0, 65, 90,
		95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13,
		13, 32, 32, 2, 0, 10, 10, 13, 13, 438, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35,
		1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0,
		43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0,
		0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0,
		0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0,
		0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1,
		0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81,
		1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0,
		89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0,
		0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0,
		0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111,
		1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0,
		0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1,
		0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129, 1, 0, 0, 0, 1, 131, 1, 0, 0, 0, 3,
		135, 1, 0, 0, 0, 5, 143, 1, 0, 0, 0, 7, 150, 1, 0, 0, 0, 9, 155, 1, 0,
		0, 0, 11, 160, 1, 0, 0, 0, 13, 164, 1, 0, 0, 0, 15, 171, 1, 0, 0, 0, 17,
		174, 1, 0, 0, 0, 19, 182, 1, 0, 0, 0, 21, 185, 1, 0, 0, 0, 23, 190, 1,
		0, 0, 0, 25, 194, 1, 0, 0, 0, 27, 200, 1, 0, 0, 0, 29, 207, 1, 0, 0, 0,
		31, 212, 1, 0, 0, 0, 33, 220, 1, 0, 0, 0, 35, 226, 1, 0, 0, 0, 37, 235,
		1, 0, 0, 0, 39, 242, 1, 0, 0, 0, 41, 244, 1, 0, 0, 0, 43, 248, 1, 0, 0,
		0, 45, 252, 1, 0, 0, 0, 47, 260, 1, 0, 0, 0, 49, 265, 1, 0, 0, 0, 51, 272,
		1, 0, 0, 0, 53, 284, 1, 0, 0, 0, 55, 287, 1, 0, 0, 0, 57, 292, 1, 0, 0,
		0, 59, 302, 1, 0, 0, 0, 61, 313, 1, 0, 0, 0, 63, 317, 1, 0, 0, 0, 65, 321,
		1, 0, 0, 0, 67, 328, 1, 0, 0, 0, 69, 330, 1, 0, 0, 0, 71, 332, 1, 0, 0,
		0, 73, 334, 1, 0, 0, 0, 75, 336, 1, 0, 0, 0, 77, 338, 1, 0, 0, 0, 79, 340,
		1, 0, 0, 0, 81, 343, 1, 0, 0, 0, 83, 346, 1, 0, 0, 0, 85, 349, 1, 0, 0,
		0, 87, 352, 1, 0, 0, 0, 89, 354, 1, 0, 0, 0, 91, 357, 1, 0, 0, 0, 93, 359,
		1, 0, 0, 0, 95, 362, 1, 0, 0, 0, 97, 364, 1, 0, 0, 0, 99, 367, 1, 0, 0,
		0, 101, 370, 1, 0, 0, 0, 103, 373, 1, 0, 0, 0, 105, 376, 1, 0, 0, 0, 107,
		379, 1, 0, 0, 0, 109, 381, 1, 0, 0, 0, 111, 383, 1, 0, 0, 0, 113, 385,
		1, 0, 0, 0, 115, 387, 1, 0, 0, 0, 117, 389, 1, 0, 0, 0, 119, 391, 1, 0,
		0, 0, 121, 393, 1, 0, 0, 0, 123, 395, 1, 0, 0, 0, 125, 398, 1, 0, 0, 0,
		127, 404, 1, 0, 0, 0, 129, 415, 1, 0, 0, 0, 131, 132, 5, 105, 0, 0, 132,
		133, 5, 110, 0, 0, 133, 134, 5, 116, 0, 0, 134, 2, 1, 0, 0, 0, 135, 136,
		5, 102, 0, 0, 136, 137, 5, 108, 0, 0, 137, 138, 5, 111, 0, 0, 138, 139,
		5, 97, 0, 0, 139, 140, 5, 116, 0, 0, 140, 141, 5, 54, 0, 0, 141, 142, 5,
		52, 0, 0, 142, 4, 1, 0, 0, 0, 143, 144, 5, 115, 0, 0, 144, 145, 5, 116,
		0, 0, 145, 146, 5, 114, 0, 0, 146, 147, 5, 105, 0, 0, 147, 148, 5, 110,
		0, 0, 148, 149, 5, 103, 0, 0, 149, 6, 1, 0, 0, 0, 150, 151, 5, 98, 0, 0,
		151, 152, 5, 111, 0, 0, 152, 153, 5, 111, 0, 0, 153, 154, 5, 108, 0, 0,
		154, 8, 1, 0, 0, 0, 155, 156, 5, 99, 0, 0, 156, 157, 5, 104, 0, 0, 157,
		158, 5, 97, 0, 0, 158, 159, 5, 114, 0, 0, 159, 10, 1, 0, 0, 0, 160, 161,
		5, 109, 0, 0, 161, 162, 5, 117, 0, 0, 162, 163, 5, 116, 0, 0, 163, 12,
		1, 0, 0, 0, 164, 165, 5, 115, 0, 0, 165, 166, 5, 116, 0, 0, 166, 167, 5,
		114, 0, 0, 167, 168, 5, 117, 0, 0, 168, 169, 5, 99, 0, 0, 169, 170, 5,
		116, 0, 0, 170, 14, 1, 0, 0, 0, 171, 172, 5, 102, 0, 0, 172, 173, 5, 110,
		0, 0, 173, 16, 1, 0, 0, 0, 174, 175, 5, 112, 0, 0, 175, 176, 5, 114, 0,
		0, 176, 177, 5, 105, 0, 0, 177, 178, 5, 110, 0, 0, 178, 179, 5, 116, 0,
		0, 179, 180, 5, 108, 0, 0, 180, 181, 5, 110, 0, 0, 181, 18, 1, 0, 0, 0,
		182, 183, 5, 105, 0, 0, 183, 184, 5, 102, 0, 0, 184, 20, 1, 0, 0, 0, 185,
		186, 5, 101, 0, 0, 186, 187, 5, 108, 0, 0, 187, 188, 5, 115, 0, 0, 188,
		189, 5, 101, 0, 0, 189, 22, 1, 0, 0, 0, 190, 191, 5, 102, 0, 0, 191, 192,
		5, 111, 0, 0, 192, 193, 5, 114, 0, 0, 193, 24, 1, 0, 0, 0, 194, 195, 5,
		119, 0, 0, 195, 196, 5, 104, 0, 0, 196, 197, 5, 105, 0, 0, 197, 198, 5,
		108, 0, 0, 198, 199, 5, 101, 0, 0, 199, 26, 1, 0, 0, 0, 200, 201, 5, 115,
		0, 0, 201, 202, 5, 119, 0, 0, 202, 203, 5, 105, 0, 0, 203, 204, 5, 116,
		0, 0, 204, 205, 5, 99, 0, 0, 205, 206, 5, 104, 0, 0, 206, 28, 1, 0, 0,
		0, 207, 208, 5, 99, 0, 0, 208, 209, 5, 97, 0, 0, 209, 210, 5, 115, 0, 0,
		210, 211, 5, 101, 0, 0, 211, 30, 1, 0, 0, 0, 212, 213, 5, 100, 0, 0, 213,
		214, 5, 101, 0, 0, 214, 215, 5, 102, 0, 0, 215, 216, 5, 97, 0, 0, 216,
		217, 5, 117, 0, 0, 217, 218, 5, 108, 0, 0, 218, 219, 5, 116, 0, 0, 219,
		32, 1, 0, 0, 0, 220, 221, 5, 98, 0, 0, 221, 222, 5, 114, 0, 0, 222, 223,
		5, 101, 0, 0, 223, 224, 5, 97, 0, 0, 224, 225, 5, 107, 0, 0, 225, 34, 1,
		0, 0, 0, 226, 227, 5, 99, 0, 0, 227, 228, 5, 111, 0, 0, 228, 229, 5, 110,
		0, 0, 229, 230, 5, 116, 0, 0, 230, 231, 5, 105, 0, 0, 231, 232, 5, 110,
		0, 0, 232, 233, 5, 117, 0, 0, 233, 234, 5, 101, 0, 0, 234, 36, 1, 0, 0,
		0, 235, 236, 5, 114, 0, 0, 236, 237, 5, 101, 0, 0, 237, 238, 5, 116, 0,
		0, 238, 239, 5, 117, 0, 0, 239, 240, 5, 114, 0, 0, 240, 241, 5, 110, 0,
		0, 241, 38, 1, 0, 0, 0, 242, 243, 5, 58, 0, 0, 243, 40, 1, 0, 0, 0, 244,
		245, 5, 108, 0, 0, 245, 246, 5, 101, 0, 0, 246, 247, 5, 110, 0, 0, 247,
		42, 1, 0, 0, 0, 248, 249, 5, 99, 0, 0, 249, 250, 5, 97, 0, 0, 250, 251,
		5, 112, 0, 0, 251, 44, 1, 0, 0, 0, 252, 253, 5, 105, 0, 0, 253, 254, 5,
		110, 0, 0, 254, 255, 5, 100, 0, 0, 255, 256, 5, 101, 0, 0, 256, 257, 5,
		120, 0, 0, 257, 258, 5, 79, 0, 0, 258, 259, 5, 102, 0, 0, 259, 46, 1, 0,
		0, 0, 260, 261, 5, 106, 0, 0, 261, 262, 5, 111, 0, 0, 262, 263, 5, 105,
		0, 0, 263, 264, 5, 110, 0, 0, 264, 48, 1, 0, 0, 0, 265, 266, 5, 97, 0,
		0, 266, 267, 5, 112, 0, 0, 267, 268, 5, 112, 0, 0, 268, 269, 5, 101, 0,
		0, 269, 270, 5, 110, 0, 0, 270, 271, 5, 100, 0, 0, 271, 50, 1, 0, 0, 0,
		272, 273, 5, 105, 0, 0, 273, 274, 5, 110, 0, 0, 274, 52, 1, 0, 0, 0, 275,
		276, 5, 116, 0, 0, 276, 277, 5, 114, 0, 0, 277, 278, 5, 117, 0, 0, 278,
		285, 5, 101, 0, 0, 279, 280, 5, 102, 0, 0, 280, 281, 5, 97, 0, 0, 281,
		282, 5, 108, 0, 0, 282, 283, 5, 115, 0, 0, 283, 285, 5, 101, 0, 0, 284,
		275, 1, 0, 0, 0, 284, 279, 1, 0, 0, 0, 285, 54, 1, 0, 0, 0, 286, 288, 7,
		0, 0, 0, 287, 286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 287, 1, 0, 0,
		0, 289, 290, 1, 0, 0, 0, 290, 56, 1, 0, 0, 0, 291, 293, 7, 0, 0, 0, 292,
		291, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 294, 295,
		1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 298, 5, 46, 0, 0, 297, 299, 7, 0,
		0, 0, 298, 297, 1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0,
		300, 301, 1, 0, 0, 0, 301, 58, 1, 0, 0, 0, 302, 308, 5, 34, 0, 0, 303,
		307, 8, 1, 0, 0, 304, 305, 5, 92, 0, 0, 305, 307, 9, 0, 0, 0, 306, 303,
		1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 307, 310, 1, 0, 0, 0, 308, 306, 1, 0,
		0, 0, 308, 309, 1, 0, 0, 0, 309, 311, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0,
		311, 312, 5, 34, 0, 0, 312, 60, 1, 0, 0, 0, 313, 314, 5, 39, 0, 0, 314,
		315, 9, 0, 0, 0, 315, 316, 5, 39, 0, 0, 316, 62, 1, 0, 0, 0, 317, 318,
		5, 110, 0, 0, 318, 319, 5, 105, 0, 0, 319, 320, 5, 108, 0, 0, 320, 64,
		1, 0, 0, 0, 321, 325, 7, 2, 0, 0, 322, 324, 7, 3, 0, 0, 323, 322, 1, 0,
		0, 0, 324, 327, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0,
		326, 66, 1, 0, 0, 0, 327, 325, 1, 0, 0, 0, 328, 329, 5, 43, 0, 0, 329,
		68, 1, 0, 0, 0, 330, 331, 5, 45, 0, 0, 331, 70, 1, 0, 0, 0, 332, 333, 5,
		42, 0, 0, 333, 72, 1, 0, 0, 0, 334, 335, 5, 47, 0, 0, 335, 74, 1, 0, 0,
		0, 336, 337, 5, 37, 0, 0, 337, 76, 1, 0, 0, 0, 338, 339, 5, 33, 0, 0, 339,
		78, 1, 0, 0, 0, 340, 341, 5, 38, 0, 0, 341, 342, 5, 38, 0, 0, 342, 80,
		1, 0, 0, 0, 343, 344, 5, 124, 0, 0, 344, 345, 5, 124, 0, 0, 345, 82, 1,
		0, 0, 0, 346, 347, 5, 61, 0, 0, 347, 348, 5, 61, 0, 0, 348, 84, 1, 0, 0,
		0, 349, 350, 5, 33, 0, 0, 350, 351, 5, 61, 0, 0, 351, 86, 1, 0, 0, 0, 352,
		353, 5, 60, 0, 0, 353, 88, 1, 0, 0, 0, 354, 355, 5, 60, 0, 0, 355, 356,
		5, 61, 0, 0, 356, 90, 1, 0, 0, 0, 357, 358, 5, 62, 0, 0, 358, 92, 1, 0,
		0, 0, 359, 360, 5, 62, 0, 0, 360, 361, 5, 61, 0, 0, 361, 94, 1, 0, 0, 0,
		362, 363, 5, 61, 0, 0, 363, 96, 1, 0, 0, 0, 364, 365, 5, 58, 0, 0, 365,
		366, 5, 61, 0, 0, 366, 98, 1, 0, 0, 0, 367, 368, 5, 43, 0, 0, 368, 369,
		5, 61, 0, 0, 369, 100, 1, 0, 0, 0, 370, 371, 5, 45, 0, 0, 371, 372, 5,
		61, 0, 0, 372, 102, 1, 0, 0, 0, 373, 374, 5, 43, 0, 0, 374, 375, 5, 43,
		0, 0, 375, 104, 1, 0, 0, 0, 376, 377, 5, 45, 0, 0, 377, 378, 5, 45, 0,
		0, 378, 106, 1, 0, 0, 0, 379, 380, 5, 40, 0, 0, 380, 108, 1, 0, 0, 0, 381,
		382, 5, 41, 0, 0, 382, 110, 1, 0, 0, 0, 383, 384, 5, 123, 0, 0, 384, 112,
		1, 0, 0, 0, 385, 386, 5, 125, 0, 0, 386, 114, 1, 0, 0, 0, 387, 388, 5,
		91, 0, 0, 388, 116, 1, 0, 0, 0, 389, 390, 5, 93, 0, 0, 390, 118, 1, 0,
		0, 0, 391, 392, 5, 46, 0, 0, 392, 120, 1, 0, 0, 0, 393, 394, 5, 44, 0,
		0, 394, 122, 1, 0, 0, 0, 395, 396, 5, 59, 0, 0, 396, 124, 1, 0, 0, 0, 397,
		399, 7, 4, 0, 0, 398, 397, 1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 398,
		1, 0, 0, 0, 400, 401, 1, 0, 0, 0, 401, 402, 1, 0, 0, 0, 402, 403, 6, 62,
		0, 0, 403, 126, 1, 0, 0, 0, 404, 405, 5, 47, 0, 0, 405, 406, 5, 47, 0,
		0, 406, 410, 1, 0, 0, 0, 407, 409, 8, 5, 0, 0, 408, 407, 1, 0, 0, 0, 409,
		412, 1, 0, 0, 0, 410, 408, 1, 0, 0, 0, 410, 411, 1, 0, 0, 0, 411, 413,
		1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 413, 414, 6, 63, 0, 0, 414, 128, 1, 0,
		0, 0, 415, 416, 5, 47, 0, 0, 416, 417, 5, 42, 0, 0, 417, 421, 1, 0, 0,
		0, 418, 420, 9, 0, 0, 0, 419, 418, 1, 0, 0, 0, 420, 423, 1, 0, 0, 0, 421,
		422, 1, 0, 0, 0, 421, 419, 1, 0, 0, 0, 422, 424, 1, 0, 0, 0, 423, 421,
		1, 0, 0, 0, 424, 425, 5, 42, 0, 0, 425, 426, 5, 47, 0, 0, 426, 427, 1,
		0, 0, 0, 427, 428, 6, 64, 0, 0, 428, 130, 1, 0, 0, 0, 11, 0, 284, 289,
		294, 300, 306, 308, 325, 400, 410, 421, 1, 6, 0, 0,
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

// VlangLexerInit initializes any static state used to implement VlangLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewVlangLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func VlangLexerInit() {
	staticData := &VlangLexerLexerStaticData
	staticData.once.Do(vlanglexerLexerInit)
}

// NewVlangLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewVlangLexer(input antlr.CharStream) *VlangLexer {
	VlangLexerInit()
	l := new(VlangLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &VlangLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Vlang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// VlangLexer tokens.
const (
	VlangLexerT__0          = 1
	VlangLexerT__1          = 2
	VlangLexerT__2          = 3
	VlangLexerT__3          = 4
	VlangLexerT__4          = 5
	VlangLexerMUT           = 6
	VlangLexerSTRUCT        = 7
	VlangLexerFUNC          = 8
	VlangLexerPRINT         = 9
	VlangLexerIF            = 10
	VlangLexerELSE          = 11
	VlangLexerFOR           = 12
	VlangLexerWHILE         = 13
	VlangLexerSWITCH        = 14
	VlangLexerCASE          = 15
	VlangLexerDEFAULT       = 16
	VlangLexerBREAK         = 17
	VlangLexerCONTINUE      = 18
	VlangLexerRETURN        = 19
	VlangLexerCOLON         = 20
	VlangLexerLEN           = 21
	VlangLexerCAP           = 22
	VlangLexerINDEXOF       = 23
	VlangLexerJOIN          = 24
	VlangLexerAPPEND        = 25
	VlangLexerIN            = 26
	VlangLexerBOOLEANO      = 27
	VlangLexerENTERO        = 28
	VlangLexerDECIMAL       = 29
	VlangLexerCADENA        = 30
	VlangLexerCARACTER      = 31
	VlangLexerNIL           = 32
	VlangLexerID            = 33
	VlangLexerPLUS          = 34
	VlangLexerMINUS         = 35
	VlangLexerMUL           = 36
	VlangLexerDIV           = 37
	VlangLexerMOD           = 38
	VlangLexerNOT           = 39
	VlangLexerAND           = 40
	VlangLexerOR            = 41
	VlangLexerEQ            = 42
	VlangLexerNEQ           = 43
	VlangLexerLT            = 44
	VlangLexerLE            = 45
	VlangLexerGT            = 46
	VlangLexerGE            = 47
	VlangLexerASSIGN        = 48
	VlangLexerASSIGN_DECL   = 49
	VlangLexerPLUSEQ        = 50
	VlangLexerMINUSEQ       = 51
	VlangLexerINC           = 52
	VlangLexerDEC           = 53
	VlangLexerLPAREN        = 54
	VlangLexerRPAREN        = 55
	VlangLexerLBRACE        = 56
	VlangLexerRBRACE        = 57
	VlangLexerLSQUARE       = 58
	VlangLexerRSQUARE       = 59
	VlangLexerDOT           = 60
	VlangLexerCOMMA         = 61
	VlangLexerSEMICOLON     = 62
	VlangLexerWS            = 63
	VlangLexerLINE_COMMENT  = 64
	VlangLexerBLOCK_COMMENT = 65
)
