# Recache
Restart fusion cache

Alright this is a tiny tool i built out of necessity in order to help me speed things up in my job, what it does 
comes down to let me start or stop a docker container in case it's running or stopped with just a couple clicks. 
As i work with a framework built on top of react and it kind of has a docker container inside which among other 
things cache its api content i need a way to somewhat shut down this functionality during development so  as i'm 
a golang beginner i thought it would be a good challenge to see if i could do it.

<table border="0" cellspacing="0" cellpadding="0" style="border-collapse: collapse; border: none;">
  <tr>
    <td><img alt="GitHub" src="https://img.shields.io/github/license/wwleak/recache?style=for-the-badge"></td>
    <td><img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/wwleak/recache?style=for-the-badge"></td>
    <td><img alt="GitHub top language" src="https://img.shields.io/github/languages/top/wwleak/recache?style=for-the-badge"></td>
    <td><img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/wwleak/recache?style=for-the-badge"></td>
    <td><img alt="GitHub stars" src="https://img.shields.io/github/stars/wwleak/recache?style=for-the-badge"></td>
  </tr>
</table>

## How to build it :rocket: 

Clone the repo

```console
foo@bar:~$ git clone https://github.com/wwleak/recache.git
```

Change directory into the src folder

```console
foo@bar:~$ cd recache/src
```

Run the build command as follows

```console
foo@bar:~$ go build -o ../bin/recache
```
