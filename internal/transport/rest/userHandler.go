package rest

import (
	"backendAleshinaFeklistova/internal/models"
	"backendAleshinaFeklistova/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

// RegisterUser
//
//	@Summary	New User registration
//	@ID			registerUser
//	@Tags		User
//	@Accept		json
//	@Param		UserData	body		models.User	true	"UserData"
//	@Success	200			{string}	string		"ok"
//	@Router		/user/register [post]
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newUser.Password, err = service.HashPassword(newUser.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	service.RegisterUser(newUser)
}

// Login
//
//	@Summary	User login
//	@ID			login
//	@Tags		User
//	@Accept		mpfd
//	@Param		Username	formData	string	true	"Username"
//	@Param		Password	formData	string	true	"Password"
//	@Success	200			{string}	string	"ok"
//	@Router		/user/login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var (
		email    string
		password string
	)
	reader, _ := r.MultipartReader()
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		formName := part.FormName()
		formValue, err := io.ReadAll(part)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		switch formName {
		case "Username":
			email = string(formValue)
		case "Password":
			password = string(formValue)
		}
	}

	if service.ExistsByEmail(email) && service.CheckPasswordHash(password, email) {
		w.Write([]byte(service.CreateJwt(email)))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetUserById
//
//	@Summary	Get user profile by id
//	@ID			getUserById
//	@Tags		User
//	@Param		id	path		int		true	"User ID"
//	@Security		ApiKeyAuth
//	@Success	200	{object}	models.User	"ok"
//	@Router		/user/{id} [get]
func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	res, err := json.Marshal(service.GetUserById(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// EditProfile
//
//	@Summary	Edit Profile
//	@ID			editProfile
//	@Tags		User
//	@Param		UserData	body		models.User	true	"UserData"
//	@Security		ApiKeyAuth
//	@Success	200			{string}	string		"ok"
//	@Router		/user/edit [put]
func EditProfile(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&user)
	service.EditProfile(user)
}

// DeleteProfile
//
//	@Summary	Delete profile
//	@ID			deleteProfile
//	@Tags		User
//	@Param		id	path		int		true	"User ID"
//	@Security		ApiKeyAuth
//	@Success	200	{string}	string	"ok"
//	@Router		/user/{id} [delete]
func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	service.DeleteProfile(id)
}

// GetAll
//
//	@Summary	Get all users
//	@ID			getAllUsers
//	@Tags		User
//	@Produce	json
//	@Security		ApiKeyAuth
//	@Success	200	{object}	[]models.User	"ok"
//	@Router		/user [get]
func GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(service.GetAllUsers())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
