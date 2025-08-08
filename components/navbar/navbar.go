package navbar

type Navbar struct {
	Username string
}

func NewNavbar(username string) *Navbar {
	return &Navbar{
		Username: username,
	}
}
