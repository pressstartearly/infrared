package mac

import "github.com/haveachin/infrared/protocol"

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
