// vim: noet sw=8 sts=8
%{
package parsing

import (
	"log"
	data "github.com/k0kubun/sandal/lang/data"
)

type token struct {
	tok int
	lit string
	pos data.Pos
}
%}

%union{
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

	tok         token
}

%type<definitions> spec
%type<definition> toplevel_body
%type<definition> data_def module_def const_def proc_def fault_def init_block ltl_spec
%type<definitions> module_body_zero
%type<definition> module_body
%type<initvars> initvars_zero initvars_one
%type<initvar> initvar
%type<statements> statements_zero
%type<statement> statement
%type<expression> expr

%type<identifiers> idents_one
%type<parameters> parameters_zero parameters_one
%type<parameter> parameter
%type<expressions> arguments_one
%type<typetypes> types_one
%type<typetype> type
%type<tags> tags_zero tags_one
%type<tag> tag
%type<blocks> blocks_one
%type<ltlexpr> ltl_expr
%type<ltlatom> ltl_atom

%token<tok> IDENTIFIER
%token<tok> NUMBER
%token<tok> COMMENT

%token<tok> ADD // +
%token<tok> SUB // -
%token<tok> MUL // *
%token<tok> QUO // /
%token<tok> REM // %

%token<tok> AND // &
%token<tok> OR  // |
%token<tok> XOR // ^
%token<tok> SHL // <<
%token<tok> SHR // >>

%token<tok> ADD_ASSIGN // +=
%token<tok> SUB_ASSIGN // -=
%token<tok> MUL_ASSIGN // *=
%token<tok> QUO_ASSIGN // /=
%token<tok> REM_ASSIGN // %=

%token<tok> AND_ASSIGN // &=
%token<tok> OR_ASSIGN  // |=
%token<tok> XOR_ASSIGN // ^=
%token<tok> SHL_ASSIGN // <<=
%token<tok> SHR_ASSIGN // >>=

%token<tok> LAND // &&
%token<tok> LOR  // ||

%token<tok> EQL    // ==
%token<tok> LSS    // <
%token<tok> GTR    // >
%token<tok> ASSIGN // =
%token<tok> NOT    // !

%token<tok> NEQ // !=
%token<tok> LEQ // <=
%token<tok> GEQ // >=

%token<tok> DATA
%token<tok> CONST
%token<tok> MODULE
%token<tok> CHANNEL
%token<tok> PROC
%token<tok> FAULT
%token<tok> VAR
%token<tok> IF
%token<tok> ELSE
%token<tok> CHOICE
%token<tok> RECV
%token<tok> TIMEOUT_RECV
%token<tok> NONBLOCK_RECV
%token<tok> PEEK
%token<tok> TIMEOUT_PEEK
%token<tok> NONBLOCK_PEEK
%token<tok> SEND
%token<tok> FOR
%token<tok> BREAK
%token<tok> IN
%token<tok> RANGE
%token<tok> TO
%token<tok> INIT
%token<tok> GOTO
%token<tok> SKIP
%token<tok> TRUE
%token<tok> FALSE
%token<tok> LTL
%token<tok> THEN
%token<tok> IFF
%token<tok> '{' '}' '(' ')' '[' ']' ',' ':' ';'

%left LOR
%left LAND
%left EQL NEQ LSS LEQ GTR GEQ
%left ADD SUB OR XOR
%left MUL QUO REM SHL SHR AND
%left THEN 'U' 'V' 'S' 'T'
%right UNARY

%%

spec	: toplevel_body
	{
		$$ = []data.Def{$1}
		if l, isLexerWrapper := yylex.(*lexerWrapper); isLexerWrapper {
			l.definitions = $$
		}
	}
	| toplevel_body spec
	{
		$$ = append([]data.Def{$1}, $2...)
		if l, isLexerWrapper := yylex.(*lexerWrapper); isLexerWrapper {
			l.definitions = $$
		}
	}

toplevel_body
	: data_def
	| module_def
	| const_def
	| proc_def
	| fault_def
	| init_block
	| ltl_spec

data_def
	: DATA IDENTIFIER '{' idents_one '}' ';'
	{
		$$ = data.DataDef{Pos: $1.pos, Name: $2.lit, Elems: $4}
	}

module_def
	: MODULE IDENTIFIER '(' parameters_zero ')' '{' module_body_zero '}' ';'
	{
		$$ = data.ModuleDef{Pos: $1.pos, Name: $2.lit, Parameters: $4, Defs: $7}
	}

