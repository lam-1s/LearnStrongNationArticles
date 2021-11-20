# LearnStrongNationArticles

Articles from xuexi.cn / 学习强国文章库

## What is this?

This is a project which automatically fetch newly-published articles on [学习强国](https://xuexi.cn/).

## How I can use them?

Head to the release page and choose what you want to download.

When the archive file was decompressed, you will see files were structured as below:

```
- lastUpdated.json
- txt
  |- 学习理论
     |- “红船精神”的时代价值.txt
     |- ...
  |- 人民心中的习近平
  |- ...
```

Articles on [学习强国](https://xuexi.cn/) were categorised into different "channels", where "channels" could have some sub-"channels". We preserved such a tree-like structure since the name of the channel is strongly related to the content of articles. 

### Integrity check

You may want to check the integrity of the downloaded archive to ensure that nobody other than us modified the file. Luckily, our files are GPG signed. You can do the verification yourself. Our GPG key id is 356434255004FABE, you can retrieve it on [OpenPGP Key Server](https://keys.openpgp.org/).

```
$ gpg --verify XXX.sig
gpg: assuming signed data in 'XXX'
gpg: Signature made Sun 21 Nov 2021 00:00:00 UTC
gpg:                using ECDSA key DD98ED60A11780F091010407356434255004FABE
gpg: Good signature from "LearnStrongNation Bot (Release Signing Only) <qamp99asj@mozmail.com>" [full]
```

## How are these articles updated?

The articles are updated at 01:15 and 13:15 (UTC) everyday.

## I have questions / suggestions...

Please open an issue to state that. Thank you in advance.

_Last updated: 21 November 2021_
