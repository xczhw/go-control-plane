//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/grpc_json_transcoder/v3/transcoder.proto

package grpc_json_transcoderv3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GrpcJsonTranscoder with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GrpcJsonTranscoder) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GrpcJsonTranscoder with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GrpcJsonTranscoderMultiError, or nil if none found.
func (m *GrpcJsonTranscoder) ValidateAll() error {
	return m.validate(true)
}

func (m *GrpcJsonTranscoder) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPrintOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GrpcJsonTranscoderValidationError{
					field:  "PrintOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GrpcJsonTranscoderValidationError{
					field:  "PrintOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPrintOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GrpcJsonTranscoderValidationError{
				field:  "PrintOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for MatchIncomingRequestRoute

	// no validation rules for AutoMapping

	// no validation rules for IgnoreUnknownQueryParameters

	// no validation rules for ConvertGrpcStatus

	if _, ok := GrpcJsonTranscoder_UrlUnescapeSpec_name[int32(m.GetUrlUnescapeSpec())]; !ok {
		err := GrpcJsonTranscoderValidationError{
			field:  "UrlUnescapeSpec",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for QueryParamUnescapePlus

	// no validation rules for MatchUnregisteredCustomVerb

	if all {
		switch v := interface{}(m.GetRequestValidationOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GrpcJsonTranscoderValidationError{
					field:  "RequestValidationOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GrpcJsonTranscoderValidationError{
					field:  "RequestValidationOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRequestValidationOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GrpcJsonTranscoderValidationError{
				field:  "RequestValidationOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for CaseInsensitiveEnumParsing

	if wrapper := m.GetMaxRequestBodySize(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			err := GrpcJsonTranscoderValidationError{
				field:  "MaxRequestBodySize",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if wrapper := m.GetMaxResponseBodySize(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			err := GrpcJsonTranscoderValidationError{
				field:  "MaxResponseBodySize",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	// no validation rules for CaptureUnknownQueryParameters

	oneofDescriptorSetPresent := false
	switch v := m.DescriptorSet.(type) {
	case *GrpcJsonTranscoder_ProtoDescriptor:
		if v == nil {
			err := GrpcJsonTranscoderValidationError{
				field:  "DescriptorSet",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofDescriptorSetPresent = true
		// no validation rules for ProtoDescriptor
	case *GrpcJsonTranscoder_ProtoDescriptorBin:
		if v == nil {
			err := GrpcJsonTranscoderValidationError{
				field:  "DescriptorSet",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofDescriptorSetPresent = true
		// no validation rules for ProtoDescriptorBin
	default:
		_ = v // ensures v is used
	}
	if !oneofDescriptorSetPresent {
		err := GrpcJsonTranscoderValidationError{
			field:  "DescriptorSet",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GrpcJsonTranscoderMultiError(errors)
	}

	return nil
}

// GrpcJsonTranscoderMultiError is an error wrapping multiple validation errors
// returned by GrpcJsonTranscoder.ValidateAll() if the designated constraints
// aren't met.
type GrpcJsonTranscoderMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrpcJsonTranscoderMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrpcJsonTranscoderMultiError) AllErrors() []error { return m }

// GrpcJsonTranscoderValidationError is the validation error returned by
// GrpcJsonTranscoder.Validate if the designated constraints aren't met.
type GrpcJsonTranscoderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrpcJsonTranscoderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrpcJsonTranscoderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrpcJsonTranscoderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrpcJsonTranscoderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrpcJsonTranscoderValidationError) ErrorName() string {
	return "GrpcJsonTranscoderValidationError"
}

// Error satisfies the builtin error interface
func (e GrpcJsonTranscoderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcJsonTranscoder.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrpcJsonTranscoderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrpcJsonTranscoderValidationError{}

// Validate checks the field values on UnknownQueryParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UnknownQueryParams) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UnknownQueryParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UnknownQueryParamsMultiError, or nil if none found.
func (m *UnknownQueryParams) ValidateAll() error {
	return m.validate(true)
}

func (m *UnknownQueryParams) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	{
		sorted_keys := make([]string, len(m.GetKey()))
		i := 0
		for key := range m.GetKey() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetKey()[key]
			_ = val

			// no validation rules for Key[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, UnknownQueryParamsValidationError{
							field:  fmt.Sprintf("Key[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, UnknownQueryParamsValidationError{
							field:  fmt.Sprintf("Key[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return UnknownQueryParamsValidationError{
						field:  fmt.Sprintf("Key[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return UnknownQueryParamsMultiError(errors)
	}

	return nil
}

// UnknownQueryParamsMultiError is an error wrapping multiple validation errors
// returned by UnknownQueryParams.ValidateAll() if the designated constraints
// aren't met.
type UnknownQueryParamsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnknownQueryParamsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnknownQueryParamsMultiError) AllErrors() []error { return m }

// UnknownQueryParamsValidationError is the validation error returned by
// UnknownQueryParams.Validate if the designated constraints aren't met.
type UnknownQueryParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnknownQueryParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnknownQueryParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnknownQueryParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnknownQueryParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnknownQueryParamsValidationError) ErrorName() string {
	return "UnknownQueryParamsValidationError"
}

// Error satisfies the builtin error interface
func (e UnknownQueryParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnknownQueryParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnknownQueryParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnknownQueryParamsValidationError{}

// Validate checks the field values on GrpcJsonTranscoder_PrintOptions with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GrpcJsonTranscoder_PrintOptions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GrpcJsonTranscoder_PrintOptions with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GrpcJsonTranscoder_PrintOptionsMultiError, or nil if none found.
func (m *GrpcJsonTranscoder_PrintOptions) ValidateAll() error {
	return m.validate(true)
}

func (m *GrpcJsonTranscoder_PrintOptions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AddWhitespace

	// no validation rules for AlwaysPrintPrimitiveFields

	// no validation rules for AlwaysPrintEnumsAsInts

	// no validation rules for PreserveProtoFieldNames

	// no validation rules for StreamNewlineDelimited

	if len(errors) > 0 {
		return GrpcJsonTranscoder_PrintOptionsMultiError(errors)
	}

	return nil
}

// GrpcJsonTranscoder_PrintOptionsMultiError is an error wrapping multiple
// validation errors returned by GrpcJsonTranscoder_PrintOptions.ValidateAll()
// if the designated constraints aren't met.
type GrpcJsonTranscoder_PrintOptionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrpcJsonTranscoder_PrintOptionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrpcJsonTranscoder_PrintOptionsMultiError) AllErrors() []error { return m }

// GrpcJsonTranscoder_PrintOptionsValidationError is the validation error
// returned by GrpcJsonTranscoder_PrintOptions.Validate if the designated
// constraints aren't met.
type GrpcJsonTranscoder_PrintOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrpcJsonTranscoder_PrintOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrpcJsonTranscoder_PrintOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrpcJsonTranscoder_PrintOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrpcJsonTranscoder_PrintOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrpcJsonTranscoder_PrintOptionsValidationError) ErrorName() string {
	return "GrpcJsonTranscoder_PrintOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e GrpcJsonTranscoder_PrintOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcJsonTranscoder_PrintOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrpcJsonTranscoder_PrintOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrpcJsonTranscoder_PrintOptionsValidationError{}

// Validate checks the field values on
// GrpcJsonTranscoder_RequestValidationOptions with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GrpcJsonTranscoder_RequestValidationOptions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// GrpcJsonTranscoder_RequestValidationOptions with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// GrpcJsonTranscoder_RequestValidationOptionsMultiError, or nil if none found.
func (m *GrpcJsonTranscoder_RequestValidationOptions) ValidateAll() error {
	return m.validate(true)
}

func (m *GrpcJsonTranscoder_RequestValidationOptions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RejectUnknownMethod

	// no validation rules for RejectUnknownQueryParameters

	// no validation rules for RejectBindingBodyFieldCollisions

	if len(errors) > 0 {
		return GrpcJsonTranscoder_RequestValidationOptionsMultiError(errors)
	}

	return nil
}

// GrpcJsonTranscoder_RequestValidationOptionsMultiError is an error wrapping
// multiple validation errors returned by
// GrpcJsonTranscoder_RequestValidationOptions.ValidateAll() if the designated
// constraints aren't met.
type GrpcJsonTranscoder_RequestValidationOptionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrpcJsonTranscoder_RequestValidationOptionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrpcJsonTranscoder_RequestValidationOptionsMultiError) AllErrors() []error { return m }

// GrpcJsonTranscoder_RequestValidationOptionsValidationError is the validation
// error returned by GrpcJsonTranscoder_RequestValidationOptions.Validate if
// the designated constraints aren't met.
type GrpcJsonTranscoder_RequestValidationOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrpcJsonTranscoder_RequestValidationOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrpcJsonTranscoder_RequestValidationOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrpcJsonTranscoder_RequestValidationOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrpcJsonTranscoder_RequestValidationOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrpcJsonTranscoder_RequestValidationOptionsValidationError) ErrorName() string {
	return "GrpcJsonTranscoder_RequestValidationOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e GrpcJsonTranscoder_RequestValidationOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcJsonTranscoder_RequestValidationOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrpcJsonTranscoder_RequestValidationOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrpcJsonTranscoder_RequestValidationOptionsValidationError{}

// Validate checks the field values on UnknownQueryParams_Values with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UnknownQueryParams_Values) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UnknownQueryParams_Values with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UnknownQueryParams_ValuesMultiError, or nil if none found.
func (m *UnknownQueryParams_Values) ValidateAll() error {
	return m.validate(true)
}

func (m *UnknownQueryParams_Values) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UnknownQueryParams_ValuesMultiError(errors)
	}

	return nil
}

// UnknownQueryParams_ValuesMultiError is an error wrapping multiple validation
// errors returned by UnknownQueryParams_Values.ValidateAll() if the
// designated constraints aren't met.
type UnknownQueryParams_ValuesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnknownQueryParams_ValuesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnknownQueryParams_ValuesMultiError) AllErrors() []error { return m }

// UnknownQueryParams_ValuesValidationError is the validation error returned by
// UnknownQueryParams_Values.Validate if the designated constraints aren't met.
type UnknownQueryParams_ValuesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnknownQueryParams_ValuesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnknownQueryParams_ValuesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnknownQueryParams_ValuesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnknownQueryParams_ValuesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnknownQueryParams_ValuesValidationError) ErrorName() string {
	return "UnknownQueryParams_ValuesValidationError"
}

// Error satisfies the builtin error interface
func (e UnknownQueryParams_ValuesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnknownQueryParams_Values.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnknownQueryParams_ValuesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnknownQueryParams_ValuesValidationError{}