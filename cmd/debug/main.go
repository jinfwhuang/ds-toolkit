package main

import (
	tmplog "log"
)

func init() {
	tmplog.SetFlags(tmplog.Llongfile)
}

/*

 */
func main() {

	ch := make(chan int, 5)

	//go func() {
	//	for {
	//		select {
	//		case msg := <-ch:
	//			tmplog.Println(msg)
	//		}
	//	}
	//}()

	for i := 1; i < 3; i++ {
		select {
		case ch <- i: {
			tmplog.Println("sent", i)
		}
		}
		//ch <- i
	}

	tmplog.Println("fff")

}






