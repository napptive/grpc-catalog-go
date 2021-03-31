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

// Validate checks the field values on RemoveApplicationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveApplicationRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetApplicationName()) < 1 {
		return RemoveApplicationRequestValidationError{
			field:  "ApplicationName",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// RemoveApplicationRequestValidationError is the validation error returned by
// RemoveApplicationRequest.Validate if the designated constraints aren't met.
type RemoveApplicationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveApplicationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveApplicationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveApplicationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveApplicationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveApplicationRequestValidationError) ErrorName() string {
	return "RemoveApplicationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveApplicationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveApplicationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveApplicationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveApplicationRequestValidationError{}

// Validate checks the field values on InfoApplicationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *InfoApplicationRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetApplicationName()) < 1 {
		return InfoApplicationRequestValidationError{
			field:  "ApplicationName",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// InfoApplicationRequestValidationError is the validation error returned by
// InfoApplicationRequest.Validate if the designated constraints aren't met.
type InfoApplicationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InfoApplicationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InfoApplicationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InfoApplicationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InfoApplicationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InfoApplicationRequestValidationError) ErrorName() string {
	return "InfoApplicationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e InfoApplicationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInfoApplicationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InfoApplicationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InfoApplicationRequestValidationError{}

// Validate checks the field values on InfoApplicationResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *InfoApplicationResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RepositoryName

	// no validation rules for ApplicationName

	// no validation rules for Tag

	// no validation rules for MetadataFile

	// no validation rules for ReadmeFile

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InfoApplicationResponseValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// InfoApplicationResponseValidationError is the validation error returned by
// InfoApplicationResponse.Validate if the designated constraints aren't met.
type InfoApplicationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InfoApplicationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InfoApplicationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InfoApplicationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InfoApplicationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InfoApplicationResponseValidationError) ErrorName() string {
	return "InfoApplicationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e InfoApplicationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInfoApplicationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InfoApplicationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InfoApplicationResponseValidationError{}

// Validate checks the field values on ApplicationSummary with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ApplicationSummary) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RepositoryName

	// no validation rules for ApplicationName

	// no validation rules for Tag

	// no validation rules for MetadataName

	return nil
}

// ApplicationSummaryValidationError is the validation error returned by
// ApplicationSummary.Validate if the designated constraints aren't met.
type ApplicationSummaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationSummaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationSummaryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationSummaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationSummaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationSummaryValidationError) ErrorName() string {
	return "ApplicationSummaryValidationError"
}

// Error satisfies the builtin error interface
func (e ApplicationSummaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationSummary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationSummaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationSummaryValidationError{}

// Validate checks the field values on ApplicationList with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ApplicationList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetApplications() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ApplicationListValidationError{
					field:  fmt.Sprintf("Applications[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ApplicationListValidationError is the validation error returned by
// ApplicationList.Validate if the designated constraints aren't met.
type ApplicationListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationListValidationError) ErrorName() string { return "ApplicationListValidationError" }

// Error satisfies the builtin error interface
func (e ApplicationListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationListValidationError{}

// Validate checks the field values on KubernetesEntities with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *KubernetesEntities) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ApiVersion

	// no validation rules for Type

	// no validation rules for Name

	return nil
}

// KubernetesEntitiesValidationError is the validation error returned by
// KubernetesEntities.Validate if the designated constraints aren't met.
type KubernetesEntitiesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KubernetesEntitiesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KubernetesEntitiesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KubernetesEntitiesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KubernetesEntitiesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KubernetesEntitiesValidationError) ErrorName() string {
	return "KubernetesEntitiesValidationError"
}

// Error satisfies the builtin error interface
func (e KubernetesEntitiesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKubernetesEntities.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KubernetesEntitiesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KubernetesEntitiesValidationError{}

// Validate checks the field values on CatalogRequirement with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CatalogRequirement) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetK8S() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CatalogRequirementValidationError{
					field:  fmt.Sprintf("K8S[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CatalogRequirementValidationError is the validation error returned by
// CatalogRequirement.Validate if the designated constraints aren't met.
type CatalogRequirementValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CatalogRequirementValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CatalogRequirementValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CatalogRequirementValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CatalogRequirementValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CatalogRequirementValidationError) ErrorName() string {
	return "CatalogRequirementValidationError"
}

// Error satisfies the builtin error interface
func (e CatalogRequirementValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCatalogRequirement.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CatalogRequirementValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CatalogRequirementValidationError{}

// Validate checks the field values on ApplicationLogo with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ApplicationLogo) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Src

	// no validation rules for Type

	// no validation rules for Size

	return nil
}

// ApplicationLogoValidationError is the validation error returned by
// ApplicationLogo.Validate if the designated constraints aren't met.
type ApplicationLogoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationLogoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationLogoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationLogoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationLogoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationLogoValidationError) ErrorName() string { return "ApplicationLogoValidationError" }

// Error satisfies the builtin error interface
func (e ApplicationLogoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationLogo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationLogoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationLogoValidationError{}

// Validate checks the field values on CatalogMetadata with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CatalogMetadata) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ApiVersion

	// no validation rules for Kind

	// no validation rules for Name

	// no validation rules for Version

	// no validation rules for Description

	// no validation rules for License

	// no validation rules for Url

	// no validation rules for Doc

	if v, ok := interface{}(m.GetRequires()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CatalogMetadataValidationError{
				field:  "Requires",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetLogo() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CatalogMetadataValidationError{
					field:  fmt.Sprintf("Logo[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CatalogMetadataValidationError is the validation error returned by
// CatalogMetadata.Validate if the designated constraints aren't met.
type CatalogMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CatalogMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CatalogMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CatalogMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CatalogMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CatalogMetadataValidationError) ErrorName() string { return "CatalogMetadataValidationError" }

// Error satisfies the builtin error interface
func (e CatalogMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCatalogMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CatalogMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CatalogMetadataValidationError{}
