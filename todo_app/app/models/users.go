package models

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

func (u *User) CreateUser() (int64, error) {
	cmd := `insert into users(
		uuid,
		name,
		email,
		password,
		created_at) values(?,?,?,?,?)`

	result, err := Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("User created with ID: %d\n", id)
	return id, err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (User, error) {
	var user User
	cmd := `SELECT * FROM users WHERE email = ?`
	err := Db.QueryRow(cmd, email).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.PassWord, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

type Session struct {
	UUID      string
	CreatedAt time.Time
}
