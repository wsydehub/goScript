// Code generated from GoScript.g4 by ANTLR 4.13.1. DO NOT EDIT.

package goScript // GoScript
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

type GoScriptParser struct {
	*antlr.BaseParser
}

var GoScriptParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func goscriptParserInit() {
	staticData := &GoScriptParserStaticData
	staticData.LiteralNames = []string{
		"", "';'", "'func'", "'('", "','", "')'", "'{'", "'}'", "':'", "'['",
		"']'", "'map'", "'connector'", "'dynamic'", "'error'", "'int'", "'uint'",
		"'float'", "'bool'", "'char'", "'string'", "'if'", "'else'", "'for'",
		"'return'", "'break'", "'continue'", "'.'", "'++'", "'--'", "'new'",
		"'&&'", "'||'", "'?'", "':='", "", "", "'null'", "", "", "", "", "'+'",
		"'-'", "'*'", "'/'", "'%'", "'>'", "'>='", "'<'", "'<='", "'='", "'=='",
		"'!='", "'!'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "BooleanLiteral", "Identifier", "NULL", "HexLiteral", "DecimalLiteral",
		"OctalLiteral", "FloatingPointLiteral", "Plus", "Minus", "Multiplication",
		"Division", "Percent", "Gt", "Get", "Lt", "Let", "Assign", "Eq", "NEq",
		"Not", "CharacterLiteral", "StringLiteral", "COMMENT", "WS", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"compilationUnit", "functionDeclaration", "formalParameters", "formalParameterDecl",
		"returnType", "block", "blockStatement", "variableDeclaration", "variableDeclarators",
		"variableDeclarator", "variableInitializer", "arrayInitializer", "mapInitializer",
		"mapEntry", "type_", "mapType", "connectorType", "dynamicType", "primitiveType",
		"statement", "ifStatement", "forStatement", "forControl", "forInit",
		"forUpdate", "returnStatement", "breakStatement", "continueStatement",
		"expressionStatement", "expression", "primary", "literal", "integerLiteral",
		"expressionList", "identifierList", "lvalue", "creator", "mapCreator",
		"arrayCreator", "creatorName", "connectorCreator", "primitiveCreator",
		"dynamicCreator",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 59, 477, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 91, 8, 0, 10, 0, 12, 0, 94, 9,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 102, 8, 1, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 5, 2, 110, 8, 2, 10, 2, 12, 2, 113, 9, 2, 3, 2, 115,
		8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4,
		127, 8, 4, 10, 4, 12, 4, 130, 9, 4, 1, 4, 1, 4, 3, 4, 134, 8, 4, 1, 5,
		1, 5, 5, 5, 138, 8, 5, 10, 5, 12, 5, 141, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 149, 8, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 5, 8,
		157, 8, 8, 10, 8, 12, 8, 160, 9, 8, 1, 9, 1, 9, 1, 9, 3, 9, 165, 8, 9,
		1, 10, 1, 10, 1, 10, 3, 10, 170, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 5,
		11, 176, 8, 11, 10, 11, 12, 11, 179, 9, 11, 3, 11, 181, 8, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 189, 8, 12, 10, 12, 12, 12, 192,
		9, 12, 1, 12, 3, 12, 195, 8, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 5, 14, 206, 8, 14, 10, 14, 12, 14, 209, 9, 14,
		1, 14, 1, 14, 1, 14, 5, 14, 214, 8, 14, 10, 14, 12, 14, 217, 9, 14, 1,
		14, 1, 14, 1, 14, 5, 14, 222, 8, 14, 10, 14, 12, 14, 225, 9, 14, 1, 14,
		1, 14, 1, 14, 5, 14, 230, 8, 14, 10, 14, 12, 14, 233, 9, 14, 3, 14, 235,
		8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 3, 19, 260, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 3, 20, 267, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22,
		3, 22, 276, 8, 22, 1, 22, 1, 22, 3, 22, 280, 8, 22, 1, 22, 1, 22, 3, 22,
		284, 8, 22, 1, 23, 1, 23, 3, 23, 288, 8, 23, 1, 24, 1, 24, 1, 25, 1, 25,
		3, 25, 294, 8, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1,
		27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 5, 29, 316, 8, 29, 10, 29, 12, 29, 319, 9, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 5, 29, 325, 8, 29, 10, 29, 12, 29, 328, 9, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 335, 8, 29, 10, 29, 12, 29, 338, 9,
		29, 3, 29, 340, 8, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 374, 8, 29, 1, 29, 1, 29, 1, 29, 5,
		29, 379, 8, 29, 10, 29, 12, 29, 382, 9, 29, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 3, 30, 390, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 3, 31, 398, 8, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 5, 33, 405, 8,
		33, 10, 33, 12, 33, 408, 9, 33, 1, 34, 1, 34, 1, 34, 5, 34, 413, 8, 34,
		10, 34, 12, 34, 416, 9, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 5, 35, 429, 8, 35, 10, 35, 12, 35, 432,
		9, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 439, 8, 36, 1, 37, 1,
		37, 1, 37, 1, 38, 1, 38, 1, 38, 4, 38, 447, 8, 38, 11, 38, 12, 38, 448,
		1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 457, 8, 39, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 3, 41, 466, 8, 41, 1, 41, 1, 41,
		1, 42, 1, 42, 1, 42, 3, 42, 473, 8, 42, 1, 42, 1, 42, 1, 42, 0, 2, 58,
		70, 43, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
		70, 72, 74, 76, 78, 80, 82, 84, 0, 8, 1, 0, 13, 14, 1, 0, 15, 20, 2, 0,
		42, 43, 54, 54, 1, 0, 44, 46, 1, 0, 42, 43, 2, 0, 47, 50, 52, 53, 1, 0,
		28, 29, 1, 0, 38, 40, 508, 0, 92, 1, 0, 0, 0, 2, 97, 1, 0, 0, 0, 4, 105,
		1, 0, 0, 0, 6, 118, 1, 0, 0, 0, 8, 133, 1, 0, 0, 0, 10, 135, 1, 0, 0, 0,
		12, 148, 1, 0, 0, 0, 14, 150, 1, 0, 0, 0, 16, 153, 1, 0, 0, 0, 18, 161,
		1, 0, 0, 0, 20, 169, 1, 0, 0, 0, 22, 171, 1, 0, 0, 0, 24, 184, 1, 0, 0,
		0, 26, 198, 1, 0, 0, 0, 28, 234, 1, 0, 0, 0, 30, 236, 1, 0, 0, 0, 32, 243,
		1, 0, 0, 0, 34, 248, 1, 0, 0, 0, 36, 250, 1, 0, 0, 0, 38, 259, 1, 0, 0,
		0, 40, 261, 1, 0, 0, 0, 42, 268, 1, 0, 0, 0, 44, 275, 1, 0, 0, 0, 46, 287,
		1, 0, 0, 0, 48, 289, 1, 0, 0, 0, 50, 291, 1, 0, 0, 0, 52, 297, 1, 0, 0,
		0, 54, 300, 1, 0, 0, 0, 56, 303, 1, 0, 0, 0, 58, 339, 1, 0, 0, 0, 60, 389,
		1, 0, 0, 0, 62, 397, 1, 0, 0, 0, 64, 399, 1, 0, 0, 0, 66, 401, 1, 0, 0,
		0, 68, 409, 1, 0, 0, 0, 70, 417, 1, 0, 0, 0, 72, 438, 1, 0, 0, 0, 74, 440,
		1, 0, 0, 0, 76, 443, 1, 0, 0, 0, 78, 456, 1, 0, 0, 0, 80, 458, 1, 0, 0,
		0, 82, 462, 1, 0, 0, 0, 84, 469, 1, 0, 0, 0, 86, 91, 3, 2, 1, 0, 87, 88,
		3, 14, 7, 0, 88, 89, 5, 1, 0, 0, 89, 91, 1, 0, 0, 0, 90, 86, 1, 0, 0, 0,
		90, 87, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1,
		0, 0, 0, 93, 95, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 96, 5, 0, 0, 1, 96,
		1, 1, 0, 0, 0, 97, 98, 5, 2, 0, 0, 98, 99, 5, 36, 0, 0, 99, 101, 3, 4,
		2, 0, 100, 102, 3, 8, 4, 0, 101, 100, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0,
		102, 103, 1, 0, 0, 0, 103, 104, 3, 10, 5, 0, 104, 3, 1, 0, 0, 0, 105, 114,
		5, 3, 0, 0, 106, 111, 3, 6, 3, 0, 107, 108, 5, 4, 0, 0, 108, 110, 3, 6,
		3, 0, 109, 107, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0,
		111, 112, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114,
		106, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 117,
		5, 5, 0, 0, 117, 5, 1, 0, 0, 0, 118, 119, 3, 28, 14, 0, 119, 120, 5, 36,
		0, 0, 120, 7, 1, 0, 0, 0, 121, 134, 3, 28, 14, 0, 122, 123, 5, 3, 0, 0,
		123, 128, 3, 28, 14, 0, 124, 125, 5, 4, 0, 0, 125, 127, 3, 28, 14, 0, 126,
		124, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 128, 129,
		1, 0, 0, 0, 129, 131, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131, 132, 5, 5,
		0, 0, 132, 134, 1, 0, 0, 0, 133, 121, 1, 0, 0, 0, 133, 122, 1, 0, 0, 0,
		134, 9, 1, 0, 0, 0, 135, 139, 5, 6, 0, 0, 136, 138, 3, 12, 6, 0, 137, 136,
		1, 0, 0, 0, 138, 141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0,
		0, 0, 140, 142, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 143, 5, 7, 0, 0,
		143, 11, 1, 0, 0, 0, 144, 145, 3, 14, 7, 0, 145, 146, 5, 1, 0, 0, 146,
		149, 1, 0, 0, 0, 147, 149, 3, 38, 19, 0, 148, 144, 1, 0, 0, 0, 148, 147,
		1, 0, 0, 0, 149, 13, 1, 0, 0, 0, 150, 151, 3, 28, 14, 0, 151, 152, 3, 16,
		8, 0, 152, 15, 1, 0, 0, 0, 153, 158, 3, 18, 9, 0, 154, 155, 5, 4, 0, 0,
		155, 157, 3, 18, 9, 0, 156, 154, 1, 0, 0, 0, 157, 160, 1, 0, 0, 0, 158,
		156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 17, 1, 0, 0, 0, 160, 158, 1,
		0, 0, 0, 161, 164, 5, 36, 0, 0, 162, 163, 5, 51, 0, 0, 163, 165, 3, 20,
		10, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 19, 1, 0, 0, 0,
		166, 170, 3, 22, 11, 0, 167, 170, 3, 24, 12, 0, 168, 170, 3, 58, 29, 0,
		169, 166, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 168, 1, 0, 0, 0, 170,
		21, 1, 0, 0, 0, 171, 180, 5, 6, 0, 0, 172, 177, 3, 20, 10, 0, 173, 174,
		5, 4, 0, 0, 174, 176, 3, 20, 10, 0, 175, 173, 1, 0, 0, 0, 176, 179, 1,
		0, 0, 0, 177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 181, 1, 0, 0,
		0, 179, 177, 1, 0, 0, 0, 180, 172, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181,
		182, 1, 0, 0, 0, 182, 183, 5, 7, 0, 0, 183, 23, 1, 0, 0, 0, 184, 185, 5,
		6, 0, 0, 185, 190, 3, 26, 13, 0, 186, 187, 5, 4, 0, 0, 187, 189, 3, 26,
		13, 0, 188, 186, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0,
		190, 191, 1, 0, 0, 0, 191, 194, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 193,
		195, 5, 4, 0, 0, 194, 193, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 196,
		1, 0, 0, 0, 196, 197, 5, 7, 0, 0, 197, 25, 1, 0, 0, 0, 198, 199, 3, 58,
		29, 0, 199, 200, 5, 8, 0, 0, 200, 201, 3, 20, 10, 0, 201, 27, 1, 0, 0,
		0, 202, 207, 3, 36, 18, 0, 203, 204, 5, 9, 0, 0, 204, 206, 5, 10, 0, 0,
		205, 203, 1, 0, 0, 0, 206, 209, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207,
		208, 1, 0, 0, 0, 208, 235, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 210, 215,
		3, 30, 15, 0, 211, 212, 5, 9, 0, 0, 212, 214, 5, 10, 0, 0, 213, 211, 1,
		0, 0, 0, 214, 217, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 215, 216, 1, 0, 0,
		0, 216, 235, 1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 218, 223, 3, 32, 16, 0,
		219, 220, 5, 9, 0, 0, 220, 222, 5, 10, 0, 0, 221, 219, 1, 0, 0, 0, 222,
		225, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 235,
		1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 226, 231, 3, 34, 17, 0, 227, 228, 5,
		9, 0, 0, 228, 230, 5, 10, 0, 0, 229, 227, 1, 0, 0, 0, 230, 233, 1, 0, 0,
		0, 231, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 235, 1, 0, 0, 0, 233,
		231, 1, 0, 0, 0, 234, 202, 1, 0, 0, 0, 234, 210, 1, 0, 0, 0, 234, 218,
		1, 0, 0, 0, 234, 226, 1, 0, 0, 0, 235, 29, 1, 0, 0, 0, 236, 237, 5, 11,
		0, 0, 237, 238, 5, 49, 0, 0, 238, 239, 3, 36, 18, 0, 239, 240, 5, 4, 0,
		0, 240, 241, 3, 28, 14, 0, 241, 242, 5, 47, 0, 0, 242, 31, 1, 0, 0, 0,
		243, 244, 5, 12, 0, 0, 244, 245, 5, 49, 0, 0, 245, 246, 5, 36, 0, 0, 246,
		247, 5, 47, 0, 0, 247, 33, 1, 0, 0, 0, 248, 249, 7, 0, 0, 0, 249, 35, 1,
		0, 0, 0, 250, 251, 7, 1, 0, 0, 251, 37, 1, 0, 0, 0, 252, 260, 3, 10, 5,
		0, 253, 260, 3, 40, 20, 0, 254, 260, 3, 42, 21, 0, 255, 260, 3, 50, 25,
		0, 256, 260, 3, 52, 26, 0, 257, 260, 3, 54, 27, 0, 258, 260, 3, 56, 28,
		0, 259, 252, 1, 0, 0, 0, 259, 253, 1, 0, 0, 0, 259, 254, 1, 0, 0, 0, 259,
		255, 1, 0, 0, 0, 259, 256, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259, 258,
		1, 0, 0, 0, 260, 39, 1, 0, 0, 0, 261, 262, 5, 21, 0, 0, 262, 263, 3, 58,
		29, 0, 263, 266, 3, 38, 19, 0, 264, 265, 5, 22, 0, 0, 265, 267, 3, 38,
		19, 0, 266, 264, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 41, 1, 0, 0, 0,
		268, 269, 5, 23, 0, 0, 269, 270, 5, 3, 0, 0, 270, 271, 3, 44, 22, 0, 271,
		272, 5, 5, 0, 0, 272, 273, 3, 38, 19, 0, 273, 43, 1, 0, 0, 0, 274, 276,
		3, 46, 23, 0, 275, 274, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 277, 1,
		0, 0, 0, 277, 279, 5, 1, 0, 0, 278, 280, 3, 58, 29, 0, 279, 278, 1, 0,
		0, 0, 279, 280, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 283, 5, 1, 0, 0,
		282, 284, 3, 48, 24, 0, 283, 282, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284,
		45, 1, 0, 0, 0, 285, 288, 3, 14, 7, 0, 286, 288, 3, 66, 33, 0, 287, 285,
		1, 0, 0, 0, 287, 286, 1, 0, 0, 0, 288, 47, 1, 0, 0, 0, 289, 290, 3, 66,
		33, 0, 290, 49, 1, 0, 0, 0, 291, 293, 5, 24, 0, 0, 292, 294, 3, 58, 29,
		0, 293, 292, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295,
		296, 5, 1, 0, 0, 296, 51, 1, 0, 0, 0, 297, 298, 5, 25, 0, 0, 298, 299,
		5, 1, 0, 0, 299, 53, 1, 0, 0, 0, 300, 301, 5, 26, 0, 0, 301, 302, 5, 1,
		0, 0, 302, 55, 1, 0, 0, 0, 303, 304, 3, 58, 29, 0, 304, 305, 5, 1, 0, 0,
		305, 57, 1, 0, 0, 0, 306, 307, 6, 29, -1, 0, 307, 340, 3, 60, 30, 0, 308,
		309, 7, 2, 0, 0, 309, 340, 3, 58, 29, 10, 310, 311, 5, 30, 0, 0, 311, 340,
		3, 72, 36, 0, 312, 317, 3, 70, 35, 0, 313, 314, 5, 4, 0, 0, 314, 316, 3,
		70, 35, 0, 315, 313, 1, 0, 0, 0, 316, 319, 1, 0, 0, 0, 317, 315, 1, 0,
		0, 0, 317, 318, 1, 0, 0, 0, 318, 320, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0,
		320, 321, 5, 51, 0, 0, 321, 326, 3, 58, 29, 0, 322, 323, 5, 4, 0, 0, 323,
		325, 3, 58, 29, 0, 324, 322, 1, 0, 0, 0, 325, 328, 1, 0, 0, 0, 326, 324,
		1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 340, 1, 0, 0, 0, 328, 326, 1, 0,
		0, 0, 329, 330, 3, 68, 34, 0, 330, 331, 5, 34, 0, 0, 331, 336, 3, 58, 29,
		0, 332, 333, 5, 4, 0, 0, 333, 335, 3, 58, 29, 0, 334, 332, 1, 0, 0, 0,
		335, 338, 1, 0, 0, 0, 336, 334, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337,
		340, 1, 0, 0, 0, 338, 336, 1, 0, 0, 0, 339, 306, 1, 0, 0, 0, 339, 308,
		1, 0, 0, 0, 339, 310, 1, 0, 0, 0, 339, 312, 1, 0, 0, 0, 339, 329, 1, 0,
		0, 0, 340, 380, 1, 0, 0, 0, 341, 342, 10, 8, 0, 0, 342, 343, 7, 3, 0, 0,
		343, 379, 3, 58, 29, 9, 344, 345, 10, 7, 0, 0, 345, 346, 7, 4, 0, 0, 346,
		379, 3, 58, 29, 8, 347, 348, 10, 6, 0, 0, 348, 349, 7, 5, 0, 0, 349, 379,
		3, 58, 29, 7, 350, 351, 10, 5, 0, 0, 351, 352, 5, 31, 0, 0, 352, 379, 3,
		58, 29, 6, 353, 354, 10, 4, 0, 0, 354, 355, 5, 32, 0, 0, 355, 379, 3, 58,
		29, 5, 356, 357, 10, 3, 0, 0, 357, 358, 5, 33, 0, 0, 358, 359, 3, 58, 29,
		0, 359, 360, 5, 8, 0, 0, 360, 361, 3, 58, 29, 4, 361, 379, 1, 0, 0, 0,
		362, 363, 10, 14, 0, 0, 363, 364, 5, 27, 0, 0, 364, 379, 5, 36, 0, 0, 365,
		366, 10, 13, 0, 0, 366, 367, 5, 9, 0, 0, 367, 368, 3, 58, 29, 0, 368, 369,
		5, 10, 0, 0, 369, 379, 1, 0, 0, 0, 370, 371, 10, 12, 0, 0, 371, 373, 5,
		3, 0, 0, 372, 374, 3, 66, 33, 0, 373, 372, 1, 0, 0, 0, 373, 374, 1, 0,
		0, 0, 374, 375, 1, 0, 0, 0, 375, 379, 5, 5, 0, 0, 376, 377, 10, 11, 0,
		0, 377, 379, 7, 6, 0, 0, 378, 341, 1, 0, 0, 0, 378, 344, 1, 0, 0, 0, 378,
		347, 1, 0, 0, 0, 378, 350, 1, 0, 0, 0, 378, 353, 1, 0, 0, 0, 378, 356,
		1, 0, 0, 0, 378, 362, 1, 0, 0, 0, 378, 365, 1, 0, 0, 0, 378, 370, 1, 0,
		0, 0, 378, 376, 1, 0, 0, 0, 379, 382, 1, 0, 0, 0, 380, 378, 1, 0, 0, 0,
		380, 381, 1, 0, 0, 0, 381, 59, 1, 0, 0, 0, 382, 380, 1, 0, 0, 0, 383, 384,
		5, 3, 0, 0, 384, 385, 3, 58, 29, 0, 385, 386, 5, 5, 0, 0, 386, 390, 1,
		0, 0, 0, 387, 390, 3, 62, 31, 0, 388, 390, 5, 36, 0, 0, 389, 383, 1, 0,
		0, 0, 389, 387, 1, 0, 0, 0, 389, 388, 1, 0, 0, 0, 390, 61, 1, 0, 0, 0,
		391, 398, 3, 64, 32, 0, 392, 398, 5, 41, 0, 0, 393, 398, 5, 55, 0, 0, 394,
		398, 5, 56, 0, 0, 395, 398, 5, 35, 0, 0, 396, 398, 5, 37, 0, 0, 397, 391,
		1, 0, 0, 0, 397, 392, 1, 0, 0, 0, 397, 393, 1, 0, 0, 0, 397, 394, 1, 0,
		0, 0, 397, 395, 1, 0, 0, 0, 397, 396, 1, 0, 0, 0, 398, 63, 1, 0, 0, 0,
		399, 400, 7, 7, 0, 0, 400, 65, 1, 0, 0, 0, 401, 406, 3, 58, 29, 0, 402,
		403, 5, 4, 0, 0, 403, 405, 3, 58, 29, 0, 404, 402, 1, 0, 0, 0, 405, 408,
		1, 0, 0, 0, 406, 404, 1, 0, 0, 0, 406, 407, 1, 0, 0, 0, 407, 67, 1, 0,
		0, 0, 408, 406, 1, 0, 0, 0, 409, 414, 5, 36, 0, 0, 410, 411, 5, 4, 0, 0,
		411, 413, 5, 36, 0, 0, 412, 410, 1, 0, 0, 0, 413, 416, 1, 0, 0, 0, 414,
		412, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415, 69, 1, 0, 0, 0, 416, 414, 1,
		0, 0, 0, 417, 418, 6, 35, -1, 0, 418, 419, 5, 36, 0, 0, 419, 430, 1, 0,
		0, 0, 420, 421, 10, 2, 0, 0, 421, 422, 5, 27, 0, 0, 422, 429, 5, 36, 0,
		0, 423, 424, 10, 1, 0, 0, 424, 425, 5, 9, 0, 0, 425, 426, 3, 58, 29, 0,
		426, 427, 5, 10, 0, 0, 427, 429, 1, 0, 0, 0, 428, 420, 1, 0, 0, 0, 428,
		423, 1, 0, 0, 0, 429, 432, 1, 0, 0, 0, 430, 428, 1, 0, 0, 0, 430, 431,
		1, 0, 0, 0, 431, 71, 1, 0, 0, 0, 432, 430, 1, 0, 0, 0, 433, 439, 3, 74,
		37, 0, 434, 439, 3, 76, 38, 0, 435, 439, 3, 80, 40, 0, 436, 439, 3, 82,
		41, 0, 437, 439, 3, 84, 42, 0, 438, 433, 1, 0, 0, 0, 438, 434, 1, 0, 0,
		0, 438, 435, 1, 0, 0, 0, 438, 436, 1, 0, 0, 0, 438, 437, 1, 0, 0, 0, 439,
		73, 1, 0, 0, 0, 440, 441, 3, 30, 15, 0, 441, 442, 3, 24, 12, 0, 442, 75,
		1, 0, 0, 0, 443, 446, 3, 78, 39, 0, 444, 445, 5, 9, 0, 0, 445, 447, 5,
		10, 0, 0, 446, 444, 1, 0, 0, 0, 447, 448, 1, 0, 0, 0, 448, 446, 1, 0, 0,
		0, 448, 449, 1, 0, 0, 0, 449, 450, 1, 0, 0, 0, 450, 451, 3, 22, 11, 0,
		451, 77, 1, 0, 0, 0, 452, 457, 3, 36, 18, 0, 453, 457, 3, 30, 15, 0, 454,
		457, 3, 32, 16, 0, 455, 457, 3, 34, 17, 0, 456, 452, 1, 0, 0, 0, 456, 453,
		1, 0, 0, 0, 456, 454, 1, 0, 0, 0, 456, 455, 1, 0, 0, 0, 457, 79, 1, 0,
		0, 0, 458, 459, 3, 32, 16, 0, 459, 460, 5, 3, 0, 0, 460, 461, 5, 5, 0,
		0, 461, 81, 1, 0, 0, 0, 462, 463, 3, 36, 18, 0, 463, 465, 5, 3, 0, 0, 464,
		466, 3, 58, 29, 0, 465, 464, 1, 0, 0, 0, 465, 466, 1, 0, 0, 0, 466, 467,
		1, 0, 0, 0, 467, 468, 5, 5, 0, 0, 468, 83, 1, 0, 0, 0, 469, 470, 3, 34,
		17, 0, 470, 472, 5, 3, 0, 0, 471, 473, 3, 58, 29, 0, 472, 471, 1, 0, 0,
		0, 472, 473, 1, 0, 0, 0, 473, 474, 1, 0, 0, 0, 474, 475, 5, 5, 0, 0, 475,
		85, 1, 0, 0, 0, 46, 90, 92, 101, 111, 114, 128, 133, 139, 148, 158, 164,
		169, 177, 180, 190, 194, 207, 215, 223, 231, 234, 259, 266, 275, 279, 283,
		287, 293, 317, 326, 336, 339, 373, 378, 380, 389, 397, 406, 414, 428, 430,
		438, 448, 456, 465, 472,
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

// GoScriptParserInit initializes any static state used to implement GoScriptParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGoScriptParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoScriptParserInit() {
	staticData := &GoScriptParserStaticData
	staticData.once.Do(goscriptParserInit)
}

// NewGoScriptParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGoScriptParser(input antlr.TokenStream) *GoScriptParser {
	GoScriptParserInit()
	this := new(GoScriptParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GoScriptParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "GoScript.g4"

	return this
}

// GoScriptParser tokens.
const (
	GoScriptParserEOF                  = antlr.TokenEOF
	GoScriptParserT__0                 = 1
	GoScriptParserT__1                 = 2
	GoScriptParserT__2                 = 3
	GoScriptParserT__3                 = 4
	GoScriptParserT__4                 = 5
	GoScriptParserT__5                 = 6
	GoScriptParserT__6                 = 7
	GoScriptParserT__7                 = 8
	GoScriptParserT__8                 = 9
	GoScriptParserT__9                 = 10
	GoScriptParserT__10                = 11
	GoScriptParserT__11                = 12
	GoScriptParserT__12                = 13
	GoScriptParserT__13                = 14
	GoScriptParserT__14                = 15
	GoScriptParserT__15                = 16
	GoScriptParserT__16                = 17
	GoScriptParserT__17                = 18
	GoScriptParserT__18                = 19
	GoScriptParserT__19                = 20
	GoScriptParserT__20                = 21
	GoScriptParserT__21                = 22
	GoScriptParserT__22                = 23
	GoScriptParserT__23                = 24
	GoScriptParserT__24                = 25
	GoScriptParserT__25                = 26
	GoScriptParserT__26                = 27
	GoScriptParserT__27                = 28
	GoScriptParserT__28                = 29
	GoScriptParserT__29                = 30
	GoScriptParserT__30                = 31
	GoScriptParserT__31                = 32
	GoScriptParserT__32                = 33
	GoScriptParserT__33                = 34
	GoScriptParserBooleanLiteral       = 35
	GoScriptParserIdentifier           = 36
	GoScriptParserNULL                 = 37
	GoScriptParserHexLiteral           = 38
	GoScriptParserDecimalLiteral       = 39
	GoScriptParserOctalLiteral         = 40
	GoScriptParserFloatingPointLiteral = 41
	GoScriptParserPlus                 = 42
	GoScriptParserMinus                = 43
	GoScriptParserMultiplication       = 44
	GoScriptParserDivision             = 45
	GoScriptParserPercent              = 46
	GoScriptParserGt                   = 47
	GoScriptParserGet                  = 48
	GoScriptParserLt                   = 49
	GoScriptParserLet                  = 50
	GoScriptParserAssign               = 51
	GoScriptParserEq                   = 52
	GoScriptParserNEq                  = 53
	GoScriptParserNot                  = 54
	GoScriptParserCharacterLiteral     = 55
	GoScriptParserStringLiteral        = 56
	GoScriptParserCOMMENT              = 57
	GoScriptParserWS                   = 58
	GoScriptParserLINE_COMMENT         = 59
)

// GoScriptParser rules.
const (
	GoScriptParserRULE_compilationUnit     = 0
	GoScriptParserRULE_functionDeclaration = 1
	GoScriptParserRULE_formalParameters    = 2
	GoScriptParserRULE_formalParameterDecl = 3
	GoScriptParserRULE_returnType          = 4
	GoScriptParserRULE_block               = 5
	GoScriptParserRULE_blockStatement      = 6
	GoScriptParserRULE_variableDeclaration = 7
	GoScriptParserRULE_variableDeclarators = 8
	GoScriptParserRULE_variableDeclarator  = 9
	GoScriptParserRULE_variableInitializer = 10
	GoScriptParserRULE_arrayInitializer    = 11
	GoScriptParserRULE_mapInitializer      = 12
	GoScriptParserRULE_mapEntry            = 13
	GoScriptParserRULE_type_               = 14
	GoScriptParserRULE_mapType             = 15
	GoScriptParserRULE_connectorType       = 16
	GoScriptParserRULE_dynamicType         = 17
	GoScriptParserRULE_primitiveType       = 18
	GoScriptParserRULE_statement           = 19
	GoScriptParserRULE_ifStatement         = 20
	GoScriptParserRULE_forStatement        = 21
	GoScriptParserRULE_forControl          = 22
	GoScriptParserRULE_forInit             = 23
	GoScriptParserRULE_forUpdate           = 24
	GoScriptParserRULE_returnStatement     = 25
	GoScriptParserRULE_breakStatement      = 26
	GoScriptParserRULE_continueStatement   = 27
	GoScriptParserRULE_expressionStatement = 28
	GoScriptParserRULE_expression          = 29
	GoScriptParserRULE_primary             = 30
	GoScriptParserRULE_literal             = 31
	GoScriptParserRULE_integerLiteral      = 32
	GoScriptParserRULE_expressionList      = 33
	GoScriptParserRULE_identifierList      = 34
	GoScriptParserRULE_lvalue              = 35
	GoScriptParserRULE_creator             = 36
	GoScriptParserRULE_mapCreator          = 37
	GoScriptParserRULE_arrayCreator        = 38
	GoScriptParserRULE_creatorName         = 39
	GoScriptParserRULE_connectorCreator    = 40
	GoScriptParserRULE_primitiveCreator    = 41
	GoScriptParserRULE_dynamicCreator      = 42
)

// ICompilationUnitContext is an interface to support dynamic dispatch.
type ICompilationUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllFunctionDeclaration() []IFunctionDeclarationContext
	FunctionDeclaration(i int) IFunctionDeclarationContext
	AllVariableDeclaration() []IVariableDeclarationContext
	VariableDeclaration(i int) IVariableDeclarationContext

	// IsCompilationUnitContext differentiates from other interfaces.
	IsCompilationUnitContext()
}

type CompilationUnitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompilationUnitContext() *CompilationUnitContext {
	var p = new(CompilationUnitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_compilationUnit
	return p
}

func InitEmptyCompilationUnitContext(p *CompilationUnitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_compilationUnit
}

func (*CompilationUnitContext) IsCompilationUnitContext() {}

func NewCompilationUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompilationUnitContext {
	var p = new(CompilationUnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_compilationUnit

	return p
}

func (s *CompilationUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *CompilationUnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(GoScriptParserEOF, 0)
}

func (s *CompilationUnitContext) AllFunctionDeclaration() []IFunctionDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IFunctionDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionDeclarationContext); ok {
			tst[i] = t.(IFunctionDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *CompilationUnitContext) FunctionDeclaration(i int) IFunctionDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
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

	return t.(IFunctionDeclarationContext)
}

func (s *CompilationUnitContext) AllVariableDeclaration() []IVariableDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IVariableDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDeclarationContext); ok {
			tst[i] = t.(IVariableDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *CompilationUnitContext) VariableDeclaration(i int) IVariableDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
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

	return t.(IVariableDeclarationContext)
}

func (s *CompilationUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompilationUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompilationUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitCompilationUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) CompilationUnit() (localctx ICompilationUnitContext) {
	localctx = NewCompilationUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GoScriptParserRULE_compilationUnit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2095108) != 0 {
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case GoScriptParserT__1:
			{
				p.SetState(86)
				p.FunctionDeclaration()
			}

		case GoScriptParserT__10, GoScriptParserT__11, GoScriptParserT__12, GoScriptParserT__13, GoScriptParserT__14, GoScriptParserT__15, GoScriptParserT__16, GoScriptParserT__17, GoScriptParserT__18, GoScriptParserT__19:
			{
				p.SetState(87)
				p.VariableDeclaration()
			}
			{
				p.SetState(88)
				p.Match(GoScriptParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(95)
		p.Match(GoScriptParserEOF)
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

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	FormalParameters() IFormalParametersContext
	Block() IBlockContext
	ReturnType() IReturnTypeContext

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_functionDeclaration
	return p
}

func InitEmptyFunctionDeclarationContext(p *FunctionDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_functionDeclaration
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(GoScriptParserIdentifier, 0)
}

func (s *FunctionDeclarationContext) FormalParameters() IFormalParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormalParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormalParametersContext)
}

