package utilities

import (
	"strings"
)

// Represents an enum for languages - default to C++
type Language int64

const (
	Cpp Language = 1 + iota
	C
	Java
)

// GetLanguageFromInt converts an int64 to Language, used
// primarily for type assertion through reflection.
// It takes in an int64 to convert.
// It returns the Language mapping for the int64.
func GetLanguageFromInt(v int64) Language {
	Sugar.Infof("Inside get lang int: %v", v)
	switch v {
	case 1:
		return Cpp
	case 2:
		return C
	case 3:
		return Java
	default:
		return Cpp
	}
}

// GetLanguage converts a Language to a corresponding
// string for return upstream.
// It takes in a language to convert.
// It returns the string version of the language.
func GetLanguage(lang Language) string {
	switch lang {
	case Cpp:
		return "C++"
	case C:
		return "C"
	case Java:
		return "Java"
	default:
		return "Unknown"
	}
}

// SetLanguage converts a string to its corresponding
// language.
// It takes in a string to convert.
// It returns the language.
func SetLanguage(lang string) Language {
	lang = strings.Title(lang)
	switch lang {
	case "C++":
		return Cpp
	case "C":
		return C
	case "Java":
		return Java
	default:
		return Cpp
	}
}
