package utils

// Utils is use for public utilities method
var Utils *utils

type utils struct{}

// IsProduction func is check local env return true or false
func (u *utils) IsProduction() bool {
	// Add logic to check env that is dev or prod
	return false
}
