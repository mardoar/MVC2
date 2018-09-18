package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Estudiante struct {
	ID       string `db:"id"       json:"id"`
	Nombre   string `db:"nombre"   json:"nombre"`
	Apellido string `db:"apellido" json:"apellido"`
	Carrera  string `db:"carrera"  json:"carrera"`
}

var DB *sqlx.DB

func OpenDB() {
	db, err := sqlx.Open("mysql", "root:@tcp(localhost:3307)/estudiantesgo")

	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
	err = DB.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func InsertEstudiante(estudiante Estudiante) (err error) {
	_, err = DB.NamedExec(`INSERT INTO estudiantes (nombre, apellido, carrera) 
	VALUES (:nombre, :apellido, :carrera)`, &estudiante)
	return
}

func ListEstudiantes() (estudiantes []Estudiante, err error) {
	err = DB.Select(&estudiantes, `SELECT id, nombre, apellido, carrera FROM estudiantes`)
	return
}

func GetEstudiante(estudianteID string) (estudiante Estudiante, err error) {
	err = DB.Get(&estudiante, `SELECT id, nombre, apellido, carrera FROM estudiantes WHERE id=?`, estudianteID)
	return
}

func UpdateEstudiante(estudiante Estudiante) (err error) {
	_, err = DB.NamedExec(`UPDATE estudiantes SET nombre=:nombre, apellido=:apellido, carrera=:carrera WHERE id=:id`, estudiante)
	return
}

func DeleteEstudiante(estudianteID string) (err error) {
	_, err = DB.Exec(`DELETE FROM estudiantes WHERE id=?`, estudianteID)
	return
}
