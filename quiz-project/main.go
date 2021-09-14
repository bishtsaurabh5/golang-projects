package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	var csvFile string
	var timeLimit int
	flag.StringVar(&csvFile, "csvFile", "", "input csv file")
	flag.IntVar(&timeLimit, "timeLimit", 10, "time limit per question for the quiz , default is 10 seconds")
	flag.Parse()
	// Open csv
	f, err := os.Open(csvFile)
	if err != nil {
		fmt.Println("Unable to open the file!!!")
		os.Exit(1)
	}
	//Since os.Open returns file which implements Reader interface
	data := csv.NewReader(f)
	lines, err := data.ReadAll()
	if err != nil {
		fmt.Println("Unable to parse csv")
		os.Exit(2)
	}

	qz := []quiz{}

	for _, line := range lines {
		tmp := quiz{}
		tmp.p = problem(line[0])
		tmp.q = answer(line[1])
		t, _ := strconv.Atoi(line[2])
		tmp.pt = points(t)
		qz = append(qz, tmp)
	}

	total := 0
	/*	for _, p := range qz {
			fmt.Printf(string(p.p))
			var ans string
			fmt.Scanf("%s\n", &ans)
			if ans == string(p.q) {
				total += int(p.pt)
			}
		}
	*/
	//Timed quiz now

	c := make(chan string, len(qz))
	for _, p := range qz {

		go func(quiz) {
			fmt.Printf(string(p.p))
			var ans string
			fmt.Scanf("%s\n", &ans)
			c <- ans
		}(p)
		// The above go routine runs separately and puts the answer to the channel
		// While we wait for the answer from the channel
		select {
		case o := <-c:
			if o == string(p.q) {
				total += int(p.pt)
			}
			break
		case <-time.After(time.Duration(timeLimit) * time.Second):
			fmt.Println("Time up !!!")
			break

		}

	}

	fmt.Println("Total Score is %s", total)

}

type problem string
type answer string
type points int

type quiz struct {
	p  problem
	q  answer
	pt points
}
