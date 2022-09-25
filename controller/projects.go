package controller

import (
	"context"
	"fmt"
	"html/template"
	"mvcweb/connection"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
)

type ProjectData struct {
	Name,Description,Image string
	StartDate,EndDate pgtype.Date
	Technologies[]string	
}

var projects []ProjectData 


func GetHome(w http.ResponseWriter, r *http.Request) {

	data, err := connection.Conn.Query(context.Background(), "SELECT name,start_date,end_date,description,technologies,image FROM public.tb_projects;")

	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	var dataResult []ProjectData

	for data.Next() {
		var project = ProjectData{}

		var err = data.Scan(&project.Name, &project.StartDate, &project.EndDate, &project.Description, &project.Technologies, &project.Image)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		dataResult = append(dataResult, project)
	}
	
	var view, templErr = template.ParseFiles("views/index.html")	
	if err != nil {
		panic(templErr.Error())
	}
	view.Execute(w, dataResult)
}

func PostAddProject(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
	name := r.PostForm.Get("name")
	description := r.PostForm.Get("description")
	startDate := r.PostForm.Get("start-date")
	endDate := r.PostForm.Get("end-date")
	techlist := r.PostForm["checkbox"]
	fmt.Println(techlist)

	query := `INSERT INTO public.tb_projects(
		name, start_date, end_date, description, technologies, image)
		VALUES ($1,$2,$3,$4,$5,$6)`

	data , err := connection.Conn.Exec(context.Background(),query, name,startDate,endDate,description,techlist,"https://images.unsplash.com/photo-1498050108023-c5249f4df085?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8Y29kaW5nfGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(data)
	

	http.Redirect(w,r,"/form-add-project", http.StatusFound)
}