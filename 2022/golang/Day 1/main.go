package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main (){
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)
	
	count := int64(0)
	elves := []int64{}
	biggestBoi := int64(0)
	SecbiggestBoi := int64(0)
	ThirdbiggestBoi := int64(0)
	biggestBoiPlace := 0
	i := 0

	for scanner.Scan() {
		if (scanner.Text() != "") {
			add, err := strconv.ParseInt(scanner.Text(), 0, 32)
			if err != nil {
				log.Fatal(err)
			}
			count = (count +  add);
		} else {
			i++;
			elves = append(elves, count);
			if count > biggestBoi {
				ThirdbiggestBoi = SecbiggestBoi
				SecbiggestBoi = biggestBoi
				biggestBoi = count
				
				biggestBoiPlace = i
			} else {
				if count > SecbiggestBoi {
					ThirdbiggestBoi = SecbiggestBoi
					SecbiggestBoi = count
				} else {
					if count > ThirdbiggestBoi{
						ThirdbiggestBoi = count
					}
				}
			}
			count = 0
		}
	}
	sumofFirstThree := biggestBoi + SecbiggestBoi + ThirdbiggestBoi
	fmt.Println(sumofFirstThree, biggestBoiPlace)
	
}