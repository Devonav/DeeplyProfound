package web

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"go-trigger-app/db"
	"go-trigger-app/model"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key") // Please use a more secure key in a real application!

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

type App struct {
	d        db.DB
	handlers map[string]http.HandlerFunc
}

func NewApp(d db.DB, cors bool) App {
	app := App{
		d:        d,
		handlers: make(map[string]http.HandlerFunc),
	}

	// This is the updated section
	var techHandler http.HandlerFunc = authMiddleware(app.GetTechnologies)
	loginHandler := app.LoginUser
	registerHandler := app.RegisterUser

	if !cors {
		techHandler = disableCors(techHandler)
		loginHandler = disableCors(loginHandler)
		registerHandler = disableCors(registerHandler)
	}

	app.handlers["/api/technologies"] = techHandler
	app.handlers["/api/login"] = loginHandler
	app.handlers["/api/register"] = registerHandler
	app.handlers["/"] = http.FileServer(http.Dir("./webapp/build")).ServeHTTP
	return app
}

func (a *App) Serve() error {
	for path, handler := range a.handlers {
		http.Handle(path, handler)
	}
	log.Println("Web server is available on port 8080")
	return http.ListenAndServe(":8080", nil)
}

func (a *App) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		sendErr(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	userID, err := a.d.CreateUser(&user, string(hashedPassword))
	if err != nil {
		sendErr(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	// Add the default leaks for the new user
	err = a.d.AddDefaultLeaksForUser(userID)
	if err != nil {
		// In a real app, you might want more robust error handling here.
		log.Printf("Failed to add default leaks for user %d: %v", userID, err)
	}

	user.ID = userID
	user.Password = "" // Don't send password back

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (a *App) LoginUser(w http.ResponseWriter, r *http.Request) {
	var creds model.User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		sendErr(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, hashedPassword, err := a.d.GetUserByUsername(creds.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			sendErr(w, http.StatusUnauthorized, "User not found")
			return
		}
		sendErr(w, http.StatusInternalServerError, "Failed to get user")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password)); err != nil {
		sendErr(w, http.StatusUnauthorized, "Invalid password")
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, "Failed to create token")
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func (a *App) GetTechnologies(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		sendErr(w, http.StatusUnauthorized, "Invalid user ID in token")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	technologies, err := a.d.GetTechnologiesByUserID(userID)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(technologies)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			sendErr(w, http.StatusUnauthorized, "Missing Authorization header")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				sendErr(w, http.StatusUnauthorized, "Invalid token signature")
				return
			}
			sendErr(w, http.StatusBadRequest, "Invalid token")
			return
		}

		if !token.Valid {
			sendErr(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		h(w, r)
	}
}
