package main

import (
    "crypto/rand"
    "flag"
    "fmt"
    "math/big"
    "strings"
)

const (
    lowercase = "abcdefghijklmnopqrstuvwxyz"
    uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    digits    = "0123456789"
    symbols   = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

// PasswordGenerator holds configuration for password generation
type PasswordGenerator struct {
    length      int
    useLower    bool
    useUpper    bool
    useDigits   bool
    useSymbols  bool
    excludeAmbiguous bool
}

// NewPasswordGenerator creates a new password generator with default settings
func NewPasswordGenerator(length int) *PasswordGenerator {
    return &PasswordGenerator{
        length:      length,
        useLower:    true,
        useUpper:    true,
        useDigits:   true,
        useSymbols:   false,
        excludeAmbiguous: false,
    }
}

// SetOptions configures the password generator
func (pg *PasswordGenerator) SetOptions(useLower, useUpper, useDigits, useSymbols, excludeAmbiguous bool) {
    pg.useLower = useLower
    pg.useUpper = useUpper
    pg.useDigits = useDigits
    pg.useSymbols = useSymbols
    pg.excludeAmbiguous = excludeAmbiguous
}

// buildCharSet builds the character set based on enabled options
func (pg *PasswordGenerator) buildCharSet() string {
    var charset strings.Builder
    
    if pg.useLower {
        charset.WriteString(lowercase)
    }
    if pg.useUpper {
        charset.WriteString(uppercase)
    }
    if pg.useDigits {
        charset.WriteString(digits)
    }
    if pg.useSymbols {
        charset.WriteString(symbols)
    }
    
    charSetStr := charset.String()
    
    if pg.excludeAmbiguous {
        // Remove ambiguous characters: O, 0, I, l, 1, etc.
        charSetStr = strings.Map(func(r rune) rune {
            switch r {
            case 'O', 'o', '0', 'I', 'i', 'l', '1':
                return -1
            default:
                return r
            }
        }, charSetStr)
    }
    
    return charSetStr
}

// Generate generates a random password
func (pg *PasswordGenerator) Generate() (string, error) {
    charset := pg.buildCharSet()
    
    if charset == "" {
        return "", fmt.Errorf("no character set selected")
    }
    
    password := make([]byte, pg.length)
    charsetLen := big.NewInt(int64(len(charset)))
    
    for i := 0; i < pg.length; i++ {
        randomIndex, err := rand.Int(rand.Reader, charsetLen)
        if err != nil {
            return "", err
        }
        password[i] = charset[randomIndex.Int64()]
    }
    
    return string(password), nil
}

// GenerateMultiple generates multiple passwords
func (pg *PasswordGenerator) GenerateMultiple(count int) ([]string, error) {
    passwords := make([]string, count)
    for i := 0; i < count; i++ {
        pwd, err := pg.Generate()
        if err != nil {
            return nil, err
        }
        passwords[i] = pwd
    }
    return passwords, nil
}

// GeneratePronounceable generates a more memorable (pronounceable) password
func GeneratePronounceable(length int) (string, error) {
    consonants := "bcdfghjklmnpqrstvwxyz"
    vowels := "aeiou"
    
    password := make([]byte, length)
    
    for i := 0; i < length; i++ {
        var charset string
        if i%2 == 0 {
            charset = consonants
        } else {
            charset = vowels
        }
        
        randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
        if err != nil {
            return "", err
        }
        password[i] = charset[randomIndex.Int64()]
    }
    
    return string(password), nil
}

// GenerateSecure generates a high-security password (14+ chars, all character types)
func GenerateSecure() (string, error) {
    generator := NewPasswordGenerator(16)
    generator.SetOptions(true, true, true, true, false)
    return generator.Generate()
}

// Example usage with command-line flags
func main() {
    // Command-line flags
    length := flag.Int("length", 12, "Length of the password")
    count := flag.Int("count", 1, "Number of passwords to generate")
    useLower := flag.Bool("lower", true, "Include lowercase letters")
    useUpper := flag.Bool("upper", true, "Include uppercase letters")
    useDigits := flag.Bool("digits", true, "Include digits")
    useSymbols := flag.Bool("symbols", false, "Include symbols")
    excludeAmb := flag.Bool("exclude-ambiguous", false, "Exclude ambiguous characters (O,0,I,l,1)")
    pronounceable := flag.Bool("pronounceable", false, "Generate pronounceable password")
    secure := flag.Bool("secure", false, "Generate high-security password (overrides other options)")
    
    flag.Parse()
    
    var passwords []string
    var err error
    
    if *secure {
        // Generate high-security password
        pwd, err := GenerateSecure()
        if err != nil {
            fmt.Printf("Error generating password: %v\n", err)
            return
        }
        passwords = []string{pwd}
    } else if *pronounceable {
        // Generate pronounceable password
        pwd, err := GeneratePronounceable(*length)
        if err != nil {
            fmt.Printf("Error generating password: %v\n", err)
            return
        }
        passwords = []string{pwd}
    } else {
        // Use the standard generator with custom options
        generator := NewPasswordGenerator(*length)
        generator.SetOptions(*useLower, *useUpper, *useDigits, *useSymbols, *excludeAmb)
        
        passwords, err = generator.GenerateMultiple(*count)
        if err != nil {
            fmt.Printf("Error generating passwords: %v\n", err)
            return
        }
    }
    
    // Print results
    fmt.Printf("Generated %d password(s):\n", len(passwords))
    for i, pwd := range passwords {
        fmt.Printf("%d. %s\n", i+1, pwd)
    }
}