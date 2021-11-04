package service

type Element string

const (
    # TODO: スピードと時間を変えられるようにしたい
    GO_FORWARD  Element = "await sphero.rollTime(250, 0, 2000, []);"
    GO_BACKWARD Element = "await sphero.rollTime(250, 180, 2000, []);"
    GO_RIGHT    Element = "await sphero.rollTime(250, 90, 2000, []);"
    GO_LEFT     Element = "await sphero.rollTime(250, 270, 2000, []);"
)
