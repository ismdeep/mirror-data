import json
import os
from common import Storage

storage = Storage('alpine-linux')


def main():
    remote_site = 'https://dl-cdn.alpinelinux.org/alpine'
    versions = [
        'v3.0', 'v3.1', 'v3.2', 'v3.3', 'v3.4', 'v3.5', 'v3.6', 'v3.7', 'v3.8', 'v3.9', 'v3.10',
        'v3.11', 'v3.12', 'v3.13', 'v3.14', 'v3.15', 'v3.16', 'v3.17', 'v3.18'
    ]
    for version in versions:
        items = json.loads(
            os.popen('rclone lsjson -R --include=\'/**.iso\' --http-url {}/{}/releases/ :http:'.format(remote_site, version)).read())
        for v in items:
            if v['IsDir']:
                continue
            storage.write('{}/{}'.format(version, v['Path']),
                          '{}/{}/releases/{}'.format(remote_site, version, v['Path']))


if __name__ == '__main__':
    main()
