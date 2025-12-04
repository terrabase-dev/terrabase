# Terrabase API Reference

## terrabase.application.v1

### ApplicationService (application.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `CreateApplication` | [CreateApplicationRequest](#createapplicationrequest-applicationv1) | [CreateApplicationResponse](#createapplicationresponse-applicationv1) | `true` | `SCOPE_ADMIN`, `SCOPE_APPLICATION_WRITE` | Create a new application |
| `GetApplication` | [GetApplicationRequest](#getapplicationrequest-applicationv1) | [GetApplicationResponse](#getapplicationresponse-applicationv1) | `true` | `SCOPE_ADMIN`, `SCOPE_APPLICATION_READ`, `SCOPE_APPLICATION_WRITE` | Retrieve details about a single application |
| `ListApplications` | [ListApplicationsRequest](#listapplicationsrequest-applicationv1) | [ListApplicationsResponse](#listapplicationsresponse-applicationv1) | `true` | `SCOPE_ADMIN`, `SCOPE_APPLICATION_READ`, `SCOPE_APPLICATION_WRITE` | List applications owned by a specific team |
| `UpdateApplication` | [UpdateApplicationRequest](#updateapplicationrequest-applicationv1) | [UpdateApplicationResponse](#updateapplicationresponse-applicationv1) | `true` | `SCOPE_ADMIN`, `SCOPE_APPLICATION_WRITE` | Change details about an application |
| `DeleteApplication` | [DeleteApplicationRequest](#deleteapplicationrequest-applicationv1) | [DeleteApplicationResponse](#deleteapplicationresponse-applicationv1) | `true` | `SCOPE_ADMIN`, `SCOPE_APPLICATION_WRITE` | Delete an application |
| `GrantTeamAccess` | [GrantTeamAccessRequest](#grantteamaccessrequest-applicationv1) | [GrantTeamAccessResponse](#grantteamaccessresponse-applicationv1) | `true` | `SCOPE_ADMIN`, `SCOPE_APPLICATION_WRITE` | Grant access to an application to a single team or multiple teams |
| `RevokeTeamAccess` | [RevokeTeamAccessRequest](#revoketeamaccessrequest-applicationv1) | [RevokeTeamAccessResponse](#revoketeamaccessresponse-applicationv1) | `true` | `SCOPE_ADMIN`, `SCOPE_APPLICATION_WRITE` | Revoke access to an application from a single team or multiple teams |

### Application (application.v1)

A Terrabase application can be deployed in multiple environments, each with their own workspace

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` | The unique ID of the application |
| `name` | string |  | `false` | The name of the application |
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
| `name` | string |  | `false` | The new name of the application |

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
| `team_access_grants` | [CreateTeamApplicationAccessGrantRequest](#createteamapplicationaccessgrantrequest-team_application_access_grantv1) | repeated | `true` | A list of team access grants |

### GrantTeamAccessResponse (application.v1)

- (no fields)

### RevokeTeamAccessRequest (application.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the application |
| `team_ids` | [TeamIds](#teamids-teamv1) |  | `true` | A list of team IDs whose access to the application should be revoked |

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
| `default_role` | [UserRole](#userrole-user_rolev1) |  | `true` | The default role for the user |

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
| `scopes` | [Scope](#scope-authzv1) | repeated | `false` | The scopes the user has |

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
| `default_role` | [UserRole](#userrole-user_rolev1) |  | `true` | The default role of the new machine user |
| `user_type` | [UserType](#usertype-userv1) |  | `true` | The type of the new machine user |
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
| `scopes` | [Scope](#scope-authzv1) | repeated | `false` | The scopes the API key has |
| `owner_id` | string |  | `false` | The ID of the user that owns the API key |
| `owner_type` | [ApiKeyOwnerType](#apikeyownertype-authv1) |  | `false` | The type of user that owns the API key |
| `created_at` | `Timestamp` |  | `false` | The time the API key was created |
| `expires_at` | `Timestamp` |  | `false` | The time the API key expires |
| `last_used_at` | `Timestamp` |  | `false` | The time the API key was last used |
| `revoked_at` | `Timestamp` |  | `false` | The time the API key was revoked |

### CreateApiKeyRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `name` | string |  | `true` | The name of the API key |
| `owner_type` | [ApiKeyOwnerType](#apikeyownertype-authv1) |  | `true` | The type of the user that owns the API key |
| `owner_id` | string |  | `true` | The ID of the user that owns the API key |
| `scopes` | [Scope](#scope-authzv1) | repeated | `true` | The scopes the API key has |
| `ttl_hours` | int64 |  | `false` | Hours until the API key expires. If unset, key does not expire. |

### CreateApiKeyResponse (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `api_key_token` | string |  | `false` | The access token for the API key |
| `api_key` | [ApiKey](#apikey-authv1) |  | `false` | The API key that was created |

### ListApiKeysRequest (auth.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `owner_type` | [ApiKeyOwnerType](#apikeyownertype-authv1) |  | `false` | The type of user that owns API keys |
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
| `scopes` | [Scope](#scope-authzv1) | repeated | `false` | The scopes the new API key should have. If unset, inherits from the existing API key. |
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

### ApiKeyOwnerType (auth.v1)

| Name | Number | Description |
| --- | --- | --- |
| `API_KEY_OWNER_TYPE_UNSPECIFIED` | `0` | Default - should not use |
| `API_KEY_OWNER_TYPE_USER` | `1` | A human user |
| `API_KEY_OWNER_TYPE_BOT` | `2` | A bot user |
| `API_KEY_OWNER_TYPE_SERVICE` | `3` | A service principal |

## terrabase.authz.v1

### Scope (authz.v1)

| Name | Number | Description |
| --- | --- | --- |
| `SCOPE_UNSPECIFIED` | `0` |  |
| `SCOPE_ADMIN` | `1` |  |
| `SCOPE_ORG_WRITE` | `2` |  |
| `SCOPE_ORG_READ` | `3` |  |
| `SCOPE_TEAM_WRITE` | `4` |  |
| `SCOPE_TEAM_READ` | `5` |  |
| `SCOPE_APPLICATION_WRITE` | `6` |  |
| `SCOPE_APPLICATION_READ` | `7` |  |
| `SCOPE_ENVIRONMENT_WRITE` | `8` |  |
| `SCOPE_ENVIRONMENT_READ` | `9` |  |
| `SCOPE_WORKSPACE_WRITE` | `10` |  |
| `SCOPE_WORKSPACE_READ` | `11` |  |

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
| `CreateEnvironment` | [CreateEnvironmentRequest](#createenvironmentrequest-environmentv1) | [CreateEnvironmentResponse](#createenvironmentresponse-environmentv1) | `false` |  | Create a new environment |
| `GetEnvironment` | [GetEnvironmentRequest](#getenvironmentrequest-environmentv1) | [GetEnvironmentResponse](#getenvironmentresponse-environmentv1) | `false` |  | Retrieve details about a specific environment |
| `ListEnvironments` | [ListEnvironmentsRequest](#listenvironmentsrequest-environmentv1) | [ListEnvironmentsResponse](#listenvironmentsresponse-environmentv1) | `false` |  | List environments that belong to an application |
| `UpdateEnvironment` | [UpdateEnvironmentRequest](#updateenvironmentrequest-environmentv1) | [UpdateEnvironmentResponse](#updateenvironmentresponse-environmentv1) | `false` |  | Change details about an environment |
| `DeleteEnvironment` | [DeleteEnvironmentRequest](#deleteenvironmentrequest-environmentv1) | [DeleteEnvironmentResponse](#deleteenvironmentresponse-environmentv1) | `false` |  | Delete an environment |

### Environment (environment.v1)

A Terrabase environment is a business environment that an application is deployed in

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` | The unique ID of the environment |
| `name` | string |  | `false` | The name of the environment |
| `application_id` | string |  | `false` | The ID of the application this environment belongs to |
| `created_at` | `Timestamp` |  | `false` | The time the environment was created |
| `updated_at` | `Timestamp` |  | `false` | The time the environment was last updated at |

### CreateEnvironmentRequest (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `name` | string |  | `true` | The name of the environment |
| `application_id` | string |  | `true` | The ID of the application the environment belongs to |
| `new_workspace` | [CreateWorkspaceRequest](#createworkspacerequest-workspacev1) |  | `false` | The configuration for the workspace that belongs to this environment |

### CreateEnvironmentResponse (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `environment` | [Environment](#environment-environmentv1) |  | `false` | The environment that was created |

### GetEnvironmentRequest (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the environment |

### GetEnvironmentResponse (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `environment` | [Environment](#environment-environmentv1) |  | `false` | The environment |

### ListEnvironmentsRequest (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `application_id` | string |  | `true` | The ID of the application to list environments for |

### ListEnvironmentsResponse (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `environments` | [Environment](#environment-environmentv1) | repeated | `false` | A list of environments |

### UpdateEnvironmentRequest (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the environment to update |
| `name` | string |  | `false` | The new name of the environment |

### UpdateEnvironmentResponse (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `environment` | [Environment](#environment-environmentv1) |  | `false` | The updated environment |

### DeleteEnvironmentRequest (environment.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the environment to delete |

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
| `subscription` | [Subscription](#subscription-organizationv1) |  | `false` | The subscription level of the organization |
| `created_at` | `Timestamp` |  | `false` | The time the organization was created |
| `updated_at` | `Timestamp` |  | `false` | The time the organization was last updated at |

### CreateOrganizationRequest (organization.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `name` | string |  | `true` | The name of the organization to create |
| `subscription` | [Subscription](#subscription-organizationv1) |  | `true` | The subscription level of the organization to create |

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
| `subscription` | [Subscription](#subscription-organizationv1) |  | `false` | The new subscription level of the organization |

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

### Subscription (organization.v1)

| Name | Number | Description |
| --- | --- | --- |
| `SUBSCRIPTION_UNSPECIFIED` | `0` | Default - should not use |
| `SUBSCRIPTION_FREE` | `1` | Free subscription |
| `SUBSCRIPTION_TEAM` | `2` | Team subscription |
| `SUBSCRIPTION_ENTERPRISE` | `3` | Enterprise subscription |

## terrabase.s3_backend_config.v1

### S3BackendConfigService (s3_backend_config.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `CreateS3BackendConfig` | [CreateS3BackendConfigRequest](#creates3backendconfigrequest-s3_backend_configv1) | [CreateS3BackendConfigResponse](#creates3backendconfigresponse-s3_backend_configv1) | `false` |  | Create a new S3 backend configuration |
| `GetS3BackendConfig` | [GetS3BackendConfigRequest](#gets3backendconfigrequest-s3_backend_configv1) | [GetS3BackendConfigResponse](#gets3backendconfigresponse-s3_backend_configv1) | `false` |  | Retrieve details about a single S3 backend configuration |
| `UpdateS3BackendConfig` | [UpdateS3BackendConfigRequest](#updates3backendconfigrequest-s3_backend_configv1) | [UpdateS3BackendConfigResponse](#updates3backendconfigresponse-s3_backend_configv1) | `false` |  | Change details about an S3 backend configuration |
| `DeleteS3BackendConfig` | [DeleteS3BackendConfigRequest](#deletes3backendconfigrequest-s3_backend_configv1) | [DeleteS3BackendConfigResponse](#deletes3backendconfigresponse-s3_backend_configv1) | `false` |  | Delete an S3 backend configuration |

### S3BackendConfig (s3_backend_config.v1)

Configuration for an S3 backend for a Terraform state file. See [Terraform documentation](https://developer.hashicorp.com/terraform/language/backend/s3) for more details.

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` | The unique ID of the S3 backend configuration |
| `workspace_id` | string |  | `false` | The ID of the workspace the S3 backend configuration belongs to |
| `bucket` | string |  | `false` | The name of the S3 Bucket where the state file is stored |
| `key` | string |  | `false` | The path to the state file inside the S3 Bucket |
| `region` | string |  | `false` | The AWS region of the S3 Bucket and DynamoDB Table (if used) |
| `dynamodb_lock` | bool |  | `false` | Whether or not to use DynamoDB state locking. Defaults to `false`, as DynamoDB state locking is deprecated and will be removed in a future Terraform version. Mutually exclusive with `s3_lock`. |
| `s3_lock` | bool |  | `false` | Whether or not to use S3 state locking. Defaults to `true`. Mutually exclusive with `dynamodb_lock`. |
| `encrypt` | bool |  | `false` | Whether or not to enable [server side encryption](https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingServerSideEncryption.html) of the state and lock files |
| `created_at` | `Timestamp` |  | `false` | The time the S3 backend configuration was created |
| `updated_at` | `Timestamp` |  | `false` | The time the S3 backend configuration was last updated at |
| `dynamodb_table` | string |  | `false` | The name of the DynamoDB Table to use for state file locking. The table must have a partition key named `LockID` with a type of `String`. Required if `dynamodb_lock` is `true`, ignored otherwise. |

### CreateS3BackendConfigRequest (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `bucket` | string |  | `true` | The name of the S3 Bucket where the state file is stored |
| `key` | string |  | `true` | The path to the state file inside the S3 Bucket |
| `region` | string |  | `true` | The AWS region of the S3 Bucket and DynamoDB Table (if used) |
| `dynamodb_lock` | bool |  | `false` | Whether or not to use DynamoDB state locking. Defaults to `false`, as DynamoDB state locking is deprecated and will be removed in a future Terraform version. Mutually exclusive with `s3_lock`. |
| `s3_lock` | bool |  | `false` | Whether or not to use S3 state locking. Defaults to `true`. Mutually exclusive with `dynamodb_lock`. |
| `encrypt` | bool |  | `true` | Whether or not to enable [server side encryption](https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingServerSideEncryption.html) of the state and lock files |
| `dynamodb_table` | string |  | `false` | The name of the DynamoDB Table to use for state file locking. The table must have a partition key named `LockID` with a type of `String`. Required if `dynamodb_lock` is `true`, ignored otherwise. |

### CreateS3BackendConfigResponse (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `s3_backend_config` | [S3BackendConfig](#s3backendconfig-s3_backend_configv1) |  | `false` | The S3 backend configuration that was created |

### GetS3BackendConfigRequest (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` | The unique ID of the S3 backend configuration |

### GetS3BackendConfigResponse (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `s3_backend_config` | [S3BackendConfig](#s3backendconfig-s3_backend_configv1) |  | `false` | The S3 backend configuration |

### UpdateS3BackendConfigRequest (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the S3 backend configuration to update |
| `workspace_id` | string |  | `false` | The ID of the new workspace the S3 backend configuration belongs to |
| `bucket` | string |  | `false` | The new name of the S3 Bucket where the state file is stored |
| `key` | string |  | `false` | The new path to the state file inside the S3 Bucket |
| `region` | string |  | `false` | The new AWS region of the S3 Bucket and DynamoDB Table (if used) |
| `dynamodb_lock` | bool |  | `false` | Whether or not to use DynamoDB state locking. Defaults to `false`, as DynamoDB state locking is deprecated and will be removed in a future Terraform version. Mutually exclusive with `s3_lock`. |
| `s3_lock` | bool |  | `false` | Whether or not to use S3 state locking. Defaults to `true`. Mutually exclusive with `dynamodb_lock`. |
| `encrypt` | bool |  | `false` | Whether or not to enable [server side encryption](https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingServerSideEncryption.html) of the state and lock files |
| `dynamodb_table` | string |  | `false` | The name of the new DynamoDB Table to use for state file locking. The table must have a partition key named `LockID` with a type of `String`. Required if `dynamodb_lock` is `true`, ignored otherwise. |

### UpdateS3BackendConfigResponse (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `s3_backend_config` | [S3BackendConfig](#s3backendconfig-s3_backend_configv1) |  | `false` | The updated S3 backend configuration |

### DeleteS3BackendConfigRequest (s3_backend_config.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` | The unique ID of the S3 backend configuration to delete |

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
| `team_ids` | string | repeated | `false` | A list of team IDs |

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

## terrabase.team_access_type.v1

### TeamAccessType (team_access_type.v1)

| Name | Number | Description |
| --- | --- | --- |
| `TEAM_ACCESS_TYPE_UNSPECIFIED` | `0` |  |
| `TEAM_ACCESS_TYPE_OWNER` | `1` |  |
| `TEAM_ACCESS_TYPE_GRANTED` | `2` |  |

## terrabase.team_application_access_grant.v1

### TeamApplicationAccessGrantService (team_application_access_grant.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `CreateTeamApplicationAccessGrant` | [CreateTeamApplicationAccessGrantRequest](#createteamapplicationaccessgrantrequest-team_application_access_grantv1) | [CreateTeamApplicationAccessGrantResponse](#createteamapplicationaccessgrantresponse-team_application_access_grantv1) | `true` | `SCOPE_ADMIN`, `SCOPE_APPLICATION_WRITE`, `SCOPE_TEAM_WRITE` |  |
| `GetTeamApplicationAccessGrant` | [GetTeamApplicationAccessGrantRequest](#getteamapplicationaccessgrantrequest-team_application_access_grantv1) | [GetTeamApplicationAccessGrantResponse](#getteamapplicationaccessgrantresponse-team_application_access_grantv1) | `true` | `SCOPE_ADMIN`, `SCOPE_APPLICATION_READ`, `SCOPE_APPLICATION_WRITE`, `SCOPE_TEAM_READ`, `SCOPE_TEAM_WRITE` |  |
| `ListTeamApplicationAccessGrants` | [ListTeamApplicationAccessGrantsRequest](#listteamapplicationaccessgrantsrequest-team_application_access_grantv1) | [ListTeamApplicationAccessGrantsResponse](#listteamapplicationaccessgrantsresponse-team_application_access_grantv1) | `true` | `SCOPE_ADMIN`, `SCOPE_APPLICATION_READ`, `SCOPE_APPLICATION_WRITE`, `SCOPE_TEAM_READ`, `SCOPE_TEAM_WRITE` |  |
| `UpdateTeamApplicationAccessGrant` | [UpdateTeamApplicationAccessGrantRequest](#updateteamapplicationaccessgrantrequest-team_application_access_grantv1) | [UpdateTeamApplicationAccessGrantResponse](#updateteamapplicationaccessgrantresponse-team_application_access_grantv1) | `true` | `SCOPE_ADMIN`, `SCOPE_APPLICATION_WRITE`, `SCOPE_TEAM_WRITE` |  |
| `DeleteTeamApplicationAccessGrant` | [DeleteTeamApplicationAccessGrantRequest](#deleteteamapplicationaccessgrantrequest-team_application_access_grantv1) | [DeleteTeamApplicationAccessGrantResponse](#deleteteamapplicationaccessgrantresponse-team_application_access_grantv1) | `true` | `SCOPE_ADMIN`, `SCOPE_APPLICATION_WRITE`, `SCOPE_TEAM_WRITE` |  |
| `DeleteTeamApplicationAccessGrantById` | [DeleteTeamApplicationAccessGrantByIdRequest](#deleteteamapplicationaccessgrantbyidrequest-team_application_access_grantv1) | [DeleteTeamApplicationAccessGrantByIdResponse](#deleteteamapplicationaccessgrantbyidresponse-team_application_access_grantv1) | `true` | `SCOPE_ADMIN`, `SCOPE_APPLICATION_WRITE`, `SCOPE_TEAM_WRITE` |  |

### TeamApplicationAccessGrant (team_application_access_grant.v1)

Grants a team access to an application

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` | The unique ID of the team access grant |
| `team_id` | string |  | `false` | The ID of the team that has access to the application |
| `application_id` | string |  | `false` | The ID of the application the team has access to |
| `access_type` | [TeamAccessType](#teamaccesstype-team_access_typev1) |  | `false` | The type of access the team has to the application |
| `created_at` | `Timestamp` |  | `false` | The time the access to the application was granted to the team |
| `updated_at` | `Timestamp` |  | `false` | The time the team's access to the application was last changed |

### CreateTeamApplicationAccessGrantRequest (team_application_access_grant.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `team_id` | string |  | `true` | The ID of the team to grant access to the application to |
| `application_id` | string |  | `true` | The ID of the application to grant the team access to |
| `access_type` | [TeamAccessType](#teamaccesstype-team_access_typev1) |  | `true` | The type of access that should be granted to the team |

### CreateTeamApplicationAccessGrantResponse (team_application_access_grant.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `team_application_access_grant` | [TeamApplicationAccessGrant](#teamapplicationaccessgrant-team_application_access_grantv1) |  | `false` | The team access grant that was created |

### GetTeamApplicationAccessGrantRequest (team_application_access_grant.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the team access grant |

### GetTeamApplicationAccessGrantResponse (team_application_access_grant.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `team_application_access_grant` | [TeamApplicationAccessGrant](#teamapplicationaccessgrant-team_application_access_grantv1) |  | `false` | The team access grant |

### ListTeamApplicationAccessGrantsRequest (team_application_access_grant.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `team_id` | string |  | `false` | The ID of a team to list all applications that the team has access to |
| `application_id` | string |  | `false` | The ID of an application to list all the teams that have access to it |
| `page_size` | int32 |  | `false` | The number of access grants on each page of results |
| `page_token` | string |  | `false` | The token to retrieve the next page of results |

### ListTeamApplicationAccessGrantsResponse (team_application_access_grant.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `team_application_access_grants` | [TeamApplicationAccessGrant](#teamapplicationaccessgrant-team_application_access_grantv1) | repeated | `false` | A list of team access grants |
| `next_page_token` | string |  | `false` | The token to retrieve the next page of results |

### UpdateTeamApplicationAccessGrantRequest (team_application_access_grant.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the team access grant |
| `access_type` | [TeamAccessType](#teamaccesstype-team_access_typev1) |  | `true` | The new type of access that should be granted to the team |

### UpdateTeamApplicationAccessGrantResponse (team_application_access_grant.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `team_application_access_grant` | [TeamApplicationAccessGrant](#teamapplicationaccessgrant-team_application_access_grantv1) |  | `false` | The updated team access grant |

### DeleteTeamApplicationAccessGrantRequest (team_application_access_grant.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `team_id` | string |  | `true` | The ID of the team to whose access grant should be deleted |
| `application_id` | string |  | `true` | The ID of the application whose access grant should be deleted |

### DeleteTeamApplicationAccessGrantResponse (team_application_access_grant.v1)

- (no fields)

### DeleteTeamApplicationAccessGrantByIdRequest (team_application_access_grant.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the team access grant |

### DeleteTeamApplicationAccessGrantByIdResponse (team_application_access_grant.v1)

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
| `default_role` | [UserRole](#userrole-user_rolev1) |  | `false` | The default role of the user |
| `created_at` | `Timestamp` |  | `false` | The time the user was created |
| `updated_at` | `Timestamp` |  | `false` | The time the user was last updated at |
| `user_type` | [UserType](#usertype-userv1) |  | `false` | The type of the user (user, bot, or service) |
| `owner_user_id` | string |  | `false` | The ID of the user that owns the machine user, if a machine user |

### UserSummary (user.v1)

UserSummary is context aware

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` | The ID of the user |
| `name` | string |  | `false` | The name of the user |
| `email` | string |  | `false` | The email address of the user, if a human user |
| `role` | [UserRole](#userrole-user_rolev1) |  | `false` | The role the user has in the context in which the RPC returning this object (or list of objects) is called from |
| `created_at` | `Timestamp` |  | `false` | The time the user was created |
| `updated_at` | `Timestamp` |  | `false` | The time the user was last updated at |
| `user_type` | [UserType](#usertype-userv1) |  | `false` | The type of the user (user, bot, or service) |
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
| `user_type` | [UserType](#usertype-userv1) |  | `false` | The type of users to list |

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
| `default_role` | [UserRole](#userrole-user_rolev1) |  | `false` | The new default role of the user |

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

### UserType (user.v1)

| Name | Number | Description |
| --- | --- | --- |
| `USER_TYPE_UNSPECIFIED` | `0` | Default - should not use |
| `USER_TYPE_USER` | `1` | A human user |
| `USER_TYPE_BOT` | `2` | A bot user - perform actions in Terrabase |
| `USER_TYPE_SERVICE` | `3` | A service principal - grant access to another application |

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
| `role` | [UserRole](#userrole-user_rolev1) |  | `true` |  |

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
| `role` | [UserRole](#userrole-user_rolev1) |  | `true` |  |

### SetUserWorkspaceRoleResponse (user_membership.v1)

- (no fields)

## terrabase.user_role.v1

### UserRole (user_role.v1)

| Name | Number | Description |
| --- | --- | --- |
| `USER_ROLE_UNSPECIFIED` | `0` |  |
| `USER_ROLE_DEVELOPER` | `2` |  |
| `USER_ROLE_MAINTAINER` | `3` |  |
| `USER_ROLE_OWNER` | `4` |  |

## terrabase.workspace.v1

### WorkspaceService (workspace.v1)

| Name | Request | Response | Authentication Required | Required Scopes | Description |
| --- | --- | --- | --- | --- | --- |
| `CreateWorkspace` | [CreateWorkspaceRequest](#createworkspacerequest-workspacev1) | [CreateWorkspaceResponse](#createworkspaceresponse-workspacev1) | `false` |  | Create a new workspace |
| `GetWorkspace` | [GetWorkspaceRequest](#getworkspacerequest-workspacev1) | [GetWorkspaceResponse](#getworkspaceresponse-workspacev1) | `false` |  | Retrieve details about a single workspace |
| `ListWorkspaces` | [ListWorkspacesRequest](#listworkspacesrequest-workspacev1) | [ListWorkspacesResponse](#listworkspacesresponse-workspacev1) | `false` |  | List workspaces owned by a specific team, or belong to a specific application |
| `UpdateWorkspace` | [UpdateWorkspaceRequest](#updateworkspacerequest-workspacev1) | [UpdateWorkspaceResponse](#updateworkspaceresponse-workspacev1) | `false` |  | Change details about a workspace |
| `DeleteWorkspace` | [DeleteWorkspaceRequest](#deleteworkspacerequest-workspacev1) | [DeleteWorkspaceResponse](#deleteworkspaceresponse-workspacev1) | `false` |  | Delete a workspace |
| `GrantTeamAccess` | [GrantTeamAccessRequest](#grantteamaccessrequest-workspacev1) | [GrantTeamAccessResponse](#grantteamaccessresponse-workspacev1) | `false` |  | Grant access to a workspace to a single team or multiple teams |
| `RevokeTeamAccess` | [RevokeTeamAccessRequest](#revoketeamaccessrequest-workspacev1) | [RevokeTeamAccessResponse](#revoketeamaccessresponse-workspacev1) | `false` |  | Revoke access to a workspace from a single team or multiple teams |

### Workspace (workspace.v1)

A Terrabase workspace represents a single Terraform state file

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `false` | The unique ID of the workspace |
| `name` | string |  | `false` | The name of the workspace |
| `backend_type` | [BackendType](#backendtype-workspacev1) |  | `false` | The type of the backend configuration of the workspace |
| `environment_id` | string |  | `false` | The ID of the application environment the workspace belongs to. Mutually exclusive with `team_id`. A workspace does not have to belong to an application environment if it is owned by a team. |
| `team_id` | string |  | `false` | The ID of the team that owns this workspace. Mutually exclusive with `environment_id`. A workspace does not have to be owned by a team if it belongs to an application environment. |
| `created_at` | `Timestamp` |  | `false` | The time the workspace was created |
| `updated_at` | `Timestamp` |  | `false` | The time the workspace was last updated at |

### CreateWorkspaceRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `name` | string |  | `true` | The name of the workspace |
| `backend_type` | [BackendType](#backendtype-workspacev1) |  | `true` | The type of the backend configuration of the workspace |
| `environment_id` | string |  | `false` | The ID of the application environment the workspace belongs to. Mutually exclusive with `team_id`. A workspace does not have to belong to an application environment if it is owned by a team. |
| `team_id` | string |  | `false` | The ID of the team that owns this workspace. Mutually exclusive with `environment_id`. A workspace does not have to be owned by a team if it belongs to an application environment. |
| `s3_backend_config` | [CreateS3BackendConfigRequest](#creates3backendconfigrequest-s3_backend_configv1) |  | `false` | The S3 backend configuration, if the backend type is S3 |

### CreateWorkspaceResponse (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `workspace` | [Workspace](#workspace-workspacev1) |  | `false` | The workspace that was created |

### GetWorkspaceRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the workspace |

### GetWorkspaceResponse (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `workspace` | [Workspace](#workspace-workspacev1) |  | `false` | The workspace |

### ListWorkspacesRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `page_size` | int32 |  | `false` | The number of workspaces on each page of results |
| `page_token` | string |  | `false` | The token to retrieve the next page of results |
| `team_id` | string |  | `false` | The ID of the team to list all workspaces who are owned by the team |
| `application_id` | string |  | `false` | The ID of the application to list all workspaces who belong to the application |

### ListWorkspacesResponse (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `workspaces` | [Workspace](#workspace-workspacev1) | repeated | `false` | A list of workspaces |
| `next_page_token` | string |  | `false` | The token to retrieve the next page of results |

### UpdateWorkspaceRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the workspace to update |
| `name` | string |  | `false` | The new name of the workspace |
| `backend_type` | [BackendType](#backendtype-workspacev1) |  | `false` | The new backend type of the workspace |
| `environment_id` | string |  | `false` | The ID of the application environment the workspace belongs to. Mutually exclusive with `team_id`. A workspace does not have to belong to an application environment if it is owned by a team. |
| `team_id` | string |  | `false` | The ID of the team that owns this workspace. Mutually exclusive with `environment_id`. A workspace does not have to be owned by a team if it belongs to an application environment. |
| `s3_backend_config` | [CreateS3BackendConfigRequest](#creates3backendconfigrequest-s3_backend_configv1) |  | `false` | The new S3 backend configuration, if the backend type is S3 and a new configuration needs to be created |

### UpdateWorkspaceResponse (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `workspace` | [Workspace](#workspace-workspacev1) |  | `false` | The updated workspace |

### DeleteWorkspaceRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | string |  | `true` | The unique ID of the workspace to delete |

### DeleteWorkspaceResponse (workspace.v1)

- (no fields)

### GrantTeamAccessRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `workspace_id` | string |  | `true` | The unique ID of the workspace |
| `team_ids` | [TeamIds](#teamids-teamv1) |  | `true` | A list of team IDs who should be granted access to the workspace |

### GrantTeamAccessResponse (workspace.v1)

- (no fields)

### RevokeTeamAccessRequest (workspace.v1)

| Name | Type | Label | Required | Description |
| --- | --- | --- | --- | --- |
| `workspace_id` | string |  | `true` | The unique ID of the workspace |
| `team_ids` | [TeamIds](#teamids-teamv1) |  | `true` | A list of team IDs whose access to the workspace should be revoked |

### RevokeTeamAccessResponse (workspace.v1)

- (no fields)

### BackendType (workspace.v1)

| Name | Number | Description |
| --- | --- | --- |
| `BACKEND_TYPE_UNSPECIFIED` | `0` | Default - should not be used |
| `BACKEND_TYPE_S3` | `1` | S3 backend |
