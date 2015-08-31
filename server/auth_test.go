package server

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestAuth(t *testing.T) {
	//Email Validation
	assert.Equal(t, nil, checkMail("aspirin2d@outlook.com"))
	assert.Equal(t, nil, checkMail("754055516@qq.com"))
	assert.Equal(t, "Error: Empty Email Address", checkMail("").Error())
	assert.Equal(t, "Error: Invalid Email Address", checkMail("aspirin2d@").Error())

	//Password Validation
	assert.Equal(t, nil, checkPassword("Passw0rd!"))
	assert.Equal(t, nil, checkPassword("Pa*:<>s21&"))
	assert.Equal(t, "Error: Empty Password", checkPassword("").Error())
	assert.Equal(t, "Error: Invalid Password", checkPassword("Hello World").Error())
	assert.Equal(t, "Error: Invalid Password", checkPassword("HelloWorld ").Error())
	assert.Equal(t, "Error: Invalid Password", checkPassword("Hello World><>>???").Error())
}
