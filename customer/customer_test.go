package customer

import "testing"

func TestSetCustomerVars(t *testing.T) {
	type args struct {
		username  string
		firstname string
		lastname  string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetCustomerVars(tt.args.username, tt.args.firstname, tt.args.lastname)
		})
	}
}

func TestShowCustomerPrompts(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShowCustomerPrompts()
		})
	}
}

func TestApply(t *testing.T) {
	type args struct {
		username  string
		firstname string
		lastname  string
		acntname  string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Apply(tt.args.username, tt.args.firstname, tt.args.lastname, tt.args.acntname)
		})
	}
}

func TestJointApp(t *testing.T) {
	type args struct {
		username  string
		firstname string
		lastname  string
		acntname  string
		username2 string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JointApp(tt.args.username, tt.args.firstname, tt.args.lastname, tt.args.acntname, tt.args.username2)
		})
	}
}

func TestCheckCustomer(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
		want2 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := CheckCustomer(tt.args.username)
			if got != tt.want {
				t.Errorf("CheckCustomer() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CheckCustomer() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("CheckCustomer() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestShowAccounts(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShowAccounts(tt.args.username)
		})
	}
}

func TestCheckOwnAccount(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckOwnAccount(tt.args.num); got != tt.want {
				t.Errorf("CheckOwnAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShowBalance(t *testing.T) {
	type args struct {
		acntnum int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShowBalance(tt.args.acntnum)
		})
	}
}

func TestVerifyAccount(t *testing.T) {
	type args struct {
		accountnumber int
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := VerifyAccount(tt.args.accountnumber)
			if got != tt.want {
				t.Errorf("VerifyAccount() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("VerifyAccount() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWithdraw(t *testing.T) {
	type args struct {
		acntnum    int
		withdrawal float64
		balance    float64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Withdraw(tt.args.acntnum, tt.args.withdrawal, tt.args.balance)
		})
	}
}

func TestDeposit(t *testing.T) {
	type args struct {
		acntnum int
		deposit float64
		balance float64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Deposit(tt.args.acntnum, tt.args.deposit, tt.args.balance)
		})
	}
}

func TestShowPendingApps(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShowPendingApps(tt.args.username)
		})
	}
}

func TestTransfer(t *testing.T) {
	type args struct {
		acntnumwithdraw int
		acntnumdeposit  int
		balanceinput    float64
		fundsinput      float64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Transfer(tt.args.acntnumwithdraw, tt.args.acntnumdeposit, tt.args.balanceinput, tt.args.fundsinput)
		})
	}
}
