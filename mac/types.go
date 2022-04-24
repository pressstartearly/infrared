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

type ClientBoundEntityVelocity struct {
	entityID protocol.VarInt
	vx       protocol.Short
	vy       protocol.Short
	vz       protocol.Short
}

type ClientBoundEntityMetadata struct {
	entityID  protocol.VarInt
	index     protocol.UnsignedByte
	valueType protocol.VarInt
	value     protocol.VarInt
}

type ClientBoundParticle struct {
	particleID   protocol.Int
	longDistance protocol.Boolean
	x            protocol.Double
	y            protocol.Double
	z            protocol.Double
}

type ClientBoundJoinGame struct {
	entityID protocol.Int
}

type movementPacket struct {
	x      protocol.Double
	y      protocol.Double
	z      protocol.Double
	ground protocol.Boolean
}

type PlayerTracker struct {
	isFlying       bool
	pose           int
	t              time.Time
	cheat          string
	lastMovementPk movementPacket
	onGround       float64
	entityID       protocol.Int
}

func NewPlayerTracker() PlayerTracker {
	playerTracker := PlayerTracker{}
	playerTracker.isFlying = false
	playerTracker.t = time.Now()
	playerTracker.cheat = ""
	playerTracker.lastMovementPk = movementPacket{}
	playerTracker.onGround = 1000.0
	return playerTracker
}
