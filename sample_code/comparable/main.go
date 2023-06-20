package main

import "fmt"

type Doubler interface {
	Double()
}

type DoubleInt int

func (di *DoubleInt) Double() {
	*di = *di * 2
}

type DoubleIntSlice []int

func (dis DoubleIntSlice) Double() {
	for i := range dis {
		dis[i] = dis[i] * 2
	}
}

func DoubleAndPrint(d Doubler) {
	d.Double()
	fmt.Println(d)
}

func DoublerCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}
func main() {
	var di DoubleInt = 10
	var di2 DoubleInt = 10
	var dis = DoubleIntSlice{1, 2, 3}
	var dis2 = DoubleIntSlice{1, 2, 3}
	// false because we are comparing pointers,
	// and they point to different values
	DoublerCompare(&di, &di2)
	// false because they have different underlying types
	DoublerCompare(&di, dis)
	// triggers a panic, because the underlying types
	// match, but are a non-comparable type
	DoublerCompare(dis, dis2)
}
