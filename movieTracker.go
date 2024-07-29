package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Constants for file paths
const (
	UNSEEN_FILE = "unseen_movies.json"
	SEEN_FILE   = "seen_movies.json"
)

// clear clears the console screen.
func clear() {
	fmt.Print("\033[H\033[2J")
}

// pressAnyKey prompts the user to press Enter to continue and then clears the screen.
func pressAnyKey() {
	fmt.Println("\nPress Enter to continue...")
	fmt.Scanln()
	clear()
}

// loadMovies loads movies from a JSON file and returns them as a slice of strings.
func loadMovies(filePath string) []string {
	var movies []string
	if _, err := os.Stat(filePath); err == nil {
		file, err := ioutil.ReadFile(filePath)
		if err == nil {
			json.Unmarshal(file, &movies)
		}
	}
	return movies
}

// saveMovies saves a slice of movies to a JSON file.
func saveMovies(filePath string, movies []string) {
	file, _ := json.MarshalIndent(movies, "", "    ")
	ioutil.WriteFile(filePath, file, 0644)
}

// addMovie adds a new movie to the unseen list.
func addMovie() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the movie name: ")
	movie, _ := reader.ReadString('\n')
	movie = strings.TrimSpace(movie)
	unseenMovies := loadMovies(UNSEEN_FILE)
	unseenMovies = append(unseenMovies, movie)
	saveMovies(UNSEEN_FILE, unseenMovies)
	fmt.Printf("'%s' has been added to the unseen movies list.\n", movie)
}

// displayUnseenMovies displays the list of unseen movies.
func displayUnseenMovies() {
	unseenMovies := loadMovies(UNSEEN_FILE)
	if len(unseenMovies) > 0 {
		fmt.Println("Unseen Movies:")
		for i, movie := range unseenMovies {
			fmt.Printf("%d. %s\n", i+1, movie)
		}
	} else {
		fmt.Println("No unseen movies found.")
	}
}

// displaySeenMovies displays the list of seen movies.
func displaySeenMovies() {
	seenMovies := loadMovies(SEEN_FILE)
	if len(seenMovies) > 0 {
		fmt.Println("Seen Movies:")
		for i, movie := range seenMovies {
			fmt.Printf("%d. %s\n", i+1, movie)
		}
	} else {
		fmt.Println("No seen movies found.")
	}
}

// markMovieAsSeen marks a movie from the unseen list as seen and moves it to the seen list.
func markMovieAsSeen() {
	displayUnseenMovies()
	unseenMovies := loadMovies(UNSEEN_FILE)
	seenMovies := loadMovies(SEEN_FILE)
	if len(unseenMovies) > 0 {
		var movieIndex int
		fmt.Print("Enter the number of the movie you have seen: ")
		fmt.Scanln(&movieIndex)
		movieIndex--
		if movieIndex >= 0 && movieIndex < len(unseenMovies) {
			movie := unseenMovies[movieIndex]
			unseenMovies = append(unseenMovies[:movieIndex], unseenMovies[movieIndex+1:]...)
			seenMovies = append(seenMovies, movie)
			saveMovies(UNSEEN_FILE, unseenMovies)
			saveMovies(SEEN_FILE, seenMovies)
			fmt.Printf("'%s' has been moved to the seen movies list.\n", movie)
		} else {
			fmt.Println("Invalid selection.")
		}
	}
}

// deleteMovie deletes a movie from either the unseen or seen list based on user choice.
func deleteMovie() {
	for {
		var choice string
		fmt.Print("Do you want to delete from unseen (u) or seen (s) list? you can also press (n) if you want to go back: ")
		fmt.Scanln(&choice)
		fmt.Println()
		if choice == "u" {
			displayUnseenMovies()
			unseenMovies := loadMovies(UNSEEN_FILE)
			if len(unseenMovies) > 0 {
				var movieIndex int
				fmt.Print("Enter the number of the movie to delete: ")
				fmt.Scanln(&movieIndex)
				movieIndex--
				if movieIndex >= 0 && movieIndex < len(unseenMovies) {
					movie := unseenMovies[movieIndex]
					unseenMovies = append(unseenMovies[:movieIndex], unseenMovies[movieIndex+1:]...)
					saveMovies(UNSEEN_FILE, unseenMovies)
					fmt.Printf("'%s' has been deleted from the unseen movies list.\n", movie)
					break
				} else {
					fmt.Println("Invalid selection.")
				}
			}
		} else if choice == "s" {
			displaySeenMovies()
			seenMovies := loadMovies(SEEN_FILE)
			if len(seenMovies) > 0 {
				var movieIndex int
				fmt.Print("Enter the number of the movie to delete: ")
				fmt.Scanln(&movieIndex)
				movieIndex--
				if movieIndex >= 0 && movieIndex < len(seenMovies) {
					movie := seenMovies[movieIndex]
					seenMovies = append(seenMovies[:movieIndex], seenMovies[movieIndex+1:]...)
					saveMovies(SEEN_FILE, seenMovies)
					fmt.Printf("'%s' has been deleted from the seen movies list.\n", movie)
					break
				} else {
					fmt.Println("Invalid selection.")
				}
			}
		} else if choice == "n" {
			break
		} else {
			fmt.Println("Invalid choice. Please enter 'u' or 's' or 'n'")
		}
	}
}

// randomUnseenMovie selects and displays a random movie from the unseen list.
func randomUnseenMovie() {
	unseenMovies := loadMovies(UNSEEN_FILE)
	if len(unseenMovies) > 0 {
		rand.Seed(time.Now().UnixNano())
		movie := unseenMovies[rand.Intn(len(unseenMovies))]
		fmt.Printf("Randomly selected unseen movie: %s\n", movie)
	} else {
		fmt.Println("No unseen movies found.")
	}
}

// displayMenu displays the main menu options to the user.
func displayMenu() {
	fmt.Print("\nMovie Watch List Menu\n\n")
	fmt.Println("1. Add a movie")
	fmt.Println("2. Get a random unseen movie")
	fmt.Println("3. Display unseen movies")
	fmt.Println("4. Display seen movies")
	fmt.Println("5. Mark a movie as seen")
	fmt.Println("6. Delete a movie")
	fmt.Print("7. Exit\n\n")
}

// main runs the main menu-driven program.
func main() {
	for {
		displayMenu()
		var choice string
		fmt.Print("Enter your choice (1-7): ")
		fmt.Scanln(&choice)
		fmt.Println()
		switch choice {
		case "1":
			addMovie()
			pressAnyKey()
		case "2":
			randomUnseenMovie()
			pressAnyKey()
		case "3":
			displayUnseenMovies()
			pressAnyKey()
		case "4":
			displaySeenMovies()
			pressAnyKey()
		case "5":
			markMovieAsSeen()
			pressAnyKey()
		case "6":
			deleteMovie()
			pressAnyKey()
		case "7":
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
