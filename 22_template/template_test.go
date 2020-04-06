package template

import "testing"

func Test_template(t *testing.T) {
	tennis := NewTennis()
	tennis.play()
	//output
	//Tennis Game Initialized! Start playing.
	//Tennis Game Started. Enjoy the game!
	//Tennis Game Finished!

	football := NewFootball()
	football.play()
	//output
	//Football Game Initialized! Start playing.
	//Football Game Started. Enjoy the game!
	//Football Game Finished!

}
