from app.deocders.DecoderHandlers import DecoderHandler
from app.interfaces.decoder import LLMResponseContext_PM
from app.settings.prompts import PromptEnums


class DecoderController:

    @staticmethod
    def get_llm_response(payload: LLMResponseContext_PM) -> str:
        code_chunks_builder = []
        for i, ctx in enumerate(payload.contexts):
            code_chunk = {"file_path": ctx.FilePath, "code_chunk": ctx.Code}
            code_chunks_builder.append(code_chunk)
        system_prompt = PromptEnums.SYSTEM_PROMPT.value
        user_prompt = PromptEnums.USER_PROMPT.value.format(
            user_query=payload.query, code_snippet=str(code_chunks_builder)
        )
        llm_response = DecoderHandler().generate_llm_reponse(
            system_prompt=system_prompt, user_prompt=user_prompt
        )
        return llm_response
