import datetime

from google.api import field_behavior_pb2 as _field_behavior_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Drift(_message.Message):
    __slots__ = ()
    class ChangesEntry(_message.Message):
        __slots__ = ()
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    RESOURCE_ID_FIELD_NUMBER: _ClassVar[int]
    CHANGES_FIELD_NUMBER: _ClassVar[int]
    resource_id: str
    changes: _containers.ScalarMap[str, str]
    def __init__(self, resource_id: _Optional[str] = ..., changes: _Optional[_Mapping[str, str]] = ...) -> None: ...

class DriftReport(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    WORKSPACE_ID_FIELD_NUMBER: _ClassVar[int]
    DRIFTED_FIELD_NUMBER: _ClassVar[int]
    CHANGES_FIELD_NUMBER: _ClassVar[int]
    DETECTED_AT_FIELD_NUMBER: _ClassVar[int]
    id: str
    workspace_id: str
    drifted: bool
    changes: _containers.RepeatedCompositeFieldContainer[Drift]
    detected_at: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., workspace_id: _Optional[str] = ..., drifted: _Optional[bool] = ..., changes: _Optional[_Iterable[_Union[Drift, _Mapping]]] = ..., detected_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class CreateDriftReportRequest(_message.Message):
    __slots__ = ()
    WORKSPACE_ID_FIELD_NUMBER: _ClassVar[int]
    DRIFTED_FIELD_NUMBER: _ClassVar[int]
    CHANGES_FIELD_NUMBER: _ClassVar[int]
    workspace_id: str
    drifted: bool
    changes: _containers.RepeatedCompositeFieldContainer[Drift]
    def __init__(self, workspace_id: _Optional[str] = ..., drifted: _Optional[bool] = ..., changes: _Optional[_Iterable[_Union[Drift, _Mapping]]] = ...) -> None: ...

class CreateDriftReportResponse(_message.Message):
    __slots__ = ()
    DRIFT_REPORT_FIELD_NUMBER: _ClassVar[int]
    drift_report: DriftReport
    def __init__(self, drift_report: _Optional[_Union[DriftReport, _Mapping]] = ...) -> None: ...

class GetDriftReportRequest(_message.Message):
    __slots__ = ()
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class GetDriftReportResponse(_message.Message):
    __slots__ = ()
    DRIFT_REPORT_FIELD_NUMBER: _ClassVar[int]
    drift_report: DriftReport
    def __init__(self, drift_report: _Optional[_Union[DriftReport, _Mapping]] = ...) -> None: ...
