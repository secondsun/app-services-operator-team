
<a name="0.8.2"></a>
## [0.8.2](https://github.com/redhat-developer/app-services-operator/compare/0.8.1...0.8.2) (2021-08-23)

### Bug Fixes

* installing service registry


<a name="0.8.1"></a>
## [0.8.1](https://github.com/redhat-developer/app-services-operator/compare/0.8.0...0.8.1) (2021-08-23)

### Bug Fixes

* apdating apis and fixing tests

### Features

* schema registry added to cloud services request


<a name="0.8.0"></a>
## [0.8.0](https://github.com/redhat-developer/app-services-operator/compare/0.0.3...0.8.0) (2021-08-23)

### Bug Fixes

* creating string constants for KafkkaConnection labels
* /s/cloud.redhat.com/console.redhat.com/g

### Features

* team
* enforcing labels on KafaConnection
* schema registry added to cloud services request


<a name="0.0.3"></a>
## [0.0.3](https://github.com/redhat-developer/app-services-operator/compare/0.0.2...0.0.3) (2021-08-11)


<a name="0.0.2"></a>
## [0.0.2](https://github.com/redhat-developer/app-services-operator/compare/0.0.1...0.0.2) (2021-08-10)


<a name="0.0.1"></a>
## 0.0.1 (2021-08-10)

### BREAKING CHANGES

