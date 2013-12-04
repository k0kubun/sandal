// vim: noet sw=8 sts=8
%{
package parsing

import (
	"log"
	data "github.com/draftcode/sandal/lang/data"
)

type token struct {
	tok int
	lit string
	pos Position
}
%}

%union{
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
	blocks      []data.BlockStatement
	initvars    []data.InitVar
	initvar     data.InitVar

	tok         token
}

%type<definitions> spec
%type<definition> toplevel_body
%type<definition> data_def module_def const_def proc_def init_block
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
%type<blocks> blocks_one

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
%token<tok> UNSTABLE
%token<tok> SKIP

%left LOR
%left LAND
%left EQL NEQ LSS LEQ GTR GEQ
%left ADD SUB OR XOR
%left MUL QUO REM SHL SHR AND
%right UNARY

%%

spec	: toplevel_body
	{
		$$ = []data.Definition{$1}
		if l, isLexerWrapper := yylex.(*lexerWrapper); isLexerWrapper {
			l.definitions = $$
		}
	}
	| toplevel_body spec
	{
		$$ = append([]data.Definition{$1}, $2...)
		if l, isLexerWrapper := yylex.(*lexerWrapper); isLexerWrapper {
			l.definitions = $$
		}
	}

toplevel_body
	: data_def
	| module_def
	| const_def
	| proc_def
	| init_block

data_def
	: DATA IDENTIFIER '{' idents_one '}' ';'
	{
		$$ = data.DataDefinition{Name: $2.lit, Elems: $4}
	}

module_def
	: MODULE IDENTIFIER '(' parameters_zero ')' '{' module_body_zero '}' ';'
	{
		$$ = data.ModuleDefinition{Name: $2.lit, Parameters: $4, Definitions: $7}
	}

module_body_zero
	:
	{
		$$ = nil
	}
	| module_body module_body_zero
	{
		$$ = append([]data.Definition{$1}, $2...)
	}

module_body
	: const_def
	| proc_def
	| init_block

const_def
	: CONST IDENTIFIER type ASSIGN expr ';' /* This should be a const expression. */
	{
		$$ = data.ConstantDefinition{Name: $2.lit, Type: $3, Expr: $5}
	}

proc_def
	: PROC IDENTIFIER '(' parameters_zero ')' '{' statements_zero '}' ';'
	{
		$$ = data.ProcDefinition{Name: $2.lit, Parameters: $4, Statements: $7}
	}