module_body_zero
	:
	{
		$$ = nil
	}
	| module_body module_body_zero
	{
		$$ = append([]data.Def{$1}, $2...)
	}

module_body
	: const_def
	| proc_def
	| init_block

const_def
	: CONST IDENTIFIER type ASSIGN expr ';' /* This should be a const expression. */
	{
		$$ = data.ConstantDef{Pos: $1.pos, Name: $2.lit, Type: $3, Expr: $5}
	}

proc_def
	: PROC IDENTIFIER '(' parameters_zero ')' '{' statements_zero '}' ';'
	{
		$$ = data.ProcDef{Pos: $1.pos, Name: $2.lit, Parameters: $4, Stmts: $7}
	}

fault_def
	: FAULT SEND '(' parameters_one ')' tag '{' statements_zero '}' ';'
	{
		$$ = data.FaultDef{Pos: $1.pos, Name: $2.lit, Parameters: $4, Tag: $6, Stmts: $8}
	}
	| FAULT RECV '(' parameters_one ')' tag '{' statements_zero '}' ';'
	{
		$$ = data.FaultDef{Pos: $1.pos, Name: $2.lit, Parameters: $4, Tag: $6, Stmts: $8}
	}

init_block
	: INIT '{' initvars_zero '}' ';'
	{
		$$ = data.InitBlock{Pos: $1.pos, Vars: $3}
	}

ltl_spec
	: LTL '{' ltl_expr ';' '}' ';'
	{
		$$ = data.LtlSpec{Expr: $3}
	}
	| LTL '{' ltl_expr '}' ';'
	{
		$$ = data.LtlSpec{Expr: $3}
	}

initvars_zero
	:
	{
		$$ = nil
	}
	| initvars_one
	{
		$$ = $1
	}

initvars_one
	: initvar
	{
		$$ = []data.InitVar{$1}
	}
	| initvar ','
	{
		$$ = []data.InitVar{$1}
	}
	| initvar ',' initvars_one
	{
		$$ = append([]data.InitVar{$1}, $3...)
	}

initvar	: IDENTIFIER ':' type tags_zero
	{
		$$ = data.ChannelVar{Pos: $1.pos, Name: $1.lit, Type: $3, Tags: $4}
	}
	| IDENTIFIER ':' IDENTIFIER '(' arguments_one ')' tags_zero
	{
		$$ = data.InstanceVar{Pos: $1.pos, Name: $1.lit, ProcDefName: $3.lit, Args: $5, Tags: $7}
	}

statements_zero
	:
	{
		$$ = nil
	}
	| statement statements_zero
	{
		$$ = append([]data.Stmt{$1}, $2...)
	}

