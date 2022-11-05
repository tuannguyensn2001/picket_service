package googlepkg

import (
	"log"
	"testing"
)

func TestGetCode(t *testing.T) {
	result, err := GetGoogleOauthToken("4/0AfgeXvve8oc_gmLkIvlvNXapLnjyGMgePIyO8ATUcte-o0WdX3x5oP78cl2E7FAKbJz3Fg")
	log.Print(result, err)
}
