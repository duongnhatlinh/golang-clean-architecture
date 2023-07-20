package tokenprovider

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtProvider struct {
	secret string
}

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	Payload TokenPayload `json:"payload"`
	jwt.StandardClaims
}

// Generate token
func (j *jwtProvider) Generate(data TokenPayload, expiry int) (*Token, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))

	if err != nil {
		return nil, err
	}

	return &Token{
		Token:   myToken,
		Expiry:  expiry,
		Created: time.Now().UTC(),
	}, nil
}

// Validate token
func (j *jwtProvider) Validate(token string) (*TokenPayload, error) {
	myToken, err := jwt.ParseWithClaims(token, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, ErrInvalidToken
	}

	if !myToken.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := myToken.Claims.(*myClaims)

	if !ok {
		return nil, ErrInvalidToken
	}

	return &claims.Payload, nil
}
