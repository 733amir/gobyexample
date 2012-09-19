package main

import ("fmt"; "net/url"; "strings")

func main() {
	u, err := url.Parse("postgres://user:pass@host.com:5432/path?k=v#frag")
	if err != nil { panic(err) }
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)
	fmt.Println(u.Host)
	split := strings.Split(u.Host, ":")
	fmt.Println(split[0])
	fmt.Println(split[1])
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
