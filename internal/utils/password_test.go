package utils

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	md5Hash := "e10adc3949ba59abbe56e057f20f883e"

	hashedPassword, err := HashPassword(md5Hash)
	if err != nil {
		t.Fatalf("HashPassword failed: %v", err)
	}

	if len(hashedPassword) != 60 {
		t.Errorf("Expected hashed password length 60, got %d", len(hashedPassword))
	}

	if hashedPassword == md5Hash {
		t.Error("Hashed password should not equal MD5 hash")
	}
}

func TestHashPasswordInvalidInput(t *testing.T) {
	_, err := HashPassword("short")
	if err != ErrPasswordTooShort {
		t.Errorf("Expected ErrPasswordTooShort, got %v", err)
	}

	_, err = HashPassword("tooolongpasswordhashtest1234567890")
	if err != ErrPasswordTooShort {
		t.Errorf("Expected ErrPasswordTooShort, got %v", err)
	}
}

func TestVerifyPassword(t *testing.T) {
	md5Hash := "e10adc3949ba59abbe56e057f20f883e"

	hashedPassword, _ := HashPassword(md5Hash)

	if !VerifyPassword(md5Hash, hashedPassword) {
		t.Error("Password verification should succeed")
	}

	if VerifyPassword("wrongmd5hash12345678901234567890", hashedPassword) {
		t.Error("Wrong password should not verify")
	}
}

func TestVerifyPasswordInvalidInput(t *testing.T) {
	hashedPassword, _ := HashPassword("e10adc3949ba59abbe56e057f20f883e")

	if VerifyPassword("short", hashedPassword) {
		t.Error("Short password should not verify")
	}
}

func TestBcryptCost(t *testing.T) {
	md5Hash := "e10adc3949ba59abbe56e057f20f883e"

	hashedPassword, _ := HashPassword(md5Hash)

	cost, err := bcrypt.Cost([]byte(hashedPassword))
	if err != nil {
		t.Fatalf("Failed to get bcrypt cost: %v", err)
	}

	if cost != BcryptCost {
		t.Errorf("Expected cost %d, got %d", BcryptCost, cost)
	}
}

func TestNeedsRehash(t *testing.T) {
	md5Hash := "e10adc3949ba59abbe56e057f20f883e"

	hashedPassword, _ := HashPassword(md5Hash)

	if NeedsRehash(hashedPassword) {
		t.Error("Fresh hash should not need rehash")
	}

	lowCostHash, _ := bcrypt.GenerateFromPassword([]byte(md5Hash), 10)
	if !NeedsRehash(string(lowCostHash)) {
		t.Error("Low cost hash should need rehash")
	}
}

func TestDifferentPasswordsProduceDifferentHashes(t *testing.T) {
	md5Hash1 := "e10adc3949ba59abbe56e057f20f883e"
	md5Hash2 := "d41d8cd98f00b204e9800998ecf8427e"

	hash1, _ := HashPassword(md5Hash1)
	hash2, _ := HashPassword(md5Hash2)

	if hash1 == hash2 {
		t.Error("Different passwords should produce different hashes")
	}
}

func TestSamePasswordProducesDifferentHashes(t *testing.T) {
	md5Hash := "e10adc3949ba59abbe56e057f20f883e"

	hash1, _ := HashPassword(md5Hash)
	hash2, _ := HashPassword(md5Hash)

	if hash1 == hash2 {
		t.Error("Same password should produce different hashes due to salt")
	}

	if !VerifyPassword(md5Hash, hash1) {
		t.Error("First hash should verify")
	}

	if !VerifyPassword(md5Hash, hash2) {
		t.Error("Second hash should verify")
	}
}
