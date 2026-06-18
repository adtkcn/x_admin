package util

import (
	"errors"
	"time"
	"x_admin/config"

	"github.com/golang-jwt/jwt/v5"
)

// UserClaims JWT Claims（用户系统，仅存必要字段）
type UserClaims struct {
	UserID       string `json:"userId"`
	TokenVersion int64  `json:"tokenVersion"`
	TokenType    string `json:"tokenType"` // access 或 refresh
	jwt.RegisteredClaims
}

var JWTUtil = &jwtUtil{}

type jwtUtil struct{}

// GenerateAccessToken 生成 access_token
func (j *jwtUtil) GenerateAccessToken(userID string, tokenVersion int64) (string, error) {
	now := time.Now()
	claims := UserClaims{
		UserID:       userID,
		TokenVersion: tokenVersion,
		TokenType:    "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(config.JWTConfig.AccessExpireSec) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    config.AppConfig.AppName,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTConfig.AccessSecret))
}

// GenerateRefreshToken 生成 refresh_token
func (j *jwtUtil) GenerateRefreshToken(userID string, tokenVersion int64) (string, error) {
	now := time.Now()
	claims := UserClaims{
		UserID:       userID,
		TokenVersion: tokenVersion,
		TokenType:    "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(config.JWTConfig.RefreshExpireSec) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    config.AppConfig.AppName,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTConfig.RefreshSecret))
}

// GenerateTokenPair 生成 access_token + refresh_token 对
func (j *jwtUtil) GenerateTokenPair(userID string, tokenVersion int64) (accessToken, refreshToken string, err error) {
	accessToken, err = j.GenerateAccessToken(userID, tokenVersion)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = j.GenerateRefreshToken(userID, tokenVersion)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

// ParseAccessToken 解析 access_token
func (j *jwtUtil) ParseAccessToken(tokenStr string) (*UserClaims, error) {
	return j.parseToken(tokenStr, config.JWTConfig.AccessSecret, "access")
}

// ParseRefreshToken 解析 refresh_token
func (j *jwtUtil) ParseRefreshToken(tokenStr string) (*UserClaims, error) {
	return j.parseToken(tokenStr, config.JWTConfig.RefreshSecret, "refresh")
}

// parseToken 内部解析
func (j *jwtUtil) parseToken(tokenStr, secret, expectedType string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims.TokenType != expectedType {
		return nil, errors.New("token type mismatch")
	}
	return claims, nil
}

// IsAccessExpiringSoon 判断 access_token 是否即将过期（剩余<30分钟）
func (j *jwtUtil) IsAccessExpiringSoon(claims *UserClaims) bool {
	if claims.ExpiresAt == nil {
		return true
	}
	remaining := time.Until(claims.ExpiresAt.Time)
	return remaining < 30*time.Minute
}
