package main

import "math/rand"

//func c(a, b int) int {
//	if b < 0 {
//		return a
//	}
//	return a + c(a, b-1)
//}

func c(a, b int) int {
	if b < 0 {
		return a
	}
	pa := &a
	pb := &b
	pc := *pa * *pb * rand.Intn(5)
	//println("pc:", pc, "pa:", *pa, "pb:", *pb)
	return pc + c(*pa, *pb-1)
}

func f(s string, a, b int) int {
	return a + b + len(s)
}
func b(ch chan int) int {
	str := "fskdjflkdjskfljdasfjfsadfasdfasfasdfdasfakldsjfkldasjflkasfldsalfjldkfjkaljfkfskdjflkdjskfljdasfjfsadfasdfasfasdfdasfakldsjfkldasjflkasfldsalfjldkfjkaljfkfskdjflkdjskfljdasfjfsadfasdfasfasdfdasfakldsjfkldasjflkasfldsalfjldkfjkaljfkfskdjflkdjskfljdasfjfsadfasdfasfasdfdasfakldsjfkldasjflkasfldsalfjldkfjkaljfk"
	a, b := 1, 20
	f(str, a, b)
	e, f := 1, 20
	c(e, f)
	g, g1 := 1, 200
	c(g, g1)
	k, b2 := 1, 200
	c(k, b2)
	a2, b := 1, 200
	c(a2, b)
	a4, b3 := 1, 200
	c(a4, b3)

	ch <- c(a, b)
	return c(a, b)
}

func main() {
	ch := make(chan int)
	go b(ch)
	<-ch
}
