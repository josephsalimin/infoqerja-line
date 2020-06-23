package constant

// Message constant for reply text from line bot
const (
	HelpMessage = `Use command below to use InfoQerja functionality:
	- !help		: to find out how to use InfoQerja
	- !add		: to add job posting to InfoQerja
	- !show		: to show job posting in InfoQerja`
	HelpShortMessage = `Please click button below to refer to available command`
	WelcomeMessage   = `Welcome to the InfoQerja Bot!!!💻💻`
	UnWelcomeMessage = `Please contact us for future improvement!!`
	InvalidMessage   = `Please enter a valid command!! Refer to !help for available command.`
	UnknownMessage   = `This bot does not respond to other things except for command 😎😎
	Please refer to !help command to use InfoQerja functionality.
	Hope you enjoy this bot 😊😊
	- Joseph Salimin 😍`
	AddMessage         = `Please add a new job 😍`
	AddTitleMessage    = `Please add a new job title😍`
	AddDescMessage     = `Please add the job description😍`
	AddDateMessage     = `Pick job deadline date😍`
	ShowMessageFail    = `Unable to show job listing saved. Please try again or contant the developer!!`
	UnavailableMessage = `Please view this in Mobile Version`
	ThankYouMessage    = `Thank you for adding job data 😎😎`
)

// Command code constant for refering to open public functionality for user to use
const (
	WelcomeCommandCode   = "!welcome--"
	UnWelcomeCommandCode = "!unwelcome--"
	ShowCommandCode      = "!show"
	HelpCommandCode      = "!help"
	AddCommandCode       = "!add"
)

// Image Url constant configuration
const (
	SearchImageURL = "https://img.icons8.com/material-outlined/1000/000000/advanced-search.png"
	AlertImageURL  = "https://img.icons8.com/material-outlined/1000/000000/warning-shield.png"
	ResumeImageURL = "https://img.icons8.com/material-outlined/1000/000000/parse-resume.png"
	ClockImageURL  = "https://img.icons8.com/material-outlined/1000/000000/clock.png"
)

// Data definition for postback action
const (
	DateData  = "DATE"
	JobIDData = "JOB-ID"
)
