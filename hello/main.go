package main

func main() {
	Hello("Chris")
}

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return "Hello, " + name
}
