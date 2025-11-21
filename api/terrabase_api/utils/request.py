from typing import Tuple

from fastapi import Query


def get_paginated_request_params() -> Tuple[int, str]:
    return (
        Query(
            description="The maximum number of records to return per page",
            default=50,
            required=False,
        ),
        Query(
            description="The token to get the next page of requests",
            default=None,
            required=False,
        ),
    )
