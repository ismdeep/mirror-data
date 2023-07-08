import json
import random
import unittest

import requests
from dateutil import parser


class Storage:
    bucket_name: str
    fp = None
    exists = set()

    def __init__(self, __bucket_name__):
        self.bucket_name = __bucket_name__
        file_path = './data/{}.txt'.format(__bucket_name__)

        with open(file_path, 'a') as f:
            print(f.name)

        with open(file_path, 'r') as f:
            for line in f.readlines():
                line = line.strip()
                v = line.split('|')
                self.exists.add(v[1])

        self.fp = open(file_path, 'a')

    def write(self, __path__, __origin_url__):
        print('{} -> {}'.format(__path__, __origin_url__))
        if __path__ in self.exists:
            return
        resp_header = requests.head(url=__origin_url__, allow_redirects=True)
        last_modified = int(parser.parse(resp_header.headers['Last-Modified']).timestamp())
        content_length = resp_header.headers['Content-Length']
        content_type = resp_header.headers['Content-Type']
        self.fp.write('{}|{}|{}|{}|{}|{}\n'.format(
            self.bucket_name,
            __path__,
            __origin_url__,
            content_length,
            content_type,
            last_modified
        ))
        self.fp.flush()
        self.exists.add(__path__)


def get_go_version(__s__: str):
    __s__ = __s__.strip()
    v = __s__.split('.')
    items = [v[0]]
    for i in range(1, len(v)):
        if '0' <= v[i][0] <= '9':
            items.append(v[i])
            continue
        break
    return '.'.join(items)


class TestGetGoVersion(unittest.TestCase):

    def test_get_go_version(self):
        self.assertEqual(get_go_version('go1.20.5.src.tar.gz.tar.gz'), 'go1.20.5')
        self.assertEqual(get_go_version('go1.8rc1.freebsd-386.tar.gz'), 'go1.8rc1')
        self.assertEqual(get_go_version('go1.21rc2.freebsd-amd64.tar.gz'), 'go1.21rc2')


def get_github_release(__bucket_name__: str, __owner__: str, __repo__: str):
    storage = Storage(__bucket_name__)
    config = json.load(open('config.json'))
    page = 0
    per_page = 100
    while True:
        page += 1
        req = requests.get(
            url='https://api.github.com/repos/{}/{}/releases?page={}&per_page={}'.format(
                __owner__, __repo__, page, per_page),
            headers={
                'Accept': 'application/vnd.github+json',
                'Authorization': 'Bearer {}'.format(random.choice(config['ghp_list'])),
                'X-GitHub-Api-Version': '2022-11-28'
            }
        )
        items = json.loads(req.text)
        for i in range(len(items)):
            tag_item = items[i]
            print('{} / {}'.format(i + 1, len(items)))
            version = tag_item['tag_name']
            for asset in tag_item['assets']:
                file_name = asset['name']
                download_url = asset['browser_download_url']
                path_dist = '{}/{}'.format(version, file_name)
                storage.write(__path__=path_dist, __origin_url__=download_url)
        if len(items) < per_page:
            break
