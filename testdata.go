package yogofn

var estados = []string{
	"Aguascalientes",
	"Baja California",
	"Baja California Sur",
	"Campeche",
	"Chiapas",
	"Chihuahua",
	"Coahuila",
	"Colima",
	"Durango",
	"Guanajuato",
	"Guerrero",
	"Hidalgo",
	"Jalisco",
	"México",
	"Michoacán",
	"Morelos",
	"Nayarit",
	"Nuevo León",
	"Oaxaca",
	"Puebla",
	"Querétaro",
	"Quintana Roo",
	"San Luis Potosí",
	"Sinaloa",
	"Sonora",
	"Tabasco",
	"Tamaulipas",
	"Tlaxcala",
	"Veracruz",
	"Yucatán",
	"Zacatecas",
}

type persona struct {
	nombre string
	edad   int
}

var adultos = []persona{
	{"Jose", 33},
	{"John", 56},
	{"Isabel", 19},
	{"Ramón", 39},
}

var ninos = []persona{
	{"Becky", 12},
	{"Javier", 8},
	{"Shanika", 11},
	{"Camila", 7},
	{"Diego", 5},
	{"Alejandro", 17},
}

var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
