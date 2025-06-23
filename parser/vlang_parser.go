// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
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

type VlangParser struct {
	*antlr.BaseParser
}

var VlangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func vlangParserInit() {
	staticData := &VlangParserStaticData
	staticData.LiteralNames = []string{
		"", "'int'", "'float64'", "'string'", "'bool'", "'char'", "'mut'", "'struct'",
		"'fn'", "'println'", "'if'", "'else'", "'for'", "'while'", "'switch'",
		"'case'", "'default'", "'break'", "'continue'", "'return'", "':'", "'len'",
		"'cap'", "'indexOf'", "'join'", "'append'", "'in'", "", "", "", "",
		"", "", "'+'", "'-'", "'*'", "'/'", "'%'", "'!'", "'&&'", "'||'", "'=='",
		"'!='", "'<'", "'<='", "'>'", "'>='", "'='", "':='", "'+='", "'-='",
		"'++'", "'--'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'.'", "','",
		"';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "MUT", "STRUCT", "FUNC", "PRINT", "IF", "ELSE",
		"FOR", "WHILE", "SWITCH", "CASE", "DEFAULT", "BREAK", "CONTINUE", "RETURN",
		"COLON", "LEN", "CAP", "INDEXOF", "JOIN", "APPEND", "IN", "BOOLEANO",
		"ENTERO", "DECIMAL", "CADENA", "CARACTER", "ID", "PLUS", "MINUS", "MUL",
		"DIV", "MOD", "NOT", "AND", "OR", "EQ", "NEQ", "LT", "LE", "GT", "GE",
		"ASSIGN", "ASSIGN_DECL", "PLUSEQ", "MINUSEQ", "INC", "DEC", "LPAREN",
		"RPAREN", "LBRACE", "RBRACE", "LSQUARE", "RSQUARE", "DOT", "COMMA",
		"SEMICOLON", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"programa", "declaraciones", "structDecl", "atributos", "atributo",
		"inicializadorStruct", "inicializadorCampo", "funcionStruct", "funcion",
		"parametrosFormales", "parametroFormal", "tipoConReferencia", "varDcl",
		"tipo", "sliceTipo", "stmt", "sentencias_control", "ifDcl", "elseifDcl",
		"elseDcl", "switchDcl", "caseBlock", "defaultBlock", "forDcl", "whileDcl",
		"bloque", "expresion", "parametros", "valor",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 64, 462, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 5, 0, 60, 8, 0, 10, 0, 12, 0,
		63, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 72, 8, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 4, 3, 81, 8, 3, 11, 3, 12, 3, 82,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 92, 8, 5, 10, 5, 12, 5,
		95, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 105, 8,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 112, 8, 7, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 3, 8, 121, 8, 8, 1, 8, 1, 8, 3, 8, 125, 8, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 132, 8, 9, 10, 9, 12, 9, 135, 9, 9, 1, 10,
		1, 10, 1, 10, 1, 11, 3, 11, 141, 8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		12, 3, 12, 148, 8, 12, 1, 12, 1, 12, 3, 12, 152, 8, 12, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 3, 13, 160, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 3, 15, 180, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 3, 15, 187, 8, 15, 1, 15, 1, 15, 3, 15, 191, 8, 15, 1, 16, 1, 16,
		1, 16, 1, 16, 3, 16, 197, 8, 16, 1, 17, 1, 17, 3, 17, 201, 8, 17, 1, 17,
		1, 17, 3, 17, 205, 8, 17, 1, 17, 1, 17, 5, 17, 209, 8, 17, 10, 17, 12,
		17, 212, 9, 17, 1, 17, 3, 17, 215, 8, 17, 1, 18, 1, 18, 1, 18, 3, 18, 220,
		8, 18, 1, 18, 1, 18, 3, 18, 224, 8, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		19, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 235, 8, 20, 10, 20, 12, 20, 238,
		9, 20, 1, 20, 3, 20, 241, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1,
		21, 5, 21, 249, 8, 21, 10, 21, 12, 21, 252, 9, 21, 1, 22, 1, 22, 1, 22,
		5, 22, 257, 8, 22, 10, 22, 12, 22, 260, 9, 22, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 281, 8, 23, 1, 24, 1, 24, 3,
		24, 285, 8, 24, 1, 24, 1, 24, 3, 24, 289, 8, 24, 1, 24, 1, 24, 1, 25, 1,
		25, 5, 25, 295, 8, 25, 10, 25, 12, 25, 298, 9, 25, 1, 25, 1, 25, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 312,
		8, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 3, 26, 335, 8, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 341, 8, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 3, 26, 354, 8, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26,
		362, 8, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 422, 8, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 442, 8, 26,
		10, 26, 12, 26, 445, 9, 26, 1, 27, 1, 27, 1, 27, 5, 27, 450, 8, 27, 10,
		27, 12, 27, 453, 9, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 460,
		8, 28, 1, 28, 0, 1, 52, 29, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 0,
		6, 1, 0, 47, 48, 2, 0, 34, 34, 38, 38, 1, 0, 35, 37, 1, 0, 33, 34, 1, 0,
		43, 46, 1, 0, 41, 42, 518, 0, 61, 1, 0, 0, 0, 2, 71, 1, 0, 0, 0, 4, 73,
		1, 0, 0, 0, 6, 80, 1, 0, 0, 0, 8, 84, 1, 0, 0, 0, 10, 88, 1, 0, 0, 0, 12,
		96, 1, 0, 0, 0, 14, 100, 1, 0, 0, 0, 16, 116, 1, 0, 0, 0, 18, 128, 1, 0,
		0, 0, 20, 136, 1, 0, 0, 0, 22, 140, 1, 0, 0, 0, 24, 144, 1, 0, 0, 0, 26,
		159, 1, 0, 0, 0, 28, 161, 1, 0, 0, 0, 30, 190, 1, 0, 0, 0, 32, 196, 1,
		0, 0, 0, 34, 198, 1, 0, 0, 0, 36, 216, 1, 0, 0, 0, 38, 227, 1, 0, 0, 0,
		40, 230, 1, 0, 0, 0, 42, 244, 1, 0, 0, 0, 44, 253, 1, 0, 0, 0, 46, 280,
		1, 0, 0, 0, 48, 282, 1, 0, 0, 0, 50, 292, 1, 0, 0, 0, 52, 421, 1, 0, 0,
		0, 54, 446, 1, 0, 0, 0, 56, 459, 1, 0, 0, 0, 58, 60, 3, 2, 1, 0, 59, 58,
		1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0,
		62, 64, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 65, 5, 0, 0, 1, 65, 1, 1, 0,
		0, 0, 66, 72, 3, 24, 12, 0, 67, 72, 3, 4, 2, 0, 68, 72, 3, 14, 7, 0, 69,
		72, 3, 16, 8, 0, 70, 72, 3, 30, 15, 0, 71, 66, 1, 0, 0, 0, 71, 67, 1, 0,
		0, 0, 71, 68, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 70, 1, 0, 0, 0, 72, 3,
		1, 0, 0, 0, 73, 74, 5, 7, 0, 0, 74, 75, 5, 32, 0, 0, 75, 76, 5, 55, 0,
		0, 76, 77, 3, 6, 3, 0, 77, 78, 5, 56, 0, 0, 78, 5, 1, 0, 0, 0, 79, 81,
		3, 8, 4, 0, 80, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0,
		82, 83, 1, 0, 0, 0, 83, 7, 1, 0, 0, 0, 84, 85, 3, 26, 13, 0, 85, 86, 5,
		32, 0, 0, 86, 87, 5, 61, 0, 0, 87, 9, 1, 0, 0, 0, 88, 93, 3, 12, 6, 0,
		89, 90, 5, 60, 0, 0, 90, 92, 3, 12, 6, 0, 91, 89, 1, 0, 0, 0, 92, 95, 1,
		0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 11, 1, 0, 0, 0, 95,
		93, 1, 0, 0, 0, 96, 97, 5, 32, 0, 0, 97, 98, 5, 20, 0, 0, 98, 99, 3, 52,
		26, 0, 99, 13, 1, 0, 0, 0, 100, 101, 5, 8, 0, 0, 101, 102, 5, 53, 0, 0,
		102, 104, 5, 32, 0, 0, 103, 105, 5, 35, 0, 0, 104, 103, 1, 0, 0, 0, 104,
		105, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 5, 32, 0, 0, 107, 108,
		5, 54, 0, 0, 108, 109, 5, 32, 0, 0, 109, 111, 5, 53, 0, 0, 110, 112, 3,
		18, 9, 0, 111, 110, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 113, 1, 0, 0,
		0, 113, 114, 5, 54, 0, 0, 114, 115, 3, 50, 25, 0, 115, 15, 1, 0, 0, 0,
		116, 117, 5, 8, 0, 0, 117, 118, 5, 32, 0, 0, 118, 120, 5, 53, 0, 0, 119,
		121, 3, 18, 9, 0, 120, 119, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 122,
		1, 0, 0, 0, 122, 124, 5, 54, 0, 0, 123, 125, 3, 26, 13, 0, 124, 123, 1,
		0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 3, 50, 25,
		0, 127, 17, 1, 0, 0, 0, 128, 133, 3, 20, 10, 0, 129, 130, 5, 60, 0, 0,
		130, 132, 3, 20, 10, 0, 131, 129, 1, 0, 0, 0, 132, 135, 1, 0, 0, 0, 133,
		131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 19, 1, 0, 0, 0, 135, 133, 1,
		0, 0, 0, 136, 137, 3, 22, 11, 0, 137, 138, 5, 32, 0, 0, 138, 21, 1, 0,
		0, 0, 139, 141, 5, 35, 0, 0, 140, 139, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0,
		141, 142, 1, 0, 0, 0, 142, 143, 3, 26, 13, 0, 143, 23, 1, 0, 0, 0, 144,
		145, 5, 6, 0, 0, 145, 147, 5, 32, 0, 0, 146, 148, 3, 26, 13, 0, 147, 146,
		1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0, 149, 150, 7, 0,
		0, 0, 150, 152, 3, 52, 26, 0, 151, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0,
		0, 152, 25, 1, 0, 0, 0, 153, 160, 3, 28, 14, 0, 154, 160, 5, 1, 0, 0, 155,
		160, 5, 2, 0, 0, 156, 160, 5, 3, 0, 0, 157, 160, 5, 4, 0, 0, 158, 160,
		5, 5, 0, 0, 159, 153, 1, 0, 0, 0, 159, 154, 1, 0, 0, 0, 159, 155, 1, 0,
		0, 0, 159, 156, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159, 158, 1, 0, 0, 0,
		160, 27, 1, 0, 0, 0, 161, 162, 5, 57, 0, 0, 162, 163, 5, 58, 0, 0, 163,
		164, 3, 26, 13, 0, 164, 29, 1, 0, 0, 0, 165, 191, 3, 52, 26, 0, 166, 167,
		5, 32, 0, 0, 167, 168, 5, 57, 0, 0, 168, 169, 3, 52, 26, 0, 169, 170, 5,
		58, 0, 0, 170, 171, 5, 57, 0, 0, 171, 172, 3, 52, 26, 0, 172, 173, 5, 58,
		0, 0, 173, 174, 5, 47, 0, 0, 174, 175, 3, 52, 26, 0, 175, 191, 1, 0, 0,
		0, 176, 177, 5, 9, 0, 0, 177, 179, 5, 53, 0, 0, 178, 180, 3, 54, 27, 0,
		179, 178, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181,
		191, 5, 54, 0, 0, 182, 191, 5, 17, 0, 0, 183, 191, 5, 18, 0, 0, 184, 186,
		5, 19, 0, 0, 185, 187, 3, 52, 26, 0, 186, 185, 1, 0, 0, 0, 186, 187, 1,
		0, 0, 0, 187, 191, 1, 0, 0, 0, 188, 191, 3, 32, 16, 0, 189, 191, 3, 50,
		25, 0, 190, 165, 1, 0, 0, 0, 190, 166, 1, 0, 0, 0, 190, 176, 1, 0, 0, 0,
		190, 182, 1, 0, 0, 0, 190, 183, 1, 0, 0, 0, 190, 184, 1, 0, 0, 0, 190,
		188, 1, 0, 0, 0, 190, 189, 1, 0, 0, 0, 191, 31, 1, 0, 0, 0, 192, 197, 3,
		34, 17, 0, 193, 197, 3, 46, 23, 0, 194, 197, 3, 48, 24, 0, 195, 197, 3,
		40, 20, 0, 196, 192, 1, 0, 0, 0, 196, 193, 1, 0, 0, 0, 196, 194, 1, 0,
		0, 0, 196, 195, 1, 0, 0, 0, 197, 33, 1, 0, 0, 0, 198, 200, 5, 10, 0, 0,
		199, 201, 5, 53, 0, 0, 200, 199, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201,
		202, 1, 0, 0, 0, 202, 204, 3, 52, 26, 0, 203, 205, 5, 54, 0, 0, 204, 203,
		1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 210, 3, 50,
		25, 0, 207, 209, 3, 36, 18, 0, 208, 207, 1, 0, 0, 0, 209, 212, 1, 0, 0,
		0, 210, 208, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 214, 1, 0, 0, 0, 212,
		210, 1, 0, 0, 0, 213, 215, 3, 38, 19, 0, 214, 213, 1, 0, 0, 0, 214, 215,
		1, 0, 0, 0, 215, 35, 1, 0, 0, 0, 216, 217, 5, 11, 0, 0, 217, 219, 5, 10,
		0, 0, 218, 220, 5, 53, 0, 0, 219, 218, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0,
		220, 221, 1, 0, 0, 0, 221, 223, 3, 52, 26, 0, 222, 224, 5, 54, 0, 0, 223,
		222, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 226,
		3, 50, 25, 0, 226, 37, 1, 0, 0, 0, 227, 228, 5, 11, 0, 0, 228, 229, 3,
		50, 25, 0, 229, 39, 1, 0, 0, 0, 230, 231, 5, 14, 0, 0, 231, 232, 3, 52,
		26, 0, 232, 236, 5, 55, 0, 0, 233, 235, 3, 42, 21, 0, 234, 233, 1, 0, 0,
		0, 235, 238, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237,
		240, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 239, 241, 3, 44, 22, 0, 240, 239,
		1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 243, 5, 56,
		0, 0, 243, 41, 1, 0, 0, 0, 244, 245, 5, 15, 0, 0, 245, 246, 3, 52, 26,
		0, 246, 250, 5, 20, 0, 0, 247, 249, 3, 2, 1, 0, 248, 247, 1, 0, 0, 0, 249,
		252, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 43, 1,
		0, 0, 0, 252, 250, 1, 0, 0, 0, 253, 254, 5, 16, 0, 0, 254, 258, 5, 20,
		0, 0, 255, 257, 3, 2, 1, 0, 256, 255, 1, 0, 0, 0, 257, 260, 1, 0, 0, 0,
		258, 256, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 45, 1, 0, 0, 0, 260, 258,
		1, 0, 0, 0, 261, 262, 5, 12, 0, 0, 262, 263, 3, 52, 26, 0, 263, 264, 3,
		50, 25, 0, 264, 281, 1, 0, 0, 0, 265, 266, 5, 12, 0, 0, 266, 267, 3, 52,
		26, 0, 267, 268, 5, 61, 0, 0, 268, 269, 3, 52, 26, 0, 269, 270, 5, 61,
		0, 0, 270, 271, 3, 52, 26, 0, 271, 272, 3, 50, 25, 0, 272, 281, 1, 0, 0,
		0, 273, 274, 5, 12, 0, 0, 274, 275, 5, 32, 0, 0, 275, 276, 5, 60, 0, 0,
		276, 277, 5, 32, 0, 0, 277, 278, 5, 26, 0, 0, 278, 279, 5, 32, 0, 0, 279,
		281, 3, 50, 25, 0, 280, 261, 1, 0, 0, 0, 280, 265, 1, 0, 0, 0, 280, 273,
		1, 0, 0, 0, 281, 47, 1, 0, 0, 0, 282, 284, 5, 13, 0, 0, 283, 285, 5, 53,
		0, 0, 284, 283, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0,
		286, 288, 3, 52, 26, 0, 287, 289, 5, 54, 0, 0, 288, 287, 1, 0, 0, 0, 288,
		289, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 291, 3, 50, 25, 0, 291, 49,
		1, 0, 0, 0, 292, 296, 5, 55, 0, 0, 293, 295, 3, 2, 1, 0, 294, 293, 1, 0,
		0, 0, 295, 298, 1, 0, 0, 0, 296, 294, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0,
		297, 299, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 299, 300, 5, 56, 0, 0, 300,
		51, 1, 0, 0, 0, 301, 302, 6, 26, -1, 0, 302, 303, 5, 32, 0, 0, 303, 304,
		5, 49, 0, 0, 304, 422, 3, 52, 26, 30, 305, 306, 5, 32, 0, 0, 306, 307,
		5, 50, 0, 0, 307, 422, 3, 52, 26, 29, 308, 309, 5, 32, 0, 0, 309, 311,
		5, 53, 0, 0, 310, 312, 3, 54, 27, 0, 311, 310, 1, 0, 0, 0, 311, 312, 1,
		0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 422, 5, 54, 0, 0, 314, 315, 5, 32,
		0, 0, 315, 316, 5, 47, 0, 0, 316, 422, 3, 52, 26, 27, 317, 318, 5, 32,
		0, 0, 318, 422, 5, 51, 0, 0, 319, 320, 5, 32, 0, 0, 320, 422, 5, 52, 0,
		0, 321, 322, 5, 32, 0, 0, 322, 323, 5, 59, 0, 0, 323, 324, 5, 32, 0, 0,
		324, 325, 5, 47, 0, 0, 325, 422, 3, 52, 26, 24, 326, 327, 5, 32, 0, 0,
		327, 328, 5, 59, 0, 0, 328, 422, 5, 32, 0, 0, 329, 330, 5, 32, 0, 0, 330,
		331, 5, 59, 0, 0, 331, 332, 5, 32, 0, 0, 332, 334, 5, 53, 0, 0, 333, 335,
		3, 54, 27, 0, 334, 333, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 336, 1,
		0, 0, 0, 336, 422, 5, 54, 0, 0, 337, 338, 5, 32, 0, 0, 338, 340, 5, 55,
		0, 0, 339, 341, 3, 10, 5, 0, 340, 339, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0,
		341, 342, 1, 0, 0, 0, 342, 422, 5, 56, 0, 0, 343, 344, 7, 1, 0, 0, 344,
		422, 3, 52, 26, 14, 345, 346, 5, 53, 0, 0, 346, 347, 3, 52, 26, 0, 347,
		348, 5, 54, 0, 0, 348, 422, 1, 0, 0, 0, 349, 422, 3, 56, 28, 0, 350, 422,
		5, 32, 0, 0, 351, 353, 5, 55, 0, 0, 352, 354, 3, 54, 27, 0, 353, 352, 1,
		0, 0, 0, 353, 354, 1, 0, 0, 0, 354, 355, 1, 0, 0, 0, 355, 422, 5, 56, 0,
		0, 356, 357, 5, 57, 0, 0, 357, 358, 5, 58, 0, 0, 358, 359, 3, 26, 13, 0,
		359, 361, 5, 55, 0, 0, 360, 362, 3, 54, 27, 0, 361, 360, 1, 0, 0, 0, 361,
		362, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 364, 5, 56, 0, 0, 364, 422,
		1, 0, 0, 0, 365, 366, 5, 23, 0, 0, 366, 367, 5, 53, 0, 0, 367, 368, 3,
		52, 26, 0, 368, 369, 5, 60, 0, 0, 369, 370, 3, 52, 26, 0, 370, 371, 5,
		54, 0, 0, 371, 422, 1, 0, 0, 0, 372, 373, 5, 24, 0, 0, 373, 374, 5, 53,
		0, 0, 374, 375, 3, 52, 26, 0, 375, 376, 5, 60, 0, 0, 376, 377, 3, 52, 26,
		0, 377, 378, 5, 54, 0, 0, 378, 422, 1, 0, 0, 0, 379, 380, 5, 21, 0, 0,
		380, 381, 5, 53, 0, 0, 381, 382, 3, 52, 26, 0, 382, 383, 5, 54, 0, 0, 383,
		422, 1, 0, 0, 0, 384, 385, 5, 25, 0, 0, 385, 386, 5, 53, 0, 0, 386, 387,
		3, 52, 26, 0, 387, 388, 5, 60, 0, 0, 388, 389, 3, 52, 26, 0, 389, 390,
		5, 54, 0, 0, 390, 422, 1, 0, 0, 0, 391, 392, 5, 32, 0, 0, 392, 393, 5,
		57, 0, 0, 393, 394, 3, 52, 26, 0, 394, 395, 5, 58, 0, 0, 395, 422, 1, 0,
		0, 0, 396, 397, 5, 32, 0, 0, 397, 398, 5, 57, 0, 0, 398, 399, 3, 52, 26,
		0, 399, 400, 5, 58, 0, 0, 400, 401, 5, 47, 0, 0, 401, 402, 3, 52, 26, 3,
		402, 422, 1, 0, 0, 0, 403, 404, 5, 32, 0, 0, 404, 405, 5, 57, 0, 0, 405,
		406, 3, 52, 26, 0, 406, 407, 5, 58, 0, 0, 407, 408, 5, 57, 0, 0, 408, 409,
		3, 52, 26, 0, 409, 410, 5, 58, 0, 0, 410, 422, 1, 0, 0, 0, 411, 412, 5,
		32, 0, 0, 412, 413, 5, 57, 0, 0, 413, 414, 3, 52, 26, 0, 414, 415, 5, 58,
		0, 0, 415, 416, 5, 57, 0, 0, 416, 417, 3, 52, 26, 0, 417, 418, 5, 58, 0,
		0, 418, 419, 5, 47, 0, 0, 419, 420, 3, 52, 26, 1, 420, 422, 1, 0, 0, 0,
		421, 301, 1, 0, 0, 0, 421, 305, 1, 0, 0, 0, 421, 308, 1, 0, 0, 0, 421,
		314, 1, 0, 0, 0, 421, 317, 1, 0, 0, 0, 421, 319, 1, 0, 0, 0, 421, 321,
		1, 0, 0, 0, 421, 326, 1, 0, 0, 0, 421, 329, 1, 0, 0, 0, 421, 337, 1, 0,
		0, 0, 421, 343, 1, 0, 0, 0, 421, 345, 1, 0, 0, 0, 421, 349, 1, 0, 0, 0,
		421, 350, 1, 0, 0, 0, 421, 351, 1, 0, 0, 0, 421, 356, 1, 0, 0, 0, 421,
		365, 1, 0, 0, 0, 421, 372, 1, 0, 0, 0, 421, 379, 1, 0, 0, 0, 421, 384,
		1, 0, 0, 0, 421, 391, 1, 0, 0, 0, 421, 396, 1, 0, 0, 0, 421, 403, 1, 0,
		0, 0, 421, 411, 1, 0, 0, 0, 422, 443, 1, 0, 0, 0, 423, 424, 10, 20, 0,
		0, 424, 425, 7, 2, 0, 0, 425, 442, 3, 52, 26, 21, 426, 427, 10, 19, 0,
		0, 427, 428, 7, 3, 0, 0, 428, 442, 3, 52, 26, 20, 429, 430, 10, 18, 0,
		0, 430, 431, 7, 4, 0, 0, 431, 442, 3, 52, 26, 19, 432, 433, 10, 17, 0,
		0, 433, 434, 7, 5, 0, 0, 434, 442, 3, 52, 26, 18, 435, 436, 10, 16, 0,
		0, 436, 437, 5, 39, 0, 0, 437, 442, 3, 52, 26, 17, 438, 439, 10, 15, 0,
		0, 439, 440, 5, 40, 0, 0, 440, 442, 3, 52, 26, 16, 441, 423, 1, 0, 0, 0,
		441, 426, 1, 0, 0, 0, 441, 429, 1, 0, 0, 0, 441, 432, 1, 0, 0, 0, 441,
		435, 1, 0, 0, 0, 441, 438, 1, 0, 0, 0, 442, 445, 1, 0, 0, 0, 443, 441,
		1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444, 53, 1, 0, 0, 0, 445, 443, 1, 0,
		0, 0, 446, 451, 3, 52, 26, 0, 447, 448, 5, 60, 0, 0, 448, 450, 3, 52, 26,
		0, 449, 447, 1, 0, 0, 0, 450, 453, 1, 0, 0, 0, 451, 449, 1, 0, 0, 0, 451,
		452, 1, 0, 0, 0, 452, 55, 1, 0, 0, 0, 453, 451, 1, 0, 0, 0, 454, 460, 5,
		28, 0, 0, 455, 460, 5, 29, 0, 0, 456, 460, 5, 30, 0, 0, 457, 460, 5, 27,
		0, 0, 458, 460, 5, 31, 0, 0, 459, 454, 1, 0, 0, 0, 459, 455, 1, 0, 0, 0,
		459, 456, 1, 0, 0, 0, 459, 457, 1, 0, 0, 0, 459, 458, 1, 0, 0, 0, 460,
		57, 1, 0, 0, 0, 41, 61, 71, 82, 93, 104, 111, 120, 124, 133, 140, 147,
		151, 159, 179, 186, 190, 196, 200, 204, 210, 214, 219, 223, 236, 240, 250,
		258, 280, 284, 288, 296, 311, 334, 340, 353, 361, 421, 441, 443, 451, 459,
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

// VlangParserInit initializes any static state used to implement VlangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVlangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VlangParserInit() {
	staticData := &VlangParserStaticData
	staticData.once.Do(vlangParserInit)
}

// NewVlangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVlangParser(input antlr.TokenStream) *VlangParser {
	VlangParserInit()
	this := new(VlangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &VlangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Vlang.g4"

	return this
}

// VlangParser tokens.
const (
	VlangParserEOF           = antlr.TokenEOF
	VlangParserT__0          = 1
	VlangParserT__1          = 2
	VlangParserT__2          = 3
	VlangParserT__3          = 4
	VlangParserT__4          = 5
	VlangParserMUT           = 6
	VlangParserSTRUCT        = 7
	VlangParserFUNC          = 8
	VlangParserPRINT         = 9
	VlangParserIF            = 10
	VlangParserELSE          = 11
	VlangParserFOR           = 12
	VlangParserWHILE         = 13
	VlangParserSWITCH        = 14
	VlangParserCASE          = 15
	VlangParserDEFAULT       = 16
	VlangParserBREAK         = 17
	VlangParserCONTINUE      = 18
	VlangParserRETURN        = 19
	VlangParserCOLON         = 20
	VlangParserLEN           = 21
	VlangParserCAP           = 22
	VlangParserINDEXOF       = 23
	VlangParserJOIN          = 24
	VlangParserAPPEND        = 25
	VlangParserIN            = 26
	VlangParserBOOLEANO      = 27
	VlangParserENTERO        = 28
	VlangParserDECIMAL       = 29
	VlangParserCADENA        = 30
	VlangParserCARACTER      = 31
	VlangParserID            = 32
	VlangParserPLUS          = 33
	VlangParserMINUS         = 34
	VlangParserMUL           = 35
	VlangParserDIV           = 36
	VlangParserMOD           = 37
	VlangParserNOT           = 38
	VlangParserAND           = 39
	VlangParserOR            = 40
	VlangParserEQ            = 41
	VlangParserNEQ           = 42
	VlangParserLT            = 43
	VlangParserLE            = 44
	VlangParserGT            = 45
	VlangParserGE            = 46
	VlangParserASSIGN        = 47
	VlangParserASSIGN_DECL   = 48
	VlangParserPLUSEQ        = 49
	VlangParserMINUSEQ       = 50
	VlangParserINC           = 51
	VlangParserDEC           = 52
	VlangParserLPAREN        = 53
	VlangParserRPAREN        = 54
	VlangParserLBRACE        = 55
	VlangParserRBRACE        = 56
	VlangParserLSQUARE       = 57
	VlangParserRSQUARE       = 58
	VlangParserDOT           = 59
	VlangParserCOMMA         = 60
	VlangParserSEMICOLON     = 61
	VlangParserWS            = 62
	VlangParserLINE_COMMENT  = 63
	VlangParserBLOCK_COMMENT = 64
)

// VlangParser rules.
const (
	VlangParserRULE_programa            = 0
	VlangParserRULE_declaraciones       = 1
	VlangParserRULE_structDecl          = 2
	VlangParserRULE_atributos           = 3
	VlangParserRULE_atributo            = 4
	VlangParserRULE_inicializadorStruct = 5
	VlangParserRULE_inicializadorCampo  = 6
	VlangParserRULE_funcionStruct       = 7
	VlangParserRULE_funcion             = 8
	VlangParserRULE_parametrosFormales  = 9
	VlangParserRULE_parametroFormal     = 10
	VlangParserRULE_tipoConReferencia   = 11
	VlangParserRULE_varDcl              = 12
	VlangParserRULE_tipo                = 13
	VlangParserRULE_sliceTipo           = 14
	VlangParserRULE_stmt                = 15
	VlangParserRULE_sentencias_control  = 16
	VlangParserRULE_ifDcl               = 17
	VlangParserRULE_elseifDcl           = 18
	VlangParserRULE_elseDcl             = 19
	VlangParserRULE_switchDcl           = 20
	VlangParserRULE_caseBlock           = 21
	VlangParserRULE_defaultBlock        = 22
	VlangParserRULE_forDcl              = 23
	VlangParserRULE_whileDcl            = 24
	VlangParserRULE_bloque              = 25
	VlangParserRULE_expresion           = 26
	VlangParserRULE_parametros          = 27
	VlangParserRULE_valor               = 28
)

// IProgramaContext is an interface to support dynamic dispatch.
type IProgramaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllDeclaraciones() []IDeclaracionesContext
	Declaraciones(i int) IDeclaracionesContext

	// IsProgramaContext differentiates from other interfaces.
	IsProgramaContext()
}

type ProgramaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramaContext() *ProgramaContext {
	var p = new(ProgramaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_programa
	return p
}

func InitEmptyProgramaContext(p *ProgramaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_programa
}

func (*ProgramaContext) IsProgramaContext() {}

func NewProgramaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramaContext {
	var p = new(ProgramaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_programa

	return p
}

func (s *ProgramaContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramaContext) EOF() antlr.TerminalNode {
	return s.GetToken(VlangParserEOF, 0)
}

func (s *ProgramaContext) AllDeclaraciones() []IDeclaracionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracionesContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracionesContext); ok {
			tst[i] = t.(IDeclaracionesContext)
			i++
		}
	}

	return tst
}

func (s *ProgramaContext) Declaraciones(i int) IDeclaracionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionesContext); ok {
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

	return t.(IDeclaracionesContext)
}

func (s *ProgramaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterPrograma(s)
	}
}

func (s *ProgramaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitPrograma(s)
	}
}

func (s *ProgramaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitPrograma(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Programa() (localctx IProgramaContext) {
	localctx = NewProgramaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VlangParserRULE_programa)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&189151484924819392) != 0 {
		{
			p.SetState(58)
			p.Declaraciones()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(64)
		p.Match(VlangParserEOF)
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

// IDeclaracionesContext is an interface to support dynamic dispatch.
type IDeclaracionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDcl() IVarDclContext
	StructDecl() IStructDeclContext
	FuncionStruct() IFuncionStructContext
	Funcion() IFuncionContext
	Stmt() IStmtContext

	// IsDeclaracionesContext differentiates from other interfaces.
	IsDeclaracionesContext()
}

type DeclaracionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracionesContext() *DeclaracionesContext {
	var p = new(DeclaracionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_declaraciones
	return p
}

func InitEmptyDeclaracionesContext(p *DeclaracionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_declaraciones
}

func (*DeclaracionesContext) IsDeclaracionesContext() {}

func NewDeclaracionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionesContext {
	var p = new(DeclaracionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_declaraciones

	return p
}

func (s *DeclaracionesContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionesContext) VarDcl() IVarDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDclContext)
}

func (s *DeclaracionesContext) StructDecl() IStructDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclContext)
}

