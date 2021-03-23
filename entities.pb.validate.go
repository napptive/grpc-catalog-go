// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: catalog/entities.proto

package grpc_catalog_go

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _entities_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on FileInfo with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *FileInfo) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetPath()) < 1 {
		return FileInfoValidationError{
			field:  "Path",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Data

	return nil
}

// FileInfoValidationError is the validation error returned by
// FileInfo.Validate if the designated constraints aren't met.
type FileInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FileInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FileInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FileInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FileInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FileInfoValidationError) ErrorName() string { return "FileInfoValidationError" }

// Error satisfies the builtin error interface
func (e FileInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFileInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FileInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FileInfoValidationError{}

// Validate checks the field values on AddApplicationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AddApplicationRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetApplicationName()) < 1 {
		return AddApplicationRequestValidationError{
			field:  "ApplicationName",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetFile()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddApplicationRequestValidationError{
				field:  "File",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// AddApplicationRequestValidationError is the validation error returned by
// AddApplicationRequest.Validate if the designated constraints aren't met.
type AddApplicationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddApplicationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddApplicationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddApplicationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddApplicationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddApplicationRequestValidationError) ErrorName() string {
	return "AddApplicationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddApplicationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddApplicationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddApplicationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddApplicationRequestValidationError{}

// Validate checks the field values on DownloadApplicationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DownloadApplicationRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetApplicationName()) < 1 {
		return DownloadApplicationRequestValidationError{
			field:  "ApplicationName",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// DownloadApplicationRequestValidationError is the validation error returned
// by DownloadApplicationRequest.Validate if the designated constraints aren't met.
type DownloadApplicationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownloadApplicationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownloadApplicationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownloadApplicationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownloadApplicationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownloadApplicationRequestValidationError) ErrorName() string {
	return "DownloadApplicationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DownloadApplicationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownloadApplicationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownloadApplicationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownloadApplicationRequestValidationError{}
