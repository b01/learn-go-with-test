package main

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return greeting(lang) + name
}

func greeting(lang string) string {
	prefix := "Hello, "

	switch lang {
	case "es":
		prefix = "Hola, "
	case "fr":
		prefix = "Bonjour, "
	}

	return prefix
}
func main () {
	Hello("World", "en")
}