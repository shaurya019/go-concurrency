import "fmt"


var signals = []string{"test"}


var wg sync.WaitGroup //pointer
var mut sync.Mutex    // pointer

func main() {
	// Here go is for gorountines
	// Goroutines are lightweight threads managed by the Go runtime. They allow you to run functions concurrently, making it easier to handle multiple tasks simultaneously without the overhead of traditional threads.

	// Key Points
	// Concurrency vs. Parallelism: Goroutines enable concurrent programming. Concurrency is about dealing with multiple tasks at once, while parallelism is about executing multiple tasks simultaneously. Goroutines can run in parallel if multiple CPU cores are available.

	// Lightweight: Goroutines are much lighter than traditional threads. They use a very small amount of memory and the Go runtime can manage hundreds of thousands of goroutines efficiently.

	// Communication: Goroutines can communicate with each other using channels, which provide a way to send and receive values between goroutines safely.

	// go help("hello")
	// help("world")

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.wait()

	// How WaitGroups Work
	// Add: Increment the WaitGroup counter to indicate that new goroutines are being added.
	// Done: Decrement the WaitGroup counter to indicate that a goroutine has finished its work.
	// Wait: Block until the WaitGroup counter goes back to zero, indicating that all goroutines have completed their work.

	// Explanation
	// worker function: This function simulates a worker by printing a start message, sleeping for one second (to simulate some work), and then printing a done message. The defer wg.Done() statement ensures that the WaitGroup counter is decremented when the goroutine finishes, even if the function exits early.

	// WaitGroup Initialization: In the main function, a WaitGroup variable wg is declared.

	// Adding Workers: A loop is used to start multiple worker goroutines. Before starting each goroutine, wg.Add(1) increments the WaitGroup counter.

	// Waiting for Completion: After starting all the worker goroutines, wg.Wait() is called to block the main function until all worker goroutines have called wg.Done() and the counter has returned to zero.

	// Output: The program prints messages indicating when each worker starts and finishes, and finally, a message indicating that all workers are done.

}

func getStatusCode(endpoint string) {
	defer wg.Done()

	//  the defer statement is used to ensure that a function call is performed later in a program's execution, typically for purposes of cleanup.

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}

// func help(string x) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3*time.Milisecond)
// 		fmt.Println(x)
// 	}
// }