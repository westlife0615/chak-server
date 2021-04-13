package handler

type loginInput struct {
	Email    string `json:email`
	Password string `json:password`
}

type AuthHandler struct {

}
