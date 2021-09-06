// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/api.proto

package api

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
)

// Validate checks the field values on CreateRecipeRequestV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateRecipeRequestV1) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUserId() <= 0 {
		return CreateRecipeRequestV1ValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return CreateRecipeRequestV1ValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		return CreateRecipeRequestV1ValidationError{
			field:  "Description",
			reason: "value length must be at least 1 runes",
		}
	}

	if len(m.GetActions()) < 1 {
		return CreateRecipeRequestV1ValidationError{
			field:  "Actions",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetActions() {
		_, _ = idx, item

		if utf8.RuneCountInString(item) < 1 {
			return CreateRecipeRequestV1ValidationError{
				field:  fmt.Sprintf("Actions[%v]", idx),
				reason: "value length must be at least 1 runes",
			}
		}

	}

	return nil
}

// CreateRecipeRequestV1ValidationError is the validation error returned by
// CreateRecipeRequestV1.Validate if the designated constraints aren't met.
type CreateRecipeRequestV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRecipeRequestV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRecipeRequestV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRecipeRequestV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRecipeRequestV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRecipeRequestV1ValidationError) ErrorName() string {
	return "CreateRecipeRequestV1ValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRecipeRequestV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRecipeRequestV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRecipeRequestV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRecipeRequestV1ValidationError{}

// Validate checks the field values on CreateRecipeResponseV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateRecipeResponseV1) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRecipeId() <= 0 {
		return CreateRecipeResponseV1ValidationError{
			field:  "RecipeId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// CreateRecipeResponseV1ValidationError is the validation error returned by
// CreateRecipeResponseV1.Validate if the designated constraints aren't met.
type CreateRecipeResponseV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRecipeResponseV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRecipeResponseV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRecipeResponseV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRecipeResponseV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRecipeResponseV1ValidationError) ErrorName() string {
	return "CreateRecipeResponseV1ValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRecipeResponseV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRecipeResponseV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRecipeResponseV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRecipeResponseV1ValidationError{}

// Validate checks the field values on CreateRecipeV1 with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CreateRecipeV1) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUserId() <= 0 {
		return CreateRecipeV1ValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return CreateRecipeV1ValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		return CreateRecipeV1ValidationError{
			field:  "Description",
			reason: "value length must be at least 1 runes",
		}
	}

	if len(m.GetActions()) < 1 {
		return CreateRecipeV1ValidationError{
			field:  "Actions",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetActions() {
		_, _ = idx, item

		if utf8.RuneCountInString(item) < 1 {
			return CreateRecipeV1ValidationError{
				field:  fmt.Sprintf("Actions[%v]", idx),
				reason: "value length must be at least 1 runes",
			}
		}

	}

	return nil
}

// CreateRecipeV1ValidationError is the validation error returned by
// CreateRecipeV1.Validate if the designated constraints aren't met.
type CreateRecipeV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRecipeV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRecipeV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRecipeV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRecipeV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRecipeV1ValidationError) ErrorName() string { return "CreateRecipeV1ValidationError" }

// Error satisfies the builtin error interface
func (e CreateRecipeV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRecipeV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRecipeV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRecipeV1ValidationError{}

// Validate checks the field values on MultiCreateRecipeRequestV1 with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateRecipeRequestV1) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetRecipes()) < 1 {
		return MultiCreateRecipeRequestV1ValidationError{
			field:  "Recipes",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetRecipes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MultiCreateRecipeRequestV1ValidationError{
					field:  fmt.Sprintf("Recipes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MultiCreateRecipeRequestV1ValidationError is the validation error returned
// by MultiCreateRecipeRequestV1.Validate if the designated constraints aren't met.
type MultiCreateRecipeRequestV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateRecipeRequestV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateRecipeRequestV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateRecipeRequestV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateRecipeRequestV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateRecipeRequestV1ValidationError) ErrorName() string {
	return "MultiCreateRecipeRequestV1ValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateRecipeRequestV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateRecipeRequestV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateRecipeRequestV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateRecipeRequestV1ValidationError{}

// Validate checks the field values on MultiCreateRecipeResponseV1 with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateRecipeResponseV1) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// MultiCreateRecipeResponseV1ValidationError is the validation error returned
// by MultiCreateRecipeResponseV1.Validate if the designated constraints
// aren't met.
type MultiCreateRecipeResponseV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateRecipeResponseV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateRecipeResponseV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateRecipeResponseV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateRecipeResponseV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateRecipeResponseV1ValidationError) ErrorName() string {
	return "MultiCreateRecipeResponseV1ValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateRecipeResponseV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateRecipeResponseV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateRecipeResponseV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateRecipeResponseV1ValidationError{}

// Validate checks the field values on DescribeRecipeRequestV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeRecipeRequestV1) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRecipeId() <= 0 {
		return DescribeRecipeRequestV1ValidationError{
			field:  "RecipeId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeRecipeRequestV1ValidationError is the validation error returned by
// DescribeRecipeRequestV1.Validate if the designated constraints aren't met.
type DescribeRecipeRequestV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeRecipeRequestV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeRecipeRequestV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeRecipeRequestV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeRecipeRequestV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeRecipeRequestV1ValidationError) ErrorName() string {
	return "DescribeRecipeRequestV1ValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeRecipeRequestV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeRecipeRequestV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeRecipeRequestV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeRecipeRequestV1ValidationError{}

// Validate checks the field values on RecipeV1 with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RecipeV1) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRecipeId() <= 0 {
		return RecipeV1ValidationError{
			field:  "RecipeId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetUserId() <= 0 {
		return RecipeV1ValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return RecipeV1ValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		return RecipeV1ValidationError{
			field:  "Description",
			reason: "value length must be at least 1 runes",
		}
	}

	if len(m.GetActions()) < 1 {
		return RecipeV1ValidationError{
			field:  "Actions",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetActions() {
		_, _ = idx, item

		if utf8.RuneCountInString(item) < 1 {
			return RecipeV1ValidationError{
				field:  fmt.Sprintf("Actions[%v]", idx),
				reason: "value length must be at least 1 runes",
			}
		}

	}

	return nil
}

// RecipeV1ValidationError is the validation error returned by
// RecipeV1.Validate if the designated constraints aren't met.
type RecipeV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RecipeV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RecipeV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RecipeV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RecipeV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RecipeV1ValidationError) ErrorName() string { return "RecipeV1ValidationError" }

// Error satisfies the builtin error interface
func (e RecipeV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRecipeV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RecipeV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RecipeV1ValidationError{}

// Validate checks the field values on DescribeRecipeResponseV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeRecipeResponseV1) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRecipe()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeRecipeResponseV1ValidationError{
				field:  "Recipe",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeRecipeResponseV1ValidationError is the validation error returned by
// DescribeRecipeResponseV1.Validate if the designated constraints aren't met.
type DescribeRecipeResponseV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeRecipeResponseV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeRecipeResponseV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeRecipeResponseV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeRecipeResponseV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeRecipeResponseV1ValidationError) ErrorName() string {
	return "DescribeRecipeResponseV1ValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeRecipeResponseV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeRecipeResponseV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeRecipeResponseV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeRecipeResponseV1ValidationError{}

// Validate checks the field values on ListRecipesRequestV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListRecipesRequestV1) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetLimit() <= 0 {
		return ListRecipesRequestV1ValidationError{
			field:  "Limit",
			reason: "value must be greater than 0",
		}
	}

	// no validation rules for Offset

	return nil
}

// ListRecipesRequestV1ValidationError is the validation error returned by
// ListRecipesRequestV1.Validate if the designated constraints aren't met.
type ListRecipesRequestV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRecipesRequestV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRecipesRequestV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRecipesRequestV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRecipesRequestV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRecipesRequestV1ValidationError) ErrorName() string {
	return "ListRecipesRequestV1ValidationError"
}

// Error satisfies the builtin error interface
func (e ListRecipesRequestV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRecipesRequestV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRecipesRequestV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRecipesRequestV1ValidationError{}

// Validate checks the field values on ListRecipesResponseV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListRecipesResponseV1) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRecipes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListRecipesResponseV1ValidationError{
					field:  fmt.Sprintf("Recipes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListRecipesResponseV1ValidationError is the validation error returned by
// ListRecipesResponseV1.Validate if the designated constraints aren't met.
type ListRecipesResponseV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRecipesResponseV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRecipesResponseV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRecipesResponseV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRecipesResponseV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRecipesResponseV1ValidationError) ErrorName() string {
	return "ListRecipesResponseV1ValidationError"
}

// Error satisfies the builtin error interface
func (e ListRecipesResponseV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRecipesResponseV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRecipesResponseV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRecipesResponseV1ValidationError{}

// Validate checks the field values on RemoveRecipeRequestV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveRecipeRequestV1) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRecipeId() <= 0 {
		return RemoveRecipeRequestV1ValidationError{
			field:  "RecipeId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemoveRecipeRequestV1ValidationError is the validation error returned by
// RemoveRecipeRequestV1.Validate if the designated constraints aren't met.
type RemoveRecipeRequestV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveRecipeRequestV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveRecipeRequestV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveRecipeRequestV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveRecipeRequestV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveRecipeRequestV1ValidationError) ErrorName() string {
	return "RemoveRecipeRequestV1ValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveRecipeRequestV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveRecipeRequestV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveRecipeRequestV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveRecipeRequestV1ValidationError{}

// Validate checks the field values on RemoveRecipesResponseV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveRecipesResponseV1) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRecipeId() <= 0 {
		return RemoveRecipesResponseV1ValidationError{
			field:  "RecipeId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemoveRecipesResponseV1ValidationError is the validation error returned by
// RemoveRecipesResponseV1.Validate if the designated constraints aren't met.
type RemoveRecipesResponseV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveRecipesResponseV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveRecipesResponseV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveRecipesResponseV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveRecipesResponseV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveRecipesResponseV1ValidationError) ErrorName() string {
	return "RemoveRecipesResponseV1ValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveRecipesResponseV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveRecipesResponseV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveRecipesResponseV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveRecipesResponseV1ValidationError{}

// Validate checks the field values on UpdateRecipeRequestV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateRecipeRequestV1) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRecipe()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateRecipeRequestV1ValidationError{
				field:  "Recipe",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateRecipeRequestV1ValidationError is the validation error returned by
// UpdateRecipeRequestV1.Validate if the designated constraints aren't met.
type UpdateRecipeRequestV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRecipeRequestV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRecipeRequestV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRecipeRequestV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRecipeRequestV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRecipeRequestV1ValidationError) ErrorName() string {
	return "UpdateRecipeRequestV1ValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRecipeRequestV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRecipeRequestV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRecipeRequestV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRecipeRequestV1ValidationError{}

// Validate checks the field values on UpdateRecipeResponseV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateRecipeResponseV1) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateRecipeResponseV1ValidationError is the validation error returned by
// UpdateRecipeResponseV1.Validate if the designated constraints aren't met.
type UpdateRecipeResponseV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRecipeResponseV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRecipeResponseV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRecipeResponseV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRecipeResponseV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRecipeResponseV1ValidationError) ErrorName() string {
	return "UpdateRecipeResponseV1ValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRecipeResponseV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRecipeResponseV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRecipeResponseV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRecipeResponseV1ValidationError{}
