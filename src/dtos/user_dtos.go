package dtos

import "github.com/asaskevich/govalidator"

type RegisterParam struct {
	Email     string `json:"email"  valid:"required~Email wajib diisi,email~Format email tidak sesuai"`
	Password  string `json:"password" valid:"required~Password wajib diisi,minstringlength(8)~Password minimal memiliki 8 karakter"`
	FirstName string `json:"first_name"  valid:"required~First Name wajib diisi"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"  valid:"required~Phone wajib diisi"`
}

func (r *RegisterParam) Validate() (err error) {
	_, err = govalidator.ValidateStruct(r)
	return
}

type LoginParam struct {
	Email    string `json:"email" valid:"required~Email wajib diisi"`
	Password string `json:"password" valid:"required~Password wajib diisi"`
}

func (l *LoginParam) Validate() (err error) {
	_, err = govalidator.ValidateStruct(l)
	return
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}
