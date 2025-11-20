import datetime

from google.api import field_behavior_pb2 as _field_behavior_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class StateVersion(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    WORKSPACE_ID_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    HASH_FIELD_NUMBER: _ClassVar[int]
    SIZE_BYTES_FIELD_NUMBER: _ClassVar[int]
    STORAGE_KEY_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    id: str
    workspace_id: str
    version: str
    hash: str
    size_bytes: int
    storage_key: str
    created_at: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., workspace_id: _Optional[str] = ..., version: _Optional[str] = ..., hash: _Optional[str] = ..., size_bytes: _Optional[int] = ..., storage_key: _Optional[str] = ..., created_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class CreateStateVersionRequest(_message.Message):
    __slots__ = ()
    WORKSPACE_ID_FIELD_NUMBER: _ClassVar[int]
    STORAGE_KEY_FIELD_NUMBER: _ClassVar[int]
    workspace_id: str
    storage_key: str
    def __init__(self, workspace_id: _Optional[str] = ..., storage_key: _Optional[str] = ...) -> None: ...

class CreateStateVersionResponse(_message.Message):
    __slots__ = ()
    STATE_VERSION_FIELD_NUMBER: _ClassVar[int]
    state_version: StateVersion
    def __init__(self, state_version: _Optional[_Union[StateVersion, _Mapping]] = ...) -> None: ...

class GetStateVersionRequest(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class GetStateVersionResponse(_message.Message):
    __slots__ = ()
    STATE_VERSION_FIELD_NUMBER: _ClassVar[int]
    state_version: StateVersion
    def __init__(self, state_version: _Optional[_Union[StateVersion, _Mapping]] = ...) -> None: ...

class ListStateVersionsRequest(_message.Message):
    __slots__ = ()
    WORKSPACE_ID_FIELD_NUMBER: _ClassVar[int]
    PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
    PAGE_TOKEN_FIELD_NUMBER: _ClassVar[int]
    workspace_id: str
    page_size: int
    page_token: str
    def __init__(self, workspace_id: _Optional[str] = ..., page_size: _Optional[int] = ..., page_token: _Optional[str] = ...) -> None: ...

class ListStateVersionsResponse(_message.Message):
    __slots__ = ()
    STATE_VERSIONS_FIELD_NUMBER: _ClassVar[int]
    NEXT_PAGE_TOKEN_FIELD_NUMBER: _ClassVar[int]
    state_versions: _containers.RepeatedCompositeFieldContainer[StateVersion]
    next_page_token: str
    def __init__(self, state_versions: _Optional[_Iterable[_Union[StateVersion, _Mapping]]] = ..., next_page_token: _Optional[str] = ...) -> None: ...

class DeleteStateVersionRequest(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class DeleteStateVersionResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
