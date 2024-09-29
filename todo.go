package main

import (
	"fmt"
	"os"
)

// Tâche représente une tâche avec un ID et un nom
type Task struct {
	ID   int
	Name string
}

// Liste des tâches
var tasks []Task
var nextID int = 1

// Fonction pour afficher le menu
func displayMenu() {
	fmt.Println("=== Gestionnaire de Tâches ===")
	fmt.Println("1. Ajouter une tâche")
	fmt.Println("2. Supprimer une tâche")
	fmt.Println("3. Afficher les tâches")
	fmt.Println("4. Quitter")
	fmt.Print("Choisissez une option: ")
}

// Fonction pour ajouter une tâche
func addTask(name string) {
	task := Task{ID: nextID, Name: name}
	tasks = append(tasks, task)
	nextID++
	fmt.Println("Tâche ajoutée avec succès.")
}

// Fonction pour supprimer une tâche
func deleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Tâche supprimée avec succès.")
			return
		}
	}
	fmt.Println("Aucune tâche trouvée avec cet ID.")
}

// Fonction pour afficher toutes les tâches
func displayTasks() {
	if len(tasks) == 0 {
		fmt.Println("Aucune tâche disponible.")
		return
	}
	fmt.Println("Liste des tâches :")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Tâche: %s\n", task.ID, task.Name)
	}
}

func main() {
	for {
		displayMenu()
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name string
			fmt.Print("Entrez le nom de la tâche: ")
			fmt.Scan(&name)
			addTask(name)
		case 2:
			var id int
			fmt.Print("Entrez l'ID de la tâche à supprimer: ")
			fmt.Scan(&id)
			deleteTask(id)
		case 3:
			displayTasks()
		case 4:
			fmt.Println("Au revoir!")
			os.Exit(0)
		default:
			fmt.Println("Option invalide, veuillez réessayer.")
		}
	}
}
