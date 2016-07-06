package errors

// Authentication-related errors
var (
	LoginRequired      = &Error{401, "You must be logged in to perform this action."}
	InvalidCredentials = &Error{401, "Invalid Credentials"}
	BadUserType        = &Error{403, "You do not have access to this feature."}

	PasswordsDoNotMatch = &Error{400, "Passwords Do Not Match"}
	UserExists          = &Error{400, "User already exists"}
)

// 400 errors
var (
	BadParameters     = &Error{400, "Bad Parameters"}
	MissingParameters = &Error{400, "Missing Parameters"}
	TagExists         = &Error{400, "Tag already exists"}
)

// 403 errors
var (
	NotNoteOwner       = &Error{403, "You are not the owner of this note"}
	NotCollectionOwner = &Error{403, "You are not the owner of this collection"}
	NotTextbookOwner   = &Error{403, "You are not the owner of this textbook"}
)

// 404 errors
var (
	NoteNotFound       = &Error{404, "Note Not Found"}
	CollectionNotFound = &Error{404, "Collection Not Found"}
)

// 5xx errors
var (
	Debug = &Error{500, "teh internets are asplode"}
	JSON  = &Error{500, "Could not convert to JSON"}
)
