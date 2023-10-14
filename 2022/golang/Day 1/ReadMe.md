# Prompt
## part 1
The jungle must be too overgrown and difficult to navigate in vehicles or access from the air; the Elves' expedition traditionally goes on foot. As your boats approach land, the Elves begin taking inventory of their supplies. One important consideration is food - in particular, the number of Calories each Elf is carrying (your puzzle input).

The Elves take turns writing down the number of Calories contained by the various meals, snacks, rations, etc. that they've brought with them, one item per line. Each Elf separates their own inventory from the previous Elf's inventory (if any) by a blank line.

For example, suppose the Elves finish writing their items' Calories and end up with the following list:

1000
2000
3000

4000

5000
6000

7000
8000
9000

10000

This list represents the Calories of the food carried by five Elves:

The first Elf is carrying food with 1000, 2000, and 3000 Calories, a total of 6000 Calories.
The second Elf is carrying one food item with 4000 Calories.
The third Elf is carrying food with 5000 and 6000 Calories, a total of 11000 Calories.
The fourth Elf is carrying food with 7000, 8000, and 9000 Calories, a total of 24000 Calories.
The fifth Elf is carrying one food item with 10000 Calories.
In case the Elves get hungry and need extra snacks, they need to know which Elf to ask: they'd like to know how many Calories are being carried by the Elf carrying the most Calories. In the example above, this is 24000 (carried by the fourth Elf).

Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?

## Part 2 

By the time you calculate the answer to the Elves' question, they've already realized that the Elf carrying the most Calories of food might eventually run out of snacks.

To avoid this unacceptable situation, the Elves would instead like to know the total Calories carried by the top three Elves carrying the most Calories. That way, even if one of those Elves runs out of snacks, they still have two backups.

In the example above, the top three Elves are the fourth Elf (with 24000 Calories), then the third Elf (with 11000 Calories), then the fifth Elf (with 10000 Calories). The sum of the Calories carried by these three elves is 45000.

Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?

# Thoughts + notes
This was fun! Had some trouble getting all the types to work together. I certainly didn't need to many variables and this is definitly the brute force way. 

## explination
```go
    import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)
```
first we do the imports, I am trying to use only standand packages in the Advent of code. I was not familiar with bufio before this challange. 

``` go
func main (){
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
```
Then we open up by opening the file, and closing it when we are done, something I am sure will be a staple of the challenges. 
``` go
	scanner := bufio.NewScanner(input)
	
	count := int64(0)
	elves := []int64{}
	biggestBoi := int64(0)
	SecbiggestBoi := int64(0)
	ThirdbiggestBoi := int64(0)
	biggestBoiPlace := 0
	i := 0
```
Here is the meat. We set up the scanner which reads each line and (through scanner.Text()) returns the text of that line. Then I set up the (far too many) variables. 
- count: holds the current max of the current elf
- elves: is an array which I ultimatly didn't need but kept for prosparaty
- biggestBoi: holds the value of the highest count
- SecbiggestBoi: holds the value of the second highest count
- ThirdbiggestBoi: holds the value of the Third highest count
- BiggestBoiPlace: holds the value of the index of the highest count
- i: iterates the index

``` go
	for scanner.Scan() {
		if (scanner.Text() != "") {
			add, err := strconv.ParseInt(scanner.Text(), 0, 32)
			if err != nil {
				log.Fatal(err)
			}
			count = (count +  add);
```
Here we kick off a for loop, which runs for each line in the txt file (as read by bufio). In that for loop we check if the text is a number or == "". If it is a number we parse is string into an int with strconv, then assign that value to $add. If there are no errors we assign $count to $count + $add. The flow is,
- take a single line of txt 
- is it blank? 
- make it a number 
- add it to count 

The result is that for a block of numbers [15277, 4285, 10110, 7055] so long as they are continus we add them together. Ultimatly this adds up all of the calories being carried by a single elf.

- it1: 15227
- it2: 15277 + 4285 ($count = 19562) 
- it3: 19562 + 10110 ($count = 29672) 
- it4: 29672 + 7055 ($count = 36727) 

``` go
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
```
However if the scanner.text() does == blank, in other words if the current line in the txt file is blank. Then we know we need a) move onto the next elf and b) check if this elfe is holding the most. 

To do this we initialise the $biggestBoi variable and check if the current count is larger the value of $biggestBoi. If it is then we know that this elf is carrying more then whichever elf was previously carrying the most. So we set $biggestBoi to the count and now $biggestBoi's value is the largest value so far encounted. 

In part two the challenge require we also check the second and third largest elves. So we just run this spagetti logic to do that. Setting up some new variables SecBiggestBoi and ThirdbiggestBoi. Then when we find a count larger then $biggestBoi we waterfall the values down.

I ran into an issue here however. I assumed initially that if we found a count larger then biggestBoi then the second largest would simply be the previous largest. This resulted in my first response being to low. I had failed to think about the edge case in which we found a count which was not higher then biggestBoi but was higher then SecbiggestBoi. For instance, if the bellow where the values for individual elves. 
- 6789: set largest to 6789
- 15227: larger then 6789, update largest to 15227, update secornd largest to 6789
- 4285: not larger then 15227, no update
- 10110: not larger then 15227, no update
- 7055: not larger then 15227

This would result in the largest being accuaratly 15227 but the second largest being 6789. 

The fix for this is just to check count against each of the values each time. 

```go

	sumofFirstThree := biggestBoi + SecbiggestBoi + ThirdbiggestBoi
	fmt.Println(sumofFirstThree, biggestBoiPlace)
}
```
Final we print the total of all the bois and that is the answer.

