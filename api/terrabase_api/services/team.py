from typing import Annotated

import grpc

from fastapi import Depends

from terrabase_api.specs.terrabase.team.v1 import team_pb2_grpc
from terrabase_api.utils import RPC_ADDR


async def get_team_stub() -> team_pb2_grpc.TeamServiceStub:
    channel = grpc.aio.insecure_channel(RPC_ADDR)

    return team_pb2_grpc.TeamServiceStub(channel)


TeamService = Annotated[team_pb2_grpc.TeamServiceStub, Depends(get_team_stub)]
