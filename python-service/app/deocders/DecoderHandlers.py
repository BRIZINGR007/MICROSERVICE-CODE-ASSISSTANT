from typing import cast, final
from zoldics_service_utils.ioc import SingletonMeta
from llama_cpp import Llama


class DecoderHandler(metaclass=SingletonMeta):
    def __init__(self) -> None:
        self.__llm = Llama(
            model_path="MaziyarPanahi/Llama-3.2-1B-Instruct-GGUF/Llama-3.2-1B-Instruct.Q8_0.gguf",
            n_ctx=4048,
        )

    @staticmethod
    def get_system_prompt(system_prompt: str) -> str:
        bos_token = "<|begin_of_text|>"
        system_start_token = "<|start_header_id|>system<|end_header_id|>"
        content_end_token = "<|eot_id|>"
        structured_prompt = (
            bos_token + system_start_token + system_prompt + content_end_token
        )
        return structured_prompt

    @staticmethod
    def get_user_prompt(user_prompt: str) -> str:
        bos_token = "<|begin_of_text|>"
        user_start_token = "<|start_header_id|>user<|end_header_id|>"
        content_end_token = "<|eot_id|>"
        structured_prompt = (
            bos_token
            + user_start_token
            + user_prompt
            + content_end_token
            + "<|start_header_id|>assistant<|end_header_id|>"
        )
        return structured_prompt

    def generate_llm_reponse(self, system_prompt: str, user_prompt: str) -> str:
        system_prompt = self.get_system_prompt(system_prompt)
        user_prompt = self.get_user_prompt(user_prompt)
        response = self.__llm.create_chat_completion(
            messages=[
                {"role": "user", "content": user_prompt},
            ],
            max_tokens=256,
        )
        return cast(str, response["choices"][0]["message"]["content"])
