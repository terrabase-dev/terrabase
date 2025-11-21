import io
import yaml

from typing import Any, Dict

from fastapi import FastAPI
from fastapi.openapi.utils import get_openapi


def configure_openapi(api: FastAPI) -> Dict[str, Any]:
    if api.openapi_schema:
        return api.openapi_schema

    openapi_schema = get_openapi(
        title=api.title,
        version=api.version,
        routes=api.routes,
        openapi_version="3.0.4",
    )

    # TODO: add logo
    # openapi_schema["info"]["x-logo"] = {"url": ""}

    api.openapi_schema = openapi_schema

    return api.openapi_schema


def get_openapi_yaml(api: FastAPI) -> str:
    openapi_json = api.openapi()
    yaml_stream = io.StringIO()
    yaml.dump(openapi_json, yaml_stream, sort_keys=False)

    return yaml_stream.getvalue()
