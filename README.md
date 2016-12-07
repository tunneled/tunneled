# Tunneled

Tunneled is an application to assist in viewing webhook deliveries. It provides
the pieces of an HTTP request in your terminal, as then come in, broken up by
colors.

It's meant to be used in conjunction with a reverse SSH tunnel like:

```
$ ssh -nNT -R 80:localhost:8000 root@bswinnerton.tunneled.computer
```

It looks something like this:

![screenshot](https://cloud.githubusercontent.com/assets/934497/20954501/eb046a5c-bc09-11e6-9be0-5798e226028d.png)
