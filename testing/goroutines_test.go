package testing

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestRunHelloWorld(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("UPS")

	time.Sleep(1 * time.Second)

}
