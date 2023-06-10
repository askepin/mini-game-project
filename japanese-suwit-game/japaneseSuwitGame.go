package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func determineWinner(playerChoice, computerChoice string) string {
	if playerChoice == computerChoice {
		return "Seri"
	} else if (playerChoice == "batu" && computerChoice == "gunting") ||
		(playerChoice == "gunting" && computerChoice == "kertas") ||
		(playerChoice == "kertas" && computerChoice == "batu") {
		return "Anda Menang !"
	} else {
		return "Computer Menang !"
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	choices := []string{"batu", "gunting", "kertas"}
	playAgain := true
	playerScore := 0
	computerScore := 0

	for playAgain {
		fmt.Println("Pilih: batu, gunting atau kertas")
		var playerChoice string
		fmt.Print("Pilihan anda: ")
		fmt.Scanln(&playerChoice)
		playerChoice = strings.ToLower(playerChoice)

		validChoice := false
		for _, choice := range choices {
			if playerChoice == choice {
				validChoice = true
				break
			}
		}

		if !validChoice {
			fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
			continue
		}

		computerChoice := choices[rand.Intn(len(choices))]
		computerChoice = strings.ToLower(computerChoice)

		fmt.Println("Anda memilih:", playerChoice)
		fmt.Println("Komputer memilih:", computerChoice)

		result := determineWinner(playerChoice, computerChoice)
		fmt.Println(result)

		if result == "Anda Menang !" {
			playerScore++
		} else if result == "Computer Menang !" {
			computerScore++
		} else {
			fmt.Println("Tidak Ada Penambahan Skor !")
		}

		fmt.Println("Player score: ", playerScore)
		fmt.Println("Computer score: ", computerScore)

		var playAgainInput string
		fmt.Print("Apakah Anda ingin bermain lagi? (ya/tidak)")
		fmt.Scanln(&playAgainInput)
		playAgainInput = strings.ToLower(playAgainInput)
		if playAgainInput != "ya" {
			playAgain = false
		}
	}
}
