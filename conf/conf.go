package conf

var (
	Enviroment		string
	LocalDB 		= new(localDB)
)

type localDB struct{
	Name 	string
	Option	string
}
