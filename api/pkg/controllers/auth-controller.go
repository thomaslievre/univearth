package controllers

import (
	"api/pkg/models"
	"api/pkg/utils"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

var NewUser models.User

var secret = os.Getenv("JWT_SECRET")

func Register(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	if _, err := utils.ParseBody(r, &data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := &models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	b := user.CreateUser()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

const (
	PW_SALT_BYTES = 32
	PW_HASH_BYTES = 64
)

func Login(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	if _, err := utils.ParseBody(r, &data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, _ := models.GetUserByEmail(data["email"])

	// check user id is valid
	if user.Id == uuid.Nil {
		res, _ := json.Marshal(map[string]interface{}{
			"message": "user not found",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
	}

	//salt := make([]byte, PW_SALT_BYTES)
	//_, err := io.ReadFull(rand.Reader, salt)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//hash, err = bcrypt.CompareHashAndPassword()

	// check password is valid
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		res, _ := json.Marshal(map[string]interface{}{
			"message": "incorrect password",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
	}

	expirationTime := time.Now().Add(time.Hour * 24)

	// generate json web token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":   user.Id,
		"ExpireAt": expirationTime.Unix(),
	})

	token, err := claims.SignedString([]byte(secret))

	if err != nil {
		res, _ := json.Marshal(map[string]interface{}{
			"message": "could not login",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})

	w.WriteHeader(http.StatusOK)
}

func Validate(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")

	if err != nil {
		res, _ := json.Marshal(map[string]interface{}{
			"message": "no token set in cookie",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		res, _ := json.Marshal(map[string]interface{}{
			"message": "not authorized",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(res)
	}

	claims := token.Claims.(*jwt.MapClaims)

	//issuer := (*claims)["Issuer"].(string)

	//uid := uuid.MustParse(issuer)
	//
	//user, _ := models.GetUserById(uid)

	res, _ := json.Marshal(claims)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
