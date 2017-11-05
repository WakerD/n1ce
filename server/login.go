package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func GetLogout(c *gin.Context) {
	httputil.DelCookie(c.Writer, c.Request, "user_sess")
	httputil.DelCookie(c.Writer, c.Request, "user_last")
	c.Redirect(303, "/")
}

func GetLoginToken(c *gin.Context) {
	in := &tokenPayLoad{}
}
