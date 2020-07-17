package main

import (
	"fmt"
	"goft-expr/src/expr"
	"log"
)

type TestReulst struct{}

func (this *TestReulst) Name() string {
	return "test-result"
}
func main() {
	exprMap := map[string]interface{}{
		"test": func(name string, age int64) (string, int) {
			log.Println("this is ", name, " and  age is :", age)
			return "shenyi", 19
		},
		"test2": func() {
			log.Println("test")
		},
		"test3": func() *TestReulst {
			return &TestReulst{}
		},
	}
	fmt.Println(expr.BeanExpr("test('shenyi',19)", exprMap))
	fmt.Println(expr.BeanExpr("test2()", exprMap))
	fmt.Println(expr.BeanExpr("test3()", exprMap)[0].(*TestReulst).Name())
}