statement
	: IDENTIFIER ':' statement /* no semicolon */
	{
		$$ = data.LabelledStmt{Pos: $1.pos, Label: $1.lit, Stmt: $3}
	}
	| '{' statements_zero '}' ';'
	{
		$$ = data.BlockStmt{Pos: $1.pos, Stmts: $2}
	}
	| VAR IDENTIFIER type ';'
	{
		$$ = data.VarDeclStmt{Pos: $1.pos, Name: $2.lit, Type: $3}
	}
	| VAR IDENTIFIER type ASSIGN expr ';'
	{
		$$ = data.VarDeclStmt{Pos: $1.pos, Name: $2.lit, Type: $3, Initializer: $5}
	}
	| IF expr '{' statements_zero '}' ';'
	{
		$$ = data.IfStmt{Pos: $1.pos, Condition: $2, TrueBranch: $4}
	}
	| IF expr '{' statements_zero '}' ELSE '{' statements_zero '}' ';'
	{
		$$ = data.IfStmt{Pos: $1.pos, Condition: $2, TrueBranch: $4, FalseBranch: $8}
	}
	| IDENTIFIER ASSIGN expr ';'
	{
		$$ = data.AssignmentStmt{Pos: $1.pos, Variable: $1.lit, Expr: $3}
	}
	| IDENTIFIER ADD_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStmt{Pos: $1.pos, Variable: $1.lit, Operator: "+", Expr: $3}
	}
	| IDENTIFIER SUB_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStmt{Pos: $1.pos, Variable: $1.lit, Operator: "-", Expr: $3}
	}
	| IDENTIFIER MUL_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStmt{Pos: $1.pos, Variable: $1.lit, Operator: "*", Expr: $3}
	}
	| IDENTIFIER QUO_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStmt{Pos: $1.pos, Variable: $1.lit, Operator: "/", Expr: $3}
	}
	| IDENTIFIER REM_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStmt{Pos: $1.pos, Variable: $1.lit, Operator: "%", Expr: $3}
	}
	| IDENTIFIER AND_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStmt{Pos: $1.pos, Variable: $1.lit, Operator: "&", Expr: $3}
	}
	| IDENTIFIER OR_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStmt{Pos: $1.pos, Variable: $1.lit, Operator: "|", Expr: $3}
	}
	| IDENTIFIER XOR_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStmt{Pos: $1.pos, Variable: $1.lit, Operator: "^", Expr: $3}
	}
	| IDENTIFIER SHL_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStmt{Pos: $1.pos, Variable: $1.lit, Operator: "<<", Expr: $3}
	}
	| IDENTIFIER SHR_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStmt{Pos: $1.pos, Variable: $1.lit, Operator: ">>", Expr: $3}
	}
	| CHOICE blocks_one ';'
	{
		$$ = data.ChoiceStmt{Pos: $1.pos, Blocks: $2}
	}
	| RECV '(' arguments_one ')' tags_zero ';'
	{
		$$ = data.RecvStmt{Pos: $1.pos, Channel: $3[0], Args: $3[1:], Tags: $5}
	}
	| PEEK '(' arguments_one ')' ';'
	{
		$$ = data.PeekStmt{Pos: $1.pos, Channel: $3[0], Args: $3[1:]}
	}
	| SEND '(' arguments_one ')' tags_zero ';'
	{
		$$ = data.SendStmt{Pos: $1.pos, Channel: $3[0], Args: $3[1:], Tags: $5}
	}
	| FOR '{' statements_zero '}' ';'
	{
		$$ = data.ForStmt{Pos: $1.pos, Stmts: $3}
	}
	| FOR IDENTIFIER IN expr '{' statements_zero '}' ';'
	{
		$$ = data.ForInStmt{Pos: $1.pos, Variable: $2.lit, Container: $4, Stmts: $6}
	}
	| FOR IDENTIFIER IN RANGE expr TO expr '{' statements_zero '}' ';'
	{
		$$ = data.ForInRangeStmt{Pos: $1.pos, Variable: $2.lit, FromExpr: $5, ToExpr: $7, Stmts: $9}
	}
	| BREAK ';'
	{
		$$ = data.BreakStmt{Pos: $1.pos}
	}
	| GOTO IDENTIFIER ';'
	{
		$$ = data.GotoStmt{Pos: $1.pos, Label: $2.lit}
	}
	| SKIP ';'
	{
		$$ = data.SkipStmt{Pos: $1.pos}
	}
	| expr ';'
	{
		$$ = data.ExprStmt{Expr: $1}
	}
	| ';'
	{
		$$ = data.NullStmt{Pos: $1.pos}
	}
	| const_def
	{
		$$ = $1.(data.Stmt)
	}

