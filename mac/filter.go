package mac

import (
	"github.com/haveachin/infrared/protocol"
	"log"
)

func Filter(pk *protocol.Packet, tracker *PlayerTracker) string {

	//Client bound log in packet.
	//Using this to get the player's entity id
	//PROB DON'T NEED THIS
	if pk.ID == 0x26 {
		packet := ClientBoundJoinGame{}
		err := pk.Scan(&packet.entityID)
		if err != nil {
			return "Well this isn't supposed to happen. ¯\\_(ツ)_/¯"
		}
		tracker.entityID = packet.entityID
		//log.Println("Player Entity ID: ", tracker.entityID)
	}

	if pk.ID == 0x4D {
		packet := ClientBoundEntityMetadata{}

		err := pk.Scan(&packet.entityID, &packet.index, &packet.valueType, &packet.value)

		if err != nil {
			return "Well this isn't supposed to happen. ¯\\_(ツ)_/¯"
		}

		if packet.index != 0xff && int(packet.entityID) == int(tracker.entityID) {

			if int(packet.valueType) == 18 {
				tracker.pose = int(packet.value)
				log.Println("Set pose to: ", int(packet.value))
			}

			log.Println(int(packet.entityID), int(tracker.entityID), int(packet.valueType), int(packet.value))
		}
		//log.Println(packet.entityID, packet.index, packet.valueType, packet.value)
	}

	if pk.ID == 0x24 {
		packet := ClientBoundParticle{}

		err := pk.Scan(&packet.particleID, &packet.longDistance, &packet.x, &packet.y, &packet.z)
		if err != nil {
			return "Well this isn't supposed to happen. ¯\\_(ツ)_/¯"
		}

		log.Println(packet)

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
