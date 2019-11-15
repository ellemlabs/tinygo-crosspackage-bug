package main

import "./package1"
import "./package2"

func main() {
	_ = package2.MyStruct {
		MyField:package1.MyConst, // fails: cannot use package1.MyConst (constant 0 of type package1.MyInt) as package1.MyInt value in struct literal
	}

	mystruct := package2.MyStruct {}
	package2.NotAWorkAround(&mystruct, package1.MyConst) // fails: cannot use package1.MyConst (constant 0 of type package1.MyInt) as package1.MyInt value in argument to package2.NotAWorkAround
	package2.WorkAround(&mystruct, package1.MyConst) // works.
}

