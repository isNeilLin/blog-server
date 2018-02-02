package conf

type mysql struct {
	username	string
	password	string
	port		string
}

var MySql mysql = mysql{"root","password","3306"}