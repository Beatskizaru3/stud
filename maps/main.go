package main

import "fmt"

type User struct {
	ID   int64
	Name string
}

func main() {
	// Создаем слайс пользователей
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 45, Name: "Bob"},
		{ID: 57, Name: "Eve"},
	}

	// Создаем map для быстрого поиска по ID
	usersMap := make(map[int64]User, len(users))

	// Заполняем map данными из слайса
	for _, user := range users {
		if _, exists := usersMap[user.ID]; !exists {
			usersMap[user.ID] = user
		}
	}

	// Поиск пользователя по ID
	searchID := int64(57)
	user, found := usersMap[searchID]
	if found {
		fmt.Printf("Найден пользователь: %v\n", user)
	} else {
		fmt.Println("Пользователь не найден")
	}
}