init_block
	: INIT '{' initvars_zero '}' ';'
	{
		$$ = data.InitBlock{Vars: $3}
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

initvar	: IDENTIFIER ':' type
	{
		$$ = data.ChannelVar{Name: $1.lit, Type: $3}
	}
	| IDENTIFIER ':' IDENTIFIER '(' arguments_one ')'
	{
		$$ = data.InstanceVar{Name: $1.lit, ProcDefName: $3.lit, Args: $5}
	}

statements_zero
	:
	{
		$$ = nil
	}
	| statement statements_zero
	{
		$$ = append([]data.Statement{$1}, $2...)
	}

statement
	: IDENTIFIER ':' statement /* no semicolon */
	{
		$$ = data.LabelledStatement{Label: $1.lit, Statement: $3}
	}
	| '{' statements_zero '}' ';'
	{
		$$ = data.BlockStatement{Statements: $2}
	}
	| VAR IDENTIFIER type ';'
	{
		$$ = data.VarDeclStatement{Name: $2.lit, Type: $3}
	}
	| VAR IDENTIFIER type ASSIGN expr ';'
	{
		$$ = data.VarDeclStatement{Name: $2.lit, Type: $3, Initializer: $5}
	}
	| IF expr '{' statements_zero '}' ';'
	{
		$$ = data.IfStatement{Condition: $2, TrueBranch: $4}
	}
	| IF expr '{' statements_zero '}' ELSE '{' statements_zero '}' ';'
	{
		$$ = data.IfStatement{Condition: $2, TrueBranch: $4, FalseBranch: $8}
	}
	| IDENTIFIER ASSIGN expr ';'
	{
		$$ = data.AssignmentStatement{Variable: $1.lit, Expr: $3}
	}
	| IDENTIFIER ADD_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStatement{Variable: $1.lit, Operator: "+", Expr: $3}
	}
	| IDENTIFIER SUB_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStatement{Variable: $1.lit, Operator: "-", Expr: $3}
	}
	| IDENTIFIER MUL_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStatement{Variable: $1.lit, Operator: "*", Expr: $3}
	}
	| IDENTIFIER QUO_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStatement{Variable: $1.lit, Operator: "/", Expr: $3}
	}
	| IDENTIFIER REM_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStatement{Variable: $1.lit, Operator: "%", Expr: $3}
	}
	| IDENTIFIER AND_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStatement{Variable: $1.lit, Operator: "&", Expr: $3}
	}
	| IDENTIFIER OR_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStatement{Variable: $1.lit, Operator: "|", Expr: $3}
	}
	| IDENTIFIER XOR_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStatement{Variable: $1.lit, Operator: "^", Expr: $3}
	}
	| IDENTIFIER SHL_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStatement{Variable: $1.lit, Operator: "<<", Expr: $3}
	}
	| IDENTIFIER SHR_ASSIGN expr ';'
	{
		$$ = data.OpAssignmentStatement{Variable: $1.lit, Operator: ">>", Expr: $3}
	}
	| CHOICE blocks_one ';'
	{
		$$ = data.ChoiceStatement{Blocks: $2}
	}
	| RECV '(' arguments_one ')' ';'
	{
		$$ = data.RecvStatement{Channel: $3[0], Args: $3[1:]}
	}
	| PEEK '(' arguments_one ')' ';'
	{
		$$ = data.PeekStatement{Channel: $3[0], Args: $3[1:]}
	}
	| SEND '(' arguments_one ')' ';'
	{
		$$ = data.SendStatement{Channel: $3[0], Args: $3[1:]}
	}
	| FOR '{' statements_zero '}' ';'
	{
		$$ = data.ForStatement{Statements: $3}
	}
	| FOR IDENTIFIER IN expr '{' statements_zero '}' ';'
	{
		$$ = data.ForInStatement{Variable: $2.lit, Container: $4, Statements: $6}
	}
	| FOR IDENTIFIER IN RANGE expr TO expr '{' statements_zero '}' ';'
	{
		$$ = data.ForInRangeStatement{Variable: $2.lit, FromExpr: $5, ToExpr: $7, Statements: $9}
	}
	| BREAK ';'
	{
		$$ = data.BreakStatement{}
	}
	| GOTO IDENTIFIER ';'
	{
		$$ = data.GotoStatement{Label: $2.lit}
	}
	| SKIP ';'
	{
		$$ = data.SkipStatement{}
	}
	| expr ';'
	{
		$$ = data.ExprStatement{Expr: $1}
	}
	| ';'
	{
		$$ = data.NullStatement{}
	}
	| const_def
	{
		$$ = $1.(data.Statement)
	}

