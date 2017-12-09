package utilities

const (
	// Default secret for JWT
	DEFAULT_TOKEN_SECRET = "47X86A1Fnr6TqkpEyK0X"
)

// The port that the backend server runs on
var PORT = GetVar("PORT", "8080")
var DEFAULT_PASSWORD = GetVar("STUDENT_PASSWORD", "password")
