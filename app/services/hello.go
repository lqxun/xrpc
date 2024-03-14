package services

func Hello(name string) (string, error) {
	return "hello, " + name, nil
}