func (s *DeclaracionesContext) FuncionStruct() IFuncionStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncionStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncionStructContext)
}

func (s *DeclaracionesContext) Funcion() IFuncionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncionContext)
}

func (s *DeclaracionesContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *DeclaracionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDeclaraciones(s)
	}
}

func (s *DeclaracionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDeclaraciones(s)
	}
}

func (s *DeclaracionesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDeclaraciones(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Declaraciones() (localctx IDeclaracionesContext) {
	localctx = NewDeclaracionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VlangParserRULE_declaraciones)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.VarDcl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.StructDecl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)
			p.FuncionStruct()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(69)
			p.Funcion()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(70)
			p.Stmt()
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

// IStructDeclContext is an interface to support dynamic dispatch.
type IStructDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStructDeclContext differentiates from other interfaces.
	IsStructDeclContext()
}

type StructDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclContext() *StructDeclContext {
	var p = new(StructDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_structDecl
	return p
}

func InitEmptyStructDeclContext(p *StructDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_structDecl
}

func (*StructDeclContext) IsStructDeclContext() {}

func NewStructDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclContext {
	var p = new(StructDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_structDecl

	return p
}

func (s *StructDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclContext) CopyAll(ctx *StructDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StructDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructDeclarationContext struct {
	StructDeclContext
}

func NewStructDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDeclarationContext {
	var p = new(StructDeclarationContext)

	InitEmptyStructDeclContext(&p.StructDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*StructDeclContext))

	return p
}

func (s *StructDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclarationContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(VlangParserSTRUCT, 0)
}

func (s *StructDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *StructDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACE, 0)
}

func (s *StructDeclarationContext) Atributos() IAtributosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtributosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtributosContext)
}

func (s *StructDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACE, 0)
}

func (s *StructDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStructDeclaration(s)
	}
}

func (s *StructDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStructDeclaration(s)
	}
}

func (s *StructDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStructDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) StructDecl() (localctx IStructDeclContext) {
	localctx = NewStructDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VlangParserRULE_structDecl)
	localctx = NewStructDeclarationContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Match(VlangParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Match(VlangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Atributos()
	}
	{
		p.SetState(77)
		p.Match(VlangParserRBRACE)
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

// IAtributosContext is an interface to support dynamic dispatch.
type IAtributosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAtributo() []IAtributoContext
	Atributo(i int) IAtributoContext

	// IsAtributosContext differentiates from other interfaces.
	IsAtributosContext()
}

type AtributosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtributosContext() *AtributosContext {
	var p = new(AtributosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_atributos
	return p
}

func InitEmptyAtributosContext(p *AtributosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_atributos
}

func (*AtributosContext) IsAtributosContext() {}

func NewAtributosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtributosContext {
	var p = new(AtributosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_atributos

	return p
}

func (s *AtributosContext) GetParser() antlr.Parser { return s.parser }

func (s *AtributosContext) AllAtributo() []IAtributoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAtributoContext); ok {
			len++
		}
	}

	tst := make([]IAtributoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAtributoContext); ok {
			tst[i] = t.(IAtributoContext)
			i++
		}
	}

	return tst
}

func (s *AtributosContext) Atributo(i int) IAtributoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtributoContext); ok {
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

	return t.(IAtributoContext)
}

func (s *AtributosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtributosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtributosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAtributos(s)
	}
}

func (s *AtributosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAtributos(s)
	}
}

func (s *AtributosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAtributos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Atributos() (localctx IAtributosContext) {
	localctx = NewAtributosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VlangParserRULE_atributos)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144115188075855934) != 0) {
		{
			p.SetState(79)
			p.Atributo()
		}

		p.SetState(82)
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

// IAtributoContext is an interface to support dynamic dispatch.
type IAtributoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tipo() ITipoContext
	ID() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsAtributoContext differentiates from other interfaces.
	IsAtributoContext()
}

type AtributoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtributoContext() *AtributoContext {
	var p = new(AtributoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_atributo
	return p
}

func InitEmptyAtributoContext(p *AtributoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_atributo
}

func (*AtributoContext) IsAtributoContext() {}

func NewAtributoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtributoContext {
	var p = new(AtributoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_atributo

	return p
}

func (s *AtributoContext) GetParser() antlr.Parser { return s.parser }

func (s *AtributoContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *AtributoContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *AtributoContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(VlangParserSEMICOLON, 0)
}

func (s *AtributoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtributoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtributoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAtributo(s)
	}
}

func (s *AtributoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAtributo(s)
	}
}

func (s *AtributoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAtributo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Atributo() (localctx IAtributoContext) {
	localctx = NewAtributoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VlangParserRULE_atributo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Tipo()
	}
	{
		p.SetState(85)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Match(VlangParserSEMICOLON)
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

// IInicializadorStructContext is an interface to support dynamic dispatch.
type IInicializadorStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInicializadorCampo() []IInicializadorCampoContext
	InicializadorCampo(i int) IInicializadorCampoContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsInicializadorStructContext differentiates from other interfaces.
	IsInicializadorStructContext()
}

type InicializadorStructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInicializadorStructContext() *InicializadorStructContext {
	var p = new(InicializadorStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_inicializadorStruct
	return p
}

func InitEmptyInicializadorStructContext(p *InicializadorStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_inicializadorStruct
}

func (*InicializadorStructContext) IsInicializadorStructContext() {}

func NewInicializadorStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InicializadorStructContext {
	var p = new(InicializadorStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_inicializadorStruct

	return p
}

func (s *InicializadorStructContext) GetParser() antlr.Parser { return s.parser }

func (s *InicializadorStructContext) AllInicializadorCampo() []IInicializadorCampoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInicializadorCampoContext); ok {
			len++
		}
	}

	tst := make([]IInicializadorCampoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInicializadorCampoContext); ok {
			tst[i] = t.(IInicializadorCampoContext)
			i++
		}
	}

	return tst
}

func (s *InicializadorStructContext) InicializadorCampo(i int) IInicializadorCampoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInicializadorCampoContext); ok {
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

	return t.(IInicializadorCampoContext)
}

func (s *InicializadorStructContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *InicializadorStructContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *InicializadorStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InicializadorStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InicializadorStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterInicializadorStruct(s)
	}
}

func (s *InicializadorStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitInicializadorStruct(s)
	}
}

func (s *InicializadorStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitInicializadorStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) InicializadorStruct() (localctx IInicializadorStructContext) {
	localctx = NewInicializadorStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VlangParserRULE_inicializadorStruct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.InicializadorCampo()
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(89)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.InicializadorCampo()
		}

		p.SetState(95)
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

// IInicializadorCampoContext is an interface to support dynamic dispatch.
type IInicializadorCampoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Expresion() IExpresionContext

	// IsInicializadorCampoContext differentiates from other interfaces.
	IsInicializadorCampoContext()
}

type InicializadorCampoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInicializadorCampoContext() *InicializadorCampoContext {
	var p = new(InicializadorCampoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_inicializadorCampo
	return p
}

func InitEmptyInicializadorCampoContext(p *InicializadorCampoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_inicializadorCampo
}

func (*InicializadorCampoContext) IsInicializadorCampoContext() {}

func NewInicializadorCampoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InicializadorCampoContext {
	var p = new(InicializadorCampoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_inicializadorCampo

	return p
}

func (s *InicializadorCampoContext) GetParser() antlr.Parser { return s.parser }

func (s *InicializadorCampoContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *InicializadorCampoContext) COLON() antlr.TerminalNode {
	return s.GetToken(VlangParserCOLON, 0)
}

func (s *InicializadorCampoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *InicializadorCampoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InicializadorCampoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InicializadorCampoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterInicializadorCampo(s)
	}
}

func (s *InicializadorCampoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitInicializadorCampo(s)
	}
}

