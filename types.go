// Package options gets command line args for a program
package options

// Option holds the structure of the config
type Option struct {
	Name        string
	Description string
	Type        string
}

// typeMap is the type defining the
// intermediate result of option-type mapping
type typeMap map[string]string

// Output represents the map returned for each option.
// The key is the options name and the value is it's value
type Output map[string]interface{}
