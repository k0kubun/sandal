//line parser.go.y:3
package parsing

import __yyfmt__ "fmt"

//line parser.go.y:3
import (
	data "github.com/k0kubun/sandal/lang/data"
	"log"
)

type token struct {
	tok int
	lit string
	pos data.Pos
}

//line parser.go.y:17
type yySymType struct {
	yys         int
	definitions []data.Definition
	definition  data.Definition
	statements  []data.Statement
	statement   data.Statement
	expressions []data.Expression
	expression  data.Expression
	parameters  []data.Parameter
	parameter   data.Parameter
	typetypes   []data.Type
	typetype    data.Type
	identifiers []string
	tags        []string
	tag         string
	blocks      []data.BlockStatement
	initvars    []data.InitVar
	initvar     data.InitVar
	ltlexpr     data.LtlExpression
	ltlatom     data.LtlAtomExpression

	tok token
}

const IDENTIFIER = 57346
const NUMBER = 57347
const COMMENT = 57348
const ADD = 57349
const SUB = 57350
const MUL = 57351
const QUO = 57352
const REM = 57353
const AND = 57354
const OR = 57355
const XOR = 57356
const SHL = 57357
const SHR = 57358
const ADD_ASSIGN = 57359
const SUB_ASSIGN = 57360
const MUL_ASSIGN = 57361
const QUO_ASSIGN = 57362
const REM_ASSIGN = 57363
const AND_ASSIGN = 57364
const OR_ASSIGN = 57365
const XOR_ASSIGN = 57366
const SHL_ASSIGN = 57367
const SHR_ASSIGN = 57368
const LAND = 57369
const LOR = 57370
const EQL = 57371
const LSS = 57372
const GTR = 57373
const ASSIGN = 57374
const NOT = 57375
const NEQ = 57376
const LEQ = 57377
const GEQ = 57378
const DATA = 57379
const CONST = 57380
const MODULE = 57381
const CHANNEL = 57382
const PROC = 57383
const FAULT = 57384
const VAR = 57385
const IF = 57386
const ELSE = 57387
const CHOICE = 57388
const RECV = 57389
const TIMEOUT_RECV = 57390
const NONBLOCK_RECV = 57391
const PEEK = 57392
const TIMEOUT_PEEK = 57393
const NONBLOCK_PEEK = 57394
const SEND = 57395
const FOR = 57396
const BREAK = 57397
const IN = 57398
const RANGE = 57399
const TO = 57400
const INIT = 57401
const GOTO = 57402
const SKIP = 57403
const TRUE = 57404
const FALSE = 57405
const LTL = 57406
const THEN = 57407
const IFF = 57408
const UNARY = 57409

