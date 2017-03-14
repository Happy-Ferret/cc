// CAUTION: Generated file - DO NOT EDIT.

// Copyright 2016 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on [0], 6.4.

package cc

import (
	"fmt"

	"github.com/cznic/golex/lex"
)

const (
	_           = iota
	scCOMMENT   // [`/*`, `*/`]
	scDEFINE    // [^#define, next token]
	scDIRECTIVE // [^#, next token]
	scHEADER    // [`#include`, next token]
)

func (l *lexer) scan() (r int) {
	c := l.Enter()

yystate0:
	yyrule := -1
	_ = yyrule
	c = l.Rule0()

	switch yyt := l.sc; yyt {
	default:
		panic(fmt.Errorf(`invalid start condition %d`, yyt))
	case 0: // start condition: INITIAL
		goto yystart1
	case 1: // start condition: COMMENT
		goto yystart147
	case 2: // start condition: DEFINE
		goto yystart152
	case 3: // start condition: DIRECTIVE
		goto yystart165
	case 4: // start condition: HEADER
		goto yystart221
	}

	goto yystate0 // silence unused label error
	goto yyAction // silence unused label error
yyAction:
	switch yyrule {
	case 1:
		goto yyrule1
	case 2:
		goto yyrule2
	case 3:
		goto yyrule3
	case 4:
		goto yyrule4
	case 5:
		goto yyrule5
	case 6:
		goto yyrule6
	case 7:
		goto yyrule7
	case 8:
		goto yyrule8
	case 9:
		goto yyrule9
	case 10:
		goto yyrule10
	case 11:
		goto yyrule11
	case 12:
		goto yyrule12
	case 13:
		goto yyrule13
	case 14:
		goto yyrule14
	case 15:
		goto yyrule15
	case 16:
		goto yyrule16
	case 17:
		goto yyrule17
	case 18:
		goto yyrule18
	case 19:
		goto yyrule19
	case 20:
		goto yyrule20
	case 21:
		goto yyrule21
	case 22:
		goto yyrule22
	case 23:
		goto yyrule23
	case 24:
		goto yyrule24
	case 25:
		goto yyrule25
	case 26:
		goto yyrule26
	case 27:
		goto yyrule27
	case 28:
		goto yyrule28
	case 29:
		goto yyrule29
	case 30:
		goto yyrule30
	case 31:
		goto yyrule31
	case 32:
		goto yyrule32
	case 33:
		goto yyrule33
	case 34:
		goto yyrule34
	case 35:
		goto yyrule35
	case 36:
		goto yyrule36
	case 37:
		goto yyrule37
	case 38:
		goto yyrule38
	case 39:
		goto yyrule39
	case 40:
		goto yyrule40
	case 41:
		goto yyrule41
	case 42:
		goto yyrule42
	case 43:
		goto yyrule43
	case 44:
		goto yyrule44
	case 45:
		goto yyrule45
	case 46:
		goto yyrule46
	case 47:
		goto yyrule47
	case 48:
		goto yyrule48
	case 49:
		goto yyrule49
	case 50:
		goto yyrule50
	case 51:
		goto yyrule51
	case 52:
		goto yyrule52
	case 53:
		goto yyrule53
	case 54:
		goto yyrule54
	case 55:
		goto yyrule55
	case 56:
		goto yyrule56
	case 57:
		goto yyrule57
	case 58:
		goto yyrule58
	case 59:
		goto yyrule59
	case 60:
		goto yyrule60
	}
	goto yystate1 // silence unused label error
yystate1:
	c = l.Next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate3
	case c == '"':
		goto yystate5
	case c == '#':
		goto yystate16
	case c == '%':
		goto yystate20
	case c == '&':
		goto yystate27
	case c == '*':
		goto yystate42
	case c == '+':
		goto yystate44
	case c == '-':
		goto yystate47
	case c == '.':
		goto yystate51
	case c == '/':
		goto yystate72
	case c == '0':
		goto yystate76
	case c == ':':
		goto yystate93
	case c == '<':
		goto yystate95
	case c == '=':
		goto yystate101
	case c == '>':
		goto yystate103
	case c == 'L':
		goto yystate117
	case c == '\'':
		goto yystate30
	case c == '\\':
		goto yystate108
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate2
	case c == '\u0080':
		goto yystate146
	case c == '^':
		goto yystate141
	case c == '|':
		goto yystate143
	case c >= '1' && c <= '9':
		goto yystate92
	case c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0084':
		goto yystate107
	}

yystate2:
	c = l.Next()
	yyrule = 1
	l.Mark()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate2
	}

yystate3:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate4
	}

yystate4:
	c = l.Next()
	yyrule = 7
	l.Mark()
	goto yyrule7

yystate5:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate6
	case c == '\\':
		goto yystate7
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate5
	}

yystate6:
	c = l.Next()
	yyrule = 60
	l.Mark()
	goto yyrule60

yystate7:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"' || c == '\'' || c >= '0' && c <= '7' || c == '?' || c == '\\' || c == 'a' || c == 'b' || c == 'f' || c == 'n' || c == 'r' || c == 't' || c == 'v':
		goto yystate5
	case c == 'U':
		goto yystate8
	case c == 'u':
		goto yystate12
	case c == 'x':
		goto yystate15
	}

yystate8:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate9
	}

yystate9:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate10
	}

