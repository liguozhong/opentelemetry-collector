// Code generated by goyacc -p expr -o pkg/logql/expr.y.go pkg/logql/expr.y. DO NOT EDIT.

//line pkg/logql/expr.y:2
package logql

import __yyfmt__ "fmt"

//line pkg/logql/expr.y:2

import (
	"github.com/prometheus/prometheus/pkg/labels"
	"time"
)

//line pkg/logql/expr.y:10
type exprSymType struct {
	yys                   int
	Expr                  Expr
	Filter                labels.MatchType
	Grouping              *grouping
	Labels                []string
	LogExpr               LogSelectorExpr
	LogRangeExpr          *logRange
	Matcher               *labels.Matcher
	Matchers              []*labels.Matcher
	RangeAggregationExpr  SampleExpr
	RangeOp               string
	Selector              []*labels.Matcher
	VectorAggregationExpr SampleExpr
	VectorOp              string
	str                   string
	duration              time.Duration
	int                   int64
}

const IDENTIFIER = 57346
const STRING = 57347
const DURATION = 57348
const MATCHERS = 57349
const LABELS = 57350
const EQ = 57351
const NEQ = 57352
const RE = 57353
const NRE = 57354
const OPEN_BRACE = 57355
const CLOSE_BRACE = 57356
const OPEN_BRACKET = 57357
const CLOSE_BRACKET = 57358
const COMMA = 57359
const DOT = 57360
const PIPE_MATCH = 57361
const PIPE_EXACT = 57362
const OPEN_PARENTHESIS = 57363
const CLOSE_PARENTHESIS = 57364
const BY = 57365
const WITHOUT = 57366
const COUNT_OVER_TIME = 57367
const RATE = 57368
const SUM = 57369
const AVG = 57370
const MAX = 57371
const MIN = 57372
const COUNT = 57373
const STDDEV = 57374
const STDVAR = 57375
const BOTTOMK = 57376
const TOPK = 57377

var exprToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENTIFIER",
	"STRING",
	"DURATION",
	"MATCHERS",
	"LABELS",
	"EQ",
	"NEQ",
	"RE",
	"NRE",
	"OPEN_BRACE",
	"CLOSE_BRACE",
	"OPEN_BRACKET",
	"CLOSE_BRACKET",
	"COMMA",
	"DOT",
	"PIPE_MATCH",
	"PIPE_EXACT",
	"OPEN_PARENTHESIS",
	"CLOSE_PARENTHESIS",
	"BY",
	"WITHOUT",
	"COUNT_OVER_TIME",
	"RATE",
	"SUM",
	"AVG",
	"MAX",
	"MIN",
	"COUNT",
	"STDDEV",
	"STDVAR",
	"BOTTOMK",
	"TOPK",
}
var exprStatenames = [...]string{}

const exprEofCode = 1
const exprErrCode = 2
const exprInitialStackSize = 16

//line pkg/logql/expr.y:140

//line yacctab:1
var exprExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 3,
	1, 2,
	-2, 0,
}

const exprPrivate = 57344

const exprLast = 123

