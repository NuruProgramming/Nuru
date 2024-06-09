package ast

import (
	"bytes"
	"strings"

	"github.com/NuruProgramming/Nuru/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) expressionNode()      {}
func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }
func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("kama")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("sivyo")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token
	Name       string
	Parameters []*Identifier
	Defaults   map[string]Expression
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}

	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

type CallExpression struct {
	Token     token.Token
	Function  Expression // can be Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

type DictLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}

func (dl *DictLiteral) expressionNode()      {}
func (dl *DictLiteral) TokenLiteral() string { return dl.Token.Literal }
func (dl *DictLiteral) String() string {
	var out bytes.Buffer
	pairs := []string{}
	for key, value := range dl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("(")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

type Assign struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ae *Assign) expressionNode()      {}
func (ae *Assign) TokenLiteral() string { return ae.Token.Literal }
func (ae *Assign) String() string {
	var out bytes.Buffer

	out.WriteString(ae.Name.String())
	out.WriteString(ae.TokenLiteral())
	out.WriteString(ae.Value.String())

	return out.String()
}

type AssignEqual struct {
	Token token.Token
	Left  *Identifier
	Value Expression
}

func (ae *AssignEqual) expressionNode()      {}
func (ae *AssignEqual) TokenLiteral() string { return ae.Token.Literal }
func (ae *AssignEqual) String() string {
	var out bytes.Buffer

	out.WriteString(ae.Left.String())
	out.WriteString(ae.TokenLiteral())
	out.WriteString(ae.Value.String())

	return out.String()
}

type AssignmentExpression struct {
	Token token.Token
	Left  Expression
	Value Expression
}

func (ae *AssignmentExpression) expressionNode()      {}
func (ae *AssignmentExpression) TokenLiteral() string { return ae.Token.Literal }
func (ae *AssignmentExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ae.Left.String())
	out.WriteString(ae.TokenLiteral())
	out.WriteString(ae.Value.String())

	return out.String()
}

type WhileExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
}

func (we *WhileExpression) expressionNode()      {}
func (we *WhileExpression) TokenLiteral() string { return we.Token.Literal }
func (we *WhileExpression) String() string {
	var out bytes.Buffer

	out.WriteString("wakati")
	out.WriteString(we.Condition.String())
	out.WriteString(" ")
	out.WriteString(we.Consequence.String())

	return out.String()
}

type Null struct {
	Token token.Token
}

func (n *Null) expressionNode()      {}
func (n *Null) TokenLiteral() string { return n.Token.Literal }
func (n *Null) String() string       { return n.Token.Literal }

type Break struct {
	Statement
	Token token.Token // the 'break' token
}

func (b *Break) expressionNode()      {}
func (b *Break) TokenLiteral() string { return b.Token.Literal }
func (b *Break) String() string       { return b.Token.Literal }

type Continue struct {
	Statement
	Token token.Token // the 'continue' token
}

func (c *Continue) expressionNode()      {}
func (c *Continue) TokenLiteral() string { return c.Token.Literal }
func (c *Continue) String() string       { return c.Token.Literal }

type PostfixExpression struct {
	Token    token.Token
	Operator string
}

func (pe *PostfixExpression) expressionNode()      {}
func (pe *PostfixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PostfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Token.Literal)
	out.WriteString(pe.Operator)
	out.WriteString(")")
	return out.String()
}

type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (fl *FloatLiteral) expressionNode()      {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FloatLiteral) String() string       { return fl.Token.Literal }

type For struct {
	Token        token.Token
	Identifier   string      // "i"
	StarterName  *Identifier // i = 0
	StarterValue Expression
	Closer       Expression // i++
	Condition    Expression // i < 1
	Block        *BlockStatement
}

type ForIn struct {
	Token    token.Token
	Key      string
	Value    string
	Iterable Expression
	Block    *BlockStatement
}

