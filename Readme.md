# Bills FarGO
## Godfrey Macasero
Bills FarGO is a banking app that is written in Golang. It uses docker and runs a postgres database using docker. You are able to log in as two different users which are customers and employees.
This app simulates banking on the terminal that has interactive text menus. It is a simple banking application that emulates things a customer would be able to do at the bank like withdrawing, depositing, and transfering money to your different accounts. The employees on the other hand can view customer information, and approve or deny applications for accounts.

# User Stories
- [x] Customers should be able to:
    - [x] Register with a username and password
    - [x] Authenticate (login) with that usename and password
    - [x] Apply to open an account
    - [x] Apply for joint accounts with other customers
    - [x] Withdraw, deposit, and transfer funds between accounts
- [x] Employees should be able to:
    - [x] View customer information
    - [x] Approve/deny open applications for accounts
    - [x] Register with a username and password
    - [x] Authenticate (login) with that usename and password
- [x] All account and user information should persist to files or a database
- [x] Usable from command line args, interactive text menus, or through HTTP
- [x] Basic validation and testing


# Instructions
Cloning from git.

Make sure you have a file structure in your go/src with a directory stucture shown below.
```
.
├── github.com
    └── gmac220
```
Then call git clone inside the ~/go/src/github.com/gmac220 directory

Make sure to have docker installed in your terminal type
```bash
sudo apt search docker
sudo apt install docker
sudo apt install docker.io
```
Have postgres image installed with docker & driver
```bash
sudo docker run postgres
go get -u github.com/lib/pq
```
Inside project file go into the database folder where Dockerfile is run
```bash
cd db
docker build -t billsfargo .
docker image ls
docker run -p 5432:5432 -d --rm --name runningbillsfargo billsfargo
```

**OPTIONAL COMMAND**

If you want to look into your table in postgres use this command
```bash
docker exec -it runningbillsfargo psql -U postgres
```

# Running the program
```bash
go run main.go
```

# Testing
When testing make sure your container is new and nothing has been added to the database previously.
Testing commands for navigation package. Optional -v for detailed run
```bash
go test ./navigation
go test -v ./navigation/
```
Need to figure out how to run only one file, but until then this is how I test the customer package.
Testing commands for customer1_test.go
```bash
go test -v ./customer/ -run SetCustomerVars
go test -v ./customer/ -run Apply
go test -v ./customer/ -run JointApp
go test -v ./customer/ -run CheckCustomer
```

Testing commands for employees package
```bash
go test ./employees
```

Testing commands for customer2_test.go
```bash
go test -v ./customer/ -run VerifyAccount
go test -v ./customer/ -run CheckOwnAccount
go test -v ./customer/ -run Deposit
go test -v ./customer/ -run ShowBalance
go test -v ./customer/ -run Withdraw
go test -v ./customer/ -run Transfer
```
### [Presentation](https://gitpitch.com/gmac220/project-0/master)

go test -timeout 30s github.com/gmac220/project-0/employees -coverprofile=/tmp/vscode-goKcYvzT/go-code-cover