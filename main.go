import "fmt"

func main() {
	// Here go is for gorountines
	// Goroutines are lightweight threads managed by the Go runtime. They allow you to run functions concurrently, making it easier to handle multiple tasks simultaneously without the overhead of traditional threads.

	// Key Points
// Concurrency vs. Parallelism: Goroutines enable concurrent programming. Concurrency is about dealing with multiple tasks at once, while parallelism is about executing multiple tasks simultaneously. Goroutines can run in parallel if multiple CPU cores are available.

// Lightweight: Goroutines are much lighter than traditional threads. They use a very small amount of memory and the Go runtime can manage hundreds of thousands of goroutines efficiently.

// Communication: Goroutines can communicate with each other using channels, which provide a way to send and receive values between goroutines safely.



	go help("hello")
	help("world")
}

func help(string x) {
	for i := 0; i < 6; i++ {
		time.Sleep(3*time.Milisecond)
		fmt.Println(x)
	}
}