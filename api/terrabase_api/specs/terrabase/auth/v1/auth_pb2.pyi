import datetime

from google.api import field_behavior_pb2 as _field_behavior_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from terrabase.user.v1 import user_pb2 as _user_pb2
from terrabase.user_role.v1 import user_role_pb2 as _user_role_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Scope(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    SCOPE_UNSPECIFIED: _ClassVar[Scope]
    SCOPE_ADMIN: _ClassVar[Scope]
    SCOPE_ORG_WRITE: _ClassVar[Scope]
    SCOPE_ORG_READ: _ClassVar[Scope]
    SCOPE_TEAM_WRITE: _ClassVar[Scope]
    SCOPE_TEAM_READ: _ClassVar[Scope]
    SCOPE_APPLICATION_WRITE: _ClassVar[Scope]
    SCOPE_APPLICATION_READ: _ClassVar[Scope]
    SCOPE_ENVIRONMENT_WRITE: _ClassVar[Scope]
    SCOPE_ENVIRONMENT_READ: _ClassVar[Scope]
    SCOPE_WORKSPACE_WRITE: _ClassVar[Scope]
    SCOPE_WORKSPACE_READ: _ClassVar[Scope]

class ApiKeyOwnerType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    API_KEY_OWNER_TYPE_UNSPECIFIED: _ClassVar[ApiKeyOwnerType]
    API_KEY_OWNER_TYPE_USER: _ClassVar[ApiKeyOwnerType]
    API_KEY_OWNER_TYPE_BOT: _ClassVar[ApiKeyOwnerType]
    API_KEY_OWNER_TYPE_SERVICE: _ClassVar[ApiKeyOwnerType]
SCOPE_UNSPECIFIED: Scope
SCOPE_ADMIN: Scope
SCOPE_ORG_WRITE: Scope
SCOPE_ORG_READ: Scope
SCOPE_TEAM_WRITE: Scope
SCOPE_TEAM_READ: Scope
SCOPE_APPLICATION_WRITE: Scope
SCOPE_APPLICATION_READ: Scope
SCOPE_ENVIRONMENT_WRITE: Scope
SCOPE_ENVIRONMENT_READ: Scope
SCOPE_WORKSPACE_WRITE: Scope
SCOPE_WORKSPACE_READ: Scope
API_KEY_OWNER_TYPE_UNSPECIFIED: ApiKeyOwnerType
API_KEY_OWNER_TYPE_USER: ApiKeyOwnerType
API_KEY_OWNER_TYPE_BOT: ApiKeyOwnerType
API_KEY_OWNER_TYPE_SERVICE: ApiKeyOwnerType

class SignupRequest(_message.Message):
    __slots__ = ()
    NAME_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    DEFAULT_ROLE_FIELD_NUMBER: _ClassVar[int]
    name: str
    email: str
    password: str
    default_role: _user_role_pb2.UserRole
    def __init__(self, name: _Optional[str] = ..., email: _Optional[str] = ..., password: _Optional[str] = ..., default_role: _Optional[_Union[_user_role_pb2.UserRole, str]] = ...) -> None: ...

class SignupResponse(_message.Message):
    __slots__ = ()
    USER_FIELD_NUMBER: _ClassVar[int]
    ACCESS_TOKEN_FIELD_NUMBER: _ClassVar[int]
    REFRESH_TOKEN_FIELD_NUMBER: _ClassVar[int]
    user: _user_pb2.User
    access_token: str
    refresh_token: str
    def __init__(self, user: _Optional[_Union[_user_pb2.User, _Mapping]] = ..., access_token: _Optional[str] = ..., refresh_token: _Optional[str] = ...) -> None: ...

class LoginRequest(_message.Message):
    __slots__ = ()
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    email: str
    password: str
    def __init__(self, email: _Optional[str] = ..., password: _Optional[str] = ...) -> None: ...

class LoginResponse(_message.Message):
    __slots__ = ()
    USER_FIELD_NUMBER: _ClassVar[int]
    ACCESS_TOKEN_FIELD_NUMBER: _ClassVar[int]
    REFRESH_TOKEN_FIELD_NUMBER: _ClassVar[int]
    user: _user_pb2.User
    access_token: str
    refresh_token: str
    def __init__(self, user: _Optional[_Union[_user_pb2.User, _Mapping]] = ..., access_token: _Optional[str] = ..., refresh_token: _Optional[str] = ...) -> None: ...

class RefreshRequest(_message.Message):
    __slots__ = ()
    REFRESH_TOKEN_FIELD_NUMBER: _ClassVar[int]
    refresh_token: str
    def __init__(self, refresh_token: _Optional[str] = ...) -> None: ...

class RefreshResponse(_message.Message):
    __slots__ = ()
    ACCESS_TOKEN_FIELD_NUMBER: _ClassVar[int]
    REFRESH_TOKEN_FIELD_NUMBER: _ClassVar[int]
    access_token: str
    refresh_token: str
    def __init__(self, access_token: _Optional[str] = ..., refresh_token: _Optional[str] = ...) -> None: ...

class WhoAmIRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class WhoAmIResponse(_message.Message):
    __slots__ = ()
    USER_FIELD_NUMBER: _ClassVar[int]
    SCOPES_FIELD_NUMBER: _ClassVar[int]
    user: _user_pb2.User
    scopes: _containers.RepeatedScalarFieldContainer[Scope]
    def __init__(self, user: _Optional[_Union[_user_pb2.User, _Mapping]] = ..., scopes: _Optional[_Iterable[_Union[Scope, str]]] = ...) -> None: ...

class LogoutRequest(_message.Message):
    __slots__ = ()
    SESSION_ID_FIELD_NUMBER: _ClassVar[int]
    session_id: str
    def __init__(self, session_id: _Optional[str] = ...) -> None: ...

class LogoutResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class CreateMachineUserRequest(_message.Message):
    __slots__ = ()
    NAME_FIELD_NUMBER: _ClassVar[int]
    DEFAULT_ROLE_FIELD_NUMBER: _ClassVar[int]
    USER_TYPE_FIELD_NUMBER: _ClassVar[int]
    OWNER_USER_ID_FIELD_NUMBER: _ClassVar[int]
    name: str
    default_role: _user_role_pb2.UserRole
    user_type: _user_pb2.UserType
    owner_user_id: str
    def __init__(self, name: _Optional[str] = ..., default_role: _Optional[_Union[_user_role_pb2.UserRole, str]] = ..., user_type: _Optional[_Union[_user_pb2.UserType, str]] = ..., owner_user_id: _Optional[str] = ...) -> None: ...

class CreateMachineUserResponse(_message.Message):
    __slots__ = ()
    MACHINE_USER_FIELD_NUMBER: _ClassVar[int]
    machine_user: _user_pb2.User
    def __init__(self, machine_user: _Optional[_Union[_user_pb2.User, _Mapping]] = ...) -> None: ...

class ApiKey(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    SCOPES_FIELD_NUMBER: _ClassVar[int]
    OWNER_ID_FIELD_NUMBER: _ClassVar[int]
    OWNER_TYPE_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    EXPIRES_AT_FIELD_NUMBER: _ClassVar[int]
    LAST_USED_AT_FIELD_NUMBER: _ClassVar[int]
    REVOKED_AT_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    scopes: _containers.RepeatedScalarFieldContainer[Scope]
    owner_id: str
    owner_type: ApiKeyOwnerType
    created_at: _timestamp_pb2.Timestamp
    expires_at: _timestamp_pb2.Timestamp
    last_used_at: _timestamp_pb2.Timestamp
    revoked_at: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., scopes: _Optional[_Iterable[_Union[Scope, str]]] = ..., owner_id: _Optional[str] = ..., owner_type: _Optional[_Union[ApiKeyOwnerType, str]] = ..., created_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., expires_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., last_used_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., revoked_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class CreateApiKeyRequest(_message.Message):
    __slots__ = ()
    NAME_FIELD_NUMBER: _ClassVar[int]
    OWNER_TYPE_FIELD_NUMBER: _ClassVar[int]
    OWNER_ID_FIELD_NUMBER: _ClassVar[int]
    SCOPES_FIELD_NUMBER: _ClassVar[int]
    TTL_HOURS_FIELD_NUMBER: _ClassVar[int]
    name: str
    owner_type: ApiKeyOwnerType
    owner_id: str
    scopes: _containers.RepeatedScalarFieldContainer[Scope]
    ttl_hours: int
    def __init__(self, name: _Optional[str] = ..., owner_type: _Optional[_Union[ApiKeyOwnerType, str]] = ..., owner_id: _Optional[str] = ..., scopes: _Optional[_Iterable[_Union[Scope, str]]] = ..., ttl_hours: _Optional[int] = ...) -> None: ...

class CreateApiKeyResponse(_message.Message):
    __slots__ = ()
    API_KEY_TOKEN_FIELD_NUMBER: _ClassVar[int]
    API_KEY_FIELD_NUMBER: _ClassVar[int]
    api_key_token: str
    api_key: ApiKey
    def __init__(self, api_key_token: _Optional[str] = ..., api_key: _Optional[_Union[ApiKey, _Mapping]] = ...) -> None: ...

class ListApiKeysRequest(_message.Message):
    __slots__ = ()
    OWNER_TYPE_FIELD_NUMBER: _ClassVar[int]
    OWNER_ID_FIELD_NUMBER: _ClassVar[int]
    owner_type: ApiKeyOwnerType
    owner_id: str
    def __init__(self, owner_type: _Optional[_Union[ApiKeyOwnerType, str]] = ..., owner_id: _Optional[str] = ...) -> None: ...

class ListApiKeysResponse(_message.Message):
    __slots__ = ()
    API_KEYS_FIELD_NUMBER: _ClassVar[int]
    api_keys: _containers.RepeatedCompositeFieldContainer[ApiKey]
    def __init__(self, api_keys: _Optional[_Iterable[_Union[ApiKey, _Mapping]]] = ...) -> None: ...

class RevokeApiKeyRequest(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    REASON_FIELD_NUMBER: _ClassVar[int]
    id: str
    reason: str
    def __init__(self, id: _Optional[str] = ..., reason: _Optional[str] = ...) -> None: ...

class RevokeApiKeyResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class RotateApiKeyRequest(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    SCOPES_FIELD_NUMBER: _ClassVar[int]
    TTL_HOURS_FIELD_NUMBER: _ClassVar[int]
    id: str
    scopes: _containers.RepeatedScalarFieldContainer[Scope]
    ttl_hours: int
    def __init__(self, id: _Optional[str] = ..., scopes: _Optional[_Iterable[_Union[Scope, str]]] = ..., ttl_hours: _Optional[int] = ...) -> None: ...

class RotateApiKeyResponse(_message.Message):
    __slots__ = ()
    API_KEY_TOKEN_FIELD_NUMBER: _ClassVar[int]
    API_KEY_FIELD_NUMBER: _ClassVar[int]
    api_key_token: str
    api_key: ApiKey
    def __init__(self, api_key_token: _Optional[str] = ..., api_key: _Optional[_Union[ApiKey, _Mapping]] = ...) -> None: ...

class Session(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    USER_AGENT_FIELD_NUMBER: _ClassVar[int]
    IP_FIELD_NUMBER: _ClassVar[int]
    EXPIRES_AT_FIELD_NUMBER: _ClassVar[int]
    LAST_USED_AT_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    id: str
    user_agent: str
    ip: str
    expires_at: _timestamp_pb2.Timestamp
    last_used_at: _timestamp_pb2.Timestamp
    created_at: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., user_agent: _Optional[str] = ..., ip: _Optional[str] = ..., expires_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., last_used_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., created_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class ListSessionsRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ListSessionsResponse(_message.Message):
    __slots__ = ()
    SESSIONS_FIELD_NUMBER: _ClassVar[int]
    sessions: _containers.RepeatedCompositeFieldContainer[Session]
    def __init__(self, sessions: _Optional[_Iterable[_Union[Session, _Mapping]]] = ...) -> None: ...
