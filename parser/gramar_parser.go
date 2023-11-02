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
		"'in'", "'print'", "'('", "')'", "';'", "','", "'while'", "'='", "'+='",
		"'-='", "'?'", "'if'", "'else'", "'func'", "'->'", "'!'", "'-'", "'/'",
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
		"pfor", "prin", "cexpr", "swhile", "asign", "pdeclara", "tipo", "pif",
		"pelse", "pfuncion", "pparams", "pparamet", "pllamada", "argumento",
		"tipoArg", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 59, 310, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 51, 8, 1, 10, 1,
		12, 1, 54, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 3, 2, 66, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 4, 4, 74, 8,
		4, 11, 4, 12, 4, 75, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3,
		5, 86, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 92, 8, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 102, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 118,
		8, 8, 3, 8, 120, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 127, 8, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 146, 8, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 158,
		8, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 175, 8, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 184, 8, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 3, 16, 192, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 205, 8, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 3, 16, 211, 8, 16, 1, 17, 1, 17, 1, 17, 5, 17,
		216, 8, 17, 10, 17, 12, 17, 219, 9, 17, 3, 17, 221, 8, 17, 1, 18, 1, 18,
		1, 18, 1, 18, 3, 18, 227, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3,
		18, 234, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 240, 8, 18, 1, 19, 1,
		19, 1, 19, 1, 19, 3, 19, 246, 8, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		3, 19, 253, 8, 19, 3, 19, 255, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		3, 20, 262, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 270,
		8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 3, 22, 285, 8, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 5, 22, 305, 8, 22, 10, 22, 12, 22, 308, 9, 22,
		1, 22, 0, 1, 44, 23, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 44, 0, 9, 1, 0, 15, 17, 1, 0, 41, 42, 1,
		0, 43, 47, 2, 0, 48, 48, 55, 55, 1, 0, 23, 24, 1, 0, 25, 27, 2, 0, 24,
		24, 28, 28, 1, 0, 29, 32, 1, 0, 33, 34, 335, 0, 46, 1, 0, 0, 0, 2, 52,
		1, 0, 0, 0, 4, 65, 1, 0, 0, 0, 6, 67, 1, 0, 0, 0, 8, 69, 1, 0, 0, 0, 10,
		80, 1, 0, 0, 0, 12, 87, 1, 0, 0, 0, 14, 93, 1, 0, 0, 0, 16, 119, 1, 0,
		0, 0, 18, 126, 1, 0, 0, 0, 20, 128, 1, 0, 0, 0, 22, 134, 1, 0, 0, 0, 24,
		157, 1, 0, 0, 0, 26, 159, 1, 0, 0, 0, 28, 174, 1, 0, 0, 0, 30, 183, 1,
		0, 0, 0, 32, 210, 1, 0, 0, 0, 34, 220, 1, 0, 0, 0, 36, 239, 1, 0, 0, 0,
		38, 254, 1, 0, 0, 0, 40, 261, 1, 0, 0, 0, 42, 269, 1, 0, 0, 0, 44, 284,
		1, 0, 0, 0, 46, 47, 3, 2, 1, 0, 47, 48, 5, 0, 0, 1, 48, 1, 1, 0, 0, 0,
		49, 51, 3, 4, 2, 0, 50, 49, 1, 0, 0, 0, 51, 54, 1, 0, 0, 0, 52, 50, 1,
		0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 3, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 55,
		66, 3, 16, 8, 0, 56, 66, 3, 20, 10, 0, 57, 66, 3, 28, 14, 0, 58, 66, 3,
		24, 12, 0, 59, 66, 3, 22, 11, 0, 60, 66, 3, 14, 7, 0, 61, 66, 3, 8, 4,
		0, 62, 66, 3, 32, 16, 0, 63, 66, 3, 38, 19, 0, 64, 66, 3, 6, 3, 0, 65,
		55, 1, 0, 0, 0, 65, 56, 1, 0, 0, 0, 65, 57, 1, 0, 0, 0, 65, 58, 1, 0, 0,
		0, 65, 59, 1, 0, 0, 0, 65, 60, 1, 0, 0, 0, 65, 61, 1, 0, 0, 0, 65, 62,
		1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 64, 1, 0, 0, 0, 66, 5, 1, 0, 0, 0,
		67, 68, 5, 54, 0, 0, 68, 7, 1, 0, 0, 0, 69, 70, 5, 1, 0, 0, 70, 71, 3,
		44, 22, 0, 71, 73, 5, 2, 0, 0, 72, 74, 3, 10, 5, 0, 73, 72, 1, 0, 0, 0,
		74, 75, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 77, 1,
		0, 0, 0, 77, 78, 3, 12, 6, 0, 78, 79, 5, 3, 0, 0, 79, 9, 1, 0, 0, 0, 80,
		81, 5, 4, 0, 0, 81, 82, 3, 44, 22, 0, 82, 83, 5, 5, 0, 0, 83, 85, 3, 2,
		1, 0, 84, 86, 5, 54, 0, 0, 85, 84, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86,
		11, 1, 0, 0, 0, 87, 88, 5, 6, 0, 0, 88, 89, 5, 5, 0, 0, 89, 91, 3, 2, 1,
		0, 90, 92, 5, 54, 0, 0, 91, 90, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 13,
		1, 0, 0, 0, 93, 94, 5, 7, 0, 0, 94, 95, 5, 55, 0, 0, 95, 101, 5, 8, 0,
		0, 96, 97, 3, 44, 22, 0, 97, 98, 5, 53, 0, 0, 98, 99, 3, 44, 22, 0, 99,
		102, 1, 0, 0, 0, 100, 102, 3, 44, 22, 0, 101, 96, 1, 0, 0, 0, 101, 100,
		1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 5, 2, 0, 0, 104, 105, 3, 2,
		1, 0, 105, 106, 5, 3, 0, 0, 106, 15, 1, 0, 0, 0, 107, 108, 5, 9, 0, 0,
		108, 109, 5, 10, 0, 0, 109, 110, 3, 44, 22, 0, 110, 111, 5, 11, 0, 0, 111,
		120, 1, 0, 0, 0, 112, 113, 5, 9, 0, 0, 113, 114, 5, 10, 0, 0, 114, 115,
		3, 18, 9, 0, 115, 117, 5, 11, 0, 0, 116, 118, 5, 12, 0, 0, 117, 116, 1,
		0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 120, 1, 0, 0, 0, 119, 107, 1, 0, 0,
		0, 119, 112, 1, 0, 0, 0, 120, 17, 1, 0, 0, 0, 121, 122, 3, 44, 22, 0, 122,
		123, 5, 13, 0, 0, 123, 124, 3, 18, 9, 0, 124, 127, 1, 0, 0, 0, 125, 127,
		3, 44, 22, 0, 126, 121, 1, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 19, 1, 0,
		0, 0, 128, 129, 5, 14, 0, 0, 129, 130, 3, 44, 22, 0, 130, 131, 5, 2, 0,
		0, 131, 132, 3, 2, 1, 0, 132, 133, 5, 3, 0, 0, 133, 21, 1, 0, 0, 0, 134,
		135, 5, 55, 0, 0, 135, 136, 7, 0, 0, 0, 136, 137, 3, 44, 22, 0, 137, 23,
		1, 0, 0, 0, 138, 139, 7, 1, 0, 0, 139, 140, 5, 55, 0, 0, 140, 141, 5, 5,
		0, 0, 141, 142, 3, 26, 13, 0, 142, 143, 5, 15, 0, 0, 143, 145, 3, 44, 22,
		0, 144, 146, 5, 12, 0, 0, 145, 144, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146,
		158, 1, 0, 0, 0, 147, 148, 7, 1, 0, 0, 148, 149, 5, 55, 0, 0, 149, 150,
		5, 5, 0, 0, 150, 151, 3, 26, 13, 0, 151, 152, 5, 18, 0, 0, 152, 158, 1,
		0, 0, 0, 153, 154, 7, 1, 0, 0, 154, 155, 5, 55, 0, 0, 155, 156, 5, 15,
		0, 0, 156, 158, 3, 44, 22, 0, 157, 138, 1, 0, 0, 0, 157, 147, 1, 0, 0,
		0, 157, 153, 1, 0, 0, 0, 158, 25, 1, 0, 0, 0, 159, 160, 7, 2, 0, 0, 160,
		27, 1, 0, 0, 0, 161, 162, 5, 19, 0, 0, 162, 163, 3, 44, 22, 0, 163, 164,
		5, 2, 0, 0, 164, 165, 3, 2, 1, 0, 165, 166, 5, 3, 0, 0, 166, 175, 1, 0,
		0, 0, 167, 168, 5, 19, 0, 0, 168, 169, 3, 44, 22, 0, 169, 170, 5, 2, 0,
		0, 170, 171, 3, 2, 1, 0, 171, 172, 5, 3, 0, 0, 172, 173, 3, 30, 15, 0,
		173, 175, 1, 0, 0, 0, 174, 161, 1, 0, 0, 0, 174, 167, 1, 0, 0, 0, 175,
		29, 1, 0, 0, 0, 176, 177, 5, 20, 0, 0, 177, 184, 3, 28, 14, 0, 178, 179,
		5, 20, 0, 0, 179, 180, 5, 2, 0, 0, 180, 181, 3, 2, 1, 0, 181, 182, 5, 3,
		0, 0, 182, 184, 1, 0, 0, 0, 183, 176, 1, 0, 0, 0, 183, 178, 1, 0, 0, 0,
		184, 31, 1, 0, 0, 0, 185, 186, 5, 21, 0, 0, 186, 187, 5, 55, 0, 0, 187,
		188, 5, 10, 0, 0, 188, 191, 5, 11, 0, 0, 189, 190, 5, 22, 0, 0, 190, 192,
		3, 26, 13, 0, 191, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 193, 1,
		0, 0, 0, 193, 194, 5, 2, 0, 0, 194, 195, 3, 2, 1, 0, 195, 196, 5, 3, 0,
		0, 196, 211, 1, 0, 0, 0, 197, 198, 5, 21, 0, 0, 198, 199, 5, 55, 0, 0,
		199, 200, 5, 10, 0, 0, 200, 201, 3, 34, 17, 0, 201, 204, 5, 11, 0, 0, 202,
		203, 5, 22, 0, 0, 203, 205, 3, 26, 13, 0, 204, 202, 1, 0, 0, 0, 204, 205,
		1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 207, 5, 2, 0, 0, 207, 208, 3, 2,
		1, 0, 208, 209, 5, 3, 0, 0, 209, 211, 1, 0, 0, 0, 210, 185, 1, 0, 0, 0,
		210, 197, 1, 0, 0, 0, 211, 33, 1, 0, 0, 0, 212, 217, 3, 36, 18, 0, 213,
		214, 5, 13, 0, 0, 214, 216, 3, 36, 18, 0, 215, 213, 1, 0, 0, 0, 216, 219,
		1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 221, 1, 0,
		0, 0, 219, 217, 1, 0, 0, 0, 220, 212, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0,
		221, 35, 1, 0, 0, 0, 222, 223, 7, 3, 0, 0, 223, 224, 5, 55, 0, 0, 224,
		226, 5, 5, 0, 0, 225, 227, 5, 49, 0, 0, 226, 225, 1, 0, 0, 0, 226, 227,
		1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 240, 3, 26, 13, 0, 229, 230, 7,
		3, 0, 0, 230, 231, 5, 55, 0, 0, 231, 233, 5, 5, 0, 0, 232, 234, 5, 49,
		0, 0, 233, 232, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0,
		235, 236, 5, 51, 0, 0, 236, 237, 3, 26, 13, 0, 237, 238, 5, 52, 0, 0, 238,
		240, 1, 0, 0, 0, 239, 222, 1, 0, 0, 0, 239, 229, 1, 0, 0, 0, 240, 37, 1,
		0, 0, 0, 241, 242, 5, 55, 0, 0, 242, 243, 5, 10, 0, 0, 243, 245, 5, 11,
		0, 0, 244, 246, 5, 12, 0, 0, 245, 244, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0,
		246, 255, 1, 0, 0, 0, 247, 248, 5, 55, 0, 0, 248, 249, 5, 10, 0, 0, 249,
		250, 3, 40, 20, 0, 250, 252, 5, 11, 0, 0, 251, 253, 5, 12, 0, 0, 252, 251,
		1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 255, 1, 0, 0, 0, 254, 241, 1, 0,
		0, 0, 254, 247, 1, 0, 0, 0, 255, 39, 1, 0, 0, 0, 256, 257, 3, 42, 21, 0,
		257, 258, 5, 13, 0, 0, 258, 259, 3, 40, 20, 0, 259, 262, 1, 0, 0, 0, 260,
		262, 3, 42, 21, 0, 261, 256, 1, 0, 0, 0, 261, 260, 1, 0, 0, 0, 262, 41,
		1, 0, 0, 0, 263, 264, 5, 50, 0, 0, 264, 270, 3, 44, 22, 0, 265, 266, 5,
		55, 0, 0, 266, 267, 5, 5, 0, 0, 267, 270, 3, 44, 22, 0, 268, 270, 3, 44,
		22, 0, 269, 263, 1, 0, 0, 0, 269, 265, 1, 0, 0, 0, 269, 268, 1, 0, 0, 0,
		270, 43, 1, 0, 0, 0, 271, 272, 6, 22, -1, 0, 272, 273, 7, 4, 0, 0, 273,
		285, 3, 44, 22, 14, 274, 275, 5, 10, 0, 0, 275, 276, 3, 44, 22, 0, 276,
		277, 5, 11, 0, 0, 277, 285, 1, 0, 0, 0, 278, 285, 5, 56, 0, 0, 279, 285,
		5, 57, 0, 0, 280, 285, 5, 40, 0, 0, 281, 285, 5, 58, 0, 0, 282, 285, 5,
		59, 0, 0, 283, 285, 5, 55, 0, 0, 284, 271, 1, 0, 0, 0, 284, 274, 1, 0,
		0, 0, 284, 278, 1, 0, 0, 0, 284, 279, 1, 0, 0, 0, 284, 280, 1, 0, 0, 0,
		284, 281, 1, 0, 0, 0, 284, 282, 1, 0, 0, 0, 284, 283, 1, 0, 0, 0, 285,
		306, 1, 0, 0, 0, 286, 287, 10, 12, 0, 0, 287, 288, 7, 5, 0, 0, 288, 305,
		3, 44, 22, 13, 289, 290, 10, 11, 0, 0, 290, 291, 7, 6, 0, 0, 291, 305,
		3, 44, 22, 12, 292, 293, 10, 10, 0, 0, 293, 294, 7, 7, 0, 0, 294, 305,
		3, 44, 22, 11, 295, 296, 10, 9, 0, 0, 296, 297, 7, 8, 0, 0, 297, 305, 3,
		44, 22, 10, 298, 299, 10, 8, 0, 0, 299, 300, 5, 35, 0, 0, 300, 305, 3,
		44, 22, 9, 301, 302, 10, 7, 0, 0, 302, 303, 5, 36, 0, 0, 303, 305, 3, 44,
		22, 8, 304, 286, 1, 0, 0, 0, 304, 289, 1, 0, 0, 0, 304, 292, 1, 0, 0, 0,
		304, 295, 1, 0, 0, 0, 304, 298, 1, 0, 0, 0, 304, 301, 1, 0, 0, 0, 305,
		308, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 45, 1,
		0, 0, 0, 308, 306, 1, 0, 0, 0, 29, 52, 65, 75, 85, 91, 101, 117, 119, 126,
		145, 157, 174, 183, 191, 204, 210, 217, 220, 226, 233, 239, 245, 252, 254,
		261, 269, 284, 304, 306,
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
	GramarParserRULE_cexpr       = 9
	GramarParserRULE_swhile      = 10
	GramarParserRULE_asign       = 11
	GramarParserRULE_pdeclara    = 12
	GramarParserRULE_tipo        = 13
	GramarParserRULE_pif         = 14
	GramarParserRULE_pelse       = 15
	GramarParserRULE_pfuncion    = 16
	GramarParserRULE_pparams     = 17
	GramarParserRULE_pparamet    = 18
	GramarParserRULE_pllamada    = 19
	GramarParserRULE_argumento   = 20
	GramarParserRULE_tipoArg     = 21
	GramarParserRULE_expr        = 22
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
		p.SetState(46)
		p.Block()
	}
	{
		p.SetState(47)
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
	p.SetState(52)
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
				p.SetState(49)
				p.Productions()
			}

		}
		p.SetState(54)
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
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.Prin()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.Swhile()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.Pif()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(58)
			p.Pdeclara()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(59)
			p.Asign()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(60)
			p.Pfor()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(61)
			p.Pswitch()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(62)
			p.Pfuncion()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(63)
			p.Pllamada()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(64)
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
		p.SetState(67)
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
		p.SetState(69)
		p.Match(GramarParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)
		p.expr(0)
	}
	{
		p.SetState(71)
		p.Match(GramarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GramarParserT__3 {
		{
			p.SetState(72)
			p.Ccase()
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(77)
		p.Pdefault()
	}
	{
		p.SetState(78)
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
		p.SetState(80)
		p.Match(GramarParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.expr(0)
	}
	{
		p.SetState(82)
		p.Match(GramarParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.Block()
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramarParserBRK {
		{
			p.SetState(84)
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
		p.SetState(87)
		p.Match(GramarParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.Match(GramarParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.Block()
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramarParserBRK {
		{
			p.SetState(90)
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
		p.SetState(93)
		p.Match(GramarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Match(GramarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Match(GramarParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(96)
			p.expr(0)
		}
		{
			p.SetState(97)
			p.Match(GramarParserPOINTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.expr(0)
		}

	case 2:
		{
			p.SetState(100)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(103)
		p.Match(GramarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.Block()
	}
	{
		p.SetState(105)
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
	Cexpr() ICexprContext

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

func (s *PrinContext) Cexpr() ICexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICexprContext)
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
	var _la int

	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.Match(GramarParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Match(GramarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.expr(0)
		}
		{
			p.SetState(110)
			p.Match(GramarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Match(GramarParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Match(GramarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Cexpr()
		}
		{
			p.SetState(115)
			p.Match(GramarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__11 {
			{
				p.SetState(116)
				p.Match(GramarParserT__11)
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

// ICexprContext is an interface to support dynamic dispatch.
type ICexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Cexpr() ICexprContext

	// IsCexprContext differentiates from other interfaces.
	IsCexprContext()
}

type CexprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCexprContext() *CexprContext {
	var p = new(CexprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_cexpr
	return p
}

func InitEmptyCexprContext(p *CexprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_cexpr
}

func (*CexprContext) IsCexprContext() {}

func NewCexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CexprContext {
	var p = new(CexprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_cexpr

	return p
}

func (s *CexprContext) GetParser() antlr.Parser { return s.parser }

func (s *CexprContext) Expr() IExprContext {
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

func (s *CexprContext) Cexpr() ICexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICexprContext)
}

func (s *CexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CexprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitCexpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Cexpr() (localctx ICexprContext) {
	localctx = NewCexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GramarParserRULE_cexpr)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.expr(0)
		}
		{
			p.SetState(122)
			p.Match(GramarParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Cexpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
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
	p.EnterRule(localctx, 20, GramarParserRULE_swhile)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(GramarParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.expr(0)
	}
	{
		p.SetState(130)
		p.Match(GramarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Block()
	}
	{
		p.SetState(132)
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
	p.EnterRule(localctx, 22, GramarParserRULE_asign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(GramarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*AsignContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&229376) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*AsignContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(136)
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
	p.EnterRule(localctx, 24, GramarParserRULE_pdeclara)
	var _la int

	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(138)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserRVAR || _la == GramarParserRLET) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(139)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Tipo()
		}
		{
			p.SetState(142)
			p.Match(GramarParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.expr(0)
		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__11 {
			{
				p.SetState(144)
				p.Match(GramarParserT__11)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(147)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserRVAR || _la == GramarParserRLET) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(148)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Tipo()
		}
		{
			p.SetState(151)
			p.Match(GramarParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(153)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserRVAR || _la == GramarParserRLET) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(154)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.Match(GramarParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
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
	p.EnterRule(localctx, 26, GramarParserRULE_tipo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
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
	p.EnterRule(localctx, 28, GramarParserRULE_pif)
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.Match(GramarParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.expr(0)
		}
		{
			p.SetState(163)
			p.Match(GramarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.Block()
		}
		{
			p.SetState(165)
			p.Match(GramarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Match(GramarParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.expr(0)
		}
		{
			p.SetState(169)
			p.Match(GramarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.Block()
		}
		{
			p.SetState(171)
			p.Match(GramarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
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
	p.EnterRule(localctx, 30, GramarParserRULE_pelse)
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)
			p.Match(GramarParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Pif()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(178)
			p.Match(GramarParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.Match(GramarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.Block()
		}
		{
			p.SetState(181)
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
	p.EnterRule(localctx, 32, GramarParserRULE_pfuncion)
	var _la int

	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.Match(GramarParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(GramarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Match(GramarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__21 {
			{
				p.SetState(189)
				p.Match(GramarParserT__21)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(190)
				p.Tipo()
			}

		}
		{
			p.SetState(193)
			p.Match(GramarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.Block()
		}
		{
			p.SetState(195)
			p.Match(GramarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(197)
			p.Match(GramarParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.Match(GramarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.Pparams()
		}
		{
			p.SetState(201)
			p.Match(GramarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__21 {
			{
				p.SetState(202)
				p.Match(GramarParserT__21)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(203)
				p.Tipo()
			}

		}
		{
			p.SetState(206)
			p.Match(GramarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.Block()
		}
		{
			p.SetState(208)
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
	p.EnterRule(localctx, 34, GramarParserRULE_pparams)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramarParserGUION || _la == GramarParserID {
		{
			p.SetState(212)
			p.Pparamet()
		}
		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GramarParserT__12 {
			{
				p.SetState(213)
				p.Match(GramarParserT__12)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(214)
				p.Pparamet()
			}

			p.SetState(219)
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
	p.EnterRule(localctx, 36, GramarParserRULE_pparamet)
	var _la int

	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(222)

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
			p.SetState(223)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserINOUT {
			{
				p.SetState(225)
				p.Match(GramarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(228)
			p.Tipo()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)

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
			p.SetState(230)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Match(GramarParserT__4)
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

		if _la == GramarParserINOUT {
			{
				p.SetState(232)
				p.Match(GramarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(235)
			p.Match(GramarParserABR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.Tipo()
		}
		{
			p.SetState(237)
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
	p.EnterRule(localctx, 38, GramarParserRULE_pllamada)
	var _la int

	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Match(GramarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.Match(GramarParserT__10)
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

		if _la == GramarParserT__11 {
			{
				p.SetState(244)
				p.Match(GramarParserT__11)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

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
			p.Match(GramarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)
			p.Argumento()
		}
		{
			p.SetState(250)
			p.Match(GramarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__11 {
			{
				p.SetState(251)
				p.Match(GramarParserT__11)
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
	p.EnterRule(localctx, 40, GramarParserRULE_argumento)
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(256)
			p.TipoArg()
		}
		{
			p.SetState(257)
			p.Match(GramarParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)
			p.Argumento()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(260)
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
	p.EnterRule(localctx, 42, GramarParserRULE_tipoArg)
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.Match(GramarParserPUNTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)

			var _x = p.expr(0)

			localctx.(*TipoArgContext).exp1 = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(265)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(266)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)

			var _x = p.expr(0)

			localctx.(*TipoArgContext).exp2 = _x
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(268)

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
	_startState := 44
	p.EnterRecursionRule(localctx, 44, GramarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(284)
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
			p.SetState(272)

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
			p.SetState(273)

			var _x = p.expr(14)

			localctx.(*AritmeticContext).right = _x
		}

	case GramarParserT__9:
		localctx = NewParentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(274)
			p.Match(GramarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)
			p.expr(0)
		}
		{
			p.SetState(276)
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
			p.SetState(278)
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
			p.SetState(279)
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
			p.SetState(280)
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
			p.SetState(281)
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
			p.SetState(282)
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
			p.SetState(283)
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
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(304)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAritmeticContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AritmeticContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(286)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(287)

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
					p.SetState(288)

					var _x = p.expr(13)

					localctx.(*AritmeticContext).right = _x
				}

			case 2:
				localctx = NewAritmeticContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AritmeticContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(289)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(290)

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
					p.SetState(291)

					var _x = p.expr(12)

					localctx.(*AritmeticContext).right = _x
				}

			case 3:
				localctx = NewAritmeticContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AritmeticContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(293)

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
					p.SetState(294)

					var _x = p.expr(11)

					localctx.(*AritmeticContext).right = _x
				}

			case 4:
				localctx = NewAritmeticContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AritmeticContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(295)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(296)

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
					p.SetState(297)

					var _x = p.expr(10)

					localctx.(*AritmeticContext).right = _x
				}

			case 5:
				localctx = NewAritmeticContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AritmeticContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(298)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(299)

					var _m = p.Match(GramarParserT__34)

					localctx.(*AritmeticContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(300)

					var _x = p.expr(9)

					localctx.(*AritmeticContext).right = _x
				}

			case 6:
				localctx = NewAritmeticContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AritmeticContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(301)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(302)

					var _m = p.Match(GramarParserT__35)

					localctx.(*AritmeticContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(303)

					var _x = p.expr(8)

					localctx.(*AritmeticContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
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
	case 22:
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
