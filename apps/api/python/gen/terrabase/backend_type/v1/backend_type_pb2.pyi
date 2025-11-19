from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class BackendType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    BACKEND_TYPE_UNSPECIFIED: _ClassVar[BackendType]
    BACKEND_TYPE_S3: _ClassVar[BackendType]
BACKEND_TYPE_UNSPECIFIED: BackendType
BACKEND_TYPE_S3: BackendType
