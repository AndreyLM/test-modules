package testmodules

import "fmt"

// Hi - say hi
func Hi(name string) string {
	return fmt.Sprintf("hi, %s!", name)
}
