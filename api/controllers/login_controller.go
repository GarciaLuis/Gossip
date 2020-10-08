package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/garcialuis/Gossip/api/auth"
	"github.com/garcialuis/Gossip/api/models"
	"github.com/garcialuis/Gossip/api/responses"
	"github.com/garcialuis/Gossip/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

// Login Handler:
// swagger:route POST /login users LoginUser
// Logs in user given the login credentials
//
//	Responses:
//		200: authToken
//		422: description: Unprocessable Entity
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, models.Token{Token: token})
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	userModel := models.User{}

	user, err := userModel.FindUserByEmail(server.DB, email)

	if err != nil {
		return "", err
	}

	err = models.VerifyPassword(user.Password, password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return auth.CreateToken(user.ID)
}
