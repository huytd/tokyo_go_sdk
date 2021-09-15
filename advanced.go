package tokyo_go_sdk

import (
	"errors"
	"math"
)

func (c *Client) HeadToPoint(x, y float64) error {
	radian := RadianBetweenPoint1ToPoint2(c.CurrentPlayer.X, c.CurrentPlayer.Y, x, y)
	return c.Rotate(radian)
}

func (c *Client) GetFirstOtherPlayer() (Player, error) {
	if len(c.CurrentGameInfo.Players) == 0 {
		return Player{}, errors.New("Don't find anyone, 0 player")
	}

	for _, player := range c.CurrentGameInfo.Players {
		if IsSamePlayers(player, c.CurrentPlayer) || IsPlayersSamePosition(player, c.CurrentPlayer) {
			continue
		}
		if c.IsDeadPlayer(player) {
			continue
		}
		return player, nil
	}

	return Player{}, errors.New("Don't find anyone, only me")
}

func (c *Client) GetClosestPlayer() (Player, float64, error) {
	others := c.GetOtherPlayers()

	shortestDistance := math.MaxFloat64
	var closestPlayer *Player
	for index := range others {
		d := DistanceBetweenTwoPlayers(c.CurrentPlayer, others[index])
		if d < shortestDistance {
			closestPlayer = &others[index]
			shortestDistance = d
		}
	}
	if closestPlayer == nil {
		return Player{}, 0, errors.New("Don't find anyone, 0 player")
	}
	return *closestPlayer, shortestDistance, nil
}

func (c *Client) GetOtherPlayers() Players {
	var others Players
	for _, player := range c.CurrentGameInfo.Players {
		if IsSamePlayers(player, c.CurrentPlayer) || IsPlayersSamePosition(player, c.CurrentPlayer) {
			continue
		}
		if c.IsDeadPlayer(player) {
			continue
		}
		others = append(others, player)
	}
	return others
}

func (c *Client) IsDeadPlayer(player Player) bool {
	for _, aDead := range c.CurrentGameInfo.Dead {
		if aDead.Player.ID == player.ID {
			return true
		}
	}
	return false
}
