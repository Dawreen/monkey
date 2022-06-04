// lexer/lexer.go
package lexer

type Lexer struct {
	input        string
	position     int  // current posiotion in input (current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

/*
	We will need to be able to "peek" further into the input and look after
	the current character to see what comes next.
*/

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
