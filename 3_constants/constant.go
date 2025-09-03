package main

const(
	port int=8080
	host string="localhost"
)

func main() {
	const pi float32=3.14

	println("Value of pi:",pi)

	println("Server is running on",host, "at port",port)

}
