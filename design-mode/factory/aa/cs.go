package aa
import "fmt"

type aa interface {
	Aaa()
}
type bb struct {

}
type cc struct {

}
type aaa interface {
	Aaas()
}
type bbb struct {

}
type ccc struct {

}
func (a bbb) Aaas()  {
	fmt.Println("bbbb")
}
func (a ccc) Aaas()  {
	fmt.Println("cccc")
}

func (a bb) Aaa()  {
	fmt.Println("bb")
}
func (a cc) Aaa()  {
	fmt.Println("cc")
}
func NewAa(t string) aa { //简单工厂
	switch t{
	case "bb":
		return bb{}
	case "cc":
		return cc{}
	}
	return nil
}
//工厂方法
type aaFactory interface {
	AaaFactory() aa
	SaaFactory() aaa
}
type bbFactory struct {
}
type ccFactory struct {
}

func (f bbFactory) AaaFactory() aa {
	return bb{}
}

func (f ccFactory) AaaFactory() aa {
	return cc{}
}

func (f bbFactory) SaaFactory() aaa {
	return bbb{}
}

func (f ccFactory) SaaFactory() aaa {
	return ccc{}
}
func NewAaFactory(t string) aaFactory {
	switch t{
	case "bb":
		return bbFactory{}
	case "cc":
		return ccFactory{}
	}
	return nil
}
