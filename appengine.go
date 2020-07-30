package aefire

import (
	"fmt"
	"os"
)

func GAEAppID() string {
	aid := os.Getenv("GOOGLE_CLOUD_PROJECT")

	if aid == "" {
		aid = os.Getenv("GCP_PROJECT")
	}

	return aid
}

func GAEServiceID() string {
	return os.Getenv("GAE_SERVICE")
}

func EndPoint() string {
	return ServiceEndPoint(GAEServiceID())
}

func ServiceEndPoint(serviceID string) string {
	if serviceID == "default" || serviceID == "" {
		return fmt.Sprintf("https://%s.appspot.com",
			GAEAppID())
	} else {
		return fmt.Sprintf("https://%s-dot-%s.appspot.com",
			serviceID,
			GAEAppID())
	}
}
