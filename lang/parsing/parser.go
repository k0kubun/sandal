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
const BLOCK = 57390
const TIMEOUT_RECV = 57391
const NONBLOCK_RECV = 57392
const PEEK = 57393
const TIMEOUT_PEEK = 57394
const NONBLOCK_PEEK = 57395
const SEND = 57396
const FOR = 57397
const BREAK = 57398
const IN = 57399
const RANGE = 57400
const TO = 57401
const INIT = 57402
const GOTO = 57403
const SKIP = 57404
const TRUE = 57405
const FALSE = 57406
const LTL = 57407
const THEN = 57408
const IFF = 57409
const UNARY = 57410

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
	"BLOCK",
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

//line parser.go.y:753

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

const yyNprod = 144
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1042

var yyAct = []int{

	134, 211, 222, 227, 223, 120, 42, 114, 54, 185,
	57, 224, 37, 93, 121, 43, 344, 360, 314, 77,
	113, 136, 5, 8, 5, 358, 354, 347, 29, 6,
	80, 81, 82, 83, 346, 341, 36, 53, 340, 336,
	334, 329, 326, 324, 286, 65, 66, 343, 279, 263,
	226, 84, 85, 86, 87, 88, 89, 90, 91, 92,
	220, 99, 313, 218, 116, 155, 45, 98, 142, 112,
	137, 71, 332, 76, 188, 97, 95, 70, 255, 139,
	76, 143, 144, 145, 146, 147, 148, 149, 150, 151,
	152, 61, 30, 74, 75, 78, 323, 62, 319, 318,
	79, 63, 317, 44, 156, 178, 179, 180, 158, 154,
	294, 293, 186, 292, 291, 46, 47, 48, 49, 50,
	51, 52, 140, 119, 118, 77, 210, 195, 32, 208,
	192, 193, 77, 117, 96, 73, 80, 81, 82, 83,
	225, 39, 72, 80, 81, 82, 83, 215, 214, 213,
	184, 183, 182, 181, 35, 34, 33, 28, 32, 359,
	31, 232, 233, 234, 235, 236, 237, 238, 239, 240,
	241, 242, 243, 244, 245, 246, 247, 248, 249, 229,
	356, 231, 186, 186, 186, 186, 351, 230, 331, 328,
	31, 251, 252, 253, 254, 327, 325, 260, 265, 266,
	267, 268, 269, 270, 271, 272, 273, 274, 275, 257,
	258, 264, 138, 320, 186, 186, 186, 277, 280, 316,
	298, 296, 284, 281, 282, 283, 186, 289, 276, 287,
	217, 194, 290, 187, 141, 288, 76, 10, 12, 11,
	94, 13, 14, 69, 350, 212, 262, 261, 259, 74,
	229, 78, 231, 191, 189, 157, 79, 186, 230, 67,
	15, 27, 26, 25, 285, 16, 295, 299, 300, 160,
	161, 162, 163, 164, 165, 166, 167, 168, 169, 312,
	297, 60, 56, 219, 315, 209, 321, 58, 77, 170,
	171, 172, 173, 174, 216, 59, 175, 176, 177, 80,
	81, 82, 83, 55, 198, 199, 200, 201, 202, 203,
	204, 205, 206, 207, 12, 330, 64, 13, 53, 197,
	333, 41, 335, 338, 122, 101, 339, 68, 105, 21,
	23, 24, 20, 19, 345, 18, 15, 22, 342, 162,
	163, 164, 165, 348, 1, 168, 169, 17, 40, 38,
	352, 228, 9, 104, 7, 4, 353, 3, 12, 2,
	0, 357, 196, 124, 125, 0, 126, 127, 0, 107,
	109, 128, 108, 110, 129, 130, 131, 0, 0, 0,
	0, 132, 133, 102, 103, 0, 0, 0, 123, 0,
	106, 0, 111, 0, 0, 0, 135, 160, 161, 162,
	163, 164, 165, 166, 167, 168, 169, 160, 161, 162,
	163, 164, 165, 166, 167, 168, 169, 170, 171, 172,
	173, 174, 0, 0, 175, 176, 177, 170, 171, 172,
	173, 174, 0, 0, 175, 176, 177, 0, 0, 0,
	160, 161, 162, 163, 164, 165, 166, 167, 168, 169,
	160, 161, 162, 163, 164, 165, 166, 167, 168, 169,
	170, 171, 172, 173, 174, 0, 311, 175, 176, 177,
	170, 171, 172, 173, 174, 0, 310, 175, 176, 177,
	0, 0, 0, 160, 161, 162, 163, 164, 165, 166,
	167, 168, 169, 160, 161, 162, 163, 164, 165, 166,
	167, 168, 169, 170, 171, 172, 173, 174, 0, 309,
	175, 176, 177, 170, 171, 172, 173, 174, 0, 308,
	175, 176, 177, 0, 0, 0, 160, 161, 162, 163,
	164, 165, 166, 167, 168, 169, 160, 161, 162, 163,
	164, 165, 166, 167, 168, 169, 170, 171, 172, 173,
	174, 0, 307, 175, 176, 177, 170, 171, 172, 173,
	174, 0, 306, 175, 176, 177, 0, 0, 0, 160,
	161, 162, 163, 164, 165, 166, 167, 168, 169, 160,
	161, 162, 163, 164, 165, 166, 167, 168, 169, 170,
	171, 172, 173, 174, 0, 305, 175, 176, 177, 170,
	171, 172, 173, 174, 0, 304, 175, 176, 177, 0,
	0, 0, 160, 161, 162, 163, 164, 165, 166, 167,
	168, 169, 160, 161, 162, 163, 164, 165, 166, 167,
	168, 169, 170, 171, 172, 173, 174, 0, 303, 175,
	176, 177, 170, 171, 172, 173, 174, 0, 302, 175,
	176, 177, 0, 0, 0, 160, 161, 162, 163, 164,
	165, 166, 167, 168, 169, 160, 161, 162, 163, 164,
	165, 166, 167, 168, 169, 170, 171, 172, 173, 174,
	0, 301, 175, 176, 177, 170, 171, 172, 173, 174,
	0, 221, 175, 176, 177, 76, 0, 0, 160, 161,
	162, 163, 164, 165, 166, 167, 168, 169, 74, 75,
	78, 0, 0, 0, 0, 79, 0, 0, 170, 171,
	172, 173, 174, 0, 159, 175, 176, 177, 0, 0,
	0, 0, 256, 160, 161, 162, 163, 164, 165, 166,
	167, 168, 169, 0, 0, 0, 0, 77, 0, 0,
	0, 0, 153, 0, 0, 0, 0, 0, 80, 81,
	82, 83, 76, 0, 190, 160, 161, 162, 163, 164,
	165, 166, 167, 168, 169, 0, 0, 78, 0, 0,
	0, 0, 79, 0, 0, 170, 171, 172, 173, 174,
	0, 0, 175, 176, 177, 160, 161, 162, 163, 164,
	165, 166, 167, 168, 169, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 77, 170, 171, 172, 173, 174,
	0, 0, 175, 176, 177, 80, 81, 82, 83, 250,
	160, 161, 162, 163, 164, 165, 166, 167, 168, 169,
	160, 161, 162, 163, 164, 165, 166, 167, 168, 169,
	170, 171, 172, 173, 174, 0, 355, 175, 176, 177,
	170, 171, 172, 173, 174, 0, 0, 175, 176, 177,
	100, 101, 0, 0, 105, 0, 0, 0, 0, 0,
	160, 161, 162, 163, 164, 165, 166, 167, 168, 169,
	0, 337, 0, 0, 0, 0, 100, 101, 0, 104,
	105, 278, 172, 173, 174, 0, 0, 175, 176, 177,
	0, 0, 0, 0, 0, 107, 109, 0, 108, 110,
	100, 101, 0, 0, 105, 104, 0, 0, 0, 102,
	103, 0, 0, 0, 0, 0, 106, 0, 111, 115,
	0, 107, 109, 0, 108, 110, 0, 0, 0, 104,
	322, 0, 0, 0, 0, 102, 103, 0, 0, 0,
	0, 0, 106, 0, 111, 107, 109, 0, 108, 110,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 102,
	103, 0, 0, 0, 0, 0, 106, 0, 111, 160,
	161, 162, 163, 164, 165, 166, 167, 168, 169, 160,
	161, 162, 163, 164, 165, 166, 167, 168, 169, 170,
	171, 172, 173, 174, 0, 0, 175, 176, 177, 170,
	0, 172, 173, 174, 0, 0, 175, 176, 177, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 349,
}
var yyPact = []int{

	200, -1000, 200, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	331, 329, 328, 325, 283, 195, 194, -1000, 193, 87,
	88, 86, 85, 84, -78, 317, 33, 299, 291, 249,
	-1000, 18, 29, 291, 291, 291, 191, 323, 174, -1000,
	3, -4, 66, -1000, 33, 33, 33, 33, 33, 33,
	33, 33, 33, -76, 171, 2, 63, -1000, 1, 88,
	916, 88, 88, 866, 62, 53, 52, 320, -1000, -6,
	317, 118, 165, -8, 33, 33, 33, 33, 33, 33,
	33, 33, 33, 33, 681, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 314, -11, 299, 187, 291, -1000, 648,
	-1000, -1000, -1000, -1000, 916, 916, 916, 83, 82, 81,
	80, 916, -1000, 164, 0, 186, 691, 185, -78, -78,
	162, 320, 287, 320, 281, 916, 177, 79, 78, 77,
	226, -13, 279, -16, 615, -1000, -1000, -1000, -1000, -78,
	70, -26, -1000, 748, 222, -47, -1000, 59, 59, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 276, -1000, -1000,
	916, 916, 916, 916, 916, 916, 916, 916, 916, 916,
	916, 916, 916, 916, 916, 916, 916, 916, -1000, -1000,
	758, 916, 916, 916, 916, 5, 658, -1000, 88, 88,
	180, 320, 179, 178, -27, -1000, 320, 916, 916, 916,
	916, 916, 916, 916, 916, 916, 916, 916, 159, 88,
	833, -28, 320, 916, 916, 916, 320, 207, -1000, -32,
	-1000, -1000, -1000, -1000, -78, 916, -1000, 158, 276, -1000,
	-1000, -1000, 330, 330, -1000, -1000, -1000, -1000, 330, 330,
	-1000, -1000, 873, 992, 726, 726, 726, 726, 726, 726,
	-1000, 43, 42, 40, 39, -1000, 916, -1000, 152, 88,
	151, 320, 320, -1000, -1000, 605, 572, 562, 529, 519,
	486, 476, 443, 433, 400, 390, -78, -14, 320, -1000,
	150, 31, 28, 27, 144, 892, -1000, -1000, 25, -33,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 127, -34, 126,
	120, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -35, -1000, 916, 119, -2, -78, -36, -78,
	-37, 823, 916, -78, -1000, -1000, -1000, -38, -41, -1000,
	262, -29, 177, -42, -1000, -49, -1000, 320, 982, -1000,
	-1000, -1000, -1000, -1000, 176, -1000, -1000, -1000, 117, 916,
	320, -50, 788, 111, -1000, 320, -51, 90, -1000, -59,
	-1000,
}
var yyPgo = []int{

	0, 344, 359, 357, 355, 21, 29, 354, 23, 352,
	3, 351, 349, 141, 348, 5, 14, 0, 8, 282,
	10, 287, 9, 20, 7, 2, 4, 11, 1, 6,
	15,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 2, 2, 2, 2,
	3, 4, 10, 10, 11, 11, 11, 5, 6, 7,
	7, 7, 8, 9, 9, 12, 12, 13, 13, 13,
	14, 14, 15, 15, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 29, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 30, 30, 18, 18, 18, 19,
	19, 20, 20, 20, 21, 22, 22, 22, 23, 23,
	23, 24, 24, 24, 24, 24, 25, 25, 26, 26,
	27, 28, 28, 28,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	6, 9, 0, 2, 1, 1, 1, 6, 9, 10,
	10, 7, 5, 6, 5, 0, 1, 1, 2, 3,
	4, 7, 0, 2, 3, 5, 4, 6, 6, 10,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 3, 6, 5, 6, 5, 8, 11, 2, 3,
	2, 2, 1, 1, 1, 1, 1, 1, 2, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4,
	4, 4, 4, 3, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 2, 2, 2,
	2, 2, 2, 2, 1, 3, 1, 2, 3, 0,
	1, 1, 2, 3, 2, 1, 2, 3, 1, 2,
	3, 1, 3, 4, 6, 7, 0, 1, 1, 2,
	2, 3, 4, 5,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	37, 39, 38, 41, 42, 60, 65, -1, 4, 4,
	4, 4, 54, 47, 48, 68, 68, 68, 70, -24,
	4, 72, 40, 70, 70, 70, -27, 90, -12, -13,
	-14, 4, -29, -30, 70, 33, 82, 83, 84, 85,
	86, 87, 88, 4, -18, 4, -19, -20, -21, 4,
	32, 73, 68, 72, -19, -20, -20, 68, 4, 69,
	74, 75, 76, 69, 27, 28, 14, 66, 29, 34,
	77, 78, 79, 80, -29, -29, -29, -29, -29, -29,
	-29, -29, -29, 89, 69, 74, 71, 74, -24, -17,
	4, 5, 63, 64, 33, 8, 70, 49, 52, 50,
	53, 72, -24, -23, -24, 73, -17, 71, 71, 71,
	-15, -16, 4, 68, 43, 44, 46, 47, 51, 54,
	55, 56, 61, 62, -17, 76, -5, 76, -13, -24,
	4, 69, 76, -29, -29, -29, -29, -29, -29, -29,
	-29, -29, -29, 71, -30, 76, -18, 68, -20, 76,
	7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
	27, 28, 29, 30, 31, 34, 35, 36, -17, -17,
	-17, 70, 70, 70, 70, -22, -17, 69, 74, 68,
	73, 68, -27, -27, 69, -15, 75, 32, 17, 18,
	19, 20, 21, 22, 23, 24, 25, 26, -15, 4,
	-17, -28, 68, 70, 70, 70, 68, 4, 76, 4,
	76, 76, -25, -26, -27, 70, 76, -10, -11, -5,
	-6, -8, -17, -17, -17, -17, -17, -17, -17, -17,
	-17, -17, -17, -17, -17, -17, -17, -17, -17, -17,
	71, -22, -22, -22, -22, 73, 74, -23, -23, 68,
	-15, 68, 68, 76, -16, -17, -17, -17, -17, -17,
	-17, -17, -17, -17, -17, -17, 69, -24, 68, 76,
	-15, -22, -22, -22, -15, 57, 76, -26, -22, 69,
	-10, 71, 71, 71, 71, -22, 69, -23, 69, -15,
	-15, 76, 76, 76, 76, 76, 76, 76, 76, 76,
	76, 76, -25, 76, 32, -15, 69, 71, 71, 71,
	69, -17, 58, 71, 76, 69, 76, 69, 69, 76,
	-17, 69, 74, -25, 76, -25, 76, 68, -17, -25,
	76, 76, 76, 76, 45, -28, 76, 76, -15, 59,
	68, 69, -17, -15, 76, 68, 69, -15, 76, 69,
	76,
}
var yyDef = []int{

	0, -2, 1, 3, 4, 5, 6, 7, 8, 9,
	0, 0, 0, 0, 0, 0, 0, 2, 0, 0,
	0, 0, 0, 0, 0, 25, 0, 0, 119, 0,
	131, 0, 0, 119, 0, 0, 0, 0, 0, 26,
	27, 0, 0, 94, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 114, 0, 116, 0, 120, 121, 0,
	0, 0, 0, 0, 0, 0, 0, 32, 140, 0,
	28, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 106, 107, 108, 109, 110,
	111, 112, 113, 0, 0, 117, 0, 122, 124, 0,
	64, 65, 66, 67, 0, 0, 0, 0, 0, 0,
	0, 0, 132, 0, 128, 0, 0, 0, 0, 0,
	0, 32, 64, 32, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 62, 63, 22, 29, 136,
	131, 0, 24, 96, 97, 98, 99, 100, 101, 102,
	103, 104, 105, 95, 115, 10, 118, 12, 123, 17,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 69,
	0, 0, 0, 0, 0, 0, 125, 133, 129, 0,
	0, 32, 0, 0, 0, 33, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 32, 0, 0, 0, 32, 0, 58, 0,
	60, 61, 30, 137, 138, 0, 23, 0, 12, 14,
	15, 16, 71, 72, 73, 74, 75, 76, 77, 78,
	79, 80, 81, 82, 83, 84, 85, 86, 87, 88,
	70, 0, 0, 0, 0, 93, 126, 130, 0, 0,
	0, 32, 32, 21, 34, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 136, 0, 32, 51,
	0, 0, 0, 0, 0, 0, 59, 139, 0, 0,
	13, 89, 90, 91, 92, 127, 134, 0, 0, 0,
	0, 40, 41, 42, 43, 44, 45, 46, 47, 48,
	49, 50, 0, 36, 0, 0, 141, 136, 0, 136,
	0, 0, 0, 136, 11, 135, 18, 0, 0, 35,
	0, 0, 142, 0, 53, 0, 55, 32, 0, 31,
	19, 20, 37, 38, 0, 143, 52, 54, 0, 0,
	32, 0, 0, 0, 56, 32, 0, 0, 39, 0,
	57,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	70, 71, 3, 3, 74, 3, 89, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 75, 76,
	3, 3, 3, 3, 90, 3, 3, 3, 3, 3,
	84, 83, 87, 3, 3, 3, 3, 3, 3, 88,
	3, 3, 3, 79, 80, 77, 78, 3, 82, 85,
	86, 72, 3, 73, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 68, 3, 69,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 81,
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
		//line parser.go.y:148
		{
			yyVAL.definitions = []data.Def{yyS[yypt-0].definition}
			if l, isLexerWrapper := yylex.(*lexerWrapper); isLexerWrapper {
				l.definitions = yyVAL.definitions
			}
		}
	case 2:
		//line parser.go.y:155
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
		//line parser.go.y:173
		{
			yyVAL.definition = data.DataDef{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-4].tok.lit, Elems: yyS[yypt-2].identifiers}
		}
	case 11:
		//line parser.go.y:179
		{
			yyVAL.definition = data.ModuleDef{Pos: yyS[yypt-8].tok.pos, Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Defs: yyS[yypt-2].definitions}
		}
	case 12:
		//line parser.go.y:185
		{
			yyVAL.definitions = nil
		}
	case 13:
		//line parser.go.y:189
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
		//line parser.go.y:200
		{
			yyVAL.definition = data.ConstantDef{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-4].tok.lit, Type: yyS[yypt-3].typetype, Expr: yyS[yypt-1].expression}
		}
	case 18:
		//line parser.go.y:206
		{
			yyVAL.definition = data.ProcDef{Pos: yyS[yypt-8].tok.pos, Name: yyS[yypt-7].tok.lit, Parameters: yyS[yypt-5].parameters, Stmts: yyS[yypt-2].statements}
		}
	case 19:
		//line parser.go.y:212
		{
			yyVAL.definition = data.FaultDef{Pos: yyS[yypt-9].tok.pos, Name: yyS[yypt-8].tok.lit, Parameters: yyS[yypt-6].parameters, Tag: yyS[yypt-4].tag, Stmts: yyS[yypt-2].statements}
		}
	case 20:
		//line parser.go.y:216
		{
			yyVAL.definition = data.FaultDef{Pos: yyS[yypt-9].tok.pos, Name: yyS[yypt-8].tok.lit, Parameters: yyS[yypt-6].parameters, Tag: yyS[yypt-4].tag, Stmts: yyS[yypt-2].statements}
		}
	case 21:
		//line parser.go.y:220
		{
			yyVAL.definition = data.FaultDef{Pos: yyS[yypt-6].tok.pos, Name: yyS[yypt-5].tok.lit, Parameters: []data.Parameter{}, Tag: yyS[yypt-4].tag, Stmts: yyS[yypt-2].statements}
		}
	case 22:
		//line parser.go.y:226
		{
			yyVAL.definition = data.InitBlock{Pos: yyS[yypt-4].tok.pos, Vars: yyS[yypt-2].initvars}
		}
	case 23:
		//line parser.go.y:232
		{
			yyVAL.definition = data.LtlSpec{Expr: yyS[yypt-3].ltlexpr}
		}
	case 24:
		//line parser.go.y:236
		{
			yyVAL.definition = data.LtlSpec{Expr: yyS[yypt-2].ltlexpr}
		}
	case 25:
		//line parser.go.y:242
		{
			yyVAL.initvars = nil
		}
	case 26:
		//line parser.go.y:246
		{
			yyVAL.initvars = yyS[yypt-0].initvars
		}
	case 27:
		//line parser.go.y:252
		{
			yyVAL.initvars = []data.InitVar{yyS[yypt-0].initvar}
		}
	case 28:
		//line parser.go.y:256
		{
			yyVAL.initvars = []data.InitVar{yyS[yypt-1].initvar}
		}
	case 29:
		//line parser.go.y:260
		{
			yyVAL.initvars = append([]data.InitVar{yyS[yypt-2].initvar}, yyS[yypt-0].initvars...)
		}
	case 30:
		//line parser.go.y:265
		{
			yyVAL.initvar = data.ChannelVar{Pos: yyS[yypt-3].tok.pos, Name: yyS[yypt-3].tok.lit, Type: yyS[yypt-1].typetype, Tags: yyS[yypt-0].tags}
		}
	case 31:
		//line parser.go.y:269
		{
			yyVAL.initvar = data.InstanceVar{Pos: yyS[yypt-6].tok.pos, Name: yyS[yypt-6].tok.lit, ProcDefName: yyS[yypt-4].tok.lit, Args: yyS[yypt-2].expressions, Tags: yyS[yypt-0].tags}
		}
	case 32:
		//line parser.go.y:275
		{
			yyVAL.statements = nil
		}
	case 33:
		//line parser.go.y:279
		{
			yyVAL.statements = append([]data.Stmt{yyS[yypt-1].statement}, yyS[yypt-0].statements...)
		}
	case 34:
		//line parser.go.y:285
		{
			yyVAL.statement = data.LabelledStmt{Pos: yyS[yypt-2].tok.pos, Label: yyS[yypt-2].tok.lit, Stmt: yyS[yypt-0].statement}
		}
	case 35:
		//line parser.go.y:289
		{
			yyVAL.statement = data.BlockStmt{Pos: yyS[yypt-4].tok.pos, Stmts: yyS[yypt-3].statements, Tags: yyS[yypt-1].tags}
		}
	case 36:
		//line parser.go.y:293
		{
			yyVAL.statement = data.VarDeclStmt{Pos: yyS[yypt-3].tok.pos, Name: yyS[yypt-2].tok.lit, Type: yyS[yypt-1].typetype}
		}
	case 37:
		//line parser.go.y:297
		{
			yyVAL.statement = data.VarDeclStmt{Pos: yyS[yypt-5].tok.pos, Name: yyS[yypt-4].tok.lit, Type: yyS[yypt-3].typetype, Initializer: yyS[yypt-1].expression}
		}
	case 38:
		//line parser.go.y:301
		{
			yyVAL.statement = data.IfStmt{Pos: yyS[yypt-5].tok.pos, Condition: yyS[yypt-4].expression, TrueBranch: yyS[yypt-2].statements}
		}
	case 39:
		//line parser.go.y:305
		{
			yyVAL.statement = data.IfStmt{Pos: yyS[yypt-9].tok.pos, Condition: yyS[yypt-8].expression, TrueBranch: yyS[yypt-6].statements, FalseBranch: yyS[yypt-2].statements}
		}
	case 40:
		//line parser.go.y:309
		{
			yyVAL.statement = data.AssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Expr: yyS[yypt-1].expression}
		}
	case 41:
		//line parser.go.y:313
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "+", Expr: yyS[yypt-1].expression}
		}
	case 42:
		//line parser.go.y:317
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "-", Expr: yyS[yypt-1].expression}
		}
	case 43:
		//line parser.go.y:321
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "*", Expr: yyS[yypt-1].expression}
		}
	case 44:
		//line parser.go.y:325
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "/", Expr: yyS[yypt-1].expression}
		}
	case 45:
		//line parser.go.y:329
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "%", Expr: yyS[yypt-1].expression}
		}
	case 46:
		//line parser.go.y:333
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "&", Expr: yyS[yypt-1].expression}
		}
	case 47:
		//line parser.go.y:337
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "|", Expr: yyS[yypt-1].expression}
		}
	case 48:
		//line parser.go.y:341
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "^", Expr: yyS[yypt-1].expression}
		}
	case 49:
		//line parser.go.y:345
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: "<<", Expr: yyS[yypt-1].expression}
		}
	case 50:
		//line parser.go.y:349
		{
			yyVAL.statement = data.OpAssignmentStmt{Pos: yyS[yypt-3].tok.pos, Variable: yyS[yypt-3].tok.lit, Operator: ">>", Expr: yyS[yypt-1].expression}
		}
	case 51:
		//line parser.go.y:353
		{
			yyVAL.statement = data.ChoiceStmt{Pos: yyS[yypt-2].tok.pos, Blocks: yyS[yypt-1].blocks}
		}
	case 52:
		//line parser.go.y:357
		{
			yyVAL.statement = data.RecvStmt{Pos: yyS[yypt-5].tok.pos, Channel: yyS[yypt-3].expressions[0], Args: yyS[yypt-3].expressions[1:], Tags: yyS[yypt-1].tags}
		}
	case 53:
		//line parser.go.y:361
		{
			yyVAL.statement = data.PeekStmt{Pos: yyS[yypt-4].tok.pos, Channel: yyS[yypt-2].expressions[0], Args: yyS[yypt-2].expressions[1:]}
		}
	case 54:
		//line parser.go.y:365
		{
			yyVAL.statement = data.SendStmt{Pos: yyS[yypt-5].tok.pos, Channel: yyS[yypt-3].expressions[0], Args: yyS[yypt-3].expressions[1:], Tags: yyS[yypt-1].tags}
		}
	case 55:
		//line parser.go.y:369
		{
			yyVAL.statement = data.ForStmt{Pos: yyS[yypt-4].tok.pos, Stmts: yyS[yypt-2].statements}
		}
	case 56:
		//line parser.go.y:373
		{
			yyVAL.statement = data.ForInStmt{Pos: yyS[yypt-7].tok.pos, Variable: yyS[yypt-6].tok.lit, Container: yyS[yypt-4].expression, Stmts: yyS[yypt-2].statements}
		}
	case 57:
		//line parser.go.y:377
		{
			yyVAL.statement = data.ForInRangeStmt{Pos: yyS[yypt-10].tok.pos, Variable: yyS[yypt-9].tok.lit, FromExpr: yyS[yypt-6].expression, ToExpr: yyS[yypt-4].expression, Stmts: yyS[yypt-2].statements}
		}
	case 58:
		//line parser.go.y:381
		{
			yyVAL.statement = data.BreakStmt{Pos: yyS[yypt-1].tok.pos}
		}
	case 59:
		//line parser.go.y:385
		{
			yyVAL.statement = data.GotoStmt{Pos: yyS[yypt-2].tok.pos, Label: yyS[yypt-1].tok.lit}
		}
	case 60:
		//line parser.go.y:389
		{
			yyVAL.statement = data.SkipStmt{Pos: yyS[yypt-1].tok.pos}
		}
	case 61:
		//line parser.go.y:393
		{
			yyVAL.statement = data.ExprStmt{Expr: yyS[yypt-1].expression}
		}
	case 62:
		//line parser.go.y:397
		{
			yyVAL.statement = data.NullStmt{Pos: yyS[yypt-0].tok.pos}
		}
	case 63:
		//line parser.go.y:401
		{
			yyVAL.statement = yyS[yypt-0].definition.(data.Stmt)
		}
	case 64:
		//line parser.go.y:406
		{
			yyVAL.expression = data.IdentifierExpr{Pos: yyS[yypt-0].tok.pos, Name: yyS[yypt-0].tok.lit}
		}
	case 65:
		//line parser.go.y:410
		{
			yyVAL.expression = data.NumberExpr{Pos: yyS[yypt-0].tok.pos, Lit: yyS[yypt-0].tok.lit}
		}
	case 66:
		//line parser.go.y:414
		{
			yyVAL.expression = data.TrueExpr{Pos: yyS[yypt-0].tok.pos}
		}
	case 67:
		//line parser.go.y:418
		{
			yyVAL.expression = data.FalseExpr{Pos: yyS[yypt-0].tok.pos}
		}
	case 68:
		//line parser.go.y:422
		{
			yyVAL.expression = data.NotExpr{Pos: yyS[yypt-1].tok.pos, SubExpr: yyS[yypt-0].expression}
		}
	case 69:
		//line parser.go.y:426
		{
			yyVAL.expression = data.UnarySubExpr{Pos: yyS[yypt-1].tok.pos, SubExpr: yyS[yypt-0].expression}
		}
	case 70:
		//line parser.go.y:430
		{
			yyVAL.expression = data.ParenExpr{Pos: yyS[yypt-2].tok.pos, SubExpr: yyS[yypt-1].expression}
		}
	case 71:
		//line parser.go.y:434
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "+", RHS: yyS[yypt-0].expression}
		}
	case 72:
		//line parser.go.y:438
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "-", RHS: yyS[yypt-0].expression}
		}
	case 73:
		//line parser.go.y:442
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "*", RHS: yyS[yypt-0].expression}
		}
	case 74:
		//line parser.go.y:446
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "/", RHS: yyS[yypt-0].expression}
		}
	case 75:
		//line parser.go.y:450
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "%", RHS: yyS[yypt-0].expression}
		}
	case 76:
		//line parser.go.y:454
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "&", RHS: yyS[yypt-0].expression}
		}
	case 77:
		//line parser.go.y:458
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "|", RHS: yyS[yypt-0].expression}
		}
	case 78:
		//line parser.go.y:462
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "^", RHS: yyS[yypt-0].expression}
		}
	case 79:
		//line parser.go.y:466
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "<<", RHS: yyS[yypt-0].expression}
		}
	case 80:
		//line parser.go.y:470
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: ">>", RHS: yyS[yypt-0].expression}
		}
	case 81:
		//line parser.go.y:474
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "&&", RHS: yyS[yypt-0].expression}
		}
	case 82:
		//line parser.go.y:478
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "||", RHS: yyS[yypt-0].expression}
		}
	case 83:
		//line parser.go.y:482
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "==", RHS: yyS[yypt-0].expression}
		}
	case 84:
		//line parser.go.y:486
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "<", RHS: yyS[yypt-0].expression}
		}
	case 85:
		//line parser.go.y:490
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: ">", RHS: yyS[yypt-0].expression}
		}
	case 86:
		//line parser.go.y:494
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "!=", RHS: yyS[yypt-0].expression}
		}
	case 87:
		//line parser.go.y:498
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: "<=", RHS: yyS[yypt-0].expression}
		}
	case 88:
		//line parser.go.y:502
		{
			yyVAL.expression = data.BinOpExpr{LHS: yyS[yypt-2].expression, Operator: ">=", RHS: yyS[yypt-0].expression}
		}
	case 89:
		//line parser.go.y:506
		{
			yyVAL.expression = data.TimeoutRecvExpr{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 90:
		//line parser.go.y:510
		{
			yyVAL.expression = data.TimeoutPeekExpr{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 91:
		//line parser.go.y:514
		{
			yyVAL.expression = data.NonblockRecvExpr{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 92:
		//line parser.go.y:518
		{
			yyVAL.expression = data.NonblockPeekExpr{Pos: yyS[yypt-3].tok.pos, Channel: yyS[yypt-1].expressions[0], Args: yyS[yypt-1].expressions[1:]}
		}
	case 93:
		//line parser.go.y:522
		{
			yyVAL.expression = data.ArrayExpr{Pos: yyS[yypt-2].tok.pos, Elems: yyS[yypt-1].expressions}
		}
	case 94:
		//line parser.go.y:529
		{
			yyVAL.ltlexpr = yyS[yypt-0].ltlatom
		}
	case 95:
		//line parser.go.y:533
		{
			yyVAL.ltlexpr = data.ParenLtlExpr{SubExpr: yyS[yypt-1].ltlexpr}
		}
	case 96:
		//line parser.go.y:537
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "&", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 97:
		//line parser.go.y:541
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "|", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 98:
		//line parser.go.y:545
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "^", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 99:
		//line parser.go.y:549
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "->", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 100:
		//line parser.go.y:553
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "=", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 101:
		//line parser.go.y:557
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "!=", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 102:
		//line parser.go.y:561
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "U", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 103:
		//line parser.go.y:565
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "V", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 104:
		//line parser.go.y:569
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "S", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 105:
		//line parser.go.y:573
		{
			yyVAL.ltlexpr = data.BinOpLtlExpr{Operator: "T", LHS: yyS[yypt-2].ltlexpr, RHS: yyS[yypt-0].ltlexpr}
		}
	case 106:
		//line parser.go.y:577
		{
			yyVAL.ltlexpr = data.UnOpLtlExpr{Operator: "!", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 107:
		//line parser.go.y:581
		{
			yyVAL.ltlexpr = data.UnOpLtlExpr{Operator: "X", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 108:
		//line parser.go.y:585
		{
			yyVAL.ltlexpr = data.UnOpLtlExpr{Operator: "G", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 109:
		//line parser.go.y:589
		{
			yyVAL.ltlexpr = data.UnOpLtlExpr{Operator: "F", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 110:
		//line parser.go.y:593
		{
			yyVAL.ltlexpr = data.UnOpLtlExpr{Operator: "Y", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 111:
		//line parser.go.y:597
		{
			yyVAL.ltlexpr = data.UnOpLtlExpr{Operator: "Z", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 112:
		//line parser.go.y:601
		{
			yyVAL.ltlexpr = data.UnOpLtlExpr{Operator: "H", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 113:
		//line parser.go.y:605
		{
			yyVAL.ltlexpr = data.UnOpLtlExpr{Operator: "O", SubExpr: yyS[yypt-0].ltlexpr}
		}
	case 114:
		//line parser.go.y:610
		{
			yyVAL.ltlatom = data.LtlAtomExpr{Names: []string{yyS[yypt-0].tok.lit}}
		}
	case 115:
		//line parser.go.y:614
		{
			yyVAL.ltlatom = data.LtlAtomExpr{Names: append([]string{yyS[yypt-2].tok.lit}, yyS[yypt-0].ltlatom.Names...)}
		}
	case 116:
		//line parser.go.y:622
		{
			yyVAL.identifiers = []string{yyS[yypt-0].tok.lit}
		}
	case 117:
		//line parser.go.y:626
		{
			yyVAL.identifiers = []string{yyS[yypt-1].tok.lit}
		}
	case 118:
		//line parser.go.y:630
		{
			yyVAL.identifiers = append([]string{yyS[yypt-2].tok.lit}, yyS[yypt-0].identifiers...)
		}
	case 119:
		//line parser.go.y:636
		{
			yyVAL.parameters = nil
		}
	case 120:
		//line parser.go.y:640
		{
			yyVAL.parameters = yyS[yypt-0].parameters
		}
	case 121:
		//line parser.go.y:646
		{
			yyVAL.parameters = []data.Parameter{yyS[yypt-0].parameter}
		}
	case 122:
		//line parser.go.y:650
		{
			yyVAL.parameters = []data.Parameter{yyS[yypt-1].parameter}
		}
	case 123:
		//line parser.go.y:654
		{
			yyVAL.parameters = append([]data.Parameter{yyS[yypt-2].parameter}, yyS[yypt-0].parameters...)
		}
	case 124:
		//line parser.go.y:660
		{
			yyVAL.parameter = data.Parameter{Name: yyS[yypt-1].tok.lit, Type: yyS[yypt-0].typetype}
		}
	case 125:
		//line parser.go.y:666
		{
			yyVAL.expressions = []data.Expr{yyS[yypt-0].expression}
		}
	case 126:
		//line parser.go.y:670
		{
			yyVAL.expressions = []data.Expr{yyS[yypt-1].expression}
		}
	case 127:
		//line parser.go.y:674
		{
			yyVAL.expressions = append([]data.Expr{yyS[yypt-2].expression}, yyS[yypt-0].expressions...)
		}
	case 128:
		//line parser.go.y:680
		{
			yyVAL.typetypes = []data.Type{yyS[yypt-0].typetype}
		}
	case 129:
		//line parser.go.y:684
		{
			yyVAL.typetypes = []data.Type{yyS[yypt-1].typetype}
		}
	case 130:
		//line parser.go.y:688
		{
			yyVAL.typetypes = append([]data.Type{yyS[yypt-2].typetype}, yyS[yypt-0].typetypes...)
		}
	case 131:
		//line parser.go.y:693
		{
			yyVAL.typetype = data.NamedType{Name: yyS[yypt-0].tok.lit}
		}
	case 132:
		//line parser.go.y:697
		{
			yyVAL.typetype = data.ArrayType{ElemType: yyS[yypt-0].typetype}
		}
	case 133:
		//line parser.go.y:701
		{
			yyVAL.typetype = data.HandshakeChannelType{Elems: yyS[yypt-1].typetypes}
		}
	case 134:
		//line parser.go.y:705
		{
			yyVAL.typetype = data.BufferedChannelType{BufferSize: nil, Elems: yyS[yypt-1].typetypes}
		}
	case 135:
		//line parser.go.y:709
		{
			yyVAL.typetype = data.BufferedChannelType{BufferSize: yyS[yypt-4].expression, Elems: yyS[yypt-1].typetypes}
		}
	case 136:
		//line parser.go.y:715
		{
			yyVAL.tags = nil
		}
	case 137:
		//line parser.go.y:719
		{
			yyVAL.tags = yyS[yypt-0].tags
		}
	case 138:
		//line parser.go.y:725
		{
			yyVAL.tags = []string{yyS[yypt-0].tag}
		}
	case 139:
		//line parser.go.y:729
		{
			yyVAL.tags = append([]string{yyS[yypt-1].tag}, yyS[yypt-0].tags...)
		}
	case 140:
		//line parser.go.y:735
		{
			yyVAL.tag = yyS[yypt-0].tok.lit
		}
	case 141:
		//line parser.go.y:741
		{
			yyVAL.blocks = []data.BlockStmt{data.BlockStmt{Pos: yyS[yypt-2].tok.pos, Stmts: yyS[yypt-1].statements}}
		}
	case 142:
		//line parser.go.y:745
		{
			yyVAL.blocks = []data.BlockStmt{data.BlockStmt{Pos: yyS[yypt-3].tok.pos, Stmts: yyS[yypt-2].statements}}
		}
	case 143:
		//line parser.go.y:749
		{
			yyVAL.blocks = append([]data.BlockStmt{data.BlockStmt{Pos: yyS[yypt-4].tok.pos, Stmts: yyS[yypt-3].statements}}, yyS[yypt-0].blocks...)
		}
	}
	goto yystack /* stack new state and value */
}
