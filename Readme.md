# Bills FarGO
## Godfrey Macasero
A banking app that has two users customers and employees. 
This app simulates banking on the terminal that has interactive text menus.
It is a simple banking application that emulates things you would be able to do at the bank like withdrawing, depositing, and transfering money to your different accounts.

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

This command is Optional
If you want to look into your table in postgres use this command
```bash
docker exec -it runningbillsfargo psql -U postgres
```

# Running the program
```bash
go run main.go
```

# Testing
Testing commands for navigation package
```bash
go test ./navigation
```

Testing commands for customer package
```bash
go test ./customer
```

Testing commands for employees package
```bash
go test ./employees
```

