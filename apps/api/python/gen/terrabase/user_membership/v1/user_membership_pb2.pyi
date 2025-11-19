from google.api import field_behavior_pb2 as _field_behavior_pb2
from terrabase.user_role.v1 import user_role_pb2 as _user_role_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AddUserToOrganizationRequest(_message.Message):
    __slots__ = ()
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    ORGANIZATION_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    organization_id: str
    def __init__(self, user_id: _Optional[str] = ..., organization_id: _Optional[str] = ...) -> None: ...

class AddUserToOrganizationResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class RemoveUserFromOrganizationRequest(_message.Message):
    __slots__ = ()
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    ORGANIZATION_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    organization_id: str
    def __init__(self, user_id: _Optional[str] = ..., organization_id: _Optional[str] = ...) -> None: ...

class RemoveUserFromOrganizationResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class AddUserToTeamRequest(_message.Message):
    __slots__ = ()
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    team_id: str
    def __init__(self, user_id: _Optional[str] = ..., team_id: _Optional[str] = ...) -> None: ...

class AddUserToTeamResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class RemoveUserFromTeamRequest(_message.Message):
    __slots__ = ()
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    team_id: str
    def __init__(self, user_id: _Optional[str] = ..., team_id: _Optional[str] = ...) -> None: ...

class RemoveUserFromTeamResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class AddUserToWorkspaceRequest(_message.Message):
    __slots__ = ()
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    WORKSPACE_ID_FIELD_NUMBER: _ClassVar[int]
    ROLE_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    workspace_id: str
    role: _user_role_pb2.UserRole
    def __init__(self, user_id: _Optional[str] = ..., workspace_id: _Optional[str] = ..., role: _Optional[_Union[_user_role_pb2.UserRole, str]] = ...) -> None: ...

class AddUserToWorkspaceResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class RemoveUserFromWorkspaceRequest(_message.Message):
    __slots__ = ()
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    WORKSPACE_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    workspace_id: str
    def __init__(self, user_id: _Optional[str] = ..., workspace_id: _Optional[str] = ...) -> None: ...

class RemoveUserFromWorkspaceResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class SetUserWorkspaceRoleRequest(_message.Message):
    __slots__ = ()
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    WORKSPACE_ID_FIELD_NUMBER: _ClassVar[int]
    ROLE_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    workspace_id: str
    role: _user_role_pb2.UserRole
    def __init__(self, user_id: _Optional[str] = ..., workspace_id: _Optional[str] = ..., role: _Optional[_Union[_user_role_pb2.UserRole, str]] = ...) -> None: ...

class SetUserWorkspaceRoleResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
