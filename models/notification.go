package models

//Notification struct of json that will send to NotifyTrack
type Notification struct {
	CpnID     int       `json:"cpnId"`
	TrackType string    `json:"trackType"`
	TrackData TrackData `json:"trackData"`
}

//TrackData struct of message to publish in NotifyTrack
type TrackData struct {
	Msg string `json:"msg"`
}
