from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class TeamAccessType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    TEAM_ACCESS_TYPE_UNSPECIFIED: _ClassVar[TeamAccessType]
    TEAM_ACCESS_TYPE_OWNER: _ClassVar[TeamAccessType]
    TEAM_ACCESS_TYPE_GRANTED: _ClassVar[TeamAccessType]
TEAM_ACCESS_TYPE_UNSPECIFIED: TeamAccessType
TEAM_ACCESS_TYPE_OWNER: TeamAccessType
TEAM_ACCESS_TYPE_GRANTED: TeamAccessType
