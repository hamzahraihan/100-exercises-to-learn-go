package validate

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewTicketValid(t *testing.T) {
	ticket, err := NewTicket("Bug", "crashes")
	if err != nil {
		t.Fatalf("got error %v, want nil", err)
	}
	if ticket.Title != "Bug" || ticket.Description != "crashes" {
		t.Fatalf("got ticket %+v, want {Title:Bug Description:crashes}", ticket)
	}
}

func TestNewTicketEmptyTitle(t *testing.T) {
	_, err := NewTicket("", "crashes")
	if err == nil {
		t.Fatal("got nil error, want a ValidationError for empty title")
	}
	var ve *ValidationError
	if !errors.As(err, &ve) {
		t.Fatalf("got %T, want *ValidationError in chain (hint: wrap with fmt.Errorf(\"new ticket: %%w\", ...))", err)
	}
	if ve.Field != "title" {
		t.Fatalf("got Field=%q, want %q (hint: set Field=\"title\" when title is empty)", ve.Field, "title")
	}
}

func TestNewTicketEmptyDesc(t *testing.T) {
	_, err := NewTicket("Bug", "")
	if err == nil {
		t.Fatal("got nil error, want a ValidationError for empty description")
	}
	var ve *ValidationError
	if !errors.As(err, &ve) {
		t.Fatalf("got %T, want *ValidationError in chain (hint: wrap with fmt.Errorf(\"new ticket: %%w\", ...))", err)
	}
	if ve.Field != "description" {
		t.Fatalf("got Field=%q, want %q (hint: set Field=\"description\" when description is empty)", ve.Field, "description")
	}
}

func TestExtractField(t *testing.T) {
	inner := &ValidationError{Field: "priority", Message: "required"}
	err := fmt.Errorf("wrap: %w", inner)
	field, ok := ExtractField(err)
	if !ok {
		t.Fatal("got ok=false, want true (hint: use errors.As to unwrap *ValidationError)")
	}
	if field != "priority" {
		t.Fatalf("got field=%q, want %q (hint: return ve.Field after errors.As succeeds)", field, "priority")
	}
}

func TestExtractFieldNil(t *testing.T) {
	field, ok := ExtractField(nil)
	if ok || field != "" {
		t.Fatalf("got (%q, %v), want (\"\", false) for nil error", field, ok)
	}
}

func TestExtractFieldMismatch(t *testing.T) {
	field, ok := ExtractField(errors.New("plain"))
	if ok || field != "" {
		t.Fatalf("got (%q, %v), want (\"\", false) for non-ValidationError", field, ok)
	}
}
