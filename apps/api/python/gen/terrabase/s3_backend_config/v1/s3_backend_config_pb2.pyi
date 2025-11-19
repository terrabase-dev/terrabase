import datetime

from google.api import field_behavior_pb2 as _field_behavior_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class S3BackendConfig(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    BUCKET_FIELD_NUMBER: _ClassVar[int]
    KEY_FIELD_NUMBER: _ClassVar[int]
    REGION_FIELD_NUMBER: _ClassVar[int]
    DYNAMODB_LOCK_FIELD_NUMBER: _ClassVar[int]
    ENCRYPT_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    id: str
    bucket: str
    key: str
    region: str
    dynamodb_lock: str
    encrypt: bool
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., bucket: _Optional[str] = ..., key: _Optional[str] = ..., region: _Optional[str] = ..., dynamodb_lock: _Optional[str] = ..., encrypt: _Optional[bool] = ..., created_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class CreateS3BackendConfigRequest(_message.Message):
    __slots__ = ()
    BUCKET_FIELD_NUMBER: _ClassVar[int]
    KEY_FIELD_NUMBER: _ClassVar[int]
    REGION_FIELD_NUMBER: _ClassVar[int]
    DYNAMODB_LOCK_FIELD_NUMBER: _ClassVar[int]
    ENCRYPT_FIELD_NUMBER: _ClassVar[int]
    bucket: str
    key: str
    region: str
    dynamodb_lock: str
    encrypt: bool
    def __init__(self, bucket: _Optional[str] = ..., key: _Optional[str] = ..., region: _Optional[str] = ..., dynamodb_lock: _Optional[str] = ..., encrypt: _Optional[bool] = ...) -> None: ...

class CreateS3BackendConfigResponse(_message.Message):
    __slots__ = ()
    S3_BACKEND_CONFIG_FIELD_NUMBER: _ClassVar[int]
    s3_backend_config: S3BackendConfig
    def __init__(self, s3_backend_config: _Optional[_Union[S3BackendConfig, _Mapping]] = ...) -> None: ...

class GetS3BackendConfigRequest(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class GetS3BackendConfigResponse(_message.Message):
    __slots__ = ()
    S3_BACKEND_CONFIG_FIELD_NUMBER: _ClassVar[int]
    s3_backend_config: S3BackendConfig
    def __init__(self, s3_backend_config: _Optional[_Union[S3BackendConfig, _Mapping]] = ...) -> None: ...

class UpdateS3BackendConfigRequest(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    BUCKET_FIELD_NUMBER: _ClassVar[int]
    KEY_FIELD_NUMBER: _ClassVar[int]
    REGION_FIELD_NUMBER: _ClassVar[int]
    DYNAMODB_LOCK_FIELD_NUMBER: _ClassVar[int]
    ENCRYPT_FIELD_NUMBER: _ClassVar[int]
    id: str
    bucket: str
    key: str
    region: str
    dynamodb_lock: str
    encrypt: bool
    def __init__(self, id: _Optional[str] = ..., bucket: _Optional[str] = ..., key: _Optional[str] = ..., region: _Optional[str] = ..., dynamodb_lock: _Optional[str] = ..., encrypt: _Optional[bool] = ...) -> None: ...

class UpdateS3BackendConfigResponse(_message.Message):
    __slots__ = ()
    S3_BACKEND_CONFIG_FIELD_NUMBER: _ClassVar[int]
    s3_backend_config: S3BackendConfig
    def __init__(self, s3_backend_config: _Optional[_Union[S3BackendConfig, _Mapping]] = ...) -> None: ...

class DeleteS3BackendConfigRequest(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class DeleteS3BackendConfigResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