var exprAct = [...]int{

	31, 5, 4, 36, 64, 10, 30, 78, 32, 33,
	32, 33, 80, 7, 45, 82, 81, 11, 12, 13,
	14, 16, 17, 15, 18, 19, 20, 21, 77, 78,
	76, 60, 44, 43, 79, 11, 12, 13, 14, 16,
	17, 15, 18, 19, 20, 21, 3, 59, 63, 62,
	57, 10, 48, 66, 28, 67, 47, 46, 29, 7,
	72, 73, 61, 74, 75, 11, 12, 13, 14, 16,
	17, 15, 18, 19, 20, 21, 42, 23, 50, 52,
	9, 39, 84, 85, 38, 27, 71, 26, 23, 70,
	49, 23, 58, 51, 24, 25, 27, 40, 26, 27,
	69, 26, 68, 83, 37, 24, 25, 65, 24, 25,
	53, 54, 55, 56, 35, 6, 37, 8, 34, 41,
	22, 2, 1,
}
var exprPact = [...]int{

	-8, -1000, -1000, 89, -1000, -1000, -1000, 38, 37, -15,
	112, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 79, -1000, -1000, -1000, -1000, -1000, 75, 38,
	10, 36, 35, 31, 76, 65, -1000, 101, -1000, -1000,
	-1000, 28, 86, 25, 9, 45, 40, 103, 103, -1000,
	-1000, 100, -1000, 97, 95, 84, 81, -1000, -1000, -13,
	-13, 40, 8, 6, 12, -1000, -10, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -6, -7, -1000, -1000, 99, -1000,
	-1000, -13, -13, -1000, -1000, -1000,
}
var exprPgo = [...]int{

	0, 122, 121, 120, 0, 4, 46, 119, 3, 118,
	2, 117, 115, 1, 80,
}
var exprR1 = [...]int{

	0, 1, 2, 2, 2, 6, 6, 6, 6, 6,
	7, 10, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 3, 3, 3, 3, 12, 12, 12, 9,
	9, 8, 8, 8, 8, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 11, 11, 5, 5, 4, 4,
}
var exprR2 = [...]int{

	0, 1, 1, 1, 1, 1, 3, 3, 3, 2,
	2, 4, 4, 4, 5, 5, 5, 5, 6, 7,
	6, 7, 1, 1, 1, 1, 3, 3, 3, 1,
	3, 3, 3, 3, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 4, 4,
}
var exprChk = [...]int{

	-1000, -1, -2, -6, -10, -13, -12, 21, -11, -14,
	13, 25, 26, 27, 28, 31, 29, 30, 32, 33,
	34, 35, -3, 2, 19, 20, 12, 10, -6, 21,
	21, -4, 23, 24, -9, 2, -8, 4, 5, 2,
	22, -7, -6, -10, -13, 4, 21, 21, 21, 14,
	2, 17, 14, 9, 10, 11, 12, 22, 6, 22,
	22, 17, -10, -13, -5, 4, -5, -8, 5, 5,
	5, 5, -4, -4, -13, -10, 22, 22, 17, 22,
	22, 22, 22, 4, -4, -4,
}
var exprDef = [...]int{

	0, -2, 1, -2, 3, 4, 5, 0, 0, 0,
	0, 44, 45, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 0, 9, 22, 23, 24, 25, 0, 0,
	0, 0, 0, 0, 0, 0, 29, 0, 6, 8,
	7, 0, 0, 0, 0, 0, 0, 0, 0, 26,
	27, 0, 28, 0, 0, 0, 0, 11, 10, 12,
	13, 0, 0, 0, 0, 46, 0, 30, 31, 32,
	33, 34, 16, 17, 0, 0, 14, 15, 0, 48,
	49, 18, 20, 47, 19, 21,
}
var exprTok1 = [...]int{

	1,
}
var exprTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35,
}
var exprTok3 = [...]int{
	0,
}

var exprErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	exprDebug        = 0
	exprErrorVerbose = false
)

type exprLexer interface {
	Lex(lval *exprSymType) int
	Error(s string)
}

type exprParser interface {
	Parse(exprLexer) int
	Lookahead() int
}

type exprParserImpl struct {
	lval  exprSymType
	stack [exprInitialStackSize]exprSymType
	char  int
}

func (p *exprParserImpl) Lookahead() int {
	return p.char
}

func exprNewParser() exprParser {
	return &exprParserImpl{}
}

const exprFlag = -1000