yystate10:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate11
	}

yystate11:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate12
	}

yystate12:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate13
	}

yystate13:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate14
	}

yystate14:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate15
	}

yystate15:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate5
	}

yystate16:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '#':
		goto yystate17
	case c == '%':
		goto yystate18
	}

yystate17:
	c = l.Next()
	yyrule = 34
	l.Mark()
	goto yyrule34

yystate18:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == ':':
		goto yystate19
	}

yystate19:
	c = l.Next()
	yyrule = 35
	l.Mark()
	goto yyrule35

yystate20:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == ':':
		goto yystate21
	case c == '=':
		goto yystate25
	case c == '>':
		goto yystate26
	}

yystate21:
	c = l.Next()
	yyrule = 8
	l.Mark()
	switch {
	default:
		goto yyrule8
	case c == '#':
		goto yystate22
	case c == '%':
		goto yystate23
	}

yystate22:
	c = l.Next()
	yyrule = 36
	l.Mark()
	goto yyrule36

yystate23:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == ':':
		goto yystate24
	}

yystate24:
	c = l.Next()
	yyrule = 37
	l.Mark()
	goto yyrule37

yystate25:
	c = l.Next()
	yyrule = 9
	l.Mark()
	goto yyrule9

yystate26:
	c = l.Next()
	yyrule = 10
	l.Mark()
	goto yyrule10

yystate27:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '&':
		goto yystate28
	case c == '=':
		goto yystate29
	}

yystate28:
	c = l.Next()
	yyrule = 11
	l.Mark()
	goto yyrule11

yystate29:
	c = l.Next()
	yyrule = 12
	l.Mark()
	goto yyrule12

yystate30:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '\\':
		goto yystate33
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate31
	}

yystate31:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate32
	case c == '\\':
		goto yystate33
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate31
	}

yystate32:
	c = l.Next()
	yyrule = 54
	l.Mark()
	goto yyrule54

yystate33:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"' || c == '\'' || c >= '0' && c <= '7' || c == '?' || c == '\\' || c == 'a' || c == 'b' || c == 'f' || c == 'n' || c == 'r' || c == 't' || c == 'v':
		goto yystate31
	case c == 'U':
		goto yystate34
	case c == 'u':
		goto yystate38
	case c == 'x':
		goto yystate41
	}

yystate34:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate35
	}

yystate35:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate36
	}

yystate36:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate37
	}

yystate37:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate38
	}

yystate38:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate39
	}

yystate39:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate40
	}

yystate40:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate41
	}

yystate41:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate31
	}

yystate42:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate43
	}

yystate43:
	c = l.Next()
	yyrule = 13
	l.Mark()
	goto yyrule13

yystate44:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '+':
		goto yystate45
	case c == '=':
		goto yystate46
	}

yystate45:
	c = l.Next()
	yyrule = 14
	l.Mark()
	goto yyrule14

yystate46:
	c = l.Next()
	yyrule = 15
	l.Mark()
	goto yyrule15

yystate47:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '-':
		goto yystate48
	case c == '=':
		goto yystate49
	case c == '>':
		goto yystate50
	}

yystate48:
	c = l.Next()
	yyrule = 16
	l.Mark()
	goto yyrule16

yystate49:
	c = l.Next()
	yyrule = 17
	l.Mark()
	goto yyrule17

yystate50:
	c = l.Next()
	yyrule = 18
	l.Mark()
	goto yyrule18

yystate51:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '.':
		goto yystate52
	case c >= '0' && c <= '9':
		goto yystate54
	}

yystate52:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '.':
		goto yystate53
	}

yystate53:
	c = l.Next()
	yyrule = 19
	l.Mark()
	goto yyrule19

yystate54:
	c = l.Next()
	yyrule = 58
	l.Mark()
	switch {
	default:
		goto yyrule58
	case c == '.' || c >= 'A' && c <= 'D' || c >= 'G' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c == 'g' || c == 'h' || c == 'j' || c == 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'e':
		goto yystate66
	case c == 'F' || c == 'L' || c == 'f' || c == 'l':
		goto yystate69
	case c == 'P' || c == 'p':
		goto yystate56
	case c == '\\':
		goto yystate57
	case c == 'i':
		goto yystate71
	case c >= '0' && c <= '9':
		goto yystate54
	}

yystate55:
	c = l.Next()
	yyrule = 59
	l.Mark()
	switch {
	default:
		goto yyrule59
	case c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate56
	case c == '\\':
		goto yystate57
	}

yystate56:
	c = l.Next()
	yyrule = 59
	l.Mark()
	switch {
	default:
		goto yyrule59
	case c == '+' || c == '-' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate56
	case c == '\\':
		goto yystate57
	}

yystate57:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'U':
		goto yystate58
	case c == 'u':
		goto yystate62
	}

yystate58:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate59
	}

yystate59:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate60
	}

yystate60:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate61
	}

yystate61:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate62
	}

yystate62:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate63
	}

yystate63:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate64
	}

yystate64:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate65
	}

yystate65:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate55
	}

yystate66:
	c = l.Next()
	yyrule = 59
	l.Mark()
	switch {
	default:
		goto yyrule59
	case c == '+' || c == '-':
		goto yystate67
	case c == '.' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate56
	case c == '\\':
		goto yystate57
	case c >= '0' && c <= '9':
		goto yystate68
	}

