package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type EnvData struct {
    DBURL   string
    JWTKey  string
    PORT  string
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func LoadEnv() *EnvData {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    cfg := &EnvData{
        DBURL: os.Getenv("DATABASE_URL"),
        JWTKey: os.Getenv("JWT_KEY"),
        PORT: os.Getenv("PORT"),
    }

    if cfg.DBURL == "" {
        log.Fatal("DATABASE_URL is not set")
    }
	if cfg.PORT == "" {
        log.Fatal("PORT is not set")
    }
    if cfg.JWTKey == "" {
        log.Fatal("JWT_KEY is not set")
    }

    return cfg
}

func SignJWTToken(userId int64, email string) (string, error) {
	var (
		key []byte
		t   *jwt.Token
	)

	env := LoadEnv()
	key = []byte(env.JWTKey)
	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":   "expense-tracker",
			"sub":   userId,
			"email": email,
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		})
	s, err := t.SignedString(key)

	return s, err
}

func VerifyJWTToken(tokenString string) (*jwt.Token, error) {
	env := LoadEnv()
	key := []byte(env.JWTKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorExpired)
	}
	
	// Check if the token has expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && claims["exp"] != nil {
		if exp, ok := claims["exp"].(float64); ok && time.Unix(int64(exp), 0).Before(time.Now()) {
			return nil, jwt.NewValidationError("token has expired", jwt.ValidationErrorExpired)
		}
	}
	return token, nil
}

func GetJWTTokenFromHeader(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return "", jwt.NewValidationError("authorization header is missing", jwt.ValidationErrorMalformed)
	}
	if len(token) < 7 || token[:7] != "Bearer " {
		return "", jwt.NewValidationError("invalid authorization header format", jwt.ValidationErrorMalformed)
	}
	token = token[7:] // Remove "Bearer " prefix
	if token == "" {
		return "", jwt.NewValidationError("token is empty", jwt.ValidationErrorMalformed)
	}
	return token, nil
}

func GetUserIdFromJWTToken(r *http.Request) (int64, error) {
	tokenString, err := GetJWTTokenFromHeader(r)
	if err != nil {
		return 0, err
	}

	token, err := VerifyJWTToken(tokenString)
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && claims["sub"] != nil {
		if userId, ok := claims["sub"].(float64); ok {
			return int64(userId), nil
		}
	}
	return 0, jwt.NewValidationError("user ID not found in token", jwt.ValidationErrorClaimsInvalid)
}