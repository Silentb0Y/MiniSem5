package funki

type UserData struct {
	Username string
	Password string
}

func AddUsers(user *[]UserData) {
	*user = append(*user, UserData{
		Username: "yash",
		Password: "yash",
	}, UserData{
		Username: "abhijeet",
		Password: "abhijeet",
	}, UserData{
		Username: "sanika",
		Password: "sanika",
	})
}
