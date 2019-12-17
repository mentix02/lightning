import os
import unittest

import requests

base_url = 'http://localhost:8080'


class BookmarkViewTest(unittest.TestCase):

    @classmethod
    def setUpClass(cls) -> None:

        if os.path.exists('token.txt'):
            with open('token.txt') as f:
                cls.token = f.read()
        else:
            with requests.post('http://127.0.0.1:8000/api/authors/authenticate/', data={
                'username': 'mentix02',
                'password': 'abcd1432',
            }) as r:
                f = open('token.txt', 'w+')
                cls.token = r.json()['token']
                f.write(cls.token)

        cls.url = f'{base_url}/bookmark/list/'

    def test_unauthenticated_bookmark_list_handler(self):

        r = requests.get(self.url)

        self.assertEqual(r.status_code, 401)
        self.assertEqual(r.json(), {'detail': 'Authentication credentials were not provided.'})

    def test_authenticated_bookmark_list_hander(self):

        r = requests.get(self.url, headers={
            'Authorization': f'Token {self.token}'
        })

        self.assertEqual(r.status_code, 200)
        self.assertTrue(isinstance(r.json(), list))

    def test_invalid_credentials_bookmark_list_handler(self):

        r = requests.get(self.url, headers={
            'Authorization': f'Token aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa'
        })

        self.assertEqual(r.status_code, 401)
        self.assertEqual(r.json(), {'detail': 'Invalid credentials.'})

    def test_ill_formed_credentials_bookmark_list_handler(self):

        r = requests.get(self.url, headers={
            'Authorization': 'Token'
        })

        self.assertEqual(r.status_code, 401)
        self.assertEqual(r.json(), {'detail': 'Authentication credentials were not provided.'})


if __name__ == '__main__':
    unittest.main()