func (s *FunctionDeclarationContext) Block() IBlockContext {
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

func (s *FunctionDeclarationContext) ReturnType() IReturnTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnTypeContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GoScriptParserRULE_functionDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(GoScriptParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.Match(GoScriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
		p.FormalParameters()
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2095112) != 0 {
		{
			p.SetState(100)
			p.ReturnType()
		}

	}
	{
		p.SetState(103)
		p.Block()
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

// IFormalParametersContext is an interface to support dynamic dispatch.
type IFormalParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFormalParameterDecl() []IFormalParameterDeclContext
	FormalParameterDecl(i int) IFormalParameterDeclContext

	// IsFormalParametersContext differentiates from other interfaces.
	IsFormalParametersContext()
}

type FormalParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParametersContext() *FormalParametersContext {
	var p = new(FormalParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_formalParameters
	return p
}

func InitEmptyFormalParametersContext(p *FormalParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_formalParameters
}

func (*FormalParametersContext) IsFormalParametersContext() {}

func NewFormalParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParametersContext {
	var p = new(FormalParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_formalParameters

	return p
}

func (s *FormalParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParametersContext) AllFormalParameterDecl() []IFormalParameterDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFormalParameterDeclContext); ok {
			len++
		}
	}

	tst := make([]IFormalParameterDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFormalParameterDeclContext); ok {
			tst[i] = t.(IFormalParameterDeclContext)
			i++
		}
	}

	return tst
}

func (s *FormalParametersContext) FormalParameterDecl(i int) IFormalParameterDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormalParameterDeclContext); ok {
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

	return t.(IFormalParameterDeclContext)
}

func (s *FormalParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterFormalParameters(s)
	}
}

func (s *FormalParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitFormalParameters(s)
	}
}