func (s *InicializadorCampoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitInicializadorCampo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) InicializadorCampo() (localctx IInicializadorCampoContext) {
	localctx = NewInicializadorCampoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VlangParserRULE_inicializadorCampo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)
		p.Match(VlangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.expresion(0)
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

// IFuncionStructContext is an interface to support dynamic dispatch.
type IFuncionStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode
	Bloque() IBloqueContext
	MUL() antlr.TerminalNode
	ParametrosFormales() IParametrosFormalesContext

	// IsFuncionStructContext differentiates from other interfaces.
	IsFuncionStructContext()
}

type FuncionStructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncionStructContext() *FuncionStructContext {
	var p = new(FuncionStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_funcionStruct
	return p
}

func InitEmptyFuncionStructContext(p *FuncionStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_funcionStruct
}

func (*FuncionStructContext) IsFuncionStructContext() {}

func NewFuncionStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionStructContext {
	var p = new(FuncionStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_funcionStruct

	return p
}

func (s *FuncionStructContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionStructContext) FUNC() antlr.TerminalNode {
	return s.GetToken(VlangParserFUNC, 0)
}

func (s *FuncionStructContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(VlangParserLPAREN)
}

func (s *FuncionStructContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, i)
}

func (s *FuncionStructContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *FuncionStructContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *FuncionStructContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(VlangParserRPAREN)
}

func (s *FuncionStructContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, i)
}

func (s *FuncionStructContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *FuncionStructContext) MUL() antlr.TerminalNode {
	return s.GetToken(VlangParserMUL, 0)
}

func (s *FuncionStructContext) ParametrosFormales() IParametrosFormalesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosFormalesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosFormalesContext)
}

func (s *FuncionStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFuncionStruct(s)
	}
}

func (s *FuncionStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFuncionStruct(s)
	}
}

func (s *FuncionStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFuncionStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) FuncionStruct() (localctx IFuncionStructContext) {
	localctx = NewFuncionStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VlangParserRULE_funcionStruct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(VlangParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserMUL {
		{
			p.SetState(103)
			p.Match(VlangParserMUL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(106)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(VlangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144115222435594302) != 0 {
		{
			p.SetState(110)
			p.ParametrosFormales()
		}

	}
	{
		p.SetState(113)
		p.Match(VlangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.Bloque()
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

// IFuncionContext is an interface to support dynamic dispatch.
type IFuncionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Bloque() IBloqueContext
	ParametrosFormales() IParametrosFormalesContext
	Tipo() ITipoContext

	// IsFuncionContext differentiates from other interfaces.
	IsFuncionContext()
}

type FuncionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncionContext() *FuncionContext {
	var p = new(FuncionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_funcion
	return p
}

func InitEmptyFuncionContext(p *FuncionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_funcion
}

func (*FuncionContext) IsFuncionContext() {}

func NewFuncionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionContext {
	var p = new(FuncionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_funcion

	return p
}

func (s *FuncionContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionContext) FUNC() antlr.TerminalNode {
	return s.GetToken(VlangParserFUNC, 0)
}

func (s *FuncionContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *FuncionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *FuncionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *FuncionContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *FuncionContext) ParametrosFormales() IParametrosFormalesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosFormalesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosFormalesContext)
}

func (s *FuncionContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *FuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFuncion(s)
	}
}

func (s *FuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFuncion(s)
	}
}

func (s *FuncionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFuncion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Funcion() (localctx IFuncionContext) {
	localctx = NewFuncionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VlangParserRULE_funcion)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(VlangParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144115222435594302) != 0 {
		{
			p.SetState(119)
			p.ParametrosFormales()
		}

	}
	{
		p.SetState(122)
		p.Match(VlangParserRPAREN)
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

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144115188075855934) != 0 {
		{
			p.SetState(123)
			p.Tipo()
		}

	}
	{
		p.SetState(126)
		p.Bloque()
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

// IParametrosFormalesContext is an interface to support dynamic dispatch.
type IParametrosFormalesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParametroFormal() []IParametroFormalContext
	ParametroFormal(i int) IParametroFormalContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParametrosFormalesContext differentiates from other interfaces.
	IsParametrosFormalesContext()
}

type ParametrosFormalesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametrosFormalesContext() *ParametrosFormalesContext {
	var p = new(ParametrosFormalesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametrosFormales
	return p
}

func InitEmptyParametrosFormalesContext(p *ParametrosFormalesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametrosFormales
}

func (*ParametrosFormalesContext) IsParametrosFormalesContext() {}

func NewParametrosFormalesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosFormalesContext {
	var p = new(ParametrosFormalesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_parametrosFormales

	return p
}

func (s *ParametrosFormalesContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosFormalesContext) AllParametroFormal() []IParametroFormalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroFormalContext); ok {
			len++
		}
	}

	tst := make([]IParametroFormalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroFormalContext); ok {
			tst[i] = t.(IParametroFormalContext)
			i++
		}
	}

	return tst
}

func (s *ParametrosFormalesContext) ParametroFormal(i int) IParametroFormalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroFormalContext); ok {
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

	return t.(IParametroFormalContext)
}

func (s *ParametrosFormalesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *ParametrosFormalesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *ParametrosFormalesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosFormalesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosFormalesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterParametrosFormales(s)
	}
}

func (s *ParametrosFormalesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitParametrosFormales(s)
	}
}

func (s *ParametrosFormalesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitParametrosFormales(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) ParametrosFormales() (localctx IParametrosFormalesContext) {
	localctx = NewParametrosFormalesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, VlangParserRULE_parametrosFormales)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.ParametroFormal()
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(129)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.ParametroFormal()
		}

		p.SetState(135)
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

// IParametroFormalContext is an interface to support dynamic dispatch.
type IParametroFormalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTipoRef returns the tipoRef token.
	GetTipoRef() antlr.Token

	// SetTipoRef sets the tipoRef token.
	SetTipoRef(antlr.Token)

	// Getter signatures
	TipoConReferencia() ITipoConReferenciaContext
	ID() antlr.TerminalNode

	// IsParametroFormalContext differentiates from other interfaces.
	IsParametroFormalContext()
}

type ParametroFormalContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	tipoRef antlr.Token
}

func NewEmptyParametroFormalContext() *ParametroFormalContext {
	var p = new(ParametroFormalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametroFormal
	return p
}

func InitEmptyParametroFormalContext(p *ParametroFormalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametroFormal
}

func (*ParametroFormalContext) IsParametroFormalContext() {}

func NewParametroFormalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametroFormalContext {
	var p = new(ParametroFormalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_parametroFormal

	return p
}

func (s *ParametroFormalContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametroFormalContext) GetTipoRef() antlr.Token { return s.tipoRef }

func (s *ParametroFormalContext) SetTipoRef(v antlr.Token) { s.tipoRef = v }

func (s *ParametroFormalContext) TipoConReferencia() ITipoConReferenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoConReferenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoConReferenciaContext)
}

func (s *ParametroFormalContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *ParametroFormalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametroFormalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametroFormalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterParametroFormal(s)
	}
}

func (s *ParametroFormalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitParametroFormal(s)
	}
}

func (s *ParametroFormalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitParametroFormal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) ParametroFormal() (localctx IParametroFormalContext) {
	localctx = NewParametroFormalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, VlangParserRULE_parametroFormal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.TipoConReferencia()
	}
	{
		p.SetState(137)

		var _m = p.Match(VlangParserID)

		localctx.(*ParametroFormalContext).tipoRef = _m
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

// ITipoConReferenciaContext is an interface to support dynamic dispatch.
type ITipoConReferenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRef returns the ref token.
	GetRef() antlr.Token

	// SetRef sets the ref token.
	SetRef(antlr.Token)

	// GetTipoTipo returns the tipoTipo rule contexts.
	GetTipoTipo() ITipoContext

	// SetTipoTipo sets the tipoTipo rule contexts.
	SetTipoTipo(ITipoContext)

	// Getter signatures
	Tipo() ITipoContext
	MUL() antlr.TerminalNode

	// IsTipoConReferenciaContext differentiates from other interfaces.
	IsTipoConReferenciaContext()
}

type TipoConReferenciaContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	ref      antlr.Token
	tipoTipo ITipoContext
}

func NewEmptyTipoConReferenciaContext() *TipoConReferenciaContext {
	var p = new(TipoConReferenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_tipoConReferencia
	return p
}

func InitEmptyTipoConReferenciaContext(p *TipoConReferenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_tipoConReferencia
}

func (*TipoConReferenciaContext) IsTipoConReferenciaContext() {}

func NewTipoConReferenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoConReferenciaContext {
	var p = new(TipoConReferenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_tipoConReferencia

	return p
}

func (s *TipoConReferenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoConReferenciaContext) GetRef() antlr.Token { return s.ref }

func (s *TipoConReferenciaContext) SetRef(v antlr.Token) { s.ref = v }

func (s *TipoConReferenciaContext) GetTipoTipo() ITipoContext { return s.tipoTipo }

func (s *TipoConReferenciaContext) SetTipoTipo(v ITipoContext) { s.tipoTipo = v }

func (s *TipoConReferenciaContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *TipoConReferenciaContext) MUL() antlr.TerminalNode {
	return s.GetToken(VlangParserMUL, 0)
}

func (s *TipoConReferenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoConReferenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoConReferenciaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterTipoConReferencia(s)
	}
}

func (s *TipoConReferenciaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitTipoConReferencia(s)
	}
}

func (s *TipoConReferenciaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitTipoConReferencia(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) TipoConReferencia() (localctx ITipoConReferenciaContext) {
	localctx = NewTipoConReferenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, VlangParserRULE_tipoConReferencia)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserMUL {
		{
			p.SetState(139)

			var _m = p.Match(VlangParserMUL)

			localctx.(*TipoConReferenciaContext).ref = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(142)

		var _x = p.Tipo()

		localctx.(*TipoConReferenciaContext).tipoTipo = _x
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

// IVarDclContext is an interface to support dynamic dispatch.
type IVarDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVarDclContext differentiates from other interfaces.
	IsVarDclContext()
}

type VarDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDclContext() *VarDclContext {
	var p = new(VarDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_varDcl
	return p
}

func InitEmptyVarDclContext(p *VarDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_varDcl
}

func (*VarDclContext) IsVarDclContext() {}

func NewVarDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDclContext {
	var p = new(VarDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_varDcl

	return p
}

func (s *VarDclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDclContext) CopyAll(ctx *VarDclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VarDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VariableDeclarationContext struct {
	VarDclContext
}

func NewVariableDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	InitEmptyVarDclContext(&p.VarDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDclContext))

	return p
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) MUT() antlr.TerminalNode {
	return s.GetToken(VlangParserMUT, 0)
}

func (s *VariableDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *VariableDeclarationContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *VariableDeclarationContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *VariableDeclarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *VariableDeclarationContext) ASSIGN_DECL() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN_DECL, 0)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) VarDcl() (localctx IVarDclContext) {
	localctx = NewVarDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, VlangParserRULE_varDcl)
	var _la int

	localctx = NewVariableDeclarationContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(VlangParserMUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(146)
			p.Tipo()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserASSIGN || _la == VlangParserASSIGN_DECL {
		{
			p.SetState(149)
			_la = p.GetTokenStream().LA(1)

			if !(_la == VlangParserASSIGN || _la == VlangParserASSIGN_DECL) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(150)
			p.expresion(0)
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

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SliceTipo() ISliceTipoContext

	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_tipo
	return p
}

func InitEmptyTipoContext(p *TipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_tipo
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) SliceTipo() ISliceTipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceTipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceTipoContext)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterTipo(s)
	}
}

func (s *TipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitTipo(s)
	}
}

