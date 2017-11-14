package utilities

// CheckError panics on an error, ending the program.
// It takes in the error to panic on if at all.
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
