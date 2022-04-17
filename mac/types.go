package mac

import (
	"github.com/haveachin/infrared/protocol"
	"time"
)

type ServerBoundPlayerPositionRotation struct {
	x      protocol.Double
	y      protocol.Double
	z      protocol.Double
	yaw    protocol.Float
	pitch  protocol.Float
	ground protocol.Boolean
}

type ServerBoundPlayerRotation struct {
	yaw    protocol.Float
	pitch  protocol.Float
	ground protocol.Boolean
}

type ServerBoundPlayerPosition struct {
	x      protocol.Double
	y      protocol.Double
	z      protocol.Double
	ground protocol.Boolean
}

type ServerBoundPlayerGround struct {
	ground protocol.Boolean
}

type ServerBoundEntityAction struct {
	entityid  protocol.VarInt
	actionid  protocol.VarInt
	jumpboost protocol.VarInt
}

type movementPacket struct {
	x      protocol.Double
	y      protocol.Double
	z      protocol.Double
	ground protocol.Boolean
}

type playerTracker struct {
	isFlying       bool
	t              time.Time
	cheat          string
	lastMovementPk movementPacket
	onGround       float64
}

func NewPlayerTracker() playerTracker {
	playerTracker := playerTracker{}
	playerTracker.isFlying = false
	playerTracker.t = time.Now()
	playerTracker.cheat = ""
	playerTracker.lastMovementPk = movementPacket{}
	playerTracker.onGround = 1000.0
	return playerTracker
}
