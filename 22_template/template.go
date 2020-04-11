package template

import "fmt"

type Game interface {
	initialize()
	startPlay()
	endPlay()
}

type GameTemplate struct {
	Game
}

func (gt *GameTemplate) play() {
	gt.initialize()
	gt.startPlay()
	gt.endPlay()
}

type Tennis struct {
	GameTemplate
}

func (c *Tennis) initialize() {
	fmt.Println("Tennis Game Initialized! Start playing.")
}

func (c *Tennis) startPlay() {
	fmt.Println("Tennis Game Started. Enjoy the game!")
}

func (c *Tennis) endPlay() {
	fmt.Println("Tennis Game Finished!")
}
func NewTennis() *Tennis {
	game := &Tennis{}
	game.GameTemplate = GameTemplate{game}
	return game
}

type Football struct {
	GameTemplate
}

func NewFootball() *Football {
	game := &Football{}
	game.GameTemplate = GameTemplate{game}
	return game
}

func (c *Football) initialize() {
	fmt.Println("Football Game Initialized! Start playing.")
}

func (c *Football) startPlay() {
	fmt.Println("Football Game Started. Enjoy the game!")
}

func (c *Football) endPlay() {
	fmt.Println("Football Game Finished!")
}
