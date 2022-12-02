package utils

import (
	"errors"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var ErrLayoutNotValid = errors.New("layout not valid")

func ParseTime(layout string, val string) (*time.Time, error) {
	var format string
	switch layout {
	case "HH:MM:SS DD/MM/YYYY":
		format = "15:04:05 02/01/2006"
	}

	if len(format) == 0 {
		return nil, ErrLayoutNotValid
	}

	result, err := time.Parse(format, val)
	if err != nil {
		return nil, err
	}

	return &result, nil

}

func ParseTimeToGrpc(t *time.Time) *timestamp.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}
