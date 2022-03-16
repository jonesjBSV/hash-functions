package preprocessing

import "testing"
import "encoding/binary"
import "ripemd160-example/preprocessing"
import "fmt"

func TestBuildMessageSchedule(t *testing.T) {

	messageBlock := preprocessing.BuildMessageBlocks([]byte("abc"))

	var schedule [16]uint32

	i := 0
	for j := 0; j < 16; j++ {

		schedule[j] = binary.LittleEndian.Uint32(messageBlock[i:])

		i += 4
	}

	fmt.Printf("Output: %x\n", schedule)

	if len(schedule) == 16 {
		for i := 0; i < 16; i++ {
			t.Errorf("%x\n", schedule[i])
		}
	}

}
