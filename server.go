package main

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/didbot/didbot-api-go/models"
)



func init() {
	models.LoadDB()
}

func main() {
	publicKeyStr := `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA0RmJg/RrFsgGZ/Ln9AfR
y9pH5mNSWHoHJZsm2wtCKfcKUkdugnMm0cxhaBq1dMVeEd3oQaOHiFde3kZy8nKe
X/jWOpPz5z03CaWO1WmK29ZB/S+8bzl++ui2RZSk00C3mQ5COvic1M6+1M+fBpeS
QZXtqeU328Np5SO3csdqBbszbeI6iZIRRMtw6AHdLx9a6A5PTVbTxn3QL2MYVatV
j/ZBrWU14Gpz9aBqy1XQDkuoRApFM4aTSBoFugpVo0Vyn6PUedHcW5tQR1dbLlEz
tPRzU9FKwN8zWEk24qhgdiggSrj6JWQtsy6bPT3P/PZ+zgzJdzChM+ic54fPGG5N
ir2NhaFkRjmj12JWUr3mi1cZjYw56QaB/JzROH4raefayEDfZUhX1P6j8s4V2R8l
+IQsGAr/D6g5NwjT9dTn8MhqlzzLskzv1HemtrejSpBVjrvMfhhMVK0bZdGezDGb
YzVPbzyCVQJh7Dzs4S7KdpSyu9excPldBb/+/7FJyA0ljc5GqCVt2uAhRn8dlx18
vVnKNFtyKk2xL/qHPqsDhyBS9UuwYTA5qO4wqCRSXV3ZWRcWZY0RTBlyEHm0DVpY
qTM0XMyI/IVOSBO8FamDXgPXjagVnPTviFtqLyvONrMuHfUN7xMS7ovPXioKEOeq
FXfPjdvNq0HyVlEvkoiExZ0CAwEAAQ==
-----END PUBLIC KEY-----`

	publicKey, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyStr))
	jwtConfig := middleware.JWTConfig{
		Skipper: func(c echo.Context) bool {
			return false
		},
		ContextKey:    "parsedToken",
		SigningKey:    publicKey,
		SigningMethod: "RS256",
	}

	// Initialize server and Set middleware
	e := echo.New()
	e.Pre(middleware.HTTPSRedirect())
	e.Use(
		middleware.RequestID(),
		middleware.Logger(),
		middleware.JWTWithConfig(jwtConfig),
		middleware.BodyLimit("1M"),
		middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}),
		SetToken)

	LoadRoutes(e)
	e.Logger.Fatal(e.StartTLS(":1334", "cert.pem", "key.pem"))

}

// SetToken middleware sets the Token struct into the request context after verifying it is unrevoked
func SetToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Parse out the token that was added to Context from the JWTWithConfig
		parsedToken 	:= c.Get("parsedToken").(*jwt.Token)
		claims 		:= parsedToken.Claims.(jwt.MapClaims)
		tokenId 	:= claims["jti"].(string)
		userId 		:= claims["sub"].(string)

		// TODO: check expiration date from claims["exp"].(string) before hitting db
		t := &models.Token{}
		token, err := t.ValidateToken(tokenId, userId)
		if err != nil {
			// TODO: log the request ID with the failed to authenticate user info
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		// TODO: log the request ID with the authenticated user info
		c.Set("token", token)
		return next(c)
	}
}