yystate67:
	c = l.Next()
	yyrule = 59
	l.Mark()
	switch {
	default:
		goto yyrule59
	case c == '.' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate56
	case c == '\\':
		goto yystate57
	case c >= '0' && c <= '9':
		goto yystate68
	}

yystate68:
	c = l.Next()
	yyrule = 58
	l.Mark()
	switch {
	default:
		goto yyrule58
	case c == '.' || c >= 'A' && c <= 'D' || c >= 'G' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c == 'g' || c == 'h' || c == 'j' || c == 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate56
	case c == 'F' || c == 'L' || c == 'f' || c == 'l':
		goto yystate69
	case c == '\\':
		goto yystate57
	case c == 'i':
		goto yystate71
	case c >= '0' && c <= '9':
		goto yystate68
	}

yystate69:
	c = l.Next()
	yyrule = 58
	l.Mark()
	switch {
	default:
		goto yyrule58
	case c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'h' || c >= 'j' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate56
	case c == '\\':
		goto yystate57
	case c == 'i':
		goto yystate70
	}

yystate70:
	c = l.Next()
	yyrule = 58
	l.Mark()
	switch {
	default:
		goto yyrule58
	case c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate56
	case c == '\\':
		goto yystate57
	}

yystate71:
	c = l.Next()
	yyrule = 58
	l.Mark()
	switch {
	default:
		goto yyrule58
	case c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'G' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'g' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate56
	case c == 'F' || c == 'L' || c == 'f' || c == 'l':
		goto yystate70
	case c == '\\':
		goto yystate57
	}

yystate72:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate73
	case c == '/':
		goto yystate74
	case c == '=':
		goto yystate75
	}

yystate73:
	c = l.Next()
	yyrule = 3
	l.Mark()
	goto yyrule3

yystate74:
	c = l.Next()
	yyrule = 2
	l.Mark()
	switch {
	default:
		goto yyrule2
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate74
	}

yystate75:
	c = l.Next()
	yyrule = 20
	l.Mark()
	goto yyrule20

yystate76:
	c = l.Next()
	yyrule = 57
	l.Mark()
	switch {
	default:
		goto yyrule57
	case c == '.':
		goto yystate54
	case c == '8' || c == '9':
		goto yystate78
	case c == 'E' || c == 'e':
		goto yystate66
	case c == 'L':
		goto yystate79
	case c == 'P' || c == 'p':
		goto yystate56
	case c == 'U' || c == 'u':
		goto yystate82
	case c == 'X' || c == 'x':
		goto yystate86
	case c == '\\':
		goto yystate57
	case c == 'l':
		goto yystate85
	case c >= '0' && c <= '7':
		goto yystate77
	case c >= 'A' && c <= 'D' || c >= 'F' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'T' || c == 'V' || c == 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 't' || c == 'v' || c == 'w' || c == 'y' || c == 'z' || c == '\u0084':
		goto yystate55
	}

yystate77:
	c = l.Next()
	yyrule = 57
	l.Mark()
	switch {
	default:
		goto yyrule57
	case c == '.':
		goto yystate54
	case c == '8' || c == '9':
		goto yystate78
	case c == 'E' || c == 'e':
		goto yystate66
	case c == 'L':
		goto yystate79
	case c == 'P' || c == 'p':
		goto yystate56
	case c == 'U' || c == 'u':
		goto yystate82
	case c == '\\':
		goto yystate57
	case c == 'l':
		goto yystate85
	case c >= '0' && c <= '7':
		goto yystate77
	case c >= 'A' && c <= 'D' || c >= 'F' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0084':
		goto yystate55
	}

yystate78:
	c = l.Next()
	yyrule = 59
	l.Mark()
	switch {
	default:
		goto yyrule59
	case c == '.':
		goto yystate54
	case c == 'E' || c == 'e':
		goto yystate66
	case c == 'P' || c == 'p':
		goto yystate56
	case c == '\\':
		goto yystate57
	case c >= '0' && c <= '9':
		goto yystate78
	case c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate55
	}

yystate79:
	c = l.Next()
	yyrule = 57
	l.Mark()
	switch {
	default:
		goto yyrule57
	case c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate56
	case c == 'L':
		goto yystate80
	case c == 'U' || c == 'u':
		goto yystate81
	case c == '\\':
		goto yystate57
	}

yystate80:
	c = l.Next()
	yyrule = 57
	l.Mark()
	switch {
	default:
		goto yyrule57
	case c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate56
	case c == 'U' || c == 'u':
		goto yystate81
	case c == '\\':
		goto yystate57
	}

yystate81:
	c = l.Next()
	yyrule = 57
	l.Mark()
	switch {
	default:
		goto yyrule57
	case c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate56
	case c == '\\':
		goto yystate57
	}

yystate82:
	c = l.Next()
	yyrule = 57
	l.Mark()
	switch {
	default:
		goto yyrule57
	case c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate56
	case c == 'L':
		goto yystate83
	case c == '\\':
		goto yystate57
	case c == 'l':
		goto yystate84
	}

yystate83:
	c = l.Next()
	yyrule = 57
	l.Mark()
	switch {
	default:
		goto yyrule57
	case c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate56
	case c == 'L':
		goto yystate81
	case c == '\\':
		goto yystate57
	}

