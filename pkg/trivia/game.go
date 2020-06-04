package trivia

// Game represents a single instance of a trivia game
type Game struct {
	Players []*Player
}

// NewGame creates and returns a new Game
func NewGame() *Game {
	g := Game{}
	return &g
}

// AddPlayer adds a new player to the game
func (g *Game) AddPlayer(player *Player) {
	g.Players = append(g.Players, player)
}

// ContainsPlayer verifies a Player is already created in the Game
func (g *Game) ContainsPlayer(username string) bool {
	for _, player := range g.Players {
		if player.Username == username {
			return true
		}
	}
	return false
}

// GetPlayer returns the Player from the Game
func (g *Game) GetPlayer(username string) *Player {
	for _, player := range g.Players {
		if player.Username == username {
			return player
		}
	}
	return nil
}

// Player represents an instance of a player in a trivia game
type Player struct {
	Username       string
	CorrectAnswers int16
	Score          int32
}

// NewPlayer creates and returns a new Player instance
func NewPlayer(username string) *Player {
	p := Player{
		Username: username,
	}
	return &p
}

// AddScore adds points to a Player's score
func (p *Player) AddScore(score int32) {
	p.Score += score
}

// AddCorrectAnswer increments the Player's correct answer
func (p *Player) AddCorrectAnswer() {
	p.CorrectAnswers++
}
