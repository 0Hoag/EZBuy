package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"indetity-service/modules"
	"indetity-service/modules/item/entity"
	"net/http"
	"strings"
	"time"
)

type InvalidatedToken struct {
	ID         string `gorm:"primaryKey"`
	ExpiryTime time.Time
}

type AuthenticationRequest struct {
	Username string
	Password string
}

type AuthenticationResponse struct {
	Token         string
	Authenticated bool
}

type LogoutRequest struct {
	Token string
}

type RefreshRequest struct {
	Token string
}

type AuthService struct {
	tokenService *TokenService
	db           *gorm.DB
}

type TokenService struct {
	signerKey       []byte
	validDuration   time.Duration
	refreshDuration time.Duration
}

type TokenUser struct {
	UserID string
	Scope  []string
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AuthenticationRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		authService := NewAuthService(
			"RREniFvR4thHKZvOOSbmIgCCi59YC/qcJwNiAbl/50w1lldEjpnBXXpd5LGkYofz",
			24*time.Hour,
			7*24*time.Hour,
			db,
		)

		response, err := authService.Authenticate(req)
		if err != nil {
			switch err {
			case entity.UserNotFound:
				c.JSON(http.StatusNotFound, err)
			case entity.UnAuthentication:
				c.JSON(http.StatusUnauthorized, err)
			default:
				c.JSON(http.StatusInternalServerError, err)
			}
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func Logout(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LogoutRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		authService := NewAuthService(
			"RREniFvR4thHKZvOOSbmIgCCi59YC/qcJwNiAbl/50w1lldEjpnBXXpd5LGkYofz",
			24*time.Hour,
			7*24*time.Hour,
			db,
		)

		if err := authService.Logout(req); err != nil {
			switch err {
			case entity.UserNotFound:
				c.JSON(http.StatusNotFound, err)
			case entity.UnAuthentication:
				c.JSON(http.StatusUnauthorized, err)
			default:
				c.JSON(http.StatusInternalServerError, err)
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully logged out",
		})
	}
}

func RefreshToken(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RefreshRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		authService := NewAuthService(
			"RREniFvR4thHKZvOOSbmIgCCi59YC/qcJwNiAbl/50w1lldEjpnBXXpd5LGkYofz",
			24*time.Hour,
			7*24*time.Hour,
			db,
		)

		response, err := authService.RefreshToken(req)
		if err != nil {
			switch err {
			case entity.UserNotFound:
				c.JSON(http.StatusNotFound, err)
			case entity.UnAuthentication:
				c.JSON(http.StatusUnauthorized, err)
			default:
				c.JSON(http.StatusInternalServerError, err)
			}
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

//// Middleware function that takes db as parameter
//func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		authHeader := c.GetHeader("Authorization")
//		if authHeader == "" {
//			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
//				"error": "Authorization header is required",
//			})
//			return
//		}
//
//		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
//			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
//				"error": "Invalid authorization format",
//			})
//			return
//		}
//
//		token := authHeader[7:]
//
//		authService := NewAuthService(
//			"your-secret-key",
//			24*time.Hour,
//			7*24*time.Hour,
//			db,
//		)
//
//		// Verify token
//		validToken, err := authService.tokenService.ValidateToken(token)
//		if err != nil {
//			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
//				"error": "Invalid token",
//			})
//			return
//		}
//
//		claims, ok := validToken.Claims.(jwt.MapClaims)
//		if !ok {
//			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
//				"error": "Invalid token claims",
//			})
//			return
//		}
//
//		c.Set("userID", claims["sub"])
//		c.Set("scope", claims["scope"])
//
//		c.Next()
//	}
//}

func NewAuthService(signerKey string, validDuration time.Duration, refreshDuration time.Duration, db *gorm.DB) *AuthService {
	return &AuthService{
		tokenService: NewTokenService(signerKey, validDuration, refreshDuration),
		db:           db,
	}
}

func (s *AuthService) Authenticate(req AuthenticationRequest) (*AuthenticationResponse, error) {
	var user entity.User
	if err := s.db.Table(entity.User{}.TableName()).Preload("Role.Permission").Where("username = ?", req.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entity.UserNotFound
		}
		return nil, err
	}

	token, err := s.tokenService.GenerateToken(TokenUser{
		UserID: user.Username,
		Scope:  s.buildScope(user),
	})
	if err != nil {
		return nil, err
	}

	return &AuthenticationResponse{
		Token:         token,
		Authenticated: true,
	}, nil
}

func (s *AuthService) Logout(req LogoutRequest) error {
	token, err := s.verifyToken(req.Token, true)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return entity.UnAuthentication
	}

	invalidToken := InvalidatedToken{
		ID:         claims["jti"].(string),
		ExpiryTime: time.Unix(int64(claims["exp"].(float64)), 0),
	}

	return s.db.Create(&invalidToken).Error
}

func (s *AuthService) RefreshToken(req RefreshRequest) (*AuthenticationResponse, error) {
	token, _ := s.verifyToken(req.Token, true)

	claims, _ := token.Claims.(jwt.MapClaims)

	invalidToken := InvalidatedToken{
		ID:         claims["jti"].(string),
		ExpiryTime: time.Unix(int64(claims["exp"].(float64)), 0),
	}
	if err := s.db.Create(&invalidToken).Error; err != nil {
		return nil, err
	}

	username := claims["sub"].(string)
	var user entity.User
	if err := s.db.Preload("Role.Permission").Where("username = ?", username).First(&user).Error; err != nil {
		return nil, entity.UnAuthentication
	}

	newToken, err := s.tokenService.GenerateToken(TokenUser{
		UserID: user.Username,
		Scope:  s.buildScope(user),
	})
	if err != nil {
		return nil, err
	}

	return &AuthenticationResponse{
		Token:         newToken,
		Authenticated: true,
	}, nil
}

func (s *AuthService) verifyToken(tokenString string, isRefresh bool) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return s.tokenService.signerKey, nil
	})

	if err != nil || !token.Valid {
		return nil, entity.UnAuthentication
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, entity.UnAuthentication
	}

	var expiryTime time.Time
	if isRefresh {
		issueTime := time.Unix(int64(claims["iat"].(float64)), 0)
		expiryTime = issueTime.Add(s.tokenService.refreshDuration)
	} else {
		expiryTime = time.Unix(int64(claims["exp"].(float64)), 0)
	}

	if time.Now().After(expiryTime) {
		return nil, entity.UnAuthentication
	}

	var invalidToken InvalidatedToken
	if err := s.db.First(&invalidToken, "id = ?", claims["jti"]).Error; err == nil {
		return nil, entity.UnAuthentication
	}

	return token, nil
}

func (s *AuthService) buildScope(user entity.User) []string {
	var scopes []string
	for _, role := range user.Role {
		scopes = append(scopes, "ROLE_"+role.Name)
		for _, perm := range role.Permission {
			scopes = append(scopes, perm.Name)
		}
	}
	return scopes
}

func NewTokenService(signerKey string, validDuration, refreshDuration time.Duration) *TokenService {
	return &TokenService{
		signerKey:       []byte(signerKey),
		validDuration:   validDuration,
		refreshDuration: refreshDuration,
	}
}

func (s *TokenService) GenerateToken(user TokenUser) (string, error) {
	now := time.Now()

	claims := jwt.MapClaims{
		"sub":   user.UserID,
		"iss":   "devteria.com",
		"iat":   now.Unix(),
		"exp":   now.Add(s.validDuration).Unix(),
		"jti":   uuid.New().String(),
		"scope": strings.Join(user.Scope, " "),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString(s.signerKey)
}
