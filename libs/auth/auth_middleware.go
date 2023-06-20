package auth

// import (
// 	"strings"

// 	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/clients"
// 	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"

// 	"firebase.google.com/go/auth"
// 	"github.com/gofiber/fiber/v2"
// )

// const AuthUserKey = "authUser"

// type AuthUser struct {
// 	UID           string
// 	ID_           string
// 	Email         string
// 	EmailVerified bool
// 	Name          string
// 	Roles         []string
// }

// type AuthMiddlewareConfig struct {
// 	Firebase   clients.FirebaseClientIF
// 	IgnoreUrls []string
// }

// func NewAuthMiddleware(cfg AuthMiddlewareConfig) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		// Check ignored URLS
// 		url := c.Method() + "::" + c.Path()
// 		if cfg.IgnoreUrls != nil && len(cfg.IgnoreUrls) > 0 {
// 			for i := range cfg.IgnoreUrls {
// 				if cfg.IgnoreUrls[i] == url {
// 					return c.Next()
// 				}
// 			}
// 		}

// 		token := strings.Split(c.Get(fiber.HeaderAuthorization), " ")
// 		if len(token) != 2 || len(token) > 0 && token[0] != "Bearer" {
// 			err := models.ErrInvalidToken
// 			return c.Status(err.HttpStatusCode).JSON(err)
// 		}

// 		user, err := cfg.Firebase.Auth(c.Context(), token[1])
// 		if err != nil {
// 			err := models.ErrInvalidToken
// 			return c.Status(err.HttpStatusCode).JSON(err)
// 		}

// 		c.Locals(AuthUserKey, parseAuthUser(user))
// 		return c.Next()
// 	}
// }

// func GetAuthUser(c *fiber.Ctx) *AuthUser {
// 	rawUser := c.Locals(AuthUserKey)
// 	return rawUser.(*AuthUser)
// }

// func parseAuthUser(token *auth.Token) *AuthUser {
// 	role, _ := token.Claims["role"].(string)
// 	email, _ := token.Claims["email"].(string)
// 	emailVerified, _ := token.Claims["email_verified"].(bool)
// 	name, _ := token.Claims["name"].(string)
// 	// _id for mps
// 	_id, _ := token.Claims["_id"].(string)

// 	user := &AuthUser{
// 		UID:           extractTokenId(token),
// 		ID_:           _id,
// 		Email:         email,
// 		EmailVerified: emailVerified,
// 		Name:          name,
// 		Roles:         strings.Split(role, "|"),
// 	}

// 	return user
// }

// func extractTokenId(token *auth.Token) string {
// 	id, ok := token.Claims["userId"].(string)
// 	if ok && id != "" {
// 		return id
// 	}

// 	id, ok = token.Claims["user_id"].(string)
// 	if ok && id != "" {
// 		return id
// 	}

// 	return token.UID
// }
