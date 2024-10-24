package basicConcepts

import (
	"basicConcepts/basics"
	_ "basicConcepts/basics"
	"fmt"
)

func main() {
	//arrays()
	//workingArrayTask()
	//slices()
	//sliceOperator()
	//buildInFunctionsSlices()
	//minCount()
	age, fullname := basics.Add(4, 5, "Olmosbek", "Urazboev")
	fmt.Println(age)
	fmt.Println(fullname)
}
