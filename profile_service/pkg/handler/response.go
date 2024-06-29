package handler

import (
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/sirupsen/logrus"
)

func newErrorResponse(message string, err error) error {
	logrus.Error(err)
	return gqlerrors.NewFormattedError(message)
}
