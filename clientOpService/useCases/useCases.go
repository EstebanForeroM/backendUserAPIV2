package usecases

func NewUser(db ClientOpDb, user User) error {
    return db.CreateUser(user)
}

func DeleteUser(db ClientOpDb, userId string) error {
    return db.DeleteUser(userId)
}

type User struct {
    Name string
    Email string
    Id string
}

type ClientOpDb interface {
    CreateUser(user User) error
    DeleteUser(userId string) error
}
