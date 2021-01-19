package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/garcialuis/Gossip/api/auth"
	"github.com/garcialuis/Gossip/api/models"
	"github.com/garcialuis/Gossip/api/responses"
	"github.com/garcialuis/Gossip/api/utils/formaterror"
	"github.com/gorilla/mux"
)

// CreateUser handler:
// swagger:route POST /users users createUser
// Creates a new user record
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//		201: userResponse
//		422: description: Unprocessable entity - unable to process input data
//		500: description: Internal Server Error
func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveUser(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)

}

// GetUsers handler:
// swagger:route GET /users users getUsers
// GetUsers returns a list of all users
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//		200: usersResponse
//		500: description: Internal Server Error
func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	users, err := user.FindAllUsers(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// GetUser Handler:
// swagger:route GET /users/{id} users getUser
// GetUser returns a user record with the specified userID
//
//	Responses:
//		400: description: Bad Request
//		200: userResponse
func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user := models.User{}
	userGotten, err := user.FindUserByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, userGotten)
}

// AuthenticatedGetUser is an Authenticated version of GetUser,
// this endpoint may return confidential data that GetUser will not
// swagger:route GET /private/users/{id} private_user GetAuthUser
// Returns authenticated user's information with additional sensitive info
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Security:
//	- api_key:
//
//	Responses:
//		401: description: Unauthorized
//		400: description: Bad Request
//		200: authenticatedUser
func (server *Server) AuthenticatedGetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Check Authentication via jwt token:
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	user := models.User{}
	userGotten, err := user.FindUserByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, userGotten)
}

// UpdateUser Handler:
// swagger:route PUT /users users UpdateUser
//	Security:
//	- api_key:
//
//	Responses:
//		422: description: Unprocessable Entity
//		401: description: Unauthorized
//		500: description: Internal Server Error
//		200: description: userResponse
func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	user.Prepare()
	err = user.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedUser, err := user.UpdateUserAccount(server.DB, uint32(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, updatedUser)

}

// UpdateUserAttributes Handler:
// swagger:route PUT /users users UpdateUserAttributes
//	Security:
//	- api_key:
//
//	Responses:
//		422: description: Unprocessable Entity
//		401: description: Unauthorized
//		500: description: Internal Server Error
//		200: description: userResponse
func (server *Server) UpdateUserAttributes(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	updatedUser, err := user.UpdateUserAttributes(server.DB, uint32(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, updatedUser)

}

// DeleteUser Handler:
// swagger:route DELETE /users users DeleteUser
//	Security:
//	- api_key:
//
//	Responses:
//		204: description: No Content
//		422: description: Unprocessable Entity
//		401: description: Unauthorized
func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	user := models.User{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if tokenID != 0 && tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	_, err = user.DeleteUser(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) GetUserBMI(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	varUID, err := strconv.ParseUint(vars["id"], 10, 32)
	uid := uint32(varUID)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user := models.User{}
	weight, height, err := user.GetWeightAndHeight(server.DB, uid)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	personInfo := server.NutriportClient.CalculateImperialBMI(weight, height)

	responses.JSON(w, http.StatusOK, personInfo)

}

func (server *Server) GetUserTEE(w http.ResponseWriter, r *http.Request) {

	activityLevel := "moderately active"

	vars := mux.Vars(r)
	varUID, err := strconv.ParseUint(vars["id"], 10, 32)
	uid := uint32(varUID)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user := models.User{}
	foundUser, err := user.FindUserByID(server.DB, uid)
	personInfo := server.NutriportClient.CalculateTotalEnergyExpenditure(int(foundUser.Age), int(foundUser.Gender), foundUser.Weight, activityLevel)

	responses.JSON(w, http.StatusOK, personInfo)
}
