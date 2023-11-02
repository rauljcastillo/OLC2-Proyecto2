// Code generated from Gramar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Gramar

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

type GramarParser struct {
	*antlr.BaseParser
}

var GramarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gramarParserInit() {
	staticData := &GramarParserStaticData
	staticData.LiteralNames = []string{
		"", "'switch'", "'{'", "'}'", "'case'", "':'", "'default'", "'for'",
		"'in'", "'print'", "'('", "')'", "'while'", "'='", "'+='", "'-='", "';'",
		"'?'", "'if'", "'else'", "'func'", "'->'", "','", "'!'", "'-'", "'/'",
		"'*'", "'%'", "'+'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'&&'",
		"'||'", "", "", "", "", "'var'", "'let'", "'Int'", "'Float'", "'String'",
		"'Bool'", "'Character'", "'_'", "'inout'", "'&'", "'['", "']'", "'...'",
		"'break'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "WS", "COMS", "COMM", "BOOLEAN", "RVAR", "RLET", "RINT",
		"RFLOAT", "RSTRING", "RBOOL", "RChar", "GUION", "INOUT", "PUNTE", "ABR",
		"CIER", "POINTS", "BRK", "ID", "INT", "DOUBLE", "CADENA", "CHAR",
	}
	staticData.RuleNames = []string{
		"s", "block", "productions", "pbrk", "pswitch", "ccase", "pdefault",
		"pfor", "prin", "swhile", "asign", "pdeclara", "tipo", "pif", "pelse",
		"pfuncion", "pparams", "pparamet", "pllamada", "argumento", "tipoArg",
		"expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 59, 292, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 49, 8, 1, 10, 1, 12, 1, 52, 9,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 64,
		8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 4, 4, 72, 8, 4, 11, 4, 12, 4,
		73, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 84, 8, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 3, 6, 90, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 3, 7, 100, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 128, 8, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3,
		11, 140, 8, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 157, 8, 13, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 166, 8, 14, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 174, 8, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 187,
		8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 193, 8, 15, 1, 16, 1, 16, 1,
		16, 5, 16, 198, 8, 16, 10, 16, 12, 16, 201, 9, 16, 3, 16, 203, 8, 16, 1,
		17, 1, 17, 1, 17, 1, 17, 3, 17, 209, 8, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 3, 17, 216, 8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 222, 8, 17,
		1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 228, 8, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 3, 18, 235, 8, 18, 3, 18, 237, 8, 18, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 3, 19, 244, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		3, 20, 252, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 267, 8, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 287, 8, 21, 10, 21, 12, 21,
		290, 9, 21, 1, 21, 0, 1, 42, 22, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 0, 9, 1, 0, 13, 15, 1, 0, 41,
		42, 1, 0, 43, 47, 2, 0, 48, 48, 55, 55, 1, 0, 23, 24, 1, 0, 25, 27, 2,
		0, 24, 24, 28, 28, 1, 0, 29, 32, 1, 0, 33, 34, 315, 0, 44, 1, 0, 0, 0,
		2, 50, 1, 0, 0, 0, 4, 63, 1, 0, 0, 0, 6, 65, 1, 0, 0, 0, 8, 67, 1, 0, 0,
		0, 10, 78, 1, 0, 0, 0, 12, 85, 1, 0, 0, 0, 14, 91, 1, 0, 0, 0, 16, 105,
		1, 0, 0, 0, 18, 110, 1, 0, 0, 0, 20, 116, 1, 0, 0, 0, 22, 139, 1, 0, 0,
		0, 24, 141, 1, 0, 0, 0, 26, 156, 1, 0, 0, 0, 28, 165, 1, 0, 0, 0, 30, 192,
		1, 0, 0, 0, 32, 202, 1, 0, 0, 0, 34, 221, 1, 0, 0, 0, 36, 236, 1, 0, 0,
		0, 38, 243, 1, 0, 0, 0, 40, 251, 1, 0, 0, 0, 42, 266, 1, 0, 0, 0, 44, 45,
		3, 2, 1, 0, 45, 46, 5, 0, 0, 1, 46, 1, 1, 0, 0, 0, 47, 49, 3, 4, 2, 0,
		48, 47, 1, 0, 0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1,
		0, 0, 0, 51, 3, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53, 64, 3, 16, 8, 0, 54,
		64, 3, 18, 9, 0, 55, 64, 3, 26, 13, 0, 56, 64, 3, 22, 11, 0, 57, 64, 3,
		20, 10, 0, 58, 64, 3, 14, 7, 0, 59, 64, 3, 8, 4, 0, 60, 64, 3, 30, 15,
		0, 61, 64, 3, 36, 18, 0, 62, 64, 3, 6, 3, 0, 63, 53, 1, 0, 0, 0, 63, 54,
		1, 0, 0, 0, 63, 55, 1, 0, 0, 0, 63, 56, 1, 0, 0, 0, 63, 57, 1, 0, 0, 0,
		63, 58, 1, 0, 0, 0, 63, 59, 1, 0, 0, 0, 63, 60, 1, 0, 0, 0, 63, 61, 1,
		0, 0, 0, 63, 62, 1, 0, 0, 0, 64, 5, 1, 0, 0, 0, 65, 66, 5, 54, 0, 0, 66,
		7, 1, 0, 0, 0, 67, 68, 5, 1, 0, 0, 68, 69, 3, 42, 21, 0, 69, 71, 5, 2,
		0, 0, 70, 72, 3, 10, 5, 0, 71, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73,
		71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 76, 3, 12,
		6, 0, 76, 77, 5, 3, 0, 0, 77, 9, 1, 0, 0, 0, 78, 79, 5, 4, 0, 0, 79, 80,
		3, 42, 21, 0, 80, 81, 5, 5, 0, 0, 81, 83, 3, 2, 1, 0, 82, 84, 5, 54, 0,
		0, 83, 82, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 11, 1, 0, 0, 0, 85, 86,
		5, 6, 0, 0, 86, 87, 5, 5, 0, 0, 87, 89, 3, 2, 1, 0, 88, 90, 5, 54, 0, 0,
		89, 88, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 13, 1, 0, 0, 0, 91, 92, 5,
		7, 0, 0, 92, 93, 5, 55, 0, 0, 93, 99, 5, 8, 0, 0, 94, 95, 3, 42, 21, 0,
		95, 96, 5, 53, 0, 0, 96, 97, 3, 42, 21, 0, 97, 100, 1, 0, 0, 0, 98, 100,
		3, 42, 21, 0, 99, 94, 1, 0, 0, 0, 99, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0,
		0, 101, 102, 5, 2, 0, 0, 102, 103, 3, 2, 1, 0, 103, 104, 5, 3, 0, 0, 104,
		15, 1, 0, 0, 0, 105, 106, 5, 9, 0, 0, 106, 107, 5, 10, 0, 0, 107, 108,
		3, 42, 21, 0, 108, 109, 5, 11, 0, 0, 109, 17, 1, 0, 0, 0, 110, 111, 5,
		12, 0, 0, 111, 112, 3, 42, 21, 0, 112, 113, 5, 2, 0, 0, 113, 114, 3, 2,
		1, 0, 114, 115, 5, 3, 0, 0, 115, 19, 1, 0, 0, 0, 116, 117, 5, 55, 0, 0,
		117, 118, 7, 0, 0, 0, 118, 119, 3, 42, 21, 0, 119, 21, 1, 0, 0, 0, 120,
		121, 7, 1, 0, 0, 121, 122, 5, 55, 0, 0, 122, 123, 5, 5, 0, 0, 123, 124,
		3, 24, 12, 0, 124, 125, 5, 13, 0, 0, 125, 127, 3, 42, 21, 0, 126, 128,
		5, 16, 0, 0, 127, 126, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 140, 1, 0,
		0, 0, 129, 130, 7, 1, 0, 0, 130, 131, 5, 55, 0, 0, 131, 132, 5, 5, 0, 0,
		132, 133, 3, 24, 12, 0, 133, 134, 5, 17, 0, 0, 134, 140, 1, 0, 0, 0, 135,
		136, 7, 1, 0, 0, 136, 137, 5, 55, 0, 0, 137, 138, 5, 13, 0, 0, 138, 140,
		3, 42, 21, 0, 139, 120, 1, 0, 0, 0, 139, 129, 1, 0, 0, 0, 139, 135, 1,
		0, 0, 0, 140, 23, 1, 0, 0, 0, 141, 142, 7, 2, 0, 0, 142, 25, 1, 0, 0, 0,
		143, 144, 5, 18, 0, 0, 144, 145, 3, 42, 21, 0, 145, 146, 5, 2, 0, 0, 146,
		147, 3, 2, 1, 0, 147, 148, 5, 3, 0, 0, 148, 157, 1, 0, 0, 0, 149, 150,
		5, 18, 0, 0, 150, 151, 3, 42, 21, 0, 151, 152, 5, 2, 0, 0, 152, 153, 3,
		2, 1, 0, 153, 154, 5, 3, 0, 0, 154, 155, 3, 28, 14, 0, 155, 157, 1, 0,
		0, 0, 156, 143, 1, 0, 0, 0, 156, 149, 1, 0, 0, 0, 157, 27, 1, 0, 0, 0,
		158, 159, 5, 19, 0, 0, 159, 166, 3, 26, 13, 0, 160, 161, 5, 19, 0, 0, 161,
		162, 5, 2, 0, 0, 162, 163, 3, 2, 1, 0, 163, 164, 5, 3, 0, 0, 164, 166,
		1, 0, 0, 0, 165, 158, 1, 0, 0, 0, 165, 160, 1, 0, 0, 0, 166, 29, 1, 0,
		0, 0, 167, 168, 5, 20, 0, 0, 168, 169, 5, 55, 0, 0, 169, 170, 5, 10, 0,
		0, 170, 173, 5, 11, 0, 0, 171, 172, 5, 21, 0, 0, 172, 174, 3, 24, 12, 0,
		173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175,
		176, 5, 2, 0, 0, 176, 177, 3, 2, 1, 0, 177, 178, 5, 3, 0, 0, 178, 193,
		1, 0, 0, 0, 179, 180, 5, 20, 0, 0, 180, 181, 5, 55, 0, 0, 181, 182, 5,
		10, 0, 0, 182, 183, 3, 32, 16, 0, 183, 186, 5, 11, 0, 0, 184, 185, 5, 21,
		0, 0, 185, 187, 3, 24, 12, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0,
		0, 187, 188, 1, 0, 0, 0, 188, 189, 5, 2, 0, 0, 189, 190, 3, 2, 1, 0, 190,
		191, 5, 3, 0, 0, 191, 193, 1, 0, 0, 0, 192, 167, 1, 0, 0, 0, 192, 179,
		1, 0, 0, 0, 193, 31, 1, 0, 0, 0, 194, 199, 3, 34, 17, 0, 195, 196, 5, 22,
		0, 0, 196, 198, 3, 34, 17, 0, 197, 195, 1, 0, 0, 0, 198, 201, 1, 0, 0,
		0, 199, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 203, 1, 0, 0, 0, 201,
		199, 1, 0, 0, 0, 202, 194, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 33, 1,
		0, 0, 0, 204, 205, 7, 3, 0, 0, 205, 206, 5, 55, 0, 0, 206, 208, 5, 5, 0,
		0, 207, 209, 5, 49, 0, 0, 208, 207, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209,
		210, 1, 0, 0, 0, 210, 222, 3, 24, 12, 0, 211, 212, 7, 3, 0, 0, 212, 213,
		5, 55, 0, 0, 213, 215, 5, 5, 0, 0, 214, 216, 5, 49, 0, 0, 215, 214, 1,
		0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 218, 5, 51, 0,
		0, 218, 219, 3, 24, 12, 0, 219, 220, 5, 52, 0, 0, 220, 222, 1, 0, 0, 0,
		221, 204, 1, 0, 0, 0, 221, 211, 1, 0, 0, 0, 222, 35, 1, 0, 0, 0, 223, 224,
		5, 55, 0, 0, 224, 225, 5, 10, 0, 0, 225, 227, 5, 11, 0, 0, 226, 228, 5,
		16, 0, 0, 227, 226, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 237, 1, 0, 0,
		0, 229, 230, 5, 55, 0, 0, 230, 231, 5, 10, 0, 0, 231, 232, 3, 38, 19, 0,
		232, 234, 5, 11, 0, 0, 233, 235, 5, 16, 0, 0, 234, 233, 1, 0, 0, 0, 234,
		235, 1, 0, 0, 0, 235, 237, 1, 0, 0, 0, 236, 223, 1, 0, 0, 0, 236, 229,
		1, 0, 0, 0, 237, 37, 1, 0, 0, 0, 238, 239, 3, 40, 20, 0, 239, 240, 5, 22,
		0, 0, 240, 241, 3, 38, 19, 0, 241, 244, 1, 0, 0, 0, 242, 244, 3, 40, 20,
		0, 243, 238, 1, 0, 0, 0, 243, 242, 1, 0, 0, 0, 244, 39, 1, 0, 0, 0, 245,
		246, 5, 50, 0, 0, 246, 252, 3, 42, 21, 0, 247, 248, 5, 55, 0, 0, 248, 249,
		5, 5, 0, 0, 249, 252, 3, 42, 21, 0, 250, 252, 3, 42, 21, 0, 251, 245, 1,
		0, 0, 0, 251, 247, 1, 0, 0, 0, 251, 250, 1, 0, 0, 0, 252, 41, 1, 0, 0,
		0, 253, 254, 6, 21, -1, 0, 254, 255, 7, 4, 0, 0, 255, 267, 3, 42, 21, 14,
		256, 257, 5, 10, 0, 0, 257, 258, 3, 42, 21, 0, 258, 259, 5, 11, 0, 0, 259,
		267, 1, 0, 0, 0, 260, 267, 5, 56, 0, 0, 261, 267, 5, 57, 0, 0, 262, 267,
		5, 40, 0, 0, 263, 267, 5, 58, 0, 0, 264, 267, 5, 59, 0, 0, 265, 267, 5,
		55, 0, 0, 266, 253, 1, 0, 0, 0, 266, 256, 1, 0, 0, 0, 266, 260, 1, 0, 0,
		0, 266, 261, 1, 0, 0, 0, 266, 262, 1, 0, 0, 0, 266, 263, 1, 0, 0, 0, 266,
		264, 1, 0, 0, 0, 266, 265, 1, 0, 0, 0, 267, 288, 1, 0, 0, 0, 268, 269,
		10, 12, 0, 0, 269, 270, 7, 5, 0, 0, 270, 287, 3, 42, 21, 13, 271, 272,
		10, 11, 0, 0, 272, 273, 7, 6, 0, 0, 273, 287, 3, 42, 21, 12, 274, 275,
		10, 10, 0, 0, 275, 276, 7, 7, 0, 0, 276, 287, 3, 42, 21, 11, 277, 278,
		10, 9, 0, 0, 278, 279, 7, 8, 0, 0, 279, 287, 3, 42, 21, 10, 280, 281, 10,
		8, 0, 0, 281, 282, 5, 35, 0, 0, 282, 287, 3, 42, 21, 9, 283, 284, 10, 7,
		0, 0, 284, 285, 5, 36, 0, 0, 285, 287, 3, 42, 21, 8, 286, 268, 1, 0, 0,
		0, 286, 271, 1, 0, 0, 0, 286, 274, 1, 0, 0, 0, 286, 277, 1, 0, 0, 0, 286,
		280, 1, 0, 0, 0, 286, 283, 1, 0, 0, 0, 287, 290, 1, 0, 0, 0, 288, 286,
		1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 43, 1, 0, 0, 0, 290, 288, 1, 0,
		0, 0, 26, 50, 63, 73, 83, 89, 99, 127, 139, 156, 165, 173, 186, 192, 199,
		202, 208, 215, 221, 227, 234, 236, 243, 251, 266, 286, 288,
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

// GramarParserInit initializes any static state used to implement GramarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGramarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GramarParserInit() {
	staticData := &GramarParserStaticData
	staticData.once.Do(gramarParserInit)
}

// NewGramarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGramarParser(input antlr.TokenStream) *GramarParser {
	GramarParserInit()
	this := new(GramarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GramarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Gramar.g4"

	return this
}

// GramarParser tokens.
const (
	GramarParserEOF     = antlr.TokenEOF
	GramarParserT__0    = 1
	GramarParserT__1    = 2
	GramarParserT__2    = 3
	GramarParserT__3    = 4
	GramarParserT__4    = 5
	GramarParserT__5    = 6
	GramarParserT__6    = 7
	GramarParserT__7    = 8
	GramarParserT__8    = 9
	GramarParserT__9    = 10
	GramarParserT__10   = 11
	GramarParserT__11   = 12
	GramarParserT__12   = 13
	GramarParserT__13   = 14
	GramarParserT__14   = 15
	GramarParserT__15   = 16
	GramarParserT__16   = 17
	GramarParserT__17   = 18
	GramarParserT__18   = 19
	GramarParserT__19   = 20
	GramarParserT__20   = 21
	GramarParserT__21   = 22
	GramarParserT__22   = 23
	GramarParserT__23   = 24
	GramarParserT__24   = 25
	GramarParserT__25   = 26
	GramarParserT__26   = 27
	GramarParserT__27   = 28
	GramarParserT__28   = 29
	GramarParserT__29   = 30
	GramarParserT__30   = 31
	GramarParserT__31   = 32
	GramarParserT__32   = 33
	GramarParserT__33   = 34
	GramarParserT__34   = 35
	GramarParserT__35   = 36
	GramarParserWS      = 37
	GramarParserCOMS    = 38
	GramarParserCOMM    = 39
	GramarParserBOOLEAN = 40
	GramarParserRVAR    = 41
	GramarParserRLET    = 42
	GramarParserRINT    = 43
	GramarParserRFLOAT  = 44
	GramarParserRSTRING = 45
	GramarParserRBOOL   = 46
	GramarParserRChar   = 47
	GramarParserGUION   = 48
	GramarParserINOUT   = 49
	GramarParserPUNTE   = 50
	GramarParserABR     = 51
	GramarParserCIER    = 52
	GramarParserPOINTS  = 53
	GramarParserBRK     = 54
	GramarParserID      = 55
	GramarParserINT     = 56
	GramarParserDOUBLE  = 57
	GramarParserCADENA  = 58
	GramarParserCHAR    = 59
)

// GramarParser rules.
const (
	GramarParserRULE_s           = 0
	GramarParserRULE_block       = 1
	GramarParserRULE_productions = 2
	GramarParserRULE_pbrk        = 3
	GramarParserRULE_pswitch     = 4
	GramarParserRULE_ccase       = 5
	GramarParserRULE_pdefault    = 6
	GramarParserRULE_pfor        = 7
	GramarParserRULE_prin        = 8
	GramarParserRULE_swhile      = 9
	GramarParserRULE_asign       = 10
	GramarParserRULE_pdeclara    = 11
	GramarParserRULE_tipo        = 12
	GramarParserRULE_pif         = 13
	GramarParserRULE_pelse       = 14
	GramarParserRULE_pfuncion    = 15
	GramarParserRULE_pparams     = 16
	GramarParserRULE_pparamet    = 17
	GramarParserRULE_pllamada    = 18
	GramarParserRULE_argumento   = 19
	GramarParserRULE_tipoArg     = 20
	GramarParserRULE_expr        = 21
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SContext) EOF() antlr.TerminalNode {
	return s.GetToken(GramarParserEOF, 0)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitS(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GramarParserRULE_s)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Block()
	}
	{
		p.SetState(45)
		p.Match(GramarParserEOF)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllProductions() []IProductionsContext
	Productions(i int) IProductionsContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllProductions() []IProductionsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProductionsContext); ok {
			len++
		}
	}

	tst := make([]IProductionsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProductionsContext); ok {
			tst[i] = t.(IProductionsContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Productions(i int) IProductionsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProductionsContext); ok {
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

	return t.(IProductionsContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GramarParserRULE_block)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(47)
				p.Productions()
			}

		}
		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
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

