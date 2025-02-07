package l5embeddedstructs

type sender struct {
	user
	rateLimit int
}

type user struct {
	name   string
	number int
}
