package main

import "net/http"

func getuser(req *http.Request) (u user) {
	c, err := req.Cookie("session")
	if err != nil {
		return user{}
	}

	un,ok := dbSessions[c.Value]
	if !ok {
		return user{}
	}

	
	err = db.QueryRow("select username,firstname,lastname from users where username = ?",un).Scan(&u.UserName,&u.FirstName,&u.LastName)
	if err != nil {
		return user{}
	}

	return u 

}

func alreadyLoggedIn(req *http.Request) bool {
	c ,err := req.Cookie("session")
	if err!= nil {
		return false
	}

	un , ok := dbSessions[c.Value]
	if !ok {
		return false 
	}

	var exists bool 
	err = db.QueryRow("select exists(select 1 from users where username = ?)",un).Scan(&exists)
	if err != nil || !exists {
		return false
	}

	return true 
	
}