func (s *TipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitTipo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, VlangParserRULE_tipo)
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserLSQUARE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.SliceTipo()
		}

	case VlangParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.Match(VlangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserT__1:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(155)
			p.Match(VlangParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserT__2:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(156)
			p.Match(VlangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserT__3:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(157)
			p.Match(VlangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserT__4:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(158)
			p.Match(VlangParserT__4)
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

// ISliceTipoContext is an interface to support dynamic dispatch.
type ISliceTipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LSQUARE() antlr.TerminalNode
	RSQUARE() antlr.TerminalNode
	Tipo() ITipoContext

	// IsSliceTipoContext differentiates from other interfaces.
	IsSliceTipoContext()
}

type SliceTipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceTipoContext() *SliceTipoContext {
	var p = new(SliceTipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_sliceTipo
	return p
}

func InitEmptySliceTipoContext(p *SliceTipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_sliceTipo
}

func (*SliceTipoContext) IsSliceTipoContext() {}

func NewSliceTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceTipoContext {
	var p = new(SliceTipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_sliceTipo

	return p
}

func (s *SliceTipoContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceTipoContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(VlangParserLSQUARE, 0)
}

func (s *SliceTipoContext) RSQUARE() antlr.TerminalNode {
	return s.GetToken(VlangParserRSQUARE, 0)
}

func (s *SliceTipoContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *SliceTipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceTipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceTipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSliceTipo(s)
	}
}

func (s *SliceTipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSliceTipo(s)
	}
}

func (s *SliceTipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSliceTipo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) SliceTipo() (localctx ISliceTipoContext) {
	localctx = NewSliceTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, VlangParserRULE_sliceTipo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(VlangParserLSQUARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Match(VlangParserRSQUARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Tipo()
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

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) CopyAll(ctx *StmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrintStatementContext struct {
	StmtContext
}

func NewPrintStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintStatementContext {
	var p = new(PrintStatementContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *PrintStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStatementContext) PRINT() antlr.TerminalNode {
	return s.GetToken(VlangParserPRINT, 0)
}

func (s *PrintStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *PrintStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *PrintStatementContext) Parametros() IParametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosContext)
}

func (s *PrintStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterPrintStatement(s)
	}
}

func (s *PrintStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitPrintStatement(s)
	}
}

func (s *PrintStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitPrintStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type BloqueStatementContext struct {
	StmtContext
}

func NewBloqueStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BloqueStatementContext {
	var p = new(BloqueStatementContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *BloqueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueStatementContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *BloqueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterBloqueStatement(s)
	}
}

func (s *BloqueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitBloqueStatement(s)
	}
}

func (s *BloqueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitBloqueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ControlStatementContext struct {
	StmtContext
}

func NewControlStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ControlStatementContext {
	var p = new(ControlStatementContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *ControlStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlStatementContext) Sentencias_control() ISentencias_controlContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentencias_controlContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentencias_controlContext)
}

func (s *ControlStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterControlStatement(s)
	}
}

func (s *ControlStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitControlStatement(s)
	}
}

func (s *ControlStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitControlStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceDobleAsignacionStmtContext struct {
	StmtContext
}

func NewSliceDobleAsignacionStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceDobleAsignacionStmtContext {
	var p = new(SliceDobleAsignacionStmtContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *SliceDobleAsignacionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceDobleAsignacionStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *SliceDobleAsignacionStmtContext) AllLSQUARE() []antlr.TerminalNode {
	return s.GetTokens(VlangParserLSQUARE)
}

func (s *SliceDobleAsignacionStmtContext) LSQUARE(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserLSQUARE, i)
}

func (s *SliceDobleAsignacionStmtContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *SliceDobleAsignacionStmtContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *SliceDobleAsignacionStmtContext) AllRSQUARE() []antlr.TerminalNode {
	return s.GetTokens(VlangParserRSQUARE)
}

func (s *SliceDobleAsignacionStmtContext) RSQUARE(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserRSQUARE, i)
}

func (s *SliceDobleAsignacionStmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *SliceDobleAsignacionStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSliceDobleAsignacionStmt(s)
	}
}

func (s *SliceDobleAsignacionStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSliceDobleAsignacionStmt(s)
	}
}

func (s *SliceDobleAsignacionStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSliceDobleAsignacionStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpresionStatementContext struct {
	StmtContext
}

func NewExpresionStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpresionStatementContext {
	var p = new(ExpresionStatementContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *ExpresionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionStatementContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpresionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterExpresionStatement(s)
	}
}

func (s *ExpresionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitExpresionStatement(s)
	}
}

func (s *ExpresionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitExpresionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakStatementContext struct {
	StmtContext
}

func NewBreakStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStatementContext {
	var p = new(BreakStatementContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(VlangParserBREAK, 0)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContinueStatementContext struct {
	StmtContext
}

func NewContinueStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(VlangParserCONTINUE, 0)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStatementContext struct {
	StmtContext
}

func NewReturnStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(VlangParserRETURN, 0)
}

func (s *ReturnStatementContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, VlangParserRULE_stmt)
	var _la int

	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpresionStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.expresion(0)
		}

	case 2:
		localctx = NewSliceDobleAsignacionStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.Match(VlangParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.expresion(0)
		}
		{
			p.SetState(169)
			p.Match(VlangParserRSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.Match(VlangParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.expresion(0)
		}
		{
			p.SetState(172)
			p.Match(VlangParserRSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.expresion(0)
		}

	case 3:
		localctx = NewPrintStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(176)
			p.Match(VlangParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Match(VlangParserLPAREN)
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

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&189151484923871232) != 0 {
			{
				p.SetState(178)
				p.Parametros()
			}

		}
		{
			p.SetState(181)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewBreakStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(182)
			p.Match(VlangParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewContinueStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(183)
			p.Match(VlangParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewReturnStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(184)
			p.Match(VlangParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(185)
				p.expresion(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 7:
		localctx = NewControlStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(188)
			p.Sentencias_control()
		}

	case 8:
		localctx = NewBloqueStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(189)
			p.Bloque()
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

// ISentencias_controlContext is an interface to support dynamic dispatch.
type ISentencias_controlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSentencias_controlContext differentiates from other interfaces.
	IsSentencias_controlContext()
}

type Sentencias_controlContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentencias_controlContext() *Sentencias_controlContext {
	var p = new(Sentencias_controlContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_sentencias_control
	return p
}

func InitEmptySentencias_controlContext(p *Sentencias_controlContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_sentencias_control
}

func (*Sentencias_controlContext) IsSentencias_controlContext() {}

func NewSentencias_controlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sentencias_controlContext {
	var p = new(Sentencias_controlContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_sentencias_control

	return p
}

func (s *Sentencias_controlContext) GetParser() antlr.Parser { return s.parser }

func (s *Sentencias_controlContext) CopyAll(ctx *Sentencias_controlContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Sentencias_controlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sentencias_controlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type While_contextContext struct {
	Sentencias_controlContext
}

func NewWhile_contextContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *While_contextContext {
	var p = new(While_contextContext)

	InitEmptySentencias_controlContext(&p.Sentencias_controlContext)
	p.parser = parser
	p.CopyAll(ctx.(*Sentencias_controlContext))

	return p
}

func (s *While_contextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_contextContext) WhileDcl() IWhileDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileDclContext)
}

func (s *While_contextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterWhile_context(s)
	}
}

func (s *While_contextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitWhile_context(s)
	}
}

func (s *While_contextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitWhile_context(s)

	default:
		return t.VisitChildren(s)
	}
}

type Switch_contextContext struct {
	Sentencias_controlContext
}

func NewSwitch_contextContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Switch_contextContext {
	var p = new(Switch_contextContext)

	InitEmptySentencias_controlContext(&p.Sentencias_controlContext)
	p.parser = parser
	p.CopyAll(ctx.(*Sentencias_controlContext))

	return p
}

func (s *Switch_contextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_contextContext) SwitchDcl() ISwitchDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchDclContext)
}

func (s *Switch_contextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSwitch_context(s)
	}
}

func (s *Switch_contextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSwitch_context(s)
	}
}

func (s *Switch_contextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSwitch_context(s)

	default:
		return t.VisitChildren(s)
	}
}

type If_contextContext struct {
	Sentencias_controlContext
}

func NewIf_contextContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *If_contextContext {
	var p = new(If_contextContext)

	InitEmptySentencias_controlContext(&p.Sentencias_controlContext)
	p.parser = parser
	p.CopyAll(ctx.(*Sentencias_controlContext))

	return p
}

func (s *If_contextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_contextContext) IfDcl() IIfDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfDclContext)
}

func (s *If_contextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIf_context(s)
	}
}

func (s *If_contextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIf_context(s)
	}
}

func (s *If_contextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIf_context(s)

	default:
		return t.VisitChildren(s)
	}
}

type For_contextContext struct {
	Sentencias_controlContext
}

func NewFor_contextContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *For_contextContext {
	var p = new(For_contextContext)

	InitEmptySentencias_controlContext(&p.Sentencias_controlContext)
	p.parser = parser
	p.CopyAll(ctx.(*Sentencias_controlContext))

	return p
}

func (s *For_contextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_contextContext) ForDcl() IForDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForDclContext)
}

func (s *For_contextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFor_context(s)
	}
}

func (s *For_contextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFor_context(s)
	}
}

func (s *For_contextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFor_context(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Sentencias_control() (localctx ISentencias_controlContext) {
	localctx = NewSentencias_controlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, VlangParserRULE_sentencias_control)
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserIF:
		localctx = NewIf_contextContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.IfDcl()
		}

	case VlangParserFOR:
		localctx = NewFor_contextContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)
			p.ForDcl()
		}

	case VlangParserWHILE:
		localctx = NewWhile_contextContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(194)
			p.WhileDcl()
		}

	case VlangParserSWITCH:
		localctx = NewSwitch_contextContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(195)
			p.SwitchDcl()
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

// IIfDclContext is an interface to support dynamic dispatch.
type IIfDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expresion() IExpresionContext
	Bloque() IBloqueContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllElseifDcl() []IElseifDclContext
	ElseifDcl(i int) IElseifDclContext
	ElseDcl() IElseDclContext

	// IsIfDclContext differentiates from other interfaces.
	IsIfDclContext()
}

type IfDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfDclContext() *IfDclContext {
	var p = new(IfDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_ifDcl
	return p
}

func InitEmptyIfDclContext(p *IfDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_ifDcl
}

func (*IfDclContext) IsIfDclContext() {}

func NewIfDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfDclContext {
	var p = new(IfDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_ifDcl

	return p
}

func (s *IfDclContext) GetParser() antlr.Parser { return s.parser }

func (s *IfDclContext) IF() antlr.TerminalNode {
	return s.GetToken(VlangParserIF, 0)
}

func (s *IfDclContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *IfDclContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *IfDclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *IfDclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *IfDclContext) AllElseifDcl() []IElseifDclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseifDclContext); ok {
			len++
		}
	}

	tst := make([]IElseifDclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseifDclContext); ok {
			tst[i] = t.(IElseifDclContext)
			i++
		}
	}

	return tst
}

func (s *IfDclContext) ElseifDcl(i int) IElseifDclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseifDclContext); ok {
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

	return t.(IElseifDclContext)
}

func (s *IfDclContext) ElseDcl() IElseDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseDclContext)
}

func (s *IfDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfDclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIfDcl(s)
	}
}

func (s *IfDclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIfDcl(s)
	}
}

func (s *IfDclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIfDcl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) IfDcl() (localctx IIfDclContext) {
	localctx = NewIfDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, VlangParserRULE_ifDcl)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(VlangParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(199)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(202)
		p.expresion(0)
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserRPAREN {
		{
			p.SetState(203)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(206)
		p.Bloque()
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(207)
				p.ElseifDcl()
			}

		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserELSE {
		{
			p.SetState(213)
			p.ElseDcl()
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

// IElseifDclContext is an interface to support dynamic dispatch.
type IElseifDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	IF() antlr.TerminalNode
	Expresion() IExpresionContext
	Bloque() IBloqueContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsElseifDclContext differentiates from other interfaces.
	IsElseifDclContext()
}

type ElseifDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseifDclContext() *ElseifDclContext {
	var p = new(ElseifDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_elseifDcl
	return p
}

func InitEmptyElseifDclContext(p *ElseifDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_elseifDcl
}

func (*ElseifDclContext) IsElseifDclContext() {}

func NewElseifDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseifDclContext {
	var p = new(ElseifDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_elseifDcl

	return p
}

func (s *ElseifDclContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseifDclContext) ELSE() antlr.TerminalNode {
	return s.GetToken(VlangParserELSE, 0)
}

func (s *ElseifDclContext) IF() antlr.TerminalNode {
	return s.GetToken(VlangParserIF, 0)
}

func (s *ElseifDclContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ElseifDclContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *ElseifDclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *ElseifDclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *ElseifDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseifDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseifDclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterElseifDcl(s)
	}
}

func (s *ElseifDclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitElseifDcl(s)
	}
}

func (s *ElseifDclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitElseifDcl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) ElseifDcl() (localctx IElseifDclContext) {
	localctx = NewElseifDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, VlangParserRULE_elseifDcl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Match(VlangParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.Match(VlangParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(218)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(221)
		p.expresion(0)
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserRPAREN {
		{
			p.SetState(222)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(225)
		p.Bloque()
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

// IElseDclContext is an interface to support dynamic dispatch.
type IElseDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	Bloque() IBloqueContext

	// IsElseDclContext differentiates from other interfaces.
	IsElseDclContext()
}

type ElseDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseDclContext() *ElseDclContext {
	var p = new(ElseDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_elseDcl
	return p
}

func InitEmptyElseDclContext(p *ElseDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_elseDcl
}

func (*ElseDclContext) IsElseDclContext() {}

func NewElseDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseDclContext {
	var p = new(ElseDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_elseDcl

	return p
}

func (s *ElseDclContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseDclContext) ELSE() antlr.TerminalNode {
	return s.GetToken(VlangParserELSE, 0)
}

func (s *ElseDclContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *ElseDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseDclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterElseDcl(s)
	}
}

func (s *ElseDclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitElseDcl(s)
	}
}

func (s *ElseDclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitElseDcl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) ElseDcl() (localctx IElseDclContext) {
	localctx = NewElseDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, VlangParserRULE_elseDcl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Match(VlangParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
		p.Bloque()
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

// ISwitchDclContext is an interface to support dynamic dispatch.
type ISwitchDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expresion() IExpresionContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllCaseBlock() []ICaseBlockContext
	CaseBlock(i int) ICaseBlockContext
	DefaultBlock() IDefaultBlockContext

	// IsSwitchDclContext differentiates from other interfaces.
	IsSwitchDclContext()
}

type SwitchDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchDclContext() *SwitchDclContext {
	var p = new(SwitchDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_switchDcl
	return p
}

func InitEmptySwitchDclContext(p *SwitchDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_switchDcl
}

func (*SwitchDclContext) IsSwitchDclContext() {}

func NewSwitchDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchDclContext {
	var p = new(SwitchDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_switchDcl

	return p
}

func (s *SwitchDclContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchDclContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(VlangParserSWITCH, 0)
}

func (s *SwitchDclContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SwitchDclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACE, 0)
}

func (s *SwitchDclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACE, 0)
}

func (s *SwitchDclContext) AllCaseBlock() []ICaseBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICaseBlockContext); ok {
			len++
		}
	}

	tst := make([]ICaseBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICaseBlockContext); ok {
			tst[i] = t.(ICaseBlockContext)
			i++
		}
	}

	return tst
}

func (s *SwitchDclContext) CaseBlock(i int) ICaseBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseBlockContext); ok {
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

	return t.(ICaseBlockContext)
}

func (s *SwitchDclContext) DefaultBlock() IDefaultBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultBlockContext)
}

func (s *SwitchDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchDclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSwitchDcl(s)
	}
}

