package services

import (
	m "../../models"
	userRepository "../../repositories/user.repository"
)

func Create(user m.User) error {

	err := userRepository.Create(user)

	if err != nil {
		return err
	}
	return nil
}

func Read() (m.Users, error) {

	var users m.Users

	return users, nil
}

func Update(user m.User, userId string) error {

	return nil
}

func Delete(userId string) error {

	return nil
}
