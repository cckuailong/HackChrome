# HackChrome

[![Build Status](https://travis-ci.com/cckuailong/HackChrome.svg?branch=master)](https://travis-ci.com/cckuailong/HackChrome)

[English ReadMe](https://github.com/cckuailong/HackChrome/blob/master/README.md) || 
[中文 ReadMe](https://github.com/cckuailong/HackChrome/blob/master/README_zh.md)

Get the User:Password from Chrome(include version < 80 and version > 80)

## Chrome version Affact

All version

## Platform

Windows

## Usage

- Download the exe file [here](https://github.com/cckuailong/HackChrome/releases/tag/v0.1)

- Open cmd or powershell

- Run

```
Hackone.exe > res.txt
```

## Demo

![demo](image/result.png)

## Theory

- version < 80

User:Password pairs were stored in the file named "Login Data".

Password was encrypted, But we can use "CryptUnprotectData" Function in "Crypt32.dll" to decrypt them.

Finally, We get the plaintext of the User:Password pairs stored in Chrome

- version > 80

Based on the Algorithm used by "version < 80", It use AES-GCM to encrypt the password via a <master key> and a <nounce>.

The <master key> can be found in the "Local State" file, and can be decypted by "CryptUnprotectData" mentioned above.

The <nounce> can be found at the begin of the encrypted_password.

Therefore, we can decrpted all the password.

- Merge the result

If someone update the Chrome recently, we need to find the two ways of User:Password pairs.

What's more, I use some rules to merge the results into an array.

## LICENSE

The Project follows MIT LICENSE.
