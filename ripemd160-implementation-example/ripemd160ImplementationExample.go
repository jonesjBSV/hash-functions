package main

import "fmt"
import preprocessing "ripemd160-example/preprocessing"
import compression "ripemd160-example/compression"
import "os"
import "bufio"
import "strings"
import "encoding/binary"

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

func PrintFinalHashValuesInHex(hashValues []uint32) string {

	finalValue := make([]byte, 20)

	binary.LittleEndian.PutUint32(finalValue[0:], hashValues[0])
	binary.LittleEndian.PutUint32(finalValue[4:], hashValues[1])
	binary.LittleEndian.PutUint32(finalValue[8:], hashValues[2])
	binary.LittleEndian.PutUint32(finalValue[12:], hashValues[3])
	binary.LittleEndian.PutUint32(finalValue[16:], hashValues[4])

	return fmt.Sprintf("%x", finalValue)
}
