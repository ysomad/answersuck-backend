// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: package/v1/package.proto

package packagev1

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

// Validate checks the field values on Package with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Package) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Package with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PackageMultiError, or nil if none found.
func (m *Package) ValidateAll() error {
	return m.validate(true)
}

func (m *Package) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Author

	// no validation rules for IsPublished

	// no validation rules for CoverUrl

	for idx, item := range m.GetStages() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PackageValidationError{
						field:  fmt.Sprintf("Stages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PackageValidationError{
						field:  fmt.Sprintf("Stages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PackageValidationError{
					field:  fmt.Sprintf("Stages[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetCreationTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PackageValidationError{
					field:  "CreationTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PackageValidationError{
					field:  "CreationTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreationTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PackageValidationError{
				field:  "CreationTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PackageValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PackageValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PackageValidationError{
				field:  "UpdateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PackageMultiError(errors)
	}

	return nil
}

// PackageMultiError is an error wrapping multiple validation errors returned
// by Package.ValidateAll() if the designated constraints aren't met.
type PackageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PackageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PackageMultiError) AllErrors() []error { return m }

// PackageValidationError is the validation error returned by Package.Validate
// if the designated constraints aren't met.
type PackageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PackageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PackageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PackageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PackageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PackageValidationError) ErrorName() string { return "PackageValidationError" }

// Error satisfies the builtin error interface
func (e PackageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPackage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PackageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PackageValidationError{}

// Validate checks the field values on Package_Stage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Package_Stage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Package_Stage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Package_StageMultiError, or
// nil if none found.
func (m *Package_Stage) ValidateAll() error {
	return m.validate(true)
}

func (m *Package_Stage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	if len(errors) > 0 {
		return Package_StageMultiError(errors)
	}

	return nil
}

// Package_StageMultiError is an error wrapping multiple validation errors
// returned by Package_Stage.ValidateAll() if the designated constraints
// aren't met.
type Package_StageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Package_StageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Package_StageMultiError) AllErrors() []error { return m }

// Package_StageValidationError is the validation error returned by
// Package_Stage.Validate if the designated constraints aren't met.
type Package_StageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Package_StageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Package_StageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Package_StageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Package_StageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Package_StageValidationError) ErrorName() string { return "Package_StageValidationError" }

// Error satisfies the builtin error interface
func (e Package_StageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPackage_Stage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Package_StageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Package_StageValidationError{}