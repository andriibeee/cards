package internal

import (
	"testing"
	"time"
)

func assertTrue(tb testing.TB, val bool) {
	tb.Helper()
	if !val {
		tb.Fatalf("got false, expected true")
	}
}

func assertFalse(tb testing.TB, val bool) {
	tb.Helper()
	if val {
		tb.Fatalf("got true, expected false")
	}
}

func TestComboValidator(t *testing.T) {
	t.Run("Valid exp", func(t *testing.T) {
		assertFalse(t, HasCardExpired("12", "2028", time.Now()))
		assertFalse(t, HasCardExpired("12", "28", time.Now()))
		assertFalse(t, HasCardExpired("8", "24", time.Now()))
	})

	t.Run("Invalid exp", func(t *testing.T) {
		assertTrue(t, HasCardExpired("", "", time.Now()))
		assertTrue(t, HasCardExpired("1", "24", time.Now()))
		assertTrue(t, HasCardExpired("12", "23", time.Now()))
		assertTrue(t, HasCardExpired("12", "2023", time.Now()))
		assertTrue(t, HasCardExpired("1", "2024", time.Now()))

	})
}

func TestMonthValidator(t *testing.T) {
	t.Run("Valid months", func(t *testing.T) {
		assertTrue(t, CheckCreditCardMonth("1"))
		assertTrue(t, CheckCreditCardMonth("2"))
		assertTrue(t, CheckCreditCardMonth("01"))
		assertTrue(t, CheckCreditCardMonth(" 02"))
		assertTrue(t, CheckCreditCardMonth("03"))
		assertTrue(t, CheckCreditCardMonth("04"))
		assertTrue(t, CheckCreditCardMonth("05"))
		assertTrue(t, CheckCreditCardMonth("09"))
		assertTrue(t, CheckCreditCardMonth("11"))
		assertTrue(t, CheckCreditCardMonth("12"))
	})

	t.Run("Invalid months", func(t *testing.T) {
		assertFalse(t, CheckCreditCardMonth(""))
		assertFalse(t, CheckCreditCardMonth("a"))
		assertFalse(t, CheckCreditCardMonth("13"))
		assertFalse(t, CheckCreditCardMonth("00"))
		assertFalse(t, CheckCreditCardMonth("0"))
	})
}

func TestNumberValidator(t *testing.T) {
	t.Run("Valid card", func(t *testing.T) {
		assertTrue(t, CheckCreditCardNumber("4111111111111111"))
		assertTrue(t, CheckCreditCardNumber("37144-96353-98431"))
		assertTrue(t, CheckCreditCardNumber("3766 808163 76961"))
		assertTrue(t, CheckCreditCardNumber("3625.960000.0004"))
		assertTrue(t, CheckCreditCardNumber("6304 000000 000000"))
		assertTrue(t, CheckCreditCardNumber("50635 169450 05047"))
		assertTrue(t, CheckCreditCardNumber("2223_00004_8400011"))
		assertTrue(t, CheckCreditCardNumber("4005519200000004"))
		assertTrue(t, CheckCreditCardNumber("4012000033330026"))
	})

	t.Run("Invalid card", func(t *testing.T) {
		assertFalse(t, CheckCreditCardNumber(""))
		assertFalse(t, CheckCreditCardNumber("test"))
		assertFalse(t, CheckCreditCardNumber("1111111111111"))
	})
}