yystate84:
	c = l.Next()
	yyrule = 57
	l.Mark()
	switch {
	default:
		goto yyrule57
	case c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate56
	case c == '\\':
		goto yystate57
	case c == 'l':
		goto yystate81
	}

yystate85:
	c = l.Next()
	yyrule = 57
	l.Mark()
	switch {
	default:
		goto yyrule57
	case c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate56
	case c == 'U' || c == 'u':
		goto yystate81
	case c == '\\':
		goto yystate57
	case c == 'l':
		goto yystate80
	}

yystate86:
	c = l.Next()
	yyrule = 59
	l.Mark()
	switch {
	default:
		goto yyrule59
	case c == '.':
		goto yystate87
	case c == 'E' || c == 'e':
		goto yystate91
	case c == 'P' || c == 'p':
		goto yystate56
	case c == '\\':
		goto yystate57
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c == 'F' || c >= 'a' && c <= 'd' || c == 'f':
		goto yystate90
	case c >= 'G' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'g' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate55
	}

yystate87:
	c = l.Next()
	yyrule = 59
	l.Mark()
	switch {
	default:
		goto yyrule59
	case c == '.' || c >= 'G' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'g' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'e':
		goto yystate89
	case c == 'P' || c == 'p':
		goto yystate56
	case c == '\\':
		goto yystate57
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c == 'F' || c >= 'a' && c <= 'd' || c == 'f':
		goto yystate88
	}

yystate88:
	c = l.Next()
	yyrule = 59
	l.Mark()
	switch {
	default:
		goto yyrule59
	case c == '.' || c >= 'G' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'g' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'e':
		goto yystate89
	case c == 'P' || c == 'p':
		goto yystate66
	case c == '\\':
		goto yystate57
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c == 'F' || c >= 'a' && c <= 'd' || c == 'f':
		goto yystate88
	}

yystate89:
	c = l.Next()
	yyrule = 59
	l.Mark()
	switch {
	default:
		goto yyrule59
	case c == '+' || c == '-' || c == '.' || c >= 'G' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'g' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == 'E' || c == 'e':
		goto yystate89
	case c == 'P' || c == 'p':
		goto yystate66
	case c == '\\':
		goto yystate57
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c == 'F' || c >= 'a' && c <= 'd' || c == 'f':
		goto yystate88
	}

yystate90:
	c = l.Next()
	yyrule = 57
	l.Mark()
	switch {
	default:
		goto yyrule57
	case c == '.':
		goto yystate88
	case c == 'E' || c == 'e':
		goto yystate91
	case c == 'L':
		goto yystate79
	case c == 'P' || c == 'p':
		goto yystate66
	case c == 'U' || c == 'u':
		goto yystate82
	case c == '\\':
		goto yystate57
	case c == 'l':
		goto yystate85
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c == 'F' || c >= 'a' && c <= 'd' || c == 'f':
		goto yystate90
	case c >= 'G' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'g' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0084':
		goto yystate55
	}

yystate91:
	c = l.Next()
	yyrule = 57
	l.Mark()
	switch {
	default:
		goto yyrule57
	case c == '+' || c == '-' || c >= 'G' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'g' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0084':
		goto yystate55
	case c == '.':
		goto yystate88
	case c == 'E' || c == 'e':
		goto yystate91
	case c == 'L':
		goto yystate79
	case c == 'P' || c == 'p':
		goto yystate66
	case c == 'U' || c == 'u':
		goto yystate82
	case c == '\\':
		goto yystate57
	case c == 'l':
		goto yystate85
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c == 'F' || c >= 'a' && c <= 'd' || c == 'f':
		goto yystate90
	}

yystate92:
	c = l.Next()
	yyrule = 57
	l.Mark()
	switch {
	default:
		goto yyrule57
	case c == '.':
		goto yystate54
	case c == 'E' || c == 'e':
		goto yystate66
	case c == 'L':
		goto yystate79
	case c == 'P' || c == 'p':
		goto yystate56
	case c == 'U' || c == 'u':
		goto yystate82
	case c == '\\':
		goto yystate57
	case c == 'l':
		goto yystate85
	case c >= '0' && c <= '9':
		goto yystate92
	case c >= 'A' && c <= 'D' || c >= 'F' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0084':
		goto yystate55
	}

yystate93:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '>':
		goto yystate94
	}

yystate94:
	c = l.Next()
	yyrule = 21
	l.Mark()
	goto yyrule21

yystate95:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '%':
		goto yystate96
	case c == ':':
		goto yystate97
	case c == '<':
		goto yystate98
	case c == '=':
		goto yystate100
	}

yystate96:
	c = l.Next()
	yyrule = 22
	l.Mark()
	goto yyrule22

yystate97:
	c = l.Next()
	yyrule = 23
	l.Mark()
	goto yyrule23

yystate98:
	c = l.Next()
	yyrule = 24
	l.Mark()
	switch {
	default:
		goto yyrule24
	case c == '=':
		goto yystate99
	}

yystate99:
	c = l.Next()
	yyrule = 25
	l.Mark()
	goto yyrule25

yystate100:
	c = l.Next()
	yyrule = 26
	l.Mark()
	goto yyrule26

yystate101:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate102
	}

yystate102:
	c = l.Next()
	yyrule = 27
	l.Mark()
	goto yyrule27

