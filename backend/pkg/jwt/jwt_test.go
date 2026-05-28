package jwt

import "testing"

const testSecret = "test-secret-key"

func TestAccessTokenRoundTrip(t *testing.T) {
	token, err := GenerateAccessToken("user-123", "ali@example.com", testSecret)
	if err != nil {
		t.Fatalf("GenerateAccessToken hata: %v", err)
	}

	claims, err := ValidateToken(token, testSecret)
	if err != nil {
		t.Fatalf("ValidateToken hata: %v", err)
	}
	if claims.UserID != "user-123" {
		t.Errorf("UserID beklenen 'user-123', alınan '%s'", claims.UserID)
	}
	if claims.Email != "ali@example.com" {
		t.Errorf("Email beklenen 'ali@example.com', alınan '%s'", claims.Email)
	}
}

func TestValidateTokenWrongSecret(t *testing.T) {
	token, err := GenerateAccessToken("user-123", "ali@example.com", testSecret)
	if err != nil {
		t.Fatalf("GenerateAccessToken hata: %v", err)
	}

	if _, err := ValidateToken(token, "yanlis-secret"); err == nil {
		t.Fatal("Yanlış secret ile token doğrulanmamalıydı")
	}
}

func TestRefreshTokenRoundTrip(t *testing.T) {
	token, err := GenerateRefreshToken("user-456", "ayse@example.com", testSecret)
	if err != nil {
		t.Fatalf("GenerateRefreshToken hata: %v", err)
	}

	claims, err := ValidateToken(token, testSecret)
	if err != nil {
		t.Fatalf("ValidateToken hata: %v", err)
	}
	if claims.UserID != "user-456" {
		t.Errorf("UserID beklenen 'user-456', alınan '%s'", claims.UserID)
	}
}
