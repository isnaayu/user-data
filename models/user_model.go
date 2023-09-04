package models

type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var users = []User{
    {ID: "1", Name: "Isna Ayu Muarofah", Email: "isnaayu2@gmail.com"},
    {ID: "2", Name: "Refal Hady", Email: "refalhady@gmail.com"},
}

// GetAllUsers mengambil semua pengguna
func GetAllUsers() []User {
    return users
}

// GetUserByID mengambil pengguna berdasarkan ID
func GetUserByID(id string) (*User, error) {
    for _, user := range users {
        if user.ID == id {
            return &user, nil
        }
    }
    return nil, nil
}
