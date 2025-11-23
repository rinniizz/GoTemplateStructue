package utils

import (
	"html"
	"regexp"
	"strings"
	"unicode"
)

// SanitizeString removes dangerous characters and patterns
func SanitizeString(input string) string {
	// Escape HTML to prevent XSS
	input = html.EscapeString(input)

	// Remove SQL injection patterns
	sqlPatterns := []string{
		`(?i)(union|select|insert|update|delete|drop|create|alter|exec|execute)`,
		`(?i)(--|;|\/\*|\*\/|xp_|sp_)`,
		`(?i)(script|javascript|onerror|onload|alert|prompt|confirm)`,
	}

	for _, pattern := range sqlPatterns {
		re := regexp.MustCompile(pattern)
		input = re.ReplaceAllString(input, "")
	}

	// Trim whitespace
	input = strings.TrimSpace(input)

	return input
}

// ValidateEmail checks if email format is valid
func ValidateEmail(email string) bool {
	if email == "" {
		return false
	}

	// RFC 5322 compliant email regex
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// ValidatePassword checks password strength
// Requirements: min 8 chars, 1 uppercase, 1 lowercase, 1 number, 1 special char
func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}

// GetPasswordStrength returns password strength level (1-5)
func GetPasswordStrength(password string) int {
	strength := 0

	if len(password) >= 8 {
		strength++
	}
	if len(password) >= 12 {
		strength++
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if hasUpper {
		strength++
	}
	if hasLower {
		strength++
	}
	if hasNumber {
		strength++
	}
	if hasSpecial {
		strength++
	}

	// Cap at 5
	if strength > 5 {
		strength = 5
	}

	return strength
}

// ValidatePhone checks phone number format (Thai format)
func ValidatePhone(phone string) bool {
	if phone == "" {
		return false
	}

	// Thai phone format: 0812345678 or +66812345678
	re := regexp.MustCompile(`^(\+66|0)[0-9]{9}$`)
	return re.MatchString(phone)
}

// ValidateURL checks if URL format is valid
func ValidateURL(url string) bool {
	if url == "" {
		return false
	}

	re := regexp.MustCompile(`^https?://[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*`)
	return re.MatchString(url)
}

// ContainsSQLInjection checks for SQL injection patterns
func ContainsSQLInjection(input string) bool {
	sqlPatterns := []string{
		`(?i)(union.*select)`,
		`(?i)(insert.*into)`,
		`(?i)(delete.*from)`,
		`(?i)(drop.*table)`,
		`(?i)(update.*set)`,
		`(?i)(exec|execute)`,
		`(?i)(script.*>)`,
		`(?i)(--|;|\/\*|\*\/)`,
	}

	for _, pattern := range sqlPatterns {
		re := regexp.MustCompile(pattern)
		if re.MatchString(input) {
			return true
		}
	}

	return false
}

// ContainsXSS checks for XSS patterns
func ContainsXSS(input string) bool {
	xssPatterns := []string{
		`(?i)(<script|</script>)`,
		`(?i)(javascript:)`,
		`(?i)(onerror|onload|onclick)`,
		`(?i)(<iframe|</iframe>)`,
		`(?i)(alert\(|prompt\(|confirm\()`,
	}

	for _, pattern := range xssPatterns {
		re := regexp.MustCompile(pattern)
		if re.MatchString(input) {
			return true
		}
	}

	return false
}
