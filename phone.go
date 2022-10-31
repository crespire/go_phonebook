package main

import (
	"fmt"
)

func list(book map[string]string) {
	if len(book) == 0 {
		fmt.Println("No entries!")
	} else {
		for k, v := range book {
			fmt.Printf("%v: %v\n", k, v)
		}
	}
}

func main() {
	fmt.Println("Welcome to a simple phonebook!")
	var phonebook = make(map[string]string)
	var response string
	var number, name, lookup string

	for {
		fmt.Println("What would you like to do?")
		fmt.Println("> [l]ist")
		fmt.Println("> [f]ind")
		fmt.Println("> [a]dd/update")
		fmt.Println("> [q]uit")
		fmt.Scanln(&response)

		if response == "l" {
			list(phonebook)
		} else if response == "f" {
			fmt.Print("Enter a name to find: ")
			fmt.Scanln(&lookup)
			if _, exist := phonebook[lookup]; exist {
				fmt.Printf("Found %v's number %v\n", lookup, phonebook[lookup])
			} else {
				fmt.Printf("Sorry, %v is not an entry.\n", lookup)
				fmt.Print("Do you want to add them? (y/n) ")
				fmt.Scan(&response)
				if response == "y" {
					fmt.Print("Enter content for entry: ")
					fmt.Scan(&number)
					phonebook[lookup] = number
				}
			}
		} else if response == "a" {
			fmt.Println("You can enter a previous entry to update it.")
			fmt.Print("Enter a name for the entry: ")
			fmt.Scanln(&name)
			fmt.Print("Enter content for the entry: ")
			fmt.Scanln(&number)
			if _, exist := phonebook[name]; exist {
				fmt.Println("Updating", name, "to:", number)
				phonebook[name] = number
			} else {
				fmt.Println("Adding", name, number)
				phonebook[name] = number
			}

		} else if response == "q" {
			break
		} else {
			fmt.Println("Command not recognized.")
		}
	}
	fmt.Println("Bye!")
}
