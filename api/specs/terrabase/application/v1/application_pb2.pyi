import datetime

from google.api import field_behavior_pb2 as _field_behavior_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Application(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., created_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class CreateApplicationRequest(_message.Message):
    __slots__ = ()
    NAME_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    name: str
    team_id: str
    def __init__(self, name: _Optional[str] = ..., team_id: _Optional[str] = ...) -> None: ...

class CreateApplicationResponse(_message.Message):
    __slots__ = ()
    APPLICATION_FIELD_NUMBER: _ClassVar[int]
    application: Application
    def __init__(self, application: _Optional[_Union[Application, _Mapping]] = ...) -> None: ...

class GetApplicationRequest(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class GetApplicationResponse(_message.Message):
    __slots__ = ()
    APPLICATION_FIELD_NUMBER: _ClassVar[int]
    application: Application
    def __init__(self, application: _Optional[_Union[Application, _Mapping]] = ...) -> None: ...

class ListApplicationsRequest(_message.Message):
    __slots__ = ()
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
    PAGE_TOKEN_FIELD_NUMBER: _ClassVar[int]
    team_id: str
    page_size: int
    page_token: str
    def __init__(self, team_id: _Optional[str] = ..., page_size: _Optional[int] = ..., page_token: _Optional[str] = ...) -> None: ...

class ListApplicationsResponse(_message.Message):
    __slots__ = ()
    APPLICATIONS_FIELD_NUMBER: _ClassVar[int]
    NEXT_PAGE_TOKEN_FIELD_NUMBER: _ClassVar[int]
    applications: _containers.RepeatedCompositeFieldContainer[Application]
    next_page_token: str
    def __init__(self, applications: _Optional[_Iterable[_Union[Application, _Mapping]]] = ..., next_page_token: _Optional[str] = ...) -> None: ...

class UpdateApplicationRequest(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    team_id: str
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., team_id: _Optional[str] = ...) -> None: ...

class UpdateApplicationResponse(_message.Message):
    __slots__ = ()
    APPLICATION_FIELD_NUMBER: _ClassVar[int]
    application: Application
    def __init__(self, application: _Optional[_Union[Application, _Mapping]] = ...) -> None: ...

class DeleteApplicationRequest(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class DeleteApplicationResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GrantTeamAccessRequest(_message.Message):
    __slots__ = ()
    APPLICATION_ID_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    application_id: str
    team_id: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, application_id: _Optional[str] = ..., team_id: _Optional[_Iterable[str]] = ...) -> None: ...

class GrantTeamAccessResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class RevokeTeamAccessRequest(_message.Message):
    __slots__ = ()
    APPLICATION_ID_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    application_id: str
    team_id: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, application_id: _Optional[str] = ..., team_id: _Optional[_Iterable[str]] = ...) -> None: ...

class RevokeTeamAccessResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