var yyToknames = []string{
	"IDENTIFIER",
	"NUMBER",
	"COMMENT",
	"ADD",
	"SUB",
	"MUL",
	"QUO",
	"REM",
	"AND",
	"OR",
	"XOR",
	"SHL",
	"SHR",
	"ADD_ASSIGN",
	"SUB_ASSIGN",
	"MUL_ASSIGN",
	"QUO_ASSIGN",
	"REM_ASSIGN",
	"AND_ASSIGN",
	"OR_ASSIGN",
	"XOR_ASSIGN",
	"SHL_ASSIGN",
	"SHR_ASSIGN",
	"LAND",
	"LOR",
	"EQL",
	"LSS",
	"GTR",
	"ASSIGN",
	"NOT",
	"NEQ",
	"LEQ",
	"GEQ",
	"DATA",
	"CONST",
	"MODULE",
	"CHANNEL",
	"PROC",
	"FAULT",
	"VAR",
	"IF",
	"ELSE",
	"CHOICE",
	"RECV",
	"TIMEOUT_RECV",
	"NONBLOCK_RECV",
	"PEEK",
	"TIMEOUT_PEEK",
	"NONBLOCK_PEEK",
	"SEND",
	"FOR",
	"BREAK",
	"IN",
	"RANGE",
	"TO",
	"INIT",
	"GOTO",
	"SKIP",
	"TRUE",
	"FALSE",
	"LTL",
	"THEN",
	"IFF",
	" {",
	" }",
	" (",
	" )",
	" [",
	" ]",
	" ,",
	" :",
	" ;",
	" U",
	" V",
	" S",
	" T",
	"UNARY",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:733

type lexerWrapper struct {
	s           *Scanner
	definitions []data.Definition
	recentLit   string
	recentPos   data.Pos
}

func (l *lexerWrapper) Lex(lval *yySymType) int {
	tok, lit, pos := l.s.Scan()
	for tok == COMMENT {
		tok, lit, pos = l.s.Scan()
	}
	if tok == EOF {
		return 0
	}
	lval.tok = token{tok: tok, lit: lit, pos: pos}
	l.recentLit = lit
	l.recentPos = pos
	return tok
}

func (l *lexerWrapper) Error(e string) {
	log.Fatalf("Line %d, Column %d: %q %s",
		l.recentPos.Line, l.recentPos.Column, l.recentLit, e)
}

func Parse(s *Scanner) []data.Definition {
	l := lexerWrapper{s: s}
	if yyParse(&l) != 0 {
		panic("Parse error")
	}
	return l.definitions
}

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 139
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 998

var yyAct = []int{

	152, 197, 243, 164, 161, 158, 100, 159, 79, 99,
	34, 151, 198, 315, 329, 48, 64, 295, 45, 327,
	323, 213, 5, 7, 5, 24, 6, 66, 67, 68,
	69, 310, 309, 44, 308, 307, 293, 280, 199, 87,
	273, 257, 91, 314, 255, 252, 250, 163, 121, 109,
	104, 58, 85, 306, 154, 102, 63, 84, 192, 98,
	294, 83, 36, 81, 57, 106, 53, 90, 52, 61,
	54, 65, 11, 300, 299, 298, 254, 201, 202, 222,
	203, 204, 93, 95, 205, 94, 96, 206, 207, 208,
	120, 144, 145, 146, 209, 210, 88, 89, 35, 124,
	122, 200, 221, 92, 247, 97, 220, 64, 219, 212,
	37, 38, 39, 40, 41, 42, 43, 25, 66, 67,
	68, 69, 103, 82, 246, 245, 162, 169, 170, 171,
	172, 173, 174, 175, 176, 177, 178, 179, 180, 181,
	182, 183, 184, 185, 186, 166, 107, 168, 150, 149,
	167, 148, 147, 27, 28, 23, 328, 63, 211, 188,
	189, 190, 191, 63, 194, 195, 325, 320, 214, 218,
	61, 62, 65, 305, 216, 301, 297, 270, 65, 256,
	249, 319, 27, 226, 26, 224, 217, 166, 153, 168,
	108, 63, 167, 80, 56, 244, 196, 157, 155, 211,
	227, 211, 240, 242, 123, 223, 225, 22, 64, 279,
	51, 60, 21, 26, 64, 20, 251, 241, 59, 66,
	67, 68, 69, 47, 215, 66, 67, 68, 69, 211,
	259, 260, 261, 262, 263, 264, 265, 266, 267, 268,
	269, 258, 64, 248, 50, 211, 274, 30, 271, 211,
	278, 46, 55, 66, 67, 68, 69, 275, 276, 277,
	281, 126, 127, 128, 129, 130, 131, 132, 133, 134,
	135, 44, 32, 211, 296, 9, 11, 10, 19, 12,
	302, 136, 137, 138, 139, 140, 18, 17, 141, 142,
	143, 16, 160, 49, 31, 11, 304, 13, 12, 29,
	1, 165, 14, 15, 312, 105, 8, 4, 3, 316,
	2, 0, 211, 317, 0, 0, 13, 0, 0, 321,
	211, 322, 0, 0, 0, 211, 326, 0, 0, 313,
	126, 127, 128, 129, 130, 131, 132, 133, 134, 135,
	126, 127, 128, 129, 130, 131, 132, 133, 134, 135,
	136, 137, 138, 139, 140, 0, 0, 141, 142, 143,
	136, 137, 138, 139, 140, 0, 0, 141, 142, 143,
	0, 0, 126, 127, 128, 129, 130, 131, 132, 133,
	134, 135, 126, 127, 128, 129, 130, 131, 132, 133,
	134, 135, 136, 137, 138, 139, 140, 0, 292, 141,
	142, 143, 136, 137, 138, 139, 140, 0, 291, 141,
	142, 143, 0, 0, 126, 127, 128, 129, 130, 131,
	132, 133, 134, 135, 126, 127, 128, 129, 130, 131,
	132, 133, 134, 135, 136, 137, 138, 139, 140, 0,
	290, 141, 142, 143, 136, 137, 138, 139, 140, 0,
	289, 141, 142, 143, 0, 0, 126, 127, 128, 129,
	130, 131, 132, 133, 134, 135, 126, 127, 128, 129,
	130, 131, 132, 133, 134, 135, 136, 137, 138, 139,
	140, 0, 288, 141, 142, 143, 136, 137, 138, 139,
	140, 0, 287, 141, 142, 143, 0, 0, 126, 127,
	128, 129, 130, 131, 132, 133, 134, 135, 126, 127,
	128, 129, 130, 131, 132, 133, 134, 135, 136, 137,
	138, 139, 140, 0, 286, 141, 142, 143, 136, 137,
	138, 139, 140, 0, 285, 141, 142, 143, 0, 0,
	126, 127, 128, 129, 130, 131, 132, 133, 134, 135,
	126, 127, 128, 129, 130, 131, 132, 133, 134, 135,
	136, 137, 138, 139, 140, 0, 284, 141, 142, 143,
	136, 137, 138, 139, 140, 0, 283, 141, 142, 143,
	0, 0, 126, 127, 128, 129, 130, 131, 132, 133,
	134, 135, 126, 127, 128, 129, 130, 131, 132, 133,
	134, 135, 136, 137, 138, 139, 140, 0, 282, 141,
	142, 143, 136, 137, 138, 139, 140, 0, 253, 141,
	142, 143, 0, 0, 126, 127, 128, 129, 130, 131,
	132, 133, 134, 135, 0, 0, 0, 0, 0, 0,
	0, 63, 0, 0, 136, 137, 138, 139, 140, 0,
	125, 141, 142, 143, 61, 62, 65, 0, 193, 126,
	127, 128, 129, 130, 131, 132, 133, 134, 135, 0,
	0, 33, 0, 0, 0, 0, 0, 0, 0, 136,
	137, 138, 139, 140, 0, 0, 141, 142, 143, 156,
	0, 0, 64, 128, 129, 130, 131, 119, 0, 134,
	135, 0, 0, 66, 67, 68, 69, 70, 71, 72,
	73, 74, 75, 76, 77, 78, 0, 0, 0, 0,
	0, 0, 187, 126, 127, 128, 129, 130, 131, 132,
	133, 134, 135, 110, 111, 112, 113, 114, 115, 116,
	117, 118, 0, 136, 137, 138, 139, 140, 0, 0,
	141, 142, 143, 126, 127, 128, 129, 130, 131, 132,
	133, 134, 135, 126, 127, 128, 129, 130, 131, 132,
	133, 134, 135, 136, 137, 138, 139, 140, 0, 0,
	141, 142, 143, 324, 0, 0, 0, 126, 127, 128,
	129, 130, 131, 132, 133, 134, 135, 0, 86, 87,
	0, 0, 91, 0, 0, 0, 0, 136, 137, 138,
	139, 140, 0, 311, 141, 142, 143, 230, 231, 232,
	233, 234, 235, 236, 237, 238, 239, 90, 0, 0,
	0, 0, 229, 86, 87, 0, 0, 91, 0, 0,
	0, 0, 93, 95, 0, 94, 96, 272, 0, 0,
	0, 0, 0, 0, 0, 0, 88, 89, 0, 0,
	86, 87, 90, 92, 91, 97, 101, 0, 0, 0,
	0, 0, 0, 0, 228, 0, 0, 93, 95, 0,
	94, 96, 0, 0, 0, 0, 303, 0, 0, 90,
	0, 88, 89, 0, 0, 0, 0, 0, 92, 0,
	97, 0, 0, 0, 93, 95, 0, 94, 96, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 88, 89,
	0, 0, 0, 0, 0, 92, 0, 97, 126, 127,
	128, 129, 130, 131, 132, 133, 134, 135, 126, 127,
	128, 129, 130, 131, 132, 133, 134, 135, 136, 137,
	138, 139, 140, 0, 0, 141, 142, 143, 136, 0,
	138, 139, 140, 0, 0, 141, 142, 143, 126, 127,
	128, 129, 130, 131, 132, 133, 134, 135, 0, 318,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	138, 139, 140, 0, 0, 141, 142, 143,
}
var yyPact = []int{

	238, -1000, 238, -1000, -1000, -1000, -1000, -1000, -1000, 287,
	283, 282, 274, 148, 145, -1000, 140, 86, 113, 85,
	268, 29, 247, 240, 178, -1000, -4, -1, 240, 126,
	-1000, -9, -23, 143, -1000, 29, 29, 29, 29, 29,
	29, 29, 29, 29, -80, 125, -10, 53, -1000, -12,
	113, 856, 113, 113, 794, 52, -25, 268, 142, 122,
	-26, 29, 29, 29, 29, 29, 29, 29, 29, 29,
	627, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 267,
	-27, 247, 137, 240, -1000, 575, -1000, -1000, -1000, -1000,
	856, 856, 856, 83, 82, 80, 79, 856, -1000, 120,
	-19, 131, 617, 130, -1000, -1000, -85, 57, -28, -1000,
	149, 42, -49, -1000, 177, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 257, -1000, -1000, 856, 856, 856, 856,
	856, 856, 856, 856, 856, 856, 856, 856, 856, 856,
	856, 856, 856, 856, -1000, -1000, 652, 856, 856, 856,
	856, -14, 585, -1000, 113, 113, 129, 34, -1000, -1000,
	-85, 220, 856, -1000, 118, 257, -1000, -1000, -1000, 684,
	684, -1000, -1000, -1000, -1000, 684, 684, -1000, -1000, 961,
	931, 756, 756, 756, 756, 756, 756, -1000, 38, 36,
	32, 9, -1000, 856, -1000, 117, 113, 115, 34, 800,
	34, 213, 856, 128, 56, 55, 35, 176, -29, 212,
	-30, 543, -1000, -1000, -1000, -1000, 6, -31, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 111, -34, -1000, 34, 856,
	856, 856, 856, 856, 856, 856, 856, 856, 856, 856,
	109, 113, 780, -35, 34, 856, 856, 856, 34, 153,
	-1000, -38, -1000, -1000, -85, -1000, -1000, -1000, -1000, 533,
	501, 491, 459, 449, 417, 407, 375, 365, 333, 323,
	-39, -15, 34, -1000, 108, 5, 4, 3, 107, 829,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 856, 105, -20, -40, -41,
	-43, -44, 746, 856, 254, -32, 128, -1000, -1000, -1000,
	-1000, 34, 921, -1000, -1000, 114, -1000, 99, 856, 34,
	-55, 716, 98, -1000, 34, -56, 88, -1000, -61, -1000,
}
var yyPgo = []int{

	0, 300, 310, 308, 307, 21, 26, 23, 306, 3,
	301, 299, 247, 294, 1, 12, 0, 18, 223, 15,
	293, 11, 9, 6, 5, 7, 292, 2, 671, 10,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 2, 2, 2, 3,
	4, 9, 9, 10, 10, 10, 5, 6, 7, 8,
	8, 11, 11, 12, 12, 12, 13, 13, 14, 14,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	28, 28, 28, 28, 28, 28, 28, 28, 28, 28,
	28, 28, 28, 28, 28, 28, 28, 28, 28, 29,
	29, 17, 17, 17, 18, 18, 19, 19, 19, 20,
	21, 21, 21, 22, 22, 22, 23, 23, 23, 23,
	23, 24, 24, 25, 25, 26, 27, 27, 27,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 1, 6,
	9, 0, 2, 1, 1, 1, 6, 9, 5, 6,
	5, 0, 1, 1, 2, 3, 4, 7, 0, 2,
	3, 4, 4, 6, 6, 10, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 3, 5, 5,
	5, 5, 8, 11, 2, 3, 2, 2, 1, 1,
	1, 1, 1, 1, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 4, 4, 4, 3,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 2, 2, 2, 2, 2, 2, 1,
	3, 1, 2, 3, 0, 1, 1, 2, 3, 2,
	1, 2, 3, 1, 2, 3, 1, 3, 4, 6,
	7, 0, 1, 1, 2, 2, 3, 4, 5,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, 37,
	39, 38, 41, 59, 64, -1, 4, 4, 4, 4,
	67, 67, 67, 69, -23, 4, 71, 40, 69, -11,
	-12, -13, 4, -28, -29, 69, 33, 81, 82, 83,
	84, 85, 86, 87, 4, -17, 4, -18, -19, -20,
	4, 32, 72, 67, 71, -18, 68, 73, 74, 75,
	68, 27, 28, 14, 65, 29, 76, 77, 78, 79,
	-28, -28, -28, -28, -28, -28, -28, -28, -28, 88,
	68, 73, 70, 73, -23, -16, 4, 5, 62, 63,
	33, 8, 69, 48, 51, 49, 52, 71, -23, -22,
	-23, 72, -16, 70, 75, -12, -23, 4, 68, 75,
	-28, -28, -28, -28, -28, -28, -28, -28, -28, 70,
	-29, 75, -17, 67, -19, 75, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 27, 28, 29, 30,
	31, 34, 35, 36, -16, -16, -16, 69, 69, 69,
	69, -21, -16, 68, 73, 67, 72, 67, -24, -25,
	-26, 89, 69, 75, -9, -10, -5, -6, -7, -16,
	-16, -16, -16, -16, -16, -16, -16, -16, -16, -16,
	-16, -16, -16, -16, -16, -16, -16, 70, -21, -21,
	-21, -21, 72, 73, -22, -22, 67, -14, -15, 4,
	67, 43, 44, 46, 47, 50, 53, 54, 55, 60,
	61, -16, 75, -5, -25, 4, -21, 68, -9, 70,
	70, 70, 70, -21, 68, -22, 68, -14, 74, 32,
	17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
	-14, 4, -16, -27, 67, 69, 69, 69, 67, 4,
	75, 4, 75, 75, 70, 75, 68, 75, -15, -16,
	-16, -16, -16, -16, -16, -16, -16, -16, -16, -16,
	68, -23, 67, 75, -14, -21, -21, -21, -14, 56,
	75, -24, 75, 75, 75, 75, 75, 75, 75, 75,
	75, 75, 75, 75, 75, 32, -14, 68, 70, 70,
	70, 68, -16, 57, -16, 68, 73, 75, 75, 75,
	75, 67, -16, 75, 75, 45, -27, -14, 58, 67,
	68, -16, -14, 75, 67, 68, -14, 75, 68, 75,
}
var yyDef = []int{

	0, -2, 1, 3, 4, 5, 6, 7, 8, 0,
	0, 0, 0, 0, 0, 2, 0, 0, 0, 0,
	21, 0, 0, 114, 0, 126, 0, 0, 114, 0,
	22, 23, 0, 0, 90, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 109, 0, 111, 0, 115, 116,
	0, 0, 0, 0, 0, 0, 0, 24, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 101, 102, 103, 104, 105, 106, 107, 108, 0,
	0, 112, 0, 117, 119, 0, 60, 61, 62, 63,
	0, 0, 0, 0, 0, 0, 0, 0, 127, 0,
	123, 0, 0, 0, 18, 25, 131, 126, 0, 20,
	92, 93, 94, 95, 96, 97, 98, 99, 100, 91,
	110, 9, 113, 11, 118, 16, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 65, 0, 0, 0, 0,
	0, 0, 120, 128, 124, 0, 0, 28, 26, 132,
	133, 0, 0, 19, 0, 11, 13, 14, 15, 67,
	68, 69, 70, 71, 72, 73, 74, 75, 76, 77,
	78, 79, 80, 81, 82, 83, 84, 66, 0, 0,
	0, 0, 89, 121, 125, 0, 0, 0, 28, 60,
	28, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 58, 59, 134, 135, 0, 0, 12, 85,
	86, 87, 88, 122, 129, 0, 0, 29, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 28, 0, 0, 0, 28, 0,
	54, 0, 56, 57, 131, 10, 130, 17, 30, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 28, 47, 0, 0, 0, 0, 0, 0,
	55, 27, 36, 37, 38, 39, 40, 41, 42, 43,
	44, 45, 46, 31, 32, 0, 0, 136, 0, 0,
	0, 0, 0, 0, 0, 0, 137, 48, 49, 50,
	51, 28, 0, 33, 34, 0, 138, 0, 0, 28,
	0, 0, 0, 52, 28, 0, 0, 35, 0, 53,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	69, 70, 3, 3, 73, 3, 88, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 74, 75,
	3, 3, 3, 3, 89, 3, 3, 3, 3, 3,
	83, 82, 86, 3, 3, 3, 3, 3, 3, 87,
	3, 3, 3, 78, 79, 76, 77, 3, 81, 84,
	85, 71, 3, 72, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 67, 3, 68,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 80,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line parser.go.y:147
		{
			yyVAL.definitions = []data.Definition{yyS[yypt-0].definition}
			if l, isLexerWrapper := yylex.(*lexerWrapper); isLexerWrapper {
				l.definitions = yyVAL.definitions
			}
		}
	case 2:
		//line parser.go.y:154
		{
			yyVAL.definitions = append([]data.Definition{yyS[yypt-1].definition}, yyS[yypt-0].definitions...)
			if l, isLexerWrapper := yylex.(*lexerWrapper); isLexerWrapper {
				l.definitions = yyVAL.definitions
			}
		}
	case 3:
		yyVAL.definition = yyS[yypt-0].definition
	case 4:
		yyVAL.definition = yyS[yypt-0].definition
	case 5:
		yyVAL.definition = yyS[yypt-0].definition
	case 6:
		yyVAL.definition = yyS[yypt-0].definition
	case 7:
		yyVAL.definition = yyS[yypt-0].definition
	case 8:
		yyVAL.definition = yyS[yypt-0].definition
	case 9:
		//line parser.go.y:171
		{
			yyVAL.definition = data.DataDefinition{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-4].tok.lit, Elems: yyS[yypt-2].identifiers}
		}
	case 10:
		//line parser.go.y:177
		{
			yyVAL.definition = data.ModuleDefinition{Pos: yyS[yypt-8].tok.pos, Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Definitions: yyS[yypt-2].definitions}
		}
	case 11:
		//line parser.go.y:183
		{
			yyVAL.definitions = nil
		}
	case 12:
		//line parser.go.y:187
		{
			yyVAL.definitions = append([]data.Definition{yyS[yypt-1].definition}, yyS[yypt-0].definitions...)
		}
	case 13:
		yyVAL.definition = yyS[yypt-0].definition
	case 14:
		yyVAL.definition = yyS[yypt-0].definition
	case 15:
		yyVAL.definition = yyS[yypt-0].definition
	case 16:
		//line parser.go.y:198
		{
			yyVAL.definition = data.ConstantDefinition{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-4].tok.lit, Type: yyS[yypt-3].typetype, Expr: yyS[yypt-1].expression}
		}
	case 17:
		//line parser.go.y:204
		{
			yyVAL.definition = data.ProcDefinition{Pos: yyS[yypt-8].tok.pos, Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Statements: yyS[yypt-2].statements}
		}
	case 18:
		//line parser.go.y:210
		{
			yyVAL.definition = data.InitBlock{Pos: yyS[yypt-4].tok.pos, Vars: yyS[yypt-2].initvars}
		}
	case 19:
		//line parser.go.y:216
		{
			yyVAL.definition = data.LtlSpec{Expr: yyS[yypt-3].ltlexpr}
		}
	case 20:
		//line parser.go.y:220
		{
			yyVAL.definition = data.LtlSpec{Expr: yyS[yypt-2].ltlexpr}
		}
	case 21:
		//line parser.go.y:226
		{
			yyVAL.initvars = nil
		}
	case 22:
		//line parser.go.y:230
		{
			yyVAL.initvars = yyS[yypt-0].initvars
		}
	case 23:
		//line parser.go.y:236
		{
			yyVAL.initvars = []data.InitVar{yyS[yypt-0].initvar}
		}
	case 24:
		//line parser.go.y:240
		{
			yyVAL.initvars = []data.InitVar{yyS[yypt-1].initvar}
		}
	case 25:
		//line parser.go.y:244
		{
			yyVAL.initvars = append([]data.InitVar{yyS[yypt-2].initvar}, yyS[yypt-0].initvars...)
		}
	case 26:
		//line parser.go.y:249
		{
			yyVAL.initvar = data.ChannelVar{Pos: yyS[yypt-3].tok.pos, Name: yyS[yypt-3].tok.lit, Type: yyS[yypt-1].typetype, Tags: yyS[yypt-0].tags}
		}
	case 27:
		//line parser.go.y:253
		{
			yyVAL.initvar = data.InstanceVar{Pos: yyS[yypt-6].tok.pos, Name: yyS[yypt-6].tok.lit, ProcDefName: yyS[yypt-4].tok.lit, Args: yyS[yypt-2].expressions, Tags: yyS[yypt-0].tags}
		}
	case 28:
		//line parser.go.y:259
		{
			yyVAL.statements = nil
		}
	case 29:
		//line parser.go.y:263
		{
			yyVAL.statements = append([]data.Statement{yyS[yypt-1].statement}, yyS[yypt-0].statements...)
		}
	case 30:
		//line parser.go.y:269
		{
			yyVAL.statement = data.LabelledStatement{Pos: yyS[yypt-2].tok.pos, Label: yyS[yypt-2].tok.lit, Statement: yyS[yypt-0].statement}
		}
	case 31:
		//line parser.go.y:273
		{
			yyVAL.statement = data.BlockStatement{Pos: yyS[yypt-3].tok.pos, Statements: yyS[yypt-2].statements}
		}
	case 32:
		//line parser.go.y:277
		{
			yyVAL.statement = data.VarDeclStatement{Pos: yyS[yypt-3].tok.pos, Name: yyS[yypt-2].tok.lit, Type: yyS[yypt-1].typetype}
		}
	case 33:
		//line parser.go.y:281
		{
			yyVAL.statement = data.VarDeclStatement{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-4].tok.lit, Type: yyS[yypt-3].typetype, Initializer: yyS[yypt-1].expression}
		}
	case 34:
		//line parser.go.y:285
		{
			yyVAL.statement = data.IfStatement{Pos: yyS[yypt-5].tok.pos, Condition: yyS[yypt-4].expression, TrueBranch: yyS[yypt-2].statements}
		}
	case 35:
		//line parser.go.y:289
		{
			yyVAL.statement = data.IfStatement{Pos: yyS[yypt-9].tok.pos, Condition: yyS[yypt-8].expression, TrueBranch: yyS[yypt-6].statements, FalseBranch: yyS[yypt-2].statements}
		}
	case 36:
		//line parser.go.y:293
		{
			yyVAL.statement = data.AssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Expr: yyS[yypt-1].expression}
		}
	case 37:
		//line parser.go.y:297
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "+", Expr: yyS[yypt-1].expression}
		}
	case 38:
		//line parser.go.y:301
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "-", Expr: yyS[yypt-1].expression}
		}
	case 39:
		//line parser.go.y:305
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "*", Expr: yyS[yypt-1].expression}
		}
	case 40:
		//line parser.go.y:309
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "/", Expr: yyS[yypt-1].expression}
		}
	case 41:
		//line parser.go.y:313
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "%", Expr: yyS[yypt-1].expression}
		}
	case 42:
		//line parser.go.y:317
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "&", Expr: yyS[yypt-1].expression}
		}
	case 43:
		//line parser.go.y:321
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "|", Expr: yyS[yypt-1].expression}
		}
	case 44:
		//line parser.go.y:325
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "^", Expr: yyS[yypt-1].expression}
		}
	case 45:
		//line parser.go.y:329
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "<<", Expr: yyS[yypt-1].expression}
		}
	case 46:
		//line parser.go.y:333
		{
			yyVAL.statement = data.OpAssignmentStatement{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: ">>", Expr: yyS[yypt-1].expression}
		}
	case 47:
		//line parser.go.y:337
		{
			yyVAL.statement = data.ChoiceStatement{Pos: yyS[yypt-2].tok.pos, Blocks: yyS[yypt-1].blocks}
		}
	case 48:
		//line parser.go.y:341
		{
			yyVAL.statement = data.RecvStatement{Pos: yyS[yypt-4].tok.pos, Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 49:
		//line parser.go.y:345
		{
			yyVAL.statement = data.PeekStatement{Pos: yyS[yypt-4].tok.pos, Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 50:
		//line parser.go.y:349
		{
			yyVAL.statement = data.SendStatement{Pos: yyS[yypt-4].tok.pos, Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 51:
		//line parser.go.y:353
		{
			yyVAL.statement = data.ForStatement{Pos: yyS[yypt-4].tok.pos, Statements: yyS[yypt-2].statements}
		}
	case 52:
		//line parser.go.y:357
		{
			yyVAL.statement = data.ForInStatement{Pos: yyS[yypt-7].tok.pos, Variable: yyS[yypt-6].tok.lit, Container: yyS[yypt-4].expression, Statements: yyS[yypt-2].statements}
		}
	case 53:
		//line parser.go.y:361
		{
			yyVAL.statement = data.ForInRangeStatement{Pos: yyS[yypt-10].tok.pos, Variable: yyS[yypt-9].tok.lit, FromExpr: yyS[yypt-6].expression, ToExpr: yyS[yypt-4].expression, Statements: yyS[yypt-2].statements}
		}
	case 54:
		//line parser.go.y:365
		{
			yyVAL.statement = data.BreakStatement{Pos: yyS[yypt-1].tok.pos}
		}
	case 55:
		//line parser.go.y:369
		{
			yyVAL.statement = data.GotoStatement{Pos: yyS[yypt-2].tok.pos, Label: yyS[yypt-1].tok.lit}
		}
	case 56:
		//line parser.go.y:373
		{
			yyVAL.statement = data.SkipStatement{Pos: yyS[yypt-1].tok.pos}
		}
	case 57:
		//line parser.go.y:377
		{
			yyVAL.statement = data.ExprStatement{Expr: yyS[yypt-1].expression}
		}
	case 58:
		//line parser.go.y:381
		{
			yyVAL.statement = data.NullStatement{Pos: yyS[yypt-0].tok.pos}
		}
	case 59:
		//line parser.go.y:385
		{
			yyVAL.statement = yyS[yypt-0].definition.(data.Statement)
		}
	case 60:
		//line parser.go.y:390
		{
			yyVAL.expression = data.IdentifierExpression{Pos: yyS[yypt-0].tok.pos, Name: yyS[yypt-0].tok.lit}
		}
	case 61:
		//line parser.go.y:394
		{
			yyVAL.expression = data.NumberExpression{Pos: yyS[yypt-0].tok.pos, Lit: yyS[yypt-0].tok.lit}
		}
	case 62:
		//line parser.go.y:398
		{
			yyVAL.expression = data.TrueExpression{Pos: yyS[yypt-0].tok.pos}
		}
	case 63:
		//line parser.go.y:402
		{
			yyVAL.expression = data.FalseExpression{Pos: yyS[yypt-0].tok.pos}
		}
	case 64:
		//line parser.go.y:406
		{
			yyVAL.expression = data.NotExpression{Pos: yyS[yypt-1].tok.pos, SubExpr: yyS[yypt-0].expression}
		}
	case 65:
		//line parser.go.y:410
		{
			yyVAL.expression = data.UnarySubExpression{Pos: yyS[yypt-1].tok.pos, SubExpr: yyS[yypt-0].expression}
		}
	case 66:
		//line parser.go.y:414
		{
			yyVAL.expression = data.ParenExpression{Pos: yyS[yypt-2].tok.pos, SubExpr: yyS[yypt-1].expression}
		}
	case 67:
		//line parser.go.y:418
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "+", RHS: yyS[yypt-0].expression}
		}
	case 68:
		//line parser.go.y:422
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "-", RHS: yyS[yypt-0].expression}
		}
	case 69:
		//line parser.go.y:426
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "*", RHS: yyS[yypt-0].expression}
		}
	case 70:
		//line parser.go.y:430
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "/", RHS: yyS[yypt-0].expression}
		}
	case 71:
		//line parser.go.y:434
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "%", RHS: yyS[yypt-0].expression}
		}
	case 72:
		//line parser.go.y:438
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "&", RHS: yyS[yypt-0].expression}
		}
	case 73:
		//line parser.go.y:442
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "|", RHS: yyS[yypt-0].expression}
		}
	case 74:
		//line parser.go.y:446
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "^", RHS: yyS[yypt-0].expression}
		}
	case 75:
		//line parser.go.y:450
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "<<", RHS: yyS[yypt-0].expression}
		}
	case 76:
		//line parser.go.y:454
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: ">>", RHS: yyS[yypt-0].expression}
		}
	case 77:
		//line parser.go.y:458
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "&&", RHS: yyS[yypt-0].expression}
		}
	case 78:
		//line parser.go.y:462
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "||", RHS: yyS[yypt-0].expression}
		}
	case 79:
		//line parser.go.y:466
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "==", RHS: yyS[yypt-0].expression}
		}
	case 80:
		//line parser.go.y:470
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "<", RHS: yyS[yypt-0].expression}
		}
	case 81:
		//line parser.go.y:474
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: ">", RHS: yyS[yypt-0].expression}
		}
	case 82:
		//line parser.go.y:478
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "!=", RHS: yyS[yypt-0].expression}
		}
	case 83:
		//line parser.go.y:482
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: "<=", RHS: yyS[yypt-0].expression}
		}
	case 84:
		//line parser.go.y:486
		{
			yyVAL.expression = data.BinOpExpression{LHS: yyS[yypt-2].expression, Operator: ">=", RHS: yyS[yypt-0].expression}
		}
	case 85:
		//line parser.go.y:490
		{
			yyVAL.expression = data.TimeoutRecvExpression{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 86:
		//line parser.go.y:494
		{
			yyVAL.expression = data.TimeoutPeekExpression{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 87:
		//line parser.go.y:498
		{
			yyVAL.expression = data.NonblockRecvExpression{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 88:
		//line parser.go.y:502
		{
			yyVAL.expression = data.NonblockPeekExpression{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 89:
		//line parser.go.y:506
		{
			yyVAL.expression = data.ArrayExpression{Pos: yyS[yypt-2].tok.pos, Elems: yyS[yypt-1].expressions}
		}
	case 90:
		//line parser.go.y:513
		{
			yyVAL.ltlexpr = yyS[yypt-0].ltlatom
		}
	case 91:
		//line parser.go.y:517
		{
			yyVAL.ltlexpr = data.ParenLtlExpression{SubExpr: yyS[yypt-1].ltlexpr}
		}
	case 92:
		//line parser.go.y:521
		{
			yyVAL.ltlexpr = data.BinOpLtlExpression{Operator: "&", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 93:
		//line parser.go.y:525
		{
			yyVAL.ltlexpr = data.BinOpLtlExpression{Operator: "|", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 94:
		//line parser.go.y:529
		{
			yyVAL.ltlexpr = data.BinOpLtlExpression{Operator: "^", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 95:
		//line parser.go.y:533
		{
			yyVAL.ltlexpr = data.BinOpLtlExpression{Operator: "->", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 96:
		//line parser.go.y:537
		{
			yyVAL.ltlexpr = data.BinOpLtlExpression{Operator: "=", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 97:
		//line parser.go.y:541
		{
			yyVAL.ltlexpr = data.BinOpLtlExpression{Operator: "U", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 98:
		//line parser.go.y:545
		{
			yyVAL.ltlexpr = data.BinOpLtlExpression{Operator: "V", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 99:
		//line parser.go.y:549
		{
			yyVAL.ltlexpr = data.BinOpLtlExpression{Operator: "S", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 100:
		//line parser.go.y:553
		{
			yyVAL.ltlexpr = data.BinOpLtlExpression{Operator: "T", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 101:
		//line parser.go.y:557
		{
			yyVAL.ltlexpr = data.UnOpLtlExpression{Operator: "!", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 102:
		//line parser.go.y:561
		{
			yyVAL.ltlexpr = data.UnOpLtlExpression{Operator: "X", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 103:
		//line parser.go.y:565
		{
			yyVAL.ltlexpr = data.UnOpLtlExpression{Operator: "G", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 104:
		//line parser.go.y:569
		{
			yyVAL.ltlexpr = data.UnOpLtlExpression{Operator: "F", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 105:
		//line parser.go.y:573
		{
			yyVAL.ltlexpr = data.UnOpLtlExpression{Operator: "Y", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 106:
		//line parser.go.y:577
		{
			yyVAL.ltlexpr = data.UnOpLtlExpression{Operator: "Z", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 107:
		//line parser.go.y:581
		{
			yyVAL.ltlexpr = data.UnOpLtlExpression{Operator: "H", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 108:
		//line parser.go.y:585
		{
			yyVAL.ltlexpr = data.UnOpLtlExpression{Operator: "O", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 109:
		//line parser.go.y:590
		{
			yyVAL.ltlatom = data.LtlAtomExpression{Names: []string{yyS[yypt-0].tok.lit}}
		}
	case 110:
		//line parser.go.y:594
		{
			yyVAL.ltlatom = data.LtlAtomExpression{Names: append([]string{yyS[yypt-2].tok.lit}, yyS[yypt-0].ltlatom.Names...)}
		}
	case 111:
		//line parser.go.y:602
		{
			yyVAL.identifiers = []string{yyS[yypt-0].tok.lit}
		}
	case 112:
		//line parser.go.y:606
		{
			yyVAL.identifiers = []string{yyS[yypt-1].tok.lit}
		}
	case 113:
		//line parser.go.y:610
		{
			yyVAL.identifiers = append([]string{yyS[yypt-2].tok.lit}, yyS[yypt-0].identifiers...)
		}
	case 114:
		//line parser.go.y:616
		{
			yyVAL.parameters = nil
		}
	case 115:
		//line parser.go.y:620
		{
			yyVAL.parameters = yyS[yypt-0].parameters
		}
	case 116:
		//line parser.go.y:626
		{
			yyVAL.parameters = []data.Parameter{yyS[yypt-0].parameter}
		}
	case 117:
		//line parser.go.y:630
		{
			yyVAL.parameters = []data.Parameter{yyS[yypt-1].parameter}
		}
	case 118:
		//line parser.go.y:634
		{
			yyVAL.parameters = append([]data.Parameter{yyS[yypt-2].parameter}, yyS[yypt-0].parameters...)
		}
	case 119:
		//line parser.go.y:640
		{
			yyVAL.parameter = data.Parameter{Name: yyS[yypt-1].tok.lit, Type: yyS[yypt-0].typetype}
		}
	case 120:
		//line parser.go.y:646
		{
			yyVAL.expressions = []data.Expression{yyS[yypt-0].expression}
		}
	case 121:
		//line parser.go.y:650
		{
			yyVAL.expressions = []data.Expression{yyS[yypt-1].expression}
		}
	case 122:
		//line parser.go.y:654
		{
			yyVAL.expressions = append([]data.Expression{yyS[yypt-2].expression}, yyS[yypt-0].expressions...)
		}
	case 123:
		//line parser.go.y:660
		{
			yyVAL.typetypes = []data.Type{yyS[yypt-0].typetype}
		}
	case 124:
		//line parser.go.y:664
		{
			yyVAL.typetypes = []data.Type{yyS[yypt-1].typetype}
		}
	case 125:
		//line parser.go.y:668
		{
			yyVAL.typetypes = append([]data.Type{yyS[yypt-2].typetype}, yyS[yypt-0].typetypes...)
		}
	case 126:
		//line parser.go.y:673
		{
			yyVAL.typetype = data.NamedType{Name: yyS[yypt-0].tok.lit}
		}
	case 127:
		//line parser.go.y:677
		{
			yyVAL.typetype = data.ArrayType{ElemType: yyS[yypt-0].typetype}
		}
	case 128:
		//line parser.go.y:681
		{
			yyVAL.typetype = data.HandshakeChannelType{Elems: yyS[yypt-1].typetypes}
		}
	case 129:
		//line parser.go.y:685
		{
			yyVAL.typetype = data.BufferedChannelType{BufferSize: nil, Elems: yyS[yypt-1].typetypes}
		}
	case 130:
		//line parser.go.y:689
		{
			yyVAL.typetype = data.BufferedChannelType{BufferSize: yyS[yypt-4].expression, Elems: yyS[yypt-1].typetypes}
		}
	case 131:
		//line parser.go.y:695
		{
			yyVAL.tags = nil
		}
	case 132:
		//line parser.go.y:699
		{
			yyVAL.tags = yyS[yypt-0].tags
		}
	case 133:
		//line parser.go.y:705
		{
			yyVAL.tags = []string{yyS[yypt-0].tag}
		}
	case 134:
		//line parser.go.y:709
		{
			yyVAL.tags = append([]string{yyS[yypt-1].tag}, yyS[yypt-0].tags...)
		}
	case 135:
		//line parser.go.y:715
		{
			yyVAL.tag = yyS[yypt-0].tok.lit
		}
	case 136:
		//line parser.go.y:721
		{
			yyVAL.blocks = []data.BlockStatement{data.BlockStatement{Pos: yyS[yypt-2].tok.pos, Statements: yyS[yypt-1].statements}}
		}
	case 137:
		//line parser.go.y:725
		{
			yyVAL.blocks = []data.BlockStatement{data.BlockStatement{Pos: yyS[yypt-3].tok.pos, Statements: yyS[yypt-2].statements}}
		}
	case 138:
		//line parser.go.y:729
		{
			yyVAL.blocks = append([]data.BlockStatement{data.BlockStatement{Pos: yyS[yypt-4].tok.pos, Statements: yyS[yypt-3].statements}}, yyS[yypt-0].blocks...)
		}
	}
	goto yystack /* stack new state and value */
}
