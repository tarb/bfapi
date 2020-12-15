package bfapi

//Login Status values
const (
	StatusSuccess    string = "SUCCESS"
	StatusLimited    string = "LIMITED_ACCESS"
	StatusRestricted string = "LOGIN_RESTRICTED"
	StatusFail       string = "FAIL"
)

// LoginError contains the error string returned when login
// is not a success
type LoginError string

// Error - implementing the error interface
func (le LoginError) Error() string { return string(le) }

// Login Errors
var (
	ErrInvalidUsernameOrPassword            LoginError = "INVALID_USERNAME_OR_PASSWORD"            // the username or password are invalid
	ErrAccountNowLocked                     LoginError = "ACCOUNT_NOW_LOCKED"                      // the account was just locked
	ErrAccountAlreadyLocked                 LoginError = "ACCOUNT_ALREADY_LOCKED"                  // the account is already locked
	ErrPendingAuth                          LoginError = "PENDING_AUTH"                            // pending authentication
	ErrTelbetTermsConditionsNa              LoginError = "TELBET_TERMS_CONDITIONS_NA"              // Telbet terms and conditions rejected
	ErrDuplicateCards                       LoginError = "DUPLICATE_CARDS"                         // duplicate cards
	ErrSecurityQuestionsWrong               LoginError = "SECURITY_QUESTION_WRONG_3X"              // the user has entered wrong the security answer 3 times
	ErrKYCSuspended                         LoginError = "KYC_SUSPEND"                             // KYC suspended
	ErrSuspended                            LoginError = "SUSPENDED"                               // the account is suspended
	ErrClosed                               LoginError = "CLOSED"                                  // the account is closed
	ErrSeldExcluded                         LoginError = "SELF_EXCLUDED"                           // the account has been self-excluded
	ErrInvalidConnToRegulatorDK             LoginError = "INVALID_CONNECTIVITY_TO_REGULATOR_DK"    // the DK regulator cannot be accessed due to some internal problems in the system behind or in at regulator; timeout cases included.
	ErrNotAuthorizedByRegulatorDK           LoginError = "NOT_AUTHORIZED_BY_REGULATOR_DK"          // the user identified by the given credentials is not authorized in the DK's jurisdictions due to the regulators' policies. Ex: the user for which this session should be created is not allowed to act(play, bet) in the DK's jurisdiction.
	ErrInvalidConnToRegulatorIT             LoginError = "INVALID_CONNECTIVITY_TO_REGULATOR_IT"    // the IT regulator cannot be accessed due to some internal problems in the system behind or in at regulator; timeout cases included.
	ErrNotAuthorizedByRegulatorIT           LoginError = "NOT_AUTHORIZED_BY_REGULATOR_IT"          // the user identified by the given credentials is not authorized in the IT's jurisdictions due to the regulators' policies. Ex: the user for which this session should be created is not allowed to act(play, bet) in the IT's jurisdiction.
	ErrSecurityRestrictedLocation           LoginError = "SECURITY_RESTRICTED_LOCATION"            // the account is restricted due to security concerns
	ErrBettingRestrictedLocation            LoginError = "BETTING_RESTRICTED_LOCATION"             // the account is accessed from a location where betting is restricted
	ErrTradingMaster                        LoginError = "TRADING_MASTER"                          // Trading Master Account
	ErrTradingMasterSuspended               LoginError = "TRADING_MASTER_SUSPENDED"                // Suspended Trading Master Account
	ErrAgentClientMaster                    LoginError = "AGENT_CLIENT_MASTER"                     // Agent Client Master
	ErrAgentClientMasterSuspended           LoginError = "AGENT_CLIENT_MASTER_SUSPENDED"           // Suspended Agent Client Master
	ErrDanishAuthorizationRequired          LoginError = "DANISH_AUTHORIZATION_REQUIRED"           // Danish authorization required
	ErrSpainMigrationRequired               LoginError = "SPAIN_MIGRATION_REQUIRED"                // Spain migration required
	ErrDenmarkMigrationRequired             LoginError = "DENMARK_MIGRATION_REQUIRED"              // Denmark migration required
	ErrSpanishTermsAcceptanceRequired       LoginError = "SPANISH_TERMS_ACCEPTANCE_REQUIRED"       // The latest Spanish terms and conditions version must be accepted. You must login to the website to accept the new conditions.
	ErrItalianContractAcceptanceRequired    LoginError = "ITALIAN_CONTRACT_ACCEPTANCE_REQUIRED"    // The latest Italian contract version must be accepted. You must login to the website to accept the new conditions.
	ErrCertAuthRequired                     LoginError = "CERT_AUTH_REQUIRED"                      // Certificate required or certificate present but could not authenticate with it
	ErrChangePasswordRequired               LoginError = "CHANGE_PASSWORD_REQUIRED"                // Change password required
	ErrPersonalMessageRequired              LoginError = "PERSONAL_MESSAGE_REQUIRED"               // Personal message required for the user
	ErrInternationalTermsAcceptanceRequired LoginError = "INTERNATIONAL_TERMS_ACCEPTANCE_REQUIRED" // The latest international terms and conditions must be accepted prior to logging in.
	ErrEmailLoginNotAllowed                 LoginError = "EMAIL_LOGIN_NOT_ALLOWED"                 // This account has not opted in to log in with the email
	ErrMultipleUsersWithSameCredential      LoginError = "MULTIPLE_USERS_WITH_SAME_CREDENTIAL"     // There is more than one account with the same credential
	ErrAccountPendingPasswordChange         LoginError = "ACCOUNT_PENDING_PASSWORD_CHANGE"         // The account must undergo password recovery to reactivate via https://identitysso.betfair.com/view/recoverpassword
	ErrTemporaryBanTooManyRequests          LoginError = "TEMPORARY_BAN_TOO_MANY_REQUESTS"         // The limit for successful login requests per minute has been exceeded. New login attempts will be banned for 20 minutes
	ErrItalianProfilingAcceptanceRequired   LoginError = "ITALIAN_PROFILING_ACCEPTANCE_REQUIRED"   // You must login to the website to accept the new conditions
)

//
type CertLoginResult struct {
	SessionToken string `json:"sessionToken"`
	Status       string `json:"loginStatus"`
}

//
type LoginResult struct {
	SessionToken string `json:"token"`
	AppKey       string `json:"product"`
	Status       string `json:"status"`
	Error        string `json:"error"`
}
