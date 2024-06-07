package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func index(w http.ResponseWriter,req *http.Request) {
	tpl.ExecuteTemplate(w,"index.tmpl",nil)
}

func gallery(w http.ResponseWriter,req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w,"gallery.tmpl",nil)

}

func person(w http.ResponseWriter,req *http.Request) {
	if !alreadyLoggedIn(req){
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}

	u := getuser(req)

	tpl.ExecuteTemplate(w,"person.tmpl",u)
}

func login(w http.ResponseWriter,req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w,req,"/person",http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")

		// now check the username in the db 
		var hashedPassword []byte 
		err := db.QueryRow("select password from users where username = ?",un).Scan(&hashedPassword)
		if err != nil {
			http.Error(w,"no user name exists of this name ",http.StatusNotFound)
			return
		}

		//compare the hashed password with typed password 

		err = bcrypt.CompareHashAndPassword(hashedPassword,[]byte(p))
		if err != nil {
			http.Error(w,"wrong password",http.StatusForbidden)
			return
		}

		// if both conditions are correct then redirecet the user into the person page 
		//after making the new login session 

		sID := uuid.NewV4().String()
		c := &http.Cookie{
			Name: "session",
			Value: sID,
		}
		http.SetCookie(w,c)

		dbSessions[c.Value] = un 

		http.Redirect(w,req,"/person",http.StatusSeeOther)
		return

	}

	tpl.ExecuteTemplate(w,"login.tmpl",nil)
}


func logout(w http.ResponseWriter,req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}
	
	
	c, err := req.Cookie("session")
	if err != nil {
		http.Error(w,"unable to get the cookie ",http.StatusSeeOther)
		return
	}

	delete(dbSessions,c.Value)


	c = &http.Cookie{
		Name: "session",
		Value: "",
		MaxAge: -1,
	}
	http.SetCookie(w,c)

	http.Redirect(w,req,"/",http.StatusSeeOther)
}

func signin(w http.ResponseWriter ,req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		// now check the username is exists in the db or not 
		var exists bool
		err := db.QueryRow("select exists (select 1 from users where username = ?)",un).Scan(&exists)
		if err != nil {
			http.Error(w,"unable to fetch the query from the database ", http.StatusInternalServerError)
			return
		}
		if exists {
			http.Error(w,"username is already taken ",http.StatusInternalServerError)
			return
		}

		sID := uuid.NewV4().String()
		c := &http.Cookie{
			Name: "session",
			Value: sID,
		}
		http.SetCookie(w,c)

		dbSessions[c.Value] = un 

		// now insert all the values inside the database 
		bs,err := bcrypt.GenerateFromPassword([]byte(p),bcrypt.MinCost)
		if err != nil {
			http.Error(w,"internal server error",http.StatusInternalServerError)
			return
		}


		_ , err = db.Exec("insert into users (username,password,firstname,lastname) values (?,?,?,?)",un,bs,f,l)
		if err != nil {
			http.Error(w,"unabel to insert the values inside the database ",http.StatusInternalServerError)
			return
		}

		http.Redirect(w,req,"/person",http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w,"signin.tmpl",nil)

}