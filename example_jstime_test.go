package jstime_test

import (
	"fmt"
	"github.com/Cherrison/jstime"
	"time"
)

var times int

func ExampleSetInterval() {

	go jstime.SetInterval(time.Second, func() {
		fmt.Printf("qps %d\n", times)
		times = 0
	})
	// Output:
	//测试qps
	//qps 38040
	//qps 40590
	//qps 40649
	//qps 40941
	//qps 40083
}

func ExampleSetInterval2() {

	jstime.SetInterval(time.Second, func() {
		fmt.Println("1s pass")
	})
	// Output:
	//1s pass
	//1s pass
	//1s pass
	//1s pass
}

func ExampleSetTimeout() {
	jstime.SetInterval(time.Second, func() {
		fmt.Println("1s later")
	})
	// Output:
	// 1s later
}