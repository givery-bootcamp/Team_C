package hash

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(p string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(password), err
}

func CompareHashPassword(hashedPassword, requestPassword string) error {
	// パスワードの文字列をハッシュ化して、既に登録されているハッシュ化したパスワードと比較します
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(requestPassword)); err != nil {
		return err
	}
	return nil
}
