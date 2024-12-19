//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/router/v2/router.proto

package routerv2

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

// Validate checks the field values on Router with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Router) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Router with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RouterMultiError, or nil if none found.
func (m *Router) ValidateAll() error {
	return m.validate(true)
}

func (m *Router) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDynamicStats()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RouterValidationError{
					field:  "DynamicStats",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RouterValidationError{
					field:  "DynamicStats",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDynamicStats()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouterValidationError{
				field:  "DynamicStats",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for StartChildSpan

	for idx, item := range m.GetUpstreamLog() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RouterValidationError{
						field:  fmt.Sprintf("UpstreamLog[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RouterValidationError{
						field:  fmt.Sprintf("UpstreamLog[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouterValidationError{
					field:  fmt.Sprintf("UpstreamLog[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for SuppressEnvoyHeaders

	for idx, item := range m.GetStrictCheckHeaders() {
		_, _ = idx, item

		if _, ok := _Router_StrictCheckHeaders_InLookup[item]; !ok {
			err := RouterValidationError{
				field:  fmt.Sprintf("StrictCheckHeaders[%v]", idx),
				reason: "value must be in list [x-envoy-upstream-rq-timeout-ms x-envoy-upstream-rq-per-try-timeout-ms x-envoy-max-retries x-envoy-retry-grpc-on x-envoy-retry-on]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	// no validation rules for RespectExpectedRqTimeout

	if len(errors) > 0 {
		return RouterMultiError(errors)
	}

	return nil
}

// RouterMultiError is an error wrapping multiple validation errors returned by
// Router.ValidateAll() if the designated constraints aren't met.
type RouterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RouterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RouterMultiError) AllErrors() []error { return m }

// RouterValidationError is the validation error returned by Router.Validate if
// the designated constraints aren't met.
type RouterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouterValidationError) ErrorName() string { return "RouterValidationError" }

// Error satisfies the builtin error interface
func (e RouterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouterValidationError{}

var _Router_StrictCheckHeaders_InLookup = map[string]struct{}{
	"x-envoy-upstream-rq-timeout-ms":         {},
	"x-envoy-upstream-rq-per-try-timeout-ms": {},
	"x-envoy-max-retries":                    {},
	"x-envoy-retry-grpc-on":                  {},
	"x-envoy-retry-on":                       {},
}