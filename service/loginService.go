package service

type LoginService interface {
	LoginUser(email string, password string) bool
}

type loginForm struct {
	email    string
	password string
}

func (form *loginForm) LoginUser(email string, password string) bool {
	return email == form.email && password == form.password
}

func StaticLoginService() LoginService {
	return &loginForm{
		email:    "westlife0615@naver.com",
		password: "bigin0520!",
	}
}
