package package2

import "../package1"

type MyStruct struct {
	MyField package1.MyInt
}

// This work around removes the compiler's abilty to do any type-checking on myfield
// thus will not result in the compile-time error.
func WorkAround(mystruct *MyStruct, myfield interface{}) {
	myfieldFromInterface,_ := myfield.(package1.MyInt)
	mystruct.MyField        = myfieldFromInterface
}


// Trying to use this funciton in main will result in the same bug.
func NotAWorkAround(mystruct *MyStruct, myfield package1.MyInt) {
	mystruct.MyField = myfield
}
