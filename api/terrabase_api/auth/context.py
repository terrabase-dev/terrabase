from __future__ import annotations

from contextvars import ContextVar
from dataclasses import dataclass, field
from typing import Any, Dict, List, Optional


@dataclass
class AuthContext:
    subject_id: str
    subject_type: str
    name: Optional[str] = None
    email: Optional[str] = None
    default_role: Optional[int] = None
    scopes: List[str] = field(default_factory=list)
    entitlements: Dict[str, List[str]] = field(default_factory=dict)
    metadata: Dict[str, Any] = field(default_factory=dict)
    token_id: Optional[str] = None
    auth_scheme: str = "bearer"
    raw_authorization: Optional[str] = None


_current_auth_context: ContextVar[Optional[AuthContext]] = ContextVar(
    "auth_context", default=None
)


@dataclass
class ClientInfo:
    ip: Optional[str] = None
    user_agent: Optional[str] = None


_current_client_info: ContextVar[Optional[ClientInfo]] = ContextVar(
    "client_info", default=None
)


def set_auth_context(ctx: Optional[AuthContext]) -> None:
    _current_auth_context.set(ctx)


def get_auth_context() -> Optional[AuthContext]:
    return _current_auth_context.get()


def set_client_info(info: Optional[ClientInfo]) -> None:
    _current_client_info.set(info)


def get_client_info() -> Optional[ClientInfo]:
    return _current_client_info.get()
