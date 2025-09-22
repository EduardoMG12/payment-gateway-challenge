package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"time"
)

func GenerateCardDetails() (cardNumber, cvc, expiryMonth, expiryYear string, err error) {
	maxCardNumber := big.NewInt(999999999999)
	randCardNumber, err := rand.Int(rand.Reader, maxCardNumber)
	if err != nil {
		return "", "", "", "", fmt.Errorf("failed to generate card number: %w", err)
	}
	cardNumber = fmt.Sprintf("%012d", randCardNumber)

	maxCvc := big.NewInt(999)
	randCvc, err := rand.Int(rand.Reader, maxCvc)
	if err != nil {
		return "", "", "", "", fmt.Errorf("failed to generate CVC: %w", err)
	}
	cvc = fmt.Sprintf("%03d", randCvc)

	now := time.Now()
	maxYears := big.NewInt(5)
	randYears, err := rand.Int(rand.Reader, maxYears)
	if err != nil {
		return "", "", "", "", fmt.Errorf("failed to generate expiry date: %w", err)
	}

	expiry := now.AddDate(int(randYears.Int64())+1, 0, 0)
	expiryMonth = expiry.Format("01")
	expiryYear = expiry.Format("06")

	return cardNumber, cvc, expiryMonth, expiryYear, nil
}

func GenerateCardToken(cardNumber, cvc, expiryMonth, expiryYear string) string {
	dataToHash := cardNumber + cvc + expiryMonth + expiryYear
	hash := sha256.Sum256([]byte(dataToHash))

	return fmt.Sprintf("%x", hash)
}
