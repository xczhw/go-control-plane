//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/fault/v3/fault.proto

package faultv3

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

// Validate checks the field values on FaultAbort with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FaultAbort) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FaultAbort with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FaultAbortMultiError, or
// nil if none found.
func (m *FaultAbort) ValidateAll() error {
	return m.validate(true)
}

func (m *FaultAbort) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPercentage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FaultAbortValidationError{
					field:  "Percentage",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FaultAbortValidationError{
					field:  "Percentage",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPercentage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FaultAbortValidationError{
				field:  "Percentage",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	oneofErrorTypePresent := false
	switch v := m.ErrorType.(type) {
	case *FaultAbort_HttpStatus:
		if v == nil {
			err := FaultAbortValidationError{
				field:  "ErrorType",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofErrorTypePresent = true

		if val := m.GetHttpStatus(); val < 200 || val >= 600 {
			err := FaultAbortValidationError{
				field:  "HttpStatus",
				reason: "value must be inside range [200, 600)",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *FaultAbort_GrpcStatus:
		if v == nil {
			err := FaultAbortValidationError{
				field:  "ErrorType",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofErrorTypePresent = true
		// no validation rules for GrpcStatus
	case *FaultAbort_HeaderAbort_:
		if v == nil {
			err := FaultAbortValidationError{
				field:  "ErrorType",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofErrorTypePresent = true

		if all {
			switch v := interface{}(m.GetHeaderAbort()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FaultAbortValidationError{
						field:  "HeaderAbort",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FaultAbortValidationError{
						field:  "HeaderAbort",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetHeaderAbort()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FaultAbortValidationError{
					field:  "HeaderAbort",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofErrorTypePresent {
		err := FaultAbortValidationError{
			field:  "ErrorType",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return FaultAbortMultiError(errors)
	}

	return nil
}

// FaultAbortMultiError is an error wrapping multiple validation errors
// returned by FaultAbort.ValidateAll() if the designated constraints aren't met.
type FaultAbortMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FaultAbortMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FaultAbortMultiError) AllErrors() []error { return m }

// FaultAbortValidationError is the validation error returned by
// FaultAbort.Validate if the designated constraints aren't met.
type FaultAbortValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultAbortValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultAbortValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultAbortValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultAbortValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultAbortValidationError) ErrorName() string { return "FaultAbortValidationError" }

// Error satisfies the builtin error interface
func (e FaultAbortValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultAbort.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultAbortValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultAbortValidationError{}

// Validate checks the field values on HTTPFault with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HTTPFault) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HTTPFault with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HTTPFaultMultiError, or nil
// if none found.
func (m *HTTPFault) ValidateAll() error {
	return m.validate(true)
}

func (m *HTTPFault) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDelay()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HTTPFaultValidationError{
					field:  "Delay",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HTTPFaultValidationError{
					field:  "Delay",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDelay()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HTTPFaultValidationError{
				field:  "Delay",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAbort()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HTTPFaultValidationError{
					field:  "Abort",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HTTPFaultValidationError{
					field:  "Abort",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAbort()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HTTPFaultValidationError{
				field:  "Abort",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UpstreamCluster

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HTTPFaultValidationError{
						field:  fmt.Sprintf("Headers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HTTPFaultValidationError{
						field:  fmt.Sprintf("Headers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HTTPFaultValidationError{
					field:  fmt.Sprintf("Headers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetMaxActiveFaults()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HTTPFaultValidationError{
					field:  "MaxActiveFaults",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HTTPFaultValidationError{
					field:  "MaxActiveFaults",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMaxActiveFaults()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HTTPFaultValidationError{
				field:  "MaxActiveFaults",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetResponseRateLimit()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HTTPFaultValidationError{
					field:  "ResponseRateLimit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HTTPFaultValidationError{
					field:  "ResponseRateLimit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResponseRateLimit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HTTPFaultValidationError{
				field:  "ResponseRateLimit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DelayPercentRuntime

	// no validation rules for AbortPercentRuntime

	// no validation rules for DelayDurationRuntime

	// no validation rules for AbortHttpStatusRuntime

	// no validation rules for MaxActiveFaultsRuntime

	// no validation rules for ResponseRateLimitPercentRuntime

	// no validation rules for AbortGrpcStatusRuntime

	// no validation rules for DisableDownstreamClusterStats

	if all {
		switch v := interface{}(m.GetFilterMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HTTPFaultValidationError{
					field:  "FilterMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HTTPFaultValidationError{
					field:  "FilterMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFilterMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HTTPFaultValidationError{
				field:  "FilterMetadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return HTTPFaultMultiError(errors)
	}

	return nil
}

// HTTPFaultMultiError is an error wrapping multiple validation errors returned
// by HTTPFault.ValidateAll() if the designated constraints aren't met.
type HTTPFaultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HTTPFaultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HTTPFaultMultiError) AllErrors() []error { return m }

// HTTPFaultValidationError is the validation error returned by
// HTTPFault.Validate if the designated constraints aren't met.
type HTTPFaultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HTTPFaultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HTTPFaultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HTTPFaultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HTTPFaultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HTTPFaultValidationError) ErrorName() string { return "HTTPFaultValidationError" }

// Error satisfies the builtin error interface
func (e HTTPFaultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHTTPFault.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HTTPFaultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HTTPFaultValidationError{}

// Validate checks the field values on FaultAbort_HeaderAbort with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FaultAbort_HeaderAbort) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FaultAbort_HeaderAbort with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FaultAbort_HeaderAbortMultiError, or nil if none found.
func (m *FaultAbort_HeaderAbort) ValidateAll() error {
	return m.validate(true)
}

func (m *FaultAbort_HeaderAbort) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return FaultAbort_HeaderAbortMultiError(errors)
	}

	return nil
}

// FaultAbort_HeaderAbortMultiError is an error wrapping multiple validation
// errors returned by FaultAbort_HeaderAbort.ValidateAll() if the designated
// constraints aren't met.
type FaultAbort_HeaderAbortMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FaultAbort_HeaderAbortMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FaultAbort_HeaderAbortMultiError) AllErrors() []error { return m }

// FaultAbort_HeaderAbortValidationError is the validation error returned by
// FaultAbort_HeaderAbort.Validate if the designated constraints aren't met.
type FaultAbort_HeaderAbortValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultAbort_HeaderAbortValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultAbort_HeaderAbortValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultAbort_HeaderAbortValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultAbort_HeaderAbortValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultAbort_HeaderAbortValidationError) ErrorName() string {
	return "FaultAbort_HeaderAbortValidationError"
}

// Error satisfies the builtin error interface
func (e FaultAbort_HeaderAbortValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultAbort_HeaderAbort.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultAbort_HeaderAbortValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultAbort_HeaderAbortValidationError{}