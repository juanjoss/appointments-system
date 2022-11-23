package ports

/*
	Custom struct used in the login handler to
	capture users request.
*/
type LoginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
