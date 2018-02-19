// Code generated from Saucy.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Saucy

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 8, 56, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 3, 2, 7, 2, 19, 10, 2, 12, 2, 14, 2, 22, 11, 2, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 31, 10, 4, 12, 4, 14, 4, 34, 11,
	4, 3, 4, 5, 4, 37, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 8, 7, 8, 51, 10, 8, 12, 8, 14, 8, 54, 11, 8, 3,
	8, 2, 2, 9, 2, 4, 6, 8, 10, 12, 14, 2, 2, 2, 53, 2, 20, 3, 2, 2, 2, 4,
	25, 3, 2, 2, 2, 6, 27, 3, 2, 2, 2, 8, 40, 3, 2, 2, 2, 10, 42, 3, 2, 2,
	2, 12, 45, 3, 2, 2, 2, 14, 47, 3, 2, 2, 2, 16, 19, 7, 7, 2, 2, 17, 19,
	5, 4, 3, 2, 18, 16, 3, 2, 2, 2, 18, 17, 3, 2, 2, 2, 19, 22, 3, 2, 2, 2,
	20, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 23, 3, 2, 2, 2, 22, 20, 3,
	2, 2, 2, 23, 24, 7, 2, 2, 3, 24, 3, 3, 2, 2, 2, 25, 26, 5, 6, 4, 2, 26,
	5, 3, 2, 2, 2, 27, 32, 5, 8, 5, 2, 28, 29, 7, 3, 2, 2, 29, 31, 5, 8, 5,
	2, 30, 28, 3, 2, 2, 2, 31, 34, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 32, 33,
	3, 2, 2, 2, 33, 36, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 35, 37, 7, 3, 2, 2,
	36, 35, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 39, 7,
	7, 2, 2, 39, 7, 3, 2, 2, 2, 40, 41, 5, 10, 6, 2, 41, 9, 3, 2, 2, 2, 42,
	43, 7, 5, 2, 2, 43, 44, 5, 12, 7, 2, 44, 11, 3, 2, 2, 2, 45, 46, 5, 14,
	8, 2, 46, 13, 3, 2, 2, 2, 47, 52, 7, 6, 2, 2, 48, 49, 7, 4, 2, 2, 49, 51,
	7, 6, 2, 2, 50, 48, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2,
	52, 53, 3, 2, 2, 2, 53, 15, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 7, 18, 20,
	32, 36, 52,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'.'", "'import'",
}
var symbolicNames = []string{
	"", "", "", "IMPORT", "NAME", "NEWLINE", "SKIP_",
}

var ruleNames = []string{
	"file_input", "stmt", "simple_stmt", "small_stmt", "import_stmt", "import_name",
	"dotted_name",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SaucyParser struct {
	*antlr.BaseParser
}

func NewSaucyParser(input antlr.TokenStream) *SaucyParser {
	this := new(SaucyParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Saucy.g4"

	return this
}

// SaucyParser tokens.
const (
	SaucyParserEOF     = antlr.TokenEOF
	SaucyParserT__0    = 1
	SaucyParserT__1    = 2
	SaucyParserIMPORT  = 3
	SaucyParserNAME    = 4
	SaucyParserNEWLINE = 5
	SaucyParserSKIP_   = 6
)

// SaucyParser rules.
const (
	SaucyParserRULE_file_input  = 0
	SaucyParserRULE_stmt        = 1
	SaucyParserRULE_simple_stmt = 2
	SaucyParserRULE_small_stmt  = 3
	SaucyParserRULE_import_stmt = 4
	SaucyParserRULE_import_name = 5
	SaucyParserRULE_dotted_name = 6
)

// IFile_inputContext is an interface to support dynamic dispatch.
type IFile_inputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFile_inputContext differentiates from other interfaces.
	IsFile_inputContext()
}

type File_inputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFile_inputContext() *File_inputContext {
	var p = new(File_inputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SaucyParserRULE_file_input
	return p
}

func (*File_inputContext) IsFile_inputContext() {}

func NewFile_inputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *File_inputContext {
	var p = new(File_inputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SaucyParserRULE_file_input

	return p
}

func (s *File_inputContext) GetParser() antlr.Parser { return s.parser }

func (s *File_inputContext) EOF() antlr.TerminalNode {
	return s.GetToken(SaucyParserEOF, 0)
}

func (s *File_inputContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(SaucyParserNEWLINE)
}

func (s *File_inputContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(SaucyParserNEWLINE, i)
}

func (s *File_inputContext) AllStmt() []IStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *File_inputContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *File_inputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *File_inputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *File_inputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SaucyListener); ok {
		listenerT.EnterFile_input(s)
	}
}

func (s *File_inputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SaucyListener); ok {
		listenerT.ExitFile_input(s)
	}
}

