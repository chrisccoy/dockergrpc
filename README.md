
# amp-client-lib
This is a GO library intended to provide access to the AMP service endpoints in an easy and repeatable way. This library will be leveraged by both amp-system-test and amp (CLI). GO was chosen for two main reasons: first, since it is a completely separate language from Node, the implementation will require that amp services are accessed in the same manner that an external application development team would be required; second, GO is a strongly typed language, allowing for system tests to quickly identify if critical interfaces have changed unexpectedly. Modules containing data structures will be maintained to act as service contracts between the consumer and provider.

## Approach
From a design perspective, the goal of amp-client-lib is to provide an easy developer experience. Underlying concerns, such as managing headers, marshaling JSON, setting up request bodies, etc will be handled on behalf of a helper module, allowing the developer to focus on interacting with the services and their respective responses.

```
	result := ampsvc.NewAmpServiceProvider().GetAmpStatus()
```
***OR***

If you intend to reuse the Service Provider for multiple calls

```
    provider := ampsvc.NewAmpServiceProvider()
    result := provider.GetAmpStatus()
    nextResult := provider.startAmpService(someService)
```
## Concepts

**AmpServiceProvider**

Module that exposes all of the available AMP services as functions to the developer

```
	result := ampsvc.NewAmpServiceProvider().startAmpService(serviceId)
```


**AmpServiceHelper**

Module that wraps internal concerns around how to prepare requests, call services, marshal data, etc

**Data Responses**

This package contains all of the typed structures representing a service response from AMP. These structures act as the binding contract between the consumer and producer.

```
type AmpStatusResponse struct {
	Name   string `json:"name"`
	Id     uint   `json:"id"`
	Status string `json:"status"`
}
```


## Package Layout

```
ampsvc
|
|
+---client
    |
    |
    +----data
```

### ampsvc
This package will contain all available AMP services via **AmpServiceProvider** and is responsible for preparing and delegating calls to the implemented function defined with the client package.

### client
This package contains the main helper module **AmpServiceHelper** that is prepared internally by the **AmpServiceProvider** and used by the service request module.

### data
This package contains all of the response types from the service in the form of GO typed structures with built-in JSON support. Some services may actually contain a structure to represent the request type, depending on complexity of the request.

## Running
The current test sample is setup to run against amp-status "/status". In order to test the code, make sure at a minimum the "/status" service is reachable. In addition, make sure AMP_URL is visible to the running process via export or as follows:

```
AMP_URL=http://localhost:32778 go test
```
