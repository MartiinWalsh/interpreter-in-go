package parser

import (
	"github.com/MartiinWalsh/interpreter-in-go/ast"
	"github.com/MartiinWalsh/interpreter-in-go/lexer"
	"github.com/MartiinWalsh/interpreter-in-go/token"
)

type Parser struct {
	l *lexer.Lexer // pointer to an instance of the lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l} // create a new instance of the parser and initialize it with a lexer

	// Read two tokens, so curToken and peekToken are both set
	p.curToken = p.peekToken      // set curToken to the first token
	p.peekToken = p.l.NextToken() // set peekToken to the second token

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
