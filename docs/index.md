# Terrabase API Reference

## terrabase.application.v1

### ApplicationService (application.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `CreateApplication` | [CreateApplicationRequest](#createapplicationrequest-applicationv1) | [CreateApplicationResponse](#createapplicationresponse-applicationv1) | `false` |  | Create a new application |
| `GetApplication` | [GetApplicationRequest](#getapplicationrequest-applicationv1) | [GetApplicationResponse](#getapplicationresponse-applicationv1) | `false` |  | Retrieve details about a single application |
| `ListApplications` | [ListApplicationsRequest](#listapplicationsrequest-applicationv1) | [ListApplicationsResponse](#listapplicationsresponse-applicationv1) | `false` |  | List applications owned by a specific team |
| `UpdateApplication` | [UpdateApplicationRequest](#updateapplicationrequest-applicationv1) | [UpdateApplicationResponse](#updateapplicationresponse-applicationv1) | `false` |  | Change details about an application |
| `DeleteApplication` | [DeleteApplicationRequest](#deleteapplicationrequest-applicationv1) | [DeleteApplicationResponse](#deleteapplicationresponse-applicationv1) | `false` |  | Delete an application |
| `GrantTeamAccess` | [GrantTeamAccessRequest](#grantteamaccessrequest-applicationv1) | [GrantTeamAccessResponse](#grantteamaccessresponse-applicationv1) | `false` |  | Grant access to an application to a single team or multiple teams |
| `RevokeTeamAccess` | [RevokeTeamAccessRequest](#revoketeamaccessrequest-applicationv1) | [RevokeTeamAccessResponse](#revoketeamaccessresponse-applicationv1) | `false` |  | Revoke access to an application from a single team or multiple teams |

### Application (application.v1)

A Terrabase application can be deployed in multiple environments, each with their own workspace

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` | The unique ID of the application |
| `name` | string |  | `false` | The name of the application |
| `team_id` | string |  | `false` | The ID of the team that owns the application |
| `created_at` | `Timestamp` |  | `false` | The time the application was created |
| `updated_at` | `Timestamp` |  | `false` | The time the application was last updated at |

### CreateApplicationRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `name` | string |  | `true` | The name of the application |
| `team_id` | string |  | `true` | The ID of the team that owns the application |

### CreateApplicationResponse (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `application` | [Application](#application-applicationv1) |  | `false` | The application that was created |

### GetApplicationRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the application |

### GetApplicationResponse (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `application` | [Application](#application-applicationv1) |  | `false` | The application |

### ListApplicationsRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `team_id` | string |  | `true` | The ID of a team that owns applications |
| `page_size` | int32 |  | `false` | The number of applications on each page of results |
| `page_token` | string |  | `false` | The token to retrieve the next page of results |

### ListApplicationsResponse (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `applications` | [Application](#application-applicationv1) | repeated | `false` | A list of applications |
| `next_page_token` | string |  | `false` | The token to retrieve the next page of results |

### UpdateApplicationRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the application to update |
| `name` | string |  | `true` | The new name of the application |
| `team_id` | string |  | `false` | The new ID of the team that owns the application |

### UpdateApplicationResponse (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `application` | [Application](#application-applicationv1) |  | `false` | The updated application |

### DeleteApplicationRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the application to delete |

### DeleteApplicationResponse (application.v1)

- (no fields)

### GrantTeamAccessRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `application_id` | string |  | `true` | The unique ID of the application |
| `team_id` | string | repeated | `true` | A list of team IDs who should be granted access to the application |

### GrantTeamAccessResponse (application.v1)

- (no fields)

### RevokeTeamAccessRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `application_id` | string |  | `true` | The unique ID of the application |
| `team_id` | string | repeated | `true` | A list of team IDs whose access to the application should be revoked |

### RevokeTeamAccessResponse (application.v1)

- (no fields)

## terrabase.auth.v1

### AuthService (auth.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `Signup` | [SignupRequest](#signuprequest-authv1) | [SignupResponse](#signupresponse-authv1) | `false` |  | Signup for a new user account in a Terrabase instance |
| `Login` | [LoginRequest](#loginrequest-authv1) | [LoginResponse](#loginresponse-authv1) | `false` |  | Login to a Terrabase instance with a user account |
| `Refresh` | [RefreshRequest](#refreshrequest-authv1) | [RefreshResponse](#refreshresponse-authv1) | `true` |  | Refresh a user's access token |
| `WhoAmI` | [WhoAmIRequest](#whoamirequest-authv1) | [WhoAmIResponse](#whoamiresponse-authv1) | `true` |  | Retrieve details about the logged in user |
| `Logout` | [LogoutRequest](#logoutrequest-authv1) | [LogoutResponse](#logoutresponse-authv1) | `true` |  | Logout of a Terrabase instance |
| `ListSessions` | [ListSessionsRequest](#listsessionsrequest-authv1) | [ListSessionsResponse](#listsessionsresponse-authv1) | `true` |  | List currently active sessions |
| `CreateMachineUser` | [CreateMachineUserRequest](#createmachineuserrequest-authv1) | [CreateMachineUserResponse](#createmachineuserresponse-authv1) | `true` | `SCOPE_ADMIN` | Create a machine user (bot or service principal) |
| `CreateApiKey` | [CreateApiKeyRequest](#createapikeyrequest-authv1) | [CreateApiKeyResponse](#createapikeyresponse-authv1) | `true` | Admin or self | Create an API key. API keys can be owned by human users or machine users |
| `ListApiKeys` | [ListApiKeysRequest](#listapikeysrequest-authv1) | [ListApiKeysResponse](#listapikeysresponse-authv1) | `true` | Admin or self | List API keys owned by a user |
| `RevokeApiKey` | [RevokeApiKeyRequest](#revokeapikeyrequest-authv1) | [RevokeApiKeyResponse](#revokeapikeyresponse-authv1) | `true` | Admin or self | Revoke an API key |
| `RotateApiKey` | [RotateApiKeyRequest](#rotateapikeyrequest-authv1) | [RotateApiKeyResponse](#rotateapikeyresponse-authv1) | `true` | Admin or self | Rotate an API key |

### SignupRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `name` | string |  | `true` | The user's full name |
| `email` | string |  | `true` | The user's email address |
| `password` | string |  | `true` | The user's password |
| `default_role` | UserRole |  | `true` | The default role for the user |

### SignupResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `user` | [User](#user-userv1) |  | `false` | The user that was created |
| `access_token` | string |  | `false` | The user's access token |
| `refresh_token` | string |  | `false` | The user's token to refresh their access token when it expires |

### LoginRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `email` | string |  | `true` | The user's email address |
| `password` | string |  | `true` | The user's password |

### LoginResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `user` | [User](#user-userv1) |  | `false` | The logged in user |
| `access_token` | string |  | `false` | The logged in user's access token |
| `refresh_token` | string |  | `false` | The logged in user's token to refresh their access token when it expires |

### RefreshRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `refresh_token` | string |  | `true` | The user's token to refresh their access token |

### RefreshResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `access_token` | string |  | `false` | A new access token for the user |
| `refresh_token` | string |  | `false` | The user's token to refresh their new access token |

### WhoAmIRequest (auth.v1)

- (no fields)

### WhoAmIResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `user` | [User](#user-userv1) |  | `false` | The logged in user |
| `scopes` | Scope | repeated | `false` | The scopes the user has |

### LogoutRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `session_id` | string |  | `false` | The session to logout of |

### LogoutResponse (auth.v1)

- (no fields)

### CreateMachineUserRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `name` | string |  | `true` | The name of the new machine user |
| `default_role` | UserRole |  | `true` | The default role of the new machine user |
| `user_type` | UserType |  | `true` | The type of the new machine user |
| `owner_user_id` | string |  | `true` | The ID of the user that owns the new machine user |

### CreateMachineUserResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `machine_user` | [User](#user-userv1) |  | `false` | The machine user that was created |

### ApiKey (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` | The unique ID of the API key |
| `name` | string |  | `false` | The name of the API key |
| `scopes` | Scope | repeated | `false` | The scopes the API key has |
| `owner_id` | string |  | `false` | The ID of the user that owns the API key |
| `owner_type` | ApiKeyOwnerType |  | `false` | The type of user that owns the API key |
| `created_at` | `Timestamp` |  | `false` | The time the API key was created |
| `expires_at` | `Timestamp` |  | `false` | The time the API key expires |
| `last_used_at` | `Timestamp` |  | `false` | The time the API key was last used |
| `revoked_at` | `Timestamp` |  | `false` | The time the API key was revoked |

### CreateApiKeyRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `name` | string |  | `true` | The name of the API key |
| `owner_type` | ApiKeyOwnerType |  | `true` | The type of the user that owns the API key |
| `owner_id` | string |  | `true` | The ID of the user that owns the API key |
| `scopes` | Scope | repeated | `true` | The scopes the API key has |
| `ttl_hours` | int64 |  | `false` | Hours until the API key expires. If unset, key does not expire. |

### CreateApiKeyResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `api_key_token` | string |  | `false` | The access token for the API key |
| `api_key` | [ApiKey](#apikey-authv1) |  | `false` | The API key that was created |

### ListApiKeysRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `owner_type` | ApiKeyOwnerType |  | `false` | The type of user that owns API keys |
| `owner_id` | string |  | `false` | The unique ID of the user that owns an API key |

### ListApiKeysResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `api_keys` | [ApiKey](#apikey-authv1) | repeated | `false` | A list of API keys |

### RevokeApiKeyRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The ID of the API key to revoke |
| `reason` | string |  | `false` | The reason for revoking the API key |

### RevokeApiKeyResponse (auth.v1)

- (no fields)

### RotateApiKeyRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The ID of the API key to rotate |
| `scopes` | Scope | repeated | `false` | The scopes the new API key should have. If unset, inherits from the existing API key. |
| `ttl_hours` | int64 |  | `false` | Hours until the new API key expires. If unset, inherits from the existing API key. |

### RotateApiKeyResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `api_key_token` | string |  | `false` | The new access token for the API key |
| `api_key` | [ApiKey](#apikey-authv1) |  | `false` | The new API key |

### Session (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` | The unique ID of the session |
| `user_agent` | string |  | `false` | The user agent the session originated from |
| `ip` | string |  | `false` | The IP address the session originated from |
| `expires_at` | `Timestamp` |  | `false` | The time the session expires at |
| `last_used_at` | `Timestamp` |  | `false` | The time the session was last used at |
| `created_at` | `Timestamp` |  | `false` | The time the session was created |

### ListSessionsRequest (auth.v1)

- (no fields)

### ListSessionsResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `sessions` | [Session](#session-authv1) | repeated | `false` | A list of active sessions |

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
| `CreateOrganization` | [CreateOrganizationRequest](#createorganizationrequest-organizationv1) | [CreateOrganizationResponse](#createorganizationresponse-organizationv1) | `true` | `SCOPE_ADMIN`, `SCOPE_ORG_WRITE` | Create a new organization |
| `GetOrganization` | [GetOrganizationRequest](#getorganizationrequest-organizationv1) | [GetOrganizationResponse](#getorganizationresponse-organizationv1) | `true` | `SCOPE_ADMIN`, `SCOPE_ORG_READ`, `SCOPE_ORG_WRITE` | Retrieve details about a single organization |
| `ListOrganizations` | [ListOrganizationsRequest](#listorganizationsrequest-organizationv1) | [ListOrganizationsResponse](#listorganizationsresponse-organizationv1) | `true` | `SCOPE_ADMIN`, `SCOPE_ORG_READ`, `SCOPE_ORG_WRITE` | List all organizations |
| `UpdateOrganization` | [UpdateOrganizationRequest](#updateorganizationrequest-organizationv1) | [UpdateOrganizationResponse](#updateorganizationresponse-organizationv1) | `true` | `SCOPE_ADMIN`, `SCOPE_ORG_WRITE` | Change details about an organization |
| `DeleteOrganization` | [DeleteOrganizationRequest](#deleteorganizationrequest-organizationv1) | [DeleteOrganizationResponse](#deleteorganizationresponse-organizationv1) | `true` | `SCOPE_ADMIN`, `SCOPE_ORG_WRITE` | Delete an organization |

### Organization (organization.v1)

A Terrabase organization is the top level grouping of resources in a Terrabase instance

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` | The unique ID of the organization |
| `name` | string |  | `false` | The name of the organization |
| `subscription` | Subscription |  | `false` | The subscription level of the organization |
| `created_at` | `Timestamp` |  | `false` | The time the organization was created |
| `updated_at` | `Timestamp` |  | `false` | The time the organization was last updated at |

### CreateOrganizationRequest (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `name` | string |  | `true` | The name of the organization to create |
| `subscription` | Subscription |  | `true` | The subscription level of the organization to create |

### CreateOrganizationResponse (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `organization` | [Organization](#organization-organizationv1) |  | `false` | The organization that was created |

### GetOrganizationRequest (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the organization |

### GetOrganizationResponse (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `organization` | [Organization](#organization-organizationv1) |  | `false` | The organization |

### ListOrganizationsRequest (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `page_size` | int32 |  | `false` | The number of organizations on each page of results |
| `page_token` | string |  | `false` | The token to retrieve the next page of results |

### ListOrganizationsResponse (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `organizations` | [Organization](#organization-organizationv1) | repeated | `false` | A list of organizations |
| `next_page_token` | string |  | `false` | The token to retrieve the next page of results |

### UpdateOrganizationRequest (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the organization to update |
| `name` | string |  | `false` | The new name of the organization |
| `subscription` | Subscription |  | `false` | The new subscription level of the organization |

### UpdateOrganizationResponse (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `organization` | [Organization](#organization-organizationv1) |  | `false` | The updated organization |

### DeleteOrganizationRequest (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the organization to delete |

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
| `CreateTeam` | [CreateTeamRequest](#createteamrequest-teamv1) | [CreateTeamResponse](#createteamresponse-teamv1) | `true` | `SCOPE_ADMIN`, `SCOPE_TEAM_WRITE` | Create a new team |
| `GetTeam` | [GetTeamRequest](#getteamrequest-teamv1) | [GetTeamResponse](#getteamresponse-teamv1) | `true` | `SCOPE_ADMIN`, `SCOPE_TEAM_READ`, `SCOPE_TEAM_WRITE` | Retrieve details about a single team |
| `ListTeams` | [ListTeamsRequest](#listteamsrequest-teamv1) | [ListTeamsResponse](#listteamsresponse-teamv1) | `true` | `SCOPE_ADMIN`, `SCOPE_TEAM_READ`, `SCOPE_TEAM_WRITE` | List all teams |
| `UpdateTeam` | [UpdateTeamRequest](#updateteamrequest-teamv1) | [UpdateTeamResponse](#updateteamresponse-teamv1) | `true` | `SCOPE_ADMIN`, `SCOPE_TEAM_WRITE` | Change details about a team |
| `DeleteTeam` | [DeleteTeamRequest](#deleteteamrequest-teamv1) | [DeleteTeamResponse](#deleteteamresponse-teamv1) | `true` | `SCOPE_ADMIN`, `SCOPE_TEAM_WRITE` | Delete a team |

### Team (team.v1)

A Terrabase team belongs to a single organization, and can have many users. Teams should likely strongly correlate with actual business units

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` | The unique ID of the team |
| `name` | string |  | `false` | The name of the team |
| `organization_id` | string |  | `false` | The ID of the organization the team belongs to |
| `created_at` | `Timestamp` |  | `false` | The time the team was created |
| `updated_at` | `Timestamp` |  | `false` | The time the team was last updated |

### TeamIds (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `team_id` | string | repeated | `false` |  |

### CreateTeamRequest (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `name` | string |  | `true` | The name of the team |
| `organization_id` | string |  | `true` | The ID of the organization the team belongs to |

### CreateTeamResponse (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `team` | [Team](#team-teamv1) |  | `false` | The team that was created |

### GetTeamRequest (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the team |

### GetTeamResponse (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `team` | [Team](#team-teamv1) |  | `false` | The team |

### ListTeamsRequest (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `page_size` | int32 |  | `false` | The number of teams on each page of results |
| `page_token` | string |  | `false` | The token to retrieve the next page of results |

### ListTeamsResponse (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `teams` | [Team](#team-teamv1) | repeated | `false` | A list of teams |
| `next_page_token` | string |  | `false` | The token to retrieve the next page of results |

### UpdateTeamRequest (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the team to update |
| `name` | string |  | `false` | The new name of the team |

### UpdateTeamResponse (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `team` | [Team](#team-teamv1) |  | `false` | The updated team |

### DeleteTeamRequest (team.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the team to delete |

### DeleteTeamResponse (team.v1)

- (no fields)

## terrabase.user.v1

### UserService (user.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `GetUser` | [GetUserRequest](#getuserrequest-userv1) | [GetUserResponse](#getuserresponse-userv1) | `true` | Admin or self | Retrieve details about a single user |
| `ListUsers` | [ListUsersRequest](#listusersrequest-userv1) | [ListUsersResponse](#listusersresponse-userv1) | `true` | `SCOPE_ADMIN` | List users who belong to an organization or team, or who have access to a specific workspace |
| `UpdateUser` | [UpdateUserRequest](#updateuserrequest-userv1) | [UpdateUserResponse](#updateuserresponse-userv1) | `true` | Admin or self | Change details about a user |
| `DeleteUser` | [DeleteUserRequest](#deleteuserrequest-userv1) | [DeleteUserResponse](#deleteuserresponse-userv1) | `true` | `SCOPE_ADMIN` | Delete a user |

### User (user.v1)

A Terrabase user is either a human user or a machine user. A machine user is either a bot or a service principal

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` | The unique ID of the user |
| `name` | string |  | `false` | The name of the user |
| `email` | string |  | `false` | The email address of the user, if a human user |
| `default_role` | UserRole |  | `false` | The default role of the user |
| `created_at` | `Timestamp` |  | `false` | The time the user was created |
| `updated_at` | `Timestamp` |  | `false` | The time the user was last updated at |
| `user_type` | UserType |  | `false` | The type of the user (user, bot, or service) |
| `owner_user_id` | string |  | `false` | The ID of the user that owns the machine user, if a machine user |

### UserSummary (user.v1)

UserSummary is context aware

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` | The ID of the user |
| `name` | string |  | `false` | The name of the user |
| `email` | string |  | `false` | The email address of the user, if a human user |
| `role` | UserRole |  | `false` | The role the user has in the context in which the RPC returning this object (or list of objects) is called from |
| `created_at` | `Timestamp` |  | `false` | The time the user was created |
| `updated_at` | `Timestamp` |  | `false` | The time the user was last updated at |
| `user_type` | UserType |  | `false` | The type of the user (user, bot, or service) |
| `owner_user_id` | string |  | `false` | The ID of the user that owns the machine user, if a machine user |

### GetUserRequest (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The ID of the user |

### GetUserResponse (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `user` | [User](#user-userv1) |  | `false` | The user |

### ListUsersRequest (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `page_size` | int32 |  | `false` | The number of users on each page of results |
| `page_token` | string |  | `false` | The token to retrieve the next page of results |
| `organization_id` | string |  | `false` | The ID of the organization to list all users who belong to the organization |
| `team_id` | string |  | `false` | The ID of the team to list all users who belong to the team |
| `workspace_id` | string |  | `false` | The ID of the workspace to list all users with access to the workspace |
| `user_type` | UserType |  | `false` | The type of users to list |

### ListUsersResponse (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `users` | [UserSummary](#usersummary-userv1) | repeated | `false` | A list of users |
| `next_page_token` | string |  | `false` | The token to retrieve the next page of results |

### UpdateUserRequest (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the user to update |
| `name` | string |  | `false` | The new name of the user |
| `email` | string |  | `false` | The new email address of the user |
| `default_role` | UserRole |  | `false` | The new default role of the user |

### UpdateUserResponse (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `user` | [User](#user-userv1) |  | `false` | The updated user |

### DeleteUserRequest (user.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the user to delete |

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
