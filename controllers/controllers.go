package controllers

// Importing the libraries
import (
	"html/template"
	"log"
	"net/http"

	"github.com/ibilalkayy/Portfolio/middleware"
	"github.com/ibilalkayy/Portfolio/models"
	"gopkg.in/gomail.v2"
)

// TemplatesList() will list all the files that are used in the template
// It will execute them and will be called from the base.html
func TemplatesList() *template.Template {
	files := []string{
		"views/templates/header.html",
		"views/templates/navbar.html",
		"views/templates/about.html",
		// "views/templates/facts.html",
		"views/templates/skills.html",
		"views/templates/resume.html",
		"views/templates/portfolio.html",
		// "views/templates/services.html",
		"views/templates/testimonials.html",
		"views/templates/contact.html",
		"views/templates/footer.html",
		"views/templates/base.html",
	}
	return template.Must(template.ParseFiles(files...))
}

// Parsing the page error page and setting a function to a template
var Error = template.Must(template.ParseFiles("views/templates/pageerror.html"))
var Tmpl = TemplatesList()

// Using the LoadEnvVariable() to fetch an email and app password
// Setting a same email address in both From and To, because sending requires app password which is not possible to take from each user.
// Showing the user email address in Reply-To section of the email
func Email(value [4]string) [4]string {
	mail := gomail.NewMessage()

	myEmail := middleware.LoadEnvVariable("EMAIL")
	appPassword := middleware.LoadEnvVariable("APP_PASSWORD")
	mail.SetHeader("From", myEmail)
	mail.SetHeader("To", myEmail)
	mail.SetHeader("Reply-To", value[1])
	mail.SetHeader("Subject", value[2])
	mail.SetBody("text/plain", value[0]+",\n\n"+value[3])

	a := gomail.NewDialer("smtp.gmail.com", 587, myEmail, appPassword)
	if err := a.DialAndSend(mail); err != nil {
		log.Fatal(err)
	}

	return value
}

// Executing the page error if page not found
// Executing the template again if the data is not posted
// Taking the form values and using them to send email
func Template(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		return Error.Execute(w, nil)
	}
	if r.Method != http.MethodPost {
		return Tmpl.Execute(w, nil)
	}

	send := models.Send{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}

	_ = send
	value := [4]string{send.Name, send.Email, send.Subject, send.Message}
	_ = Email(value)
	// database.Connection(send.Name, send.Email, send.Subject, send.Message)
	sm := struct{ Success bool }{true}
	return Tmpl.Execute(w, sm)
}