func exprTokname(c int) string {
	if c >= 1 && c-1 < len(exprToknames) {
		if exprToknames[c-1] != "" {
			return exprToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func exprStatname(s int) string {
	if s >= 0 && s < len(exprStatenames) {
		if exprStatenames[s] != "" {
			return exprStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func exprErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !exprErrorVerbose {
		return "syntax error"
	}

	for _, e := range exprErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + exprTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := exprPact[state]
	for tok := TOKSTART; tok-1 < len(exprToknames); tok++ {
		if n := base + tok; n >= 0 && n < exprLast && exprChk[exprAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if exprDef[state] == -2 {
		i := 0
		for exprExca[i] != -1 || exprExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; exprExca[i] >= 0; i += 2 {
			tok := exprExca[i]
			if tok < TOKSTART || exprExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if exprExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += exprTokname(tok)
	}
	return res
}

func exprlex1(lex exprLexer, lval *exprSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = exprTok1[0]
		goto out
	}
	if char < len(exprTok1) {
		token = exprTok1[char]
		goto out
	}
	if char >= exprPrivate {
		if char < exprPrivate+len(exprTok2) {
			token = exprTok2[char-exprPrivate]
			goto out
		}
	}
	for i := 0; i < len(exprTok3); i += 2 {
		token = exprTok3[i+0]
		if token == char {
			token = exprTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = exprTok2[1] /* unknown char */
	}
	if exprDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", exprTokname(token), uint(char))
	}
	return char, token
}

func exprParse(exprlex exprLexer) int {
	return exprNewParser().Parse(exprlex)
}

func (exprrcvr *exprParserImpl) Parse(exprlex exprLexer) int {
	var exprn int
	var exprVAL exprSymType
	var exprDollar []exprSymType
	_ = exprDollar // silence set and not used
	exprS := exprrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	exprstate := 0
	exprrcvr.char = -1
	exprtoken := -1 // exprrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		exprstate = -1
		exprrcvr.char = -1
		exprtoken = -1
	}()
	exprp := -1
	goto exprstack

ret0:
	return 0

ret1:
	return 1

exprstack:
	/* put a state and value onto the stack */
	if exprDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", exprTokname(exprtoken), exprStatname(exprstate))
	}

	exprp++
	if exprp >= len(exprS) {
		nyys := make([]exprSymType, len(exprS)*2)
		copy(nyys, exprS)
		exprS = nyys
	}
	exprS[exprp] = exprVAL
	exprS[exprp].yys = exprstate

exprnewstate:
	exprn = exprPact[exprstate]
	if exprn <= exprFlag {
		goto exprdefault /* simple state */
	}
	if exprrcvr.char < 0 {
		exprrcvr.char, exprtoken = exprlex1(exprlex, &exprrcvr.lval)
	}
	exprn += exprtoken
	if exprn < 0 || exprn >= exprLast {
		goto exprdefault
	}
	exprn = exprAct[exprn]
	if exprChk[exprn] == exprtoken { /* valid shift */
		exprrcvr.char = -1
		exprtoken = -1
		exprVAL = exprrcvr.lval
		exprstate = exprn
		if Errflag > 0 {
			Errflag--
		}
		goto exprstack
	}

exprdefault:
	/* default state action */
	exprn = exprDef[exprstate]
	if exprn == -2 {
		if exprrcvr.char < 0 {
			exprrcvr.char, exprtoken = exprlex1(exprlex, &exprrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if exprExca[xi+0] == -1 && exprExca[xi+1] == exprstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			exprn = exprExca[xi+0]
			if exprn < 0 || exprn == exprtoken {
				break
			}
		}
		exprn = exprExca[xi+1]
		if exprn < 0 {
			goto ret0
		}
	}
	if exprn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			exprlex.Error(exprErrorMessage(exprstate, exprtoken))
			Nerrs++
			if exprDebug >= 1 {
				__yyfmt__.Printf("%s", exprStatname(exprstate))
				__yyfmt__.Printf(" saw %s\n", exprTokname(exprtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for exprp >= 0 {
				exprn = exprPact[exprS[exprp].yys] + exprErrCode
				if exprn >= 0 && exprn < exprLast {
					exprstate = exprAct[exprn] /* simulate a shift of "error" */
					if exprChk[exprstate] == exprErrCode {
						goto exprstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if exprDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", exprS[exprp].yys)
				}
				exprp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if exprDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", exprTokname(exprtoken))
			}
			if exprtoken == exprEofCode {
				goto ret1
			}
			exprrcvr.char = -1
			exprtoken = -1
			goto exprnewstate /* try again in the same state */
		}
	}

	/* reduction by production exprn */
	if exprDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", exprn, exprStatname(exprstate))
	}

	exprnt := exprn
	exprpt := exprp
	_ = exprpt // guard against "declared and not used"

	exprp -= exprR2[exprn]
	// exprp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if exprp+1 >= len(exprS) {
		nyys := make([]exprSymType, len(exprS)*2)
		copy(nyys, exprS)
		exprS = nyys
	}
	exprVAL = exprS[exprp+1]

	/* consult goto table to find next state */
	exprn = exprR1[exprn]
	exprg := exprPgo[exprn]
	exprj := exprg + exprS[exprp].yys + 1

	if exprj >= exprLast {
		exprstate = exprAct[exprg]
	} else {
		exprstate = exprAct[exprj]
		if exprChk[exprstate] != -exprn {
			exprstate = exprAct[exprg]
		}
	}
	// dummy call; replaced with literal code
	switch exprnt {

	case 1:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:52
		{
			exprlex.(*lexer).expr = exprDollar[1].Expr
		}
	case 2:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:55
		{
			exprVAL.Expr = exprDollar[1].LogExpr
		}
	case 3:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:56
		{
			exprVAL.Expr = exprDollar[1].RangeAggregationExpr
		}
	case 4:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:57
		{
			exprVAL.Expr = exprDollar[1].VectorAggregationExpr
		}
	case 5:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:61
		{
			exprVAL.LogExpr = newMatcherExpr(exprDollar[1].Selector)
		}
	case 6:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:62
		{
			exprVAL.LogExpr = NewFilterExpr(exprDollar[1].LogExpr, exprDollar[2].Filter, exprDollar[3].str)
		}
	case 7:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:63
		{
			exprVAL.LogExpr = exprDollar[2].LogExpr
		}
	case 10:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line pkg/logql/expr.y:69
		{
			exprVAL.LogRangeExpr = mustNewRange(exprDollar[1].LogExpr, exprDollar[2].duration)
		}
	case 11:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:72
		{
			exprVAL.RangeAggregationExpr = newRangeAggregationExpr(exprDollar[3].LogRangeExpr, exprDollar[1].RangeOp)
		}
	case 12:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:76
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[3].RangeAggregationExpr, exprDollar[1].VectorOp, nil, nil)
		}
	case 13:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:77
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[3].VectorAggregationExpr, exprDollar[1].VectorOp, nil, nil)
		}
	case 14:
		exprDollar = exprS[exprpt-5 : exprpt+1]
//line pkg/logql/expr.y:78
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[4].RangeAggregationExpr, exprDollar[1].VectorOp, exprDollar[2].Grouping, nil)
		}
	case 15:
		exprDollar = exprS[exprpt-5 : exprpt+1]
//line pkg/logql/expr.y:79
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[4].VectorAggregationExpr, exprDollar[1].VectorOp, exprDollar[2].Grouping, nil)
		}
	case 16:
		exprDollar = exprS[exprpt-5 : exprpt+1]
