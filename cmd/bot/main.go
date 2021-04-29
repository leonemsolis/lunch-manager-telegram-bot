package main

import (
	"bufio"
	"github.com/leonemsolis/qzhub_obed_manager_bot/pkg/telegram"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

func main() {
	lines, err := readLines("keys.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	bot, err := tb.NewBot(tb.Settings{
		Token:       lines[0],
		Poller:      &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	b := telegram.NewBot(bot, lines[1])


	b.Init()
	b.Bot.Start()
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}