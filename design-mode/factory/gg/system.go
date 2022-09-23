package gg
import "fmt"
type animal interface {
	Eat(food string)
}

type tiger struct {
}
type cat struct {
}
func (e tiger) Eat(food string) {
	fmt.Println("tiger eat ",food)
}
func (e cat) Eat(food string) {
	fmt.Println("cat eat",food)
}
/*func NewAnimal(z string) animal {
	switch z {
	case "tiger":
		return tiger{}
	case "cat":
		return cat{}
	}
	return nil
}*/

type ranimal interface {
	Feel()
}
func (e tiger) Feel() {
	fmt.Println("tiger Feel ")
}
func (e cat) Feel() {
	fmt.Println("cat Feel")
}
type CAnimal interface {
	SeeAnimal() animal
	RaiseAnimal() ranimal
}
type tigerAnimal struct {

}

func (t tigerAnimal) SeeAnimal() animal {
	return tiger{}
}
func (t tigerAnimal) RaiseAnimal() ranimal {
	return tiger{}
}
func NewCreateAnimal(a string) CAnimal {
	switch a {
	case "tiger":
		return tigerAnimal{}
	}
	return nil
}
