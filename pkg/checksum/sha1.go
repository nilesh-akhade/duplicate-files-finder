package checksum

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
)

type sha1ChecksumCalc struct {
}

func NewSHA1Checksum() ChecksumCalc {
	return &sha1ChecksumCalc{}
}

func (m *sha1ChecksumCalc) Calculate(file string) (string, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	hash := sha1.New()
	if _, err := hash.Write(bytes); err != nil {
		return "", fmt.Errorf("error calculating checksum: %w", err)
	}
	hashSum := hash.Sum(nil)
	return fmt.Sprintf("%x", hashSum), nil
}
