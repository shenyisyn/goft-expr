# Goft-Expr
* golang表达式解析器和执行
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

##安装
* go get -u github.com/shenyisyn/goft-expr

## 支持的参数
* 字符串(使用英文单引号，如 'abc')
* 数字（如 +123,-1,统一是int64类型,目前不支持科学计数)
* bool形态 , true或false(小写)
* 函数嵌套（字符串的形式)

##使用方法

* 执行函数
```go
type TestReulst struct{}
func (this *TestReulst) Name() string {
	return "test-result"
}
func main(){
	//定义个 map，里面包含了若干函数
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
	fmt.Println(expr.BeanExpr("test5(test2('abc'))", exprMap)) //嵌套执行
	}
```
返回值
```go
2020/07/18 00:32:26 this is  shenyi  and  age is : 19
[shenyi 19]
[test2_t2]
test-result
[1 true]
2020/07/18 00:32:26 abc
[]


```

## License
© jtthink, 2020~time.Now
Released under the [Apache License 2.0](https://github.com/shenyisyn/goft-expr/blob/master/LICENSE)