// IProductionsContext is an interface to support dynamic dispatch.
type IProductionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Prin() IPrinContext
	Swhile() ISwhileContext
	Pif() IPifContext
	Pdeclara() IPdeclaraContext
	Asign() IAsignContext
	Pfor() IPforContext
	Pswitch() IPswitchContext
	Pfuncion() IPfuncionContext
	Pllamada() IPllamadaContext
	Pbrk() IPbrkContext

	// IsProductionsContext differentiates from other interfaces.
	IsProductionsContext()
}

type ProductionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProductionsContext() *ProductionsContext {
	var p = new(ProductionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_productions
	return p
}

func InitEmptyProductionsContext(p *ProductionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_productions
}

func (*ProductionsContext) IsProductionsContext() {}

func NewProductionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProductionsContext {
	var p = new(ProductionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_productions

	return p
}

func (s *ProductionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ProductionsContext) Prin() IPrinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrinContext)
}

func (s *ProductionsContext) Swhile() ISwhileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwhileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwhileContext)
}

func (s *ProductionsContext) Pif() IPifContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPifContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPifContext)
}

func (s *ProductionsContext) Pdeclara() IPdeclaraContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPdeclaraContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPdeclaraContext)
}

func (s *ProductionsContext) Asign() IAsignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignContext)
}

func (s *ProductionsContext) Pfor() IPforContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPforContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPforContext)
}

