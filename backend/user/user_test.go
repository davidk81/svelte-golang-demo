package user

import "testing"

func TestPasswordSalt(t *testing.T) {
	saltedpw, err := hashAndSaltPassword([]byte("test password"))
	if err != nil {
		t.Error(err)
	}
	err = verifyPassword(saltedpw, []byte("test password"))
	if err != nil {
		t.Errorf("same password failed compare test")
	}
	err = verifyPassword(saltedpw, []byte("wrong password"))
	if err == nil {
		t.Errorf("different password should fail compare test but did not")
	}
}
