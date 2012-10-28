noSSH
=====

noSSH is a zero-interaction SSH-Honeypot.

Description
-----------
noSSH is a small SSH Server that listens on port 2222 for new connections. As soon as a username/password pair is received the login attempt gets logged and will be denied.

The purpose of this project is to create a list of common username/password combinations automated bots and attackers use on SSH services.

Warning
-------
The program wasn't tested so be careful!!!

Binaries
--------
Under `Downloads` there are binaries for

- Linux
  - i386 (packed with UPX)
  - amd64
  - arm (packed with UPX)
- Windows
  - i386 (packed with UPX)
  - amd64

Run
---
- Generate a SSH keypair with `utils/generate_keys.sh`
- Move the keys to the noSSH root directory
- Run `noSSH`

