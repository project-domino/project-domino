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
	InvalidPage       = &Error{400, "Page number is not valid"}
	InvalidItems      = &Error{400, "Item count is not valid"}
	CommentNesting    = &Error{400, "A comment can only be nested one level deep"}
)

// 403 errors
var (
	NotNoteOwner       = &Error{403, "You are not the owner of this note"}
	NotCollectionOwner = &Error{403, "You are not the owner of this collection"}
	NotTextbookOwner   = &Error{403, "You are not the owner of this textbook"}
	EmailNotVerified   = &Error{403, "Your email is not verified"}
)

// 404 errors
var (
	NotFound           = &Error{404, "Page Not Found"}
	NoteNotFound       = &Error{404, "Note Not Found"}
	CollectionNotFound = &Error{404, "Collection Not Found"}
	UserNotFound       = &Error{404, "User Not Found"}
	CommentNotFound    = &Error{404, "Comment Not Found"}
	ResetCodeNotFound  = &Error{404, "Reset Code Not Found"}
)

// 5xx errors
var (
	InternalError = &Error{500, "Server Error"}
	DB            = &Error{500, "Database Error"}
	Debug         = &Error{500, "teh internets are asplode"}
	JSON          = &Error{500, "Could not convert to JSON"}
)
