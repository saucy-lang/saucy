// Code generated from Saucy.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Saucy

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSaucyListener is a complete listener for a parse tree produced by SaucyParser.
type BaseSaucyListener struct{}

var _ SaucyListener = &BaseSaucyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSaucyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSaucyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSaucyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSaucyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile_input is called when production file_input is entered.
func (s *BaseSaucyListener) EnterFile_input(ctx *File_inputContext) {}

// ExitFile_input is called when production file_input is exited.
func (s *BaseSaucyListener) ExitFile_input(ctx *File_inputContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseSaucyListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseSaucyListener) ExitStmt(ctx *StmtContext) {}

// EnterSimple_stmt is called when production simple_stmt is entered.
func (s *BaseSaucyListener) EnterSimple_stmt(ctx *Simple_stmtContext) {}

// ExitSimple_stmt is called when production simple_stmt is exited.
func (s *BaseSaucyListener) ExitSimple_stmt(ctx *Simple_stmtContext) {}

// EnterSmall_stmt is called when production small_stmt is entered.
func (s *BaseSaucyListener) EnterSmall_stmt(ctx *Small_stmtContext) {}

// ExitSmall_stmt is called when production small_stmt is exited.
func (s *BaseSaucyListener) ExitSmall_stmt(ctx *Small_stmtContext) {}

// EnterImport_stmt is called when production import_stmt is entered.
func (s *BaseSaucyListener) EnterImport_stmt(ctx *Import_stmtContext) {}

// ExitImport_stmt is called when production import_stmt is exited.
func (s *BaseSaucyListener) ExitImport_stmt(ctx *Import_stmtContext) {}

// EnterImport_name is called when production import_name is entered.
func (s *BaseSaucyListener) EnterImport_name(ctx *Import_nameContext) {}

// ExitImport_name is called when production import_name is exited.
func (s *BaseSaucyListener) ExitImport_name(ctx *Import_nameContext) {}

// EnterDotted_name is called when production dotted_name is entered.
func (s *BaseSaucyListener) EnterDotted_name(ctx *Dotted_nameContext) {}

// ExitDotted_name is called when production dotted_name is exited.
func (s *BaseSaucyListener) ExitDotted_name(ctx *Dotted_nameContext) {}
