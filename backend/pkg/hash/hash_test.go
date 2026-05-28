package hash

import "testing"

func TestHashPasswordRoundTrip(t *testing.T) {
	password := "S3cret!Parola"

	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword hata verdi: %v", err)
	}
	if hashed == password {
		t.Fatal("Hash, açık parolayla aynı olmamalı")
	}
	if !CheckPassword(password, hashed) {
		t.Fatal("CheckPassword doğru parolayı reddetti")
	}
}

func TestCheckPasswordWrong(t *testing.T) {
	hashed, err := HashPassword("dogru")
	if err != nil {
		t.Fatalf("HashPassword hata verdi: %v", err)
	}
	if CheckPassword("yanlis", hashed) {
		t.Fatal("CheckPassword yanlış parolayı kabul etti")
	}
}
