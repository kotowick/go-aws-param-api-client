// Package utils provides a library of commond methods.
package utils

import (
	"fmt"
)

// IsSet returns true or false if a variable is not empty
//
// Example:
//     // Check for validity
//     svc := utils.IsSet(arg)
//
func IsSet(arg string) bool {
	if arg != "" {
		return true
	}
	return false
}

// EvalBool returns true or false if the argument is true or false. In addition,
// it outputs a message if the variable is not set.
//
// Example:
//     // Evaluate boolean value.
//     svc := utils.EvalBool(arg, name)
//
func EvalBool(arg bool, name string) bool {
	if !arg {
		fmt.Printf("%s is not set.\n", name)
		return false
	}
	return true
}
