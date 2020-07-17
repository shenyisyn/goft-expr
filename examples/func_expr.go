package examples

import (
	"fmt"
	"github.com/shenyisyn/goft-expr/src/expr"
	"log"
)

type TestReulst struct{}

func (this *TestReulst) Name() string {
	return "test-result"
}
func funcExpr() {
	exprMap := map[string]interface{}{
		"test": func(name string, age int64) (string, int) {
			log.Println("this is ", name, " and  age is :", age)
			return "shenyi", 19
		},
		"test2": func(str string) string {
			return "test2_" + str
		},
		"test3": func() *TestReulst {
			return &TestReulst{}
		},
		"test4": func(b bool) (int, bool) {
			return 1, b
		},
		"test5": func(f string) {
			log.Println(f)
		},
	}
	fmt.Println(expr.BeanExpr("test('shenyi',19)", exprMap))
	fmt.Println(expr.BeanExpr("test2('t2')", exprMap))
	fmt.Println(expr.BeanExpr("test3()", exprMap)[0].(*TestReulst).Name())
	fmt.Println(expr.BeanExpr("test4(true)", exprMap))
	fmt.Println(expr.BeanExpr("test5(test2('abc'))", exprMap))
}
