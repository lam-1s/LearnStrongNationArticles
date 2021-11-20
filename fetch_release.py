import json
import os

import requests

if __name__ == '__main__':
    resp = requests.get('https://api.github.com/repos/lam-1s/LearnStrongNationArticles/releases')
    releases = json.loads(resp.text)
    dlUrl = ""
    for i in releases[0]["assets"]:
        if i["browser_download_url"].find(".tar.xz") != -1:
            dlUrl = i["browser_download_url"]

    print("Fetched latest release:")
    print(dlUrl)

    print("Fetching asset...")
    xzResp = requests.get(dlUrl)
    print("Fetched, status is " + str(xzResp.status_code))
    print("File length " + str(len(xzResp.content)))
    xzOut = open("release.tar.xz", "wb")
    xzOut.write(xzResp.content)
    print("File written.")
