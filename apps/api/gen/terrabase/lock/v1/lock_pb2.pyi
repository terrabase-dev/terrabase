import datetime

from google.api import field_behavior_pb2 as _field_behavior_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from terrabase.user.v1 import user_pb2 as _user_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Lock(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    WORKSPACE_ID_FIELD_NUMBER: _ClassVar[int]
    OWNER_FIELD_NUMBER: _ClassVar[int]
    INFO_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    EXPIRES_AT_FIELD_NUMBER: _ClassVar[int]
    id: str
    workspace_id: str
    owner: _user_pb2.User
    info: str
    created_at: _timestamp_pb2.Timestamp
    expires_at: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., workspace_id: _Optional[str] = ..., owner: _Optional[_Union[_user_pb2.User, _Mapping]] = ..., info: _Optional[str] = ..., created_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., expires_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class CreateLockRequest(_message.Message):
    __slots__ = ()
    WORKSPACE_ID_FIELD_NUMBER: _ClassVar[int]
    OWNER_FIELD_NUMBER: _ClassVar[int]
    INFO_FIELD_NUMBER: _ClassVar[int]
    workspace_id: str
    owner: _user_pb2.User
    info: str
    def __init__(self, workspace_id: _Optional[str] = ..., owner: _Optional[_Union[_user_pb2.User, _Mapping]] = ..., info: _Optional[str] = ...) -> None: ...

class CreateLockResponse(_message.Message):
    __slots__ = ()
    LOCK_FIELD_NUMBER: _ClassVar[int]
    lock: Lock
    def __init__(self, lock: _Optional[_Union[Lock, _Mapping]] = ...) -> None: ...

class GetLockRequest(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class GetLockResponse(_message.Message):
    __slots__ = ()
    LOCK_FIELD_NUMBER: _ClassVar[int]
    lock: Lock
    def __init__(self, lock: _Optional[_Union[Lock, _Mapping]] = ...) -> None: ...

class DeleteLockRequest(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class DeleteLockResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