func (s *ProductionsContext) Pswitch() IPswitchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPswitchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPswitchContext)
}

func (s *ProductionsContext) Pfuncion() IPfuncionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPfuncionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPfuncionContext)
}

func (s *ProductionsContext) Pllamada() IPllamadaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPllamadaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPllamadaContext)
}

func (s *ProductionsContext) Pbrk() IPbrkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPbrkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPbrkContext)
}

func (s *ProductionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProductionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProductionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitProductions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Productions() (localctx IProductionsContext) {
	localctx = NewProductionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GramarParserRULE_productions)
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Prin()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Swhile()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)
			p.Pif()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)
			p.Pdeclara()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(57)
			p.Asign()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(58)
			p.Pfor()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(59)
			p.Pswitch()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(60)
			p.Pfuncion()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(61)
			p.Pllamada()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(62)
			p.Pbrk()
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

// IPbrkContext is an interface to support dynamic dispatch.
type IPbrkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BRK() antlr.TerminalNode

	// IsPbrkContext differentiates from other interfaces.
	IsPbrkContext()
}

type PbrkContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPbrkContext() *PbrkContext {
	var p = new(PbrkContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pbrk
	return p
}

func InitEmptyPbrkContext(p *PbrkContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pbrk
}

func (*PbrkContext) IsPbrkContext() {}

func NewPbrkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PbrkContext {
	var p = new(PbrkContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pbrk

	return p
}

func (s *PbrkContext) GetParser() antlr.Parser { return s.parser }

func (s *PbrkContext) BRK() antlr.TerminalNode {
	return s.GetToken(GramarParserBRK, 0)
}

func (s *PbrkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PbrkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PbrkContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPbrk(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pbrk() (localctx IPbrkContext) {
	localctx = NewPbrkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GramarParserRULE_pbrk)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(GramarParserBRK)
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

// IPswitchContext is an interface to support dynamic dispatch.
type IPswitchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Pdefault() IPdefaultContext
	AllCcase() []ICcaseContext
	Ccase(i int) ICcaseContext

	// IsPswitchContext differentiates from other interfaces.
	IsPswitchContext()
}

type PswitchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPswitchContext() *PswitchContext {
	var p = new(PswitchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pswitch
	return p
}

func InitEmptyPswitchContext(p *PswitchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pswitch
}

func (*PswitchContext) IsPswitchContext() {}

func NewPswitchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PswitchContext {
	var p = new(PswitchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pswitch

	return p
}

func (s *PswitchContext) GetParser() antlr.Parser { return s.parser }

func (s *PswitchContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PswitchContext) Pdefault() IPdefaultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPdefaultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPdefaultContext)
}

func (s *PswitchContext) AllCcase() []ICcaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICcaseContext); ok {
			len++
		}
	}

	tst := make([]ICcaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICcaseContext); ok {
			tst[i] = t.(ICcaseContext)
			i++
		}
	}

	return tst
}

