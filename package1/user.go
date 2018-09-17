package package1

type User struct {
	Name string
	addr int
	age int
	*test1
}

func (u User)GetAge() int {
	return u.age
}

func (u *User)SetAge(new_age int)  {
	u.SetU(string(new_age))
	
	u.age = new_age
}

func (u User)GetAddr() int {
	return u.addr
}

func (u *User)SetAddr(new_addr int)  {
	u.addr = new_addr
}

func NewUser(name string, addr, age int) *User {
	return &User{
		Name:name,
		addr:addr,
		age:age,
	}
}

type test1 struct {
	U string
}

func (t test1)GetU() string {
	return t.U
}
func (t *test1)SetU(n_u string) {
	t.U = n_u
}