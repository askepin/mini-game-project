package main

import (
	"fmt"
	"io/ioutil"
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
		return "Anda Menang!"
	} else {
		return "Computer Menang!"
	}
}

func playGame(totalRounds int) {
	rand.Seed(time.Now().UnixNano())

	choices := []string{"batu", "gunting", "kertas"}
	playerScore := 0
	computerScore := 0

	for round := 1; round <= totalRounds; round++ {
		fmt.Println("Round", round)
		fmt.Println("Pilih: batu, gunting, atau kertas")
		var playerChoice string
		fmt.Print("Pilihan Anda: ")
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
			round--
			continue
		}

		computerChoice := choices[rand.Intn(len(choices))]
		computerChoice = strings.ToLower(computerChoice)

		fmt.Println("Anda memilih:", playerChoice)
		fmt.Println("Komputer memilih:", computerChoice)

		result := determineWinner(playerChoice, computerChoice)
		fmt.Println(result)

		if result == "Anda Menang!" {
			playerScore++
		} else if result == "Computer Menang!" {
			computerScore++
		} else {
			fmt.Println("Ronde berakhir seri! Ulangi ronde yang sama.")
			round-- // Mengulangi ronde yang sama
		}

		fmt.Println("Skor Player:", playerScore)
		fmt.Println("Skor Komputer:", computerScore)
		fmt.Println()

		// Periksa jika permainan selesai
		if playerScore > totalRounds/2 || computerScore > totalRounds/2 {
			break
		}
	}

	fmt.Println("Permainan selesai!")

	if playerScore > computerScore {
		fmt.Println("Anda memenangkan permainan!")
	} else if playerScore < computerScore {
		fmt.Println("Komputer memenangkan permainan!")
	} else {
		fmt.Println("Permainan berakhir seri!")
	}

	// Simpan log data permainan ke dalam file
	saveGameLog(playerScore, computerScore, totalRounds)

	// Pilihan bermain kembali atau tidak
	fmt.Print("Apakah Anda ingin bermain kembali? (ya/tidak): ")
	var playAgainInput string
	fmt.Scanln(&playAgainInput)
	playAgainInput = strings.ToLower(playAgainInput)
	if playAgainInput == "ya" {
		fmt.Println()
		main() // Panggil kembali fungsi main untuk memulai permainan baru
	} else {
		fmt.Println("Terima kasih telah bermain!")
	}
}

func saveGameLog(playerScore, computerScore, totalRounds int) {
	logData := fmt.Sprintf("Skor Player: %d\nSkor Komputer: %d\nTotal Ronde: %d\n", playerScore, computerScore, totalRounds)
	err := ioutil.WriteFile("game_log.txt", []byte(logData), 0644)
	if err != nil {
		fmt.Println("Gagal menyimpan log data permainan:", err)
		return
	}

	fmt.Println("Log data permainan telah disimpan dalam file game_log.txt")
}

func main() {
	var totalRounds int
	fmt.Println("Pilih tipe permainan:")
	fmt.Println("1. Best of Three")
	fmt.Println("2. Best of Five")
	fmt.Println("3. Best of Seven")
	fmt.Print("Pilihan Anda: ")
	fmt.Scanln(&totalRounds)

	switch totalRounds {
	case 1:
		playGame(3)
	case 2:
		playGame(5)
	case 3:
		playGame(7)
	default:
		fmt.Println("Pilihan tidak valid. Silakan jalankan program kembali dan pilih angka antara 1-3.")
	}
}
