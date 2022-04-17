package mac

import (
	"github.com/haveachin/infrared/protocol"
	"log"
)

func Filter(pk *protocol.Packet, tracker *PlayerTracker) string {
	if pk.ID == 0x0F {
		log.Println("ClientBound Packet Detected")
	}
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
		tracker.cheat = movementCheck(&mvpk, tracker)

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
		tracker.cheat = movementCheck(&mvpk, tracker)

	}

	if pk.ID == 0x13 {

		player := ServerBoundPlayerRotation{}

		err := pk.Scan(&player.yaw, &player.pitch, &player.ground)
		if err != nil {
			return "Well this isn't supposed to happen. ¯\\_(ツ)_/¯"
		}

		groundCheck(bool(player.ground), tracker)

	}

	if pk.ID == 0x14 {

		player := ServerBoundPlayerGround{}

		err := pk.Scan(&player.ground)
		if err != nil {
			return "Well this isn't supposed to happen. ¯\\_(ツ)_/¯"
		}

		groundCheck(bool(player.ground), tracker)
	}

	// set player flying
	if pk.ID == 0x1B {
		action := ServerBoundEntityAction{}

		err := pk.Scan(&action.entityid, &action.actionid, &action.jumpboost)
		if err != nil {
			return "Well this isn't supposed to happen. ¯\\_(ツ)_/¯"
		}

		if action.actionid == 8 {
			tracker.isFlying = !tracker.isFlying
		}
	}
	return tracker.cheat
}
