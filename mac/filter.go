package mac

import (
	"github.com/haveachin/infrared/protocol"
	"time"
)

var isFlying = false
var t = time.Now()
var cheat string
var lastMovementPk movementPacket
var onGround float64 = 1000.0

func Filter(pk *protocol.Packet) string {
	if pk.ID == 0x12 {
		player := ServerBoundPlayerPositionRotation{}

		err := pk.Scan(&player.x, &player.y, &player.z, &player.yaw, &player.pitch, &player.ground)
		if err != nil {
			return "Well this isn't supposed to happen. ¯\\_(ツ)_/¯"
		}

		//make movement packet
		mvpk := movementPacket{
			x:      player.x,
			y:      player.y,
			z:      player.z,
			ground: player.ground,
		}

		// call filter
		cheat = movementCheck(&mvpk, &onGround, &lastMovementPk)

	}

	if pk.ID == 0x11 {
		player := ServerBoundPlayerPosition{}

		err := pk.Scan(&player.x, &player.y, &player.z, &player.ground)
		if err != nil {
			return "Well this isn't supposed to happen. ¯\\_(ツ)_/¯"
		}

		//make movement packet
		mvpk := movementPacket{
			x:      player.x,
			y:      player.y,
			z:      player.z,
			ground: player.ground,
		}

		// call filter
		cheat = movementCheck(&mvpk, &onGround, &lastMovementPk)

	}

	if pk.ID == 0x13 {

		player := ServerBoundPlayerRotation{}

		err := pk.Scan(&player.yaw, &player.pitch, &player.ground)
		if err != nil {
			return "Well this isn't supposed to happen. ¯\\_(ツ)_/¯"
		}

		groundCheck(bool(player.ground), &onGround)

	}

	if pk.ID == 0x14 {

		player := ServerBoundPlayerGround{}

		err := pk.Scan(&player.ground)
		if err != nil {
			return "Well this isn't supposed to happen. ¯\\_(ツ)_/¯"
		}

		groundCheck(bool(player.ground), &onGround)
	}

	// set player flying
	if pk.ID == 0x1B {
		action := ServerBoundEntityAction{}

		err := pk.Scan(&action.entityid, &action.actionid, &action.jumpboost)
		if err != nil {
			return "Well this isn't supposed to happen. ¯\\_(ツ)_/¯"
		}

		if action.actionid == 8 {
			isFlying = !isFlying
		}
	}
	return cheat
}
