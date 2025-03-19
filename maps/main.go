package main

import "fmt"

type User struct {
	ID   int64
	Name string
}

func main() {
	
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 45, Name: "Bob"},
		{ID: 57, Name: "Eve"},
	}

	
	usersMap := make(map[int64]User, len(users))

	
	for _, user := range users {
		if _, exists := usersMap[user.ID]; !exists {
			usersMap[user.ID] = user
		}
	}

	
	searchID := int64(57)
	user, found := usersMap[searchID]
	if found {
		fmt.Printf("Найден пользователь: %v\n", user)
	} else {
		fmt.Println("Пользователь не найден")
	}
}
