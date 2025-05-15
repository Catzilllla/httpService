import requests
import json
import os
from typing import Dict, Any
import yaml
from pathlib import Path
from typing import Dict, Any


API_URL = "https://jsonplaceholder.typicode.com/posts"

# Пример данных для POST-запроса
EXAMPLE_DATA = {
    "name": "John Doe",
    "email": "john@example.com",
    "age": 30
}

def send_post_request(url: str, data: Dict[str, Any]) -> Dict[str, Any]:
    """
    Отправляет POST-запрос с JSON-данными.
    Возвращает ответ сервера в виде словаря.
    """
    try:
        # Устанавливаем заголовок JSON и отправляем POST
        response = requests.post(
            url,
            json=data,  # автоматически сериализует в JSON
            headers={"Content-Type": "application/json"},
            timeout=10  # таймаут 10 секунд
        )
        response.raise_for_status()  # Проверка на ошибки HTTP (4xx, 5xx)
        return response.json()  # Парсим JSON-ответ
    except requests.exceptions.RequestException as e:
        print(f"Ошибка POST-запроса: {e}")
        return {"error": str(e)}

def send_get_request(url: str) -> Dict[str, Any]:
    """
    Отправляет GET-запрос.
    Возвращает ответ сервера в виде словаря.
    """
    try:
        response = requests.get(url, timeout=10)
        response.raise_for_status()
        return response.json()
    except requests.exceptions.RequestException as e:
        print(f"Ошибка GET-запроса: {e}")
        return {"error": str(e)}

def load_config(config_path: str) -> Dict[str, Any]:
    path = Path(config_path)

    if not path.exists():
        raise FileNotFoundError(f"Конфиг не найден: {config_path}")
    
    with open(path, 'r', encoding='utf-8') as f:
        config = yaml.safe_load(f)
    
    return config


if __name__ == "__main__":
    try:
        
        
        # Получаем настройки сервера
        server_host = config["server"]["host"]
        server_port = config["server"]["port"]
        
        print(f"Сервер: {server_host}:{server_port}")
        print("Полный конфиг:", config)
        
    except FileNotFoundError as e:
        print(f"Ошибка: {e}")
    except yaml.YAMLError as e:
        print(f"Ошибка YAML: {e}")
    except KeyError as e:
        print(f"Ошибка: Нет ключа в конфиге - {e}")

def funcSendPOSTrequest(url):
    pass

def funcSendGETrequest(url):
    pass

if __name__ == "__main__":
    int_input = input("GET: 1 | POST: 2")
    adr_iport = input("Enter address:port")

    match int_input:
        case 1:
            print("Case: GET")
            respPost = funcSendPOSTrequest()
        case 2:
            print("Case: POST")

    print("Send POST...")
    post_response = send_post_request(API_URL, EXAMPLE_DATA)
    print("Ответ сервера (POST):")
    print(json.dumps(post_response, indent=2))  # Красивый вывод JSON

    print("\nОтправка GET-запроса...")
    get_response = send_get_request(API_URL)
    print("Ответ сервера (GET):")
    print(json.dumps(get_response, indent=2))


