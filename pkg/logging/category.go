package logging

type Category string
type SubCategory string
type ExtraKey string

const (
	General    Category = "General"
	Internal   Category = "Internal"
	Validation Category = "Validation"
	Database   Category = "Database"
	Cache      Category = "Cache"
	File       Category = "File"
)

const (
	//General
	StartUp         SubCategory = "StartUp"
	ExternalService SubCategory = "ExternalService"

	//Internal
	API                 SubCategory = "API"
	HashPassword        SubCategory = "HashPassword"
	DefaultRoleNotFound SubCategory = "DefaultRoleNotFound"

	//Validation
	MobileNumberValidation SubCategory = "MobileNumberValidation"
	PasswordValidation     SubCategory = "PasswordValidation"
	EmailValidation        SubCategory = "EmailValidation"

	//Database
	SelectQuery         SubCategory = "SelectQuery"
	InsertQuery         SubCategory = "InsertQuery"
	UpdateQuery         SubCategory = "UpdateQuery"
	DeleteQuery         SubCategory = "DeleteQuery"
	RollbackTransaction SubCategory = "RollbackTransaction"
	CommitTransaction   SubCategory = "CommitTransaction"

	//Cache
	GetCache SubCategory = "GetCache"
	SetCache SubCategory = "SetCache"

	//File
	ReadFile  SubCategory = "ReadFile"
	WriteFile SubCategory = "WriteFile"
)

const (
	AppName      ExtraKey = "AppName"
	LoggerName   ExtraKey = "LoggerName"
	ClientIP     ExtraKey = "ClientIP"
	HostIP       ExtraKey = "HostIP"
	Method       ExtraKey = "Method"
	Path         ExtraKey = "Path"
	StatusCode   ExtraKey = "StatusCode"
	ErrorMessage ExtraKey = "ErrorMessage"
	BodySize     ExtraKey = "BodySize"
	Latency      ExtraKey = "Latency"
)