yystate103:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate104
	case c == '>':
		goto yystate105
	}

yystate104:
	c = l.Next()
	yyrule = 28
	l.Mark()
	goto yyrule28

yystate105:
	c = l.Next()
	yyrule = 29
	l.Mark()
	switch {
	default:
		goto yyrule29
	case c == '=':
		goto yystate106
	}

yystate106:
	c = l.Next()
	yyrule = 30
	l.Mark()
	goto yyrule30

yystate107:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	}

yystate108:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'U':
		goto yystate109
	case c == 'u':
		goto yystate113
	}

yystate109:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate110
	}

yystate110:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate111
	}

yystate111:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate112
	}

yystate112:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate113
	}

yystate113:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate114
	}

yystate114:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate115
	}

yystate115:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate116
	}

yystate116:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate107
	}

yystate117:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '"':
		goto yystate118
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\'':
		goto yystate129
	case c == '\\':
		goto yystate108
	}

yystate118:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate119
	case c == '\\':
		goto yystate120
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate118
	}

yystate119:
	c = l.Next()
	yyrule = 53
	l.Mark()
	goto yyrule53

yystate120:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"' || c == '\'' || c >= '0' && c <= '7' || c == '?' || c == '\\' || c == 'a' || c == 'b' || c == 'f' || c == 'n' || c == 'r' || c == 't' || c == 'v':
		goto yystate118
	case c == 'U':
		goto yystate121
	case c == 'u':
		goto yystate125
	case c == 'x':
		goto yystate128
	}

yystate121:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate122
	}

yystate122:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate123
	}

yystate123:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate124
	}

yystate124:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate125
	}

yystate125:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate126
	}

yystate126:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate127
	}

yystate127:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate128
	}

yystate128:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate118
	}

yystate129:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '\\':
		goto yystate132
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate130
	}

yystate130:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate131
	case c == '\\':
		goto yystate132
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate130
	}

yystate131:
	c = l.Next()
	yyrule = 52
	l.Mark()
	goto yyrule52

yystate132:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"' || c == '\'' || c >= '0' && c <= '7' || c == '?' || c == '\\' || c == 'a' || c == 'b' || c == 'f' || c == 'n' || c == 'r' || c == 't' || c == 'v':
		goto yystate130
	case c == 'U':
		goto yystate133
	case c == 'u':
		goto yystate137
	case c == 'x':
		goto yystate140
	}

yystate133:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate134
	}

yystate134:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate135
	}

yystate135:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate136
	}

yystate136:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate137
	}

yystate137:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate138
	}

yystate138:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate139
	}

yystate139:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate140
	}

yystate140:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate130
	}

yystate141:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate142
	}

yystate142:
	c = l.Next()
	yyrule = 31
	l.Mark()
	goto yyrule31

yystate143:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate144
	case c == '|':
		goto yystate145
	}

yystate144:
	c = l.Next()
	yyrule = 32
	l.Mark()
	goto yyrule32

yystate145:
	c = l.Next()
	yyrule = 33
	l.Mark()
	goto yyrule33

yystate146:
	c = l.Next()
	yyrule = 6
	l.Mark()
	goto yyrule6

	goto yystate147 // silence unused label error
yystate147:
	c = l.Next()
yystart147:
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate149
	case c == '\u0080':
		goto yystate151
	case c >= '\x01' && c <= ')' || c >= '+' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate148
	}

yystate148:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate149
	case c >= '\x01' && c <= ')' || c >= '+' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate148
	}

yystate149:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate149
	case c == '/':
		goto yystate150
	case c >= '\x01' && c <= ')' || c >= '+' && c <= '.' || c >= '0' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate148
	}

yystate150:
	c = l.Next()
	yyrule = 4
	l.Mark()
	goto yyrule4

yystate151:
	c = l.Next()
	yyrule = 5
	l.Mark()
	goto yyrule5

	goto yystate152 // silence unused label error
yystate152:
	c = l.Next()
yystart152:
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate3
	case c == '"':
		goto yystate5
	case c == '#':
		goto yystate16
	case c == '%':
		goto yystate20
	case c == '&':
		goto yystate27
	case c == '*':
		goto yystate42
	case c == '+':
		goto yystate44
	case c == '-':
		goto yystate47
	case c == '.':
		goto yystate51
	case c == '/':
		goto yystate72
	case c == '0':
		goto yystate76
	case c == ':':
		goto yystate93
	case c == '<':
		goto yystate95
	case c == '=':
		goto yystate101
	case c == '>':
		goto yystate103
	case c == 'L':
		goto yystate164
	case c == '\'':
		goto yystate30
	case c == '\\':
		goto yystate155
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate2
	case c == '\u0080':
		goto yystate146
	case c == '^':
		goto yystate141
	case c == '|':
		goto yystate143
	case c >= '1' && c <= '9':
		goto yystate92
	case c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0084':
		goto yystate153
	}

yystate153:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate153
	case c == '(':
		goto yystate154
	case c == '\\':
		goto yystate155
	}

yystate154:
	c = l.Next()
	yyrule = 56
	l.Mark()
	goto yyrule56

yystate155:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'U':
		goto yystate156
	case c == 'u':
		goto yystate160
	}

yystate156:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate157
	}

