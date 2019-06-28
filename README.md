# class-review-backend

# Go Setup
1. Download and install go
2. Figure out your GOPATH
3. cd to $GOPATH/src amd clone this repo there
4. cd into class-reviews
5. Install govendor `go get -u github.com/kardianos/govendor`
6. run 'govendor sync' (You may have to reset terminal on mac if you get govendor command not found)



# Local Database (MySQL) Setup
Install MySQL Server
https://dev.mysql.com/downloads/mysql/


(MySQL Workbench or Shell) 
https://dev.mysql.com/downloads/workbench/

1. Open MySQL Workbench
2. Click your local connection and authenticate using the root password you used when setting up the MySQL server.
3. Go to query tab and run `CREATE SCHEMA 'classreviews';`
4. Create a new user under users and privileges with login name 'classreviews' and a password to remember later make sure authentication type is `Standard`.
6. Go to schema privileges tab and add entry and select classreviews
7. Select all privileges and apply to classreviews

# Env Setup
1. Create a new file name .env in the `class-review-backend/`
2. Add this as the content but replace password with the password you created for your MySQL User

```DB_NAME=classreviews
DB_USER=classreviews
DB_PASSWORD=password
DB_SERVER=127.0.0.1
DB_PORT=3306
```

# Running
1. Run `go run main.go`
2. Navigate to 'http://127.0.0.1:8080/' 
