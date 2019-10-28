// Package options gets command line args for a program
package options

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// Parse function returns the option-value map for the cli args
func Parse(config []Option) (Output, error) {
	args := os.Args[1:]
	optionTypeMap := convertToMap(config)
	values := make(Output)

	i := 0

	for i < len(args) {
		// check if the string starts with --
		if !strings.HasPrefix(args[i], "--") {
			return nil, errors.New("Arguments need to start with '--'")
		}

		option := args[i][2:]
		optionType, optionExists := optionTypeMap[option]

		// check if option exists in the map
		if !optionExists {
			return nil, errors.New("invalid argument " + option)
		}

		switch optionType {
		case "Bool":
			values[option] = true
		case "Int":
			i++
			value, err := strconv.Atoi(args[i])
			if err != nil {
				return nil, errors.New("invalid integer passed for argument " + option)
			}
			values[option] = value
		case "String":
			i++
			values[option] = args[i]
		}

		i++
	}
	return values, nil
}
