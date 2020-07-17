package expr

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/shenyisyn/goft-expr/src/BeanExprLib"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func BeanExpr(expr string, exprMap map[string]interface{}) ResultSet {
	is := antlr.NewInputStream(expr)
	lexer := BeanExprLib.NewBeanExprLexer(is)
	ts := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := BeanExprLib.NewBeanExprParser(ts)
	lis := &beanExprListener{exprMap: exprMap}
	antlr.ParseTreeWalkerDefault.Walk(lis, p.Start())
	return lis.Run()
}

type ResultSet []interface{}

func newResultSet() ResultSet {
	return make(ResultSet, 0)
}
func (this ResultSet) IsEmpty() bool {
	return len(this) == 0
}
func result(values []reflect.Value) ResultSet {
	ret := newResultSet()
	if values == nil || len(values) == 0 {
		return ret
	}
	for _, v := range values {
		ret = append(ret, v.Interface())
	}
	return ret
}

type beanExprListener struct {
	*BeanExprLib.BaseBeanExprListener
	funcName   string
	args       []reflect.Value
	methodName string //方法名 user.abc.bcd.getage()
	execType   uint8  //执行类型  0代表函数 也就是默认值， 1代表struct执行
	exprMap    map[string]interface{}
}

func (this *beanExprListener) ExitMethodCall(ctx *BeanExprLib.MethodCallContext) {
	this.execType = 1
	this.methodName = ctx.GetStart().GetText()

}
func (this *beanExprListener) ExitFuncCall(ctx *BeanExprLib.FuncCallContext) {
	this.funcName = ctx.GetStart().GetText()
}
func (this *beanExprListener) ExitFuncArgs(ctx *BeanExprLib.FuncArgsContext) {
	for i := 0; i < ctx.GetChildCount(); i++ {
		token := ctx.GetChild(i).GetPayload().(*antlr.CommonToken)
		var value reflect.Value
		switch token.GetTokenType() {
		case BeanExprLib.BeanExprLexerStringArg:
			stringArg := strings.Trim(token.GetText(), "'")
			value = reflect.ValueOf(stringArg)
			break
		case BeanExprLib.BeanExprLexerIntArg:
			v, err := strconv.ParseInt(token.GetText(), 10, 64)
			if err != nil {
				panic("parse int64 error")
			}
			value = reflect.ValueOf(v)
			break
		case BeanExprLib.BeanExprLexerFloatArg:
			v, err := strconv.ParseFloat(token.GetText(), 64)
			if err != nil {
				panic("parse float64 error")
			}
			value = reflect.ValueOf(v)
			break
		default:
			continue
		}
		this.args = append(this.args, value)
	}

}
func (this *beanExprListener) findField(method string, v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if field := v.FieldByName(method); field.IsValid() {
		return field
	}
	return reflect.Value{}
}
func (this *beanExprListener) Run() ResultSet {
	if this.exprMap == nil {
		panic("exprMap required")
	}
	switch this.execType {
	case 0:
		if f, ok := this.exprMap[this.funcName]; ok {
			v := reflect.ValueOf(f)
			if v.Kind() == reflect.Func {
				return result(v.Call(this.args))
			}
		}
		break
	case 1: // struct方法执行
		ms := strings.Split(this.methodName, ".")
		if obj, ok := this.exprMap[ms[0]]; ok {
			objv := reflect.ValueOf(obj)
			current := objv
			for i := 1; i < len(ms); i++ {
				if i == len(ms)-1 { //最后一个是方法名
					if method := current.MethodByName(ms[i]); !method.IsValid() {
						panic("method error:" + ms[i])
					} else {
						return result(method.Call(this.args))
					}
				}
				field := this.findField(ms[i], current)
				if field.IsValid() {
					current = field
				} else {
					panic("field error:" + ms[i])
				}
			}
		}
	default:
		log.Println("nothing to do")
	}
	return newResultSet()

}
