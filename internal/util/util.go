package util

import (
	"log/slog"
	"os/user"
)

const VoidString = ""

func IsVoidString(s string) bool {
	return s == VoidString
}

func GetEffectiveUserHomeFolder() string {
	currentUser, err := user.Current()
	if err != nil {
		slog.Error("Unable to locate HOME dir for effective user: %s\n", err)
		return VoidString
	}

	return currentUser.Name
}
