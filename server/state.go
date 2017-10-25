package server

type Status struct{
	LeftWindowUp bool `json:"leftWindowUp"`
	LeftWindowDown bool `json:"leftWindowDown"`
	RightWindowUp bool `json:"rightWindowUp"`
	RightWindowDown bool `json:"rightWindowDown"`
	LeftDoorLocked bool `json:"leftDoorLocked"`
	RightDoorLocket bool `json:"rightDoorLocked"`
	TrunkOpen bool `json:"trunkOpen"`
	HazardOn bool `json:"hazardOn"`
	PopupOn bool `json:"popupOn"`
	HornOn bool `json:"hornOn"`
	HornMode string `json:"hornMode"`
	McACOn bool `json:"mcACOn"`
	CarACOn bool `json:"carACOn"`
	McFanOn bool `json:"mcFanOn"`
	CarFanOn bool `json:"carFanOn"`
	McAccessoryOn bool `json:"mcAccessoryOn"`
	CarAccessoryOn bool `json:"carAccessoryOn"`
	McIG1On bool `json:"mcIG1On"`
	CarIG1On bool `json:"carIG1On"`
	McIG2On bool `json:"mcIG2On"`
	CarIG2On bool `json:"carIG2On"`
	EngineRunning bool `json:"engineRunning"`
	BluetoothConnected bool `json:"bluetoothConnected"`
}