func (p *SaucyParser) File_input() (localctx IFile_inputContext) {
	localctx = NewFile_inputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SaucyParserRULE_file_input)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SaucyParserIMPORT || _la == SaucyParserNEWLINE {
		p.SetState(16)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SaucyParserNEWLINE:
			{
				p.SetState(14)
				p.Match(SaucyParserNEWLINE)
			}

		case SaucyParserIMPORT:
			{
				p.SetState(15)
				p.Stmt()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(21)
		p.Match(SaucyParserEOF)
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SaucyParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SaucyParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Simple_stmt() ISimple_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_stmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SaucyListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SaucyListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *SaucyParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SaucyParserRULE_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)
		p.Simple_stmt()
	}

	return localctx
}

// ISimple_stmtContext is an interface to support dynamic dispatch.
type ISimple_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_stmtContext differentiates from other interfaces.
	IsSimple_stmtContext()
}

type Simple_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_stmtContext() *Simple_stmtContext {
	var p = new(Simple_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SaucyParserRULE_simple_stmt
	return p
}

func (*Simple_stmtContext) IsSimple_stmtContext() {}

func NewSimple_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_stmtContext {
	var p = new(Simple_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SaucyParserRULE_simple_stmt

	return p
}

func (s *Simple_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_stmtContext) AllSmall_stmt() []ISmall_stmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISmall_stmtContext)(nil)).Elem())
	var tst = make([]ISmall_stmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISmall_stmtContext)
		}
	}

	return tst
}

func (s *Simple_stmtContext) Small_stmt(i int) ISmall_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISmall_stmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISmall_stmtContext)
}

func (s *Simple_stmtContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(SaucyParserNEWLINE, 0)
}

func (s *Simple_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SaucyListener); ok {
		listenerT.EnterSimple_stmt(s)
	}
}

func (s *Simple_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SaucyListener); ok {
		listenerT.ExitSimple_stmt(s)
	}
}

func (p *SaucyParser) Simple_stmt() (localctx ISimple_stmtContext) {
	localctx = NewSimple_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SaucyParserRULE_simple_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Small_stmt()
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(26)
				p.Match(SaucyParserT__0)
			}
			{
				p.SetState(27)
				p.Small_stmt()
			}

		}
		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SaucyParserT__0 {
		{
			p.SetState(33)
			p.Match(SaucyParserT__0)
		}

	}
	{
		p.SetState(36)
		p.Match(SaucyParserNEWLINE)
	}

	return localctx
}

// ISmall_stmtContext is an interface to support dynamic dispatch.
type ISmall_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSmall_stmtContext differentiates from other interfaces.
	IsSmall_stmtContext()
}

type Small_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySmall_stmtContext() *Small_stmtContext {
	var p = new(Small_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SaucyParserRULE_small_stmt
	return p
}

func (*Small_stmtContext) IsSmall_stmtContext() {}

func NewSmall_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Small_stmtContext {
	var p = new(Small_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SaucyParserRULE_small_stmt

	return p
}

func (s *Small_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Small_stmtContext) Import_stmt() IImport_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImport_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImport_stmtContext)
}

