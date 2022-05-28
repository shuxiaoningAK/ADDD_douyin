package util

import uuid "github.com/satori/go.uuid"

func NextUuid() string {
	return uuid.NewV4().String()
}
