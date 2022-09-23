package main

import (
	"fmt"
	"dd/ff"
	dg "dd/gg"
	cs "dd/aa"
	di "dd/di"
)

func main() {
	m := make(map[int]int, 10)
	for i := 1; i<= 10; i++ {
		m[i] = i
	}

	for k, v := range(m) {
		go func(k int,v int) {
			fmt.Println("k ->", k, "v ->", v)
		}(k,v)
	}

	got := ff.NewIRuleConfigParser("json")
	if got!= nil {
		got.Parse("11")
	}
	fmt.Println(got)
	fmt.Println("data")

	animal := dg.NewCreateAnimal("tiger")
	if animal!= nil {
		animal.SeeAnimal().Eat("banner")
		animal.RaiseAnimal().Feel()
	}
	bb := cs.NewAa("bb")
	bb.Aaa()

	bbb := cs.NewAaFactory("cc")
	cc := bbb.SaaFactory()
	cc.Aaas()

	//todo di
	container := di.New()
	if err := container.Provide(NewA); err != nil {
		panic(err)
	}
	if err := container.Provide(NewB); err != nil {
		panic(err)
	}
	if err := container.Provide(NewC); err != nil {
		panic(err)
	}

	err := container.Invoke(func(a *A) {
		fmt.Printf("%+v: %d", a, a.B.C.Num)
	})
	if err != nil {
		panic(err)
	}
}
// A 依赖关系 A -> B -> C
type A struct {
	B *B
}

// NewA NewA
func NewA(b *B) *A {
	return &A{
		B: b,
	}
}

// B B
type B struct {
	C *C
}

// NewB NewB
func NewB(c *C) *B {
	return &B{C: c}
}

// C C
type C struct {
	Num int
}

// NewC NewC
func NewC() *C {
	return &C{
		Num: 1,
	}
}