func (s *SwitchDclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSwitchDcl(s)
	}
}

func (s *SwitchDclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSwitchDcl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) SwitchDcl() (localctx ISwitchDclContext) {
	localctx = NewSwitchDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, VlangParserRULE_switchDcl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(VlangParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
		p.expresion(0)
	}
	{
		p.SetState(232)
		p.Match(VlangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCASE {
		{
			p.SetState(233)
			p.CaseBlock()
		}

		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserDEFAULT {
		{
			p.SetState(239)
			p.DefaultBlock()
		}

	}
	{
		p.SetState(242)
		p.Match(VlangParserRBRACE)
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

// ICaseBlockContext is an interface to support dynamic dispatch.
type ICaseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Expresion() IExpresionContext
	COLON() antlr.TerminalNode
	AllDeclaraciones() []IDeclaracionesContext
	Declaraciones(i int) IDeclaracionesContext

	// IsCaseBlockContext differentiates from other interfaces.
	IsCaseBlockContext()
}

type CaseBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseBlockContext() *CaseBlockContext {
	var p = new(CaseBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_caseBlock
	return p
}

func InitEmptyCaseBlockContext(p *CaseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_caseBlock
}

func (*CaseBlockContext) IsCaseBlockContext() {}

func NewCaseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseBlockContext {
	var p = new(CaseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_caseBlock

	return p
}

func (s *CaseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseBlockContext) CASE() antlr.TerminalNode {
	return s.GetToken(VlangParserCASE, 0)
}

func (s *CaseBlockContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *CaseBlockContext) COLON() antlr.TerminalNode {
	return s.GetToken(VlangParserCOLON, 0)
}

func (s *CaseBlockContext) AllDeclaraciones() []IDeclaracionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracionesContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracionesContext); ok {
			tst[i] = t.(IDeclaracionesContext)
			i++
		}
	}

	return tst
}

func (s *CaseBlockContext) Declaraciones(i int) IDeclaracionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionesContext); ok {
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

	return t.(IDeclaracionesContext)
}

func (s *CaseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterCaseBlock(s)
	}
}

func (s *CaseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitCaseBlock(s)
	}
}

func (s *CaseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitCaseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) CaseBlock() (localctx ICaseBlockContext) {
	localctx = NewCaseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, VlangParserRULE_caseBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		p.Match(VlangParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.expresion(0)
	}
	{
		p.SetState(246)
		p.Match(VlangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&189151484924819392) != 0 {
		{
			p.SetState(247)
			p.Declaraciones()
		}

		p.SetState(252)
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

// IDefaultBlockContext is an interface to support dynamic dispatch.
type IDefaultBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	COLON() antlr.TerminalNode
	AllDeclaraciones() []IDeclaracionesContext
	Declaraciones(i int) IDeclaracionesContext

	// IsDefaultBlockContext differentiates from other interfaces.
	IsDefaultBlockContext()
}

type DefaultBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultBlockContext() *DefaultBlockContext {
	var p = new(DefaultBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_defaultBlock
	return p
}

func InitEmptyDefaultBlockContext(p *DefaultBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_defaultBlock
}

func (*DefaultBlockContext) IsDefaultBlockContext() {}

func NewDefaultBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultBlockContext {
	var p = new(DefaultBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_defaultBlock

	return p
}

func (s *DefaultBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultBlockContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(VlangParserDEFAULT, 0)
}

func (s *DefaultBlockContext) COLON() antlr.TerminalNode {
	return s.GetToken(VlangParserCOLON, 0)
}

func (s *DefaultBlockContext) AllDeclaraciones() []IDeclaracionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracionesContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracionesContext); ok {
			tst[i] = t.(IDeclaracionesContext)
			i++
		}
	}

	return tst
}

func (s *DefaultBlockContext) Declaraciones(i int) IDeclaracionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionesContext); ok {
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

	return t.(IDeclaracionesContext)
}

func (s *DefaultBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDefaultBlock(s)
	}
}

func (s *DefaultBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDefaultBlock(s)
	}
}

func (s *DefaultBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDefaultBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) DefaultBlock() (localctx IDefaultBlockContext) {
	localctx = NewDefaultBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, VlangParserRULE_defaultBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Match(VlangParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(254)
		p.Match(VlangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&189151484924819392) != 0 {
		{
			p.SetState(255)
			p.Declaraciones()
		}

		p.SetState(260)
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

// IForDclContext is an interface to support dynamic dispatch.
type IForDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsForDclContext differentiates from other interfaces.
	IsForDclContext()
}

type ForDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForDclContext() *ForDclContext {
	var p = new(ForDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_forDcl
	return p
}

func InitEmptyForDclContext(p *ForDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_forDcl
}

func (*ForDclContext) IsForDclContext() {}

func NewForDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForDclContext {
	var p = new(ForDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_forDcl

	return p
}

func (s *ForDclContext) GetParser() antlr.Parser { return s.parser }

func (s *ForDclContext) CopyAll(ctx *ForDclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ForDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForForeachContext struct {
	ForDclContext
}

func NewForForeachContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForForeachContext {
	var p = new(ForForeachContext)

	InitEmptyForDclContext(&p.ForDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForDclContext))

	return p
}

func (s *ForForeachContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForForeachContext) FOR() antlr.TerminalNode {
	return s.GetToken(VlangParserFOR, 0)
}

func (s *ForForeachContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *ForForeachContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *ForForeachContext) COMMA() antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, 0)
}

func (s *ForForeachContext) IN() antlr.TerminalNode {
	return s.GetToken(VlangParserIN, 0)
}

func (s *ForForeachContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *ForForeachContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterForForeach(s)
	}
}

func (s *ForForeachContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitForForeach(s)
	}
}

func (s *ForForeachContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitForForeach(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForCondicionalContext struct {
	ForDclContext
}

func NewForCondicionalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForCondicionalContext {
	var p = new(ForCondicionalContext)

	InitEmptyForDclContext(&p.ForDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForDclContext))

	return p
}

func (s *ForCondicionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForCondicionalContext) FOR() antlr.TerminalNode {
	return s.GetToken(VlangParserFOR, 0)
}

func (s *ForCondicionalContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ForCondicionalContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *ForCondicionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterForCondicional(s)
	}
}

func (s *ForCondicionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitForCondicional(s)
	}
}

func (s *ForCondicionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitForCondicional(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForClasicoContext struct {
	ForDclContext
}

func NewForClasicoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForClasicoContext {
	var p = new(ForClasicoContext)

	InitEmptyForDclContext(&p.ForDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForDclContext))

	return p
}

func (s *ForClasicoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForClasicoContext) FOR() antlr.TerminalNode {
	return s.GetToken(VlangParserFOR, 0)
}

func (s *ForClasicoContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ForClasicoContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *ForClasicoContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(VlangParserSEMICOLON)
}

func (s *ForClasicoContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserSEMICOLON, i)
}

func (s *ForClasicoContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *ForClasicoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterForClasico(s)
	}
}

func (s *ForClasicoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitForClasico(s)
	}
}

func (s *ForClasicoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitForClasico(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) ForDcl() (localctx IForDclContext) {
	localctx = NewForDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, VlangParserRULE_forDcl)
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForCondicionalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(261)
			p.Match(VlangParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
			p.expresion(0)
		}
		{
			p.SetState(263)
			p.Bloque()
		}

	case 2:
		localctx = NewForClasicoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(265)
			p.Match(VlangParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(266)
			p.expresion(0)
		}
		{
			p.SetState(267)
			p.Match(VlangParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.expresion(0)
		}
		{
			p.SetState(269)
			p.Match(VlangParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.expresion(0)
		}
		{
			p.SetState(271)
			p.Bloque()
		}

	case 3:
		localctx = NewForForeachContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(273)
			p.Match(VlangParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(274)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)
			p.Match(VlangParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(278)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
			p.Bloque()
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

// IWhileDclContext is an interface to support dynamic dispatch.
type IWhileDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expresion() IExpresionContext
	Bloque() IBloqueContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsWhileDclContext differentiates from other interfaces.
	IsWhileDclContext()
}

type WhileDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileDclContext() *WhileDclContext {
	var p = new(WhileDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_whileDcl
	return p
}

func InitEmptyWhileDclContext(p *WhileDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_whileDcl
}

func (*WhileDclContext) IsWhileDclContext() {}

func NewWhileDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileDclContext {
	var p = new(WhileDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_whileDcl

	return p
}

func (s *WhileDclContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileDclContext) WHILE() antlr.TerminalNode {
	return s.GetToken(VlangParserWHILE, 0)
}

func (s *WhileDclContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *WhileDclContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *WhileDclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *WhileDclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *WhileDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileDclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterWhileDcl(s)
	}
}

func (s *WhileDclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitWhileDcl(s)
	}
}

func (s *WhileDclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitWhileDcl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) WhileDcl() (localctx IWhileDclContext) {
	localctx = NewWhileDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, VlangParserRULE_whileDcl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(VlangParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(283)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(286)
		p.expresion(0)
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserRPAREN {
		{
			p.SetState(287)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(290)
		p.Bloque()
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

// IBloqueContext is an interface to support dynamic dispatch.
type IBloqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllDeclaraciones() []IDeclaracionesContext
	Declaraciones(i int) IDeclaracionesContext

	// IsBloqueContext differentiates from other interfaces.
	IsBloqueContext()
}

type BloqueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueContext() *BloqueContext {
	var p = new(BloqueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_bloque
	return p
}

func InitEmptyBloqueContext(p *BloqueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_bloque
}

func (*BloqueContext) IsBloqueContext() {}

func NewBloqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueContext {
	var p = new(BloqueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_bloque

	return p
}

func (s *BloqueContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACE, 0)
}

func (s *BloqueContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACE, 0)
}

func (s *BloqueContext) AllDeclaraciones() []IDeclaracionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracionesContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracionesContext); ok {
			tst[i] = t.(IDeclaracionesContext)
			i++
		}
	}

	return tst
}

func (s *BloqueContext) Declaraciones(i int) IDeclaracionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionesContext); ok {
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

	return t.(IDeclaracionesContext)
}

func (s *BloqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterBloque(s)
	}
}

func (s *BloqueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitBloque(s)
	}
}

func (s *BloqueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitBloque(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Bloque() (localctx IBloqueContext) {
	localctx = NewBloqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, VlangParserRULE_bloque)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(292)
		p.Match(VlangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&189151484924819392) != 0 {
		{
			p.SetState(293)
			p.Declaraciones()
		}

		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(299)
		p.Match(VlangParserRBRACE)
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

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_expresion
	return p
}

func InitEmptyExpresionContext(p *ExpresionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_expresion
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) CopyAll(ctx *ExpresionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AsignacionAtributoContext struct {
	ExpresionContext
}

func NewAsignacionAtributoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionAtributoContext {
	var p = new(AsignacionAtributoContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *AsignacionAtributoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionAtributoContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *AsignacionAtributoContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *AsignacionAtributoContext) DOT() antlr.TerminalNode {
	return s.GetToken(VlangParserDOT, 0)
}

func (s *AsignacionAtributoContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *AsignacionAtributoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionAtributoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAsignacionAtributo(s)
	}
}

func (s *AsignacionAtributoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAsignacionAtributo(s)
	}
}

func (s *AsignacionAtributoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAsignacionAtributo(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultdivmodContext struct {
	ExpresionContext
	op antlr.Token
}

func NewMultdivmodContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultdivmodContext {
	var p = new(MultdivmodContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *MultdivmodContext) GetOp() antlr.Token { return s.op }

func (s *MultdivmodContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultdivmodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultdivmodContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *MultdivmodContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *MultdivmodContext) MUL() antlr.TerminalNode {
	return s.GetToken(VlangParserMUL, 0)
}

func (s *MultdivmodContext) DIV() antlr.TerminalNode {
	return s.GetToken(VlangParserDIV, 0)
}

func (s *MultdivmodContext) MOD() antlr.TerminalNode {
	return s.GetToken(VlangParserMOD, 0)
}

func (s *MultdivmodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterMultdivmod(s)
	}
}

func (s *MultdivmodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitMultdivmod(s)
	}
}

func (s *MultdivmodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitMultdivmod(s)

	default:
		return t.VisitChildren(s)
	}
}

type LenExprContext struct {
	ExpresionContext
}

func NewLenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LenExprContext {
	var p = new(LenExprContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *LenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LenExprContext) LEN() antlr.TerminalNode {
	return s.GetToken(VlangParserLEN, 0)
}

func (s *LenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *LenExprContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *LenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *LenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterLenExpr(s)
	}
}

func (s *LenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitLenExpr(s)
	}
}

func (s *LenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitLenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceDobleAccesoContext struct {
	ExpresionContext
}

func NewSliceDobleAccesoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceDobleAccesoContext {
	var p = new(SliceDobleAccesoContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *SliceDobleAccesoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceDobleAccesoContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *SliceDobleAccesoContext) AllLSQUARE() []antlr.TerminalNode {
	return s.GetTokens(VlangParserLSQUARE)
}

func (s *SliceDobleAccesoContext) LSQUARE(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserLSQUARE, i)
}

func (s *SliceDobleAccesoContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *SliceDobleAccesoContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *SliceDobleAccesoContext) AllRSQUARE() []antlr.TerminalNode {
	return s.GetTokens(VlangParserRSQUARE)
}

func (s *SliceDobleAccesoContext) RSQUARE(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserRSQUARE, i)
}

func (s *SliceDobleAccesoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSliceDobleAcceso(s)
	}
}

