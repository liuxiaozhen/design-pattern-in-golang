package interpreter

import (
	"strconv"
	"strings"
)

type Expression interface {
	Interpret() int
}

type ValueExpression struct {
	val int
}

func (n *ValueExpression) Interpret() int {
	return n.val
}

type AddExpression struct {
	left, right Expression
}

func (n *AddExpression) Interpret() int {
	return n.left.Interpret() + n.right.Interpret()
}

type MinusExpression struct {
	left, right Expression
}

func (n *MinusExpression) Interpret() int {
	return n.left.Interpret() - n.right.Interpret()
}

type Parser struct {
	exp   []string
	index int
	prev  Expression
}

func (p *Parser) Parse(exp string) {
	p.exp = strings.Split(exp, " ")

	for {
		if p.index >= len(p.exp) {
			return
		}
		switch p.exp[p.index] {
		case "+":
			p.prev = p.newAddExpression()
		case "-":
			p.prev = p.newMinusExpression()
		default:
			p.prev = p.newValueExpression()
		}
	}
}

func (p *Parser) newAddExpression() Expression {
	p.index++
	return &AddExpression{
		left:  p.prev,
		right: p.newValueExpression(),
	}
}

func (p *Parser) newMinusExpression() Expression {
	p.index++
	return &MinusExpression{
		left:  p.prev,
		right: p.newValueExpression(),
	}
}

func (p *Parser) newValueExpression() Expression {
	v, _ := strconv.Atoi(p.exp[p.index])
	p.index++
	return &ValueExpression{
		val: v,
	}
}

func (p *Parser) Result() Expression {
	return p.prev
}
