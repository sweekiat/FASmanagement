# FASManagementSystem
FASManagementSystem

# Please look at the submission where you will find a env.txt file with the credentials needed to get access to my aws rds database!!! Or contact me at sk.wong.2021@business.smu.edu.sg

### install all the imports

go mod tidy

### run the app (assuming u are in the main repo)

go run cmd/api/main.go

### assumptions

The sample data given doesnt show applicant with marital status, so I handled it in the create new applicant function.
The last user story was quite ambiguous but I implemented basic functionality to add a scheme to a user.
All inputs are valid. If given more time I would implement DTOs to ensure inputs are correct.

### other info

fasDB.sql is how I created my db as well as seeded it with the sample data
FASManagementSystem.postman_collection.json is the different requests I use to ensure the functions are working correctly
