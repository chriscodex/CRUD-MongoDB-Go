package user_service_test

import (
	userService "mongodb-go/services/user.service"
	"testing"
	"time"

	m "mongodb-go/models"
)

func TestCreate(t *testing.T) {

	user := m.User{
		Name:      "Christian",
		Email:     "christian@email.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := userService.Create(user)

	if err != nil {
		t.Error("La prueba de creacion de usuario, ha fallado")
		t.Fail()
	} else {
		t.Log("La prueba ha sido exitosa!")
	}
}

func TestRead(t *testing.T) {
	users, err := userService.Read()
	if err != nil {
		t.Error("La consulta de usuarios ha fallado")
		t.Fail()
	}
	if len(users) == 0 {
		t.Error("La consulta no ha retornado datos")
	} else {
		t.Log()
	}
}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}
