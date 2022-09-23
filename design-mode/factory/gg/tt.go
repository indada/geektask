package gg
// IRuleConfigParser IRuleConfigParser
type IRuleConfigParser interface { //规则配置解析
	Parse(data []byte)//解析方法
}

// jsonRuleConfigParser jsonRuleConfigParser
type jsonRuleConfigParser struct{} //json规则配置解析

// Parse Parse
func (j jsonRuleConfigParser) Parse(data []byte) { //json规则配置解析 实现解析方法
	panic("implement me")
}

// ISystemConfigParser ISystemConfigParser
type ISystemConfigParser interface { //系统配置解析
	ParseSystem(data []byte) // 系统解析
}

// jsonSystemConfigParser jsonSystemConfigParser
type jsonSystemConfigParser struct{} //json系统配置解析

// Parse Parse
func (j jsonSystemConfigParser) ParseSystem(data []byte) { //json系统配置解析 实现系统解析
	panic("implement me")
}

// IConfigParserFactory 工厂方法接口
type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISystemConfigParser
}

type jsonConfigParserFactory struct{} //json工厂方法解析接口

func (j jsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

func (j jsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return jsonSystemConfigParser{}
}
