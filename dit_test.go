package dit

import (
	"testing"
	"time"
)

func TestNewDIT(t *testing.T) {
	expected := DIT(51015)
	got, err := NewDIT(5, 10, 15)
	if err != nil {
		t.Fatal(err)
	}
	if expected != got {
		t.Fatalf("expected:%d got:%d", expected, got)
	}
}

func TestDecOOB(t *testing.T) {
	underOver := []int{-1, 10}
	for _, ou := range underOver {
		_, err := NewDIT(ou, 0, 0)
		switch et := err.(type) {
		case ErrDecOOB:
			return
		default:
			t.Fatalf("expected ErrDecOOB for %d, got %T", ou, et)
		}
	}
}

func TestDecimOOB(t *testing.T) {
	underOver := []int{-1, 100}
	for _, ou := range underOver {
		_, err := NewDIT(0, ou, 0)
		switch et := err.(type) {
		case ErrDecimOOB:
			return
		default:
			t.Fatalf("expected ErrDecimOOB for %d, got %T", ou, et)
		}
	}
}

func TestDesekOOB(t *testing.T) {
	underOver := []int{-1, 100}
	for _, ou := range underOver {
		_, err := NewDIT(0, 0, ou)
		switch et := err.(type) {
		case ErrDesekOOB:
			return
		default:
			t.Fatalf("expected ErrDesekOOB for %d, got %T", ou, et)
		}
	}
}

func TestTimeToDIT(t *testing.T) {
	//loc, err := time.LoadLocation(tzName)
	la, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		t.Fatal(err)
	}
	dl, err := time.LoadLocation(tzName)
	if err != nil {
		t.Fatal(err)
	}
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		t.Fatal(err)
	}

	cases := map[time.Time]DIT{
		time.Date(2022, 10, 14, 0, 0, 1, 0, la):  79168,
		time.Date(2022, 10, 14, 0, 0, 1, 0, dl):  1,
		time.Date(2022, 10, 14, 0, 0, 1, 0, utc): 50001,
		time.Date(2022, 10, 14, 0, 0, 0, 0, dl):  0,
	}
	for tt, expected := range cases {
		got := TimeToDIT(tt)
		if got != expected {
			t.Fatalf("expected:%d got:%d", expected, got)
		}
	}
}
