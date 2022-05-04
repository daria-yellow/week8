package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func worker(id int, jobs <-chan float64, wg *sync.WaitGroup) {
	wg.Add(1)
	fmt.Printf("worker:%v spawning\n", id)
	for j := range jobs {
		fmt.Printf("worker:%v sleep:%.1f\n", id, j)
		str := int(j * 10000)
		time.Sleep(time.Millisecond * time.Duration(str))
	}
	fmt.Printf("worker:%v stopping\n", id)
	wg.Done()
}

func Run(poolSize int) {
	var wg sync.WaitGroup
	jobs := make(chan float64, poolSize)
	reader := bufio.NewScanner(os.Stdin)
	i := 1
	for reader.Scan() {
		time := reader.Text()
		s, _ := strconv.ParseFloat(strings.Trim(time, "\n"), 64)
		jobs <- s
		if poolSize >= i {
			go worker(i, jobs, &wg)
			i += 1
		}
	}
	close(jobs)
	wg.Wait()
}
