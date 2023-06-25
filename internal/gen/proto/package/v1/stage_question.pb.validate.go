// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: package/v1/stage_question.proto

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

// Validate checks the field values on StageQuestion with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StageQuestion) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StageQuestion with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StageQuestionMultiError, or
// nil if none found.
func (m *StageQuestion) ValidateAll() error {
	return m.validate(true)
}

func (m *StageQuestion) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for StageId

	// no validation rules for TopicId

	if all {
		switch v := interface{}(m.GetQuestion()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StageQuestionValidationError{
					field:  "Question",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StageQuestionValidationError{
					field:  "Question",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQuestion()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StageQuestionValidationError{
				field:  "Question",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for QuestionType

	// no validation rules for QuestionCost

	if all {
		switch v := interface{}(m.GetAnswer()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StageQuestionValidationError{
					field:  "Answer",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StageQuestionValidationError{
					field:  "Answer",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAnswer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StageQuestionValidationError{
				field:  "Answer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for AnswerTime

	// no validation rules for HostComment

	// no validation rules for SecretTopic

	// no validation rules for SecretCost

	// no validation rules for TransferType

	// no validation rules for IsKeepable

	if len(errors) > 0 {
		return StageQuestionMultiError(errors)
	}

	return nil
}

// StageQuestionMultiError is an error wrapping multiple validation errors
// returned by StageQuestion.ValidateAll() if the designated constraints
// aren't met.
type StageQuestionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StageQuestionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StageQuestionMultiError) AllErrors() []error { return m }

// StageQuestionValidationError is the validation error returned by
// StageQuestion.Validate if the designated constraints aren't met.
type StageQuestionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StageQuestionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StageQuestionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StageQuestionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StageQuestionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StageQuestionValidationError) ErrorName() string { return "StageQuestionValidationError" }

// Error satisfies the builtin error interface
func (e StageQuestionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStageQuestion.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StageQuestionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StageQuestionValidationError{}

// Validate checks the field values on StageQuestion_Question with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StageQuestion_Question) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StageQuestion_Question with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StageQuestion_QuestionMultiError, or nil if none found.
func (m *StageQuestion_Question) ValidateAll() error {
	return m.validate(true)
}

func (m *StageQuestion_Question) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Text

	// no validation rules for Author

	// no validation rules for MediaUrl

	if all {
		switch v := interface{}(m.GetCreationTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StageQuestion_QuestionValidationError{
					field:  "CreationTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StageQuestion_QuestionValidationError{
					field:  "CreationTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreationTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StageQuestion_QuestionValidationError{
				field:  "CreationTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return StageQuestion_QuestionMultiError(errors)
	}

	return nil
}

// StageQuestion_QuestionMultiError is an error wrapping multiple validation
// errors returned by StageQuestion_Question.ValidateAll() if the designated
// constraints aren't met.
type StageQuestion_QuestionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StageQuestion_QuestionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StageQuestion_QuestionMultiError) AllErrors() []error { return m }

// StageQuestion_QuestionValidationError is the validation error returned by
// StageQuestion_Question.Validate if the designated constraints aren't met.
type StageQuestion_QuestionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StageQuestion_QuestionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StageQuestion_QuestionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StageQuestion_QuestionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StageQuestion_QuestionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StageQuestion_QuestionValidationError) ErrorName() string {
	return "StageQuestion_QuestionValidationError"
}

// Error satisfies the builtin error interface
func (e StageQuestion_QuestionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStageQuestion_Question.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StageQuestion_QuestionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StageQuestion_QuestionValidationError{}

// Validate checks the field values on StageQuestion_Answer with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StageQuestion_Answer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StageQuestion_Answer with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StageQuestion_AnswerMultiError, or nil if none found.
func (m *StageQuestion_Answer) ValidateAll() error {
	return m.validate(true)
}

func (m *StageQuestion_Answer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Text

	// no validation rules for Author

	// no validation rules for MediaUrl

	if all {
		switch v := interface{}(m.GetCreationTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StageQuestion_AnswerValidationError{
					field:  "CreationTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StageQuestion_AnswerValidationError{
					field:  "CreationTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreationTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StageQuestion_AnswerValidationError{
				field:  "CreationTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return StageQuestion_AnswerMultiError(errors)
	}

	return nil
}

// StageQuestion_AnswerMultiError is an error wrapping multiple validation
// errors returned by StageQuestion_Answer.ValidateAll() if the designated
// constraints aren't met.
type StageQuestion_AnswerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StageQuestion_AnswerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StageQuestion_AnswerMultiError) AllErrors() []error { return m }

// StageQuestion_AnswerValidationError is the validation error returned by
// StageQuestion_Answer.Validate if the designated constraints aren't met.
type StageQuestion_AnswerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StageQuestion_AnswerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StageQuestion_AnswerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StageQuestion_AnswerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StageQuestion_AnswerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StageQuestion_AnswerValidationError) ErrorName() string {
	return "StageQuestion_AnswerValidationError"
}

// Error satisfies the builtin error interface
func (e StageQuestion_AnswerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStageQuestion_Answer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StageQuestion_AnswerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StageQuestion_AnswerValidationError{}
