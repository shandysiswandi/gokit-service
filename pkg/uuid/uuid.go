package uuid

import lib "github.com/google/uuid"

func IsValidUUID(s string) bool {
	_, err := lib.Parse(s)
	return err == nil
}
