package pongo2trans

import (
	"bytes"
	pongo2 "gopkg.in/flosch/pongo2.v3"
)

type tagTransNode struct {
	translator      ITranslator
	asName          string
	valueExpression pongo2.IEvaluator
}

func (self *tagTransNode) Execute(ctx *pongo2.ExecutionContext, buffer *bytes.Buffer) *pongo2.Error {
	// Execute expression
	value, err := self.valueExpression.Evaluate(ctx)
	if err != nil {
		return err
	}

	// Result will be the same string as input
	result := value.String()

	// Then we will try to translate it
	if self.translator != nil {
		if res, ok := self.translator.Translate(value.String(), nil); ok {
			result = res
		}
	}

	// Set variable
	if self.asName != "" {
		ctx.Private[self.asName] = result
	} else {
		buffer.WriteString(result)
	}

	return nil
}

func tagTransParser(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	node := &tagTransNode{
		translator: translator,
	}

	// Save source
	valueExpr, err := arguments.ParseExpression()
	if err != nil {
		return nil, err
	}
	node.valueExpression = valueExpr

	// As
	if arguments.Remaining() > 0 {
		if arguments.MatchOne(pongo2.TokenKeyword, "as") != nil {
			nameToken := arguments.MatchType(pongo2.TokenIdentifier)
			if nameToken == nil {
				return nil, arguments.Error("Name (identifier) expected after 'as'.", nil)
			}
			node.asName = nameToken.Val
		}
	}

	if arguments.Remaining() > 0 {
		return nil, arguments.Error("Malformed trans-tag.", nil)
	}

	return node, nil
}

func init() {
	pongo2.RegisterTag("trans", pongo2.TagParser(tagTransParser))
}
