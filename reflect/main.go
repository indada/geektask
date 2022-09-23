package main

import (
	"fmt"
	"reflect"
	)
type Demo struct {

}

func fh(i int) (int,error) {
	return i,nil
}
func isError(t reflect.Type) bool {
	if t.Kind() != reflect.Interface {
		return false
	}
	//todo ???
	return t.Implements(reflect.TypeOf((*error)(nil)).Elem())
}
//反射
func main() {
	dd := reflect.TypeOf(Demo{})
	fmt.Println("type: ", dd.Name(), dd.Kind()) //type:  Demo struct
	fmt.Println("value: ",reflect.ValueOf(Demo{})) //value:  {}
	var ii int16 = 88
	reflect_type(ii)
	fmt.Println("---Implements--------")
	fmt.Println(reflect.TypeOf((*error)(nil)))
	fhv := reflect.ValueOf(fh)
	fhvt := fhv.Type()
	results := make([]reflect.Type,fhvt.NumOut())
	for i := 0; i < fhvt.NumOut(); i++ {
		results[i] = fhvt.Out(i)
	}
	for _, result := range results {
		fmt.Println(isError(result))
	}
	fmt.Println("---Implements---------")
	fmt.Println(reflect.TypeOf((*error)(nil)).Kind())
	fmt.Println(reflect.TypeOf(reflect.TypeOf((*error)(nil)).Elem()))
	var ff float64 = 8.8
	reflect_set_type(&ff)

	fmt.Println("---反射获取结构体成员类型----")
	type cat struct {
		Name string
		Type int `json:"type" id:"100"`
	}
	//创建cat的实例
	ins := cat{Name: "mimi",Type: 2}

	//获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	fmt.Println(typeOfCat.FieldByName("Name"))
	for i := 0; i < typeOfCat.NumField(); i++ {
		fieldType := typeOfCat.Field(i)
		fmt.Println(fieldType.Tag.Get("id"))
		fmt.Println(fieldType.Name,fieldType.Tag)

		//通过字段名找到对应信息
		if catType,ok := typeOfCat.FieldByName(fieldType.Name);ok{
			fmt.Println(catType)
		}
	}
	fmt.Println("====反射获取值===")
	var a int = 1024
	fmt.Println(a)
	//获取反射值对象
	valueOfa := reflect.ValueOf(ins)
	valueOfaa := reflect.ValueOf(a)
	fmt.Println(valueOfa.Interface().(cat).Name) //通过类型断言转换为int类型
	fmt.Println(valueOfaa.Int()) //通过强制类型转换

	fmt.Println("___反射访问结构体成员的值____")
	type dummy struct {
		a int
		b string
		//嵌入字段
		float32
		bool
		next *dummy
	}
	//值包装结构体
	d := reflect.ValueOf(dummy{
		next: &dummy{},
	})
	td := reflect.TypeOf(dummy{
		next: &dummy{},
	})
	//获取字段数量
	valNumField := d.NumField()
	fmt.Println("num：",valNumField)
	for i := 0; i < valNumField; i++ {
		fmt.Println("字段类型：",d.Field(i).Type())
		fmt.Println("字段key：",td.Field(i).Name)
		fmt.Println("字段值：",d.Field(i))
	}

	fmt.Println("+++判断反射值的空和有效性+++")
	//空指针
	var q *int
	fmt.Println("var q *int:",reflect.ValueOf(q).IsNil())
	fmt.Println("nil:",reflect.ValueOf(nil).IsValid())

	//通过反射调用函数
	//将函数包装为反射值对象
	funcValue := reflect.ValueOf(add)
	//构造函数参数，传入两个整形值
	paramList := []reflect.Value{reflect.ValueOf(10),reflect.ValueOf(11)}
	//反射调用函数
	retList := funcValue.Call(paramList)
	fmt.Println(retList[0].Int())
	//反射结构体调用函数

	getValue := reflect.ValueOf(typeCall{1,"dada"})
	methodValue := getValue.MethodByName("Ccc")
	argList := []reflect.Value{reflect.ValueOf(1000),reflect.ValueOf("daaaaa")}
	methodValue.Call(argList)
}
type typeCall struct {
	Id int
	Name string
}
func (t typeCall) Ccc(a int,b string)  {
	fmt.Println("t:",t.Id,t.Name,a,b)
}
func add(s int,d int) int {
	return s+d
}
	//反射interface类型
func reflect_type(a interface{}){
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	k := t.Kind()
	switch k {
	case reflect.Float64:
		fmt.Println(v.Float())
	case reflect.Int16:
		fmt.Println(v.Int()) //88
	}
}

//修改反射类型 只能修改指针类型
func reflect_set_type(a interface{}) {
	v := reflect.ValueOf(a)
	t := reflect.TypeOf(a)
	//取元素
	te := t.Elem()
	fmt.Println("te ",te)
	fmt.Println("te.Kind ",te.Kind())
	fmt.Println("te.Name ",te.Name())
	k := t.Kind()
	fmt.Println("Kind ",k)
	switch k {
	case reflect.Float64:
		v.SetFloat(64.64)
		fmt.Println(v.Float())
	case reflect.Ptr:
		fmt.Println("Ptr")
		v.Elem().SetFloat(6.6)
		fmt.Println(v.Elem().Float())
	}
}
