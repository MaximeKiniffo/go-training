package profile

// Introduction returns a string introducing the person with the given name and programming language.
func Introduction(name string, language string) string {
	return "Je m'appelle " + name + ", je suis developpeur " + language
}

// Welcome returns a welcome message for the given name.
func Welcome(name string) string {
	return "Bienvenue, " + name + "!"
}
