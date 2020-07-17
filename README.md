# Goft-Expr
golang表达式解析器和执行
## 目前功能
* 函数表达式解析
* struct方法执行

## 后续功能(开发中)
* 注解解析（类Java)
* 简易规则引擎

##版本要求
* golang >=1.12
 
##教学视频
* http://b.jtthink.com/read.php?tid=549&fid=2

##使用方法
执行函数
```go
type TestReulst struct {}
func(this *TestReulst) Name() string{
	return "test-result"
}
exprMap:= map[string]interface{}{
		"test": func(name string,age int64) (string,int) {
			log.Println("this is ",name," and  age is :",age)
			return "shenyi",19
		},
		"test2": func()  {
			log.Println("test")
		},
		"test3": func() *TestReulst {
			return &TestReulst{}
		},
	}

fmt.Println(expr.BeanExpr("test('shenyi',19)",exprMap))
fmt.Println(expr.BeanExpr("test2()",exprMap))
fmt.Println(expr.BeanExpr("test3()",exprMap)[0].(*TestReulst).Name())
```
返回值
```go
 this is  shenyi  and  age is : 19
 [shenyi 19]
test
 []
 test-result

```

## License
© jtthink, 2020~time.Now
Released under the [Apache License 2.0](https://github.com/shenyisyn/goft-expr/blob/master/LICENSE)