func (fi *ForIn) expressionNode()      {}
func (fi *ForIn) TokenLiteral() string { return fi.Token.Literal }
func (fi *ForIn) String() string {
	var out bytes.Buffer

	out.WriteString("kwa ")
	if fi.Key != "" {
		out.WriteString(fi.Key + ", ")
	}
	out.WriteString(fi.Value + " ")
	out.WriteString("ktk ")
	out.WriteString(fi.Iterable.String() + " {\n")
	out.WriteString("\t" + fi.Block.String())
	out.WriteString("\n}")

	return out.String()
}

type CaseExpression struct {
	Token   token.Token
	Default bool
	Expr    []Expression
	Block   *BlockStatement
}

func (ce *CaseExpression) expressionNode()      {}
func (ce *CaseExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CaseExpression) String() string {
	var out bytes.Buffer

	if ce.Default {
		out.WriteString("kawaida ")
	} else {
		out.WriteString("ikiwa ")

		tmp := []string{}
		for _, exp := range ce.Expr {
			tmp = append(tmp, exp.String())
		}
		out.WriteString(strings.Join(tmp, ","))
	}
	out.WriteString(ce.Block.String())
	return out.String()
}

type SwitchExpression struct {
	Token   token.Token
	Value   Expression
	Choices []*CaseExpression
}

func (se *SwitchExpression) expressionNode()      {}
func (se *SwitchExpression) TokenLiteral() string { return se.Token.Literal }
func (se *SwitchExpression) String() string {
	var out bytes.Buffer
	out.WriteString("\nbadili (")
	out.WriteString(se.Value.String())
	out.WriteString(")\n{\n")

	for _, tmp := range se.Choices {
		if tmp != nil {
			out.WriteString(tmp.String())
		}
	}
	out.WriteString("}\n")

	return out.String()
}

type MethodExpression struct {
	Token     token.Token
	Object    Expression
	Method    Expression
	Arguments []Expression
	Defaults  map[string]Expression
}

func (me *MethodExpression) expressionNode()      {}
func (me *MethodExpression) TokenLiteral() string { return me.Token.Literal }
func (me *MethodExpression) String() string {
	var out bytes.Buffer
	out.WriteString(me.Object.String())
	out.WriteString(".")
	out.WriteString(me.Method.String())

	return out.String()
}

type Import struct {
	Token       token.Token
	Identifiers map[string]*Identifier
}

func (i *Import) expressionNode()      {}
func (i *Import) TokenLiteral() string { return i.Token.Literal }
func (i *Import) String() string {
	var out bytes.Buffer
	out.WriteString("tumia ")
	for k := range i.Identifiers {
		out.WriteString(k + " ")
	}
	return out.String()
}

type PackageBlock struct {
	Token      token.Token
	Statements []Statement
}

func (pb *PackageBlock) statementNode()       {}
func (pb *PackageBlock) TokenLiteral() string { return pb.Token.Literal }
func (pb *PackageBlock) String() string {
	var out bytes.Buffer

	for _, s := range pb.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type Package struct {
	Token token.Token
	Name  *Identifier
	Block *BlockStatement
}

func (p *Package) expressionNode()      {}
func (p *Package) TokenLiteral() string { return p.Token.Literal }
func (p *Package) String() string {
	var out bytes.Buffer

	out.WriteString("pakeji " + p.Name.Value + "\n")
	out.WriteString("::\n")
	for _, s := range p.Block.Statements {
		out.WriteString(s.String())
	}
	out.WriteString("\n::")

	return out.String()
}

type At struct {
	Token token.Token
}

func (a *At) expressionNode()      {}
func (a *At) TokenLiteral() string { return a.Token.Literal }
func (a *At) String() string       { return "@" }

type PropertyAssignment struct {
	Token token.Token // the '=' token
	Name  *PropertyExpression
	Value Expression
}

func (pa *PropertyAssignment) expressionNode()      {}
func (pa *PropertyAssignment) TokenLiteral() string { return pa.Token.Literal }
func (pa *PropertyAssignment) String() string       { return "Ngl I'm tired" }

type PropertyExpression struct {
	Expression
	Token    token.Token // The . token
	Object   Expression
	Property Expression
}

func (pe *PropertyExpression) expressionNode()      {}
func (pe *PropertyExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PropertyExpression) String() string       { return "Ngl I'm tired part two" }
