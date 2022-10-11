package pointers

func AppendGreeting(s *string) {
	*s = "Hi, " + *s
}
