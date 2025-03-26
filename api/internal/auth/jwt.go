package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/configuration"
	"slices"
	"strconv"
	"time"
)

type JwtService struct {
	key []byte
}

func NewJwtService() *JwtService {
	return &JwtService{
		key: []byte("this is some key that should be changed"),
	}
}

type JwtClaims struct {
	Subject        int64
	Issuer         string
	Audience       []string
	ExpirationTime *jwt.NumericDate
	IssuedAt       *jwt.NumericDate
	NotBefore      *jwt.NumericDate
}

func NewJwtClaims(userID int64) *JwtClaims {
	now := time.Now()
	return &JwtClaims{
		Subject:        userID,
		Issuer:         configuration.JwtIssuer,
		Audience:       []string{configuration.JwtAudience},
		ExpirationTime: jwt.NewNumericDate(now.Add(1 * time.Hour)),
		IssuedAt:       jwt.NewNumericDate(now),
		NotBefore:      jwt.NewNumericDate(now.Add(-5 * time.Minute)),
	}
}

func (jc *JwtClaims) ToClaimsMap() jwt.MapClaims {
	return jwt.MapClaims{
		"sub": strconv.FormatInt(jc.Subject, 10),
		"iss": jc.Issuer,
		"aud": jc.Audience,
		"exp": jc.ExpirationTime,
		"nbf": jc.NotBefore,
		"iat": jc.IssuedAt,
	}
}

func (jc *JwtClaims) validateIssuer() bool {
	return jc.Issuer == configuration.JwtIssuer
}
func (jc *JwtClaims) validateAudience() bool {
	return slices.Contains(jc.Audience, configuration.JwtAudience)
}
func (jc *JwtClaims) validateExpirationTime() bool {
	return jc.ExpirationTime.Time.After(time.Now())
}
func (jc *JwtClaims) validateNotBefore() bool {
	return jc.NotBefore.Time.Before(time.Now())
}
func (jc *JwtClaims) Validate() bool {
	return jc.validateIssuer() &&
		jc.validateAudience() &&
		jc.validateExpirationTime() &&
		jc.validateNotBefore()
}

func stripNil(a []error) []error {
	b := make([]error, 0)
	for _, item := range a {
		if item != nil {
			b = append(b, item)
		}
	}
	return b
}

func ParseClaims(claims jwt.Claims) (*JwtClaims, error) {
	errs := make([]error, 0)
	jwtClaims := &JwtClaims{}
	subject, err := claims.GetSubject()
	errs = append(errs, err)
	issuer, err := claims.GetIssuer()
	errs = append(errs, err)
	audience, err := claims.GetAudience()
	errs = append(errs, err)
	expirationTime, err := claims.GetExpirationTime()
	errs = append(errs, err)
	issuedAt, err := claims.GetIssuedAt()
	errs = append(errs, err)
	notBefore, err := claims.GetNotBefore()
	errs = append(errs, err)
	id, err := strconv.ParseInt(subject, 10, 64)
	errs = append(errs, err)

	errs = stripNil(errs)

	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Println(e)
		}
		return nil, errors.New("jwt validation failed")
	}

	jwtClaims.Subject = id
	jwtClaims.Issuer = issuer
	jwtClaims.Audience = audience
	jwtClaims.ExpirationTime = expirationTime
	jwtClaims.NotBefore = notBefore
	jwtClaims.IssuedAt = issuedAt

	return jwtClaims, nil
}

func (js *JwtService) CreateToken(userID int64) (string, error) {
	jwtClaims := NewJwtClaims(userID)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		// https://datatracker.ietf.org/doc/html/rfc7519#section-4.1
		jwtClaims.ToClaimsMap(),
	)
	s, err := t.SignedString(js.key)
	return s, err
}

func (js *JwtService) ValidateToken(tokenString string) (*JwtClaims, error) {
	extractedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return js.key, nil
	})
	if err != nil {
		return nil, errors.New("failed to validate token")
	}
	return ParseClaims(extractedToken.Claims)
}

func UserIDFromCtx(ctx context.Context) (int64, error) {
	claims, ok := ctx.Value("user").(*JwtClaims)
	fmt.Println("claims: ", claims)
	if !ok {
		return 0, errors.New("no userID in context")
	}
	return claims.Subject, nil
}
