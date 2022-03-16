package compression

import "fmt"

import "testing"
import "sha256-example/constants"
import "sha256-example/preprocessing"
import "sha256-example/logicalFunctions"

func TestCompression(t *testing.T) {

	messageBlocks := preprocessing.BuildMessageBlocks([]byte("abc"))

	fmt.Printf("Output: %b\n", []byte("abc"))

	schedule := preprocessing.BuildMessageSchedule(messageBlocks)

	initialValues := constants.InitializeRegisters()

	//registers := initialValues

	constants := constants.GetConstants()

	//registers
	a := initialValues[0]
	b := initialValues[1]
	c := initialValues[2]
	d := initialValues[3]
	e := initialValues[4]
	f := initialValues[5]
	g := initialValues[6]
	h := initialValues[7]

	for i := 0; i < 64; i++ {

		//calc temp words
		tempWord1 := h + logicalFunctions.BigSigmaOne(e) + logicalFunctions.Choice(e, f, g) + constants[i] + schedule[i]
		tempWord2 := logicalFunctions.BigSigmaZero(a) + logicalFunctions.Majority(a, b, c)

		h = g
		g = f
		f = e
		e = d + tempWord1
		d = c
		c = b
		b = a
		a = tempWord1 + tempWord2
	}

	fmt.Printf("Output: [%x, %x, %x, %x, %x, %x, %x, %x]\n", a, b, c, d, e, f, g, h)

	hv0 := a + initialValues[0]
	hv1 := b + initialValues[1]
	hv2 := c + initialValues[2]
	hv3 := d + initialValues[3]
	hv4 := e + initialValues[4]
	hv5 := f + initialValues[5]
	hv6 := g + initialValues[6]
	hv7 := h + initialValues[7]

	fmt.Printf("Output: [%x, %x, %x, %x, %x, %x, %x, %x]\n", hv0, hv1, hv2, hv3, hv4, hv5, hv6, hv7)

	result := fmt.Sprintf("%x%x%x%x%x%x%x%x", hv0, hv1, hv2, hv3, hv4, hv5, hv6, hv7)

	expected := "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"

	if result != expected {
		fmt.Printf("Expected: %v\nResult:   %v\n", expected, result)
	}
}
