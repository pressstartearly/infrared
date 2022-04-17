package mac

import (
	"github.com/haveachin/infrared/protocol"
	"math"
	"time"
)

func movementCheck(player *movementPacket, tracker *PlayerTracker) string {
	if tracker.lastMovementPk.x != -1.1821173054061957e+116 { // make sure this isn't the original spawn packets
		// Calculates deltas
		deltaY := (player.y - tracker.lastMovementPk.y) / protocol.Double(time.Now().UnixMilli()-tracker.t.UnixMilli())
		deltaX := math.Abs(float64(player.x-tracker.lastMovementPk.x)) / float64(time.Now().UnixMilli()-tracker.t.UnixMilli())
		deltaZ := math.Abs(float64(player.z-tracker.lastMovementPk.z)) / float64(time.Now().UnixMilli()-tracker.t.UnixMilli())

		//Does ground check calculations
		groundCheck(bool(player.ground), tracker)

		// Checks for illegal vertical movements
		if tracker.onGround <= 0 && deltaY > 0 && !tracker.isFlying {
			tracker.cheat = "Detected vertical movement cheat."
		}

		// Checks for illegal horizontal movement
		if deltaX+deltaZ > 0.025 && deltaY >= 0.0 && !tracker.isFlying && !bool(player.ground) && tracker.onGround <= 0 {
			tracker.cheat = "Detected horizontal movement cheat."
		}

		//log.Println("Position", deltaX, deltaY, deltaZ, deltaX+deltaZ, !isFlying, !bool(player.ground), !bool(lastMovementPk.ground))
	}
	tracker.lastMovementPk = *player
	return tracker.cheat
}

func groundCheck(ground bool, tracker *PlayerTracker) {
	if ground && tracker.onGround <= 1000.0 {
		tracker.onGround = 1000.0
		tracker.isFlying = false
	} else if tracker.onGround <= 1000.0 || tracker.onGround >= 0 {
		tracker.onGround = tracker.onGround - (float64(time.Now().UnixMilli() - tracker.t.UnixMilli()))
	}
	tracker.t = time.Now()
}

//
//func flyCheck(pk *protocol.Packet, ground *float32) string {
//
//}
