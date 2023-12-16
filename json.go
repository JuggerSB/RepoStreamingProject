package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("tu_clave_secreta")

// User struct representa un usuario en la aplicación
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims struct representa las reclamaciones del token JWT
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	r := gin.Default()

	// Ruta para el inicio de sesión
	r.POST("/login", loginHandler)

	// Ruta protegida que requiere un token válido
	r.GET("/reproducir", authMiddleware, reproduccionHandler)

	// Inicia el servidor en el puerto 8080
	r.Run(":8080")
}

func loginHandler(c *gin.Context) {
	// Lógica para verificar las credenciales del usuario
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Credenciales inválidas"})
		return
	}

	// Verifica las credenciales (solo un ejemplo, implementa tu lógica de autenticación)
	if user.Username == "usuario" && user.Password == "contraseña" {
		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &Claims{
			Username: user.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
	}
}

func authMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticación no proporcionado"})
		c.Abort()
		return
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticación inválido"})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticación no válido"})
		c.Abort()
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticación no válido"})
		c.Abort()
		return
	}

	// Puedes acceder a la información del usuario a través de claims.Username
	c.Set("username", claims.Username)

	c.Next()
}

func reproduccionHandler(c *gin.Context) {
	// Lógica de reproducción que requiere autenticación
	// Puedes acceder al nombre de usuario del usuario autenticado usando c.GetString("username")
	c.JSON(http.StatusOK, gin.H{"message": "Reproduciendo contenido multimedia"})
}
