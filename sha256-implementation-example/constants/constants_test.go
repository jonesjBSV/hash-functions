package constants

import "testing"
import "math"
import "fmt"

var First_sixty_four_primes [64]int = [64]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311}

func TestGetConstants(t *testing.T) {

	constants := []uint32{}

	for _, element := range First_sixty_four_primes {

		root := math.Pow(float64(element), (1 / 3.0))
		value, _ := math.Modf(root * (1 << 32))
		hexValue := uint32(value)

		constants = append(constants, hexValue)

	}

	expected := uint32(0x428a2f98)

	if constants[0] != expected {
		t.Errorf("Expected value: %x\n Result: %x", expected, constants[0])
	}

	fmt.Printf("%x\n", constants)

}

func TestInitializeRegisters(t *testing.T) {

	registerValues := []uint32{}

	for i := 0; i < 8; i++ {
		root := math.Pow(float64(First_sixty_four_primes[i]), (1 / 2.0))
		value, _ := math.Modf(root * (1 << 32))
		hexValue := uint32(value)

		registerValues = append(registerValues, hexValue)
	}

	expected := uint32(0x6a09e667)

	if registerValues[0] != expected {
		t.Errorf("Expected value: %x\n Result: %x", expected, registerValues[0])
	}

	fmt.Printf("%x\n", registerValues)

	testValue := registerValues[0] | registerValues[1]

	fmt.Printf("%32b\n %32b\n %32b", registerValues[0], registerValues[1], testValue)

}
