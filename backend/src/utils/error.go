// Copyright 2022 Arrikto Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"fmt"

	api "github.com/feast-dev/feast/backend/api/go_client"

	"github.com/golang/glog"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserError struct {
	// Error for internal debugging.
	internalError error
	// Error message for the external client.
	externalMessage string
	// Status code for the external client.
	externalStatusCode codes.Code
}

func newUserError(internalError error, externalMessage string,
	externalStatusCode codes.Code) *UserError {
	return &UserError{
		internalError:      internalError,
		externalMessage:    externalMessage,
		externalStatusCode: externalStatusCode,
	}
}

func NewInternalServerError(err error, internalMessageFormat string,
	a ...interface{}) *UserError {
	internalMessage := fmt.Sprintf(internalMessageFormat, a...)
	return newUserError(
		errors.Wrapf(err, fmt.Sprintf("InternalServerError: %v", internalMessage)),
		"Internal Server Error",
		codes.Internal)
}

func NewResourceNotFoundError(resourceType string, resourceName string) *UserError {
	externalMessage := fmt.Sprintf("%s %s not found.", resourceType, resourceName)
	return newUserError(
		errors.New(fmt.Sprintf("ResourceNotFoundError: %v", externalMessage)),
		externalMessage,
		codes.NotFound)
}

func NewInvalidInputError(messageFormat string, a ...interface{}) *UserError {
	message := fmt.Sprintf(messageFormat, a...)
	return newUserError(errors.Errorf("InvalidInputError: %v", message), message, codes.InvalidArgument)
}

func NewAlreadyExistError(messageFormat string, a ...interface{}) *UserError {
	message := fmt.Sprintf(messageFormat, a...)
	return newUserError(errors.Errorf("AlreadyExistError: %v", message), message, codes.AlreadyExists)
}

func NewBadRequestError(err error, externalFormat string, a ...interface{}) *UserError {
	externalMessage := fmt.Sprintf(externalFormat, a...)
	return newUserError(
		errors.Wrapf(err, fmt.Sprintf("BadRequestError: %v", externalMessage)),
		externalMessage,
		codes.Aborted)
}

func NewUnauthenticatedError(err error, externalFormat string, a ...interface{}) *UserError {
	externalMessage := fmt.Sprintf(externalFormat, a...)
	return newUserError(
		errors.Wrapf(err, fmt.Sprintf("Unauthenticated: %v", externalMessage)),
		externalMessage,
		codes.Unauthenticated)
}

func NewPermissionDeniedError(err error, externalFormat string, a ...interface{}) *UserError {
	externalMessage := fmt.Sprintf(externalFormat, a...)
	return newUserError(
		errors.Wrapf(err, fmt.Sprintf("PermissionDenied: %v", externalMessage)),
		externalMessage,
		codes.PermissionDenied)
}

func (e *UserError) ExternalMessage() string {
	return e.externalMessage
}

func (e *UserError) ExternalStatusCode() codes.Code {
	return e.externalStatusCode
}

func (e *UserError) Error() string {
	return e.internalError.Error()
}

func (e *UserError) Cause() error {
	return e.internalError
}

func (e *UserError) String() string {
	return fmt.Sprintf("%v (code: %v): %+v", e.externalMessage, e.externalStatusCode,
		e.internalError)
}

func (e *UserError) wrapf(format string, args ...interface{}) *UserError {
	return newUserError(errors.Wrapf(e.internalError, format, args...),
		e.externalMessage, e.externalStatusCode)
}

func (e *UserError) wrap(message string) *UserError {
	return newUserError(errors.Wrap(e.internalError, message),
		e.externalMessage, e.externalStatusCode)
}

func (e *UserError) Log() {
	switch e.externalStatusCode {
	case codes.Aborted, codes.InvalidArgument, codes.NotFound, codes.Internal:
		glog.Infof("%+v", e.internalError)
	default:
		glog.Errorf("%+v", e.internalError)
	}
}

func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	switch err := err.(type) {
	case *UserError:
		return err.wrapf(format, args...)
	default:
		return errors.Wrapf(err, format, args...)
	}
}

func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}

	switch err := err.(type) {
	case *UserError:
		return err.wrap(message)
	default:
		return errors.Wrapf(err, message)
	}
}

func LogError(err error) {
	switch err := err.(type) {
	case *UserError:
		err.Log()
	default:
		// We log all the details.
		glog.Errorf("InternalError: %+v", err)
	}
}

func ToGRPCError(err error) error {
	switch err := err.(type) {
	case *UserError:
		userError := err
		stat := status.New(userError.externalStatusCode, userError.internalError.Error())
		statWithDetail, statErr := stat.
			WithDetails(&api.Error{
				ErrorMessage: userError.externalMessage,
				ErrorDetails: userError.internalError.Error(),
			})

		if statErr != nil {
			// Failed to stream error message as proto.
			glog.Errorf("Failed to stream gRPC error. Error to be streamed: %v Error: %v",
				userError.String(), statErr)
			return stat.Err()
		}
		return statWithDetail.Err()
	default:
		externalMessage := fmt.Sprintf("Internal error: %+v", err)
		stat := status.New(codes.Internal, externalMessage)
		statWithDetail, statErr := stat.
			WithDetails(&api.Error{ErrorMessage: externalMessage, ErrorDetails: externalMessage})
		if statErr != nil {
			// Failed to stream error message as proto.
			glog.Errorf("Failed to stream gRPC error. Error to be streamed: %v Error: %v",
				externalMessage, statErr)
			return stat.Err()
		}
		return statWithDetail.Err()
	}
}

// TerminateIfError Check if error is nil. Terminate if not.
func TerminateIfError(err error) {
	if err != nil {
		glog.Fatalf("%v", err)
	}
}