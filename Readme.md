# Bills FarGO
## Godfrey Macasero
A banking app that has two users customers and employees. 
This app simulates banking on the terminal with CLI like commands that allows basic bank functions like withdrawing, depositing, and linking to your different accounts with what you have.

# User Stories
- [] Customers should be able to:
    - [x] Register with a username and password
    - [x] Authenticate (login) with that usename and password
    - [x] Apply to open an account
    - [x] Apply for joint accounts with other customers
    - [] Withdraw, deposit, and transfer funds between accounts
- [x] Employees should be able to:
    - [x] View customer information
    - [x] Approve/deny open applications for accounts
    - [x] Register with a username and password
    - [x] Authenticate (login) with that usename and password
- [x] All account and user information should persist to files or a database
- [] Usable from command line args, interactive text menus, or through HTTP
- [] Basic validation and testing


# Instructions
Insert environment, build, and execution documentation here.
Make sure to have docker installed
go get -u github.com/lib/pq
Inside project the database folder where Dockerfile is run
```bash
cd db
docker build -t billsfargo .
docker image ls
docker run -p 5432:5432 -d --rm --name runningbillsfargo billsfargo
docker exec -it runningbillsfargo psql -U postgres
```
