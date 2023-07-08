import requests
from lxml import html
from common import Storage
from common import get_go_version

storage = Storage('go')


def main():
    items = html.fromstring(requests.get(url='https://go.dev/dl/').content).xpath('''//a[@class="download"]/@href''')
    cnt = len(items)
    index = 0
    for item in items:
        item = item.strip()
        index += 1
        if '/dl/' not in item:
            continue
        file_name = item[4:]
        version = get_go_version(file_name)
        print('{} / {}    {}'.format(index, cnt, file_name))
        path_all = 'all/{}'.format(file_name)
        path_dist = 'dist/{}/{}'.format(version, file_name)
        dl_url = 'https://go.dev/dl/{}'.format(file_name)
        storage.write(__path__=path_all, __origin_url__=dl_url)
        storage.write(__path__=path_dist, __origin_url__=dl_url)


if __name__ == '__main__':
    main()
