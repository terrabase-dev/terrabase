from typing import Annotated

import grpc

from fastapi import Depends

from terrabase_api.specs.terrabase.user.v1 import user_pb2_grpc
from terrabase_api.utils import RPC_ADDR


async def get_user_stub() -> user_pb2_grpc.UserServiceStub:
    channel = grpc.aio.insecure_channel(RPC_ADDR)

    return user_pb2_grpc.UserServiceStub(channel)


UserService = Annotated[user_pb2_grpc.UserServiceStub, Depends(get_user_stub)]
