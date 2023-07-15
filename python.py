import json
import os
from common import Storage

storage = Storage('python')


def is_compress_file(__path__):
    return '.tar.xz' in __path__


def exec_cmd_json(__cmd__):
    return json.loads(os.popen(__cmd__).read())


def main():
    remote_site = 'https://www.python.org/ftp/python/'
    versions = exec_cmd_json('rclone lsjson --http-url {} :http:'.format(remote_site))
    idx = 0
    for version in versions:
        idx += 1
        if not version['IsDir'] or not ('0' <= version['Path'][0] <= '9'):
            continue
        print('{} / {}    {}'.format(idx, len(versions), version['Path']))
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