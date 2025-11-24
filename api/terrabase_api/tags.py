from typing import Dict, List, TypedDict, cast


class TagMetadata(TypedDict):
    name: str
    description: str


tags_metadata: List[TagMetadata] = [
    {
        "name": "auth",
        "description": "Manage user/machine authentication and authorization",
    },
    {"name": "organization", "description": "Manage Terrabase organizations"},
]

openapi_tags = cast(List[Dict[str, str]], tags_metadata)