expr	: IDENTIFIER
	{
		$$ = data.IdentifierExpr{Pos: $1.pos, Name: $1.lit}
	}
	| NUMBER
	{
		$$ = data.NumberExpr{Pos: $1.pos, Lit: $1.lit}
	}
	| TRUE
	{
		$$ = data.TrueExpr{Pos: $1.pos}
	}
	| FALSE
	{
		$$ = data.FalseExpr{Pos: $1.pos}
	}
	| NOT expr      %prec UNARY
	{
		$$ = data.NotExpr{Pos: $1.pos, SubExpr: $2}
	}
	| SUB expr      %prec UNARY
	{
		$$ = data.UnarySubExpr{Pos: $1.pos, SubExpr: $2}
	}
	| '(' expr ')'
	{
		$$ = data.ParenExpr{Pos: $1.pos, SubExpr: $2}
	}
	| expr ADD expr
	{
		$$ = data.BinOpExpr{LHS: $1, Operator: "+", RHS: $3}
	}
	| expr SUB expr
	{
		$$ = data.BinOpExpr{LHS: $1, Operator: "-", RHS: $3}
	}
	| expr MUL expr
	{
		$$ = data.BinOpExpr{LHS: $1, Operator: "*", RHS: $3}
	}
	| expr QUO expr
	{
		$$ = data.BinOpExpr{LHS: $1, Operator: "/", RHS: $3}
	}
	| expr REM expr
	{
		$$ = data.BinOpExpr{LHS: $1, Operator: "%", RHS: $3}
	}
	| expr AND expr
	{
		$$ = data.BinOpExpr{LHS: $1, Operator: "&", RHS: $3}
	}
	| expr OR expr
	{
		$$ = data.BinOpExpr{LHS: $1, Operator: "|", RHS: $3}
	}
	| expr XOR expr
	{
		$$ = data.BinOpExpr{LHS: $1, Operator: "^", RHS: $3}
	}
	| expr SHL expr
	{
		$$ = data.BinOpExpr{LHS: $1, Operator: "<<", RHS: $3}
	}
	| expr SHR expr
	{
		$$ = data.BinOpExpr{LHS: $1, Operator: ">>", RHS: $3}
	}
	| expr LAND expr
	{
		$$ = data.BinOpExpr{LHS: $1, Operator: "&&", RHS: $3}
	}
	| expr LOR expr
	{
		$$ = data.BinOpExpr{LHS: $1, Operator: "||", RHS: $3}
	}
	| expr EQL expr
	{
		$$ = data.BinOpExpr{LHS: $1, Operator: "==", RHS: $3}
	}
	| expr LSS expr
	{
		$$ = data.BinOpExpr{LHS: $1, Operator: "<", RHS: $3}
	}
	| expr GTR expr
	{
		$$ = data.BinOpExpr{LHS: $1, Operator: ">", RHS: $3}
	}
	| expr NEQ expr
	{
		$$ = data.BinOpExpr{LHS: $1, Operator: "!=", RHS: $3}
	}
	| expr LEQ expr
	{
		$$ = data.BinOpExpr{LHS: $1, Operator: "<=", RHS: $3}
	}
	| expr GEQ expr
	{
		$$ = data.BinOpExpr{LHS: $1, Operator: ">=", RHS: $3}
	}
	| TIMEOUT_RECV '(' arguments_one ')'
	{
		$$ = data.TimeoutRecvExpr{Pos: $1.pos, Channel: $3[0], Args: $3[1:]}
	}
	| TIMEOUT_PEEK '(' arguments_one ')'
	{
		$$ = data.TimeoutPeekExpr{Pos: $1.pos, Channel: $3[0], Args: $3[1:]}
	}
	| NONBLOCK_RECV '(' arguments_one ')'
	{
		$$ = data.NonblockRecvExpr{Pos: $1.pos, Channel: $3[0], Args: $3[1:]}
	}
	| NONBLOCK_PEEK '(' arguments_one ')'
	{
		$$ = data.NonblockPeekExpr{Pos: $1.pos, Channel: $3[0], Args: $3[1:]}
	}
	| '[' arguments_one ']'
	{
		$$ = data.ArrayExpr{Pos: $1.pos, Elems: $2}
	}

/* ======================================== */

ltl_expr: ltl_atom
	{
		$$ = $1
	}
	| '(' ltl_expr ')'
	{
		$$ = data.ParenLtlExpr{SubExpr: $2}
	}
	| ltl_expr LAND ltl_expr
	{
		$$ = data.BinOpLtlExpr{Operator: "&", LHS: $1, RHS: $3}
	}
	| ltl_expr LOR ltl_expr
	{
		$$ = data.BinOpLtlExpr{Operator: "|", LHS: $1, RHS: $3}
	}
	| ltl_expr XOR ltl_expr
	{
		$$ = data.BinOpLtlExpr{Operator: "^", LHS: $1, RHS: $3}
	}
	| ltl_expr THEN ltl_expr
	{
		$$ = data.BinOpLtlExpr{Operator: "->", LHS: $1, RHS: $3}
	}
	| ltl_expr EQL ltl_expr
	{
		$$ = data.BinOpLtlExpr{Operator: "=", LHS: $1, RHS: $3}
	}
	| ltl_expr 'U' ltl_expr
	{
		$$ = data.BinOpLtlExpr{Operator: "U", LHS: $1, RHS: $3}
	}
	| ltl_expr 'V' ltl_expr
	{
		$$ = data.BinOpLtlExpr{Operator: "V", LHS: $1, RHS: $3}
	}
	| ltl_expr 'S' ltl_expr
	{
		$$ = data.BinOpLtlExpr{Operator: "S", LHS: $1, RHS: $3}
	}
	| ltl_expr 'T' ltl_expr
	{
		$$ = data.BinOpLtlExpr{Operator: "T", LHS: $1, RHS: $3}
	}
	| NOT ltl_expr      %prec UNARY
	{
		$$ = data.UnOpLtlExpr{Operator: "!", SubExpr: $2}
	}
	| 'X' ltl_expr      %prec UNARY
	{
		$$ = data.UnOpLtlExpr{Operator: "X", SubExpr: $2}
	}
	| 'G' ltl_expr      %prec UNARY
	{
		$$ = data.UnOpLtlExpr{Operator: "G", SubExpr: $2}
	}
	| 'F' ltl_expr      %prec UNARY
	{
		$$ = data.UnOpLtlExpr{Operator: "F", SubExpr: $2}
	}
	| 'Y' ltl_expr      %prec UNARY
	{
		$$ = data.UnOpLtlExpr{Operator: "Y", SubExpr: $2}
	}
	| 'Z' ltl_expr      %prec UNARY
	{
		$$ = data.UnOpLtlExpr{Operator: "Z", SubExpr: $2}
	}
	| 'H' ltl_expr      %prec UNARY
	{
		$$ = data.UnOpLtlExpr{Operator: "H", SubExpr: $2}
	}
	| 'O' ltl_expr      %prec UNARY
	{
		$$ = data.UnOpLtlExpr{Operator: "O", SubExpr: $2}
	}

