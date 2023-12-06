package main

import (
	"fmt"

	"modernc.org/cc/v4"
)

// Those are heavily inspired by github.com/xlab/c-for-go/blob/master/translator/ast_walker.go

func rawIdentifierOf(dd *cc.DirectDeclarator) string {
	switch dd.Case {
	case cc.DirectDeclaratorIdent:
		return dd.Token.SrcStr()
	case cc.DirectDeclaratorDecl:
		return rawIdentifierOf(dd.Declarator.DirectDeclarator)
	default:
		return rawIdentifierOf(dd.DirectDeclarator)
	}
}

func identifierOf(dd *cc.DirectDeclarator) (string, error) {
	id := rawIdentifierOf(dd)
	if id == "" {
		return "", fmt.Errorf("empty identifier for %+v", dd)
	}
	return id, nil
}

func paramName(n int, p *cc.Parameter) string {
	if len(p.Name()) == 0 {
		return fmt.Sprintf("arg%d", n)
	}
	return p.Name()
}

func memberName(n int, f *cc.Field) string {
	if len(f.Name()) == 0 {
		return fmt.Sprintf("field%d", n)
	}
	return f.Name()
}
