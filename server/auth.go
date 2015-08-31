package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"regexp"
)

const ()

var (
	nameRegexp, _ = regexp.Compile("^[a-zA-z\u4e00-\u9fa5][a-zA-Z0-9\u4e00-\u9fa5]{3,9}$")
	pwdRegexp, _  = regexp.Compile(`^[^\s]{8,15}$`)
	mailRegexp, _ = regexp.Compile(`^([0-9a-zA-Z]([-\.\w]*[0-9a-zA-Z])*@([0-9a-zA-Z][-\w]*[0-9a-zA-Z]\.)+[a-zA-Z]{2,9})$`)
)

type User struct {
	Email string `bson:"username"`
	Hash  string `bson:"hash"`
}

type AuthResult struct {
	Status int    `json:"status"`
	OK     bool   `json:"ok"`
	Info   string `json:"msg"`
}

func Login(c *gin.Context) {
}

func Register(c *gin.Context) {
	//var usr *User = &User{}
	var res *AuthResult = &AuthResult{Status: http.StatusOK}

	//always return the result to client
	defer c.JSON(res.Status, res)

	//get form data
	mail := c.PostForm("Email")
	pwd := c.PostForm("Password")

	log.Println("User Register:", mail, "|", pwd)

	if err := checkMail(mail); err != nil {
		res.Status = http.StatusNotAcceptable
		res.Info = err.Error()
		return
	}

	if err := checkPassword(pwd); err != nil {
		res.Status = http.StatusNotAcceptable
		res.Info = err.Error()
		return
	}

	n, err := FindUserExistedByMail(mail)

	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Info = "Error: Internal Server Error"
		return
	}

	if n > 0 {
		res.Status = http.StatusNotAcceptable
		res.Info = "Error: User Existed"
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	if err != nil {
		res.Status = http.StatusNotAcceptable
		res.Info = "Error: Couldn't Save Password"
		return
	}

	user := &User{Email: mail, Hash: string(hash)}
	err = SaveUser(user)
	if err != nil {
		res.Status = http.StatusNotAcceptable
		res.Info = "Error: Couldn't Register User"
		return
	}

	res.Status = http.StatusOK
	res.OK = true
	res.Info = "Registed"

	return
}

func checkMail(mail string) (err error) {
	if mail == "" {
		err = errors.New("Error: Empty Email Address")
		return
	}

	if matched := mailRegexp.MatchString(mail); matched == false {
		err = errors.New("Error: Invalid Email Address")
		return
	}

	return nil
}

func checkPassword(pwd string) (err error) {
	if pwd == "" {
		err = errors.New("Error: Empty Password")
		return
	}

	if matched := pwdRegexp.MatchString(pwd); matched == false {
		err = errors.New("Error: Invalid Password")
		return
	}

	return nil
}

func checkUsername(name string) (err error) {
	if name == "" {
		err = errors.New("Error: Empty Username")
		return
	}

	if matched := nameRegexp.MatchString(name); matched == false {
		err = errors.New("Error: Invalid Username")
		return
	}
	return nil
}
