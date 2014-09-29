package astrotime

import (
	"testing"
	"time"
)

// Fixed location used for testing.
const TEST_LAT = 14
const TEST_LONG = 100
const TEST_ZONE = "Asia/Bangkok"

func compareApproxTimes(value, expected time.Time) bool {
	approx := time.Minute * 1

	return value.After(expected.Add(-1*approx)) && value.Before(expected.Add(approx))
}

func TestCalcSunrise(t *testing.T) {
	location, _ := time.LoadLocation(TEST_ZONE)
	expectedSunrise := time.Date(2014, time.June, 12, 5, 52, 0, 0, location)

	// Test well before sunrise.
	now := time.Date(2014, time.June, 12, 2, 57, 12, 0, location)
	sunrise := CalcSunrise(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunrise, expectedSunrise) {
		t.Error("Unexpected sunrise result:", sunrise, ", expected: ", expectedSunrise)
	}

	// Test 1 minute after sunrise.
	now = expectedSunrise.Add(time.Minute)
	sunrise = CalcSunrise(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunrise, expectedSunrise) {
		t.Error("Unexpected sunrise result:", sunrise, ", expected: ", expectedSunrise)
	}

	// Test well after sunrise.
	now = time.Date(2014, time.June, 12, 10, 57, 12, 0, location)
	sunrise = CalcSunrise(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunrise, expectedSunrise) {
		t.Error("Unexpected sunrise result:", sunrise, ", expected: ", expectedSunrise)
	}

	// Test late in day.
	now = time.Date(2014, time.June, 12, 22, 57, 12, 0, location)
	sunrise = CalcSunrise(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunrise, expectedSunrise) {
		t.Error("Unexpected sunrise result:", sunrise, ", expected: ", expectedSunrise)
	}
}

func TestNextSunrise(t *testing.T) {
	location, _ := time.LoadLocation(TEST_ZONE)
	expectedSunriseToday := time.Date(2014, time.June, 12, 5, 51, 35, 0, location)
	expectedSunriseTomorrow := time.Date(2014, time.June, 13, 5, 51, 42, 0, location)

	// Test well before sunrise.
	now := time.Date(2014, time.June, 12, 2, 57, 12, 0, location)
	sunrise := NextSunrise(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunrise, expectedSunriseToday) {
		t.Error("[1] Unexpected sunrise result:", sunrise, ", expected: ", expectedSunriseToday)
	}

	// Test 1 minute before sunrise.
	now = expectedSunriseToday.Add(-1 * time.Minute)
	sunrise = NextSunrise(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunrise, expectedSunriseToday) {
		t.Error("[2] Unexpected sunrise result:", sunrise, ", expected: ", expectedSunriseToday)
	}

	// Test 1 minute after sunrise.
	now = expectedSunriseToday.Add(time.Minute)
	sunrise = NextSunrise(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunrise, expectedSunriseTomorrow) {
		t.Error("[3] Unexpected sunrise result:", sunrise, ", expected: ", expectedSunriseTomorrow)
	}

	// Test well after sunrise.
	now = time.Date(2014, time.June, 12, 10, 57, 12, 0, location)
	sunrise = NextSunrise(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunrise, expectedSunriseTomorrow) {
		t.Error("[4] Unexpected sunrise result:", sunrise, ", expected: ", expectedSunriseTomorrow)
	}

	// Test late in day.
	now = time.Date(2014, time.June, 12, 22, 57, 12, 0, location)
	sunrise = NextSunrise(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunrise, expectedSunriseTomorrow) {
		t.Error("[5] Unexpected sunrise result:", sunrise, ", expected: ", expectedSunriseTomorrow)
	}
}

func TestCalcSunset(t *testing.T) {
	location, _ := time.LoadLocation(TEST_ZONE)
	expectedSunset := time.Date(2014, time.June, 12, 18, 48, 0, 0, location)

	// Test well before sunset.
	now := time.Date(2014, time.June, 12, 2, 57, 12, 0, location)
	sunset := CalcSunset(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunset, expectedSunset) {
		t.Error("Unexpected sunset result:", sunset, ", expected: ", expectedSunset)
	}

	// Test 1 minute after sunset.
	now = expectedSunset.Add(time.Minute)
	sunset = CalcSunset(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunset, expectedSunset) {
		t.Error("Unexpected sunset result:", sunset, ", expected: ", expectedSunset)
	}

	// Test well after sunset.
	now = time.Date(2014, time.June, 12, 10, 57, 12, 0, location)
	sunset = CalcSunset(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunset, expectedSunset) {
		t.Error("Unexpected sunset result:", sunset, ", expected: ", expectedSunset)
	}

	// Test late in day.
	now = time.Date(2014, time.June, 12, 22, 57, 12, 0, location)
	sunset = CalcSunset(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunset, expectedSunset) {
		t.Error("Unexpected sunset result:", sunset, ", expected: ", expectedSunset)
	}
}

func TestNextSunset(t *testing.T) {

	location, _ := time.LoadLocation(TEST_ZONE)
	expectedSunsetToday := time.Date(2014, time.June, 12, 18, 48, 12, 30, location)
	expectedSunsetTomorrow := time.Date(2014, time.June, 13, 18, 48, 23, 0, location)

	// Test well before sunset.
	now := time.Date(2014, time.June, 12, 2, 57, 12, 0, location)
	sunset := NextSunset(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunset, expectedSunsetToday) {
		t.Error("[1] Unexpected sunset result:", sunset, ", expected: ", expectedSunsetToday)
	}

	// Test 1 minute before sunset.
	now = expectedSunsetToday.Add(-1 * time.Minute)
	sunset = NextSunset(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunset, expectedSunsetToday) {
		t.Error("[2] Unexpected sunset result:", sunset, ", expected: ", expectedSunsetToday)
	}

	// Test 1 minute after sunset.
	now = expectedSunsetToday.Add(time.Minute)
	sunset = NextSunset(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunset, expectedSunsetTomorrow) {
		t.Error("[3] Unexpected sunset result:", sunset, ", expected: ", expectedSunsetTomorrow)
	}

	// Test well after sunset.
	now = time.Date(2014, time.June, 12, 22, 57, 12, 0, location)
	sunset = NextSunset(now, TEST_LAT, TEST_LONG)
	if !compareApproxTimes(sunset, expectedSunsetTomorrow) {
		t.Error("[4] Unexpected sunset result:", sunset, ", expected: ", expectedSunsetTomorrow)
	}
}
