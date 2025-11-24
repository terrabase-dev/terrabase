from typing import Dict

from fastapi import APIRouter, Body
from starlette.status import (
    HTTP_400_BAD_REQUEST,
    HTTP_401_UNAUTHORIZED,
    HTTP_500_INTERNAL_SERVER_ERROR,
)

from terrabase_api.services import AuthService
from terrabase_api.specs.terrabase.auth.v1 import auth_pb2
from terrabase_api.types import ErrorResponse
from terrabase_api.utils import call, parse

auth_router = APIRouter(
    prefix="/auth",
    responses={
        HTTP_400_BAD_REQUEST: {"model": ErrorResponse},
        HTTP_401_UNAUTHORIZED: {"model": ErrorResponse},
        HTTP_500_INTERNAL_SERVER_ERROR: {"model": ErrorResponse},
    },
    tags=["auth"],
)


@auth_router.post("/signup", response_model=None)
async def signup(service: AuthService, req: Dict = Body(...)):
    return await call(service.Signup, parse(auth_pb2.SignupRequest, req))


@auth_router.post("/login", response_model=None)
async def login(service: AuthService, req: Dict = Body(...)):
    return await call(service.Login, parse(auth_pb2.LoginRequest, req))


@auth_router.post("/refresh", response_model=None)
async def refresh(service: AuthService, req: Dict = Body(...)):
    return await call(service.Refresh, parse(auth_pb2.RefreshRequest, req))


@auth_router.get("/me", response_model=None)
async def whoami(service: AuthService):
    return await call(service.WhoAmI, auth_pb2.WhoAmIRequest())


@auth_router.post("/logout", response_model=None)
async def logout(service: AuthService, req: Dict = Body(default={})):
    return await call(service.Logout, parse(auth_pb2.LogoutRequest, req))


@auth_router.get("/sessions", response_model=None)
async def list_sessions(service: AuthService):
    return await call(service.ListSessions, auth_pb2.ListSessionsRequest())


@auth_router.post("/machine-user", response_model=None)
async def create_machine_user(service: AuthService, req: Dict = Body(...)):
    return await call(
        service.CreateMachineUser, parse(auth_pb2.CreateMachineUserRequest, req)
    )


@auth_router.post("/api-keys", response_model=None)
async def create_api_key(service: AuthService, req: Dict = Body(...)):
    return await call(service.CreateApiKey, parse(auth_pb2.CreateApiKeyRequest, req))


@auth_router.get("/api-keys", response_model=None)
async def list_api_keys(
    service: AuthService,
    owner_type: str | None = None,
    owner_id: str | None = None,
):
    return await call(
        service.ListApiKeys,
        auth_pb2.ListApiKeysRequest(owner_type=owner_type, owner_id=owner_id or ""),
    )


@auth_router.post("/api-keys/{id}/revoke", response_model=None)
async def revoke_api_key(service: AuthService, id: str, req: Dict = Body(default={})):
    reason = req.get("reason")
    return await call(
        service.RevokeApiKey, auth_pb2.RevokeApiKeyRequest(id=id, reason=reason or "")
    )


@auth_router.post("/api-keys/{id}/rotate", response_model=None)
async def rotate_api_key(service: AuthService, id: str, req: Dict = Body(default={})):
    parsed = parse(auth_pb2.RotateApiKeyRequest, {"id": id, **req})
    parsed.id = id
    return await call(service.RotateApiKey, parsed)
