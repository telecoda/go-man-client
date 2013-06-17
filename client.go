package main

import (
	"encoding/json"
	"fmt"
	"github.com/telecoda/go-man/models"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {

	fmt.Println("go-man-client starting")

	start := time.Now()
	//response, err := http.Get("http://localhost:8080/games")

	board, err := startNewGame()

	if err != nil {
		fmt.Println(err)
		return
	}
	printBoard(board)

	var gameId = board.Id

	totalFrames := 1000

	for i := 0; i < totalFrames; i++ {
		fmt.Println("GameId:", gameId)
		fmt.Println("PlayerId:", board.MainPlayer.Id)
		fmt.Println("Frames displayed:", i)
		board, err = getGame(gameId)

		if err != nil {
			fmt.Println(err)
			return
		}
		printBoard(board)

	}
	duration := time.Now().Sub(start)

	fmt.Println(duration)

	fps := (float64(totalFrames) / duration.Seconds())

	fmt.Println("FPS:", fps)

}

func startNewGame() (*models.GameBoard, error) {

	response, err := http.Post("http://localhost:8080/games", "application/json", nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// get body
	jsonBody := make([]byte, 2000)

	count, err := response.Body.Read(jsonBody)

	response.Body.Close()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// reslice to actual bytes read
	jsonBody = jsonBody[:count]

	return convertJsonToBoard(jsonBody)
}

func getGame(gameId string) (*models.GameBoard, error) {

	response, err := http.Get("http://localhost:8080/games/" + gameId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// get body
	jsonBody := make([]byte, 2000)

	count, err := response.Body.Read(jsonBody)

	response.Body.Close()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// reslice to actual bytes read
	jsonBody = jsonBody[:count]

	return convertJsonToBoard(jsonBody)
}

func convertJsonToBoard(jsonBody []byte) (*models.GameBoard, error) {

	var board models.GameBoard

	err := json.Unmarshal(jsonBody, &board)

	if err != nil {
		fmt.Println("Error unmarshalling json", err)
		return nil, err
	}

	return &board, nil

}

func printBoard(board *models.GameBoard) {

	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	fmt.Println("Id:", board.Id)
	fmt.Println("Name:", board.Name)
	fmt.Println("Player:", board.MainPlayer.Location)

	// upload board with players location
	board.BoardCells[board.MainPlayer.Location.Y][board.MainPlayer.Location.X] = byte('M')

	for _, row := range board.BoardCells {
		//for _, cell := range row {
		//fmt.Print(string(cell))
		//}
		fmt.Println(string(row))
	}

}
