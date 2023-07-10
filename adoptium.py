import json
import os
from common import Storage

storage = Storage('adoptium')


def main():
    remote_site = 'https://mirrors.tuna.tsinghua.edu.cn/Adoptium/'
    items = json.loads(
        os.popen('rclone lsjson -R --http-url {} :http:'.format(remote_site)).read())
    for v in items:
        if v['IsDir']:
            continue
        if '.tar.gz' in v['Path'] or '.zip' in v['Path']:
            storage.write(v['Path'], '{}{}'.format(remote_site, v['Path']))


if __name__ == '__main__':
    main()
