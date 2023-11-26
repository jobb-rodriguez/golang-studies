package main

import "fmt"

func computeCostPerMessage(message string) float64 {
	return float64(len(message)) * 0.01
}

func getMessageCosts(messages []string) []float64 {
	messageLength := len(messages)
	messageCosts := make([]float64, messageLength)
	for i := 0; i < messageLength; i++ {
		messageCosts[i] = computeCostPerMessage(messages[i])
	}
	return messageCosts
}

// don'te edit below this line

func test(messages []string) {
	costs := getMessageCosts(messages)
	fmt.Println("Message:")
	for i := 0; i < len(messages); i++ {
		fmt.Printf(" - %v\n", messages[i])
	}
	fmt.Println("Costs:")
	for i := 0; i < len(costs); i++ {
		fmt.Printf(" - %.2f\n", costs[i])
	}
	fmt.Println("==== END REPORT ====")
}

func main() {
	test([]string{
		"Welcome to the movies!",
		"Enjoy your popcorn!",
		"Please don't talk during the movie!",
	})
	test([]string{
		"I don't want to be here anymore",
		"Can we go home?",
		"I'm hungry",
		"I'm bored",
	})
	test([]string{
		"Hello",
		"Hi",
		"Hey",
		"Hi there",
		"Hey there",
		"Hi there",
		"Hello there",
		"Hey there",
		"Hello there",
		"General Kenobi",
	})
}
