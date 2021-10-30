# [my-portfolio](https://ibilalkayy.herokuapp.com/)

![Screenshot 2021-10-30 at 21-29-21 Bilal Khan](https://user-images.githubusercontent.com/64713734/139542036-27755020-e953-408f-8750-abf512702e0e.png)

## Intro
This repository contains the code of my portfolio website. It contains five pages that are Home, About, Resume, Portfolio, and Contact. The template of this app is taken from [bootstrapmademe.com](https://bootstrapmade.com/free-html-bootstrap-template-my-resume/) that is available for free to use.

---

## App Structure

This software is written in Golang, HTML, CSS, and JavaScript. The directory structure is based on **MVC model**. The software contains six directories in which Go code, Database, and website template is written separately. These directories are.

- **controllers**: This directory contains one file `controllers.go`. This file handles the main functionality of different functions like executing the template, sending an email using SMTP.
    
- **database**: This directory contains one file `db.go`. This file contains the code of a MySQL database that will connect, and insert the data.
    
- **middleware**: This directory has one file `middleware.go`. It handles the errors in a function and it works on a function to call the values from .env file based on a key.
   
- **models**: This directory has one file `models.go`. It contains a structure which has the name, email, subject and body to be used.

- **routes**: This directory has one file `routes.go`. It handles the paths that should be visited and uses middleware to check for errors.

- **views**: This directory contains two subdirectories **templates** and **static**. The first subdirectory contains all the HTML files. The second subdirectory contains all the CSS files, JavaScript and images also. 


---

## Files

- **.gitignore**: It contains all the files that should be ignored.
- **Procfile**: This file is used to deploy an app on heroku server.
- **main.go**: It calls the routes and use a port to run an app.

## Author Info

- YouTube - [ibilalkayy](https://www.youtube.com/channel/UCBLTfRg0Rgm4FtXkvql7DRQ)
- LinkedIn - [ibilalkayy](https://www.linkedin.com/in/ibilalkayy/)
- Twitter - [ibilalkayy](https://twitter.com/ibilalkayy)

[Back to Top](#my-portfolio)
