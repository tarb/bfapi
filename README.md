# BFAPI - Betfair Api

Implementation of Betfair's Exchange, Account and Stream Api


#### Logging in

There are two ways to login and you must choose one of them before calling any other functions, but both login method acts predominantly the same way. For non-interactive logins (logins from automated apps where a human is not present during the login process) Betfair recommend using the CertLogin method. This method requires you created and upload a certificate through which you will encrypt your traffic.
```go
if token, err := bfapi.CertLogin(user, pass, appKey, certFile, keyFile); err != nil {
    log.Fatalf("Failed to login: %v", err)
}
```
Where users have provided you the username and password dynamically Betfair recommend the simpler interactive Login method.
```go
if token, err := bfapi.Login(user, pass, appkey); err != nil {
    log.Fatalf("Failed to login: %v", err)
}
```
The login expires every 3 hours but KeepAlive can be called to extend this time (by another 3 hours). KeepAlive will block for the supplied duration and is best called from a go routine.
```go
go bfapi.KeepAlive(3 * time.Hour)
```

#### Methods

ListMarketCatalogue can be used to retrieve a broad overview of a selection of markets

Example of ListMarketCatalogue from outside package
```go
arg := bfapi.ListMarketCatalogueArg{
    Filter: MarketFilter{
        EventTypes: []string{"4339"}, // greyhounds id code
        Countries:  []string{"GB"},
        TypeCodes:  []string{"WIN"},
    },
    Sort:             "FIRST_TO_START",
    MarketProjection: []string{"MARKET_START_TIME", "RUNNER_DESCRIPTION", "EVENT"},
    MaxResults:       1000,
}

result, err := bfapi.ListMarketCatalogue(arg)
```

### TODO

* Documentation - Function/TypeDef comments
* Add more Betfair methods and types
