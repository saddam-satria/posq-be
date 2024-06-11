package commons

import "golang.org/x/crypto/bcrypt"

func HashedBcrypt(textString string) (string, error) {
	hashedText, err := bcrypt.GenerateFromPassword([]byte(textString), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedText), nil
}

func VerifyHashed(plainText string, hashedText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(plainText))

	return err == nil
}