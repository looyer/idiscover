package main

import (
	"log"
	"os"
	"runtime"
	"runtime/debug"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic:%v \nstack:\n%v", err, string(debug.Stack()))
			os.Exit(-1)
		}
	}()

	log.Printf("~~~~~~~~~ go-libuse ~~~~~~~~~")
	log.Printf("now routine num:%v", runtime.NumGoroutine())

	/////////////////////////////////////////////////////////////////
	//Panic
	panic("a problem")
	/////////////////////////////////////////////////////////////////

	// /////////////////////////////////////////////////////////////////
	// //Sorting by Functions
	// fruits := []string{"apple", "pear", "graph", "cherry"}
	// lenCmp := func(a, b string) int {
	// 	return cmp.Compare(len(a), len(b))
	// }
	// slices.SortFunc(fruits, lenCmp)
	// log.Printf("fruits:%v", fruits)
	// /////////////////////////////////////////////////////////////////

	// /////////////////////////////////////////////////////////////////
	// //Sorting
	// strs := []string{"carrot", "apple", "banana"}
	// slices.Sort(strs)
	// log.Printf("strs:%v", strs)

	// ints := []int{3, 7, 1, 5, 4}
	// slices.Sort(ints)
	// slices.Reverse(ints)
	// log.Printf("ints:%v, isSorted:%v", ints, slices.IsSorted(ints))
	// /////////////////////////////////////////////////////////////////

}