* Add metadata container for the KafkaConnection
* Replace resources to remove "managed" part. ([#122](https://github.com/redhat-developer/app-services-operator/issues/122))
* bootstrapServerHost as root value
* change API for managed services requests

### Bug Fixes

* /s/cloud.redhat.com/console.redhat.com/g ([#249](https://github.com/redhat-developer/app-services-operator/issues/249))
* using new operatorhub repository ([#248](https://github.com/redhat-developer/app-services-operator/issues/248))
* publishing openapi and fixing generation
* Remove invalid workflow
* add website build
* add missing README to quickstarts
* change extensively long operator name ([#209](https://github.com/redhat-developer/app-services-operator/issues/209))
* update open api spec ([#206](https://github.com/redhat-developer/app-services-operator/issues/206))
* name of the project and general information
* provide operator hub instllation docs ([#200](https://github.com/redhat-developer/app-services-operator/issues/200))
* add reference to service binding operator
* update readme with shorter into
* push tags for release process
* update to trigger builds only on release
* add logging capability to tell us what env is used by operator
* add explicit env variable pointing to the environment
* update default url to production ([#178](https://github.com/redhat-developer/app-services-operator/issues/178))
* update case for KafkaConnect unit tests ([#187](https://github.com/redhat-developer/app-services-operator/issues/187))
* url for the kafka in the UI
* change email in olm
* add information to description about limited beta
* improve error handling to include message from server
* add lowercase versions
* minor issue with OLM spec
* change references to the organisation
* add OAuth Bearer and SASL PLAIN binding support
* Add NPE protections to the operator with proper validation messsages
* userKafkas Update, removing broken equals impl
* catching uncaught exceptions and setting status on finished
* removing memory limits
* update olm descriptions based on the requirements
* improve error messages for error handling
* add missing  x-kubernetes-preserve-unknown-fields to service accounts request
* add extra 503 handling
* add extra 503 handling
* minor typo causing an NPE
* actually checking finished status now...
* catching wider error type
* catching wider error type
* add autoignore variables ([#137](https://github.com/redhat-developer/app-services-operator/issues/137))
* Add ability to preserve custom fields for status of every object.
* remove dekorate integration
* delete outdated api mock ([#127](https://github.com/redhat-developer/app-services-operator/issues/127))
* minor changes
* Adding resource limits to CSV
* updates
* uiRef got lost in merge and typo
* move from plaintext to plain sasl mechanism
* swap SASL values passed for the status
* delete unneeded folder
* order of the service binding definitions
* typo in bootstrapServerHost ([#95](https://github.com/redhat-developer/app-services-operator/issues/95))
* bash script inconsistency
* fix olm script for if statement ([#88](https://github.com/redhat-developer/app-services-operator/issues/88))
* imports in java code
* Update to the official quay images and references
* allow to release exact versions and split into prerelease and release flows
* delete invalid builds from master
*  all namespaces watched by default, removing
* update sdk
* minor changes
* backwards compatability for old type name
* tests were not running. Marking stubs as disabled
* renaming condition and forcing finished to false
* updating crds to v1 crd definition
* leave non breaking manifest on the top
* invalid abbreviation
* update readme comment
* rollback operator-sdk version
* Update binding spec
* fixing connection controller
* update binding spec
* update binding configuration
* add service binding annotations and additional fields ([#49](https://github.com/redhat-developer/app-services-operator/issues/49))
* remove reset field
* add contributing
* remove not needed examples folder
* add additional properties required by UI to the kafka list and change format to array instead of map
* remove complexity around date format
* move examples to the root folder
* simplify development by using latest for the registry and controller image
* rename modules to source
* update image version and catalog registry
* add java specific build
* module configuration
* Improve contributing guide
* just bunch of renames
* change description field to better name
* Description of the operator in OLM
* major bug fixes
* wip
* small scheduler change
* removing unused crds
* use constant for secret
* duplicate value
* minor readme updates
* remove golang codebase
* change quarkus port
* adding mocking and readme
* adding offline token to serviceaccountrequestspec
* move to operator in order to registrer resources
* add api mock
* correcting usage of update control
* fixing build failures
* CRD draft
* resources
* improved types
* CRD draft
* generator setup
* generated API
* api changes
* minor fixes
* add base for operator
* add documentation for resource cleanup
* remove invalid link
* delete not needed file
* Adding basic documentation
* add permissions for CR
* add wider permissions for finalizers
* add push index image
* finalizers update permissions
* add secrets rbac
* group for deployments in rbac control
* update service version
* add info about indexing
* removal procedure
* add additional catalog source elements
* Update image and makefile
* package manifests
* remove replaces
* change bundle
* update catalog source
* change version
* swap images
* bundle details
* disable login
* use secret first
* add metadata
* add more information to
* setup image name
* add arch to builds
* add extreamly complex container
* update docs
* Dependency game
* Updated version of the controller
* add dummy deployment for testing
* Cleanup in operator (use rhoas as name)
* add status changes
* update readme
* reflect values for CLI POC
* remove namespace
* improve documentation
* metadata update
* add kustomization
* bundle
* minor improvements
* add annotations
* add managed kafka instalation instruction
* add controllers folder to finalize docker build
* adding binding example
* add useful commands
* use temporary image
* add binding metadata
* formatting for the types
* minor fixes for the operator
* add extra comments for clarity
* rename title
* add build support
* Update CRD
* Update CRD
* Add basic readme
* initial template for the operator

### Documentation

* add group ID config to quickstarts ([#242](https://github.com/redhat-developer/app-services-operator/issues/242))
* Updates to Operator installation procedure ([#221](https://github.com/redhat-developer/app-services-operator/issues/221))
* Added quick starts initially created for dev sandbox to the operator, so they will be automatically installed on any OpenShift instance on which the rhoas operator is installed. ([#215](https://github.com/redhat-developer/app-services-operator/issues/215))

### Features

* changing org names from com.openshift.cloud to cloud.redhat.com ([#255](https://github.com/redhat-developer/app-services-operator/issues/255))
* added new RHOSAK icons from brand to Quick Starts ([#247](https://github.com/redhat-developer/app-services-operator/issues/247))
* new icon
* install quickstarts on boot
* using new api package
* migration to app-services-sdk
* initial test class and config
* refactoring controllers to put retry and condition handling in outer structure
* adding last generation to conditions
* memory limits, updated icon, minor fixes and tweaks, autoupdate ([#126](https://github.com/redhat-developer/app-services-operator/issues/126))
* Quarkus 1.12.0 and SDK 1.7.4
* adding a new uiRef status to MKC
* maybe conditions
* conditions
* automatic OLM updates
* olm "works"
* refacotring to use dekorate
* managedsarequest
* wip managedserviceaccountcontroller
* offline token exchange
* managedkafkarequest controller works
* refreshing mkrequest
* MKRequest Controller
* wip managedkafkaresponse
* base path and token are configurable
* wip
* k8sclient, upgrade operatorsdk, upgrad quarks
* openapi generated by Java

