from typing import Annotated

import grpc

from fastapi import Depends

from terrabase_api.specs.terrabase.auth.v1 import auth_pb2_grpc
from terrabase_api.utils import RPC_ADDR


async def get_auth_stub() -> auth_pb2_grpc.AuthServiceStub:
    channel = grpc.aio.insecure_channel(RPC_ADDR)

    return auth_pb2_grpc.AuthServiceStub(channel)


AuthService = Annotated[auth_pb2_grpc.AuthServiceStub, Depends(get_auth_stub)]
