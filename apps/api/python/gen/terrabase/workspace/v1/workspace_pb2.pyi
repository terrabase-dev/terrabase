import datetime

from google.api import field_behavior_pb2 as _field_behavior_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from terrabase.backend_type.v1 import backend_type_pb2 as _backend_type_pb2
from terrabase.s3_backend_config.v1 import s3_backend_config_pb2 as _s3_backend_config_pb2
from terrabase.team.v1 import team_pb2 as _team_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Workspace(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    BACKEND_TYPE_FIELD_NUMBER: _ClassVar[int]
    ENVIRONMENT_ID_FIELD_NUMBER: _ClassVar[int]
    S3_BACKEND_CONFIG_ID_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    backend_type: _backend_type_pb2.BackendType
    environment_id: str
    s3_backend_config_id: str
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., backend_type: _Optional[_Union[_backend_type_pb2.BackendType, str]] = ..., environment_id: _Optional[str] = ..., s3_backend_config_id: _Optional[str] = ..., created_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class CreateWorkspaceRequest(_message.Message):
    __slots__ = ()
    NAME_FIELD_NUMBER: _ClassVar[int]
    BACKEND_TYPE_FIELD_NUMBER: _ClassVar[int]
    ENVIRONMENT_ID_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    S3_BACKEND_CONFIG_FIELD_NUMBER: _ClassVar[int]
    name: str
    backend_type: _backend_type_pb2.BackendType
    environment_id: str
    team_id: str
    s3_backend_config: _s3_backend_config_pb2.CreateS3BackendConfigRequest
    def __init__(self, name: _Optional[str] = ..., backend_type: _Optional[_Union[_backend_type_pb2.BackendType, str]] = ..., environment_id: _Optional[str] = ..., team_id: _Optional[str] = ..., s3_backend_config: _Optional[_Union[_s3_backend_config_pb2.CreateS3BackendConfigRequest, _Mapping]] = ...) -> None: ...

class CreateWorkspaceResponse(_message.Message):
    __slots__ = ()
    WORKSPACE_FIELD_NUMBER: _ClassVar[int]
    workspace: Workspace
    def __init__(self, workspace: _Optional[_Union[Workspace, _Mapping]] = ...) -> None: ...

class GetWorkspaceRequest(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class GetWorkspaceResponse(_message.Message):
    __slots__ = ()
    WORKSPACE_FIELD_NUMBER: _ClassVar[int]
    workspace: Workspace
    def __init__(self, workspace: _Optional[_Union[Workspace, _Mapping]] = ...) -> None: ...

class ListWorkspacesRequest(_message.Message):
    __slots__ = ()
    PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
    PAGE_TOKEN_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    ENVIRONMENT_ID_FIELD_NUMBER: _ClassVar[int]
    page_size: int
    page_token: str
    team_id: str
    environment_id: str
    def __init__(self, page_size: _Optional[int] = ..., page_token: _Optional[str] = ..., team_id: _Optional[str] = ..., environment_id: _Optional[str] = ...) -> None: ...

class ListWorkspacesResponse(_message.Message):
    __slots__ = ()
    WORKSPACES_FIELD_NUMBER: _ClassVar[int]
    NEXT_PAGE_TOKEN_FIELD_NUMBER: _ClassVar[int]
    workspaces: _containers.RepeatedCompositeFieldContainer[Workspace]
    next_page_token: str
    def __init__(self, workspaces: _Optional[_Iterable[_Union[Workspace, _Mapping]]] = ..., next_page_token: _Optional[str] = ...) -> None: ...

class UpdateWorkspaceRequest(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    BACKEND_TYPE_FIELD_NUMBER: _ClassVar[int]
    ENVIRONMENT_ID_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    S3_BACKEND_CONFIG_ID_FIELD_NUMBER: _ClassVar[int]
    S3_BACKEND_CONFIG_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    backend_type: _backend_type_pb2.BackendType
    environment_id: str
    team_id: str
    s3_backend_config_id: str
    s3_backend_config: _s3_backend_config_pb2.CreateS3BackendConfigRequest
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., backend_type: _Optional[_Union[_backend_type_pb2.BackendType, str]] = ..., environment_id: _Optional[str] = ..., team_id: _Optional[str] = ..., s3_backend_config_id: _Optional[str] = ..., s3_backend_config: _Optional[_Union[_s3_backend_config_pb2.CreateS3BackendConfigRequest, _Mapping]] = ...) -> None: ...

class UpdateWorkspaceResponse(_message.Message):
    __slots__ = ()
    WORKSPACE_FIELD_NUMBER: _ClassVar[int]
    workspace: Workspace
    def __init__(self, workspace: _Optional[_Union[Workspace, _Mapping]] = ...) -> None: ...

class DeleteWorkspaceRequest(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class DeleteWorkspaceResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GrantTeamAccessRequest(_message.Message):
    __slots__ = ()
    WORKSPACE_ID_FIELD_NUMBER: _ClassVar[int]
    TEAM_IDS_FIELD_NUMBER: _ClassVar[int]
    workspace_id: str
    team_ids: _team_pb2.TeamIds
    def __init__(self, workspace_id: _Optional[str] = ..., team_ids: _Optional[_Union[_team_pb2.TeamIds, _Mapping]] = ...) -> None: ...

class GrantTeamAccessResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class RevokeTeamAccessRequest(_message.Message):
    __slots__ = ()
    WORKSPACE_ID_FIELD_NUMBER: _ClassVar[int]
    TEAM_IDS_FIELD_NUMBER: _ClassVar[int]
    workspace_id: str
    team_ids: _team_pb2.TeamIds
    def __init__(self, workspace_id: _Optional[str] = ..., team_ids: _Optional[_Union[_team_pb2.TeamIds, _Mapping]] = ...) -> None: ...

class RevokeTeamAccessResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