func (s *PswitchContext) Ccase(i int) ICcaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICcaseContext); ok {
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

	return t.(ICcaseContext)
}

func (s *PswitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PswitchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PswitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPswitch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pswitch() (localctx IPswitchContext) {
	localctx = NewPswitchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GramarParserRULE_pswitch)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(GramarParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.expr(0)
	}
	{
		p.SetState(69)
		p.Match(GramarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GramarParserT__3 {
		{
			p.SetState(70)
			p.Ccase()
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(75)
		p.Pdefault()
	}
	{
		p.SetState(76)
		p.Match(GramarParserT__2)
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

// ICcaseContext is an interface to support dynamic dispatch.
type ICcaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Block() IBlockContext
	BRK() antlr.TerminalNode

	// IsCcaseContext differentiates from other interfaces.
	IsCcaseContext()
}

type CcaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCcaseContext() *CcaseContext {
	var p = new(CcaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_ccase
	return p
}

func InitEmptyCcaseContext(p *CcaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_ccase
}

func (*CcaseContext) IsCcaseContext() {}

func NewCcaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CcaseContext {
	var p = new(CcaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_ccase

	return p
}

func (s *CcaseContext) GetParser() antlr.Parser { return s.parser }

func (s *CcaseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CcaseContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CcaseContext) BRK() antlr.TerminalNode {
	return s.GetToken(GramarParserBRK, 0)
}

func (s *CcaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CcaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CcaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitCcase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Ccase() (localctx ICcaseContext) {
	localctx = NewCcaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GramarParserRULE_ccase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(GramarParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.expr(0)
	}
	{
		p.SetState(80)
		p.Match(GramarParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Block()
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramarParserBRK {
		{
			p.SetState(82)
			p.Match(GramarParserBRK)
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

// IPdefaultContext is an interface to support dynamic dispatch.
type IPdefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	BRK() antlr.TerminalNode

	// IsPdefaultContext differentiates from other interfaces.
	IsPdefaultContext()
}

type PdefaultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPdefaultContext() *PdefaultContext {
	var p = new(PdefaultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pdefault
	return p
}

func InitEmptyPdefaultContext(p *PdefaultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pdefault
}

func (*PdefaultContext) IsPdefaultContext() {}

func NewPdefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PdefaultContext {
	var p = new(PdefaultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pdefault

	return p
}

func (s *PdefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *PdefaultContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *PdefaultContext) BRK() antlr.TerminalNode {
	return s.GetToken(GramarParserBRK, 0)
}

func (s *PdefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PdefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PdefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPdefault(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pdefault() (localctx IPdefaultContext) {
	localctx = NewPdefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GramarParserRULE_pdefault)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(GramarParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Match(GramarParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Block()
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramarParserBRK {
		{
			p.SetState(88)
			p.Match(GramarParserBRK)
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

// IPforContext is an interface to support dynamic dispatch.
type IPforContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Block() IBlockContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	POINTS() antlr.TerminalNode

	// IsPforContext differentiates from other interfaces.
	IsPforContext()
}

type PforContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPforContext() *PforContext {
	var p = new(PforContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pfor
	return p
}

func InitEmptyPforContext(p *PforContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pfor
}

func (*PforContext) IsPforContext() {}

func NewPforContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PforContext {
	var p = new(PforContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pfor

	return p
}

func (s *PforContext) GetParser() antlr.Parser { return s.parser }

func (s *PforContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *PforContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *PforContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *PforContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *PforContext) POINTS() antlr.TerminalNode {
	return s.GetToken(GramarParserPOINTS, 0)
}

func (s *PforContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PforContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PforContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPfor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pfor() (localctx IPforContext) {
	localctx = NewPforContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GramarParserRULE_pfor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(GramarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
		p.Match(GramarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(GramarParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(94)
			p.expr(0)
		}
		{
			p.SetState(95)
			p.Match(GramarParserPOINTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.expr(0)
		}

	case 2:
		{
			p.SetState(98)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(101)
		p.Match(GramarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Block()
	}
	{
		p.SetState(103)
		p.Match(GramarParserT__2)
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

// IPrinContext is an interface to support dynamic dispatch.
type IPrinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsPrinContext differentiates from other interfaces.
	IsPrinContext()
}

type PrinContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrinContext() *PrinContext {
	var p = new(PrinContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_prin
	return p
}

func InitEmptyPrinContext(p *PrinContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_prin
}

func (*PrinContext) IsPrinContext() {}

func NewPrinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrinContext {
	var p = new(PrinContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_prin

	return p
}

func (s *PrinContext) GetParser() antlr.Parser { return s.parser }

func (s *PrinContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPrin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Prin() (localctx IPrinContext) {
	localctx = NewPrinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GramarParserRULE_prin)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(GramarParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(GramarParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.expr(0)
	}
	{
		p.SetState(108)
		p.Match(GramarParserT__10)
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

// ISwhileContext is an interface to support dynamic dispatch.
type ISwhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Block() IBlockContext

	// IsSwhileContext differentiates from other interfaces.
	IsSwhileContext()
}

type SwhileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwhileContext() *SwhileContext {
	var p = new(SwhileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_swhile
	return p
}

func InitEmptySwhileContext(p *SwhileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_swhile
}

func (*SwhileContext) IsSwhileContext() {}

func NewSwhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwhileContext {
	var p = new(SwhileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_swhile

	return p
}

func (s *SwhileContext) GetParser() antlr.Parser { return s.parser }

func (s *SwhileContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwhileContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SwhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitSwhile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Swhile() (localctx ISwhileContext) {
	localctx = NewSwhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GramarParserRULE_swhile)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(GramarParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.expr(0)
	}
	{
		p.SetState(112)
		p.Match(GramarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Block()
	}
	{
		p.SetState(114)
		p.Match(GramarParserT__2)
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

// IAsignContext is an interface to support dynamic dispatch.
type IAsignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	Expr() IExprContext

	// IsAsignContext differentiates from other interfaces.
	IsAsignContext()
}

type AsignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyAsignContext() *AsignContext {
	var p = new(AsignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_asign
	return p
}

func InitEmptyAsignContext(p *AsignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_asign
}

func (*AsignContext) IsAsignContext() {}

func NewAsignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignContext {
	var p = new(AsignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_asign

	return p
}

func (s *AsignContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignContext) GetOp() antlr.Token { return s.op }

func (s *AsignContext) SetOp(v antlr.Token) { s.op = v }

func (s *AsignContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *AsignContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitAsign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Asign() (localctx IAsignContext) {
	localctx = NewAsignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GramarParserRULE_asign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(GramarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*AsignContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&57344) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*AsignContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(118)
		p.expr(0)
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

// IPdeclaraContext is an interface to support dynamic dispatch.
type IPdeclaraContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Tipo() ITipoContext
	Expr() IExprContext
	RVAR() antlr.TerminalNode
	RLET() antlr.TerminalNode

	// IsPdeclaraContext differentiates from other interfaces.
	IsPdeclaraContext()
}

type PdeclaraContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPdeclaraContext() *PdeclaraContext {
	var p = new(PdeclaraContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pdeclara
	return p
}

func InitEmptyPdeclaraContext(p *PdeclaraContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pdeclara
}

func (*PdeclaraContext) IsPdeclaraContext() {}

func NewPdeclaraContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PdeclaraContext {
	var p = new(PdeclaraContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pdeclara

	return p
}

func (s *PdeclaraContext) GetParser() antlr.Parser { return s.parser }

func (s *PdeclaraContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *PdeclaraContext) Tipo() ITipoContext {
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

func (s *PdeclaraContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PdeclaraContext) RVAR() antlr.TerminalNode {
	return s.GetToken(GramarParserRVAR, 0)
}

func (s *PdeclaraContext) RLET() antlr.TerminalNode {
	return s.GetToken(GramarParserRLET, 0)
}

func (s *PdeclaraContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PdeclaraContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PdeclaraContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPdeclara(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pdeclara() (localctx IPdeclaraContext) {
	localctx = NewPdeclaraContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GramarParserRULE_pdeclara)
	var _la int

	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserRVAR || _la == GramarParserRLET) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(121)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Tipo()
		}
		{
			p.SetState(124)
			p.Match(GramarParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.expr(0)
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__15 {
			{
				p.SetState(126)
				p.Match(GramarParserT__15)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserRVAR || _la == GramarParserRLET) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(130)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Tipo()
		}
		{
			p.SetState(133)
			p.Match(GramarParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(135)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserRVAR || _la == GramarParserRLET) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(136)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.Match(GramarParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.expr(0)
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

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RINT() antlr.TerminalNode
	RFLOAT() antlr.TerminalNode
	RSTRING() antlr.TerminalNode
	RBOOL() antlr.TerminalNode
	RChar() antlr.TerminalNode

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
	p.RuleIndex = GramarParserRULE_tipo
	return p
}

func InitEmptyTipoContext(p *TipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_tipo
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) RINT() antlr.TerminalNode {
	return s.GetToken(GramarParserRINT, 0)
}

func (s *TipoContext) RFLOAT() antlr.TerminalNode {
	return s.GetToken(GramarParserRFLOAT, 0)
}

func (s *TipoContext) RSTRING() antlr.TerminalNode {
	return s.GetToken(GramarParserRSTRING, 0)
}

func (s *TipoContext) RBOOL() antlr.TerminalNode {
	return s.GetToken(GramarParserRBOOL, 0)
}

func (s *TipoContext) RChar() antlr.TerminalNode {
	return s.GetToken(GramarParserRChar, 0)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitTipo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GramarParserRULE_tipo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&272678883688448) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// IPifContext is an interface to support dynamic dispatch.
type IPifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Block() IBlockContext
	Pelse() IPelseContext

	// IsPifContext differentiates from other interfaces.
	IsPifContext()
}

type PifContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPifContext() *PifContext {
	var p = new(PifContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pif
	return p
}

func InitEmptyPifContext(p *PifContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pif
}

func (*PifContext) IsPifContext() {}

func NewPifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PifContext {
	var p = new(PifContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pif

	return p
}

func (s *PifContext) GetParser() antlr.Parser { return s.parser }

func (s *PifContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PifContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *PifContext) Pelse() IPelseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPelseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPelseContext)
}

func (s *PifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PifContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPif(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pif() (localctx IPifContext) {
	localctx = NewPifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GramarParserRULE_pif)
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(143)
			p.Match(GramarParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.expr(0)
		}
		{
			p.SetState(145)
			p.Match(GramarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Block()
		}
		{
			p.SetState(147)
			p.Match(GramarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(149)
			p.Match(GramarParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.expr(0)
		}
		{
			p.SetState(151)
			p.Match(GramarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.Block()
		}
		{
			p.SetState(153)
			p.Match(GramarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Pelse()
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

// IPelseContext is an interface to support dynamic dispatch.
type IPelseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Pif() IPifContext
	Block() IBlockContext

	// IsPelseContext differentiates from other interfaces.
	IsPelseContext()
}

type PelseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPelseContext() *PelseContext {
	var p = new(PelseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pelse
	return p
}

func InitEmptyPelseContext(p *PelseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pelse
}

func (*PelseContext) IsPelseContext() {}

func NewPelseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PelseContext {
	var p = new(PelseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pelse

	return p
}

func (s *PelseContext) GetParser() antlr.Parser { return s.parser }

func (s *PelseContext) Pif() IPifContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPifContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPifContext)
}

func (s *PelseContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *PelseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PelseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PelseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPelse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pelse() (localctx IPelseContext) {
	localctx = NewPelseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GramarParserRULE_pelse)
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Match(GramarParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.Pif()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(160)
			p.Match(GramarParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Match(GramarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Block()
		}
		{
			p.SetState(163)
			p.Match(GramarParserT__2)
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

// IPfuncionContext is an interface to support dynamic dispatch.
type IPfuncionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Block() IBlockContext
	Tipo() ITipoContext
	Pparams() IPparamsContext

	// IsPfuncionContext differentiates from other interfaces.
	IsPfuncionContext()
}

type PfuncionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPfuncionContext() *PfuncionContext {
	var p = new(PfuncionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pfuncion
	return p
}

func InitEmptyPfuncionContext(p *PfuncionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pfuncion
}

func (*PfuncionContext) IsPfuncionContext() {}

func NewPfuncionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PfuncionContext {
	var p = new(PfuncionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pfuncion

	return p
}

func (s *PfuncionContext) GetParser() antlr.Parser { return s.parser }

func (s *PfuncionContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *PfuncionContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *PfuncionContext) Tipo() ITipoContext {
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

func (s *PfuncionContext) Pparams() IPparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPparamsContext)
}

func (s *PfuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PfuncionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PfuncionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPfuncion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pfuncion() (localctx IPfuncionContext) {
	localctx = NewPfuncionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GramarParserRULE_pfuncion)
	var _la int

	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.Match(GramarParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Match(GramarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.Match(GramarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__20 {
			{
				p.SetState(171)
				p.Match(GramarParserT__20)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(172)
				p.Tipo()
			}

		}
		{
			p.SetState(175)
			p.Match(GramarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Block()
		}
		{
			p.SetState(177)
			p.Match(GramarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(179)
			p.Match(GramarParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)
			p.Match(GramarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Pparams()
		}
		{
			p.SetState(183)
			p.Match(GramarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__20 {
			{
				p.SetState(184)
				p.Match(GramarParserT__20)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(185)
				p.Tipo()
			}

		}
		{
			p.SetState(188)
			p.Match(GramarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.Block()
		}
		{
			p.SetState(190)
			p.Match(GramarParserT__2)
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

// IPparamsContext is an interface to support dynamic dispatch.
type IPparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPparamet() []IPparametContext
	Pparamet(i int) IPparametContext

	// IsPparamsContext differentiates from other interfaces.
	IsPparamsContext()
}

type PparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPparamsContext() *PparamsContext {
	var p = new(PparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pparams
	return p
}

func InitEmptyPparamsContext(p *PparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pparams
}

func (*PparamsContext) IsPparamsContext() {}

func NewPparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PparamsContext {
	var p = new(PparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pparams

	return p
}

func (s *PparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *PparamsContext) AllPparamet() []IPparametContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPparametContext); ok {
			len++
		}
	}

	tst := make([]IPparametContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPparametContext); ok {
			tst[i] = t.(IPparametContext)
			i++
		}
	}

	return tst
}

func (s *PparamsContext) Pparamet(i int) IPparametContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPparametContext); ok {
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

	return t.(IPparametContext)
}

func (s *PparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PparamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPparams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pparams() (localctx IPparamsContext) {
	localctx = NewPparamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GramarParserRULE_pparams)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramarParserGUION || _la == GramarParserID {
		{
			p.SetState(194)
			p.Pparamet()
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GramarParserT__21 {
			{
				p.SetState(195)
				p.Match(GramarParserT__21)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(196)
				p.Pparamet()
			}

			p.SetState(201)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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

// IPparametContext is an interface to support dynamic dispatch.
type IPparametContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPid returns the pid token.
	GetPid() antlr.Token

	// SetPid sets the pid token.
	SetPid(antlr.Token)

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	Tipo() ITipoContext
	GUION() antlr.TerminalNode
	INOUT() antlr.TerminalNode
	ABR() antlr.TerminalNode
	CIER() antlr.TerminalNode

	// IsPparametContext differentiates from other interfaces.
	IsPparametContext()
}

type PparametContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	pid    antlr.Token
}

func NewEmptyPparametContext() *PparametContext {
	var p = new(PparametContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pparamet
	return p
}

func InitEmptyPparametContext(p *PparametContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pparamet
}

func (*PparametContext) IsPparametContext() {}

func NewPparametContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PparametContext {
	var p = new(PparametContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pparamet

	return p
}

func (s *PparametContext) GetParser() antlr.Parser { return s.parser }

func (s *PparametContext) GetPid() antlr.Token { return s.pid }

func (s *PparametContext) SetPid(v antlr.Token) { s.pid = v }

func (s *PparametContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GramarParserID)
}

func (s *PparametContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GramarParserID, i)
}

func (s *PparametContext) Tipo() ITipoContext {
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

func (s *PparametContext) GUION() antlr.TerminalNode {
	return s.GetToken(GramarParserGUION, 0)
}

func (s *PparametContext) INOUT() antlr.TerminalNode {
	return s.GetToken(GramarParserINOUT, 0)
}

func (s *PparametContext) ABR() antlr.TerminalNode {
	return s.GetToken(GramarParserABR, 0)
}

func (s *PparametContext) CIER() antlr.TerminalNode {
	return s.GetToken(GramarParserCIER, 0)
}

func (s *PparametContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PparametContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PparametContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPparamet(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pparamet() (localctx IPparametContext) {
	localctx = NewPparametContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GramarParserRULE_pparamet)
	var _la int

	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(204)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*PparametContext).pid = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserGUION || _la == GramarParserID) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*PparametContext).pid = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(205)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserINOUT {
			{
				p.SetState(207)
				p.Match(GramarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(210)
			p.Tipo()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(211)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*PparametContext).pid = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserGUION || _la == GramarParserID) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*PparametContext).pid = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(212)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserINOUT {
			{
				p.SetState(214)
				p.Match(GramarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(217)
			p.Match(GramarParserABR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.Tipo()
		}
		{
			p.SetState(219)
			p.Match(GramarParserCIER)
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

// IPllamadaContext is an interface to support dynamic dispatch.
type IPllamadaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Argumento() IArgumentoContext

	// IsPllamadaContext differentiates from other interfaces.
	IsPllamadaContext()
}

type PllamadaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPllamadaContext() *PllamadaContext {
	var p = new(PllamadaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pllamada
	return p
}

func InitEmptyPllamadaContext(p *PllamadaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pllamada
}

func (*PllamadaContext) IsPllamadaContext() {}

func NewPllamadaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PllamadaContext {
	var p = new(PllamadaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pllamada

	return p
}

func (s *PllamadaContext) GetParser() antlr.Parser { return s.parser }

func (s *PllamadaContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *PllamadaContext) Argumento() IArgumentoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentoContext)
}

func (s *PllamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PllamadaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PllamadaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPllamada(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pllamada() (localctx IPllamadaContext) {
	localctx = NewPllamadaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GramarParserRULE_pllamada)
	var _la int

	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.Match(GramarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.Match(GramarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__15 {
			{
				p.SetState(226)
				p.Match(GramarParserT__15)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)
			p.Match(GramarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Argumento()
		}
		{
			p.SetState(232)
			p.Match(GramarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__15 {
			{
				p.SetState(233)
				p.Match(GramarParserT__15)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
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

// IArgumentoContext is an interface to support dynamic dispatch.
type IArgumentoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TipoArg() ITipoArgContext
	Argumento() IArgumentoContext

	// IsArgumentoContext differentiates from other interfaces.
	IsArgumentoContext()
}

type ArgumentoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentoContext() *ArgumentoContext {
	var p = new(ArgumentoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_argumento
	return p
}

func InitEmptyArgumentoContext(p *ArgumentoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_argumento
}

func (*ArgumentoContext) IsArgumentoContext() {}

func NewArgumentoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentoContext {
	var p = new(ArgumentoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_argumento

	return p
}

func (s *ArgumentoContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentoContext) TipoArg() ITipoArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoArgContext)
}

func (s *ArgumentoContext) Argumento() IArgumentoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentoContext)
}

func (s *ArgumentoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitArgumento(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Argumento() (localctx IArgumentoContext) {
	localctx = NewArgumentoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GramarParserRULE_argumento)
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(238)
			p.TipoArg()
		}
		{
			p.SetState(239)
			p.Match(GramarParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)
			p.Argumento()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)
			p.TipoArg()
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

// ITipoArgContext is an interface to support dynamic dispatch.
type ITipoArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExp1 returns the exp1 rule contexts.
	GetExp1() IExprContext

	// GetExp2 returns the exp2 rule contexts.
	GetExp2() IExprContext

	// GetExp3 returns the exp3 rule contexts.
	GetExp3() IExprContext

	// SetExp1 sets the exp1 rule contexts.
	SetExp1(IExprContext)

	// SetExp2 sets the exp2 rule contexts.
	SetExp2(IExprContext)

	// SetExp3 sets the exp3 rule contexts.
	SetExp3(IExprContext)

	// Getter signatures
	PUNTE() antlr.TerminalNode
	Expr() IExprContext
	ID() antlr.TerminalNode

	// IsTipoArgContext differentiates from other interfaces.
	IsTipoArgContext()
}

type TipoArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	exp1   IExprContext
	exp2   IExprContext
	exp3   IExprContext
}

func NewEmptyTipoArgContext() *TipoArgContext {
	var p = new(TipoArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_tipoArg
	return p
}

func InitEmptyTipoArgContext(p *TipoArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_tipoArg
}

func (*TipoArgContext) IsTipoArgContext() {}

func NewTipoArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoArgContext {
	var p = new(TipoArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_tipoArg

	return p
}

func (s *TipoArgContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoArgContext) GetExp1() IExprContext { return s.exp1 }

func (s *TipoArgContext) GetExp2() IExprContext { return s.exp2 }

func (s *TipoArgContext) GetExp3() IExprContext { return s.exp3 }

func (s *TipoArgContext) SetExp1(v IExprContext) { s.exp1 = v }

func (s *TipoArgContext) SetExp2(v IExprContext) { s.exp2 = v }

func (s *TipoArgContext) SetExp3(v IExprContext) { s.exp3 = v }

func (s *TipoArgContext) PUNTE() antlr.TerminalNode {
	return s.GetToken(GramarParserPUNTE, 0)
}

func (s *TipoArgContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TipoArgContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *TipoArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitTipoArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) TipoArg() (localctx ITipoArgContext) {
	localctx = NewTipoArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GramarParserRULE_tipoArg)
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(245)
			p.Match(GramarParserPUNTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)

			var _x = p.expr(0)

			localctx.(*TipoArgContext).exp1 = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(247)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(248)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)

			var _x = p.expr(0)

			localctx.(*TipoArgContext).exp2 = _x
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(250)

			var _x = p.expr(0)

			localctx.(*TipoArgContext).exp3 = _x
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParentContext struct {
	ExprContext
}

func NewParentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentContext {
	var p = new(ParentContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitParent(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralContext struct {
	ExprContext
}

func NewLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralContext {
	var p = new(LiteralContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(GramarParserINT, 0)
}

func (s *LiteralContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(GramarParserDOUBLE, 0)
}

func (s *LiteralContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(GramarParserBOOLEAN, 0)
}

func (s *LiteralContext) CADENA() antlr.TerminalNode {
	return s.GetToken(GramarParserCADENA, 0)
}

func (s *LiteralContext) CHAR() antlr.TerminalNode {
	return s.GetToken(GramarParserCHAR, 0)
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type AccesoContext struct {
	ExprContext
}

func NewAccesoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccesoContext {
	var p = new(AccesoContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AccesoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesoContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *AccesoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitAcceso(s)

	default:
		return t.VisitChildren(s)
	}
}

type AritmeticContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewAritmeticContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AritmeticContext {
	var p = new(AritmeticContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AritmeticContext) GetOp() antlr.Token { return s.op }

func (s *AritmeticContext) SetOp(v antlr.Token) { s.op = v }

func (s *AritmeticContext) GetLeft() IExprContext { return s.left }

func (s *AritmeticContext) GetRight() IExprContext { return s.right }

func (s *AritmeticContext) SetLeft(v IExprContext) { s.left = v }

func (s *AritmeticContext) SetRight(v IExprContext) { s.right = v }

func (s *AritmeticContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AritmeticContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AritmeticContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *AritmeticContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitAritmetic(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *GramarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, GramarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GramarParserT__22, GramarParserT__23:
		localctx = NewAritmeticContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(254)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AritmeticContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserT__22 || _la == GramarParserT__23) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AritmeticContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(255)

			var _x = p.expr(14)

			localctx.(*AritmeticContext).right = _x
		}

	case GramarParserT__9:
		localctx = NewParentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(256)
			p.Match(GramarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(257)
			p.expr(0)
		}
		{
			p.SetState(258)
			p.Match(GramarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GramarParserINT:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(260)
			p.Match(GramarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GramarParserDOUBLE:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(261)
			p.Match(GramarParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GramarParserBOOLEAN:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(262)
			p.Match(GramarParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GramarParserCADENA:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(263)
			p.Match(GramarParserCADENA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GramarParserCHAR:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(264)
			p.Match(GramarParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GramarParserID:
		localctx = NewAccesoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(265)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(286)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAritmeticContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AritmeticContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(268)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(269)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AritmeticContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&234881024) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AritmeticContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(270)

					var _x = p.expr(13)

					localctx.(*AritmeticContext).right = _x
				}

			case 2:
				localctx = NewAritmeticContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AritmeticContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(271)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(272)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AritmeticContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramarParserT__23 || _la == GramarParserT__27) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AritmeticContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(273)

					var _x = p.expr(12)

					localctx.(*AritmeticContext).right = _x
				}

			case 3:
				localctx = NewAritmeticContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AritmeticContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(274)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(275)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AritmeticContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8053063680) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AritmeticContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(276)

					var _x = p.expr(11)

					localctx.(*AritmeticContext).right = _x
				}

			case 4:
				localctx = NewAritmeticContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AritmeticContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(277)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(278)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AritmeticContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramarParserT__32 || _la == GramarParserT__33) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AritmeticContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(279)

					var _x = p.expr(10)

					localctx.(*AritmeticContext).right = _x
				}

			case 5:
				localctx = NewAritmeticContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AritmeticContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(280)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(281)

					var _m = p.Match(GramarParserT__34)

					localctx.(*AritmeticContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(282)

					var _x = p.expr(9)

					localctx.(*AritmeticContext).right = _x
				}

			case 6:
				localctx = NewAritmeticContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AritmeticContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(283)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(284)

					var _m = p.Match(GramarParserT__35)

					localctx.(*AritmeticContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(285)

					var _x = p.expr(8)

					localctx.(*AritmeticContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
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

func (p *GramarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 21:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GramarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
