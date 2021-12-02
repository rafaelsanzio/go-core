package errs

var ErrFmt = "err: [%v]"

// common
var (
	ErrLoadingTimeZone       = _new("CMN000", "error loading timezone data")
	ErrMarshalingJson        = _new("CMN001", "error marshaling json")
	ErrUnmarshalingJson      = _new("CMN002", "error unmarshaling json")
	ErrParsingTime           = _new("CMN003", "error parsing time")
	ErrNoEntityIdProvided    = _new("CMN004", "entity ID is required but none was provided")
	ErrNoDateProvided        = _new("CMN005", "error no date provided")
	ErrNoPayloadData         = _new("CMN006", "error event contains no payload data")
	ErrRepoMockAction        = _new("CMN007", "error repo mock action")
	ErrUnknownErrorType      = _new("CMN008", "error unknown error type")
	ErrInvalidDate           = _new("CMN009", "error invalid date format")
	ErrConvertingStringToInt = _new("CMN010", "error converting string to int")
)

// pkg/api
var (
	ErrResponseWriter = _new("API000", "error writing to response writer")
)

// pkg/config
var (
	ErrCreatingParamStore    = _new("CFG002", "unable to create param store service")
	ErrUnknownConfigProvider = _new("CFG004", "error unknown config provider")
)

// pkg/store
var (
	ErrCursor               = _new("STR000", "error using cursor")
	ErrDecodeCursor         = _new("STR001", "unable to decode cursor value to non-pointer variable")
	ErrClosingCursor        = _new("STR002", "error closing cursor")
	ErrMongoConnect         = _new("STR003", "error connecting to mongo")
	ErrMongoInsertOne       = _new("STR004", "error inserting one mongo document")
	ErrMongoFindOne         = _new("STR005", "error finding one mongo document")
	ErrMongoFind            = _new("STR006", "error finding mongo document(s)")
	ErrMongoUpdateOne       = _new("STR007", "error updating one mongo document")
	ErrMongoDeleteOne       = _new("STR008", "error deleting one mongo document")
	ErrNotDocumentInterface = _new("STR009", "cannot insert value that doesn't implement Document interface")
	ErrDecodingInsertedId   = _new("STR010", "error decoding inserted ID")
	ErrParseRequestURI      = _new("STR011", "error parsing request uri")
	ErrMarshalingBson       = _new("STR012", "error marshaling bson")
	ErrUnmarshalingBson     = _new("STR013", "error unmarshaling bson")
)
