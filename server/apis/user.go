package apis

import (
	"fmt"
	"log"
	"app/server/database"
)

type User struct {
	Id                                  int
	Name, Password, Icon, Info, Academy string
	Level                               int64
}

func Create(users []User) (createCount int, err error){
	stmt, err := database.SqlDB.Prepare(
		"Create into " +
		"user(name, password, icon, info, academy, level) " +
		"values(?, ?, ?, ?, ?, ?)")

	defer stmt.Close()

	if err != nil {
		log.Fatalln("users Create sql error", err)
	}

	for _, user := range users {

		rs, err := stmt.Exec(user.Name, user.Password, user.Icon, user.Info, user.Academy, user.Level)

		fmt.Println(rs)
		if err != nil {
			log.Fatalln("user Create error", err)
		} else {
			createCount++
		}
	}

	return createCount, err
}

func Delete(ids []int) (deleteCount int, err error) {
	stmt, err := database.SqlDB.Prepare("delete from user where id = ?")
	if err != nil {
		log.Fatalln("users delete sql error", err)
	}

	for id := range ids {
		rs, err := stmt.Exec(id)
		if err != nil {
			log.Fatalln("user delete error", err)
		} else {
			deleteCount++
		}
		fmt.Print(rs)
	}

	return
}

func Update(user User) (err error) {
	stmt, err := database.SqlDB.Prepare("update user set name=?, icon=?, info=?, academy=? where id=?")

	if err != nil {
		log.Fatalln("update user sql error", err)
	}
	stmt.Exec(user.Name, user.Icon, user.Info, user.Academy, user.Id)

	return
}

func Retrieve() (users []User, err error){

	rows, err := database.SqlDB.Query("select * from user")
	defer rows.Close()
	if err != nil {
		log.Fatalln("users retrieve sql error", err)
	}

	var id int
	var name, password, icon, info, academy string
	var level int64
	for rows.Next(){
		err = rows.Scan(&id, &name, &password, &icon, &info, &academy, &level)
		if err != nil {
			log.Fatalln("user retrieve error", err)
		} else {
			user := User{
				Id: id,
				Name: name,
				Password:password,
				Icon: icon,
				Info:info,
				Academy:academy,
				Level:level,
			}

			users = append(users, user)
		}
	}

	return
}



