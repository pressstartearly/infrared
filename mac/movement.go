package mac

import (
	"github.com/haveachin/infrared/protocol"
	"math"
	"time"
)

func movementCheck(player *movementPacket, onGround *float64, lastMovementPk *movementPacket) string {
	if lastMovementPk.x != -1.1821173054061957e+116 { // make sure this isn't the original spawn packets
		// Calculates deltas
		deltaY := (player.y - lastMovementPk.y) / protocol.Double(time.Now().UnixMilli()-t.UnixMilli())
		deltaX := math.Abs(float64(player.x-lastMovementPk.x)) / float64(time.Now().UnixMilli()-t.UnixMilli())
		deltaZ := math.Abs(float64(player.z-lastMovementPk.z)) / float64(time.Now().UnixMilli()-t.UnixMilli())

		//Does ground check calculations
		groundCheck(bool(player.ground), onGround)

		// Checks for illegal vertical movements
		if *onGround <= 0 && deltaY > 0 && !isFlying {
			cheat = "Detected vertical movement cheat."
		}

		// Checks for illegal horizontal movement
		if deltaX+deltaZ > 0.025 && deltaY >= 0.0 && !isFlying && !bool(player.ground) && *onGround <= 0 {
			cheat = "Detected horizontal movement cheat."
		}

		//log.Println("Position", deltaX, deltaY, deltaZ, deltaX+deltaZ, !isFlying, !bool(player.ground), !bool(lastMovementPk.ground))
	}
	*lastMovementPk = *player
	return cheat
}

func groundCheck(ground bool, onGround *float64) {
	if ground && *onGround <= 1000.0 {
		*onGround = 1000.0
		isFlying = false
	} else if *onGround <= 1000.0 || *onGround >= 0 {
		*onGround = *onGround - (float64(time.Now().UnixMilli() - t.UnixMilli()))
	}
	t = time.Now()
}

//
//func flyCheck(pk *protocol.Packet, ground *float32) string {
//
//}
