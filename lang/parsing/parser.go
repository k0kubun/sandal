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
	definitions []data.Def
	definition  data.Def
	statements  []data.Stmt
	statement   data.Stmt
	expressions []data.Expr
	expression  data.Expr
	parameters  []data.Parameter
	parameter   data.Parameter
	typetypes   []data.Type
	typetype    data.Type
	identifiers []string
	tags        []string
	tag         string
	blocks      []data.BlockStmt
	initvars    []data.InitVar
	initvar     data.InitVar
	ltlexpr     data.LtlExpr
	ltlatom     data.LtlAtomExpr

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

//line parser.go.y:744

type lexerWrapper struct {
	s           *Scanner
	definitions []data.Def
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

func Parse(s *Scanner) []data.Def {
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

const yyNprod = 142
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1041

var yyAct = []int{

	223, 257, 171, 108, 107, 209, 176, 172, 173, 169,
	87, 335, 210, 313, 351, 349, 345, 51, 161, 338,
	54, 8, 225, 5, 28, 5, 337, 40, 330, 6,
	244, 245, 246, 247, 248, 249, 250, 251, 252, 253,
	328, 334, 323, 322, 311, 243, 211, 95, 296, 50,
	99, 289, 72, 273, 62, 63, 312, 271, 93, 266,
	92, 110, 106, 74, 75, 76, 77, 264, 175, 131,
	116, 119, 114, 66, 326, 98, 164, 91, 42, 89,
	12, 65, 204, 58, 318, 213, 214, 242, 215, 216,
	101, 103, 217, 102, 104, 218, 219, 220, 29, 154,
	155, 156, 221, 222, 96, 97, 162, 132, 317, 212,
	316, 100, 134, 105, 41, 130, 59, 224, 270, 261,
	60, 168, 170, 236, 235, 234, 43, 44, 45, 46,
	47, 48, 49, 233, 31, 117, 113, 181, 182, 183,
	184, 185, 186, 187, 188, 189, 190, 191, 192, 193,
	194, 195, 196, 197, 198, 180, 178, 71, 162, 162,
	162, 162, 112, 179, 260, 30, 111, 90, 259, 206,
	207, 31, 73, 174, 71, 162, 200, 201, 202, 203,
	160, 229, 159, 158, 232, 157, 34, 69, 70, 73,
	33, 32, 27, 230, 350, 347, 342, 325, 319, 180,
	178, 315, 30, 298, 297, 286, 162, 179, 72, 272,
	240, 341, 71, 239, 238, 256, 241, 231, 254, 74,
	75, 76, 77, 163, 237, 72, 118, 88, 68, 263,
	64, 258, 268, 228, 269, 67, 74, 75, 76, 77,
	226, 208, 167, 165, 275, 276, 277, 278, 279, 280,
	281, 282, 283, 284, 285, 274, 133, 26, 25, 287,
	162, 162, 162, 72, 290, 36, 24, 295, 294, 57,
	265, 255, 227, 299, 74, 75, 76, 77, 291, 292,
	293, 136, 137, 138, 139, 140, 141, 142, 143, 144,
	145, 23, 262, 53, 314, 12, 320, 22, 13, 56,
	71, 146, 147, 148, 149, 150, 52, 50, 151, 152,
	153, 39, 38, 69, 324, 73, 15, 21, 20, 327,
	19, 329, 332, 18, 1, 55, 61, 17, 336, 37,
	35, 115, 138, 139, 140, 141, 177, 339, 144, 145,
	9, 343, 7, 4, 3, 2, 0, 344, 0, 333,
	0, 72, 348, 78, 79, 80, 81, 82, 83, 84,
	85, 86, 74, 75, 76, 77, 0, 0, 0, 0,
	136, 137, 138, 139, 140, 141, 142, 143, 144, 145,
	0, 120, 121, 122, 123, 124, 125, 126, 127, 128,
	146, 147, 148, 149, 150, 0, 0, 151, 152, 153,
	136, 137, 138, 139, 140, 141, 142, 143, 144, 145,
	10, 12, 11, 0, 13, 14, 0, 0, 0, 0,
	146, 147, 148, 149, 150, 0, 0, 151, 152, 153,
	0, 0, 15, 0, 0, 0, 0, 16, 310, 0,
	0, 0, 136, 137, 138, 139, 140, 141, 142, 143,
	144, 145, 136, 137, 138, 139, 140, 141, 142, 143,
	144, 145, 146, 147, 148, 149, 150, 0, 309, 151,
	152, 153, 146, 147, 148, 149, 150, 0, 0, 151,
	152, 153, 0, 0, 136, 137, 138, 139, 140, 141,
	142, 143, 144, 145, 136, 137, 138, 139, 140, 141,
	142, 143, 144, 145, 146, 147, 148, 149, 150, 0,
	308, 151, 152, 153, 146, 147, 148, 149, 150, 0,
	307, 151, 152, 153, 0, 0, 136, 137, 138, 139,
	140, 141, 142, 143, 144, 145, 136, 137, 138, 139,
	140, 141, 142, 143, 144, 145, 146, 147, 148, 149,
	150, 0, 306, 151, 152, 153, 146, 147, 148, 149,
	150, 0, 305, 151, 152, 153, 0, 0, 136, 137,
	138, 139, 140, 141, 142, 143, 144, 145, 136, 137,
	138, 139, 140, 141, 142, 143, 144, 145, 146, 147,
	148, 149, 150, 0, 304, 151, 152, 153, 146, 147,
	148, 149, 150, 0, 303, 151, 152, 153, 0, 0,
	136, 137, 138, 139, 140, 141, 142, 143, 144, 145,
	136, 137, 138, 139, 140, 141, 142, 143, 144, 145,
	146, 147, 148, 149, 150, 0, 302, 151, 152, 153,
	146, 147, 148, 149, 150, 0, 301, 151, 152, 153,
	0, 0, 136, 137, 138, 139, 140, 141, 142, 143,
	144, 145, 136, 137, 138, 139, 140, 141, 142, 143,
	144, 145, 146, 147, 148, 149, 150, 0, 300, 151,
	152, 153, 146, 147, 148, 149, 150, 0, 267, 151,
	152, 153, 0, 0, 136, 137, 138, 139, 140, 141,
	142, 143, 144, 145, 0, 0, 0, 0, 0, 0,
	0, 71, 0, 0, 146, 147, 148, 149, 150, 0,
	135, 151, 152, 153, 69, 70, 73, 0, 205, 136,
	137, 138, 139, 140, 141, 142, 143, 144, 145, 136,
	137, 138, 139, 140, 141, 142, 143, 144, 145, 146,
	147, 148, 149, 150, 0, 0, 151, 152, 153, 166,
	0, 0, 72, 0, 0, 0, 0, 129, 0, 0,
	0, 0, 0, 74, 75, 76, 77, 136, 137, 138,
	139, 140, 141, 142, 143, 144, 145, 0, 0, 0,
	0, 0, 199, 0, 0, 0, 0, 146, 147, 148,
	149, 150, 0, 0, 151, 152, 153, 136, 137, 138,
	139, 140, 141, 142, 143, 144, 145, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 146, 147, 148,
	149, 150, 0, 0, 151, 152, 153, 346, 0, 0,
	0, 136, 137, 138, 139, 140, 141, 142, 143, 144,
	145, 94, 95, 0, 0, 99, 0, 0, 0, 0,
	0, 146, 147, 148, 149, 150, 0, 331, 151, 152,
	153, 0, 0, 0, 0, 0, 0, 94, 95, 0,
	98, 99, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 101, 103, 0, 102, 104,
	0, 288, 0, 94, 95, 0, 98, 99, 0, 96,
	97, 0, 0, 0, 0, 0, 100, 0, 105, 109,
	0, 101, 103, 0, 102, 104, 0, 0, 0, 0,
	321, 0, 98, 0, 0, 96, 97, 0, 0, 0,
	0, 0, 100, 0, 105, 0, 0, 101, 103, 0,
	102, 104, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 96, 97, 0, 0, 0, 0, 0, 100, 0,
	105, 136, 137, 138, 139, 140, 141, 142, 143, 144,
	145, 136, 137, 138, 139, 140, 141, 142, 143, 144,
	145, 146, 147, 148, 149, 150, 0, 0, 151, 152,
	153, 146, 0, 148, 149, 150, 0, 0, 151, 152,
	153, 136, 137, 138, 139, 140, 141, 142, 143, 144,
	145, 0, 340, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 148, 149, 150, 0, 0, 151, 152,
	153,
}
var yyPact = []int{

	373, -1000, 373, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	319, 316, 314, 313, 244, 199, 191, -1000, 190, 123,
	94, 122, 121, 117, 308, 45, 302, 295, 237, -1000,
	11, 49, 295, 295, 295, 162, -1000, 8, -1, 160,
	-1000, 45, 45, 45, 45, 45, 45, 45, 45, 45,
	-78, 159, 6, 97, -1000, 4, 94, 899, 94, 94,
	847, 96, 92, 66, -3, 308, 131, 158, -4, 45,
	45, 45, 45, 45, 45, 45, 45, 45, 697, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 303, -6, 302,
	189, 295, -1000, 645, -1000, -1000, -1000, -1000, 899, 899,
	899, 116, 114, 113, 111, 899, -1000, 155, 3, 176,
	687, 175, -80, -80, -1000, -1000, -80, 104, -7, -1000,
	143, 286, -13, -1000, 198, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 257, -1000, -1000, 899, 899, 899, 899,
	899, 899, 899, 899, 899, 899, 899, 899, 899, 899,
	899, 899, 899, 899, -1000, -1000, 722, 899, 899, 899,
	899, 10, 655, -1000, 94, 94, 174, 42, 173, 268,
	166, -1000, -1000, -80, 899, -1000, 149, 257, -1000, -1000,
	-1000, 323, 323, -1000, -1000, -1000, -1000, 323, 323, -1000,
	-1000, 1004, 974, 732, 732, 732, 732, 732, 732, -1000,
	63, 55, 54, 53, -1000, 899, -1000, 146, 94, 142,
	42, 13, 42, 267, 899, 164, 99, 95, 50, 225,
	-8, 266, -16, 613, -1000, -1000, 42, -1000, 42, -1000,
	48, -18, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 141,
	-22, -1000, 42, 899, 899, 899, 899, 899, 899, 899,
	899, 899, 899, 899, 137, 94, 834, -24, 42, 899,
	899, 899, 42, 211, -1000, -27, -1000, -1000, 136, 135,
	-80, -1000, -1000, -1000, -1000, 603, 571, 561, 529, 519,
	487, 477, 445, 435, 393, 363, -31, -19, 42, -1000,
	133, 40, 38, 14, 130, 873, -1000, -32, -33, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 899, 129, 1, -80, -35, -80, -47,
	800, 899, -1000, -1000, 274, -34, 164, -49, -1000, -56,
	-1000, 42, 964, -1000, -1000, 144, -1000, -1000, -1000, 128,
	899, 42, -59, 770, 127, -1000, 42, -60, 126, -1000,
	-61, -1000,
}
var yyPgo = []int{

	0, 324, 345, 344, 343, 22, 29, 342, 21, 340,
	6, 336, 330, 265, 329, 5, 12, 0, 17, 293,
	20, 325, 18, 4, 3, 2, 7, 8, 1, 311,
	27,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 2, 2, 2, 2,
	3, 4, 10, 10, 11, 11, 11, 5, 6, 7,
	7, 8, 9, 9, 12, 12, 13, 13, 13, 14,
	14, 15, 15, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 29, 29, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 29, 29, 29,
	29, 29, 30, 30, 18, 18, 18, 19, 19, 20,
	20, 20, 21, 22, 22, 22, 23, 23, 23, 24,
	24, 24, 24, 24, 25, 25, 26, 26, 27, 28,
	28, 28,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	6, 9, 0, 2, 1, 1, 1, 6, 9, 10,
	10, 5, 6, 5, 0, 1, 1, 2, 3, 4,
	7, 0, 2, 3, 4, 4, 6, 6, 10, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	3, 6, 5, 6, 5, 8, 11, 2, 3, 2,
	2, 1, 1, 1, 1, 1, 1, 2, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 4,
	4, 4, 3, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 2, 2, 2, 2, 2,
	2, 2, 1, 3, 1, 2, 3, 0, 1, 1,
	2, 3, 2, 1, 2, 3, 1, 2, 3, 1,
	3, 4, 6, 7, 0, 1, 1, 2, 2, 3,
	4, 5,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	37, 39, 38, 41, 42, 59, 64, -1, 4, 4,
	4, 4, 53, 47, 67, 67, 67, 69, -24, 4,
	71, 40, 69, 69, 69, -12, -13, -14, 4, -29,
	-30, 69, 33, 81, 82, 83, 84, 85, 86, 87,
	4, -18, 4, -19, -20, -21, 4, 32, 72, 67,
	71, -19, -20, -20, 68, 73, 74, 75, 68, 27,
	28, 14, 65, 29, 76, 77, 78, 79, -29, -29,
	-29, -29, -29, -29, -29, -29, -29, 88, 68, 73,
	70, 73, -24, -17, 4, 5, 62, 63, 33, 8,
	69, 48, 51, 49, 52, 71, -24, -23, -24, 72,
	-17, 70, 70, 70, 75, -13, -24, 4, 68, 75,
	-29, -29, -29, -29, -29, -29, -29, -29, -29, 70,
	-30, 75, -18, 67, -20, 75, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 27, 28, 29, 30,
	31, 34, 35, 36, -17, -17, -17, 69, 69, 69,
	69, -22, -17, 68, 73, 67, 72, 67, -27, 89,
	-27, -25, -26, -27, 69, 75, -10, -11, -5, -6,
	-8, -17, -17, -17, -17, -17, -17, -17, -17, -17,
	-17, -17, -17, -17, -17, -17, -17, -17, -17, 70,
	-22, -22, -22, -22, 72, 73, -23, -23, 67, -15,
	-16, 4, 67, 43, 44, 46, 47, 50, 53, 54,
	55, 60, 61, -17, 75, -5, 67, 4, 67, -26,
	-22, 68, -10, 70, 70, 70, 70, -22, 68, -23,
	68, -15, 74, 32, 17, 18, 19, 20, 21, 22,
	23, 24, 25, 26, -15, 4, -17, -28, 67, 69,
	69, 69, 67, 4, 75, 4, 75, 75, -15, -15,
	70, 75, 68, 75, -16, -17, -17, -17, -17, -17,
	-17, -17, -17, -17, -17, -17, 68, -24, 67, 75,
	-15, -22, -22, -22, -15, 56, 75, 68, 68, -25,
	75, 75, 75, 75, 75, 75, 75, 75, 75, 75,
	75, 75, 75, 32, -15, 68, 70, 70, 70, 68,
	-17, 57, 75, 75, -17, 68, 73, -25, 75, -25,
	75, 67, -17, 75, 75, 45, -28, 75, 75, -15,
	58, 67, 68, -17, -15, 75, 67, 68, -15, 75,
	68, 75,
}
var yyDef = []int{

	0, -2, 1, 3, 4, 5, 6, 7, 8, 9,
	0, 0, 0, 0, 0, 0, 0, 2, 0, 0,
	0, 0, 0, 0, 24, 0, 0, 117, 0, 129,
	0, 0, 117, 0, 0, 0, 25, 26, 0, 0,
	93, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	112, 0, 114, 0, 118, 119, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 27, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 104,
	105, 106, 107, 108, 109, 110, 111, 0, 0, 115,
	0, 120, 122, 0, 63, 64, 65, 66, 0, 0,
	0, 0, 0, 0, 0, 0, 130, 0, 126, 0,
	0, 0, 0, 0, 21, 28, 134, 129, 0, 23,
	95, 96, 97, 98, 99, 100, 101, 102, 103, 94,
	113, 10, 116, 12, 121, 17, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 67, 68, 0, 0, 0, 0,
	0, 0, 123, 131, 127, 0, 0, 31, 0, 0,
	0, 29, 135, 136, 0, 22, 0, 12, 14, 15,
	16, 70, 71, 72, 73, 74, 75, 76, 77, 78,
	79, 80, 81, 82, 83, 84, 85, 86, 87, 69,
	0, 0, 0, 0, 92, 124, 128, 0, 0, 0,
	31, 63, 31, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 61, 62, 31, 138, 31, 137,
	0, 0, 13, 88, 89, 90, 91, 125, 132, 0,
	0, 32, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 31, 0,
	0, 0, 31, 0, 57, 0, 59, 60, 0, 0,
	134, 11, 133, 18, 33, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 31, 50,
	0, 0, 0, 0, 0, 0, 58, 0, 0, 30,
	39, 40, 41, 42, 43, 44, 45, 46, 47, 48,
	49, 34, 35, 0, 0, 139, 134, 0, 134, 0,
	0, 0, 19, 20, 0, 0, 140, 0, 52, 0,
	54, 31, 0, 36, 37, 0, 141, 51, 53, 0,
	0, 31, 0, 0, 0, 55, 31, 0, 0, 38,
	0, 56,
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
			yyVAL.definitions = []data.Def{yyS[yypt-0].definition}
			if l, isLexerWrapper := yylex.(*lexerWrapper); isLexerWrapper {
				l.definitions = yyVAL.definitions
			}
		}
	case 2:
		//line parser.go.y:154
		{
			yyVAL.definitions = append([]data.Def{yyS[yypt-1].definition}, yyS[yypt-0].definitions...)
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
		yyVAL.definition = yyS[yypt-0].definition
	case 10:
		//line parser.go.y:172
		{
			yyVAL.definition = data.DataDef{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-4].tok.lit, Elems: yyS[yypt-2].identifiers}
		}
	case 11:
		//line parser.go.y:178
		{
			yyVAL.definition = data.ModuleDef{Pos: yyS[yypt-8].tok.pos, Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Defs: yyS[yypt-2].definitions}
		}
	case 12:
		//line parser.go.y:184
		{
			yyVAL.definitions = nil
		}
	case 13:
		//line parser.go.y:188
		{
			yyVAL.definitions = append([]data.Def{yyS[yypt-1].definition}, yyS[yypt-0].definitions...)
		}
	case 14:
		yyVAL.definition = yyS[yypt-0].definition
	case 15:
		yyVAL.definition = yyS[yypt-0].definition
	case 16:
		yyVAL.definition = yyS[yypt-0].definition
	case 17:
		//line parser.go.y:199
		{
			yyVAL.definition = data.ConstantDef{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-4].tok.lit, Type: yyS[yypt-3].typetype, Expr: yyS[yypt-1].expression}
		}
	case 18:
		//line parser.go.y:205
		{
			yyVAL.definition = data.ProcDef{Pos: yyS[yypt-8].tok.pos, Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Stmts: yyS[yypt-2].statements}
		}
	case 19:
		//line parser.go.y:211
		{
			yyVAL.definition = data.FaultDef{Pos: yyS[yypt-9].tok.pos, Name: yyS[yypt-8].tok.lit, Parameters: yyS[yypt-6].parameters, Tag: yyS[yypt-4].tag, Stmts: yyS[yypt-2].statements}
		}
	case 20:
		//line parser.go.y:215
		{
			yyVAL.definition = data.FaultDef{}
		}
	case 21:
		//line parser.go.y:221
		{
			yyVAL.definition = data.InitBlock{Pos: yyS[yypt-4].tok.pos, Vars: yyS[yypt-2].initvars}
		}
	case 22:
		//line parser.go.y:227
		{
			yyVAL.definition = data.LtlSpec{Expr: yyS[yypt-3].ltlexpr}
		}
	case 23:
		//line parser.go.y:231
		{
			yyVAL.definition = data.LtlSpec{Expr: yyS[yypt-2].ltlexpr}
		}
	case 24:
		//line parser.go.y:237
		{
			yyVAL.initvars = nil
		}
	case 25:
		//line parser.go.y:241
		{
			yyVAL.initvars = yyS[yypt-0].initvars
		}
	case 26:
		//line parser.go.y:247
		{
			yyVAL.initvars = []data.InitVar{yyS[yypt-0].initvar}
		}
	case 27:
		//line parser.go.y:251
		{
			yyVAL.initvars = []data.InitVar{yyS[yypt-1].initvar}
		}
	case 28:
		//line parser.go.y:255
		{
			yyVAL.initvars = append([]data.InitVar{yyS[yypt-2].initvar}, yyS[yypt-0].initvars...)
		}
	case 29:
		//line parser.go.y:260
		{
			yyVAL.initvar = data.ChannelVar{Pos: yyS[yypt-3].tok.pos, Name: yyS[yypt-3].tok.lit, Type: yyS[yypt-1].typetype, Tags: yyS[yypt-0].tags}
		}
	case 30:
		//line parser.go.y:264
		{
			yyVAL.initvar = data.InstanceVar{Pos: yyS[yypt-6].tok.pos, Name: yyS[yypt-6].tok.lit, ProcDefName: yyS[yypt-4].tok.lit, Args: yyS[yypt-2].expressions, Tags: yyS[yypt-0].tags}
		}
	case 31:
		//line parser.go.y:270
		{
			yyVAL.statements = nil
		}
	case 32:
		//line parser.go.y:274
		{
			yyVAL.statements = append([]data.Stmt{yyS[yypt-1].statement}, yyS[yypt-0].statements...)
		}
	case 33:
		//line parser.go.y:280
		{
			yyVAL.statement = data.LabelledStmt{Pos: yyS[yypt-2].tok.pos, Label: yyS[yypt-2].tok.lit, Stmt: yyS[yypt-0].statement}
		}
	case 34:
		//line parser.go.y:284
		{
			yyVAL.statement = data.BlockStmt{Pos: yyS[yypt-3].tok.pos, Stmts: yyS[yypt-2].statements}
		}
	case 35:
		//line parser.go.y:288
		{
			yyVAL.statement = data.VarDeclStmt{Pos: yyS[yypt-3].tok.pos, Name: yyS[yypt-2].tok.lit, Type: yyS[yypt-1].typetype}
		}
	case 36:
		//line parser.go.y:292
		{
			yyVAL.statement = data.VarDeclStmt{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-4].tok.lit, Type: yyS[yypt-3].typetype, Initializer: yyS[yypt-1].expression}
		}
	case 37:
		//line parser.go.y:296
		{
			yyVAL.statement = data.IfStmt{Pos: yyS[yypt-5].tok.pos, Condition: yyS[yypt-4].expression, TrueBranch: yyS[yypt-2].statements}
		}
	case 38:
		//line parser.go.y:300
		{
			yyVAL.statement = data.IfStmt{Pos: yyS[yypt-9].tok.pos, Condition: yyS[yypt-8].expression, TrueBranch: yyS[yypt-6].statements, FalseBranch: yyS[yypt-2].statements}
		}
	case 39:
		//line parser.go.y:304
		{
			yyVAL.statement = data.AssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Expr: yyS[yypt-1].expression}
		}
	case 40:
		//line parser.go.y:308
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "+", Expr: yyS[yypt-1].expression}
		}
	case 41:
		//line parser.go.y:312
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "-", Expr: yyS[yypt-1].expression}
		}
	case 42:
		//line parser.go.y:316
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "*", Expr: yyS[yypt-1].expression}
		}
	case 43:
		//line parser.go.y:320
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "/", Expr: yyS[yypt-1].expression}
		}
	case 44:
		//line parser.go.y:324
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "%", Expr: yyS[yypt-1].expression}
		}
	case 45:
		//line parser.go.y:328
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "&", Expr: yyS[yypt-1].expression}
		}
	case 46:
		//line parser.go.y:332
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "|", Expr: yyS[yypt-1].expression}
		}
	case 47:
		//line parser.go.y:336
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "^", Expr: yyS[yypt-1].expression}
		}
	case 48:
		//line parser.go.y:340
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "<<", Expr: yyS[yypt-1].expression}
		}
	case 49:
		//line parser.go.y:344
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: ">>", Expr: yyS[yypt-1].expression}
		}
	case 50:
		//line parser.go.y:348
		{
			yyVAL.statement = data.ChoiceStmt{Pos: yyS[yypt-2].tok.pos, Blocks: yyS[yypt-1].blocks}
		}
	case 51:
		//line parser.go.y:352
		{
			yyVAL.statement = data.RecvStmt{Pos: yyS[yypt-5].tok.pos, Channel: yyS[yypt-3].expressions[0], Args: yyS[yypt-3].expressions[1:], Tags: yyS[yypt-1].tags}
		}
	case 52:
		//line parser.go.y:356
		{
			yyVAL.statement = data.PeekStmt{Pos: yyS[yypt-4].tok.pos, Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 53:
		//line parser.go.y:360
		{
			yyVAL.statement = data.SendStmt{Pos: yyS[yypt-5].tok.pos, Channel: yyS[yypt-3].expressions[0], Args: yyS[yypt-3].expressions[1:], Tags: yyS[yypt-1].tags}
		}
	case 54:
		//line parser.go.y:364
		{
			yyVAL.statement = data.ForStmt{Pos: yyS[yypt-4].tok.pos, Stmts: yyS[yypt-2].statements}
		}
	case 55:
		//line parser.go.y:368
		{
			yyVAL.statement = data.ForInStmt{Pos: yyS[yypt-7].tok.pos, Variable: yyS[yypt-6].tok.lit, Container: yyS[yypt-4].expression, Stmts: yyS[yypt-2].statements}
		}
	case 56:
		//line parser.go.y:372
		{
			yyVAL.statement = data.ForInRangeStmt{Pos: yyS[yypt-10].tok.pos, Variable: yyS[yypt-9].tok.lit, FromExpr: yyS[yypt-6].expression, ToExpr: yyS[yypt-4].expression, Stmts: yyS[yypt-2].statements}
		}
	case 57:
		//line parser.go.y:376
		{
			yyVAL.statement = data.BreakStmt{Pos: yyS[yypt-1].tok.pos}
		}
	case 58:
		//line parser.go.y:380
		{
			yyVAL.statement = data.GotoStmt{Pos: yyS[yypt-2].tok.pos, Label: yyS[yypt-1].tok.lit}
		}
	case 59:
		//line parser.go.y:384
		{
			yyVAL.statement = data.SkipStmt{Pos: yyS[yypt-1].tok.pos}
		}
	case 60:
		//line parser.go.y:388
		{
			yyVAL.statement = data.ExprStmt{Expr: yyS[yypt-1].expression}
		}
	case 61:
		//line parser.go.y:392
		{
			yyVAL.statement = data.NullStmt{Pos: yyS[yypt-0].tok.pos}
		}
	case 62:
		//line parser.go.y:396
		{
			yyVAL.statement = yyS[yypt-0].definition.(data.Stmt)
		}
	case 63:
		//line parser.go.y:401
		{
			yyVAL.expression = data.IdentifierExpr{Pos: yyS[yypt-0].tok.pos, Name: yyS[yypt-0].tok.lit}
		}
	case 64:
		//line parser.go.y:405
		{
			yyVAL.expression = data.NumberExpr{Pos: yyS[yypt-0].tok.pos, Lit: yyS[yypt-0].tok.lit}
		}
	case 65:
		//line parser.go.y:409
		{
			yyVAL.expression = data.TrueExpr{Pos: yyS[yypt-0].tok.pos}
		}
	case 66:
		//line parser.go.y:413
		{
			yyVAL.expression = data.FalseExpr{Pos: yyS[yypt-0].tok.pos}
		}
	case 67:
		//line parser.go.y:417
		{
			yyVAL.expression = data.NotExpr{Pos: yyS[yypt-1].tok.pos, SubExpr: yyS[yypt-0].expression}
		}
	case 68:
		//line parser.go.y:421
		{
			yyVAL.expression = data.UnarySubExpr{Pos: yyS[yypt-1].tok.pos, SubExpr: yyS[yypt-0].expression}
		}
	case 69:
		//line parser.go.y:425
		{
			yyVAL.expression = data.ParenExpr{Pos: yyS[yypt-2].tok.pos, SubExpr: yyS[yypt-1].expression}
		}
	case 70:
		//line parser.go.y:429
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "+", RHS: yyS[yypt-0].expression}
		}
	case 71:
		//line parser.go.y:433
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "-", RHS: yyS[yypt-0].expression}
		}
	case 72:
		//line parser.go.y:437
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "*", RHS: yyS[yypt-0].expression}
		}
	case 73:
		//line parser.go.y:441
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "/", RHS: yyS[yypt-0].expression}
		}
	case 74:
		//line parser.go.y:445
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "%", RHS: yyS[yypt-0].expression}
		}
	case 75:
		//line parser.go.y:449
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "&", RHS: yyS[yypt-0].expression}
		}
	case 76:
		//line parser.go.y:453
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "|", RHS: yyS[yypt-0].expression}
		}
	case 77:
		//line parser.go.y:457
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "^", RHS: yyS[yypt-0].expression}
		}
	case 78:
		//line parser.go.y:461
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "<<", RHS: yyS[yypt-0].expression}
		}
	case 79:
		//line parser.go.y:465
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: ">>", RHS: yyS[yypt-0].expression}
		}
	case 80:
		//line parser.go.y:469
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "&&", RHS: yyS[yypt-0].expression}
		}
	case 81:
		//line parser.go.y:473
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "||", RHS: yyS[yypt-0].expression}
		}
	case 82:
		//line parser.go.y:477
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "==", RHS: yyS[yypt-0].expression}
		}
	case 83:
		//line parser.go.y:481
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "<", RHS: yyS[yypt-0].expression}
		}
	case 84:
		//line parser.go.y:485
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: ">", RHS: yyS[yypt-0].expression}
		}
	case 85:
		//line parser.go.y:489
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "!=", RHS: yyS[yypt-0].expression}
		}
	case 86:
		//line parser.go.y:493
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "<=", RHS: yyS[yypt-0].expression}
		}
	case 87:
		//line parser.go.y:497
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: ">=", RHS: yyS[yypt-0].expression}
		}
	case 88:
		//line parser.go.y:501
		{
			yyVAL.expression = data.TimeoutRecvExpr{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 89:
		//line parser.go.y:505
		{
			yyVAL.expression = data.TimeoutPeekExpr{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 90:
		//line parser.go.y:509
		{
			yyVAL.expression = data.NonblockRecvExpr{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 91:
		//line parser.go.y:513
		{
			yyVAL.expression = data.NonblockPeekExpr{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 92:
		//line parser.go.y:517
		{
			yyVAL.expression = data.ArrayExpr{Pos: yyS[yypt-2].tok.pos, Elems: yyS[yypt-1].expressions}
		}
	case 93:
		//line parser.go.y:524
		{
			yyVAL.ltlexpr = yyS[yypt-0].ltlatom
		}
	case 94:
		//line parser.go.y:528
		{
			yyVAL.ltlexpr = data.ParenLtlExpr{SubExpr: yyS[yypt-1].ltlexpr}
		}
	case 95:
		//line parser.go.y:532
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "&", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 96:
		//line parser.go.y:536
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "|", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 97:
		//line parser.go.y:540
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "^", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 98:
		//line parser.go.y:544
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "->", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 99:
		//line parser.go.y:548
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "=", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 100:
		//line parser.go.y:552
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "U", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 101:
		//line parser.go.y:556
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "V", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 102:
		//line parser.go.y:560
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "S", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 103:
		//line parser.go.y:564
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "T", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 104:
		//line parser.go.y:568
		{
			yyVAL.ltlexpr = data.UnOpLtlExpr{Operator: "!", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 105:
		//line parser.go.y:572
		{
			yyVAL.ltlexpr = data.UnOpLtlExpr{Operator: "X", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 106:
		//line parser.go.y:576
		{
			yyVAL.ltlexpr = data.UnOpLtlExpr{Operator: "G", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 107:
		//line parser.go.y:580
		{
			yyVAL.ltlexpr = data.UnOpLtlExpr{Operator: "F", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 108:
		//line parser.go.y:584
		{
			yyVAL.ltlexpr = data.UnOpLtlExpr{Operator: "Y", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 109:
		//line parser.go.y:588
		{
			yyVAL.ltlexpr = data.UnOpLtlExpr{Operator: "Z", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 110:
		//line parser.go.y:592
		{
			yyVAL.ltlexpr = data.UnOpLtlExpr{Operator: "H", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 111:
		//line parser.go.y:596
		{
			yyVAL.ltlexpr = data.UnOpLtlExpr{Operator: "O", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 112:
		//line parser.go.y:601
		{
			yyVAL.ltlatom = data.LtlAtomExpr{Names: []string{yyS[yypt-0].tok.lit}}
		}
	case 113:
		//line parser.go.y:605
		{
			yyVAL.ltlatom = data.LtlAtomExpr{Names: append([]string{yyS[yypt-2].tok.lit}, yyS[yypt-0].ltlatom.Names...)}
		}
	case 114:
		//line parser.go.y:613
		{
			yyVAL.identifiers = []string{yyS[yypt-0].tok.lit}
		}
	case 115:
		//line parser.go.y:617
		{
			yyVAL.identifiers = []string{yyS[yypt-1].tok.lit}
		}
	case 116:
		//line parser.go.y:621
		{
			yyVAL.identifiers = append([]string{yyS[yypt-2].tok.lit}, yyS[yypt-0].identifiers...)
		}
	case 117:
		//line parser.go.y:627
		{
			yyVAL.parameters = nil
		}
	case 118:
		//line parser.go.y:631
		{
			yyVAL.parameters = yyS[yypt-0].parameters
		}
	case 119:
		//line parser.go.y:637
		{
			yyVAL.parameters = []data.Parameter{yyS[yypt-0].parameter}
		}
	case 120:
		//line parser.go.y:641
		{
			yyVAL.parameters = []data.Parameter{yyS[yypt-1].parameter}
		}
	case 121:
		//line parser.go.y:645
		{
			yyVAL.parameters = append([]data.Parameter{yyS[yypt-2].parameter}, yyS[yypt-0].parameters...)
		}
	case 122:
		//line parser.go.y:651
		{
			yyVAL.parameter = data.Parameter{Name: yyS[yypt-1].tok.lit, Type: yyS[yypt-0].typetype}
		}
	case 123:
		//line parser.go.y:657
		{
			yyVAL.expressions = []data.Expr{yyS[yypt-0].expression}
		}
	case 124:
		//line parser.go.y:661
		{
			yyVAL.expressions = []data.Expr{yyS[yypt-1].expression}
		}
	case 125:
		//line parser.go.y:665
		{
			yyVAL.expressions = append([]data.Expr{yyS[yypt-2].expression}, yyS[yypt-0].expressions...)
		}
	case 126:
		//line parser.go.y:671
		{
			yyVAL.typetypes = []data.Type{yyS[yypt-0].typetype}
		}
	case 127:
		//line parser.go.y:675
		{
			yyVAL.typetypes = []data.Type{yyS[yypt-1].typetype}
		}
	case 128:
		//line parser.go.y:679
		{
			yyVAL.typetypes = append([]data.Type{yyS[yypt-2].typetype}, yyS[yypt-0].typetypes...)
		}
	case 129:
		//line parser.go.y:684
		{
			yyVAL.typetype = data.NamedType{Name: yyS[yypt-0].tok.lit}
		}
	case 130:
		//line parser.go.y:688
		{
			yyVAL.typetype = data.ArrayType{ElemType: yyS[yypt-0].typetype}
		}
	case 131:
		//line parser.go.y:692
		{
			yyVAL.typetype = data.HandshakeChannelType{Elems: yyS[yypt-1].typetypes}
		}
	case 132:
		//line parser.go.y:696
		{
			yyVAL.typetype = data.BufferedChannelType{BufferSize: nil, Elems: yyS[yypt-1].typetypes}
		}
	case 133:
		//line parser.go.y:700
		{
			yyVAL.typetype = data.BufferedChannelType{BufferSize: yyS[yypt-4].expression, Elems: yyS[yypt-1].typetypes}
		}
	case 134:
		//line parser.go.y:706
		{
			yyVAL.tags = nil
		}
	case 135:
		//line parser.go.y:710
		{
			yyVAL.tags = yyS[yypt-0].tags
		}
	case 136:
		//line parser.go.y:716
		{
			yyVAL.tags = []string{yyS[yypt-0].tag}
		}
	case 137:
		//line parser.go.y:720
		{
			yyVAL.tags = append([]string{yyS[yypt-1].tag}, yyS[yypt-0].tags...)
		}
	case 138:
		//line parser.go.y:726
		{
			yyVAL.tag = yyS[yypt-0].tok.lit
		}
	case 139:
		//line parser.go.y:732
		{
			yyVAL.blocks = []data.BlockStmt{data.BlockStmt{Pos: yyS[yypt-2].tok.pos, Stmts: yyS[yypt-1].statements}}
		}
	case 140:
		//line parser.go.y:736
		{
			yyVAL.blocks = []data.BlockStmt{data.BlockStmt{Pos: yyS[yypt-3].tok.pos, Stmts: yyS[yypt-2].statements}}
		}
	case 141:
		//line parser.go.y:740
		{
			yyVAL.blocks = append([]data.BlockStmt{data.BlockStmt{Pos: yyS[yypt-4].tok.pos, Stmts: yyS[yypt-3].statements}}, yyS[yypt-0].blocks...)
		}
	}
	goto yystack /* stack new state and value */
}
