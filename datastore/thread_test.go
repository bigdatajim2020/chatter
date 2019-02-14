package datastore

import "testing"

func Test_CreateThread(t *testing.T) {
	defer teardown()
	if err := users[0].New(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	thread, err := users[0].NewThread("My first thread")
	if err != nil {
		t.Error(err, "Cannot create thread")
	}
	if thread.UserID != users[0].ID {
		t.Error("User not linked with thread")
	}
}
