// Package options gets command line args for a program
package options

func convertToMap(options []Option) typeMap {
	optionTypeMap := make(typeMap)
	for _, option := range options {
		optionTypeMap[option.Name] = option.Type
	}
	return optionTypeMap
}
