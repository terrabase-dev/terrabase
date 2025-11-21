import sys
import tomllib

from pathlib import Path

# Ensure generated protobuf modules (which use the proto package name as the
# import root, e.g., `terrabase.organization.v1`) are importable by adding the
# specs directory to sys.path.
_current_path = Path(__file__).resolve()
_specs_path = _current_path.parent / "specs"
_pyproject_path = _current_path.parent.parent / "pyproject.toml"

if _specs_path.exists():
    _specs_str = str(_specs_path)

    if _specs_str not in sys.path:
        sys.path.insert(0, _specs_str)

with open(_pyproject_path, "r") as f:
    _pyproject = tomllib.loads(f.read())

__version__ = _pyproject["project"]["version"]

del _current_path, _specs_path, _pyproject_path, _pyproject
