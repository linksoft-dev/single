package auth

import (
	"github.com/waykiss/wkcomps/validation"
	"golang.org/x/text/language"
)

// validate perform validationrelated to the user model
func validate(m *Model) error {
	v := validation.NewValidation(language.BrazilianPortuguese)
	v.IsIn("Status", m.Status.String(), StatusAll[:4]...)
	v.IsFilled("Name", m.Name, 2, 120)
	if v.IsValidEmailFormat("Email", m.Email) {
		v.IsFilled("Email", m.Email, 8, 120)
	}
	//validate hashed password
	v.IsByteLength("Password", m.Password, 60, 60)
	v.InRangeInt("Age", int(m.Age), 10, 100)
	return v.Error()
}

//checkPasswordPolicy check the password is following the policy
func checkPasswordPolicy(pwd string) (r error) {
	v := validation.NewValidation(language.BrazilianPortuguese)
	v.IsFilled("Senha", pwd, 4, 30)
	return v.Error()
}