expr	: IDENTIFIER
	{
		$$ = data.IdentifierExpression{Name: $1.lit}
	}
	| NUMBER
	{
		$$ = data.NumberExpression{Lit: $1.lit}
	}
	| NOT expr      %prec UNARY
	{
		$$ = data.NotExpression{SubExpr: $2}
	}
	| SUB expr      %prec UNARY
	{
		$$ = data.UnarySubExpression{SubExpr: $2}
	}
	| '(' expr ')'
	{
		$$ = data.ParenExpression{SubExpr: $2}
	}
	| expr ADD expr
	{
		$$ = data.BinOpExpression{LHS: $1, Operator: "+", RHS: $3}
	}
	| expr SUB expr
	{
		$$ = data.BinOpExpression{LHS: $1, Operator: "-", RHS: $3}
	}
	| expr MUL expr
	{
		$$ = data.BinOpExpression{LHS: $1, Operator: "*", RHS: $3}
	}
	| expr QUO expr
	{
		$$ = data.BinOpExpression{LHS: $1, Operator: "/", RHS: $3}
	}
	| expr REM expr
	{
		$$ = data.BinOpExpression{LHS: $1, Operator: "%", RHS: $3}
	}
	| expr AND expr
	{
		$$ = data.BinOpExpression{LHS: $1, Operator: "&", RHS: $3}
	}
	| expr OR expr
	{
		$$ = data.BinOpExpression{LHS: $1, Operator: "|", RHS: $3}
	}
	| expr XOR expr
	{
		$$ = data.BinOpExpression{LHS: $1, Operator: "^", RHS: $3}
	}
	| expr SHL expr
	{
		$$ = data.BinOpExpression{LHS: $1, Operator: "<<", RHS: $3}
	}
	| expr SHR expr
	{
		$$ = data.BinOpExpression{LHS: $1, Operator: ">>", RHS: $3}
	}
	| expr LAND expr
	{
		$$ = data.BinOpExpression{LHS: $1, Operator: "&&", RHS: $3}
	}
	| expr LOR expr
	{
		$$ = data.BinOpExpression{LHS: $1, Operator: "||", RHS: $3}
	}
	| expr EQL expr
	{
		$$ = data.BinOpExpression{LHS: $1, Operator: "==", RHS: $3}
	}
	| expr LSS expr
	{
		$$ = data.BinOpExpression{LHS: $1, Operator: "<", RHS: $3}
	}
	| expr GTR expr
	{
		$$ = data.BinOpExpression{LHS: $1, Operator: ">", RHS: $3}
	}
	| expr NEQ expr
	{
		$$ = data.BinOpExpression{LHS: $1, Operator: "!=", RHS: $3}
	}
	| expr LEQ expr
	{
		$$ = data.BinOpExpression{LHS: $1, Operator: "<=", RHS: $3}
	}
	| expr GEQ expr
	{
		$$ = data.BinOpExpression{LHS: $1, Operator: ">=", RHS: $3}
	}
	| TIMEOUT_RECV '(' arguments_one ')'
	{
		$$ = data.TimeoutRecvExpression{Channel: $3[0], Args: $3[1:]}
	}
	| TIMEOUT_PEEK '(' arguments_one ')'
	{
		$$ = data.TimeoutPeekExpression{Channel: $3[0], Args: $3[1:]}
	}
	| NONBLOCK_RECV '(' arguments_one ')'
	{
		$$ = data.NonblockRecvExpression{Channel: $3[0], Args: $3[1:]}
	}
	| NONBLOCK_PEEK '(' arguments_one ')'
	{
		$$ = data.NonblockPeekExpression{Channel: $3[0], Args: $3[1:]}
	}
	| '[' arguments_one ']'
	{
		$$ = data.ArrayExpression{Elems: $2}
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
		$$ = []data.Expression{$1}
	}
	| expr ','
	{
		$$ = []data.Expression{$1}
	}
	| expr ',' arguments_one
	{
		$$ = append([]data.Expression{$1}, $3...)
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
		$$ = data.HandshakeChannelType{IsUnstable: false, Elems: $3}
	}
	| UNSTABLE CHANNEL '{' types_one '}'
	{
		$$ = data.HandshakeChannelType{IsUnstable: true, Elems: $4}
	}
	| CHANNEL '[' ']' '{' types_one '}'
	{
		$$ = data.BufferedChannelType{IsUnstable: false, BufferSize: nil, Elems: $5}
	}
	| CHANNEL '[' expr ']' '{' types_one '}'
	{
		$$ = data.BufferedChannelType{IsUnstable: false, BufferSize: $3, Elems: $6}
	}
	| UNSTABLE CHANNEL '[' ']' '{' types_one '}'
	{
		$$ = data.BufferedChannelType{IsUnstable: true, BufferSize: nil, Elems: $6}
	}
	| UNSTABLE CHANNEL '[' expr ']' '{' types_one '}'
	{
		$$ = data.BufferedChannelType{IsUnstable: true, BufferSize: $4, Elems: $7}
	}

blocks_one
	: '{' statements_zero '}'
	{
		$$ = []data.BlockStatement{data.BlockStatement{Statements: $2}}
	}
	| '{' statements_zero '}' ','
	{
		$$ = []data.BlockStatement{data.BlockStatement{Statements: $2}}
	}
	| '{' statements_zero '}' ',' blocks_one
	{
		$$ = append([]data.BlockStatement{data.BlockStatement{Statements: $2}}, $5...)
	}

%%

type lexerWrapper struct {
	s           *Scanner
	definitions []data.Definition
	recentLit   string
	recentPos   Position
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
