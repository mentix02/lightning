import unittest

import requests

base_url = 'http://localhost:8080'


class AuthenticationTest(unittest.TestCase):

    def test_valid_creds_authentication(self):
        r = requests.post(f'{base_url}/authors/authenticate/', data={
            'username': 'mentix02',
            'password': 'abcd1432'
        })
        self.assertEqual(r.status_code, 200)
        self.assertRegex(r.json()['token'], r'^[A-Fa-f0-9]{40}$')

    def test_invalid_username_authentication(self):
        r = requests.post(f'{base_url}/authors/authenticate/', data={
            'username': 'abcd',
            'password': 'abcd1432'
        })
        self.assertEqual(r.status_code, 401)
        self.assertEqual(r.json(), {
            'detail': 'Invalid credentials.'
        })

    def test_invalid_password_authentication(self):
        r = requests.post(f'{base_url}/authors/authenticate/', data={
            'username': 'mentix02',
            'password': 'aaaa'
        })
        self.assertEqual(r.status_code, 401)
        self.assertEqual(r.json(), {
            'detail': 'Invalid credentials.'
        })


if __name__ == '__main__':
    unittest.main()
