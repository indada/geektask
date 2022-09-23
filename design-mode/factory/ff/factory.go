package ff

import "fmt"

type IRuleConfigParser interface { //规则解析器
	Parse(data string)
}
type jsonRuleConfigParser struct { //json规则解析器
}
type yamRuleConfigParser struct { //yam规则解析器
}
func (j yamRuleConfigParser) Parse(data string)  {
	fmt.Println(data)
}
func (j jsonRuleConfigParser) Parse(data string)  {
	fmt.Println(data)
}
func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return jsonRuleConfigParser{}
	case "yam":
		return yamRuleConfigParser{}
	}
	return nil
}