func (s *SliceDobleAccesoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSliceDobleAcceso(s)
	}
}

func (s *SliceDobleAccesoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSliceDobleAcceso(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionRestaContext struct {
	ExpresionContext
}

func NewAsignacionRestaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionRestaContext {
	var p = new(AsignacionRestaContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *AsignacionRestaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionRestaContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *AsignacionRestaContext) DEC() antlr.TerminalNode {
	return s.GetToken(VlangParserDEC, 0)
}

func (s *AsignacionRestaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAsignacionResta(s)
	}
}

func (s *AsignacionRestaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAsignacionResta(s)
	}
}

func (s *AsignacionRestaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAsignacionResta(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionforContext struct {
	ExpresionContext
}

func NewAsignacionforContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionforContext {
	var p = new(AsignacionforContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *AsignacionforContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionforContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *AsignacionforContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *AsignacionforContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionforContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAsignacionfor(s)
	}
}

func (s *AsignacionforContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAsignacionfor(s)
	}
}

func (s *AsignacionforContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAsignacionfor(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralSliceExprContext struct {
	ExpresionContext
}

func NewLiteralSliceExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralSliceExprContext {
	var p = new(LiteralSliceExprContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *LiteralSliceExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralSliceExprContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACE, 0)
}

func (s *LiteralSliceExprContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACE, 0)
}

func (s *LiteralSliceExprContext) Parametros() IParametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosContext)
}

func (s *LiteralSliceExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterLiteralSliceExpr(s)
	}
}

func (s *LiteralSliceExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitLiteralSliceExpr(s)
	}
}

func (s *LiteralSliceExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitLiteralSliceExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnarioContext struct {
	ExpresionContext
	op antlr.Token
}

func NewUnarioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnarioContext {
	var p = new(UnarioContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *UnarioContext) GetOp() antlr.Token { return s.op }

func (s *UnarioContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnarioContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *UnarioContext) NOT() antlr.TerminalNode {
	return s.GetToken(VlangParserNOT, 0)
}

func (s *UnarioContext) MINUS() antlr.TerminalNode {
	return s.GetToken(VlangParserMINUS, 0)
}

func (s *UnarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterUnario(s)
	}
}

func (s *UnarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitUnario(s)
	}
}

func (s *UnarioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitUnario(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceAsignacionContext struct {
	ExpresionContext
}

func NewSliceAsignacionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceAsignacionContext {
	var p = new(SliceAsignacionContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *SliceAsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceAsignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *SliceAsignacionContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(VlangParserLSQUARE, 0)
}

func (s *SliceAsignacionContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *SliceAsignacionContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *SliceAsignacionContext) RSQUARE() antlr.TerminalNode {
	return s.GetToken(VlangParserRSQUARE, 0)
}

func (s *SliceAsignacionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *SliceAsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSliceAsignacion(s)
	}
}

func (s *SliceAsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSliceAsignacion(s)
	}
}

func (s *SliceAsignacionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSliceAsignacion(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParentesisexpreContext struct {
	ExpresionContext
}

func NewParentesisexpreContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentesisexpreContext {
	var p = new(ParentesisexpreContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ParentesisexpreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentesisexpreContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *ParentesisexpreContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ParentesisexpreContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *ParentesisexpreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterParentesisexpre(s)
	}
}

func (s *ParentesisexpreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitParentesisexpre(s)
	}
}

func (s *ParentesisexpreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitParentesisexpre(s)

	default:
		return t.VisitChildren(s)
	}
}

type JoinExprContext struct {
	ExpresionContext
}

func NewJoinExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JoinExprContext {
	var p = new(JoinExprContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *JoinExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinExprContext) JOIN() antlr.TerminalNode {
	return s.GetToken(VlangParserJOIN, 0)
}

func (s *JoinExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *JoinExprContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *JoinExprContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *JoinExprContext) COMMA() antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, 0)
}

func (s *JoinExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *JoinExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterJoinExpr(s)
	}
}

func (s *JoinExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitJoinExpr(s)
	}
}

func (s *JoinExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitJoinExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LlamadaMetodoStructContext struct {
	ExpresionContext
}

func NewLlamadaMetodoStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LlamadaMetodoStructContext {
	var p = new(LlamadaMetodoStructContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *LlamadaMetodoStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaMetodoStructContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *LlamadaMetodoStructContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *LlamadaMetodoStructContext) DOT() antlr.TerminalNode {
	return s.GetToken(VlangParserDOT, 0)
}

func (s *LlamadaMetodoStructContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *LlamadaMetodoStructContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *LlamadaMetodoStructContext) Parametros() IParametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosContext)
}

func (s *LlamadaMetodoStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterLlamadaMetodoStruct(s)
	}
}

func (s *LlamadaMetodoStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitLlamadaMetodoStruct(s)
	}
}

func (s *LlamadaMetodoStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitLlamadaMetodoStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

type SumresContext struct {
	ExpresionContext
	op antlr.Token
}

func NewSumresContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SumresContext {
	var p = new(SumresContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *SumresContext) GetOp() antlr.Token { return s.op }

func (s *SumresContext) SetOp(v antlr.Token) { s.op = v }

func (s *SumresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumresContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *SumresContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *SumresContext) PLUS() antlr.TerminalNode {
	return s.GetToken(VlangParserPLUS, 0)
}

func (s *SumresContext) MINUS() antlr.TerminalNode {
	return s.GetToken(VlangParserMINUS, 0)
}

func (s *SumresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSumres(s)
	}
}

func (s *SumresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSumres(s)
	}
}

func (s *SumresContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSumres(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceCreacionContext struct {
	ExpresionContext
}

func NewSliceCreacionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceCreacionContext {
	var p = new(SliceCreacionContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *SliceCreacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceCreacionContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(VlangParserLSQUARE, 0)
}

func (s *SliceCreacionContext) RSQUARE() antlr.TerminalNode {
	return s.GetToken(VlangParserRSQUARE, 0)
}

func (s *SliceCreacionContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *SliceCreacionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACE, 0)
}

func (s *SliceCreacionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACE, 0)
}

func (s *SliceCreacionContext) Parametros() IParametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosContext)
}

func (s *SliceCreacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSliceCreacion(s)
	}
}

func (s *SliceCreacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSliceCreacion(s)
	}
}

func (s *SliceCreacionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSliceCreacion(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdContext struct {
	ExpresionContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitId(s)
	}
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

type AppendExprContext struct {
	ExpresionContext
}

func NewAppendExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppendExprContext {
	var p = new(AppendExprContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *AppendExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendExprContext) APPEND() antlr.TerminalNode {
	return s.GetToken(VlangParserAPPEND, 0)
}

func (s *AppendExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *AppendExprContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *AppendExprContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *AppendExprContext) COMMA() antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, 0)
}

func (s *AppendExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *AppendExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAppendExpr(s)
	}
}

func (s *AppendExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAppendExpr(s)
	}
}

func (s *AppendExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAppendExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionCompuestaRestaContext struct {
	ExpresionContext
}

func NewAsignacionCompuestaRestaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionCompuestaRestaContext {
	var p = new(AsignacionCompuestaRestaContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *AsignacionCompuestaRestaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionCompuestaRestaContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *AsignacionCompuestaRestaContext) MINUSEQ() antlr.TerminalNode {
	return s.GetToken(VlangParserMINUSEQ, 0)
}

func (s *AsignacionCompuestaRestaContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionCompuestaRestaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAsignacionCompuestaResta(s)
	}
}

func (s *AsignacionCompuestaRestaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAsignacionCompuestaResta(s)
	}
}

func (s *AsignacionCompuestaRestaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAsignacionCompuestaResta(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorexprContext struct {
	ExpresionContext
}

func NewValorexprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorexprContext {
	var p = new(ValorexprContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ValorexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorexprContext) Valor() IValorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValorContext)
}

func (s *ValorexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorexpr(s)
	}
}

func (s *ValorexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorexpr(s)
	}
}

func (s *ValorexprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorexpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IgualdadContext struct {
	ExpresionContext
	op antlr.Token
}

func NewIgualdadContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IgualdadContext {
	var p = new(IgualdadContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *IgualdadContext) GetOp() antlr.Token { return s.op }

func (s *IgualdadContext) SetOp(v antlr.Token) { s.op = v }

func (s *IgualdadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IgualdadContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *IgualdadContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *IgualdadContext) EQ() antlr.TerminalNode {
	return s.GetToken(VlangParserEQ, 0)
}

func (s *IgualdadContext) NEQ() antlr.TerminalNode {
	return s.GetToken(VlangParserNEQ, 0)
}

func (s *IgualdadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIgualdad(s)
	}
}

func (s *IgualdadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIgualdad(s)
	}
}

func (s *IgualdadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIgualdad(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	ExpresionContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *OrExprContext) OR() antlr.TerminalNode {
	return s.GetToken(VlangParserOR, 0)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceDobleAsignacionContext struct {
	ExpresionContext
}

func NewSliceDobleAsignacionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceDobleAsignacionContext {
	var p = new(SliceDobleAsignacionContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *SliceDobleAsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceDobleAsignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *SliceDobleAsignacionContext) AllLSQUARE() []antlr.TerminalNode {
	return s.GetTokens(VlangParserLSQUARE)
}

func (s *SliceDobleAsignacionContext) LSQUARE(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserLSQUARE, i)
}

func (s *SliceDobleAsignacionContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *SliceDobleAsignacionContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *SliceDobleAsignacionContext) AllRSQUARE() []antlr.TerminalNode {
	return s.GetTokens(VlangParserRSQUARE)
}

func (s *SliceDobleAsignacionContext) RSQUARE(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserRSQUARE, i)
}

func (s *SliceDobleAsignacionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *SliceDobleAsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSliceDobleAsignacion(s)
	}
}

func (s *SliceDobleAsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSliceDobleAsignacion(s)
	}
}

func (s *SliceDobleAsignacionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSliceDobleAsignacion(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceAccesoContext struct {
	ExpresionContext
}

func NewSliceAccesoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceAccesoContext {
	var p = new(SliceAccesoContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *SliceAccesoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceAccesoContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *SliceAccesoContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(VlangParserLSQUARE, 0)
}

func (s *SliceAccesoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SliceAccesoContext) RSQUARE() antlr.TerminalNode {
	return s.GetToken(VlangParserRSQUARE, 0)
}

func (s *SliceAccesoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSliceAcceso(s)
	}
}

func (s *SliceAccesoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSliceAcceso(s)
	}
}

func (s *SliceAccesoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSliceAcceso(s)

	default:
		return t.VisitChildren(s)
	}
}

type RelacionalesContext struct {
	ExpresionContext
	op antlr.Token
}

func NewRelacionalesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelacionalesContext {
	var p = new(RelacionalesContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *RelacionalesContext) GetOp() antlr.Token { return s.op }

func (s *RelacionalesContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelacionalesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelacionalesContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *RelacionalesContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *RelacionalesContext) LT() antlr.TerminalNode {
	return s.GetToken(VlangParserLT, 0)
}

func (s *RelacionalesContext) LE() antlr.TerminalNode {
	return s.GetToken(VlangParserLE, 0)
}

func (s *RelacionalesContext) GE() antlr.TerminalNode {
	return s.GetToken(VlangParserGE, 0)
}

func (s *RelacionalesContext) GT() antlr.TerminalNode {
	return s.GetToken(VlangParserGT, 0)
}

func (s *RelacionalesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterRelacionales(s)
	}
}

func (s *RelacionalesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitRelacionales(s)
	}
}

func (s *RelacionalesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitRelacionales(s)

	default:
		return t.VisitChildren(s)
	}
}

type LlamadaFuncionContext struct {
	ExpresionContext
}

func NewLlamadaFuncionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LlamadaFuncionContext {
	var p = new(LlamadaFuncionContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *LlamadaFuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaFuncionContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *LlamadaFuncionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *LlamadaFuncionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *LlamadaFuncionContext) Parametros() IParametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosContext)
}

func (s *LlamadaFuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterLlamadaFuncion(s)
	}
}

