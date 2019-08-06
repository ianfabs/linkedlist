package main

import "fmt"

func main() {
	ll := IntSLList{}
	ll.addToHead(6)
	ll.addToHead(4)
	ll.addToTail(20)
	fmt.Println(ll.min())
	ll.printAll()
}
