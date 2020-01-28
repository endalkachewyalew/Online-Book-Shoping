package main

import (
	"fmt"
	"time"
	"github.com/lib/pq"
)


func getuser(Email string) (User, error) {
	//Retrieve
	res := User{}

	var id int
	var fullname string
	var gender string
	var dob pq.NullTime
	var country string
	var city string
	var phonenumber int
	var email string
	var password string

	err := db.QueryRow(`SELECT id, fullname, gender, dob, country, city, phone_number, email, password FROM users where email = '$1'`, Email).Scan(&id, &fullname, &gender, &dob, &country, &city, &phonenumber, &email, &password)
	if err == nil {
		res = User{ID: id, FullName: fullname, Gender: gender, DOB: dob.Time, Country: country, City: city, PhoneNumber: phonenumber, Email: email, Password: password}
	}

	return res, err
}
func allUsers() ([]User, error) {
	//Retrieve
	users := []User{}

	rows, err := db.Query(`SELECT id, fullname, gender, dob, country, city, phone_number, email, password FROM users order by id`)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			var id int
			var fullname string
			var gender string
			var dob pq.NullTime
			var country string
			var city string
			var phonenumber int
			var email string
			var password string

			err = rows.Scan(&id, &fullname, &gender, &dob, &country, &city, &phonenumber, &email, &password)
			if err == nil {
				currentUser := User{ID: id, FullName: fullname, Gender: gender, DOB: dob.Time, Country: country, City: city, PhoneNumber: phonenumber, Email: email, Password: password}

				if dob.Valid {
					currentUser.DOB = dob.Time
				}

				users = append(users, currentUser)
			} else {
				return users, err
			}
		}
	} else {
		return users, err
	}

	return users, err
}
func insertUser(fullname string, gender string, dob time.Time, country string, city string, phonenumber int, email string, password string) (int, error) {
	//Create
	var ID int
	err := db.QueryRow(`INSERT INTO users(fullname, gender, dob, country, city, phone_number, email, password) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`, fullname, gender, dob, country, city, phonenumber, email, password).Scan(&ID)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Last inserted ID: %v\n", ID)
	return ID, err
}

func updateUser(id int, fullname string, gender string, dob time.Time, country string, city string, phonenumber int, email string, password string) (int, error) {
	//Create
	res, err := db.Exec(`UPDATE books set fullname=$1, gender=$2, dob=$3, country=$4, city=$5, phone_number=$6, email=$7, password=$8 where id=$9 RETURNING id`, fullname, gender, dob, country, city, phonenumber, email, password, id)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), err
}

func removeUser(email string) (int, error) {
	//Delete
	res, err := db.Exec(`delete from users where email = '$1'`, email)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil
}