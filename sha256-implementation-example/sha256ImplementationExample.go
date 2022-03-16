package main

import "fmt"
import preprocessing "sha256-example/preprocessing"
import compression "sha256-example/compression"
import "os"
import "bufio"
import "strings"

func main() {

	fi, _ := os.Stdin.Stat()

	if (fi.Mode() & os.ModeCharDevice) == 0 {

		reader := bufio.NewReader(os.Stdin)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		fmt.Println(GetHash(input))
	} else {

		if len(os.Args) > 1 && len(os.Args) < 3 {

			fmt.Println(GetHash(os.Args[1]))

		} else {
			fmt.Println("Please provide a single input string")
		}
	}

}

func GetHash(input string) string {

	message := []byte(input)

	messageBlocks := preprocessing.BuildMessageBlocks(message)

	hashValues := compression.InitialValues

	for i := 0; i < len(messageBlocks)/64; i++ {

		messageBlock := make([]byte, 64)

		index := i * 64

		copy(messageBlock, messageBlocks[index:])

		messageSchedule := preprocessing.BuildMessageSchedule(messageBlock)

		hashValues = compression.Compression(messageSchedule,
			hashValues)
	}

	return PrintFinalHashValuesInHex(hashValues)

}

func PrintFinalHashValuesInHex(finalHashValues []uint32) string {

	hashValue := fmt.Sprintf("%x%x%x%x%x%x%x%x", finalHashValues[0],
		finalHashValues[1], finalHashValues[2], finalHashValues[3],
		finalHashValues[4], finalHashValues[5], finalHashValues[6],
		finalHashValues[7])

	return hashValue

}
