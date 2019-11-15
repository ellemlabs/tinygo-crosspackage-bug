package test

import "./package1"
import "./package2"

func main() {
	_ = package2.MyStruct {
		MyField:package1.MyConst, // fails: cannot use package1.MyConst (constant 0 of type package1.MyInt) as package1.MyInt value in struct literal
	}
}