func (s *FormalParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitFormalParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) FormalParameters() (localctx IFormalParametersContext) {
	localctx = NewFormalParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GoScriptParserRULE_formalParameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(GoScriptParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2095104) != 0 {
		{
			p.SetState(106)
			p.FormalParameterDecl()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GoScriptParserT__3 {
			{
				p.SetState(107)
				p.Match(GoScriptParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(108)
				p.FormalParameterDecl()
			}

			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(116)
		p.Match(GoScriptParserT__4)
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

// IFormalParameterDeclContext is an interface to support dynamic dispatch.
type IFormalParameterDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() IType_Context
	Identifier() antlr.TerminalNode

	// IsFormalParameterDeclContext differentiates from other interfaces.
	IsFormalParameterDeclContext()
}

type FormalParameterDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParameterDeclContext() *FormalParameterDeclContext {
	var p = new(FormalParameterDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_formalParameterDecl
	return p
}

func InitEmptyFormalParameterDeclContext(p *FormalParameterDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_formalParameterDecl
}

func (*FormalParameterDeclContext) IsFormalParameterDeclContext() {}

func NewFormalParameterDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParameterDeclContext {
	var p = new(FormalParameterDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_formalParameterDecl

	return p
}

func (s *FormalParameterDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParameterDeclContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FormalParameterDeclContext) Identifier() antlr.TerminalNode {
	return s.GetToken(GoScriptParserIdentifier, 0)
}

func (s *FormalParameterDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParameterDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParameterDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterFormalParameterDecl(s)
	}
}

func (s *FormalParameterDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitFormalParameterDecl(s)
	}
}

func (s *FormalParameterDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitFormalParameterDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) FormalParameterDecl() (localctx IFormalParameterDeclContext) {
	localctx = NewFormalParameterDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GoScriptParserRULE_formalParameterDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Type_()
	}
	{
		p.SetState(119)
		p.Match(GoScriptParserIdentifier)
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

// IReturnTypeContext is an interface to support dynamic dispatch.
type IReturnTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllType_() []IType_Context
	Type_(i int) IType_Context

	// IsReturnTypeContext differentiates from other interfaces.
	IsReturnTypeContext()
}

type ReturnTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnTypeContext() *ReturnTypeContext {
	var p = new(ReturnTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_returnType
	return p
}

func InitEmptyReturnTypeContext(p *ReturnTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_returnType
}

func (*ReturnTypeContext) IsReturnTypeContext() {}

func NewReturnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnTypeContext {
	var p = new(ReturnTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_returnType

	return p
}

func (s *ReturnTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnTypeContext) AllType_() []IType_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType_Context); ok {
			len++
		}
	}

	tst := make([]IType_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType_Context); ok {
			tst[i] = t.(IType_Context)
			i++
		}
	}

	return tst
}

func (s *ReturnTypeContext) Type_(i int) IType_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
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

	return t.(IType_Context)
}

func (s *ReturnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterReturnType(s)
	}
}

func (s *ReturnTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitReturnType(s)
	}
}

func (s *ReturnTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitReturnType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ReturnType() (localctx IReturnTypeContext) {
	localctx = NewReturnTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GoScriptParserRULE_returnType)
	var _la int

	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__10, GoScriptParserT__11, GoScriptParserT__12, GoScriptParserT__13, GoScriptParserT__14, GoScriptParserT__15, GoScriptParserT__16, GoScriptParserT__17, GoScriptParserT__18, GoScriptParserT__19:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Type_()
		}

	case GoScriptParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Match(GoScriptParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Type_()
		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GoScriptParserT__3 {
			{
				p.SetState(124)
				p.Match(GoScriptParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(125)
				p.Type_()
			}

			p.SetState(130)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(131)
			p.Match(GoScriptParserT__4)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBlockStatement() []IBlockStatementContext
	BlockStatement(i int) IBlockStatementContext

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
	p.RuleIndex = GoScriptParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllBlockStatement() []IBlockStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockStatementContext); ok {
			len++
		}
	}

	tst := make([]IBlockStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockStatementContext); ok {
			tst[i] = t.(IBlockStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) BlockStatement(i int) IBlockStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
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

	return t.(IBlockStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GoScriptParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(GoScriptParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126118348596443208) != 0 {
		{
			p.SetState(136)
			p.BlockStatement()
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(142)
		p.Match(GoScriptParserT__6)
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

// IBlockStatementContext is an interface to support dynamic dispatch.
type IBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VariableDeclaration() IVariableDeclarationContext
	Statement() IStatementContext

	// IsBlockStatementContext differentiates from other interfaces.
	IsBlockStatementContext()
}

type BlockStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementContext() *BlockStatementContext {
	var p = new(BlockStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_blockStatement
	return p
}

func InitEmptyBlockStatementContext(p *BlockStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_blockStatement
}

func (*BlockStatementContext) IsBlockStatementContext() {}

func NewBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementContext {
	var p = new(BlockStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_blockStatement

	return p
}

func (s *BlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *BlockStatementContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterBlockStatement(s)
	}
}

func (s *BlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitBlockStatement(s)
	}
}

func (s *BlockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitBlockStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) BlockStatement() (localctx IBlockStatementContext) {
	localctx = NewBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GoScriptParserRULE_blockStatement)
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__10, GoScriptParserT__11, GoScriptParserT__12, GoScriptParserT__13, GoScriptParserT__14, GoScriptParserT__15, GoScriptParserT__16, GoScriptParserT__17, GoScriptParserT__18, GoScriptParserT__19:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.VariableDeclaration()
		}
		{
			p.SetState(145)
			p.Match(GoScriptParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoScriptParserT__2, GoScriptParserT__5, GoScriptParserT__20, GoScriptParserT__22, GoScriptParserT__23, GoScriptParserT__24, GoScriptParserT__25, GoScriptParserT__29, GoScriptParserBooleanLiteral, GoScriptParserIdentifier, GoScriptParserNULL, GoScriptParserHexLiteral, GoScriptParserDecimalLiteral, GoScriptParserOctalLiteral, GoScriptParserFloatingPointLiteral, GoScriptParserPlus, GoScriptParserMinus, GoScriptParserNot, GoScriptParserCharacterLiteral, GoScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(147)
			p.Statement()
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

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() IType_Context
	VariableDeclarators() IVariableDeclaratorsContext

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_variableDeclaration
	return p
}

func InitEmptyVariableDeclarationContext(p *VariableDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_variableDeclaration
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *VariableDeclarationContext) VariableDeclarators() IVariableDeclaratorsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclaratorsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorsContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GoScriptParserRULE_variableDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Type_()
	}
	{
		p.SetState(151)
		p.VariableDeclarators()
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

// IVariableDeclaratorsContext is an interface to support dynamic dispatch.
type IVariableDeclaratorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVariableDeclarator() []IVariableDeclaratorContext
	VariableDeclarator(i int) IVariableDeclaratorContext

	// IsVariableDeclaratorsContext differentiates from other interfaces.
	IsVariableDeclaratorsContext()
}

type VariableDeclaratorsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorsContext() *VariableDeclaratorsContext {
	var p = new(VariableDeclaratorsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_variableDeclarators
	return p
}

func InitEmptyVariableDeclaratorsContext(p *VariableDeclaratorsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_variableDeclarators
}

func (*VariableDeclaratorsContext) IsVariableDeclaratorsContext() {}

func NewVariableDeclaratorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorsContext {
	var p = new(VariableDeclaratorsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_variableDeclarators

	return p
}

func (s *VariableDeclaratorsContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorsContext) AllVariableDeclarator() []IVariableDeclaratorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDeclaratorContext); ok {
			len++
		}
	}

	tst := make([]IVariableDeclaratorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDeclaratorContext); ok {
			tst[i] = t.(IVariableDeclaratorContext)
			i++
		}
	}

	return tst
}

func (s *VariableDeclaratorsContext) VariableDeclarator(i int) IVariableDeclaratorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclaratorContext); ok {
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

	return t.(IVariableDeclaratorContext)
}

func (s *VariableDeclaratorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterVariableDeclarators(s)
	}
}

func (s *VariableDeclaratorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitVariableDeclarators(s)
	}
}

func (s *VariableDeclaratorsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitVariableDeclarators(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) VariableDeclarators() (localctx IVariableDeclaratorsContext) {
	localctx = NewVariableDeclaratorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GoScriptParserRULE_variableDeclarators)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.VariableDeclarator()
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoScriptParserT__3 {
		{
			p.SetState(154)
			p.Match(GoScriptParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.VariableDeclarator()
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

// IVariableDeclaratorContext is an interface to support dynamic dispatch.
type IVariableDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Assign() antlr.TerminalNode
	VariableInitializer() IVariableInitializerContext

	// IsVariableDeclaratorContext differentiates from other interfaces.
	IsVariableDeclaratorContext()
}

type VariableDeclaratorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorContext() *VariableDeclaratorContext {
	var p = new(VariableDeclaratorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_variableDeclarator
	return p
}

func InitEmptyVariableDeclaratorContext(p *VariableDeclaratorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_variableDeclarator
}

func (*VariableDeclaratorContext) IsVariableDeclaratorContext() {}

func NewVariableDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorContext {
	var p = new(VariableDeclaratorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_variableDeclarator

	return p
}

func (s *VariableDeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorContext) Identifier() antlr.TerminalNode {
	return s.GetToken(GoScriptParserIdentifier, 0)
}

func (s *VariableDeclaratorContext) Assign() antlr.TerminalNode {
	return s.GetToken(GoScriptParserAssign, 0)
}

func (s *VariableDeclaratorContext) VariableInitializer() IVariableInitializerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableInitializerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableInitializerContext)
}

func (s *VariableDeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterVariableDeclarator(s)
	}
}

func (s *VariableDeclaratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitVariableDeclarator(s)
	}
}

func (s *VariableDeclaratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitVariableDeclarator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) VariableDeclarator() (localctx IVariableDeclaratorContext) {
	localctx = NewVariableDeclaratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GoScriptParserRULE_variableDeclarator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(GoScriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GoScriptParserAssign {
		{
			p.SetState(162)
			p.Match(GoScriptParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.VariableInitializer()
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

// IVariableInitializerContext is an interface to support dynamic dispatch.
type IVariableInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ArrayInitializer() IArrayInitializerContext
	MapInitializer() IMapInitializerContext
	Expression() IExpressionContext

	// IsVariableInitializerContext differentiates from other interfaces.
	IsVariableInitializerContext()
}

type VariableInitializerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableInitializerContext() *VariableInitializerContext {
	var p = new(VariableInitializerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_variableInitializer
	return p
}

func InitEmptyVariableInitializerContext(p *VariableInitializerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_variableInitializer
}

func (*VariableInitializerContext) IsVariableInitializerContext() {}

func NewVariableInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableInitializerContext {
	var p = new(VariableInitializerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_variableInitializer

	return p
}

func (s *VariableInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableInitializerContext) ArrayInitializer() IArrayInitializerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayInitializerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayInitializerContext)
}

func (s *VariableInitializerContext) MapInitializer() IMapInitializerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapInitializerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapInitializerContext)
}

func (s *VariableInitializerContext) Expression() IExpressionContext {
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

func (s *VariableInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterVariableInitializer(s)
	}
}

func (s *VariableInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitVariableInitializer(s)
	}
}

func (s *VariableInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitVariableInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) VariableInitializer() (localctx IVariableInitializerContext) {
	localctx = NewVariableInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GoScriptParserRULE_variableInitializer)
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.ArrayInitializer()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.MapInitializer()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(168)
			p.expression(0)
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

// IArrayInitializerContext is an interface to support dynamic dispatch.
type IArrayInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVariableInitializer() []IVariableInitializerContext
	VariableInitializer(i int) IVariableInitializerContext

	// IsArrayInitializerContext differentiates from other interfaces.
	IsArrayInitializerContext()
}

type ArrayInitializerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayInitializerContext() *ArrayInitializerContext {
	var p = new(ArrayInitializerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_arrayInitializer
	return p
}

func InitEmptyArrayInitializerContext(p *ArrayInitializerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_arrayInitializer
}

func (*ArrayInitializerContext) IsArrayInitializerContext() {}

func NewArrayInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayInitializerContext {
	var p = new(ArrayInitializerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_arrayInitializer

	return p
}

func (s *ArrayInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayInitializerContext) AllVariableInitializer() []IVariableInitializerContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableInitializerContext); ok {
			len++
		}
	}

	tst := make([]IVariableInitializerContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableInitializerContext); ok {
			tst[i] = t.(IVariableInitializerContext)
			i++
		}
	}

	return tst
}

