package main

func main() {
	msgCh := make(chan int, 3)

	msgCh <- 1
	msgCh <- 2
	msgCh <- 3
	msgCh <- 4

	// go func ()  {
	// 	fmt.Println(<-msgCh)
	// }()
}
