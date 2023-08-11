import json
import requests
from common import Storage

storage = Storage('jetbrains')

user_agent = 'Mozilla/5.0 (Macintosh; ' \
             'Intel Mac OS X 10_15_7) ' \
             'AppleWebKit/537.36 (KHTML, like Gecko) ' \
             'Chrome/115.0.0.0 Safari/537.36'


def fetch(__code__: str, __product_name__: str):
    req = requests.get(
        url="https://data.services.jetbrains.com/products?code={}&release.type={}&fields={}".format(
            __code__,
            'eap%2Crc%2Crelease',
            'distributions%2Clink%2Cname%2Creleases'
        ),
        headers={
            'Usage-Agent': user_agent
        }
    )
    obj = json.loads(req.text)
    for o in obj:
        # name = o['name']
        for release in o['releases']:
            type_name = release['type']
            if type_name != 'release':
                continue
            major_version = release['majorVersion']
            version = release['version']
            for arch in release['downloads']:
                download = release['downloads'][arch]
                origin_link = download['link']
                s = origin_link.split('/')
                file_name = s[len(s) - 1]
                if len(file_name) > 5 and file_name[len(file_name) - 5:] == ".json":
                    continue
                link = '{}/{}/{}/{}/{}/{}'.format(__product_name__, type_name, major_version, version, arch, file_name)
                storage.write(link, origin_link)


def main():
    fetch('CL', 'CLion')  # CLion
    fetch('DG', 'DataGrip')  # DataGrip
    fetch('DS', 'DataSpell')  # DataSpell
    fetch('GO', 'GoLand')  # GoLand
    fetch('PC', 'PyCharm')  # PyCharm
    fetch('PS', 'PhpStorm')  # PhpStorm
    fetch('WS', 'WebStorm')  # WebStorm


if __name__ == '__main__':
    main()
