package directedgraphs

import (
	"strconv"
	"fmt"
)

type Path struct {
	initial string
	target string
	weight int
}

func newPath(p string) *Path {
	result:= new(Path)
	result.initial = string([]rune(p)[0])
	result.target = string([]rune(p)[1])
	result.weight, _ = strconv.Atoi(string([]rune(p)[2]))
	return result
}


func Problem1(input []string) string {
	paths := parse(input)
	i:=gimmeWeight(paths, "A", "B")
	d:=gimmeWeight(paths, "B", "C")
	if (i==-1 || d==-1) {
		return "NO SUCH ROUTE"
	}
	return strconv.Itoa(i+d)
}

func Problem2(input []string) string {
	paths := parse(input)
	i:=gimmeWeight(paths, "A", "D")
	if (i < 0) {
		return "NO SUCH ROUTE"
	}
	return strconv.Itoa(i)
}

func Problem3(input []string) string {
	paths := parse(input)
	i:=gimmeWeight(paths, "A", "D")
	d:=gimmeWeight(paths, "D", "C")
	if (i==-1 || d==-1) {
		return "NO SUCH ROUTE"
	}
	return strconv.Itoa(i+d)
}

func Problem4(input []string) string {
	paths := parse(input)
	i:=gimmeWeight(paths, "A", "E")
	d:=gimmeWeight(paths, "E", "B")
	f:=gimmeWeight(paths, "B", "C")
	k:=gimmeWeight(paths, "C", "D")
	if (i==-1 || d==-1 || f==-1 || k ==-1 ) {
		return "NO SUCH ROUTE"
	}
	return strconv.Itoa(i+d+f+k)
}

func Problem5(input []string) string {
	paths := parse(input)
	i:=gimmeWeight(paths, "A", "E")
	d:=gimmeWeight(paths, "E", "D")
	if (i==-1 || d==-1) {
		return "NO SUCH ROUTE"

	}
	return strconv.Itoa(i+d)
}

func gimmeWeight(paths []Path, source string, target string) int {
	for i:=0; i<len(paths); i++ {
		if(paths[i].initial==source && paths[i].target==target) {
			return paths[i].weight
		}
	}
	return -1
}


func parse(input []string) (paths []Path) {
	for i:=0; i<len(input); i++ {
		paths=append(paths, *newPath(input[i]))
	}
	return
}

func printMap(paths []Path) {
	for i:=0; i<len(paths); i++ {
		fmt.Printf("%s %s %d \n", paths[i].initial, paths[i].target, paths[i].weight)
	}
}