//line pkg/logql/expr.y:80
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[3].RangeAggregationExpr, exprDollar[1].VectorOp, exprDollar[5].Grouping, nil)
		}
	case 17:
		exprDollar = exprS[exprpt-5 : exprpt+1]
//line pkg/logql/expr.y:81
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[3].VectorAggregationExpr, exprDollar[1].VectorOp, exprDollar[5].Grouping, nil)
		}
	case 18:
		exprDollar = exprS[exprpt-6 : exprpt+1]
//line pkg/logql/expr.y:83
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[5].VectorAggregationExpr, exprDollar[1].VectorOp, nil, &exprDollar[3].str)
		}
	case 19:
		exprDollar = exprS[exprpt-7 : exprpt+1]
//line pkg/logql/expr.y:84
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[5].VectorAggregationExpr, exprDollar[1].VectorOp, exprDollar[7].Grouping, &exprDollar[3].str)
		}
	case 20:
		exprDollar = exprS[exprpt-6 : exprpt+1]
//line pkg/logql/expr.y:85
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[5].RangeAggregationExpr, exprDollar[1].VectorOp, nil, &exprDollar[3].str)
		}
	case 21:
		exprDollar = exprS[exprpt-7 : exprpt+1]
//line pkg/logql/expr.y:86
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[5].RangeAggregationExpr, exprDollar[1].VectorOp, exprDollar[7].Grouping, &exprDollar[3].str)
		}
	case 22:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:90
		{
			exprVAL.Filter = labels.MatchRegexp
		}
	case 23:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:91
		{
			exprVAL.Filter = labels.MatchEqual
		}
	case 24:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:92
		{
			exprVAL.Filter = labels.MatchNotRegexp
		}
	case 25:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:93
		{
			exprVAL.Filter = labels.MatchNotEqual
		}
	case 26:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:97
		{
			exprVAL.Selector = exprDollar[2].Matchers
		}
	case 27:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:98
		{
			exprVAL.Selector = exprDollar[2].Matchers
		}
	case 28:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:99
		{
		}
	case 29:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:103
		{
			exprVAL.Matchers = []*labels.Matcher{exprDollar[1].Matcher}
		}
	case 30:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:104
		{
			exprVAL.Matchers = append(exprDollar[1].Matchers, exprDollar[3].Matcher)
		}
	case 31:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:108
		{
			exprVAL.Matcher = mustNewMatcher(labels.MatchEqual, exprDollar[1].str, exprDollar[3].str)
		}
	case 32:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:109
		{
			exprVAL.Matcher = mustNewMatcher(labels.MatchNotEqual, exprDollar[1].str, exprDollar[3].str)
		}
	case 33:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:110
		{
			exprVAL.Matcher = mustNewMatcher(labels.MatchRegexp, exprDollar[1].str, exprDollar[3].str)
		}
	case 34:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:111
		{
			exprVAL.Matcher = mustNewMatcher(labels.MatchNotRegexp, exprDollar[1].str, exprDollar[3].str)
		}
	case 35:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:115
		{
			exprVAL.VectorOp = OpTypeSum
		}
	case 36:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:116
		{
			exprVAL.VectorOp = OpTypeAvg
		}
	case 37:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:117
		{
			exprVAL.VectorOp = OpTypeCount
		}
	case 38:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:118
		{
			exprVAL.VectorOp = OpTypeMax
		}
	case 39:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:119
		{
			exprVAL.VectorOp = OpTypeMin
		}
	case 40:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:120
		{
			exprVAL.VectorOp = OpTypeStddev
		}
	case 41:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:121
		{
			exprVAL.VectorOp = OpTypeStdvar
		}
	case 42:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:122
		{
			exprVAL.VectorOp = OpTypeBottomK
		}
	case 43:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:123
		{
			exprVAL.VectorOp = OpTypeTopK
		}
	case 44:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:127
		{
			exprVAL.RangeOp = OpTypeCountOverTime
		}
	case 45:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:128
		{
			exprVAL.RangeOp = OpTypeRate
		}
	case 46:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line pkg/logql/expr.y:132
		{
			exprVAL.Labels = []string{exprDollar[1].str}
		}
	case 47:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line pkg/logql/expr.y:133
		{
			exprVAL.Labels = append(exprDollar[1].Labels, exprDollar[3].str)
		}
	case 48:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:137
		{
			exprVAL.Grouping = &grouping{without: false, groups: exprDollar[3].Labels}
		}
	case 49:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line pkg/logql/expr.y:138
		{
			exprVAL.Grouping = &grouping{without: true, groups: exprDollar[3].Labels}
		}
	}
	goto exprstack /* stack new state and value */
}