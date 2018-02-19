// Code generated from Saucy.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Saucy

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SaucyListener is a complete listener for a parse tree produced by SaucyParser.
type SaucyListener interface {
	antlr.ParseTreeListener

	// EnterFile_input is called when entering the file_input production.
	EnterFile_input(c *File_inputContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterSimple_stmt is called when entering the simple_stmt production.
	EnterSimple_stmt(c *Simple_stmtContext)

	// EnterSmall_stmt is called when entering the small_stmt production.
	EnterSmall_stmt(c *Small_stmtContext)

	// EnterImport_stmt is called when entering the import_stmt production.
	EnterImport_stmt(c *Import_stmtContext)

	// EnterImport_name is called when entering the import_name production.
	EnterImport_name(c *Import_nameContext)

	// EnterDotted_name is called when entering the dotted_name production.
	EnterDotted_name(c *Dotted_nameContext)

	// ExitFile_input is called when exiting the file_input production.
	ExitFile_input(c *File_inputContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitSimple_stmt is called when exiting the simple_stmt production.
	ExitSimple_stmt(c *Simple_stmtContext)

	// ExitSmall_stmt is called when exiting the small_stmt production.
	ExitSmall_stmt(c *Small_stmtContext)

	// ExitImport_stmt is called when exiting the import_stmt production.
	ExitImport_stmt(c *Import_stmtContext)

	// ExitImport_name is called when exiting the import_name production.
	ExitImport_name(c *Import_nameContext)

	// ExitDotted_name is called when exiting the dotted_name production.
	ExitDotted_name(c *Dotted_nameContext)
}