ltl_atom: IDENTIFIER
	{
		$$ = data.LtlAtomExpr{Names: []string{$1.lit}}
	}
	| IDENTIFIER '.' ltl_atom
	{
		$$ = data.LtlAtomExpr{Names: append([]string{$1.lit}, $3.Names...)}
	}

/* ======================================== */

idents_one
	: IDENTIFIER
	{
		$$ = []string{$1.lit}
	}
	| IDENTIFIER ','
	{
		$$ = []string{$1.lit}
	}
	| IDENTIFIER ',' idents_one
	{
		$$ = append([]string{$1.lit}, $3...)
	}

parameters_zero
	:
	{
		$$ = nil
	}
	| parameters_one
	{
		$$ = $1
	}

parameters_one
	: parameter
	{
		$$ = []data.Parameter{$1}
	}
	| parameter ','
	{
		$$ = []data.Parameter{$1}
	}
	| parameter ',' parameters_one
	{
		$$ = append([]data.Parameter{$1}, $3...)
	}

parameter
	: IDENTIFIER type
	{
		$$ = data.Parameter{Name: $1.lit, Type: $2}
	}

arguments_one
	: expr
	{
		$$ = []data.Expr{$1}
	}
	| expr ','
	{
		$$ = []data.Expr{$1}
	}
	| expr ',' arguments_one
	{
		$$ = append([]data.Expr{$1}, $3...)
	}

types_one
	: type
	{
		$$ = []data.Type{$1}
	}
	| type ','
	{
		$$ = []data.Type{$1}
	}
	| type ',' types_one
	{
		$$ = append([]data.Type{$1}, $3...)
	}

type	: IDENTIFIER
	{
		$$ = data.NamedType{Name: $1.lit}
	}
	| '[' ']' type
	{
		$$ = data.ArrayType{ElemType: $3}
	}
	| CHANNEL '{' types_one '}'
	{
		$$ = data.HandshakeChannelType{Elems: $3}
	}
	| CHANNEL '[' ']' '{' types_one '}'
	{
		$$ = data.BufferedChannelType{BufferSize: nil, Elems: $5}
	}
	| CHANNEL '[' expr ']' '{' types_one '}'
	{
		$$ = data.BufferedChannelType{BufferSize: $3, Elems: $6}
	}

tags_zero
	:
	{
		$$ = nil
	}
	| tags_one
	{
		$$ = $1
	}

tags_one
	: tag
	{
		$$ = []string{$1}
	}
	| tag tags_one
	{
		$$ = append([]string{$1}, $2...)
	}

tag
	: '@' IDENTIFIER
	{
		$$ = $2.lit
	}

blocks_one
	: '{' statements_zero '}'
	{
		$$ = []data.BlockStmt{data.BlockStmt{Pos: $1.pos, Stmts: $2}}
	}
	| '{' statements_zero '}' ','
	{
		$$ = []data.BlockStmt{data.BlockStmt{Pos: $1.pos, Stmts: $2}}
	}
	| '{' statements_zero '}' ',' blocks_one
	{
		$$ = append([]data.BlockStmt{data.BlockStmt{Pos: $1.pos, Stmts: $2}}, $5...)
	}

%%

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
