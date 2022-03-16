package preprocessing

import "encoding/binary"

func BuildMessageSchedule(messageBlock []byte) [16]uint32 {

	var schedule [16]uint32

	i := 0
	for j := 0; j < 16; j++ {

		schedule[j] = binary.LittleEndian.Uint32(messageBlock[i:])

		i += 4
	}

	return schedule

}
