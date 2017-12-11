package utilities

const (
	// Default secret for JWT
	DEFAULT_TOKEN_SECRET = "47X86A1Fnr6TqkpEyK0X"
	// Default storage size for Multipart forms
	MAX_STORAGE = 32 << 20
	// Time format for deadline
	TIME_FORMAT = "01-02-06 15:04 (MST)"
)

// The port that the backend server runs on
var PORT = GetVar("PORT", "8080")
var DEFAULT_PASSWORD = GetVar("STUDENT_PASSWORD", "password")
