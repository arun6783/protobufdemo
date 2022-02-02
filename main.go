package main

import (
	"fmt"

	"github.com/arun6783/protobufdemo/src/protos/generated/protobufdemo/complex"
	"github.com/arun6783/protobufdemo/src/protos/generated/protobufdemo/enumerations"
	"github.com/arun6783/protobufdemo/src/protos/generated/protobufdemo/simple"
)

func main() {

	doSimple()
	enumsDemo()
	complexDemo()

}

func complexDemo() {

	cm := complex.ComplexMessage{
		OneDummy:      &complex.DummyMessage{Id: 1, Name: "Arun"},
		MultipleDummy: []*complex.DummyMessage{{Id: 2, Name: "inside comples"}, {Id: 3, Name: "inside comples 3"}},
	}

	fmt.Println(cm)
}

func enumsDemo() {
	em := enumerations.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumerations.DayOfTheWeek_MONDAY,
	}

	fmt.Println(em)
}

func doSimple() {
	sm := simple.SimpleMessage{
		IsSimple:   true,
		Name:       "My Simple message",
		SampleList: []int32{1, 3, 4, 5, 6, 8},
	}

	fmt.Println(sm)

	sm.Name = "i renamed you"

	fmt.Println(sm)

	fmt.Println("the id is ", sm.GetId())
}
