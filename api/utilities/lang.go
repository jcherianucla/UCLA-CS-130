package utilities

// Represents an enum for languages
type Language int64

const (
	Cpp Language = 1 + iota
	C
	Java
)

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
	switch lang {
	case "C++":
		return Cpp
	case "C":
		return C
	case "Java":
		return Java
	default:
		return -1
	}
}
