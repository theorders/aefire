package aefire

import (
	"strings"
	"time"
)

// var uid: String? = fauth.uid
//     var iid: String = FirebaseInstanceId.getInstance().id
//     var token: String? = null
//     var phoneNumber: String? = null
//     var createdAt = Timestamp.now()
//     var locationUpdatedAt: Timestamp? = null
//     var location:GeoPoint? = null
//     val trackers = arrayListOf<String>()
//     var gpsState = GpsState.turnedOff
//     var model: Model? = null

type Device struct {
	UID         *string    `json:"uid,omitempty" firestore:"uid,omitempty"`
	IID         string     `json:"iid" firestore:"iid"`
	Token       *string    `json:"token,omitempty" firestore:"token,omitempty"`
	PhoneNumber *string    `json:"phoneNumber,omitempty" firestore:"phoneNumber,omitempty"`
	CreatedAt   *time.Time `json:"createdAt" firestore:"createdAt"`
	UserAgent   string     `json:"userAgent" firestore:"userAgent"`
}

func (d *Device) DocId() string {
	return d.IID
}

type FcmToken string

func (t *FcmToken) IID() string {
	s := string(*t)

	return s[:strings.Index(s, ":")]
}