func (s *ArrayInitializerContext) VariableInitializer(i int) IVariableInitializerContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableInitializerContext); ok {
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

	return t.(IVariableInitializerContext)
}

func (s *ArrayInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterArrayInitializer(s)
	}
}

func (s *ArrayInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitArrayInitializer(s)
	}
}

func (s *ArrayInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitArrayInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ArrayInitializer() (localctx IArrayInitializerContext) {
	localctx = NewArrayInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GoScriptParserRULE_arrayInitializer)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(GoScriptParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126118348466421832) != 0 {
		{
			p.SetState(172)
			p.VariableInitializer()
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GoScriptParserT__3 {
			{
				p.SetState(173)
				p.Match(GoScriptParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(174)
				p.VariableInitializer()
			}

			p.SetState(179)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(182)
		p.Match(GoScriptParserT__6)
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

// IMapInitializerContext is an interface to support dynamic dispatch.
type IMapInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMapEntry() []IMapEntryContext
	MapEntry(i int) IMapEntryContext

	// IsMapInitializerContext differentiates from other interfaces.
	IsMapInitializerContext()
}

type MapInitializerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapInitializerContext() *MapInitializerContext {
	var p = new(MapInitializerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_mapInitializer
	return p
}

func InitEmptyMapInitializerContext(p *MapInitializerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_mapInitializer
}

func (*MapInitializerContext) IsMapInitializerContext() {}

func NewMapInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapInitializerContext {
	var p = new(MapInitializerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_mapInitializer

	return p
}

func (s *MapInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *MapInitializerContext) AllMapEntry() []IMapEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMapEntryContext); ok {
			len++
		}
	}

	tst := make([]IMapEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMapEntryContext); ok {
			tst[i] = t.(IMapEntryContext)
			i++
		}
	}

	return tst
}

func (s *MapInitializerContext) MapEntry(i int) IMapEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapEntryContext); ok {
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

	return t.(IMapEntryContext)
}

func (s *MapInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterMapInitializer(s)
	}
}

func (s *MapInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitMapInitializer(s)
	}
}

func (s *MapInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitMapInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) MapInitializer() (localctx IMapInitializerContext) {
	localctx = NewMapInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GoScriptParserRULE_mapInitializer)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(GoScriptParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.MapEntry()
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(186)
				p.Match(GoScriptParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(187)
				p.MapEntry()
			}

		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GoScriptParserT__3 {
		{
			p.SetState(193)
			p.Match(GoScriptParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(196)
		p.Match(GoScriptParserT__6)
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

// IMapEntryContext is an interface to support dynamic dispatch.
type IMapEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	VariableInitializer() IVariableInitializerContext

	// IsMapEntryContext differentiates from other interfaces.
	IsMapEntryContext()
}

type MapEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapEntryContext() *MapEntryContext {
	var p = new(MapEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_mapEntry
	return p
}

func InitEmptyMapEntryContext(p *MapEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_mapEntry
}

func (*MapEntryContext) IsMapEntryContext() {}

func NewMapEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapEntryContext {
	var p = new(MapEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_mapEntry

	return p
}

func (s *MapEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *MapEntryContext) Expression() IExpressionContext {
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

func (s *MapEntryContext) VariableInitializer() IVariableInitializerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableInitializerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableInitializerContext)
}

func (s *MapEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterMapEntry(s)
	}
}

func (s *MapEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitMapEntry(s)
	}
}

func (s *MapEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitMapEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) MapEntry() (localctx IMapEntryContext) {
	localctx = NewMapEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GoScriptParserRULE_mapEntry)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.expression(0)
	}
	{
		p.SetState(199)
		p.Match(GoScriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		p.VariableInitializer()
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

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimitiveType() IPrimitiveTypeContext
	MapType() IMapTypeContext
	ConnectorType() IConnectorTypeContext
	DynamicType() IDynamicTypeContext

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_type_
	return p
}

func InitEmptyType_Context(p *Type_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_type_
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *Type_Context) MapType() IMapTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapTypeContext)
}

func (s *Type_Context) ConnectorType() IConnectorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConnectorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConnectorTypeContext)
}

func (s *Type_Context) DynamicType() IDynamicTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicTypeContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitType_(s)
	}
}

func (s *Type_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitType_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GoScriptParserRULE_type_)
	var _la int

	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__14, GoScriptParserT__15, GoScriptParserT__16, GoScriptParserT__17, GoScriptParserT__18, GoScriptParserT__19:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
			p.PrimitiveType()
		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GoScriptParserT__8 {
			{
				p.SetState(203)
				p.Match(GoScriptParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(204)
				p.Match(GoScriptParserT__9)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(209)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case GoScriptParserT__10:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(210)
			p.MapType()
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GoScriptParserT__8 {
			{
				p.SetState(211)
				p.Match(GoScriptParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(212)
				p.Match(GoScriptParserT__9)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(217)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case GoScriptParserT__11:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(218)
			p.ConnectorType()
		}
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GoScriptParserT__8 {
			{
				p.SetState(219)
				p.Match(GoScriptParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(220)
				p.Match(GoScriptParserT__9)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(225)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case GoScriptParserT__12, GoScriptParserT__13:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(226)
			p.DynamicType()
		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GoScriptParserT__8 {
			{
				p.SetState(227)
				p.Match(GoScriptParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(228)
				p.Match(GoScriptParserT__9)
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

// IMapTypeContext is an interface to support dynamic dispatch.
type IMapTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lt() antlr.TerminalNode
	PrimitiveType() IPrimitiveTypeContext
	Type_() IType_Context
	Gt() antlr.TerminalNode

	// IsMapTypeContext differentiates from other interfaces.
	IsMapTypeContext()
}

type MapTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapTypeContext() *MapTypeContext {
	var p = new(MapTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_mapType
	return p
}

func InitEmptyMapTypeContext(p *MapTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_mapType
}

func (*MapTypeContext) IsMapTypeContext() {}

func NewMapTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTypeContext {
	var p = new(MapTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_mapType

	return p
}

func (s *MapTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MapTypeContext) Lt() antlr.TerminalNode {
	return s.GetToken(GoScriptParserLt, 0)
}

func (s *MapTypeContext) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *MapTypeContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *MapTypeContext) Gt() antlr.TerminalNode {
	return s.GetToken(GoScriptParserGt, 0)
}

func (s *MapTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterMapType(s)
	}
}

func (s *MapTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitMapType(s)
	}
}

func (s *MapTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitMapType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) MapType() (localctx IMapTypeContext) {
	localctx = NewMapTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GoScriptParserRULE_mapType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(GoScriptParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Match(GoScriptParserLt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)
		p.PrimitiveType()
	}
	{
		p.SetState(239)
		p.Match(GoScriptParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(240)
		p.Type_()
	}
	{
		p.SetState(241)
		p.Match(GoScriptParserGt)
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

// IConnectorTypeContext is an interface to support dynamic dispatch.
type IConnectorTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lt() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Gt() antlr.TerminalNode

	// IsConnectorTypeContext differentiates from other interfaces.
	IsConnectorTypeContext()
}

type ConnectorTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConnectorTypeContext() *ConnectorTypeContext {
	var p = new(ConnectorTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_connectorType
	return p
}

func InitEmptyConnectorTypeContext(p *ConnectorTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_connectorType
}

func (*ConnectorTypeContext) IsConnectorTypeContext() {}

func NewConnectorTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConnectorTypeContext {
	var p = new(ConnectorTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_connectorType

	return p
}

func (s *ConnectorTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ConnectorTypeContext) Lt() antlr.TerminalNode {
	return s.GetToken(GoScriptParserLt, 0)
}

func (s *ConnectorTypeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(GoScriptParserIdentifier, 0)
}

func (s *ConnectorTypeContext) Gt() antlr.TerminalNode {
	return s.GetToken(GoScriptParserGt, 0)
}

func (s *ConnectorTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConnectorTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConnectorTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterConnectorType(s)
	}
}

func (s *ConnectorTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitConnectorType(s)
	}
}

func (s *ConnectorTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitConnectorType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ConnectorType() (localctx IConnectorTypeContext) {
	localctx = NewConnectorTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GoScriptParserRULE_connectorType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(GoScriptParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Match(GoScriptParserLt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.Match(GoScriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)
		p.Match(GoScriptParserGt)
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

// IDynamicTypeContext is an interface to support dynamic dispatch.
type IDynamicTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDynamicTypeContext differentiates from other interfaces.
	IsDynamicTypeContext()
}

type DynamicTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDynamicTypeContext() *DynamicTypeContext {
	var p = new(DynamicTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_dynamicType
	return p
}

func InitEmptyDynamicTypeContext(p *DynamicTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_dynamicType
}

func (*DynamicTypeContext) IsDynamicTypeContext() {}

func NewDynamicTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DynamicTypeContext {
	var p = new(DynamicTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_dynamicType

	return p
}

func (s *DynamicTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *DynamicTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DynamicTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterDynamicType(s)
	}
}

func (s *DynamicTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitDynamicType(s)
	}
}

func (s *DynamicTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitDynamicType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) DynamicType() (localctx IDynamicTypeContext) {
	localctx = NewDynamicTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GoScriptParserRULE_dynamicType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GoScriptParserT__12 || _la == GoScriptParserT__13) {
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

// IPrimitiveTypeContext is an interface to support dynamic dispatch.
type IPrimitiveTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimitiveTypeContext differentiates from other interfaces.
	IsPrimitiveTypeContext()
}

type PrimitiveTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveTypeContext() *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_primitiveType
	return p
}

func InitEmptyPrimitiveTypeContext(p *PrimitiveTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_primitiveType
}

func (*PrimitiveTypeContext) IsPrimitiveTypeContext() {}

func NewPrimitiveTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_primitiveType

	return p
}

func (s *PrimitiveTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *PrimitiveTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterPrimitiveType(s)
	}
}

func (s *PrimitiveTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitPrimitiveType(s)
	}
}

func (s *PrimitiveTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitPrimitiveType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) PrimitiveType() (localctx IPrimitiveTypeContext) {
	localctx = NewPrimitiveTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GoScriptParserRULE_primitiveType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2064384) != 0) {
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	IfStatement() IIfStatementContext
	ForStatement() IForStatementContext
	ReturnStatement() IReturnStatementContext
	BreakStatement() IBreakStatementContext
	ContinueStatement() IContinueStatementContext
	ExpressionStatement() IExpressionStatementContext

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
	p.RuleIndex = GoScriptParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Block() IBlockContext {
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

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
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

func (s *StatementContext) BreakStatement() IBreakStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStatementContext)
}

func (s *StatementContext) ContinueStatement() IContinueStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStatementContext)
}

func (s *StatementContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GoScriptParserRULE_statement)
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(252)
			p.Block()
		}

	case GoScriptParserT__20:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)
			p.IfStatement()
		}

	case GoScriptParserT__22:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(254)
			p.ForStatement()
		}

	case GoScriptParserT__23:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(255)
			p.ReturnStatement()
		}

	case GoScriptParserT__24:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(256)
			p.BreakStatement()
		}

	case GoScriptParserT__25:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(257)
			p.ContinueStatement()
		}

	case GoScriptParserT__2, GoScriptParserT__29, GoScriptParserBooleanLiteral, GoScriptParserIdentifier, GoScriptParserNULL, GoScriptParserHexLiteral, GoScriptParserDecimalLiteral, GoScriptParserOctalLiteral, GoScriptParserFloatingPointLiteral, GoScriptParserPlus, GoScriptParserMinus, GoScriptParserNot, GoScriptParserCharacterLiteral, GoScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(258)
			p.ExpressionStatement()
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

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) Expression() IExpressionContext {
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

func (s *IfStatementContext) AllStatement() []IStatementContext {
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

func (s *IfStatementContext) Statement(i int) IStatementContext {
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

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GoScriptParserRULE_ifStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Match(GoScriptParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.expression(0)
	}
	{
		p.SetState(263)
		p.Statement()
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(264)
			p.Match(GoScriptParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)
			p.Statement()
		}

	} else if p.HasError() { // JIM
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

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ForControl() IForControlContext
	Statement() IStatementContext

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_forStatement
	return p
}

func InitEmptyForStatementContext(p *ForStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_forStatement
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) ForControl() IForControlContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForControlContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForControlContext)
}

