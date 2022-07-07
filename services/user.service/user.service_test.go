package user_service_test

import (
	userService "mongodb-go/services/user.service"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	m "mongodb-go/models"
)

var userId string

func TestCreate(t *testing.T) {
	oid := primitive.NewObjectID()
	userId = oid.Hex()

	user := m.User{
		ID:        oid,
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
		t.Fail()
	} else {
		t.Log("La prueba de consulta de usuarios ha sido exitosa")
	}
}

func TestUpdate(t *testing.T) {
	user := m.User{
		Name:  "Gonzalo",
		Email: "email@email.com",
	}
	err := userService.Update(user, userId)
	if err != nil {
		t.Error("La actualizacion de usuario ha fallado")
		t.Fail()
	} else {
		t.Log("La prueba de actualizacion de usuario ha sido exitosa")
	}
}

func TestDelete(t *testing.T) {
	err := userService.Delete(userId)
	if err != nil {
		t.Error("La eliminacion de usuario ha fallado")
	} else {
		t.Log("La prueba de eliminación ha sido exitosa")
	}
}
