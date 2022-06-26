package constants

type errorMessages struct {
	EmailOrPasswordIncorrect string
	UserAlreadyExists        string
	UserDoesNotExists        string
}

var Errormessages errorMessages = errorMessages{
	EmailOrPasswordIncorrect: "Email or password incorrect",
	UserAlreadyExists:        "User already exists",
	UserDoesNotExists:        "User does not exists",
}
