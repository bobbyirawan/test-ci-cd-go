package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldBobby(t *testing.T) {
	result := HelloWorld("Bobby")
	if result != "Hello Bobby" {
		// unit test failed
		t.Error("Result must be 'Bello Bobby ")
	}

	fmt.Println("TestHelloWorldBobby Done")
}

func TestHelloWorldIrawan(t *testing.T) {
	result := HelloWorld("Irawan")
	if result != "Hello Irawan" {
		t.Fatal("Result must be 'Hello Irawan'")
	}
	fmt.Println("TestHelloWorldIrawan Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Bobby")
	assert.Equal(t, "Hello Bobby", result, "Result  must be 'Hello Bobby'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Bobby")
	require.Equal(t, "Hello Bobby", result, "Result Must be Hello Bobby")
	fmt.Println("TestHelloWorld with Require Done")
}
