import os

import requests


def get_key() -> str:
    if os.path.exists('token.txt'):
        with open('token.txt') as f:
            return f.read()
    else:
        with requests.post('http://localhost:8000/api/authors/authenticate/', data={
            'username': 'mentix02',
            'password': 'abcd1432',
        }) as r:
            f = open('token.txt', 'w+')
            return r.json()['token']
