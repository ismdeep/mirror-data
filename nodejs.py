import json
import os
from common import Storage

storage = Storage('nodejs')


def is_compress_file(__path__):
    return '.tar.gz' in __path__ \
        or '.tar.xz' in __path__ \
        or '.zip' in __path__ \
        or '.7z' in __path__


def is_ignored_path(__path__):
    return 'latest' in __path__ \
        or 'npm' in __path__ \
        or 'patch' in __path__ \
        or '-isaacs-manual' in __path__


def exec_cmd_json(__cmd__):
    return json.loads(os.popen(__cmd__).read())


def main():
    remote_site = 'https://nodejs.org/dist/'
    versions = exec_cmd_json('rclone lsjson --http-url {} :http:'.format(remote_site))
    idx = 0
    for version in versions:
        idx += 1
        print('{} / {}    {}'.format(idx, len(versions), version['Path']))
        if not version['IsDir'] or is_ignored_path(version['Path']):
            continue
        try:
            items = exec_cmd_json('rclone lsjson --http-url {}{}/ :http:'.format(remote_site, version['Path']))
            for v in items:
                if not is_compress_file(v['Path']):
                    continue
                path = '{}/{}'.format(version['Path'], v['Path'])
                origin_url = '{}{}/{}'.format(remote_site, version['Path'], v['Path'])
                storage.write(path, origin_url)
        except Exception as e:
            print('Err: {}'.format(e))


if __name__ == '__main__':
    main()
