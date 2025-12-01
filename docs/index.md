# Terrabase API Reference

## terrabase.application.v1

### ApplicationService (application.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `CreateApplication` | [CreateApplicationRequest](#createapplicationrequest-applicationv1) | [CreateApplicationResponse](#createapplicationresponse-applicationv1) | `false` |  |  |
| `GetApplication` | [GetApplicationRequest](#getapplicationrequest-applicationv1) | [GetApplicationResponse](#getapplicationresponse-applicationv1) | `false` |  |  |
| `ListApplications` | [ListApplicationsRequest](#listapplicationsrequest-applicationv1) | [ListApplicationsResponse](#listapplicationsresponse-applicationv1) | `false` |  |  |
| `UpdateApplication` | [UpdateApplicationRequest](#updateapplicationrequest-applicationv1) | [UpdateApplicationResponse](#updateapplicationresponse-applicationv1) | `false` |  |  |
| `DeleteApplication` | [DeleteApplicationRequest](#deleteapplicationrequest-applicationv1) | [DeleteApplicationResponse](#deleteapplicationresponse-applicationv1) | `false` |  |  |
| `GrantTeamAccess` | [GrantTeamAccessRequest](#grantteamaccessrequest-applicationv1) | [GrantTeamAccessResponse](#grantteamaccessresponse-applicationv1) | `false` |  |  |
| `RevokeTeamAccess` | [RevokeTeamAccessRequest](#revoketeamaccessrequest-applicationv1) | [RevokeTeamAccessResponse](#revoketeamaccessresponse-applicationv1) | `false` |  |  |

### Application (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` | The ID of the application |
| `name` | string |  | `false` |  |
| `created_at` | `Timestamp` |  | `false` |  |
| `updated_at` | `Timestamp` |  | `false` |  |

### CreateApplicationRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `name` | string |  | `true` |  |
| `team_id` | string |  | `true` |  |

### CreateApplicationResponse (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `application` | [Application](#application-applicationv1) |  | `false` |  |

### GetApplicationRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |

### GetApplicationResponse (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `application` | [Application](#application-applicationv1) |  | `false` |  |

### ListApplicationsRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `team_id` | string |  | `true` |  |
| `page_size` | int32 |  | `false` |  |
| `page_token` | string |  | `false` |  |

### ListApplicationsResponse (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `applications` | [Application](#application-applicationv1) | repeated | `false` |  |
| `next_page_token` | string |  | `false` |  |

### UpdateApplicationRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |
| `name` | string |  | `true` |  |
| `team_id` | string |  | `false` |  |

### UpdateApplicationResponse (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `application` | [Application](#application-applicationv1) |  | `false` |  |

### DeleteApplicationRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |

### DeleteApplicationResponse (application.v1)

- (no fields)

### GrantTeamAccessRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `application_id` | string |  | `true` |  |
| `team_id` | string | repeated | `true` |  |

### GrantTeamAccessResponse (application.v1)

- (no fields)

### RevokeTeamAccessRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `application_id` | string |  | `true` |  |
| `team_id` | string | repeated | `true` |  |

### RevokeTeamAccessResponse (application.v1)

- (no fields)

## terrabase.auth.v1

### AuthService (auth.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `Signup` | [SignupRequest](#signuprequest-authv1) | [SignupResponse](#signupresponse-authv1) | `false` |  |  |
| `Login` | [LoginRequest](#loginrequest-authv1) | [LoginResponse](#loginresponse-authv1) | `false` |  |  |
| `Refresh` | [RefreshRequest](#refreshrequest-authv1) | [RefreshResponse](#refreshresponse-authv1) | `true` |  |  |
| `WhoAmI` | [WhoAmIRequest](#whoamirequest-authv1) | [WhoAmIResponse](#whoamiresponse-authv1) | `true` |  |  |
| `Logout` | [LogoutRequest](#logoutrequest-authv1) | [LogoutResponse](#logoutresponse-authv1) | `true` |  |  |
| `ListSessions` | [ListSessionsRequest](#listsessionsrequest-authv1) | [ListSessionsResponse](#listsessionsresponse-authv1) | `true` |  |  |
| `CreateMachineUser` | [CreateMachineUserRequest](#createmachineuserrequest-authv1) | [CreateMachineUserResponse](#createmachineuserresponse-authv1) | `true` | `SCOPE_ADMIN` |  |
| `CreateApiKey` | [CreateApiKeyRequest](#createapikeyrequest-authv1) | [CreateApiKeyResponse](#createapikeyresponse-authv1) | `true` | Admin or self |  |
| `ListApiKeys` | [ListApiKeysRequest](#listapikeysrequest-authv1) | [ListApiKeysResponse](#listapikeysresponse-authv1) | `true` | Admin or self |  |
| `RevokeApiKey` | [RevokeApiKeyRequest](#revokeapikeyrequest-authv1) | [RevokeApiKeyResponse](#revokeapikeyresponse-authv1) | `true` | Admin or self |  |
| `RotateApiKey` | [RotateApiKeyRequest](#rotateapikeyrequest-authv1) | [RotateApiKeyResponse](#rotateapikeyresponse-authv1) | `true` | Admin or self |  |

### SignupRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `name` | string |  | `true` |  |
| `email` | string |  | `true` |  |
| `password` | string |  | `true` |  |
| `default_role` | UserRole |  | `true` |  |

### SignupResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `user` | [User](#user-userv1) |  | `false` |  |
| `access_token` | string |  | `false` |  |
| `refresh_token` | string |  | `false` |  |

### LoginRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `email` | string |  | `true` |  |
| `password` | string |  | `true` |  |

### LoginResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `user` | [User](#user-userv1) |  | `false` |  |
| `access_token` | string |  | `false` |  |
| `refresh_token` | string |  | `false` |  |

### RefreshRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `refresh_token` | string |  | `true` |  |

### RefreshResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `access_token` | string |  | `false` |  |
| `refresh_token` | string |  | `false` |  |

### WhoAmIRequest (auth.v1)

- (no fields)

### WhoAmIResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `user` | [User](#user-userv1) |  | `false` |  |
| `scopes` | Scope | repeated | `false` |  |

### LogoutRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `session_id` | string |  | `false` | if empty, use current access token jti |

### LogoutResponse (auth.v1)

- (no fields)

### CreateMachineUserRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `name` | string |  | `true` |  |
| `default_role` | UserRole |  | `true` |  |
| `user_type` | UserType |  | `true` |  |
| `owner_user_id` | string |  | `true` |  |

### CreateMachineUserResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `machine_user` | [User](#user-userv1) |  | `false` |  |

### ApiKey (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` |  |
| `name` | string |  | `false` |  |
| `scopes` | Scope | repeated | `false` |  |
| `owner_id` | string |  | `false` |  |
| `owner_type` | ApiKeyOwnerType |  | `false` |  |
| `created_at` | `Timestamp` |  | `false` |  |
| `expires_at` | `Timestamp` |  | `false` |  |
| `last_used_at` | `Timestamp` |  | `false` |  |
| `revoked_at` | `Timestamp` |  | `false` |  |

### CreateApiKeyRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `name` | string |  | `true` |  |
| `owner_type` | ApiKeyOwnerType |  | `true` |  |
| `owner_id` | string |  | `true` |  |
| `scopes` | Scope | repeated | `true` |  |
| `ttl_hours` | int64 |  | `false` | Hours until expiry; if unset, key does not expire. |

### CreateApiKeyResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `api_key_token` | string |  | `false` |  |
| `api_key` | [ApiKey](#apikey-authv1) |  | `false` |  |

### ListApiKeysRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `owner_type` | ApiKeyOwnerType |  | `false` |  |
| `owner_id` | string |  | `false` |  |

### ListApiKeysResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `api_keys` | [ApiKey](#apikey-authv1) | repeated | `false` |  |

### RevokeApiKeyRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |
| `reason` | string |  | `false` |  |

### RevokeApiKeyResponse (auth.v1)

- (no fields)

### RotateApiKeyRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |
| `scopes` | Scope | repeated | `false` | Optional: inherit existing scopes and ttl when unset. |
| `ttl_hours` | int64 |  | `false` |  |

### RotateApiKeyResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `api_key_token` | string |  | `false` |  |
| `api_key` | [ApiKey](#apikey-authv1) |  | `false` |  |

### Session (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` |  |
| `user_agent` | string |  | `false` |  |
| `ip` | string |  | `false` |  |
| `expires_at` | `Timestamp` |  | `false` |  |
| `last_used_at` | `Timestamp` |  | `false` |  |
| `created_at` | `Timestamp` |  | `false` |  |

### ListSessionsRequest (auth.v1)

- (no fields)

### ListSessionsResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `sessions` | [Session](#session-authv1) | repeated | `false` |  |

## terrabase.drift_report.v1

### DriftReportService (drift_report.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `CreateDriftReport` | [CreateDriftReportRequest](#createdriftreportrequest-drift_reportv1) | [CreateDriftReportResponse](#createdriftreportresponse-drift_reportv1) | `false` |  |  |
| `GetDriftReport` | [GetDriftReportRequest](#getdriftreportrequest-drift_reportv1) | [GetDriftReportResponse](#getdriftreportresponse-drift_reportv1) | `false` |  |  |

### Drift (drift_report.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `resource_id` | string |  | `false` |  |
| `changes` | [ChangesEntry](#changesentry-drift_reportv1) | repeated | `false` |  |

### ChangesEntry (drift_report.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `key` | string |  | `false` |  |
| `value` | string |  | `false` |  |

### DriftReport (drift_report.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` |  |
| `workspace_id` | string |  | `false` |  |
| `drifted` | bool |  | `false` |  |
| `changes` | [Drift](#drift-drift_reportv1) | repeated | `false` |  |
| `detected_at` | `Timestamp` |  | `false` |  |

### CreateDriftReportRequest (drift_report.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `workspace_id` | string |  | `true` |  |
| `drifted` | bool |  | `true` |  |
| `changes` | [Drift](#drift-drift_reportv1) | repeated | `true` |  |

### CreateDriftReportResponse (drift_report.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `drift_report` | [DriftReport](#driftreport-drift_reportv1) |  | `false` |  |

### GetDriftReportRequest (drift_report.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |

### GetDriftReportResponse (drift_report.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `drift_report` | [DriftReport](#driftreport-drift_reportv1) |  | `false` |  |

## terrabase.environment.v1

### EnvironmentService (environment.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `CreateEnvironment` | [CreateEnvironmentRequest](#createenvironmentrequest-environmentv1) | [CreateEnvironmentResponse](#createenvironmentresponse-environmentv1) | `false` |  |  |
| `GetEnvironment` | [GetEnvironmentRequest](#getenvironmentrequest-environmentv1) | [GetEnvironmentResponse](#getenvironmentresponse-environmentv1) | `false` |  |  |
| `ListEnvironments` | [ListEnvironmentsRequest](#listenvironmentsrequest-environmentv1) | [ListEnvironmentsResponse](#listenvironmentsresponse-environmentv1) | `false` |  |  |
| `UpdateEnvironment` | [UpdateEnvironmentRequest](#updateenvironmentrequest-environmentv1) | [UpdateEnvironmentResponse](#updateenvironmentresponse-environmentv1) | `false` |  |  |
| `DeleteEnvironment` | [DeleteEnvironmentRequest](#deleteenvironmentrequest-environmentv1) | [DeleteEnvironmentResponse](#deleteenvironmentresponse-environmentv1) | `false` |  |  |

### Environment (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` |  |
| `name` | string |  | `false` |  |
| `application_id` | string |  | `false` |  |
| `created_at` | `Timestamp` |  | `false` |  |
| `updated_at` | `Timestamp` |  | `false` |  |

### CreateEnvironmentRequest (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `name` | string |  | `true` |  |
| `application_id` | string |  | `true` |  |
| `new_workspace` | [CreateWorkspaceRequest](#createworkspacerequest-workspacev1) |  | `false` |  |

### CreateEnvironmentResponse (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `environment` | [Environment](#environment-environmentv1) |  | `false` |  |

### GetEnvironmentRequest (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |

### GetEnvironmentResponse (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `environment` | [Environment](#environment-environmentv1) |  | `false` |  |

### ListEnvironmentsRequest (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `application_id` | string |  | `true` |  |
| `page_size` | int32 |  | `false` |  |
| `page_token` | string |  | `false` |  |

### ListEnvironmentsResponse (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `environments` | [Environment](#environment-environmentv1) | repeated | `false` |  |
| `next_page_token` | string |  | `false` |  |

### UpdateEnvironmentRequest (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |
| `name` | string |  | `false` |  |

### UpdateEnvironmentResponse (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `environment` | [Environment](#environment-environmentv1) |  | `false` |  |

### DeleteEnvironmentRequest (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |

### DeleteEnvironmentResponse (environment.v1)

- (no fields)

## terrabase.lock.v1

### LockService (lock.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `CreateLock` | [CreateLockRequest](#createlockrequest-lockv1) | [CreateLockResponse](#createlockresponse-lockv1) | `false` |  |  |
| `GetLock` | [GetLockRequest](#getlockrequest-lockv1) | [GetLockResponse](#getlockresponse-lockv1) | `false` |  |  |
| `DeleteLock` | [DeleteLockRequest](#deletelockrequest-lockv1) | [DeleteLockResponse](#deletelockresponse-lockv1) | `false` |  |  |

### Lock (lock.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` |  |
| `workspace_id` | string |  | `false` |  |
| `owner` | [User](#user-userv1) |  | `false` |  |
| `info` | string |  | `false` |  |
| `created_at` | `Timestamp` |  | `false` |  |
| `expires_at` | `Timestamp` |  | `false` |  |

### CreateLockRequest (lock.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `workspace_id` | string |  | `true` |  |
| `owner` | [User](#user-userv1) |  | `true` |  |
| `info` | string |  | `false` |  |

### CreateLockResponse (lock.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `lock` | [Lock](#lock-lockv1) |  | `false` |  |

### GetLockRequest (lock.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |

### GetLockResponse (lock.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `lock` | [Lock](#lock-lockv1) |  | `false` |  |

### DeleteLockRequest (lock.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |

### DeleteLockResponse (lock.v1)

- (no fields)

## terrabase.organization.v1

### OrganizationService (organization.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `CreateOrganization` | [CreateOrganizationRequest](#createorganizationrequest-organizationv1) | [CreateOrganizationResponse](#createorganizationresponse-organizationv1) | `true` | `SCOPE_ADMIN`, `SCOPE_ORG_WRITE` |  |
| `GetOrganization` | [GetOrganizationRequest](#getorganizationrequest-organizationv1) | [GetOrganizationResponse](#getorganizationresponse-organizationv1) | `true` | `SCOPE_ADMIN`, `SCOPE_ORG_READ`, `SCOPE_ORG_WRITE` |  |
| `ListOrganizations` | [ListOrganizationsRequest](#listorganizationsrequest-organizationv1) | [ListOrganizationsResponse](#listorganizationsresponse-organizationv1) | `true` | `SCOPE_ADMIN`, `SCOPE_ORG_READ`, `SCOPE_ORG_WRITE` |  |
| `UpdateOrganization` | [UpdateOrganizationRequest](#updateorganizationrequest-organizationv1) | [UpdateOrganizationResponse](#updateorganizationresponse-organizationv1) | `true` | `SCOPE_ADMIN`, `SCOPE_ORG_WRITE` |  |
| `DeleteOrganization` | [DeleteOrganizationRequest](#deleteorganizationrequest-organizationv1) | [DeleteOrganizationResponse](#deleteorganizationresponse-organizationv1) | `true` | `SCOPE_ADMIN`, `SCOPE_ORG_WRITE` |  |

### Organization (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` |  |
| `name` | string |  | `false` |  |
| `subscription` | Subscription |  | `false` |  |
| `created_at` | `Timestamp` |  | `false` |  |
| `updated_at` | `Timestamp` |  | `false` |  |

### CreateOrganizationRequest (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `name` | string |  | `true` |  |
| `subscription` | Subscription |  | `true` |  |

### CreateOrganizationResponse (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `organization` | [Organization](#organization-organizationv1) |  | `false` |  |

### GetOrganizationRequest (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |

### GetOrganizationResponse (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `organization` | [Organization](#organization-organizationv1) |  | `false` |  |

### ListOrganizationsRequest (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `page_size` | int32 |  | `false` |  |
| `page_token` | string |  | `false` |  |

### ListOrganizationsResponse (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `organizations` | [Organization](#organization-organizationv1) | repeated | `false` |  |
| `next_page_token` | string |  | `false` |  |

### UpdateOrganizationRequest (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |
| `name` | string |  | `false` |  |
| `subscription` | Subscription |  | `false` |  |

### UpdateOrganizationResponse (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `organization` | [Organization](#organization-organizationv1) |  | `false` |  |

### DeleteOrganizationRequest (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |

### DeleteOrganizationResponse (organization.v1)

- (no fields)

## terrabase.s3_backend_config.v1

### S3BackendConfigService (s3_backend_config.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `CreateS3BackendConfig` | [CreateS3BackendConfigRequest](#creates3backendconfigrequest-s3_backend_configv1) | [CreateS3BackendConfigResponse](#creates3backendconfigresponse-s3_backend_configv1) | `false` |  |  |
| `GetS3BackendConfig` | [GetS3BackendConfigRequest](#gets3backendconfigrequest-s3_backend_configv1) | [GetS3BackendConfigResponse](#gets3backendconfigresponse-s3_backend_configv1) | `false` |  |  |
| `UpdateS3BackendConfig` | [UpdateS3BackendConfigRequest](#updates3backendconfigrequest-s3_backend_configv1) | [UpdateS3BackendConfigResponse](#updates3backendconfigresponse-s3_backend_configv1) | `false` |  |  |
| `DeleteS3BackendConfig` | [DeleteS3BackendConfigRequest](#deletes3backendconfigrequest-s3_backend_configv1) | [DeleteS3BackendConfigResponse](#deletes3backendconfigresponse-s3_backend_configv1) | `false` |  |  |

### S3BackendConfig (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` |  |
| `bucket` | string |  | `false` |  |
| `key` | string |  | `false` |  |
| `region` | string |  | `false` |  |
| `dynamodb_lock` | string |  | `false` |  |
| `encrypt` | bool |  | `false` |  |
| `created_at` | `Timestamp` |  | `false` |  |
| `updated_at` | `Timestamp` |  | `false` |  |

### CreateS3BackendConfigRequest (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `bucket` | string |  | `true` |  |
| `key` | string |  | `true` |  |
| `region` | string |  | `true` |  |
| `dynamodb_lock` | string |  | `false` |  |
| `encrypt` | bool |  | `true` |  |

### CreateS3BackendConfigResponse (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `s3_backend_config` | [S3BackendConfig](#s3backendconfig-s3_backend_configv1) |  | `false` |  |

### GetS3BackendConfigRequest (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` |  |

### GetS3BackendConfigResponse (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `s3_backend_config` | [S3BackendConfig](#s3backendconfig-s3_backend_configv1) |  | `false` |  |

### UpdateS3BackendConfigRequest (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |
| `bucket` | string |  | `false` |  |
| `key` | string |  | `false` |  |
| `region` | string |  | `false` |  |
| `dynamodb_lock` | string |  | `false` |  |
| `encrypt` | bool |  | `false` |  |

### UpdateS3BackendConfigResponse (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `s3_backend_config` | [S3BackendConfig](#s3backendconfig-s3_backend_configv1) |  | `false` |  |

### DeleteS3BackendConfigRequest (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` |  |

### DeleteS3BackendConfigResponse (s3_backend_config.v1)

- (no fields)

## terrabase.state_version.v1

### StateVersionService (state_version.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `CreateStateVersion` | [CreateStateVersionRequest](#createstateversionrequest-state_versionv1) | [CreateStateVersionResponse](#createstateversionresponse-state_versionv1) | `false` |  |  |
| `GetStateVersion` | [GetStateVersionRequest](#getstateversionrequest-state_versionv1) | [GetStateVersionResponse](#getstateversionresponse-state_versionv1) | `false` |  |  |
| `ListStateVersions` | [ListStateVersionsRequest](#liststateversionsrequest-state_versionv1) | [ListStateVersionsResponse](#liststateversionsresponse-state_versionv1) | `false` |  |  |
| `DeleteStateVersion` | [DeleteStateVersionRequest](#deletestateversionrequest-state_versionv1) | [DeleteStateVersionResponse](#deletestateversionresponse-state_versionv1) | `false` |  |  |

### StateVersion (state_version.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` |  |
| `workspace_id` | string |  | `false` |  |
| `version` | string |  | `false` |  |
| `hash` | string |  | `false` |  |
| `size_bytes` | int64 |  | `false` |  |
| `storage_key` | string |  | `false` |  |
| `created_at` | `Timestamp` |  | `false` |  |

### CreateStateVersionRequest (state_version.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `workspace_id` | string |  | `true` |  |
| `storage_key` | string |  | `true` |  |

### CreateStateVersionResponse (state_version.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `state_version` | [StateVersion](#stateversion-state_versionv1) |  | `false` |  |

### GetStateVersionRequest (state_version.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |

### GetStateVersionResponse (state_version.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `state_version` | [StateVersion](#stateversion-state_versionv1) |  | `false` |  |

### ListStateVersionsRequest (state_version.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `workspace_id` | string |  | `true` |  |
| `page_size` | int32 |  | `false` |  |
| `page_token` | string |  | `false` |  |

### ListStateVersionsResponse (state_version.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `state_versions` | [StateVersion](#stateversion-state_versionv1) | repeated | `false` |  |
| `next_page_token` | string |  | `false` |  |

### DeleteStateVersionRequest (state_version.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |

### DeleteStateVersionResponse (state_version.v1)

- (no fields)

## terrabase.team.v1

### TeamService (team.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `CreateTeam` | [CreateTeamRequest](#createteamrequest-teamv1) | [CreateTeamResponse](#createteamresponse-teamv1) | `true` | `SCOPE_ADMIN`, `SCOPE_TEAM_WRITE` |  |
| `GetTeam` | [GetTeamRequest](#getteamrequest-teamv1) | [GetTeamResponse](#getteamresponse-teamv1) | `true` | `SCOPE_ADMIN`, `SCOPE_TEAM_READ`, `SCOPE_TEAM_WRITE` |  |
| `ListTeams` | [ListTeamsRequest](#listteamsrequest-teamv1) | [ListTeamsResponse](#listteamsresponse-teamv1) | `true` | `SCOPE_ADMIN`, `SCOPE_TEAM_READ`, `SCOPE_TEAM_WRITE` |  |
| `UpdateTeam` | [UpdateTeamRequest](#updateteamrequest-teamv1) | [UpdateTeamResponse](#updateteamresponse-teamv1) | `true` | `SCOPE_ADMIN`, `SCOPE_TEAM_WRITE` |  |
| `DeleteTeam` | [DeleteTeamRequest](#deleteteamrequest-teamv1) | [DeleteTeamResponse](#deleteteamresponse-teamv1) | `true` | `SCOPE_ADMIN`, `SCOPE_TEAM_WRITE` |  |

### Team (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` |  |
| `name` | string |  | `false` |  |
| `organization_id` | string |  | `false` |  |
| `created_at` | `Timestamp` |  | `false` |  |
| `updated_at` | `Timestamp` |  | `false` |  |

### TeamIds (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `team_id` | string | repeated | `false` |  |

### CreateTeamRequest (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `name` | string |  | `true` |  |
| `organization_id` | string |  | `true` |  |

### CreateTeamResponse (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `team` | [Team](#team-teamv1) |  | `false` |  |

### GetTeamRequest (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |

### GetTeamResponse (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `team` | [Team](#team-teamv1) |  | `false` |  |

### ListTeamsRequest (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `page_size` | int32 |  | `false` |  |
| `page_token` | string |  | `false` |  |

### ListTeamsResponse (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `teams` | [Team](#team-teamv1) | repeated | `false` |  |
| `next_page_token` | string |  | `false` |  |

### UpdateTeamRequest (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |
| `name` | string |  | `false` |  |

### UpdateTeamResponse (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `team` | [Team](#team-teamv1) |  | `false` |  |

### DeleteTeamRequest (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |

### DeleteTeamResponse (team.v1)

- (no fields)

## terrabase.user.v1

### UserService (user.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `GetUser` | [GetUserRequest](#getuserrequest-userv1) | [GetUserResponse](#getuserresponse-userv1) | `true` | Admin or self |  |
| `ListUsers` | [ListUsersRequest](#listusersrequest-userv1) | [ListUsersResponse](#listusersresponse-userv1) | `true` | `SCOPE_ADMIN` |  |
| `UpdateUser` | [UpdateUserRequest](#updateuserrequest-userv1) | [UpdateUserResponse](#updateuserresponse-userv1) | `true` | Admin or self |  |
| `DeleteUser` | [DeleteUserRequest](#deleteuserrequest-userv1) | [DeleteUserResponse](#deleteuserresponse-userv1) | `true` | `SCOPE_ADMIN` |  |

### User (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` |  |
| `name` | string |  | `false` |  |
| `email` | string |  | `false` |  |
| `default_role` | UserRole |  | `false` |  |
| `created_at` | `Timestamp` |  | `false` |  |
| `updated_at` | `Timestamp` |  | `false` |  |
| `user_type` | UserType |  | `false` |  |
| `owner_user_id` | string |  | `false` |  |

### UserSummary (user.v1)

UserSummary is context aware- i.e. the role will be the user's effective role in the context (organization, team, workspace) of which the ListUsers rpc is called

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` |  |
| `name` | string |  | `false` |  |
| `email` | string |  | `false` |  |
| `role` | UserRole |  | `false` |  |
| `created_at` | `Timestamp` |  | `false` |  |
| `updated_at` | `Timestamp` |  | `false` |  |
| `user_type` | UserType |  | `false` |  |
| `owner_user_id` | string |  | `false` |  |

### GetUserRequest (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |

### GetUserResponse (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `user` | [User](#user-userv1) |  | `false` |  |

### ListUsersRequest (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `page_size` | int32 |  | `false` |  |
| `page_token` | string |  | `false` |  |
| `organization_id` | string |  | `false` |  |
| `team_id` | string |  | `false` |  |
| `workspace_id` | string |  | `false` |  |
| `user_type` | UserType |  | `false` |  |

### ListUsersResponse (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `users` | [UserSummary](#usersummary-userv1) | repeated | `false` |  |
| `next_page_token` | string |  | `false` |  |

### UpdateUserRequest (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |
| `name` | string |  | `false` |  |
| `email` | string |  | `false` |  |
| `default_role` | UserRole |  | `false` |  |

### UpdateUserResponse (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `user` | [User](#user-userv1) |  | `false` |  |

### DeleteUserRequest (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |

### DeleteUserResponse (user.v1)

- (no fields)

## terrabase.user_membership.v1

### UserMembershipService (user_membership.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `AddUserToOrganization` | [AddUserToOrganizationRequest](#addusertoorganizationrequest-user_membershipv1) | [AddUserToOrganizationResponse](#addusertoorganizationresponse-user_membershipv1) | `false` |  |  |
| `RemoveUserFromOrganization` | [RemoveUserFromOrganizationRequest](#removeuserfromorganizationrequest-user_membershipv1) | [RemoveUserFromOrganizationResponse](#removeuserfromorganizationresponse-user_membershipv1) | `false` |  |  |
| `AddUserToTeam` | [AddUserToTeamRequest](#addusertoteamrequest-user_membershipv1) | [AddUserToTeamResponse](#addusertoteamresponse-user_membershipv1) | `false` |  |  |
| `RemoveUserFromTeam` | [RemoveUserFromTeamRequest](#removeuserfromteamrequest-user_membershipv1) | [RemoveUserFromTeamResponse](#removeuserfromteamresponse-user_membershipv1) | `false` |  |  |
| `AddUserToWorkspace` | [AddUserToWorkspaceRequest](#addusertoworkspacerequest-user_membershipv1) | [AddUserToWorkspaceResponse](#addusertoworkspaceresponse-user_membershipv1) | `false` |  |  |
| `RemoveUserFromWorkspace` | [RemoveUserFromWorkspaceRequest](#removeuserfromworkspacerequest-user_membershipv1) | [RemoveUserFromWorkspaceResponse](#removeuserfromworkspaceresponse-user_membershipv1) | `false` |  |  |
| `SetUserWorkspaceRole` | [SetUserWorkspaceRoleRequest](#setuserworkspacerolerequest-user_membershipv1) | [SetUserWorkspaceRoleResponse](#setuserworkspaceroleresponse-user_membershipv1) | `false` |  |  |

### AddUserToOrganizationRequest (user_membership.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `user_id` | string |  | `true` |  |
| `organization_id` | string |  | `true` |  |

### AddUserToOrganizationResponse (user_membership.v1)

- (no fields)

### RemoveUserFromOrganizationRequest (user_membership.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `user_id` | string |  | `true` |  |
| `organization_id` | string |  | `true` |  |

### RemoveUserFromOrganizationResponse (user_membership.v1)

- (no fields)

### AddUserToTeamRequest (user_membership.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `user_id` | string |  | `true` |  |
| `team_id` | string |  | `true` |  |

### AddUserToTeamResponse (user_membership.v1)

- (no fields)

### RemoveUserFromTeamRequest (user_membership.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `user_id` | string |  | `true` |  |
| `team_id` | string |  | `true` |  |

### RemoveUserFromTeamResponse (user_membership.v1)

- (no fields)

### AddUserToWorkspaceRequest (user_membership.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `user_id` | string |  | `true` |  |
| `workspace_id` | string |  | `true` |  |
| `role` | UserRole |  | `true` |  |

### AddUserToWorkspaceResponse (user_membership.v1)

- (no fields)

### RemoveUserFromWorkspaceRequest (user_membership.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `user_id` | string |  | `true` |  |
| `workspace_id` | string |  | `true` |  |

### RemoveUserFromWorkspaceResponse (user_membership.v1)

- (no fields)

### SetUserWorkspaceRoleRequest (user_membership.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `user_id` | string |  | `true` |  |
| `workspace_id` | string |  | `true` |  |
| `role` | UserRole |  | `true` |  |

### SetUserWorkspaceRoleResponse (user_membership.v1)

- (no fields)

## terrabase.workspace.v1

### WorkspaceService (workspace.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `CreateWorkspace` | [CreateWorkspaceRequest](#createworkspacerequest-workspacev1) | [CreateWorkspaceResponse](#createworkspaceresponse-workspacev1) | `false` |  |  |
| `GetWorkspace` | [GetWorkspaceRequest](#getworkspacerequest-workspacev1) | [GetWorkspaceResponse](#getworkspaceresponse-workspacev1) | `false` |  |  |
| `ListWorkspaces` | [ListWorkspacesRequest](#listworkspacesrequest-workspacev1) | [ListWorkspacesResponse](#listworkspacesresponse-workspacev1) | `false` |  |  |
| `UpdateWorkspace` | [UpdateWorkspaceRequest](#updateworkspacerequest-workspacev1) | [UpdateWorkspaceResponse](#updateworkspaceresponse-workspacev1) | `false` |  |  |
| `DeleteWorkspace` | [DeleteWorkspaceRequest](#deleteworkspacerequest-workspacev1) | [DeleteWorkspaceResponse](#deleteworkspaceresponse-workspacev1) | `false` |  |  |
| `GrantTeamAccess` | [GrantTeamAccessRequest](#grantteamaccessrequest-workspacev1) | [GrantTeamAccessResponse](#grantteamaccessresponse-workspacev1) | `false` |  |  |
| `RevokeTeamAccess` | [RevokeTeamAccessRequest](#revoketeamaccessrequest-workspacev1) | [RevokeTeamAccessResponse](#revoketeamaccessresponse-workspacev1) | `false` |  |  |

### Workspace (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` |  |
| `name` | string |  | `false` |  |
| `backend_type` | BackendType |  | `false` |  |
| `environment_id` | string |  | `false` |  |
| `s3_backend_config_id` | string |  | `false` |  |
| `created_at` | `Timestamp` |  | `false` |  |
| `updated_at` | `Timestamp` |  | `false` |  |

### CreateWorkspaceRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `name` | string |  | `true` |  |
| `backend_type` | BackendType |  | `true` |  |
| `environment_id` | string |  | `false` |  |
| `team_id` | string |  | `false` |  |
| `s3_backend_config` | [CreateS3BackendConfigRequest](#creates3backendconfigrequest-s3_backend_configv1) |  | `false` |  |

### CreateWorkspaceResponse (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `workspace` | [Workspace](#workspace-workspacev1) |  | `false` |  |

### GetWorkspaceRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |

### GetWorkspaceResponse (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `workspace` | [Workspace](#workspace-workspacev1) |  | `false` |  |

### ListWorkspacesRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `page_size` | int32 |  | `false` |  |
| `page_token` | string |  | `false` |  |
| `team_id` | string |  | `false` |  |
| `environment_id` | string |  | `false` |  |

### ListWorkspacesResponse (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `workspaces` | [Workspace](#workspace-workspacev1) | repeated | `false` |  |
| `next_page_token` | string |  | `false` |  |

### UpdateWorkspaceRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |
| `name` | string |  | `false` |  |
| `backend_type` | BackendType |  | `false` |  |
| `environment_id` | string |  | `false` |  |
| `team_id` | string |  | `false` |  |
| `s3_backend_config_id` | string |  | `false` |  |
| `s3_backend_config` | [CreateS3BackendConfigRequest](#creates3backendconfigrequest-s3_backend_configv1) |  | `false` |  |

### UpdateWorkspaceResponse (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `workspace` | [Workspace](#workspace-workspacev1) |  | `false` |  |

### DeleteWorkspaceRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` |  |

### DeleteWorkspaceResponse (workspace.v1)

- (no fields)

### GrantTeamAccessRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `workspace_id` | string |  | `true` |  |
| `team_ids` | [TeamIds](#teamids-teamv1) |  | `true` |  |

### GrantTeamAccessResponse (workspace.v1)

- (no fields)

### RevokeTeamAccessRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `workspace_id` | string |  | `true` |  |
| `team_ids` | [TeamIds](#teamids-teamv1) |  | `true` |  |

### RevokeTeamAccessResponse (workspace.v1)

- (no fields)