func (s *Small_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Small_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Small_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SaucyListener); ok {
		listenerT.EnterSmall_stmt(s)
	}
}

func (s *Small_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SaucyListener); ok {
		listenerT.ExitSmall_stmt(s)
	}
}

func (p *SaucyParser) Small_stmt() (localctx ISmall_stmtContext) {
	localctx = NewSmall_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SaucyParserRULE_small_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Import_stmt()
	}

	return localctx
}

// IImport_stmtContext is an interface to support dynamic dispatch.
type IImport_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImport_stmtContext differentiates from other interfaces.
	IsImport_stmtContext()
}

type Import_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImport_stmtContext() *Import_stmtContext {
	var p = new(Import_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SaucyParserRULE_import_stmt
	return p
}

func (*Import_stmtContext) IsImport_stmtContext() {}

func NewImport_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Import_stmtContext {
	var p = new(Import_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SaucyParserRULE_import_stmt

	return p
}

func (s *Import_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Import_stmtContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(SaucyParserIMPORT, 0)
}

func (s *Import_stmtContext) Import_name() IImport_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImport_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImport_nameContext)
}

func (s *Import_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Import_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Import_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SaucyListener); ok {
		listenerT.EnterImport_stmt(s)
	}
}

func (s *Import_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SaucyListener); ok {
		listenerT.ExitImport_stmt(s)
	}
}

func (p *SaucyParser) Import_stmt() (localctx IImport_stmtContext) {
	localctx = NewImport_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SaucyParserRULE_import_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(SaucyParserIMPORT)
	}
	{
		p.SetState(41)
		p.Import_name()
	}

	return localctx
}

// IImport_nameContext is an interface to support dynamic dispatch.
type IImport_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImport_nameContext differentiates from other interfaces.
	IsImport_nameContext()
}

type Import_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImport_nameContext() *Import_nameContext {
	var p = new(Import_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SaucyParserRULE_import_name
	return p
}

func (*Import_nameContext) IsImport_nameContext() {}

func NewImport_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Import_nameContext {
	var p = new(Import_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SaucyParserRULE_import_name

	return p
}

func (s *Import_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Import_nameContext) Dotted_name() IDotted_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDotted_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDotted_nameContext)
}

func (s *Import_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Import_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Import_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SaucyListener); ok {
		listenerT.EnterImport_name(s)
	}
}

func (s *Import_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SaucyListener); ok {
		listenerT.ExitImport_name(s)
	}
}

func (p *SaucyParser) Import_name() (localctx IImport_nameContext) {
	localctx = NewImport_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SaucyParserRULE_import_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Dotted_name()
	}

	return localctx
}

// IDotted_nameContext is an interface to support dynamic dispatch.
type IDotted_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDotted_nameContext differentiates from other interfaces.
	IsDotted_nameContext()
}

type Dotted_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDotted_nameContext() *Dotted_nameContext {
	var p = new(Dotted_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SaucyParserRULE_dotted_name
	return p
}

func (*Dotted_nameContext) IsDotted_nameContext() {}

func NewDotted_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dotted_nameContext {
	var p = new(Dotted_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SaucyParserRULE_dotted_name

	return p
}

func (s *Dotted_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Dotted_nameContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(SaucyParserNAME)
}

func (s *Dotted_nameContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(SaucyParserNAME, i)
}

func (s *Dotted_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dotted_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dotted_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SaucyListener); ok {
		listenerT.EnterDotted_name(s)
	}
}

func (s *Dotted_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SaucyListener); ok {
		listenerT.ExitDotted_name(s)
	}
}

func (p *SaucyParser) Dotted_name() (localctx IDotted_nameContext) {
	localctx = NewDotted_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SaucyParserRULE_dotted_name)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(SaucyParserNAME)
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SaucyParserT__1 {
		{
			p.SetState(46)
			p.Match(SaucyParserT__1)
		}
		{
			p.SetState(47)
			p.Match(SaucyParserNAME)
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
