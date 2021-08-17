package pasargad
/**
 * Thanks to phemmer for the gist:
 * https://gist.github.com/phemmer/fea012d087ff65819645
 */
import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/big"
)

type XMLRSAKey struct {
	Modulus  string
	Exponent string
	P        string
	Q        string
	DP       string
	DQ       string
	InverseQ string
	D        string
}

func (m *PasargadPaymentAPI)b64d(str string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println(err)
	}
	return []byte(decoded)
}

func (m *PasargadPaymentAPI)b64bigint(str string) *big.Int {
	bint := &big.Int{}
	bint.SetBytes(m.b64d(str))
	return bint
}

func (m *PasargadPaymentAPI) convertXmlToKey() (block []byte, err error) {
	xmlbs, err := ioutil.ReadFile(m.certificationFile)
	if err != nil {
		fmt.Println(err)
	}

	if decoded, err := base64.StdEncoding.DecodeString(string(xmlbs)); err == nil {
		xmlbs = decoded
	}

	xrk := XMLRSAKey{}
	error := xml.Unmarshal(xmlbs, &xrk)
	if error != nil {
		fmt.Println(error)
	}
	key := &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: m.b64bigint(xrk.Modulus),
			E: int(m.b64bigint(xrk.Exponent).Int64()),
		},
		D:      m.b64bigint(xrk.D),
		Primes: []*big.Int{m.b64bigint(xrk.P), m.b64bigint(xrk.Q)},
		Precomputed: rsa.PrecomputedValues{
			Dp:        m.b64bigint(xrk.DP),
			Dq:        m.b64bigint(xrk.DQ),
			Qinv:      m.b64bigint(xrk.InverseQ),
			CRTValues: ([]rsa.CRTValue)(nil),
		},
	}

	pemkey := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}
	block = pem.EncodeToMemory(pemkey)
	return block, nil
}
