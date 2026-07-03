package twofer

import "fmt"

// ShareWith returns the two-fer phrase for the given name, defaulting to "you".
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
