package utilities

import (
	"strings"
)

// Represents an enum for languages
type Language int64

const (
	Cpp Language = 1 + iota
	C
	Java
)

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
