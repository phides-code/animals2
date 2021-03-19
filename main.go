/*Write a program which allows the user to create a set of animals and to get information 
about those animals. Each animal has a name and can be either a cow, bird, or snake. With each 
command, the user can either create a new animal of one of the three types, or the user can 
request information about an animal that he/she has already created. Each animal has a unique 
name, defined by the user. Note that the user can define animals of a chosen type, but the 
types of animals are restricted to either cow, bird, or snake. The following table contains 
the three types of animals and their associated data.

Your program should present the user with a prompt, “>”, to indicate that the user can type a 
request. Your program should accept one command at a time from the user, print out a response, 
and print out a new prompt on a new line. Your program should continue in this loop forever. 
Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is 
“newanimal”. The second string is an arbitrary string which will be the name of the new 
animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”.  
Your program should process each newanimal command by creating the new animal and printing 
“Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”. 
The second string is the name of the animal. The third string is the name of the information 
requested about the animal, either “eat”, “move”, or “speak”. Your program should process each 
query command by printing out the requested data.*/

package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

// map of given names to animal types
var namesMap = make(map[string]string)

type Cow struct{ food, locomotion, noise string }
type Bird struct{ food, locomotion, noise string }
type Snake struct{ food, locomotion, noise string }

func (c Cow) Eat()   { fmt.Println(c.food) }
func (c Cow) Move()  { fmt.Println(c.locomotion) }
func (c Cow) Speak() { fmt.Println(c.noise) }

func (b Bird) Eat()   { fmt.Println(b.food) }
func (b Bird) Move()  { fmt.Println(b.locomotion) }
func (b Bird) Speak() { fmt.Println(b.noise) }

func (s Snake) Eat()   { fmt.Println(s.food) }
func (s Snake) Move()  { fmt.Println(s.locomotion) }
func (s Snake) Speak() { fmt.Println(s.noise) }

func newCow(givenName string)   { namesMap[givenName] = "cow" }
func newBird(givenName string)  { namesMap[givenName] = "bird" }
func newSnake(givenName string) { namesMap[givenName] = "snake" }

func main() {
	// slice of 3 strings for the arguments
	args := make([]string, 3)

	// initialize a map of strings/functions for animal actions
	animalActions := initAnimalActions()

	// initialize a map of strings/functions for making new animals
	animalMaker := initAnimalMaker()

	for {
		fmt.Printf("> ")

		fmt.Scan(&args[0]) // command "newanimal", or "query"
		fmt.Scan(&args[1]) // given name
		fmt.Scan(&args[2]) // animal type, or action

		switch {
		case strings.ToLower(args[0]) == "newanimal":
			// call the right animal maker function, pass the given name
			animalMaker[args[2]](args[1])
			fmt.Println("Created it!")

		case strings.ToLower(args[0]) == "query":
			// use the given name to get the animal type from namesMap
			animalType := namesMap[args[1]]
			animalAction := args[2]

			// call the requested action for this type
			animalActions[animalType][animalAction]()
		}
	}
}

func initAnimalActions() map[string]map[string]func() {
	cow := Cow{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Bird{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Snake{food: "mice", locomotion: "slither", noise: "hsss"}

	cowActions := map[string]func(){
		"eat": cow.Eat, "move": cow.Move, "speak": cow.Speak,
	}
	birdActions := map[string]func(){
		"eat": bird.Eat, "move": bird.Move, "speak": bird.Speak,
	}
	snakeActions := map[string]func(){
		"eat": snake.Eat, "move": snake.Move, "speak": snake.Speak,
	}

	return map[string]map[string]func(){
		"cow": cowActions, "bird": birdActions, "snake": snakeActions,
	}
}

func initAnimalMaker() map[string]func(string) {
	return map[string]func(string){
		"cow":   newCow,
		"bird":  newBird,
		"snake": newSnake,
	}
}
