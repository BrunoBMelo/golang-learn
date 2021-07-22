package accounts

import "testing"

func TestBuildAccounObject(t *testing.T) {
	params := []struct {
		username string
		email    string
		password string
		result   bool
	}{
		{"Bruno B. de Melo", "bruno.b.melo@live.com", "@Xptoz94", true},
		{"", "", "", false},
		{"", "", "1", false},
	}

	for _, param := range params {
		result := CreateNewAccount(param.username, param.email, param.password)

		if result != param.result {
			t.Errorf("[Failed]: TestBuildAccounObject\n[Exception]: %t\n[Expected]: %t\n[Data]: %v", result, param.result, param)
		}
	}
}
