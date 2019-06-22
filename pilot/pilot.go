package main

type Pilot struct {
	Name string
	Life int
	Age int
	Aircraft int

}

func main() {
	var donnie Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = student.AIRCRAFT1

	fmt.Println(donnie)
}