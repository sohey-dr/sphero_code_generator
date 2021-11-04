package service

type Element string

const (
    GO_FORWARD  Element = "await sphero.rollTime(SPEED, 0, PATROL_TIME, []);"
    GO_BACKWARD Element = "await sphero.rollTime(SPEED, 180, PATROL_TIME, []);"
)