func (s *ForStatementContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitForStatement(s)
	}
}

func (s *ForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GoScriptParserRULE_forStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(GoScriptParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.Match(GoScriptParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(270)
		p.ForControl()
	}
	{
		p.SetState(271)
		p.Match(GoScriptParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(272)
		p.Statement()
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

// IForControlContext is an interface to support dynamic dispatch.
type IForControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ForInit() IForInitContext
	Expression() IExpressionContext
	ForUpdate() IForUpdateContext

	// IsForControlContext differentiates from other interfaces.
	IsForControlContext()
}

type ForControlContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForControlContext() *ForControlContext {
	var p = new(ForControlContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_forControl
	return p
}

func InitEmptyForControlContext(p *ForControlContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_forControl
}

func (*ForControlContext) IsForControlContext() {}

func NewForControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForControlContext {
	var p = new(ForControlContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_forControl

	return p
}

func (s *ForControlContext) GetParser() antlr.Parser { return s.parser }

func (s *ForControlContext) ForInit() IForInitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForInitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForInitContext)
}

func (s *ForControlContext) Expression() IExpressionContext {
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

func (s *ForControlContext) ForUpdate() IForUpdateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForUpdateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForUpdateContext)
}

func (s *ForControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForControlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterForControl(s)
	}
}

func (s *ForControlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitForControl(s)
	}
}

