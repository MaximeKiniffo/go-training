package main

import "fmt"

// Notifier est une interface qui définit la méthode Notify
type Notifier interface {
	Notify() string
}

// ConsoleNotifier est une structure qui implémente l'interface Notifier
type ConsoleNotifier struct{}

// EmailNotifier est une structure qui implémente l'interface Notifier
type EmailNotifier struct {
	Address string
}

// SMSNotifier est une structure qui implémente l'interface Notifier
type SMSNotifier struct {
	Number string
}

// Notify est la méthode qui implémente l'interface Notifier pour SMSNotifier
func (s SMSNotifier) Notify() string {
	return fmt.Sprintf("Notification par SMS envoyée au numéro %s", s.Number)
}

// Notify est la méthode qui implémente l'interface Notifier pour EmailNotifier
func (e EmailNotifier) Notify() string {
	return fmt.Sprintf("Notification par email envoyée à %s", e.Address)
}

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
	emailNotifier := EmailNotifier{Address: "john@example.com"}
	smsNotifier := SMSNotifier{Number: "1234567890"}

	taskService1 := TaskService{notifier: consoleNotifier}
	taskService2 := TaskService{notifier: emailNotifier}
	taskService3 := TaskService{notifier: smsNotifier}

	fmt.Println(taskService1.Complete("apprendre les dépendances"))
	fmt.Println(taskService2.Complete("apprendre les interfaces implicites"))
	fmt.Println(taskService3.Complete("apprendre les notifications"))
}
