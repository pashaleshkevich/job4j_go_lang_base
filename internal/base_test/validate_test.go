package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	t.Run("When ValidateRequest is nil", func(t *testing.T) {
		t.Parallel()

		var in *base.ValidateRequest
		rsl := base.Validate(in)
		expected := []string{"ValidateRequest is nil"}
		assert.Equal(t, expected, rsl)
	})

	t.Run("When UserID is empty", func(t *testing.T) {
		t.Parallel()

		in := &base.ValidateRequest{
			Title:       "There is a title here",
			Description: "And here is his description",
		}
		rsl := base.Validate(in)
		expected := []string{"UserID is empty"}
		assert.Equal(t, expected, rsl)
	})

	t.Run("When Title is empty", func(t *testing.T) {
		t.Parallel()

		in := &base.ValidateRequest{
			UserID:      "UserID is definitely required",
			Description: "And here is his description",
		}
		rsl := base.Validate(in)
		expected := []string{"Title is empty"}
		assert.Equal(t, expected, rsl)
	})

	t.Run("When Description is empty", func(t *testing.T) {
		t.Parallel()

		in := &base.ValidateRequest{
			UserID: "UserID is definitely required",
			Title:  "There is a title here",
		}
		rsl := base.Validate(in)
		expected := []string{"Description is empty"}
		assert.Equal(t, expected, rsl)
	})

	t.Run("When UserID and Description are empty", func(t *testing.T) {
		t.Parallel()

		in := &base.ValidateRequest{
			Title: "There is a title here",
		}
		rsl := base.Validate(in)
		expected := []string{"UserID is empty", "Description is empty"}
		assert.Equal(t, expected, rsl)
	})

	t.Run("When UserID and Title are empty", func(t *testing.T) {
		t.Parallel()

		in := &base.ValidateRequest{
			Description: "And here is his description",
		}
		rsl := base.Validate(in)
		expected := []string{"UserID is empty", "Title is empty"}
		assert.Equal(t, expected, rsl)
	})

	t.Run("When all fields are empty", func(t *testing.T) {
		t.Parallel()

		in := &base.ValidateRequest{}
		rsl := base.Validate(in)
		expected := []string{"UserID is empty", "Title is empty", "Description is empty"}
		assert.Equal(t, expected, rsl)
	})
}
