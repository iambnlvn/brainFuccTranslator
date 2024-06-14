package main

import (
	"fmt"
	"io/ioutil"
	"os"
)
//making sure the memo tape starts clean
func initializeTape(tape []int8) {
    for i := range tape {
        tape[i] = 0
    }
}

func main() {
    const tapeSize = 30000
    var (
        A             [tapeSize]int8 
        dataPointer   int            
        instructionPtr int           
    )

    initializeTape(A[:])

	inputBytes, err := ioutil.ReadFile(os.Args[1])
	    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
	inputString := string(inputBytes)

	for instructionPtr < len(inputString) {
		switch inputString[instructionPtr] {
		case '>':
			if dataPointer < tapeSize-1 {
			dataPointer++
			}
        case '<':
            if dataPointer > 0 {
                dataPointer--
            }
        case '+':
            A[dataPointer]++
        case '-':
            A[dataPointer]--
        case '.':
            fmt.Printf("%c", A[dataPointer])
		case ',':
			var inputChar byte
            fmt.Scanf("%c", &inputChar)
            A[dataPointer] = int8(inputChar)

		case '[':
			if A[dataPointer] == 0 {
				stack := 0
				for j := instructionPtr + 1; ; j++ {
					if inputString[j] == '[' {
						stack++
					} else if inputString[j] == ']' {
						if stack != 0 {
							stack--
						} else {
							instructionPtr = j
							break
						}
					}
				}
			}

		case ']':
			if A[dataPointer] != 0 {
				stack := 0
				for j := instructionPtr - 1; ; j-- {
					if inputString[j] == ']' {
						stack++
					} else if inputString[j] == '[' {
						if stack != 0 {
							stack--
						} else {
							instructionPtr = j
							break
						}
					}
				}
			}
		}
		instructionPtr++
	}
}
