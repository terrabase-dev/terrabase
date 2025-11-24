from typing import Dict

from fastapi import APIRouter, Body, Path
from starlette.status import (
    HTTP_400_BAD_REQUEST,
    HTTP_401_UNAUTHORIZED,
    HTTP_500_INTERNAL_SERVER_ERROR,
)

from terrabase_api.services import UserService
from terrabase_api.specs.terrabase.user.v1 import user_pb2
from terrabase_api.types import ErrorResponse
from terrabase_api.utils import call, parse

user_router = APIRouter(
    prefix="/user",
    responses={
        HTTP_400_BAD_REQUEST: {"model": ErrorResponse},
        HTTP_401_UNAUTHORIZED: {"model": ErrorResponse},
        HTTP_500_INTERNAL_SERVER_ERROR: {"model": ErrorResponse},
    },
    tags=["auth"],
)


@user_router.get("/{id}", response_model=None)
async def get_user(
    service: UserService, id: str = Path(description="The ID of the user to retrieve")
):
    return await call(service.GetUser, user_pb2.GetUserRequest(id=id))


@user_router.post("/{id}", response_model=None)
async def update_user(
    service: UserService,
    req: Dict = Body(...),
    id: str = Path(description="The ID of the user to update"),
):
    update_req = parse(user_pb2.UpdateUserRequest, req)
    update_req.id = id

    return await call(service.UpdateUser, update_req)


@user_router.delete("/{id}", response_model=None)
async def delete_user(
    service: UserService, id: str = Path(description="The ID of the user to delete")
):
    return await call(service.DeleteUser, user_pb2.DeleteUserRequest(id=id))
