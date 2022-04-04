// Package dit supports the conversion of Decimal Internet Time from
// time.Time.
package dit

// MIT License
//
// Copyright (c) 2022 Lars Lehtonen
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

import (
	"fmt"
	"math"
	"time"
)

// DIT represents decimal internet time and is stored as decims since "midnight."
type DIT int

// Desek is 1/100 000 of a day.
type Desek uint8

// Decim is 1/1 000 of a day.
type Decim uint8

// Dec is 1/10 of a day.
type Dec uint8

// ErrDecOOB represents an out-of-bounds dec.
type ErrDecOOB struct{ Got int }

// ErrDecimOOB represents an out-of-bounds decim.
type ErrDecimOOB struct{ Got int }

// ErrDesekOOB respresent and out-of-bounds Desek.
type ErrDesekOOB struct{ Got int }

var dateline *time.Location

const (
	hourSeconds   = 3600
	minuteSeconds = 60
	maxDec        = 9
	maxDecim      = 99
	maxDesek      = 99
	tzName        = "Etc/GMT+12"
)

func init() {
	var err error
	dateline, err = time.LoadLocation(tzName)
	if err != nil {
		// if we didn't get this right at build time we're doomed.
		panic(err)
	}
}

// ErrDecOOB error string.
func (e ErrDecOOB) Error() string {
	return fmt.Sprintf("dec %d cannot be <0 or > %d", e.Got, maxDec)
}

// ErrDecimOOB error string.
func (e ErrDecimOOB) Error() string {
	return fmt.Sprintf("decim %d cannot be <0 or > %d", e.Got, maxDecim)
}

// ErrDesekOOB error string.
func (e ErrDesekOOB) Error() string {
	return fmt.Sprintf("desek %d cannot be <0 or > %d", e.Got, maxDesek)
}

// Dec returns the dec from a DIT.
func (d *DIT) Dec() Dec {
	return Dec(math.Floor(float64(*d) / 10000))
}

// Decim returns the decim from a DIT.
func (d *DIT) Decim() Decim {
	return Decim(math.Floor(float64(*d)/100 - float64(d.Dec()*100)))
}

// Desek returns the desek from a DIT.
func (d *DIT) Desek() Desek {
	return Desek(float64(*d) - (float64(d.Dec()) * 10000) - (float64(d.Decim() * 100)))
}

// String represents a DIT in DD.MM.SS notation.
func (d DIT) String() string {
	return fmt.Sprintf("%d.%02d.%02d", d.Dec(), d.Decim(), d.Desek())
}

// TimeToDIT coverts a time.Time into DIT.
func TimeToDIT(t time.Time) DIT {
	// time at dateline
	tdl := t.In(dateline)
	todaySeconds := tdl.Hour()*hourSeconds + tdl.Minute()*minuteSeconds + tdl.Second()
	return DIT(math.Round(float64(todaySeconds) / .864))
}

// NewDIT creates a DIT from the provided dec, decim, and desek.
func NewDIT(dec, decim, desek int) (DIT, error) {
	if dec > maxDec || dec < 0 {
		return 0, ErrDecOOB{dec}
	}
	if decim > maxDecim || decim < 0 {
		return 0, ErrDecimOOB{decim}
	}
	if desek > maxDesek || desek < 0 {
		return 0, ErrDesekOOB{desek}
	}
	return DIT(int(dec*10000) + int(decim*100) + int(desek)), nil
}
