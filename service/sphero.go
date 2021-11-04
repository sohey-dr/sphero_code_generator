package service

type Element string

const (
    GO_FORWARD  Element = "await toy.rollTime(SPEED, 0, PATROL_TIME, []);"
    GO_BACKWARD Element = "await toy.rollTime(SPEED, 180, PATROL_TIME, []);"
)