yystate157:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate158
	}

yystate158:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate159
	}

yystate159:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate160
	}

yystate160:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate161
	}

yystate161:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate162
	}

yystate162:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate163
	}

yystate163:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate153
	}

yystate164:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '"':
		goto yystate118
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate153
	case c == '(':
		goto yystate154
	case c == '\'':
		goto yystate129
	case c == '\\':
		goto yystate155
	}

	goto yystate165 // silence unused label error
yystate165:
	c = l.Next()
yystart165:
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate3
	case c == '"':
		goto yystate5
	case c == '#':
		goto yystate16
	case c == '%':
		goto yystate20
	case c == '&':
		goto yystate27
	case c == '*':
		goto yystate42
	case c == '+':
		goto yystate44
	case c == '-':
		goto yystate47
	case c == '.':
		goto yystate51
	case c == '/':
		goto yystate72
	case c == '0':
		goto yystate76
	case c == ':':
		goto yystate93
	case c == '<':
		goto yystate95
	case c == '=':
		goto yystate101
	case c == '>':
		goto yystate103
	case c == 'L':
		goto yystate117
	case c == '\'':
		goto yystate30
	case c == '\\':
		goto yystate108
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate2
	case c == '\u0080':
		goto yystate146
	case c == '^':
		goto yystate141
	case c == 'd':
		goto yystate166
	case c == 'e':
		goto yystate172
	case c == 'i':
		goto yystate186
	case c == 'l':
		goto yystate206
	case c == 'p':
		goto yystate210
	case c == 'u':
		goto yystate216
	case c == '|':
		goto yystate143
	case c >= '1' && c <= '9':
		goto yystate92
	case c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'f' && c <= 'h' || c == 'j' || c == 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0084':
		goto yystate107
	}

yystate166:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'e':
		goto yystate167
	}

yystate167:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'f':
		goto yystate168
	}

yystate168:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'i':
		goto yystate169
	}

yystate169:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'n':
		goto yystate170
	}

yystate170:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'e':
		goto yystate171
	}

yystate171:
	c = l.Next()
	yyrule = 38
	l.Mark()
	switch {
	default:
		goto yyrule38
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	}

yystate172:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c == 'm' || c >= 'o' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'l':
		goto yystate173
	case c == 'n':
		goto yystate178
	case c == 'r':
		goto yystate182
	}

yystate173:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'i':
		goto yystate174
	case c == 's':
		goto yystate176
	}

yystate174:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'f':
		goto yystate175
	}

yystate175:
	c = l.Next()
	yyrule = 39
	l.Mark()
	switch {
	default:
		goto yyrule39
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	}

yystate176:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'e':
		goto yystate177
	}

yystate177:
	c = l.Next()
	yyrule = 40
	l.Mark()
	switch {
	default:
		goto yyrule40
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	}

yystate178:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'd':
		goto yystate179
	}

yystate179:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'i':
		goto yystate180
	}

yystate180:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'f':
		goto yystate181
	}

yystate181:
	c = l.Next()
	yyrule = 41
	l.Mark()
	switch {
	default:
		goto yyrule41
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	}

yystate182:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'r':
		goto yystate183
	}

yystate183:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'o':
		goto yystate184
	}

yystate184:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'r':
		goto yystate185
	}

yystate185:
	c = l.Next()
	yyrule = 42
	l.Mark()
	switch {
	default:
		goto yyrule42
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	}

yystate186:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'f':
		goto yystate187
	case c == 'n':
		goto yystate195
	}

yystate187:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'd':
		goto yystate188
	case c == 'n':
		goto yystate191
	}

yystate188:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'e':
		goto yystate189
	}

yystate189:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'f':
		goto yystate190
	}

yystate190:
	c = l.Next()
	yyrule = 44
	l.Mark()
	switch {
	default:
		goto yyrule44
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	}

yystate191:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'd':
		goto yystate192
	}

yystate192:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'e':
		goto yystate193
	}

yystate193:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'f':
		goto yystate194
	}

yystate194:
	c = l.Next()
	yyrule = 45
	l.Mark()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	}

yystate195:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'c':
		goto yystate196
	}

yystate196:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'l':
		goto yystate197
	}

yystate197:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'u':
		goto yystate198
	}

yystate198:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'd':
		goto yystate199
	}

yystate199:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'e':
		goto yystate200
	}

yystate200:
	c = l.Next()
	yyrule = 46
	l.Mark()
	switch {
	default:
		goto yyrule46
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == '_':
		goto yystate201
	}

yystate201:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'n':
		goto yystate202
	}

yystate202:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'e':
		goto yystate203
	}

yystate203:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'x':
		goto yystate204
	}

yystate204:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 't':
		goto yystate205
	}

yystate205:
	c = l.Next()
	yyrule = 47
	l.Mark()
	switch {
	default:
		goto yyrule47
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	}

yystate206:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'i':
		goto yystate207
	}

yystate207:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'n':
		goto yystate208
	}

yystate208:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'e':
		goto yystate209
	}

yystate209:
	c = l.Next()
	yyrule = 48
	l.Mark()
	switch {
	default:
		goto yyrule48
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	}

yystate210:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'r':
		goto yystate211
	}

yystate211:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'a':
		goto yystate212
	}

yystate212:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'g':
		goto yystate213
	}

