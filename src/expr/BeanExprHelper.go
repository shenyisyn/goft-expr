package expr

import (
	"github.com/shenyisyn/goft-expr/src/BeanExprLib"
	"reflect"
	"strconv"
	"strings"
)

func getValueByTokenType(t int, text string, this *beanExprListener) reflect.Value {
	var value reflect.Value
	switch t {
	case BeanExprLib.BeanExprLexerStringArg:
		stringArg := strings.Trim(text, "'")
		value = reflect.ValueOf(stringArg)
		break
	case BeanExprLib.BeanExprLexerIntArg:
		v, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			panic("parse int64 error")
		}
		value = reflect.ValueOf(v)
		break
	case BeanExprLib.BeanExprLexerFloatArg:
		v, err := strconv.ParseFloat(text, 64)
		if err != nil {
			panic("parse float64 error")
		}
		value = reflect.ValueOf(v)
		break
	case BeanExprLib.BeanExprLexerBoolArg:
		if text == "true" {
			value = reflect.ValueOf(true)
		} else {
			value = reflect.ValueOf(false)
		}
		break
	case BeanExprLib.BeanExprLexerNilArg:
		value = reflect.ValueOf(nil)
		break
	case BeanExprLib.BeanExprLexerFuncName | BeanExprLib.BeanExprLexerMethodName: //代表是函数
		get_ret := BeanExpr(text, this.exprMap)
		if get_ret.IsEmpty() {
			value = reflect.ValueOf(nil)
		} else {
			value = reflect.ValueOf(get_ret[0])
		}
		break
	default:
		return reflect.Value{}
	}
	return value
}
