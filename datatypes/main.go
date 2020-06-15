package main

import (
	"fmt"
)

func createAllPerms(varTypes []string, maxLen int) {
	if maxLen != 5 {
		panic("Maxlen should be five")
	}
	extendedVarTypes := varTypes

	for _, type1 := range varTypes {
		extendedVarTypes = append(extendedVarTypes, type1+"[]")
		for i := 0; i < 5; i++ {
			extendedVarTypes = append(extendedVarTypes, type1+fmt.Sprintf("[%d]", i))
		}
	}
	fmt.Println(len(extendedVarTypes))
	/*
		fmt.Println("var extendedVarTypes = [][]string{ ")
		for _, type1 := range extendedVarTypes {
			for _, type2 := range extendedVarTypes {
				for _, type3 := range extendedVarTypes {
					for _, type4 := range extendedVarTypes {
						for _, type5 := range extendedVarTypes {
							fmt.Println("{ \"" + type1 + "\", \"" + type2 + "\", \"" + type3 + "\", \"" + type4 + "\", \"" + type5 + "\"},")
						}
					}
				}
			}
		}
		fmt.Println("}")
	*/
}

func main() {
	//varTypes := []string{"bool", "address"}

	varTypes := []string{"bool", "address", "bytes", "string",
		"uint8", "int8", "uint8", "int8", "uint16", "int16",
		"uint24", "int24", "uint32", "int32", "uint40", "int40", "uint48", "int48", "uint56", "int56",
		"uint64", "int64", "uint72", "int72", "uint80", "int80", "uint88", "int88", "uint96", "int96",
		"uint104", "int104", "uint112", "int112", "uint120", "int120", "uint128", "int128", "uint136", "int136",
		"uint144", "int144", "uint152", "int152", "uint160", "int160", "uint168", "int168", "uint176", "int176",
		"uint184", "int184", "uint192", "int192", "uint200", "int200", "uint208", "int208", "uint216", "int216",
		"uint224", "int224", "uint232", "int232", "uint240", "int240", "uint248", "int248", "uint256", "int256",
		"bytes1", "bytes2", "bytes3", "bytes4", "bytes5", "bytes6", "bytes7", "bytes8", "bytes9", "bytes10", "bytes11",
		"bytes12", "bytes13", "bytes14", "bytes15", "bytes16", "bytes17", "bytes18", "bytes19", "bytes20", "bytes21",
		"bytes22", "bytes23", "bytes24", "bytes25", "bytes26", "bytes27", "bytes28", "bytes29", "bytes30", "bytes31",
		"bytes32", "bytes"}

	createAllPerms(varTypes, 5)
	fmt.Println(len(varTypes))
}
