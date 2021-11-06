package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
)

const (
	username = "root"
	password = "root_password"
	hostname = "127.0.0.1:3306"
	dbname   = "golangdb"
)

type Company struct {
	Id   int
	Name string
	City string
}

func connString() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func dbConnection() (db *sql.DB) {
	db, err := sql.Open("mysql", connString())
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/add", Add)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/export", Export)
	http.HandleFunc("/search", Search)
	http.ListenAndServe(":8080", nil)
}

var tmpl = template.Must(template.ParseGlob("pages/*"))

func Search(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()
	sname := r.FormValue("name")
	searchQuery, err := db.Query("SELECT * FROM Company WHERE name like ?", "%"+sname+"%")
	if err != nil {
		log.Fatal(err)
	}
	defer searchQuery.Close()

	cmp := Company{}
	res := []Company{}

	var (
		id         int
		name, city string
	)
	for searchQuery.Next() {
		err = searchQuery.Scan(&id, &name, &city)
		if err != nil {
			log.Fatal(err)
		}
		cmp.Id = id
		cmp.Name = name
		cmp.City = city
		res = append(res, cmp)
	}

	tmpl.ExecuteTemplate(w, "index.html", res)
}

func Export(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()
	selDB, err := db.Query("SELECT * FROM Company ORDER BY id DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer selDB.Close()

	cmp := Company{}
	res := []Company{}
	var (
		id         int
		name, city string
	)
	for selDB.Next() {
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			log.Fatal(err)
		}
		cmp.Id = id
		cmp.Name = name
		cmp.City = city
		res = append(res, cmp)
	}

	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "Id")
	f.SetCellValue("Sheet1", "B1", "Name")
	f.SetCellValue("Sheet1", "C1", "City")

	row := 2
	for _, v := range res {
		rowNum := strconv.Itoa(row)
		f.SetCellValue("Sheet1", "A"+rowNum, v.Id)
		f.SetCellValue("Sheet1", "B"+rowNum, v.Name)
		f.SetCellValue("Sheet1", "C"+rowNum, v.City)
		row++
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment;filename=ExportedData.xlsx")
	w.Header().Set("File-Name", "ExportedData.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	f.Write(w)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()
	id := r.URL.Query().Get("id")
	delQuery, err := db.Prepare("DELETE FROM Company WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer delQuery.Close()
	delQuery.Exec(id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()
	eid := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Company WHERE id=?", eid)
	if err != nil {
		log.Fatal(err)
	}
	defer selDB.Close()

	var (
		id         int
		name, city string
	)
	cmp := Company{}
	for selDB.Next() {
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			log.Fatal(err)
		}
		cmp.Id = id
		cmp.Name = name
		cmp.City = city
	}

	tmpl.ExecuteTemplate(w, "edit.html", cmp)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		if id == "" {
			insQuery, err := db.Prepare("INSERT INTO Company(name, city) VALUES(?,?)")
			if err != nil {
				log.Fatal(err)
			}
			defer insQuery.Close()
			insQuery.Exec(name, city)
		} else {
			updQuery, err := db.Prepare("UPDATE Company SET name=?, city=? WHERE id=?")
			if err != nil {
				panic(err.Error())
			}
			defer updQuery.Close()
			updQuery.Exec(name, city, id)
		}
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Add(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "add.html", nil)
}
func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()
	sid := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Company WHERE id=?", sid)
	if err != nil {
		log.Fatal(err)
	}
	defer selDB.Close()

	var (
		id         int
		name, city string
	)
	cmp := Company{}
	for selDB.Next() {
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			log.Fatal(err)
		}
		cmp.Id = id
		cmp.Name = name
		cmp.City = city
	}

	tmpl.ExecuteTemplate(w, "show.html", cmp)
}

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()
	selDB, err := db.Query("SELECT * FROM Company ORDER BY id DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer selDB.Close()

	var (
		id         int
		name, city string
	)
	cmp := Company{}
	res := []Company{}
	for selDB.Next() {
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		cmp.Id = id
		cmp.Name = name
		cmp.City = city
		res = append(res, cmp)
	}

	tmpl.ExecuteTemplate(w, "index.html", res)
}
