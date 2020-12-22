# kogito-operator

POC for [KOGITO-3951](https://issues.redhat.com/browse/KOGITO-3951) to implement a sub-module approach.
In this POC three Go module has been implemented.
* kogito-operator
 * community-kogito-operator
 * product-kogito-operator

##### kogito-operator
Kogito-operator is the parent Go module.

##### community-kogito-operator
community-kogito-operator refers to the community version of Kogito Operator. It primarily consist of three component
* Operator-sdk
* Internal services
* Core

###### Operator-sdk component
This component has been generated using operator-sdk api's.
```shell script
operator-sdk init --domain=vajain.com --repo=github.com/vaibhavjainwiz/kogito-operator
operator-sdk create api --group cache --version app --kind KogitoRuntime --resource=true --controller=true
```

###### Internal services
Internal service package is used to put those functions which involve direct usage of struts exposed by operator-sdk.
In POC, the function `FetchKogitoRuntimeService` has been put into an internal package because it directly uses the `KogitoRuntime` strut.

###### Core
Apart from `operator-sdk` and `internal` everything will go to the `core` package. Methods and apis in the `core` package will be shared with `product-kogito-operator` as well.
In POC, function `SetupRBAC` is placed in core package because it's not requires struct created by `operator-sdk` and it also need to be shared with `product-kogito-operator`

##### product-kogito-operator
product-kogito-operator refers to the product version of Kogito Operator. It primarily consist of two component
* Operator-sdk
* Internal services

`product-kogito-operator` import `community-kogito-operator` as Go module dependency. It uses that dependency to call its `core` package services and apis. 
In POC, `product-kogito-operator` is called `SetupRBAC` function which exists in `community-kogito-operator`.

##Releases

Both `product-kogito-operator` and `community-kogito-operator` released separately.
In POC, 2 tags `v0.0.1` and `v0.0.3` has been release for `community-kogito-operator` whereas just 1 tag `v0.0.1` of `product-kogito-operator` released which refer to `v0.0.3` of `community-kogito-operator` 
