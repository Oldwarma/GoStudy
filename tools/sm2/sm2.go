package signature

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/pem"
	"errors"
	"github.com/tjfoc/gmsm/pkcs12"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"math/big"
)

type Sm2Sign struct {
	publicKey  *sm2.PublicKey
	privateKey *sm2.PrivateKey
}

func NewSm2Sign(publicKeyPath, pw, privateKeyPath string) (*Sm2Sign, error) {
	privateKey, err := LoadSmp12(privateKeyPath, pw)
	if err != nil {
		return nil, err
	}
	publicKey, err := LoadSmCer(publicKeyPath)
	if err != nil {
		return nil, err
	}
	return &Sm2Sign{
		publicKey:  publicKey,
		privateKey: privateKey,
	}, nil
}

func (ss *Sm2Sign) PrivateKey() *sm2.PrivateKey {

	return ss.privateKey
}

func (ss *Sm2Sign) PublicKey() *sm2.PublicKey {

	return ss.publicKey
}

func (ss *Sm2Sign) SignSm2(msg []byte) ([]byte, error) {
	logx.Info(msg, "SignSm2==")
	r, s, err := sm2.Sm2Sign(ss.privateKey, msg, nil, rand.Reader)
	if err != nil {
		logx.Info(msg, "SignSm2"+err.Error())
		return nil, err
	}

	sign := append(r.Bytes(), s.Bytes()...)

	return sign, nil
}

func (ss *Sm2Sign) VerifySm2(msg, sign []byte) (bool, error) {
	pr := big.Int{}
	ps := big.Int{}
	sigLen := len(sign)
	pr.SetBytes(sign[:(sigLen / 2)])
	ps.SetBytes(sign[(sigLen / 2):])

	b := sm2.Sm2Verify(ss.publicKey, msg, nil, &pr, &ps)

	return b, nil
}

func (ss *Sm2Sign) Encrypt(data []byte, model int) ([]byte, error) {

	return sm2.Encrypt(ss.publicKey, data, rand.Reader, model)
}

func (ss *Sm2Sign) Decrypt(data []byte, model int) ([]byte, error) {

	return sm2.Decrypt(ss.privateKey, data, model)
}

// LoadSmp12 加载sm .P12 的私钥证书文件
func LoadSmp12(filePath, password string) (*sm2.PrivateKey, error) {
	//根据密码读取P12证书
	_, priv, err := pkcs12.SM2P12Decrypt(filePath, password)
	if err != nil {
		return nil, err
	}
	return priv, nil
}

// LoadSmCer 加载sm .cer的公钥证书文件
func LoadSmCer(filepath string) (*sm2.PublicKey, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	derBlock, _ := pem.Decode(file)
	Certificate, err := x509.ParseCertificate(derBlock.Bytes)
	if err != nil {
		return nil, err
	}
	pubKey, ok := Certificate.PublicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("certificate conversion failed")
	}
	sm2PubKey := &sm2.PublicKey{Curve: pubKey.Curve, X: pubKey.X, Y: pubKey.Y}

	return sm2PubKey, nil
}
