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
	} else {
		t.Log("La prueba ha sido exitosa!")
	}
}

func TestRead(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}
