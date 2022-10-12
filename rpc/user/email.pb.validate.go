// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user/email.proto

package user

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

// define the regex for a UUID once up-front
var _email_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on SendEmailVerificationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SendEmailVerificationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendEmailVerificationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SendEmailVerificationRequestMultiError, or nil if none found.
func (m *SendEmailVerificationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SendEmailVerificationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetAccountId()); err != nil {
		err = SendEmailVerificationRequestValidationError{
			field:  "AccountId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SendEmailVerificationRequestMultiError(errors)
	}

	return nil
}

func (m *SendEmailVerificationRequest) _validateUuid(uuid string) error {
	if matched := _email_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// SendEmailVerificationRequestMultiError is an error wrapping multiple
// validation errors returned by SendEmailVerificationRequest.ValidateAll() if
// the designated constraints aren't met.
type SendEmailVerificationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendEmailVerificationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendEmailVerificationRequestMultiError) AllErrors() []error { return m }

// SendEmailVerificationRequestValidationError is the validation error returned
// by SendEmailVerificationRequest.Validate if the designated constraints
// aren't met.
type SendEmailVerificationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendEmailVerificationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendEmailVerificationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendEmailVerificationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendEmailVerificationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendEmailVerificationRequestValidationError) ErrorName() string {
	return "SendEmailVerificationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SendEmailVerificationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendEmailVerificationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendEmailVerificationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendEmailVerificationRequestValidationError{}

// Validate checks the field values on UpdateEmailRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateEmailRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateEmailRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateEmailRequestMultiError, or nil if none found.
func (m *UpdateEmailRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateEmailRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetAccountId()); err != nil {
		err = UpdateEmailRequestValidationError{
			field:  "AccountId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateEmail(m.GetNewEmail()); err != nil {
		err = UpdateEmailRequestValidationError{
			field:  "NewEmail",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateEmailRequestMultiError(errors)
	}

	return nil
}

func (m *UpdateEmailRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *UpdateEmailRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

func (m *UpdateEmailRequest) _validateUuid(uuid string) error {
	if matched := _email_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// UpdateEmailRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateEmailRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateEmailRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateEmailRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateEmailRequestMultiError) AllErrors() []error { return m }

// UpdateEmailRequestValidationError is the validation error returned by
// UpdateEmailRequest.Validate if the designated constraints aren't met.
type UpdateEmailRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateEmailRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateEmailRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateEmailRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateEmailRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateEmailRequestValidationError) ErrorName() string {
	return "UpdateEmailRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateEmailRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateEmailRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateEmailRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateEmailRequestValidationError{}

// Validate checks the field values on VerifyEmailRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VerifyEmailRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyEmailRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerifyEmailRequestMultiError, or nil if none found.
func (m *VerifyEmailRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyEmailRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for VerificationCode

	if len(errors) > 0 {
		return VerifyEmailRequestMultiError(errors)
	}

	return nil
}

// VerifyEmailRequestMultiError is an error wrapping multiple validation errors
// returned by VerifyEmailRequest.ValidateAll() if the designated constraints
// aren't met.
type VerifyEmailRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyEmailRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyEmailRequestMultiError) AllErrors() []error { return m }

// VerifyEmailRequestValidationError is the validation error returned by
// VerifyEmailRequest.Validate if the designated constraints aren't met.
type VerifyEmailRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyEmailRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyEmailRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyEmailRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyEmailRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyEmailRequestValidationError) ErrorName() string {
	return "VerifyEmailRequestValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyEmailRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyEmailRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyEmailRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyEmailRequestValidationError{}