func (s *LlamadaFuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitLlamadaFuncion(s)
	}
}

func (s *LlamadaFuncionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitLlamadaFuncion(s)

	default:
		return t.VisitChildren(s)
	}
}

type AccesoAtributoContext struct {
	ExpresionContext
}

func NewAccesoAtributoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccesoAtributoContext {
	var p = new(AccesoAtributoContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *AccesoAtributoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesoAtributoContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *AccesoAtributoContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *AccesoAtributoContext) DOT() antlr.TerminalNode {
	return s.GetToken(VlangParserDOT, 0)
}

func (s *AccesoAtributoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAccesoAtributo(s)
	}
}

func (s *AccesoAtributoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAccesoAtributo(s)
	}
}

func (s *AccesoAtributoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAccesoAtributo(s)

	default:
		return t.VisitChildren(s)
	}
}

type InstanciaStructContext struct {
	ExpresionContext
}

func NewInstanciaStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InstanciaStructContext {
	var p = new(InstanciaStructContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *InstanciaStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanciaStructContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *InstanciaStructContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACE, 0)
}

func (s *InstanciaStructContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACE, 0)
}

func (s *InstanciaStructContext) InicializadorStruct() IInicializadorStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInicializadorStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInicializadorStructContext)
}

func (s *InstanciaStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterInstanciaStruct(s)
	}
}

func (s *InstanciaStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitInstanciaStruct(s)
	}
}

func (s *InstanciaStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitInstanciaStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionCompuestaSumaContext struct {
	ExpresionContext
}

func NewAsignacionCompuestaSumaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionCompuestaSumaContext {
	var p = new(AsignacionCompuestaSumaContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *AsignacionCompuestaSumaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionCompuestaSumaContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *AsignacionCompuestaSumaContext) PLUSEQ() antlr.TerminalNode {
	return s.GetToken(VlangParserPLUSEQ, 0)
}

func (s *AsignacionCompuestaSumaContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionCompuestaSumaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAsignacionCompuestaSuma(s)
	}
}

func (s *AsignacionCompuestaSumaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAsignacionCompuestaSuma(s)
	}
}

func (s *AsignacionCompuestaSumaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAsignacionCompuestaSuma(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexOfExprContext struct {
	ExpresionContext
}

func NewIndexOfExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexOfExprContext {
	var p = new(IndexOfExprContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *IndexOfExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexOfExprContext) INDEXOF() antlr.TerminalNode {
	return s.GetToken(VlangParserINDEXOF, 0)
}

func (s *IndexOfExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *IndexOfExprContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *IndexOfExprContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *IndexOfExprContext) COMMA() antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, 0)
}

func (s *IndexOfExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *IndexOfExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIndexOfExpr(s)
	}
}

func (s *IndexOfExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIndexOfExpr(s)
	}
}

func (s *IndexOfExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIndexOfExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionSumaContext struct {
	ExpresionContext
}

func NewAsignacionSumaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionSumaContext {
	var p = new(AsignacionSumaContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *AsignacionSumaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionSumaContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *AsignacionSumaContext) INC() antlr.TerminalNode {
	return s.GetToken(VlangParserINC, 0)
}

func (s *AsignacionSumaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAsignacionSuma(s)
	}
}

func (s *AsignacionSumaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAsignacionSuma(s)
	}
}

func (s *AsignacionSumaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAsignacionSuma(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExprContext struct {
	ExpresionContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *AndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(VlangParserAND, 0)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *VlangParser) expresion(_p int) (localctx IExpresionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 52
	p.EnterRecursionRule(localctx, 52, VlangParserRULE_expresion, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAsignacionCompuestaSumaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(302)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)
			p.Match(VlangParserPLUSEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)
			p.expresion(30)
		}

	case 2:
		localctx = NewAsignacionCompuestaRestaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(305)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.Match(VlangParserMINUSEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)
			p.expresion(29)
		}

	case 3:
		localctx = NewLlamadaFuncionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(308)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&189151484923871232) != 0 {
			{
				p.SetState(310)
				p.Parametros()
			}

		}
		{
			p.SetState(313)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewAsignacionforContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(314)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(315)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)
			p.expresion(27)
		}

	case 5:
		localctx = NewAsignacionSumaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(317)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(318)
			p.Match(VlangParserINC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewAsignacionRestaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(319)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)
			p.Match(VlangParserDEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewAsignacionAtributoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(321)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)
			p.Match(VlangParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(325)
			p.expresion(24)
		}

	case 8:
		localctx = NewAccesoAtributoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(326)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)
			p.Match(VlangParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewLlamadaMetodoStructContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(329)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)
			p.Match(VlangParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&189151484923871232) != 0 {
			{
				p.SetState(333)
				p.Parametros()
			}

		}
		{
			p.SetState(336)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewInstanciaStructContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(337)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.Match(VlangParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VlangParserID {
			{
				p.SetState(339)
				p.InicializadorStruct()
			}

		}
		{
			p.SetState(342)
			p.Match(VlangParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewUnarioContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(343)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnarioContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == VlangParserMINUS || _la == VlangParserNOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnarioContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(344)
			p.expresion(14)
		}

	case 12:
		localctx = NewParentesisexpreContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(345)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)
			p.expresion(0)
		}
		{
			p.SetState(347)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		localctx = NewValorexprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(349)
			p.Valor()
		}

	case 14:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(350)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		localctx = NewLiteralSliceExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(351)
			p.Match(VlangParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(353)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&189151484923871232) != 0 {
			{
				p.SetState(352)
				p.Parametros()
			}

		}
		{
			p.SetState(355)
			p.Match(VlangParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		localctx = NewSliceCreacionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(356)
			p.Match(VlangParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.Match(VlangParserRSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(358)
			p.Tipo()
		}
		{
			p.SetState(359)
			p.Match(VlangParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(361)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&189151484923871232) != 0 {
			{
				p.SetState(360)
				p.Parametros()
			}

		}
		{
			p.SetState(363)
			p.Match(VlangParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 17:
		localctx = NewIndexOfExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(365)
			p.Match(VlangParserINDEXOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(366)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)
			p.expresion(0)
		}
		{
			p.SetState(368)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)
			p.expresion(0)
		}
		{
			p.SetState(370)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 18:
		localctx = NewJoinExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(372)
			p.Match(VlangParserJOIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(374)
			p.expresion(0)
		}
		{
			p.SetState(375)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)
			p.expresion(0)
		}
		{
			p.SetState(377)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 19:
		localctx = NewLenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(379)
			p.Match(VlangParserLEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)
			p.expresion(0)
		}
		{
			p.SetState(382)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 20:
		localctx = NewAppendExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(384)
			p.Match(VlangParserAPPEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(385)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
			p.expresion(0)
		}
		{
			p.SetState(387)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(388)
			p.expresion(0)
		}
		{
			p.SetState(389)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 21:
		localctx = NewSliceAccesoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(391)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(392)
			p.Match(VlangParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.expresion(0)
		}
		{
			p.SetState(394)
			p.Match(VlangParserRSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 22:
		localctx = NewSliceAsignacionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(396)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)
			p.Match(VlangParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)
			p.expresion(0)
		}
		{
			p.SetState(399)
			p.Match(VlangParserRSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(400)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(401)
			p.expresion(3)
		}

	case 23:
		localctx = NewSliceDobleAccesoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(403)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(404)
			p.Match(VlangParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.expresion(0)
		}
		{
			p.SetState(406)
			p.Match(VlangParserRSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)
			p.Match(VlangParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(408)
			p.expresion(0)
		}
		{
			p.SetState(409)
			p.Match(VlangParserRSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 24:
		localctx = NewSliceDobleAsignacionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(411)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(412)
			p.Match(VlangParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(413)
			p.expresion(0)
		}
		{
			p.SetState(414)
			p.Match(VlangParserRSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(415)
			p.Match(VlangParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)
			p.expresion(0)
		}
		{
			p.SetState(417)
			p.Match(VlangParserRSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.expresion(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(441)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultdivmodContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(423)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(424)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultdivmodContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&240518168576) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultdivmodContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(425)
					p.expresion(21)
				}

			case 2:
				localctx = NewSumresContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(426)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(427)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*SumresContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VlangParserPLUS || _la == VlangParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*SumresContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(428)
					p.expresion(20)
				}

			case 3:
				localctx = NewRelacionalesContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(429)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(430)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelacionalesContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&131941395333120) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelacionalesContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(431)
					p.expresion(19)
				}

			case 4:
				localctx = NewIgualdadContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(432)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(433)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*IgualdadContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VlangParserEQ || _la == VlangParserNEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*IgualdadContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(434)
					p.expresion(18)
				}

			case 5:
				localctx = NewAndExprContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(435)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(436)
					p.Match(VlangParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(437)
					p.expresion(17)
				}

			case 6:
				localctx = NewOrExprContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(438)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(439)
					p.Match(VlangParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(440)
					p.expresion(16)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(445)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
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

// IParametrosContext is an interface to support dynamic dispatch.
type IParametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParametrosContext differentiates from other interfaces.
	IsParametrosContext()
}

type ParametrosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametrosContext() *ParametrosContext {
	var p = new(ParametrosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametros
	return p
}

func InitEmptyParametrosContext(p *ParametrosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametros
}

func (*ParametrosContext) IsParametrosContext() {}

func NewParametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosContext {
	var p = new(ParametrosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_parametros

	return p
}

func (s *ParametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ParametrosContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *ParametrosContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *ParametrosContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *ParametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterParametros(s)
	}
}

func (s *ParametrosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitParametros(s)
	}
}

func (s *ParametrosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitParametros(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Parametros() (localctx IParametrosContext) {
	localctx = NewParametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, VlangParserRULE_parametros)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(446)
		p.expresion(0)
	}
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(447)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(448)
			p.expresion(0)
		}

		p.SetState(453)
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

// IValorContext is an interface to support dynamic dispatch.
type IValorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValorContext differentiates from other interfaces.
	IsValorContext()
}

type ValorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValorContext() *ValorContext {
	var p = new(ValorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_valor
	return p
}

func InitEmptyValorContext(p *ValorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_valor
}

func (*ValorContext) IsValorContext() {}

func NewValorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValorContext {
	var p = new(ValorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_valor

	return p
}

func (s *ValorContext) GetParser() antlr.Parser { return s.parser }

func (s *ValorContext) CopyAll(ctx *ValorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValorDecimalContext struct {
	ValorContext
}

func NewValorDecimalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorDecimalContext {
	var p = new(ValorDecimalContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorDecimalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorDecimalContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(VlangParserDECIMAL, 0)
}

func (s *ValorDecimalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorDecimal(s)
	}
}

func (s *ValorDecimalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorDecimal(s)
	}
}

func (s *ValorDecimalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorDecimal(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorEnteroContext struct {
	ValorContext
}

func NewValorEnteroContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorEnteroContext {
	var p = new(ValorEnteroContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorEnteroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorEnteroContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(VlangParserENTERO, 0)
}

func (s *ValorEnteroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorEntero(s)
	}
}

func (s *ValorEnteroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorEntero(s)
	}
}

func (s *ValorEnteroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorEntero(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorBooleanoContext struct {
	ValorContext
}

func NewValorBooleanoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorBooleanoContext {
	var p = new(ValorBooleanoContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorBooleanoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorBooleanoContext) BOOLEANO() antlr.TerminalNode {
	return s.GetToken(VlangParserBOOLEANO, 0)
}

func (s *ValorBooleanoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorBooleano(s)
	}
}

func (s *ValorBooleanoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorBooleano(s)
	}
}

func (s *ValorBooleanoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorBooleano(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorCaracterContext struct {
	ValorContext
}

func NewValorCaracterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorCaracterContext {
	var p = new(ValorCaracterContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorCaracterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorCaracterContext) CARACTER() antlr.TerminalNode {
	return s.GetToken(VlangParserCARACTER, 0)
}

func (s *ValorCaracterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorCaracter(s)
	}
}

func (s *ValorCaracterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorCaracter(s)
	}
}

func (s *ValorCaracterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorCaracter(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorCadenaContext struct {
	ValorContext
}

func NewValorCadenaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorCadenaContext {
	var p = new(ValorCadenaContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorCadenaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorCadenaContext) CADENA() antlr.TerminalNode {
	return s.GetToken(VlangParserCADENA, 0)
}

func (s *ValorCadenaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorCadena(s)
	}
}

func (s *ValorCadenaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorCadena(s)
	}
}

func (s *ValorCadenaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorCadena(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Valor() (localctx IValorContext) {
	localctx = NewValorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, VlangParserRULE_valor)
	p.SetState(459)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserENTERO:
		localctx = NewValorEnteroContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(454)
			p.Match(VlangParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserDECIMAL:
		localctx = NewValorDecimalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(455)
			p.Match(VlangParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserCADENA:
		localctx = NewValorCadenaContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(456)
			p.Match(VlangParserCADENA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserBOOLEANO:
		localctx = NewValorBooleanoContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(457)
			p.Match(VlangParserBOOLEANO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserCARACTER:
		localctx = NewValorCaracterContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(458)
			p.Match(VlangParserCARACTER)
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

func (p *VlangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 26:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *VlangParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 15)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
