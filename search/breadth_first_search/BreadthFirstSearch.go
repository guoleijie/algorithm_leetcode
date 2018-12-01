package main

import "fmt"

//广度优先算法，适用于有向图
func main() {
	//新增一个有向图。
	//将一段放入数组中，循环将二段放入数组中
	map_variable := make(map[string][]string)
	map_variable["A"] = []string{"B", "C"}
	map_variable["B"] = []string{"C", "D"}
	map_variable["C"] = []string{"D", "D"}
	map_variable["D"] = []string{"E"}

	start := "A"
	end := "E"
	runpath := map_variable[start]
	for i := 0; i < len(runpath); i++ {
		if end == runpath[i] {
			fmt.Println("finished")
			return
		}
		runpath = append(runpath, map_variable[runpath[i]]...)
	}
}

//A   B     D      E
//    C
