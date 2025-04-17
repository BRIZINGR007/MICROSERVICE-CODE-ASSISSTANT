from typing import List
from pydantic import BaseModel


class Contexts_PM(BaseModel):
    Code: str
    FilePath: str


class LLMResponseContext_PM(BaseModel):
    query: str
    contexts: List[Contexts_PM]
