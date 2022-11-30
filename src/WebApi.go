package main;

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetAllPostsRoute(c *gin.Context) {
	All := GetAllPosts()
	
	var res Response = MakeServerResponse(200, All)

	c.JSON(http.StatusOK, res)
}

func getUserPostsRoute(c *gin.Context) {
	var id string = GetFieldFromContext(c, "id_")
	fmt.Println(id)
	var UserPosts []Post = getUserPostById(id);
	var res Response = MakeServerResponse(200, UserPosts)
	c.JSON(http.StatusOK, res)
}


func getUsersRoute(c *gin.Context) {
	var q string = GetFieldFromContext(c, "q")
	var Users []User;
	
	if q != "" {
		Users = getUsersByQuery(q)
	} else {
		Users = getUsers()
	}
	
	var res Response = MakeServerResponse(200, Users)

	c.JSON(http.StatusOK, res)
}

func getUserByIdRoute(c *gin.Context) {
	var uuid string = c.Param("uuid")
	var User User = getUserById(uuid)
	var res Response = MakeServerResponse(200, User)
	c.JSON(http.StatusOK, res)
}

/* AUTHENTICATION AND OPERATIONS */
/*
Implemented: Login, Sign Up.
Not implemented: data access (Update, add, delete)
*/
func login(c *gin.Context) {
	var LoginForm UserLogin;
	
	c.BindJSON(&LoginForm);
	
	var resp Response
	
	if len(LoginForm.Token) > 0 {
		resp = AuthenticateUserJWT(LoginForm.Token)
	} else {
		if len(LoginForm.Password) > 0 && len(LoginForm.Email) > 0 {
			User, err := AuthenticateUserByEmailAndPwd(LoginForm.Password, LoginForm.Email)
			
			if err.Ok {
				resp = MakeServerResponse(200, User)
			} else {
				resp = MakeServerResponse(500, err.Text)
				fmt.Println("", resp.Data)
			}

		} else {
			resp = MakeServerResponse(500, "Missing request attributes, Email or password not specified.")
		}
	}

	c.JSON(http.StatusOK, resp);
}


func signUp(c *gin.Context) {
	var newUser User
	
	c.BindJSON(&newUser);
	
	if isEmpty(newUser.Email) || isEmpty(newUser.PasswordHash) || isEmpty(newUser.UserName) {
		c.JSON(http.StatusOK, MakeServerResponse(500, "The server could not get the Email, password or user name. please check your request then try again L86"))
	} else {
		newUser.setDefaults();
		// Hash the password.
		newUser.PasswordHash = sha256_(newUser.PasswordHash)
		fmt.Println(newUser.PasswordHash)
		var Resp Response = AddUser(newUser) // Creates the user and sets the Token.
		c.JSON(http.StatusOK, Resp)
	}
	
}