func (s *ForControlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitForControl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ForControl() (localctx IForControlContext) {
	localctx = NewForControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GoScriptParserRULE_forControl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126118348468516872) != 0 {
		{
			p.SetState(274)
			p.ForInit()
		}

	}
	{
		p.SetState(277)
		p.Match(GoScriptParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126118348466421768) != 0 {
		{
			p.SetState(278)
			p.expression(0)
		}

	}
	{
		p.SetState(281)
		p.Match(GoScriptParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126118348466421768) != 0 {
		{
			p.SetState(282)
			p.ForUpdate()
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

// IForInitContext is an interface to support dynamic dispatch.
type IForInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VariableDeclaration() IVariableDeclarationContext
	ExpressionList() IExpressionListContext

	// IsForInitContext differentiates from other interfaces.
	IsForInitContext()
}

type ForInitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForInitContext() *ForInitContext {
	var p = new(ForInitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_forInit
	return p
}

func InitEmptyForInitContext(p *ForInitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_forInit
}

func (*ForInitContext) IsForInitContext() {}

func NewForInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForInitContext {
	var p = new(ForInitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_forInit

	return p
}

func (s *ForInitContext) GetParser() antlr.Parser { return s.parser }

func (s *ForInitContext) VariableDeclaration() IVariableDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *ForInitContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ForInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterForInit(s)
	}
}

func (s *ForInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitForInit(s)
	}
}

func (s *ForInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitForInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ForInit() (localctx IForInitContext) {
	localctx = NewForInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GoScriptParserRULE_forInit)
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__10, GoScriptParserT__11, GoScriptParserT__12, GoScriptParserT__13, GoScriptParserT__14, GoScriptParserT__15, GoScriptParserT__16, GoScriptParserT__17, GoScriptParserT__18, GoScriptParserT__19:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)
			p.VariableDeclaration()
		}

	case GoScriptParserT__2, GoScriptParserT__29, GoScriptParserBooleanLiteral, GoScriptParserIdentifier, GoScriptParserNULL, GoScriptParserHexLiteral, GoScriptParserDecimalLiteral, GoScriptParserOctalLiteral, GoScriptParserFloatingPointLiteral, GoScriptParserPlus, GoScriptParserMinus, GoScriptParserNot, GoScriptParserCharacterLiteral, GoScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(286)
			p.ExpressionList()
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

// IForUpdateContext is an interface to support dynamic dispatch.
type IForUpdateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionList() IExpressionListContext

	// IsForUpdateContext differentiates from other interfaces.
	IsForUpdateContext()
}

type ForUpdateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForUpdateContext() *ForUpdateContext {
	var p = new(ForUpdateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_forUpdate
	return p
}

func InitEmptyForUpdateContext(p *ForUpdateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_forUpdate
}

func (*ForUpdateContext) IsForUpdateContext() {}

func NewForUpdateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForUpdateContext {
	var p = new(ForUpdateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_forUpdate

	return p
}

func (s *ForUpdateContext) GetParser() antlr.Parser { return s.parser }

func (s *ForUpdateContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ForUpdateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForUpdateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForUpdateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterForUpdate(s)
	}
}

func (s *ForUpdateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitForUpdate(s)
	}
}

func (s *ForUpdateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitForUpdate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ForUpdate() (localctx IForUpdateContext) {
	localctx = NewForUpdateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GoScriptParserRULE_forUpdate)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.ExpressionList()
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
	Expression() IExpressionContext

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
	p.RuleIndex = GoScriptParserRULE_returnStatement
	return p
}

func InitEmptyReturnStatementContext(p *ReturnStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_returnStatement
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) Expression() IExpressionContext {
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

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GoScriptParserRULE_returnStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Match(GoScriptParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126118348466421768) != 0 {
		{
			p.SetState(292)
			p.expression(0)
		}

	}
	{
		p.SetState(295)
		p.Match(GoScriptParserT__0)
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

// IBreakStatementContext is an interface to support dynamic dispatch.
type IBreakStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBreakStatementContext differentiates from other interfaces.
	IsBreakStatementContext()
}

type BreakStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStatementContext() *BreakStatementContext {
	var p = new(BreakStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_breakStatement
	return p
}

func InitEmptyBreakStatementContext(p *BreakStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_breakStatement
}

func (*BreakStatementContext) IsBreakStatementContext() {}

func NewBreakStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStatementContext {
	var p = new(BreakStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_breakStatement

	return p
}

func (s *BreakStatementContext) GetParser() antlr.Parser { return s.parser }
func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) BreakStatement() (localctx IBreakStatementContext) {
	localctx = NewBreakStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GoScriptParserRULE_breakStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Match(GoScriptParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(298)
		p.Match(GoScriptParserT__0)
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

// IContinueStatementContext is an interface to support dynamic dispatch.
type IContinueStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsContinueStatementContext differentiates from other interfaces.
	IsContinueStatementContext()
}

type ContinueStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStatementContext() *ContinueStatementContext {
	var p = new(ContinueStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_continueStatement
	return p
}

func InitEmptyContinueStatementContext(p *ContinueStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_continueStatement
}

func (*ContinueStatementContext) IsContinueStatementContext() {}

func NewContinueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_continueStatement

	return p
}

func (s *ContinueStatementContext) GetParser() antlr.Parser { return s.parser }
func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ContinueStatement() (localctx IContinueStatementContext) {
	localctx = NewContinueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GoScriptParserRULE_continueStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.Match(GoScriptParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(301)
		p.Match(GoScriptParserT__0)
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

// IExpressionStatementContext is an interface to support dynamic dispatch.
type IExpressionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsExpressionStatementContext differentiates from other interfaces.
	IsExpressionStatementContext()
}

type ExpressionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStatementContext() *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_expressionStatement
	return p
}

func InitEmptyExpressionStatementContext(p *ExpressionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_expressionStatement
}

func (*ExpressionStatementContext) IsExpressionStatementContext() {}

func NewExpressionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_expressionStatement

	return p
}

func (s *ExpressionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStatementContext) Expression() IExpressionContext {
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

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitExpressionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ExpressionStatement() (localctx IExpressionStatementContext) {
	localctx = NewExpressionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GoScriptParserRULE_expressionStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		p.expression(0)
	}
	{
		p.SetState(304)
		p.Match(GoScriptParserT__0)
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
	p.RuleIndex = GoScriptParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_expression

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

type CreateAndAssignExprContext struct {
	ExpressionContext
}

func NewCreateAndAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CreateAndAssignExprContext {
	var p = new(CreateAndAssignExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CreateAndAssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateAndAssignExprContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *CreateAndAssignExprContext) AllExpression() []IExpressionContext {
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

func (s *CreateAndAssignExprContext) Expression(i int) IExpressionContext {
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

func (s *CreateAndAssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterCreateAndAssignExpr(s)
	}
}

func (s *CreateAndAssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitCreateAndAssignExpr(s)
	}
}

func (s *CreateAndAssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitCreateAndAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulExprContext struct {
	ExpressionContext
}

func NewMulExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulExprContext {
	var p = new(MulExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MulExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExprContext) AllExpression() []IExpressionContext {
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

func (s *MulExprContext) Expression(i int) IExpressionContext {
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

func (s *MulExprContext) Multiplication() antlr.TerminalNode {
	return s.GetToken(GoScriptParserMultiplication, 0)
}

func (s *MulExprContext) Division() antlr.TerminalNode {
	return s.GetToken(GoScriptParserDivision, 0)
}

func (s *MulExprContext) Percent() antlr.TerminalNode {
	return s.GetToken(GoScriptParserPercent, 0)
}

func (s *MulExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterMulExpr(s)
	}
}

func (s *MulExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitMulExpr(s)
	}
}

func (s *MulExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitMulExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExprContext struct {
	ExpressionContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpression() []IExpressionContext {
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

func (s *AndExprContext) Expression(i int) IExpressionContext {
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

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddExprContext struct {
	ExpressionContext
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) AllExpression() []IExpressionContext {
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

func (s *AddExprContext) Expression(i int) IExpressionContext {
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

func (s *AddExprContext) Plus() antlr.TerminalNode {
	return s.GetToken(GoScriptParserPlus, 0)
}

func (s *AddExprContext) Minus() antlr.TerminalNode {
	return s.GetToken(GoScriptParserMinus, 0)
}

func (s *AddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterAddExpr(s)
	}
}

func (s *AddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitAddExpr(s)
	}
}

func (s *AddExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitAddExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConditionalExprContext struct {
	ExpressionContext
}

func NewConditionalExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionalExprContext {
	var p = new(ConditionalExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ConditionalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalExprContext) AllExpression() []IExpressionContext {
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

func (s *ConditionalExprContext) Expression(i int) IExpressionContext {
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

func (s *ConditionalExprContext) Let() antlr.TerminalNode {
	return s.GetToken(GoScriptParserLet, 0)
}

func (s *ConditionalExprContext) Get() antlr.TerminalNode {
	return s.GetToken(GoScriptParserGet, 0)
}

func (s *ConditionalExprContext) Gt() antlr.TerminalNode {
	return s.GetToken(GoScriptParserGt, 0)
}

func (s *ConditionalExprContext) Lt() antlr.TerminalNode {
	return s.GetToken(GoScriptParserLt, 0)
}

func (s *ConditionalExprContext) Eq() antlr.TerminalNode {
	return s.GetToken(GoScriptParserEq, 0)
}

func (s *ConditionalExprContext) NEq() antlr.TerminalNode {
	return s.GetToken(GoScriptParserNEq, 0)
}

func (s *ConditionalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterConditionalExpr(s)
	}
}

func (s *ConditionalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitConditionalExpr(s)
	}
}

func (s *ConditionalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitConditionalExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExprContext struct {
	ExpressionContext
}

func NewUnaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExprContext {
	var p = new(UnaryExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) Expression() IExpressionContext {
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

func (s *UnaryExprContext) Plus() antlr.TerminalNode {
	return s.GetToken(GoScriptParserPlus, 0)
}

func (s *UnaryExprContext) Minus() antlr.TerminalNode {
	return s.GetToken(GoScriptParserMinus, 0)
}

func (s *UnaryExprContext) Not() antlr.TerminalNode {
	return s.GetToken(GoScriptParserNot, 0)
}

func (s *UnaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterUnaryExpr(s)
	}
}

func (s *UnaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitUnaryExpr(s)
	}
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	ExpressionContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpression() []IExpressionContext {
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

func (s *OrExprContext) Expression(i int) IExpressionContext {
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

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexExprContext struct {
	ExpressionContext
}

func NewIndexExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexExprContext {
	var p = new(IndexExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IndexExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExprContext) AllExpression() []IExpressionContext {
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

func (s *IndexExprContext) Expression(i int) IExpressionContext {
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

func (s *IndexExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterIndexExpr(s)
	}
}

func (s *IndexExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitIndexExpr(s)
	}
}

func (s *IndexExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitIndexExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignExprContext struct {
	ExpressionContext
}

func NewAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignExprContext {
	var p = new(AssignExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExprContext) AllLvalue() []ILvalueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILvalueContext); ok {
			len++
		}
	}

	tst := make([]ILvalueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILvalueContext); ok {
			tst[i] = t.(ILvalueContext)
			i++
		}
	}

	return tst
}

func (s *AssignExprContext) Lvalue(i int) ILvalueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
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

	return t.(ILvalueContext)
}

func (s *AssignExprContext) Assign() antlr.TerminalNode {
	return s.GetToken(GoScriptParserAssign, 0)
}

func (s *AssignExprContext) AllExpression() []IExpressionContext {
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

func (s *AssignExprContext) Expression(i int) IExpressionContext {
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

func (s *AssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterAssignExpr(s)
	}
}

func (s *AssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitAssignExpr(s)
	}
}

func (s *AssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelectorExprContext struct {
	ExpressionContext
}

func NewSelectorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectorExprContext {
	var p = new(SelectorExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SelectorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorExprContext) Expression() IExpressionContext {
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

func (s *SelectorExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(GoScriptParserIdentifier, 0)
}

func (s *SelectorExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterSelectorExpr(s)
	}
}

func (s *SelectorExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitSelectorExpr(s)
	}
}

func (s *SelectorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitSelectorExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CreateExprContext struct {
	ExpressionContext
}

func NewCreateExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CreateExprContext {
	var p = new(CreateExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CreateExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateExprContext) Creator() ICreatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreatorContext)
}

func (s *CreateExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterCreateExpr(s)
	}
}

func (s *CreateExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitCreateExpr(s)
	}
}

func (s *CreateExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitCreateExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelfAddExprContext struct {
	ExpressionContext
}

func NewSelfAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelfAddExprContext {
	var p = new(SelfAddExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SelfAddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelfAddExprContext) Expression() IExpressionContext {
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

func (s *SelfAddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterSelfAddExpr(s)
	}
}

func (s *SelfAddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitSelfAddExpr(s)
	}
}

func (s *SelfAddExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitSelfAddExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExprContext struct {
	ExpressionContext
}

func NewPrimaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallExprContext struct {
	ExpressionContext
}

func NewCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExprContext {
	var p = new(CallExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) Expression() IExpressionContext {
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

func (s *CallExprContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *CallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterCallExpr(s)
	}
}

func (s *CallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitCallExpr(s)
	}
}

func (s *CallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TernaryExprContext struct {
	ExpressionContext
}

func NewTernaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryExprContext {
	var p = new(TernaryExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *TernaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryExprContext) AllExpression() []IExpressionContext {
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

func (s *TernaryExprContext) Expression(i int) IExpressionContext {
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

func (s *TernaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterTernaryExpr(s)
	}
}

func (s *TernaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitTernaryExpr(s)
	}
}

func (s *TernaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitTernaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *GoScriptParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 58
	p.EnterRecursionRule(localctx, 58, GoScriptParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPrimaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(307)
			p.Primary()
		}

	case 2:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(308)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18027592649015296) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(309)
			p.expression(10)
		}

	case 3:
		localctx = NewCreateExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(310)
			p.Match(GoScriptParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(311)
			p.Creator()
		}

	case 4:
		localctx = NewAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(312)
			p.lvalue(0)
		}
		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GoScriptParserT__3 {
			{
				p.SetState(313)
				p.Match(GoScriptParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(314)
				p.lvalue(0)
			}

			p.SetState(319)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(320)
			p.Match(GoScriptParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)
			p.expression(0)
		}
		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(322)
					p.Match(GoScriptParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(323)
					p.expression(0)
				}

			}
			p.SetState(328)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 5:
		localctx = NewCreateAndAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(329)
			p.IdentifierList()
		}
		{
			p.SetState(330)
			p.Match(GoScriptParserT__33)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.expression(0)
		}
		p.SetState(336)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(332)
					p.Match(GoScriptParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(333)
					p.expression(0)
				}

			}
			p.SetState(338)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(380)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(378)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(341)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(342)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&123145302310912) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(343)
					p.expression(9)
				}

			case 2:
				localctx = NewAddExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(344)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(345)
					_la = p.GetTokenStream().LA(1)

					if !(_la == GoScriptParserPlus || _la == GoScriptParserMinus) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(346)
					p.expression(8)
				}

			case 3:
				localctx = NewConditionalExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(347)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(348)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15621861207441408) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(349)
					p.expression(7)
				}

			case 4:
				localctx = NewAndExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(350)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(351)
					p.Match(GoScriptParserT__30)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(352)
					p.expression(6)
				}

			case 5:
				localctx = NewOrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(353)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(354)
					p.Match(GoScriptParserT__31)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(355)
					p.expression(5)
				}

			case 6:
				localctx = NewTernaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(356)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(357)
					p.Match(GoScriptParserT__32)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(358)
					p.expression(0)
				}
				{
					p.SetState(359)
					p.Match(GoScriptParserT__7)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(360)
					p.expression(4)
				}

			case 7:
				localctx = NewSelectorExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(362)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(363)
					p.Match(GoScriptParserT__26)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(364)
					p.Match(GoScriptParserIdentifier)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 8:
				localctx = NewIndexExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(365)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(366)
					p.Match(GoScriptParserT__8)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(367)
					p.expression(0)
				}
				{
					p.SetState(368)
					p.Match(GoScriptParserT__9)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 9:
				localctx = NewCallExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(370)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(371)
					p.Match(GoScriptParserT__2)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(373)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126118348466421768) != 0 {
					{
						p.SetState(372)
						p.ExpressionList()
					}

				}
				{
					p.SetState(375)
					p.Match(GoScriptParserT__4)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 10:
				localctx = NewSelfAddExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(376)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(377)
					_la = p.GetTokenStream().LA(1)

					if !(_la == GoScriptParserT__27 || _la == GoScriptParserT__28) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(382)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
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

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Literal() ILiteralContext
	Identifier() antlr.TerminalNode

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) Expression() IExpressionContext {
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

func (s *PrimaryContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PrimaryContext) Identifier() antlr.TerminalNode {
	return s.GetToken(GoScriptParserIdentifier, 0)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GoScriptParserRULE_primary)
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(383)
			p.Match(GoScriptParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(384)
			p.expression(0)
		}
		{
			p.SetState(385)
			p.Match(GoScriptParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoScriptParserBooleanLiteral, GoScriptParserNULL, GoScriptParserHexLiteral, GoScriptParserDecimalLiteral, GoScriptParserOctalLiteral, GoScriptParserFloatingPointLiteral, GoScriptParserCharacterLiteral, GoScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(387)
			p.Literal()
		}

	case GoScriptParserIdentifier:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(388)
			p.Match(GoScriptParserIdentifier)
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

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntegerLiteral() IIntegerLiteralContext
	FloatingPointLiteral() antlr.TerminalNode
	CharacterLiteral() antlr.TerminalNode
	StringLiteral() antlr.TerminalNode
	BooleanLiteral() antlr.TerminalNode
	NULL() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) IntegerLiteral() IIntegerLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *LiteralContext) FloatingPointLiteral() antlr.TerminalNode {
	return s.GetToken(GoScriptParserFloatingPointLiteral, 0)
}

func (s *LiteralContext) CharacterLiteral() antlr.TerminalNode {
	return s.GetToken(GoScriptParserCharacterLiteral, 0)
}

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(GoScriptParserStringLiteral, 0)
}

func (s *LiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(GoScriptParserBooleanLiteral, 0)
}

func (s *LiteralContext) NULL() antlr.TerminalNode {
	return s.GetToken(GoScriptParserNULL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GoScriptParserRULE_literal)
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserHexLiteral, GoScriptParserDecimalLiteral, GoScriptParserOctalLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(391)
			p.IntegerLiteral()
		}

	case GoScriptParserFloatingPointLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(392)
			p.Match(GoScriptParserFloatingPointLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoScriptParserCharacterLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(393)
			p.Match(GoScriptParserCharacterLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(394)
			p.Match(GoScriptParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoScriptParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(395)
			p.Match(GoScriptParserBooleanLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoScriptParserNULL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(396)
			p.Match(GoScriptParserNULL)
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

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HexLiteral() antlr.TerminalNode
	OctalLiteral() antlr.TerminalNode
	DecimalLiteral() antlr.TerminalNode

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_integerLiteral
	return p
}

func InitEmptyIntegerLiteralContext(p *IntegerLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_integerLiteral
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) HexLiteral() antlr.TerminalNode {
	return s.GetToken(GoScriptParserHexLiteral, 0)
}

func (s *IntegerLiteralContext) OctalLiteral() antlr.TerminalNode {
	return s.GetToken(GoScriptParserOctalLiteral, 0)
}

func (s *IntegerLiteralContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(GoScriptParserDecimalLiteral, 0)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitIntegerLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GoScriptParserRULE_integerLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1924145348608) != 0) {
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

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_expressionList
	return p
}

func InitEmptyExpressionListContext(p *ExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_expressionList
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
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

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
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

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GoScriptParserRULE_expressionList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(401)
		p.expression(0)
	}
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoScriptParserT__3 {
		{
			p.SetState(402)
			p.Match(GoScriptParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
			p.expression(0)
		}

		p.SetState(408)
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

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode

	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_identifierList
	return p
}

func InitEmptyIdentifierListContext(p *IdentifierListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_identifierList
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(GoScriptParserIdentifier)
}

func (s *IdentifierListContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(GoScriptParserIdentifier, i)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterIdentifierList(s)
	}
}

func (s *IdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitIdentifierList(s)
	}
}

func (s *IdentifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitIdentifierList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GoScriptParserRULE_identifierList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(409)
		p.Match(GoScriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoScriptParserT__3 {
		{
			p.SetState(410)
			p.Match(GoScriptParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(411)
			p.Match(GoScriptParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(416)
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

// ILvalueContext is an interface to support dynamic dispatch.
type ILvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Lvalue() ILvalueContext
	Expression() IExpressionContext

	// IsLvalueContext differentiates from other interfaces.
	IsLvalueContext()
}

type LvalueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLvalueContext() *LvalueContext {
	var p = new(LvalueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_lvalue
	return p
}

func InitEmptyLvalueContext(p *LvalueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_lvalue
}

func (*LvalueContext) IsLvalueContext() {}

func NewLvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LvalueContext {
	var p = new(LvalueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_lvalue

	return p
}

func (s *LvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *LvalueContext) Identifier() antlr.TerminalNode {
	return s.GetToken(GoScriptParserIdentifier, 0)
}

func (s *LvalueContext) Lvalue() ILvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *LvalueContext) Expression() IExpressionContext {
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

func (s *LvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterLvalue(s)
	}
}

func (s *LvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitLvalue(s)
	}
}

func (s *LvalueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitLvalue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) Lvalue() (localctx ILvalueContext) {
	return p.lvalue(0)
}

func (p *GoScriptParser) lvalue(_p int) (localctx ILvalueContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewLvalueContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILvalueContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 70
	p.EnterRecursionRule(localctx, 70, GoScriptParserRULE_lvalue, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(418)
		p.Match(GoScriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(428)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
			case 1:
				localctx = NewLvalueContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_lvalue)
				p.SetState(420)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(421)
					p.Match(GoScriptParserT__26)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(422)
					p.Match(GoScriptParserIdentifier)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 2:
				localctx = NewLvalueContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_lvalue)
				p.SetState(423)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(424)
					p.Match(GoScriptParserT__8)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(425)
					p.expression(0)
				}
				{
					p.SetState(426)
					p.Match(GoScriptParserT__9)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
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

// ICreatorContext is an interface to support dynamic dispatch.
type ICreatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MapCreator() IMapCreatorContext
	ArrayCreator() IArrayCreatorContext
	ConnectorCreator() IConnectorCreatorContext
	PrimitiveCreator() IPrimitiveCreatorContext
	DynamicCreator() IDynamicCreatorContext

	// IsCreatorContext differentiates from other interfaces.
	IsCreatorContext()
}

type CreatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreatorContext() *CreatorContext {
	var p = new(CreatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_creator
	return p
}

func InitEmptyCreatorContext(p *CreatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_creator
}

func (*CreatorContext) IsCreatorContext() {}

func NewCreatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreatorContext {
	var p = new(CreatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_creator

	return p
}

func (s *CreatorContext) GetParser() antlr.Parser { return s.parser }

func (s *CreatorContext) MapCreator() IMapCreatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapCreatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapCreatorContext)
}

func (s *CreatorContext) ArrayCreator() IArrayCreatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayCreatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayCreatorContext)
}

func (s *CreatorContext) ConnectorCreator() IConnectorCreatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConnectorCreatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConnectorCreatorContext)
}

func (s *CreatorContext) PrimitiveCreator() IPrimitiveCreatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveCreatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveCreatorContext)
}

func (s *CreatorContext) DynamicCreator() IDynamicCreatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicCreatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicCreatorContext)
}

func (s *CreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterCreator(s)
	}
}

func (s *CreatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitCreator(s)
	}
}

func (s *CreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) Creator() (localctx ICreatorContext) {
	localctx = NewCreatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, GoScriptParserRULE_creator)
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(433)
			p.MapCreator()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(434)
			p.ArrayCreator()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(435)
			p.ConnectorCreator()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(436)
			p.PrimitiveCreator()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(437)
			p.DynamicCreator()
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

// IMapCreatorContext is an interface to support dynamic dispatch.
type IMapCreatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MapType() IMapTypeContext
	MapInitializer() IMapInitializerContext

	// IsMapCreatorContext differentiates from other interfaces.
	IsMapCreatorContext()
}

type MapCreatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapCreatorContext() *MapCreatorContext {
	var p = new(MapCreatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_mapCreator
	return p
}

func InitEmptyMapCreatorContext(p *MapCreatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_mapCreator
}

func (*MapCreatorContext) IsMapCreatorContext() {}

func NewMapCreatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapCreatorContext {
	var p = new(MapCreatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_mapCreator

	return p
}

func (s *MapCreatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MapCreatorContext) MapType() IMapTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapTypeContext)
}

func (s *MapCreatorContext) MapInitializer() IMapInitializerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapInitializerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapInitializerContext)
}

func (s *MapCreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapCreatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapCreatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterMapCreator(s)
	}
}

func (s *MapCreatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitMapCreator(s)
	}
}

func (s *MapCreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitMapCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) MapCreator() (localctx IMapCreatorContext) {
	localctx = NewMapCreatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GoScriptParserRULE_mapCreator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(440)
		p.MapType()
	}
	{
		p.SetState(441)
		p.MapInitializer()
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

// IArrayCreatorContext is an interface to support dynamic dispatch.
type IArrayCreatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CreatorName() ICreatorNameContext
	ArrayInitializer() IArrayInitializerContext

	// IsArrayCreatorContext differentiates from other interfaces.
	IsArrayCreatorContext()
}

type ArrayCreatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayCreatorContext() *ArrayCreatorContext {
	var p = new(ArrayCreatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_arrayCreator
	return p
}

func InitEmptyArrayCreatorContext(p *ArrayCreatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_arrayCreator
}

func (*ArrayCreatorContext) IsArrayCreatorContext() {}

func NewArrayCreatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayCreatorContext {
	var p = new(ArrayCreatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_arrayCreator

	return p
}

func (s *ArrayCreatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayCreatorContext) CreatorName() ICreatorNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreatorNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreatorNameContext)
}

func (s *ArrayCreatorContext) ArrayInitializer() IArrayInitializerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayInitializerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayInitializerContext)
}

func (s *ArrayCreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayCreatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayCreatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterArrayCreator(s)
	}
}

func (s *ArrayCreatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitArrayCreator(s)
	}
}

func (s *ArrayCreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitArrayCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ArrayCreator() (localctx IArrayCreatorContext) {
	localctx = NewArrayCreatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, GoScriptParserRULE_arrayCreator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(443)
		p.CreatorName()
	}
	p.SetState(446)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoScriptParserT__8 {
		{
			p.SetState(444)
			p.Match(GoScriptParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)
			p.Match(GoScriptParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(448)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(450)
		p.ArrayInitializer()
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

// ICreatorNameContext is an interface to support dynamic dispatch.
type ICreatorNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimitiveType() IPrimitiveTypeContext
	MapType() IMapTypeContext
	ConnectorType() IConnectorTypeContext
	DynamicType() IDynamicTypeContext

	// IsCreatorNameContext differentiates from other interfaces.
	IsCreatorNameContext()
}

type CreatorNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreatorNameContext() *CreatorNameContext {
	var p = new(CreatorNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_creatorName
	return p
}

func InitEmptyCreatorNameContext(p *CreatorNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_creatorName
}

func (*CreatorNameContext) IsCreatorNameContext() {}

func NewCreatorNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreatorNameContext {
	var p = new(CreatorNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_creatorName

	return p
}

func (s *CreatorNameContext) GetParser() antlr.Parser { return s.parser }

func (s *CreatorNameContext) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *CreatorNameContext) MapType() IMapTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapTypeContext)
}

func (s *CreatorNameContext) ConnectorType() IConnectorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConnectorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConnectorTypeContext)
}

func (s *CreatorNameContext) DynamicType() IDynamicTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicTypeContext)
}

func (s *CreatorNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreatorNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreatorNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterCreatorName(s)
	}
}

func (s *CreatorNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitCreatorName(s)
	}
}

func (s *CreatorNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitCreatorName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) CreatorName() (localctx ICreatorNameContext) {
	localctx = NewCreatorNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, GoScriptParserRULE_creatorName)
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__14, GoScriptParserT__15, GoScriptParserT__16, GoScriptParserT__17, GoScriptParserT__18, GoScriptParserT__19:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(452)
			p.PrimitiveType()
		}

	case GoScriptParserT__10:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(453)
			p.MapType()
		}

	case GoScriptParserT__11:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(454)
			p.ConnectorType()
		}

	case GoScriptParserT__12, GoScriptParserT__13:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(455)
			p.DynamicType()
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

// IConnectorCreatorContext is an interface to support dynamic dispatch.
type IConnectorCreatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConnectorType() IConnectorTypeContext

	// IsConnectorCreatorContext differentiates from other interfaces.
	IsConnectorCreatorContext()
}

type ConnectorCreatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConnectorCreatorContext() *ConnectorCreatorContext {
	var p = new(ConnectorCreatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_connectorCreator
	return p
}

func InitEmptyConnectorCreatorContext(p *ConnectorCreatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_connectorCreator
}

func (*ConnectorCreatorContext) IsConnectorCreatorContext() {}

func NewConnectorCreatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConnectorCreatorContext {
	var p = new(ConnectorCreatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_connectorCreator

	return p
}

func (s *ConnectorCreatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ConnectorCreatorContext) ConnectorType() IConnectorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConnectorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConnectorTypeContext)
}

func (s *ConnectorCreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConnectorCreatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConnectorCreatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterConnectorCreator(s)
	}
}

func (s *ConnectorCreatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitConnectorCreator(s)
	}
}

func (s *ConnectorCreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitConnectorCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ConnectorCreator() (localctx IConnectorCreatorContext) {
	localctx = NewConnectorCreatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, GoScriptParserRULE_connectorCreator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(458)
		p.ConnectorType()
	}
	{
		p.SetState(459)
		p.Match(GoScriptParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(460)
		p.Match(GoScriptParserT__4)
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

// IPrimitiveCreatorContext is an interface to support dynamic dispatch.
type IPrimitiveCreatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimitiveType() IPrimitiveTypeContext
	Expression() IExpressionContext

	// IsPrimitiveCreatorContext differentiates from other interfaces.
	IsPrimitiveCreatorContext()
}

type PrimitiveCreatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveCreatorContext() *PrimitiveCreatorContext {
	var p = new(PrimitiveCreatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_primitiveCreator
	return p
}

func InitEmptyPrimitiveCreatorContext(p *PrimitiveCreatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_primitiveCreator
}

func (*PrimitiveCreatorContext) IsPrimitiveCreatorContext() {}

func NewPrimitiveCreatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveCreatorContext {
	var p = new(PrimitiveCreatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_primitiveCreator

	return p
}

func (s *PrimitiveCreatorContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveCreatorContext) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *PrimitiveCreatorContext) Expression() IExpressionContext {
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

func (s *PrimitiveCreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveCreatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveCreatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterPrimitiveCreator(s)
	}
}

func (s *PrimitiveCreatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitPrimitiveCreator(s)
	}
}

func (s *PrimitiveCreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitPrimitiveCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) PrimitiveCreator() (localctx IPrimitiveCreatorContext) {
	localctx = NewPrimitiveCreatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, GoScriptParserRULE_primitiveCreator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(462)
		p.PrimitiveType()
	}
	{
		p.SetState(463)
		p.Match(GoScriptParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(465)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126118348466421768) != 0 {
		{
			p.SetState(464)
			p.expression(0)
		}

	}
	{
		p.SetState(467)
		p.Match(GoScriptParserT__4)
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

// IDynamicCreatorContext is an interface to support dynamic dispatch.
type IDynamicCreatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DynamicType() IDynamicTypeContext
	Expression() IExpressionContext

	// IsDynamicCreatorContext differentiates from other interfaces.
	IsDynamicCreatorContext()
}

type DynamicCreatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDynamicCreatorContext() *DynamicCreatorContext {
	var p = new(DynamicCreatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_dynamicCreator
	return p
}

func InitEmptyDynamicCreatorContext(p *DynamicCreatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoScriptParserRULE_dynamicCreator
}

func (*DynamicCreatorContext) IsDynamicCreatorContext() {}

func NewDynamicCreatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DynamicCreatorContext {
	var p = new(DynamicCreatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_dynamicCreator

	return p
}

func (s *DynamicCreatorContext) GetParser() antlr.Parser { return s.parser }

func (s *DynamicCreatorContext) DynamicType() IDynamicTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicTypeContext)
}

func (s *DynamicCreatorContext) Expression() IExpressionContext {
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

func (s *DynamicCreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicCreatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DynamicCreatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.EnterDynamicCreator(s)
	}
}

func (s *DynamicCreatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoScriptListener); ok {
		listenerT.ExitDynamicCreator(s)
	}
}

func (s *DynamicCreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitDynamicCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) DynamicCreator() (localctx IDynamicCreatorContext) {
	localctx = NewDynamicCreatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, GoScriptParserRULE_dynamicCreator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(469)
		p.DynamicType()
	}
	{
		p.SetState(470)
		p.Match(GoScriptParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126118348466421768) != 0 {
		{
			p.SetState(471)
			p.expression(0)
		}

	}
	{
		p.SetState(474)
		p.Match(GoScriptParserT__4)
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

func (p *GoScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 29:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 35:
		var t *LvalueContext = nil
		if localctx != nil {
			t = localctx.(*LvalueContext)
		}
		return p.Lvalue_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GoScriptParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 11)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GoScriptParser) Lvalue_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
