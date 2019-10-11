package navigation

import "testing"

func TestWelcome(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Welcome()
		})
	}
}

func TestSelection(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Selection()
		})
	}
}

func TestSignIn(t *testing.T) {
	type args struct {
		username string
		password string
		employee bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SignIn(tt.args.username, tt.args.password, tt.args.employee)
		})
	}
}

func TestCreateAccount(t *testing.T) {
	type args struct {
		firstname string
		lastname  string
		username  string
		password  string
		choice    string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateAccount(tt.args.firstname, tt.args.lastname, tt.args.username, tt.args.password, tt.args.choice)
		})
	}
}

func TestSttyCommand(t *testing.T) {
	type args struct {
		flag string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SttyCommand(tt.args.flag)
		})
	}
}
