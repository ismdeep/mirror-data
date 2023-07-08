import json
import random

import requests
from common import Storage

storage = Storage('ctop')
config = json.load(open('config.json'))


def get_release_list():
    page = 0
    per_page = 100
    while True:
        page += 1
        req = requests.get(
            url='https://api.github.com/repos/bcicen/ctop/releases?page={}&per_page={}'.format(page, per_page),
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


def main():
    get_release_list()


if __name__ == '__main__':
    main()
