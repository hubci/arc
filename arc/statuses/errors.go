package statuses

import "fmt"

/*
 * NameTakenError occurs when a status page is registered with a name that is
 * already registered.
 */
type NameTakenError struct {
	Name string
}

/*
 * Error returns the string representation of NameTakenError.
 */
func (e *NameTakenError) Error() string {
	return fmt.Sprintf("Failed to register the name %s. It has already been taken.", e.Name)
}

/*
 * newNameTakenError creates a new NameTakenError.
 */
func newNameTakenError(name string) *NameTakenError {
	return &NameTakenError{Name: name}
}
