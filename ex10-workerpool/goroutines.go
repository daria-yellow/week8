package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func worker(id int, jobs <-chan float64, results chan<- float64) {
	fmt.Printf("worker:%v spawning\n", id)
	for j := range jobs {
		fmt.Printf("worker:%v sleep:%v\n", id, j)
		str := int(j * 10000)
		time.Sleep(time.Millisecond * time.Duration(str))
		results <- j
	}
	fmt.Printf("worker:%v stopping\n", id)
}

func Run(poolSize int) {
	reader := bufio.NewScanner(os.Stdin)
	var a []string

	for reader.Scan() {
		time := reader.Text()
		a = append(a, time)
	}

	jobs := make(chan float64, poolSize)
	results := make(chan float64, poolSize)

	for w := 1; w <= poolSize; w++ {
		if len(a) < w {
			break
		} else {
			go worker(w, jobs, results)
		}
	}

	for j := 1; j <= len(a); j++ {
		s, _ := strconv.ParseFloat(strings.Trim(a[j-1], "\n"), 64)
		jobs <- s
	}

	close(jobs)

	for j := 1; j <= len(a); j++ {
		<-results
	}
}