yystate213:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'm':
		goto yystate214
	}

yystate214:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'a':
		goto yystate215
	}

yystate215:
	c = l.Next()
	yyrule = 49
	l.Mark()
	switch {
	default:
		goto yyrule49
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	}

yystate216:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'n':
		goto yystate217
	}

yystate217:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'd':
		goto yystate218
	}

yystate218:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'e':
		goto yystate219
	}

yystate219:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	case c == 'f':
		goto yystate220
	}

yystate220:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate107
	case c == '\\':
		goto yystate108
	}

	goto yystate221 // silence unused label error
yystate221:
	c = l.Next()
yystart221:
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate3
	case c == '"':
		goto yystate222
	case c == '#':
		goto yystate16
	case c == '%':
		goto yystate20
	case c == '&':
		goto yystate27
	case c == '*':
		goto yystate42
	case c == '+':
		goto yystate44
	case c == '-':
		goto yystate47
	case c == '.':
		goto yystate51
	case c == '/':
		goto yystate72
	case c == '0':
		goto yystate76
	case c == ':':
		goto yystate93
	case c == '<':
		goto yystate237
	case c == '=':
		goto yystate101
	case c == '>':
		goto yystate103
	case c == 'L':
		goto yystate117
	case c == '\'':
		goto yystate30
	case c == '\\':
		goto yystate108
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate2
	case c == '\u0080':
		goto yystate146
	case c == '^':
		goto yystate141
	case c == '|':
		goto yystate143
	case c >= '1' && c <= '9':
		goto yystate92
	case c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0084':
		goto yystate107
	}

yystate222:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate6
	case c == '\\':
		goto yystate225
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate223
	}

yystate223:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate224
	case c == '\\':
		goto yystate225
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate223
	}

yystate224:
	c = l.Next()
	yyrule = 51
	l.Mark()
	goto yyrule51

yystate225:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate228
	case c == 'U':
		goto yystate229
	case c == '\'' || c >= '0' && c <= '7' || c == '?' || c == '\\' || c == 'a' || c == 'b' || c == 'f' || c == 'n' || c == 'r' || c == 't' || c == 'v':
		goto yystate223
	case c == 'u':
		goto yystate233
	case c == 'x':
		goto yystate236
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '&' || c >= '(' && c <= '/' || c >= '8' && c <= '>' || c >= '@' && c <= 'T' || c >= 'V' && c <= '[' || c >= ']' && c <= '`' || c >= 'c' && c <= 'e' || c >= 'g' && c <= 'm' || c >= 'o' && c <= 'q' || c == 's' || c == 'w' || c >= 'y' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate226
	}

yystate226:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate227
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate226
	}

yystate227:
	c = l.Next()
	yyrule = 51
	l.Mark()
	goto yyrule51

yystate228:
	c = l.Next()
	yyrule = 51
	l.Mark()
	switch {
	default:
		goto yyrule51
	case c == '"':
		goto yystate6
	case c == '\\':
		goto yystate7
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate5
	}

yystate229:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate227
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate230
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '/' || c >= ':' && c <= '@' || c >= 'G' && c <= '`' || c >= 'g' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate226
	}

yystate230:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate227
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate231
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '/' || c >= ':' && c <= '@' || c >= 'G' && c <= '`' || c >= 'g' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate226
	}

yystate231:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate227
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate232
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '/' || c >= ':' && c <= '@' || c >= 'G' && c <= '`' || c >= 'g' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate226
	}

yystate232:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate227
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate233
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '/' || c >= ':' && c <= '@' || c >= 'G' && c <= '`' || c >= 'g' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate226
	}

yystate233:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate227
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate234
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '/' || c >= ':' && c <= '@' || c >= 'G' && c <= '`' || c >= 'g' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate226
	}

yystate234:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate227
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate235
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '/' || c >= ':' && c <= '@' || c >= 'G' && c <= '`' || c >= 'g' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate226
	}

yystate235:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate227
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate236
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '/' || c >= ':' && c <= '@' || c >= 'G' && c <= '`' || c >= 'g' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate226
	}

yystate236:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate227
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate223
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '/' || c >= ':' && c <= '@' || c >= 'G' && c <= '`' || c >= 'g' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate226
	}

yystate237:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '%':
		goto yystate239
	case c == ':':
		goto yystate240
	case c == '<':
		goto yystate241
	case c == '=':
		goto yystate243
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '$' || c >= '&' && c <= '9' || c == ';' || c >= '?' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate238
	}

yystate238:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '>':
		goto yystate227
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '=' || c >= '?' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate238
	}

yystate239:
	c = l.Next()
	yyrule = 22
	l.Mark()
	switch {
	default:
		goto yyrule22
	case c == '>':
		goto yystate227
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '=' || c >= '?' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate238
	}

yystate240:
	c = l.Next()
	yyrule = 23
	l.Mark()
	switch {
	default:
		goto yyrule23
	case c == '>':
		goto yystate227
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '=' || c >= '?' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate238
	}

yystate241:
	c = l.Next()
	yyrule = 24
	l.Mark()
	switch {
	default:
		goto yyrule24
	case c == '=':
		goto yystate242
	case c == '>':
		goto yystate227
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '<' || c >= '?' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate238
	}

