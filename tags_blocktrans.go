package pongo2trans

import (
	"github.com/flosch/pongo2"
	"bytes"
)

type tagBlockTransNode struct {
	translator	ITranslator
	wrapper		*pongo2.NodeWrapper
	withPairs	map[string]pongo2.IEvaluator
}

func (self *tagBlockTransNode) Execute(ctx *pongo2.ExecutionContext, buffer *bytes.Buffer) error {
	buf := bytes.NewBuffer([]byte{})
	if err := self.wrapper.Execute(ctx, buf); err != nil {
		return err
	}

	// Result will be the same string as input
	result := buf.String()

	// Then we will try to translate it
	if self.translator != nil {
		args := make(map[string]interface {})

		// Put all custom with-pairs into the map
		for key, value := range self.withPairs {
			val, err := value.Evaluate(ctx)
			if err != nil {
				return err
			}

			// Set variable
			if val.IsBool() {
				args[key] = val.Bool()
			} else if val.IsString() {
				args[key] = val.String()
			} else if val.IsInteger() {
				args[key] = val.Integer()
			} else if val.IsFloat() {
				args[key] = val.Float()
			} else if val.IsNil() {
				args[key] = nil
			} else {
				args[key] = val.Interface()
			}
		}

		if res, ok := self.translator.Translate(result, args); ok {
			result = res
		}
	}

	buffer.WriteString(result)
	return nil
}

func tagBlockTransParser(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, error) {
	node := &tagBlockTransNode{
		translator:	translator,
		withPairs:	make(map[string]pongo2.IEvaluator),
	}

	// Wrap till the end
	wrapper, endArgs, err := doc.WrapUntilTag("endblocktrans")
	if err != nil {
		return nil, err
	}
	node.wrapper = wrapper

	// Arguments for ending tag is not allowed here
	if endArgs.Count() > 0 {
		return nil, endArgs.Error("Arguments not allowed here.", nil)
	}

	// OMG, here is some arguments
	if arguments.Count() > 0 {
		// Check `with` keywork
		if arguments.MatchOne(pongo2.TokenIdentifier, "with") == nil {
			return nil, arguments.Error("Expected 'with' keyword.", nil)
		}

		// Ok got it, lets check each variable
		for arguments.Remaining() > 0 {
			var key string

			// Name of variable
			keyToken := arguments.MatchType(pongo2.TokenIdentifier)
			if keyToken == nil {
				stringToken := arguments.MatchType(pongo2.TokenString)
				if stringToken == nil {
					return nil, arguments.Error("Expected an identifier or string identifier", nil)
				} else {
					key = stringToken.Val
				}
			} else {
				key = keyToken.Val
			}

			// You what it is
			if arguments.Match(pongo2.TokenSymbol, "=") == nil {
				return nil, arguments.Error("Expected '='.", nil)
			}

			// And expression
			valueExpr, err := arguments.ParseExpression()
			if err != nil {
				return nil, err
			}
			node.withPairs[key] = valueExpr
		}
	}

	// Check Remaining arguments
	if arguments.Remaining() > 0 {
		return nil, arguments.Error("Malformed blocktrans-tag.", nil)
	}

	return node, nil
}

func init() {
	pongo2.RegisterTag("blocktrans", pongo2.TagParser(tagBlockTransParser))
}
