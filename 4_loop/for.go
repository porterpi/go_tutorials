package main

func main() {
	// Simple for loop
	for i := 0; i < 10; i++ {
		println(i)
	}

	// Infinite loop
	// j := 0
	// for {
	// 	println(j)
	// }


	// For loop with continue
	k:=1
	for ; k<5; k++ {
		if k==2 {
			continue
		}
		println(k)
	}

	// For loop with break
	l:=1
	for ; l<5; l++ {
		if l==3 {
			break
		}
		println(l)
	}
}