yystate242:
	c = l.Next()
	yyrule = 25
	l.Mark()
	switch {
	default:
		goto yyrule25
	case c == '>':
		goto yystate227
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '=' || c >= '?' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate238
	}

yystate243:
	c = l.Next()
	yyrule = 26
	l.Mark()
	switch {
	default:
		goto yyrule26
	case c == '>':
		goto yystate227
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '=' || c >= '?' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate238
	}

yyrule1: // [ \t\f\v]+
	{
		return ' '
	}
yyrule2: // "//".*
	{
		l.comment(false)
		return ' '
	}
yyrule3: // "/*"
	{
		l.commentPos0 = l.First.Pos()
		l.push(scCOMMENT)
		goto yystate0
	}
yyrule4: // {comment-close}
	{
		l.pop()
		l.First = lex.NewChar(l.commentPos0, l.First.Rune)
		l.comment(true)
		return ' '
	}
yyrule5: // {eof}
	{
		l.report.Err(l.commentPos0, commentNotClosed)
		l.pop()
		return rune2class(lex.RuneEOF)
	}
yyrule6: // {eof}
	{
		return rune2class(lex.RuneEOF)
	}
yyrule7: // "!="
	{
		return NEQ
	}
yyrule8: // "%:"
	{
		return '#'
	}
yyrule9: // "%="
	{
		return MODASSIGN
	}
yyrule10: // "%>"
	{
		return '}'
	}
yyrule11: // "&&"
	{
		return ANDAND
	}
yyrule12: // "&="
	{
		return ANDASSIGN
	}
yyrule13: // "*="
	{
		return MULASSIGN
	}
yyrule14: // "++"
	{
		return INC
	}
yyrule15: // "+="
	{
		return ADDASSIGN
	}
yyrule16: // "--"
	{
		return DEC
	}
yyrule17: // "-="
	{
		return SUBASSIGN
	}
yyrule18: // "->"
	{
		return ARROW
	}
yyrule19: // "..."
	{
		return DDD
	}
yyrule20: // "/="
	{
		return DIVASSIGN
	}
yyrule21: // ":>"
	{
		return ']'
	}
yyrule22: // "<%"
	{
		return '{'
	}
yyrule23: // "<:"
	{
		return '['
	}
yyrule24: // "<<"
	{
		return LSH
	}
yyrule25: // "<<="
	{
		return LSHASSIGN
	}
yyrule26: // "<="
	{
		return LEQ
	}
yyrule27: // "=="
	{
		return EQ
	}
yyrule28: // ">="
	{
		return GEQ
	}
yyrule29: // ">>"
	{
		return RSH
	}
yyrule30: // ">>="
	{
		return RSHASSIGN
	}
yyrule31: // "^="
	{
		return XORASSIGN
	}
yyrule32: // "|="
	{
		return ORASSIGN
	}
yyrule33: // "||"
	{
		return OROR
	}
yyrule34: // "##"
yyrule35: // "#%:"
yyrule36: // "%:#"
yyrule37: // "%:%:"
	{
		return PPPASTE
	}
yyrule38: // "define"
	{
		l.pop()
		return PPDEFINE
		goto yystate0
	}
yyrule39: // "elif"
	{
		l.pop()
		return PPELIF
		goto yystate0
	}
yyrule40: // "else"
	{
		l.pop()
		return PPELSE
		goto yystate0
	}
yyrule41: // "endif"
	{
		l.pop()
		return PPENDIF
		goto yystate0
	}
yyrule42: // "error"
	{
		l.pop()
		return PPERROR
		goto yystate0
	}
yyrule43: // "if"
	{
		l.pop()
		return PPIF
		goto yystate0
	}
yyrule44: // "ifdef"
	{
		l.pop()
		return PPIFDEF
		goto yystate0
	}
yyrule45: // "ifndef"
	{
		l.pop()
		return PPIFNDEF
		goto yystate0
	}
yyrule46: // "include"
	{
		l.pop()
		return PPINCLUDE
		goto yystate0
	}
yyrule47: // "include_next"
	{
		l.pop()
		return PPINCLUDE_NEXT
		goto yystate0
	}
yyrule48: // "line"
	{
		l.pop()
		return PPLINE
		goto yystate0
	}
yyrule49: // "pragma"
	{
		l.pop()
		return PPPRAGMA
		goto yystate0
	}
yyrule50: // "undef"
	{
		l.pop()
		return PPUNDEF
		goto yystate0
	}
yyrule51: // {header-name}
	{
		l.sc = scINITIAL
		return PPHEADER_NAME
	}
yyrule52: // L{character-constant}
	{
		return LONGCHARCONST
	}
yyrule53: // L{string-literal}
	{
		return LONGSTRINGLITERAL
	}
yyrule54: // {character-constant}
	{
		return CHARCONST
	}
yyrule55: // {identifier}
	{
		return IDENTIFIER
	}
yyrule56: // {identifier}"("
	{
		return IDENTIFIER_LPAREN
	}
yyrule57: // {integer-constant}
	{
		return INTCONST
	}
yyrule58: // {floating-constant}
	{
		return FLOATCONST
	}
yyrule59: // {pp-number}
	{
		return PPNUMBER
	}
yyrule60: // {string-literal}
	{
		return STRINGLITERAL
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	if c, ok := l.Abort(); ok {
		return c
	}

	goto yyAction
}
