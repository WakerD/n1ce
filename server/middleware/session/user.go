package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func SetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *model.User

		t, err := token.ParseRequest(c.Request, func(t *token.Token) (string, error) {
			var err error
			user, err = store.GetUserLogin(c, t.Text)
			return user.Hash, err
		})
		if err == nil {
			confv := c.MustGet("config")
			if conf, ok := confv.(*model.Settings); ok {
				user.Admin = conf.IsAdmin(user)
			}
			c.Set("user", user)

			// if this is a session token (ie not the API token)
			// this means the user is accessing with a web browser,
			// so we should implement CSRF protection measures.
			if t.Kind == token.SessToken {
				err = token.CheckCsrf(c.Request, func(t *token.Token) (string, error) {
					return user.Hash, nil
				})
				// if csrf token validation fails, exit immediately
				// with a not authorized error.
				if err != nil {
					c.AbortWithStatus(http.StatusUnauthorized)
					return
				}
			}
		}
		c.Next()
	}
}
