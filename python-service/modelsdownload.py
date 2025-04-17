import requests
from pathlib import Path


import requests
from pathlib import Path


def download_hf_file(model_repo: str, filename: str, output_dir: str):
    base_url = f"https://huggingface.co/{model_repo}/resolve/main/{filename}"
    output_path = Path(output_dir) / filename
    output_path.parent.mkdir(parents=True, exist_ok=True)

    if output_path.exists():
        print(f"File already exists: {output_path}")
        return

    response = requests.get(base_url, stream=True)
    if response.status_code == 200:
        with open(output_path, "wb") as f:
            for chunk in response.iter_content(chunk_size=8192):
                f.write(chunk)
        print(f"Downloaded: {filename}")
    else:
        print(f"Failed to download {filename}: {response.status_code}")


def download_bge_model_files():
    files = [
        "config.json",
        "pytorch_model.bin",
        "tokenizer.json",
        "tokenizer_config.json",
        "vocab.txt",
    ]
    model_repo = "BAAI/bge-small-en-v1.5"
    output_dir = "BAAI/bge-small-en-v1.5"

    for file in files:
        download_hf_file(model_repo, file, output_dir)


def download_model(url, output_path):
    response = requests.get(url, stream=True)
    output_path = Path(output_path)
    output_path.parent.mkdir(parents=True, exist_ok=True)

    if response.status_code == 200:
        with open(output_path, "wb") as f:
            for chunk in response.iter_content(chunk_size=8192):
                f.write(chunk)
        print(f"Model downloaded to {output_path}")
    else:
        print(f"Failed to download model: {response.status_code}")


if __name__ == "__main__":
    model_url = "https://huggingface.co/bartowski/Llama-3.2-1B-Instruct-GGUF/resolve/main/Llama-3.2-1B-Instruct-IQ3_M.gguf"
    output_file = (
        "bartowski/Llama-3.2-1B-Instruct-GGUF/Llama-3.2-1B-Instruct-IQ3_M.gguf"
    )

    download_model(model_url, output_file)
    download_bge_model_files()
