// Code generated from Saucy.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 8, 99, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 6, 5, 36, 10, 5,
	13, 5, 14, 5, 37, 3, 6, 5, 6, 41, 10, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7,
	3, 7, 5, 7, 49, 10, 7, 3, 7, 3, 7, 3, 8, 6, 8, 54, 10, 8, 13, 8, 14, 8,
	55, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 62, 10, 9, 12, 9, 14, 9, 65, 11, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 7, 9, 71, 10, 9, 12, 9, 14, 9, 74, 11, 9, 3, 9, 3,
	9, 5, 9, 78, 10, 9, 3, 10, 3, 10, 5, 10, 82, 10, 10, 3, 10, 5, 10, 85,
	10, 10, 3, 10, 3, 10, 5, 10, 89, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 7,
	11, 95, 10, 11, 12, 11, 14, 11, 98, 11, 11, 3, 72, 2, 12, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 2, 17, 2, 19, 2, 21, 2, 3, 2, 5, 6, 2, 50, 59,
	67, 92, 97, 97, 99, 124, 4, 2, 11, 11, 34, 34, 4, 2, 12, 12, 14, 15, 2,
	107, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 3, 23, 3, 2, 2, 2, 5, 25, 3,
	2, 2, 2, 7, 27, 3, 2, 2, 2, 9, 35, 3, 2, 2, 2, 11, 40, 3, 2, 2, 2, 13,
	48, 3, 2, 2, 2, 15, 53, 3, 2, 2, 2, 17, 77, 3, 2, 2, 2, 19, 79, 3, 2, 2,
	2, 21, 90, 3, 2, 2, 2, 23, 24, 7, 61, 2, 2, 24, 4, 3, 2, 2, 2, 25, 26,
	7, 48, 2, 2, 26, 6, 3, 2, 2, 2, 27, 28, 7, 107, 2, 2, 28, 29, 7, 111, 2,
	2, 29, 30, 7, 114, 2, 2, 30, 31, 7, 113, 2, 2, 31, 32, 7, 116, 2, 2, 32,
	33, 7, 118, 2, 2, 33, 8, 3, 2, 2, 2, 34, 36, 9, 2, 2, 2, 35, 34, 3, 2,
	2, 2, 36, 37, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 10,
	3, 2, 2, 2, 39, 41, 7, 15, 2, 2, 40, 39, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2,
	41, 42, 3, 2, 2, 2, 42, 43, 7, 12, 2, 2, 43, 12, 3, 2, 2, 2, 44, 49, 5,
	15, 8, 2, 45, 49, 5, 17, 9, 2, 46, 49, 5, 19, 10, 2, 47, 49, 5, 21, 11,
	2, 48, 44, 3, 2, 2, 2, 48, 45, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 47,
	3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 51, 8, 7, 2, 2, 51, 14, 3, 2, 2, 2,
	52, 54, 9, 3, 2, 2, 53, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 53, 3,
	2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 16, 3, 2, 2, 2, 57, 58, 7, 49, 2, 2, 58,
	59, 7, 49, 2, 2, 59, 63, 3, 2, 2, 2, 60, 62, 10, 4, 2, 2, 61, 60, 3, 2,
	2, 2, 62, 65, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 78,
	3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 66, 67, 7, 49, 2, 2, 67, 68, 7, 44, 2,
	2, 68, 72, 3, 2, 2, 2, 69, 71, 11, 2, 2, 2, 70, 69, 3, 2, 2, 2, 71, 74,
	3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 73, 75, 3, 2, 2, 2,
	74, 72, 3, 2, 2, 2, 75, 76, 7, 44, 2, 2, 76, 78, 7, 49, 2, 2, 77, 57, 3,
	2, 2, 2, 77, 66, 3, 2, 2, 2, 78, 18, 3, 2, 2, 2, 79, 81, 7, 94, 2, 2, 80,
	82, 5, 15, 8, 2, 81, 80, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 88, 3, 2,
	2, 2, 83, 85, 7, 15, 2, 2, 84, 83, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85,
	86, 3, 2, 2, 2, 86, 89, 7, 12, 2, 2, 87, 89, 4, 14, 15, 2, 88, 84, 3, 2,
	2, 2, 88, 87, 3, 2, 2, 2, 89, 20, 3, 2, 2, 2, 90, 91, 7, 37, 2, 2, 91,
	92, 7, 35, 2, 2, 92, 96, 3, 2, 2, 2, 93, 95, 10, 4, 2, 2, 94, 93, 3, 2,
	2, 2, 95, 98, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 22,
	3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 14, 2, 37, 40, 48, 55, 63, 72, 77, 81,
	84, 88, 96, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "'.'", "'import'",
}

var lexerSymbolicNames = []string{
	"", "", "", "IMPORT", "NAME", "NEWLINE", "SKIP_",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "IMPORT", "NAME", "NEWLINE", "SKIP_", "SPACES", "COMMENT",
	"LINE_JOINING", "SHEBANG",
}

type SaucyLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewSaucyLexer(input antlr.CharStream) *SaucyLexer {

	l := new(SaucyLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Saucy.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SaucyLexer tokens.
const (
	SaucyLexerT__0    = 1
	SaucyLexerT__1    = 2
	SaucyLexerIMPORT  = 3
	SaucyLexerNAME    = 4
	SaucyLexerNEWLINE = 5
	SaucyLexerSKIP_   = 6
)
