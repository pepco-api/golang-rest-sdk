package pasargad
/**
 * Thanks to Github User @Aaron0 for go-rsa-sign repository that inspired us:
 *    @URL: https://github.com/AaronO/go-rsa-sign
 */
import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/hex"
)

type Signer struct {
	Key *rsa.PrivateKey
}

func NewSigner(pemKey []byte) (*Signer, error) {
	key, err := parsePrivateKey(pemKey)
	if err != nil {
		return nil, err
	}

	return &Signer{key}, nil
}

func (s *Signer) Sign(data []byte) ([]byte, error) {
	hash := crypto.SHA1
	h := hash.New()
	h.Write(data)
	hashed := h.Sum(nil)

	return rsa.SignPKCS1v15(rand.Reader, s.Key, hash, hashed)
}

func (s *Signer) SignHex(data []byte) (string, error) {
	sig, err := s.Sign(data)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sig), nil
}

func (s *Signer) SignBase64(data []byte) (string, error) {
	sig, err := s.Sign(data)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sig), nil
}
