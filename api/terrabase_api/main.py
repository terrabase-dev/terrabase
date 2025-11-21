from importlib.metadata import version

from fastapi import FastAPI

terrabase_api = FastAPI(
    title="Terrabase API", version=version("terrabase-api"), root_path="/api"
)
