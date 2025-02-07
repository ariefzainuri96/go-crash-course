package c1

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	user := User{
		Name: name,
		Membership: Membership{
			Type: membershipType,
		},
	}

	if membershipType == "premium" {
		user.Membership.MessageCharLimit = 1000
	} else {
		user.Membership.MessageCharLimit = 100
	}

	return user
}
