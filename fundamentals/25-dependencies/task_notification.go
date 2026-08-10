package main

import "fmt"

// Notifier est une interface qui définit la méthode Notify
type Notifier interface {
	Notify() string
}

// ConsoleNotifier est une structure qui implémente l'interface Notifier
type ConsoleNotifier struct{}

// Notify est la méthode qui implémente l'interface Notifier pour ConsoleNotifier
func (c ConsoleNotifier) Notify() string {
	return "Notification envoyée à la console"
}

// TaskService est une structure qui dépend de l'interface Notifier
type TaskService struct {
	notifier Notifier
}

// Complete est une méthode de TaskService qui utilise l'interface Notifier pour envoyer une notification
func (t TaskService) Complete(taskName string) string {
	return "Tâche terminée : " + taskName + " — " + t.notifier.Notify()
}

// main est la fonction principale qui crée une instance de TaskService avec un ConsoleNotifier et complète des tâches
func main() {
	consoleNotifier := ConsoleNotifier{}
	taskService := TaskService{notifier: consoleNotifier}
	fmt.Println(taskService.Complete("apprendre les dépendances"))
	fmt.Println(taskService.Complete("apprendre les interfaces implicites"))
}
