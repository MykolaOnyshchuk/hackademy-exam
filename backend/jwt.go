package main
import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"github.com/openware/rango/pkg/auth"
	"net/http"
)
type JWTService struct {
	keys *auth.KeyStore
}

func NewJWTService(privKeyPath, pubKeyPath string) (*JWTService, error) {
	keys, err := auth.LoadOrGenerateKeys(privKeyPath, pubKeyPath)
	if err != nil {
		return nil, err
	}
	return &JWTService{keys: keys}, nil
}

func (j *JWTService) GenearateJWT(u User) (string, error) {
	return auth.ForgeToken("empty", u.email, "empty", 0, j.keys.
		PrivateKey, nil)
}
func (j *JWTService) ParseJWT(jwt string) (auth.Auth, error) {
	return auth.ParseAndValidate(jwt, j.keys.PublicKey)
}

type JWTParams struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
func (u *InMemoryStorage) JWT(
	w http.ResponseWriter,
	r *http.Request,
	jwtService *JWTService,
) {
	params := &JWTParams{}
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		handleError(errors.New("could not read params"), w)
		return
	}
	passwordDigest := md5.New().Sum([]byte(params.Password))
	user, err := u.Get(params.Email)
	if err != nil {
		handleError(err, w)
		return
	}
	if string(passwordDigest) != user.password {
		handleError(errors.New("invalid login params"), w)
		return
	}
	token, err := jwtService.GenearateJWT(user)
	if err != nil {
		handleError(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}
