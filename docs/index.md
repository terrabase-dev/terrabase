# Terrabase API Reference

## terrabase.application.v1

### ApplicationService (application.v1)

#### CreateApplication

- Auth: not required
- Request: [CreateApplicationRequest](#createapplicationrequest-applicationv1)
- Response: [CreateApplicationResponse](#createapplicationresponse-applicationv1)

#### GetApplication

- Auth: not required
- Request: [GetApplicationRequest](#getapplicationrequest-applicationv1)
- Response: [GetApplicationResponse](#getapplicationresponse-applicationv1)

#### ListApplications

- Auth: not required
- Request: [ListApplicationsRequest](#listapplicationsrequest-applicationv1)
- Response: [ListApplicationsResponse](#listapplicationsresponse-applicationv1)

#### UpdateApplication

- Auth: not required
- Request: [UpdateApplicationRequest](#updateapplicationrequest-applicationv1)
- Response: [UpdateApplicationResponse](#updateapplicationresponse-applicationv1)

#### DeleteApplication

- Auth: not required
- Request: [DeleteApplicationRequest](#deleteapplicationrequest-applicationv1)
- Response: [DeleteApplicationResponse](#deleteapplicationresponse-applicationv1)

#### GrantTeamAccess

- Auth: not required
- Request: [GrantTeamAccessRequest](#grantteamaccessrequest-applicationv1)
- Response: [GrantTeamAccessResponse](#grantteamaccessresponse-applicationv1)

#### RevokeTeamAccess

- Auth: not required
- Request: [RevokeTeamAccessRequest](#revoketeamaccessrequest-applicationv1)
- Response: [RevokeTeamAccessResponse](#revoketeamaccessresponse-applicationv1)

#### Application (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | no | The ID of the application |
| name | string |  | no |  |
| created_at | `Timestamp` |  | no |  |
| updated_at | `Timestamp` |  | no |  |

#### CreateApplicationRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| name | string |  | yes |  |
| team_id | string |  | yes |  |

#### CreateApplicationResponse (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| application | [Application](#application-applicationv1) |  | no |  |

#### GetApplicationRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |

#### GetApplicationResponse (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| application | [Application](#application-applicationv1) |  | no |  |

#### ListApplicationsRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| team_id | string |  | yes |  |
| page_size | int32 |  | no |  |
| page_token | string |  | no |  |

#### ListApplicationsResponse (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| applications | [Application](#application-applicationv1) | repeated | no |  |
| next_page_token | string |  | no |  |

#### UpdateApplicationRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |
| name | string |  | yes |  |
| team_id | string |  | no |  |

#### UpdateApplicationResponse (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| application | [Application](#application-applicationv1) |  | no |  |

#### DeleteApplicationRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |

#### DeleteApplicationResponse (application.v1)

- (no fields)

#### GrantTeamAccessRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| application_id | string |  | yes |  |
| team_id | string | repeated | yes |  |

#### GrantTeamAccessResponse (application.v1)

- (no fields)

#### RevokeTeamAccessRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| application_id | string |  | yes |  |
| team_id | string | repeated | yes |  |

#### RevokeTeamAccessResponse (application.v1)

- (no fields)

## terrabase.auth.v1

### AuthService (auth.v1)

#### Signup

- Auth: not required
- Request: [SignupRequest](#signuprequest-authv1)
- Response: [SignupResponse](#signupresponse-authv1)

#### Login

- Auth: not required
- Request: [LoginRequest](#loginrequest-authv1)
- Response: [LoginResponse](#loginresponse-authv1)

#### Refresh

- Auth: required
- Request: [RefreshRequest](#refreshrequest-authv1)
- Response: [RefreshResponse](#refreshresponse-authv1)

#### WhoAmI

- Auth: required
- Request: [WhoAmIRequest](#whoamirequest-authv1)
- Response: [WhoAmIResponse](#whoamiresponse-authv1)

#### Logout

- Auth: required
- Request: [LogoutRequest](#logoutrequest-authv1)
- Response: [LogoutResponse](#logoutresponse-authv1)

#### ListSessions

- Auth: required
- Request: [ListSessionsRequest](#listsessionsrequest-authv1)
- Response: [ListSessionsResponse](#listsessionsresponse-authv1)

#### CreateMachineUser

- Auth: required
- Scopes: SCOPE_ADMIN
- Request: [CreateMachineUserRequest](#createmachineuserrequest-authv1)
- Response: [CreateMachineUserResponse](#createmachineuserresponse-authv1)

#### CreateApiKey

- Auth: required
- Request: [CreateApiKeyRequest](#createapikeyrequest-authv1)
- Response: [CreateApiKeyResponse](#createapikeyresponse-authv1)

#### ListApiKeys

- Auth: required
- Request: [ListApiKeysRequest](#listapikeysrequest-authv1)
- Response: [ListApiKeysResponse](#listapikeysresponse-authv1)

#### RevokeApiKey

- Auth: required
- Request: [RevokeApiKeyRequest](#revokeapikeyrequest-authv1)
- Response: [RevokeApiKeyResponse](#revokeapikeyresponse-authv1)

#### RotateApiKey

- Auth: required
- Request: [RotateApiKeyRequest](#rotateapikeyrequest-authv1)
- Response: [RotateApiKeyResponse](#rotateapikeyresponse-authv1)

#### SignupRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| name | string |  | yes |  |
| email | string |  | yes |  |
| password | string |  | yes |  |
| default_role | UserRole |  | yes |  |

#### SignupResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| user | [User](#user-userv1) |  | no |  |
| access_token | string |  | no |  |
| refresh_token | string |  | no |  |

#### LoginRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| email | string |  | yes |  |
| password | string |  | yes |  |

#### LoginResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| user | [User](#user-userv1) |  | no |  |
| access_token | string |  | no |  |
| refresh_token | string |  | no |  |

#### RefreshRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| refresh_token | string |  | yes |  |

#### RefreshResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| access_token | string |  | no |  |
| refresh_token | string |  | no |  |

#### WhoAmIRequest (auth.v1)

- (no fields)

#### WhoAmIResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| user | [User](#user-userv1) |  | no |  |
| scopes | Scope | repeated | no |  |

#### LogoutRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| session_id | string |  | no | if empty, use current access token jti |

#### LogoutResponse (auth.v1)

- (no fields)

#### CreateMachineUserRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| name | string |  | yes |  |
| default_role | UserRole |  | yes |  |
| user_type | UserType |  | yes |  |
| owner_user_id | string |  | yes |  |

#### CreateMachineUserResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| machine_user | [User](#user-userv1) |  | no |  |

#### ApiKey (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | no |  |
| name | string |  | no |  |
| scopes | Scope | repeated | no |  |
| owner_id | string |  | no |  |
| owner_type | ApiKeyOwnerType |  | no |  |
| created_at | `Timestamp` |  | no |  |
| expires_at | `Timestamp` |  | no |  |
| last_used_at | `Timestamp` |  | no |  |
| revoked_at | `Timestamp` |  | no |  |

#### CreateApiKeyRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| name | string |  | yes |  |
| owner_type | ApiKeyOwnerType |  | yes |  |
| owner_id | string |  | yes |  |
| scopes | Scope | repeated | yes |  |
| ttl_hours | int64 |  | no | Hours until expiry; if unset, key does not expire. |

#### CreateApiKeyResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| api_key_token | string |  | no |  |
| api_key | [ApiKey](#apikey-authv1) |  | no |  |

#### ListApiKeysRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| owner_type | ApiKeyOwnerType |  | no |  |
| owner_id | string |  | no |  |

#### ListApiKeysResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| api_keys | [ApiKey](#apikey-authv1) | repeated | no |  |

#### RevokeApiKeyRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |
| reason | string |  | no |  |

#### RevokeApiKeyResponse (auth.v1)

- (no fields)

#### RotateApiKeyRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |
| scopes | Scope | repeated | no | Optional: inherit existing scopes and ttl when unset. |
| ttl_hours | int64 |  | no |  |

#### RotateApiKeyResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| api_key_token | string |  | no |  |
| api_key | [ApiKey](#apikey-authv1) |  | no |  |

#### Session (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | no |  |
| user_agent | string |  | no |  |
| ip | string |  | no |  |
| expires_at | `Timestamp` |  | no |  |
| last_used_at | `Timestamp` |  | no |  |
| created_at | `Timestamp` |  | no |  |

#### ListSessionsRequest (auth.v1)

- (no fields)

#### ListSessionsResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| sessions | [Session](#session-authv1) | repeated | no |  |

## terrabase.drift_report.v1

### DriftReportService (drift_report.v1)

#### CreateDriftReport

- Auth: not required
- Request: [CreateDriftReportRequest](#createdriftreportrequest-drift_reportv1)
- Response: [CreateDriftReportResponse](#createdriftreportresponse-drift_reportv1)

#### GetDriftReport

- Auth: not required
- Request: [GetDriftReportRequest](#getdriftreportrequest-drift_reportv1)
- Response: [GetDriftReportResponse](#getdriftreportresponse-drift_reportv1)

#### Drift (drift_report.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| resource_id | string |  | no |  |
| changes | [ChangesEntry](#changesentry-drift_reportv1) | repeated | no |  |

#### ChangesEntry (drift_report.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| key | string |  | no |  |
| value | string |  | no |  |

#### DriftReport (drift_report.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | no |  |
| workspace_id | string |  | no |  |
| drifted | bool |  | no |  |
| changes | [Drift](#drift-drift_reportv1) | repeated | no |  |
| detected_at | `Timestamp` |  | no |  |

#### CreateDriftReportRequest (drift_report.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| workspace_id | string |  | yes |  |
| drifted | bool |  | yes |  |
| changes | [Drift](#drift-drift_reportv1) | repeated | yes |  |

#### CreateDriftReportResponse (drift_report.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| drift_report | [DriftReport](#driftreport-drift_reportv1) |  | no |  |

#### GetDriftReportRequest (drift_report.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |

#### GetDriftReportResponse (drift_report.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| drift_report | [DriftReport](#driftreport-drift_reportv1) |  | no |  |

## terrabase.environment.v1

### EnvironmentService (environment.v1)

#### CreateEnvironment

- Auth: not required
- Request: [CreateEnvironmentRequest](#createenvironmentrequest-environmentv1)
- Response: [CreateEnvironmentResponse](#createenvironmentresponse-environmentv1)

#### GetEnvironment

- Auth: not required
- Request: [GetEnvironmentRequest](#getenvironmentrequest-environmentv1)
- Response: [GetEnvironmentResponse](#getenvironmentresponse-environmentv1)

#### ListEnvironments

- Auth: not required
- Request: [ListEnvironmentsRequest](#listenvironmentsrequest-environmentv1)
- Response: [ListEnvironmentsResponse](#listenvironmentsresponse-environmentv1)

#### UpdateEnvironment

- Auth: not required
- Request: [UpdateEnvironmentRequest](#updateenvironmentrequest-environmentv1)
- Response: [UpdateEnvironmentResponse](#updateenvironmentresponse-environmentv1)

#### DeleteEnvironment

- Auth: not required
- Request: [DeleteEnvironmentRequest](#deleteenvironmentrequest-environmentv1)
- Response: [DeleteEnvironmentResponse](#deleteenvironmentresponse-environmentv1)

#### Environment (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | no |  |
| name | string |  | no |  |
| application_id | string |  | no |  |
| created_at | `Timestamp` |  | no |  |
| updated_at | `Timestamp` |  | no |  |

#### CreateEnvironmentRequest (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| name | string |  | yes |  |
| application_id | string |  | yes |  |
| new_workspace | [CreateWorkspaceRequest](#createworkspacerequest-workspacev1) |  | no |  |

#### CreateEnvironmentResponse (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| environment | [Environment](#environment-environmentv1) |  | no |  |

#### GetEnvironmentRequest (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |

#### GetEnvironmentResponse (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| environment | [Environment](#environment-environmentv1) |  | no |  |

#### ListEnvironmentsRequest (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| application_id | string |  | yes |  |
| page_size | int32 |  | no |  |
| page_token | string |  | no |  |

#### ListEnvironmentsResponse (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| environments | [Environment](#environment-environmentv1) | repeated | no |  |
| next_page_token | string |  | no |  |

#### UpdateEnvironmentRequest (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |
| name | string |  | no |  |

#### UpdateEnvironmentResponse (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| environment | [Environment](#environment-environmentv1) |  | no |  |

#### DeleteEnvironmentRequest (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |

#### DeleteEnvironmentResponse (environment.v1)

- (no fields)

## terrabase.lock.v1

### LockService (lock.v1)

#### CreateLock

- Auth: not required
- Request: [CreateLockRequest](#createlockrequest-lockv1)
- Response: [CreateLockResponse](#createlockresponse-lockv1)

#### GetLock

- Auth: not required
- Request: [GetLockRequest](#getlockrequest-lockv1)
- Response: [GetLockResponse](#getlockresponse-lockv1)

#### DeleteLock

- Auth: not required
- Request: [DeleteLockRequest](#deletelockrequest-lockv1)
- Response: [DeleteLockResponse](#deletelockresponse-lockv1)

#### Lock (lock.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | no |  |
| workspace_id | string |  | no |  |
| owner | [User](#user-userv1) |  | no |  |
| info | string |  | no |  |
| created_at | `Timestamp` |  | no |  |
| expires_at | `Timestamp` |  | no |  |

#### CreateLockRequest (lock.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| workspace_id | string |  | yes |  |
| owner | [User](#user-userv1) |  | yes |  |
| info | string |  | no |  |

#### CreateLockResponse (lock.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| lock | [Lock](#lock-lockv1) |  | no |  |

#### GetLockRequest (lock.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |

#### GetLockResponse (lock.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| lock | [Lock](#lock-lockv1) |  | no |  |

#### DeleteLockRequest (lock.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |

#### DeleteLockResponse (lock.v1)

- (no fields)

## terrabase.organization.v1

### OrganizationService (organization.v1)

#### CreateOrganization

- Auth: required
- Scopes: SCOPE_ADMIN, SCOPE_ORG_WRITE
- Errors: Unauthenticated, PermissionDenied
- Request: [CreateOrganizationRequest](#createorganizationrequest-organizationv1)
- Response: [CreateOrganizationResponse](#createorganizationresponse-organizationv1)

#### GetOrganization

- Auth: required
- Scopes: SCOPE_ADMIN, SCOPE_ORG_READ, SCOPE_ORG_WRITE
- Errors: Unauthenticated, PermissionDenied
- Request: [GetOrganizationRequest](#getorganizationrequest-organizationv1)
- Response: [GetOrganizationResponse](#getorganizationresponse-organizationv1)

#### ListOrganizations

- Auth: required
- Scopes: SCOPE_ADMIN, SCOPE_ORG_READ, SCOPE_ORG_WRITE
- Errors: Unauthenticated, PermissionDenied
- Request: [ListOrganizationsRequest](#listorganizationsrequest-organizationv1)
- Response: [ListOrganizationsResponse](#listorganizationsresponse-organizationv1)

#### UpdateOrganization

- Auth: required
- Scopes: SCOPE_ADMIN, SCOPE_ORG_WRITE
- Errors: Unauthenticated, PermissionDenied
- Request: [UpdateOrganizationRequest](#updateorganizationrequest-organizationv1)
- Response: [UpdateOrganizationResponse](#updateorganizationresponse-organizationv1)

#### DeleteOrganization

- Auth: required
- Scopes: SCOPE_ADMIN, SCOPE_ORG_WRITE
- Errors: Unauthenticated, PermissionDenied
- Request: [DeleteOrganizationRequest](#deleteorganizationrequest-organizationv1)
- Response: [DeleteOrganizationResponse](#deleteorganizationresponse-organizationv1)

#### Organization (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | no |  |
| name | string |  | no |  |
| subscription | Subscription |  | no |  |
| created_at | `Timestamp` |  | no |  |
| updated_at | `Timestamp` |  | no |  |

#### CreateOrganizationRequest (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| name | string |  | yes |  |
| subscription | Subscription |  | yes |  |

#### CreateOrganizationResponse (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| organization | [Organization](#organization-organizationv1) |  | no |  |

#### GetOrganizationRequest (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |

#### GetOrganizationResponse (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| organization | [Organization](#organization-organizationv1) |  | no |  |

#### ListOrganizationsRequest (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| page_size | int32 |  | no |  |
| page_token | string |  | no |  |

#### ListOrganizationsResponse (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| organizations | [Organization](#organization-organizationv1) | repeated | no |  |
| next_page_token | string |  | no |  |

#### UpdateOrganizationRequest (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |
| name | string |  | no |  |
| subscription | Subscription |  | no |  |

#### UpdateOrganizationResponse (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| organization | [Organization](#organization-organizationv1) |  | no |  |

#### DeleteOrganizationRequest (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |

#### DeleteOrganizationResponse (organization.v1)

- (no fields)

## terrabase.s3_backend_config.v1

### S3BackendConfigService (s3_backend_config.v1)

#### CreateS3BackendConfig

- Auth: not required
- Request: [CreateS3BackendConfigRequest](#creates3backendconfigrequest-s3_backend_configv1)
- Response: [CreateS3BackendConfigResponse](#creates3backendconfigresponse-s3_backend_configv1)

#### GetS3BackendConfig

- Auth: not required
- Request: [GetS3BackendConfigRequest](#gets3backendconfigrequest-s3_backend_configv1)
- Response: [GetS3BackendConfigResponse](#gets3backendconfigresponse-s3_backend_configv1)

#### UpdateS3BackendConfig

- Auth: not required
- Request: [UpdateS3BackendConfigRequest](#updates3backendconfigrequest-s3_backend_configv1)
- Response: [UpdateS3BackendConfigResponse](#updates3backendconfigresponse-s3_backend_configv1)

#### DeleteS3BackendConfig

- Auth: not required
- Request: [DeleteS3BackendConfigRequest](#deletes3backendconfigrequest-s3_backend_configv1)
- Response: [DeleteS3BackendConfigResponse](#deletes3backendconfigresponse-s3_backend_configv1)

#### S3BackendConfig (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | no |  |
| bucket | string |  | no |  |
| key | string |  | no |  |
| region | string |  | no |  |
| dynamodb_lock | string |  | no |  |
| encrypt | bool |  | no |  |
| created_at | `Timestamp` |  | no |  |
| updated_at | `Timestamp` |  | no |  |

#### CreateS3BackendConfigRequest (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| bucket | string |  | yes |  |
| key | string |  | yes |  |
| region | string |  | yes |  |
| dynamodb_lock | string |  | no |  |
| encrypt | bool |  | yes |  |

#### CreateS3BackendConfigResponse (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| s3_backend_config | [S3BackendConfig](#s3backendconfig-s3_backend_configv1) |  | no |  |

#### GetS3BackendConfigRequest (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | no |  |

#### GetS3BackendConfigResponse (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| s3_backend_config | [S3BackendConfig](#s3backendconfig-s3_backend_configv1) |  | no |  |

#### UpdateS3BackendConfigRequest (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |
| bucket | string |  | no |  |
| key | string |  | no |  |
| region | string |  | no |  |
| dynamodb_lock | string |  | no |  |
| encrypt | bool |  | no |  |

#### UpdateS3BackendConfigResponse (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| s3_backend_config | [S3BackendConfig](#s3backendconfig-s3_backend_configv1) |  | no |  |

#### DeleteS3BackendConfigRequest (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | no |  |

#### DeleteS3BackendConfigResponse (s3_backend_config.v1)

- (no fields)

## terrabase.state_version.v1

### StateVersionService (state_version.v1)

#### CreateStateVersion

- Auth: not required
- Request: [CreateStateVersionRequest](#createstateversionrequest-state_versionv1)
- Response: [CreateStateVersionResponse](#createstateversionresponse-state_versionv1)

#### GetStateVersion

- Auth: not required
- Request: [GetStateVersionRequest](#getstateversionrequest-state_versionv1)
- Response: [GetStateVersionResponse](#getstateversionresponse-state_versionv1)

#### ListStateVersions

- Auth: not required
- Request: [ListStateVersionsRequest](#liststateversionsrequest-state_versionv1)
- Response: [ListStateVersionsResponse](#liststateversionsresponse-state_versionv1)

#### DeleteStateVersion

- Auth: not required
- Request: [DeleteStateVersionRequest](#deletestateversionrequest-state_versionv1)
- Response: [DeleteStateVersionResponse](#deletestateversionresponse-state_versionv1)

#### StateVersion (state_version.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | no |  |
| workspace_id | string |  | no |  |
| version | string |  | no |  |
| hash | string |  | no |  |
| size_bytes | int64 |  | no |  |
| storage_key | string |  | no |  |
| created_at | `Timestamp` |  | no |  |

#### CreateStateVersionRequest (state_version.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| workspace_id | string |  | yes |  |
| storage_key | string |  | yes |  |

#### CreateStateVersionResponse (state_version.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| state_version | [StateVersion](#stateversion-state_versionv1) |  | no |  |

#### GetStateVersionRequest (state_version.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |

#### GetStateVersionResponse (state_version.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| state_version | [StateVersion](#stateversion-state_versionv1) |  | no |  |

#### ListStateVersionsRequest (state_version.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| workspace_id | string |  | yes |  |
| page_size | int32 |  | no |  |
| page_token | string |  | no |  |

#### ListStateVersionsResponse (state_version.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| state_versions | [StateVersion](#stateversion-state_versionv1) | repeated | no |  |
| next_page_token | string |  | no |  |

#### DeleteStateVersionRequest (state_version.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |

#### DeleteStateVersionResponse (state_version.v1)

- (no fields)

## terrabase.team.v1

### TeamService (team.v1)

#### CreateTeam

- Auth: required
- Scopes: SCOPE_ADMIN, SCOPE_TEAM_WRITE
- Request: [CreateTeamRequest](#createteamrequest-teamv1)
- Response: [CreateTeamResponse](#createteamresponse-teamv1)

#### GetTeam

- Auth: required
- Scopes: SCOPE_ADMIN, SCOPE_TEAM_READ, SCOPE_TEAM_WRITE
- Request: [GetTeamRequest](#getteamrequest-teamv1)
- Response: [GetTeamResponse](#getteamresponse-teamv1)

#### ListTeams

- Auth: required
- Scopes: SCOPE_ADMIN, SCOPE_TEAM_READ, SCOPE_TEAM_WRITE
- Request: [ListTeamsRequest](#listteamsrequest-teamv1)
- Response: [ListTeamsResponse](#listteamsresponse-teamv1)

#### UpdateTeam

- Auth: required
- Scopes: SCOPE_ADMIN, SCOPE_TEAM_WRITE
- Request: [UpdateTeamRequest](#updateteamrequest-teamv1)
- Response: [UpdateTeamResponse](#updateteamresponse-teamv1)

#### DeleteTeam

- Auth: required
- Scopes: SCOPE_ADMIN, SCOPE_TEAM_WRITE
- Request: [DeleteTeamRequest](#deleteteamrequest-teamv1)
- Response: [DeleteTeamResponse](#deleteteamresponse-teamv1)

#### Team (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | no |  |
| name | string |  | no |  |
| organization_id | string |  | no |  |
| created_at | `Timestamp` |  | no |  |
| updated_at | `Timestamp` |  | no |  |

#### TeamIds (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| team_id | string | repeated | no |  |

#### CreateTeamRequest (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| name | string |  | yes |  |
| organization_id | string |  | yes |  |

#### CreateTeamResponse (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| team | [Team](#team-teamv1) |  | no |  |

#### GetTeamRequest (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |

#### GetTeamResponse (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| team | [Team](#team-teamv1) |  | no |  |

#### ListTeamsRequest (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| page_size | int32 |  | no |  |
| page_token | string |  | no |  |

#### ListTeamsResponse (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| teams | [Team](#team-teamv1) | repeated | no |  |
| next_page_token | string |  | no |  |

#### UpdateTeamRequest (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |
| name | string |  | no |  |

#### UpdateTeamResponse (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| team | [Team](#team-teamv1) |  | no |  |

#### DeleteTeamRequest (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |

#### DeleteTeamResponse (team.v1)

- (no fields)

## terrabase.user.v1

### UserService (user.v1)

#### GetUser

- Auth: required
- Request: [GetUserRequest](#getuserrequest-userv1)
- Response: [GetUserResponse](#getuserresponse-userv1)

#### ListUsers

- Auth: required
- Scopes: SCOPE_ADMIN
- Request: [ListUsersRequest](#listusersrequest-userv1)
- Response: [ListUsersResponse](#listusersresponse-userv1)

#### UpdateUser

- Auth: required
- Request: [UpdateUserRequest](#updateuserrequest-userv1)
- Response: [UpdateUserResponse](#updateuserresponse-userv1)

#### DeleteUser

- Auth: required
- Scopes: SCOPE_ADMIN
- Request: [DeleteUserRequest](#deleteuserrequest-userv1)
- Response: [DeleteUserResponse](#deleteuserresponse-userv1)

#### User (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | no |  |
| name | string |  | no |  |
| email | string |  | no |  |
| default_role | UserRole |  | no |  |
| created_at | `Timestamp` |  | no |  |
| updated_at | `Timestamp` |  | no |  |
| user_type | UserType |  | no |  |
| owner_user_id | string |  | no |  |

#### UserSummary (user.v1)

UserSummary is context aware- i.e. the role will be the user's effective role in the context (organization, team, workspace) of which the ListUsers rpc is called

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | no |  |
| name | string |  | no |  |
| email | string |  | no |  |
| role | UserRole |  | no |  |
| created_at | `Timestamp` |  | no |  |
| updated_at | `Timestamp` |  | no |  |
| user_type | UserType |  | no |  |
| owner_user_id | string |  | no |  |

#### GetUserRequest (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |

#### GetUserResponse (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| user | [User](#user-userv1) |  | no |  |

#### ListUsersRequest (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| page_size | int32 |  | no |  |
| page_token | string |  | no |  |
| organization_id | string |  | no |  |
| team_id | string |  | no |  |
| workspace_id | string |  | no |  |
| user_type | UserType |  | no |  |

#### ListUsersResponse (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| users | [UserSummary](#usersummary-userv1) | repeated | no |  |
| next_page_token | string |  | no |  |

#### UpdateUserRequest (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |
| name | string |  | no |  |
| email | string |  | no |  |
| default_role | UserRole |  | no |  |

#### UpdateUserResponse (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| user | [User](#user-userv1) |  | no |  |

#### DeleteUserRequest (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |

#### DeleteUserResponse (user.v1)

- (no fields)

## terrabase.user_membership.v1

### UserMembershipService (user_membership.v1)

#### AddUserToOrganization

- Auth: not required
- Request: [AddUserToOrganizationRequest](#addusertoorganizationrequest-user_membershipv1)
- Response: [AddUserToOrganizationResponse](#addusertoorganizationresponse-user_membershipv1)

#### RemoveUserFromOrganization

- Auth: not required
- Request: [RemoveUserFromOrganizationRequest](#removeuserfromorganizationrequest-user_membershipv1)
- Response: [RemoveUserFromOrganizationResponse](#removeuserfromorganizationresponse-user_membershipv1)

#### AddUserToTeam

- Auth: not required
- Request: [AddUserToTeamRequest](#addusertoteamrequest-user_membershipv1)
- Response: [AddUserToTeamResponse](#addusertoteamresponse-user_membershipv1)

#### RemoveUserFromTeam

- Auth: not required
- Request: [RemoveUserFromTeamRequest](#removeuserfromteamrequest-user_membershipv1)
- Response: [RemoveUserFromTeamResponse](#removeuserfromteamresponse-user_membershipv1)

#### AddUserToWorkspace

- Auth: not required
- Request: [AddUserToWorkspaceRequest](#addusertoworkspacerequest-user_membershipv1)
- Response: [AddUserToWorkspaceResponse](#addusertoworkspaceresponse-user_membershipv1)

#### RemoveUserFromWorkspace

- Auth: not required
- Request: [RemoveUserFromWorkspaceRequest](#removeuserfromworkspacerequest-user_membershipv1)
- Response: [RemoveUserFromWorkspaceResponse](#removeuserfromworkspaceresponse-user_membershipv1)

#### SetUserWorkspaceRole

- Auth: not required
- Request: [SetUserWorkspaceRoleRequest](#setuserworkspacerolerequest-user_membershipv1)
- Response: [SetUserWorkspaceRoleResponse](#setuserworkspaceroleresponse-user_membershipv1)

#### AddUserToOrganizationRequest (user_membership.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| user_id | string |  | yes |  |
| organization_id | string |  | yes |  |

#### AddUserToOrganizationResponse (user_membership.v1)

- (no fields)

#### RemoveUserFromOrganizationRequest (user_membership.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| user_id | string |  | yes |  |
| organization_id | string |  | yes |  |

#### RemoveUserFromOrganizationResponse (user_membership.v1)

- (no fields)

#### AddUserToTeamRequest (user_membership.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| user_id | string |  | yes |  |
| team_id | string |  | yes |  |

#### AddUserToTeamResponse (user_membership.v1)

- (no fields)

#### RemoveUserFromTeamRequest (user_membership.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| user_id | string |  | yes |  |
| team_id | string |  | yes |  |

#### RemoveUserFromTeamResponse (user_membership.v1)

- (no fields)

#### AddUserToWorkspaceRequest (user_membership.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| user_id | string |  | yes |  |
| workspace_id | string |  | yes |  |
| role | UserRole |  | yes |  |

#### AddUserToWorkspaceResponse (user_membership.v1)

- (no fields)

#### RemoveUserFromWorkspaceRequest (user_membership.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| user_id | string |  | yes |  |
| workspace_id | string |  | yes |  |

#### RemoveUserFromWorkspaceResponse (user_membership.v1)

- (no fields)

#### SetUserWorkspaceRoleRequest (user_membership.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| user_id | string |  | yes |  |
| workspace_id | string |  | yes |  |
| role | UserRole |  | yes |  |

#### SetUserWorkspaceRoleResponse (user_membership.v1)

- (no fields)

## terrabase.workspace.v1

### WorkspaceService (workspace.v1)

#### CreateWorkspace

- Auth: not required
- Request: [CreateWorkspaceRequest](#createworkspacerequest-workspacev1)
- Response: [CreateWorkspaceResponse](#createworkspaceresponse-workspacev1)

#### GetWorkspace

- Auth: not required
- Request: [GetWorkspaceRequest](#getworkspacerequest-workspacev1)
- Response: [GetWorkspaceResponse](#getworkspaceresponse-workspacev1)

#### ListWorkspaces

- Auth: not required
- Request: [ListWorkspacesRequest](#listworkspacesrequest-workspacev1)
- Response: [ListWorkspacesResponse](#listworkspacesresponse-workspacev1)

#### UpdateWorkspace

- Auth: not required
- Request: [UpdateWorkspaceRequest](#updateworkspacerequest-workspacev1)
- Response: [UpdateWorkspaceResponse](#updateworkspaceresponse-workspacev1)

#### DeleteWorkspace

- Auth: not required
- Request: [DeleteWorkspaceRequest](#deleteworkspacerequest-workspacev1)
- Response: [DeleteWorkspaceResponse](#deleteworkspaceresponse-workspacev1)

#### GrantTeamAccess

- Auth: not required
- Request: [GrantTeamAccessRequest](#grantteamaccessrequest-workspacev1)
- Response: [GrantTeamAccessResponse](#grantteamaccessresponse-workspacev1)

#### RevokeTeamAccess

- Auth: not required
- Request: [RevokeTeamAccessRequest](#revoketeamaccessrequest-workspacev1)
- Response: [RevokeTeamAccessResponse](#revoketeamaccessresponse-workspacev1)

#### Workspace (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | no |  |
| name | string |  | no |  |
| backend_type | BackendType |  | no |  |
| environment_id | string |  | no |  |
| s3_backend_config_id | string |  | no |  |
| created_at | `Timestamp` |  | no |  |
| updated_at | `Timestamp` |  | no |  |

#### CreateWorkspaceRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| name | string |  | yes |  |
| backend_type | BackendType |  | yes |  |
| environment_id | string |  | no |  |
| team_id | string |  | no |  |
| s3_backend_config | [CreateS3BackendConfigRequest](#creates3backendconfigrequest-s3_backend_configv1) |  | no |  |

#### CreateWorkspaceResponse (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| workspace | [Workspace](#workspace-workspacev1) |  | no |  |

#### GetWorkspaceRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |

#### GetWorkspaceResponse (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| workspace | [Workspace](#workspace-workspacev1) |  | no |  |

#### ListWorkspacesRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| page_size | int32 |  | no |  |
| page_token | string |  | no |  |
| team_id | string |  | no |  |
| environment_id | string |  | no |  |

#### ListWorkspacesResponse (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| workspaces | [Workspace](#workspace-workspacev1) | repeated | no |  |
| next_page_token | string |  | no |  |

#### UpdateWorkspaceRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |
| name | string |  | no |  |
| backend_type | BackendType |  | no |  |
| environment_id | string |  | no |  |
| team_id | string |  | no |  |
| s3_backend_config_id | string |  | no |  |
| s3_backend_config | [CreateS3BackendConfigRequest](#creates3backendconfigrequest-s3_backend_configv1) |  | no |  |

#### UpdateWorkspaceResponse (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| workspace | [Workspace](#workspace-workspacev1) |  | no |  |

#### DeleteWorkspaceRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| id | string |  | yes |  |

#### DeleteWorkspaceResponse (workspace.v1)

- (no fields)

#### GrantTeamAccessRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| workspace_id | string |  | yes |  |
| team_ids | [TeamIds](#teamids-teamv1) |  | yes |  |

#### GrantTeamAccessResponse (workspace.v1)

- (no fields)

#### RevokeTeamAccessRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| workspace_id | string |  | yes |  |
| team_ids | [TeamIds](#teamids-teamv1) |  | yes |  |

#### RevokeTeamAccessResponse (workspace.v1)

- (no fields)
