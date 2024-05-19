package monopoly

type GameState uint8

const (
	GAME_TURN_START GameState = iota
	GAME_ROLLED_DICE
	GAME_MOVED_TO_NEW_FIELD
	GAME_TURN
)

func (gs GameState) String() string {
	switch gs {
	case GAME_TURN_START:
		return "begin turn"
	case GAME_ROLLED_DICE:
		return "rolled dice"
	case GAME_MOVED_TO_NEW_FIELD:
		return "moved to new field"
	case GAME_TURN:
		return "normal turn"
	default:
		return "UNKNOWN"
	}
}

func (gs GameState) GoString() string {
	switch gs {
	case GAME_TURN_START:
		return "GAME_TURN_START"
	case GAME_ROLLED_DICE:
		return "GAME_ROLLED_DICE"
	case GAME_MOVED_TO_NEW_FIELD:
		return "GAME_MOVED_TO_NEW_FIELD"
	case GAME_TURN:
		return "GAME_TURN"
	default:
		return "UNKNOWN"
	}
}
