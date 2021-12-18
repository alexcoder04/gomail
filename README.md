
# gomail

One day my friend asked me if it's possible to automate sending e-mails, so I
decided it's a great opportunity to practise Go and here it is - a tiny program
that can send e-mails via SMTP.

**It was created for educational purposes only, don't use it to spread spam.**

## How to install

### Stable release

Download the executable for your operating system from
https://github.com/alexcoder04/gomail/releases/latest.

### Compiling yourself

```sh
git clone https://github.com/alexcoder04/gomail
cd gomail
make linux # or make windows
```

## How to use (important!)

This program is pretty much in alpha status, currently it's configured through
configuration files, which must be located in the same directory as the
executable is run in. Those are:

### `settings.yml`

The host address can vary depending on your e-mail provider, most likely it's
something like *mail.domain.tld* or *smtp.domain.tld*. The port also is
different for different e-mail providers.

```yml
From: youraddress@example.com
Addr: smtp.example.com:587
Host: smtp.example.com
Subject: Hello friends!
```

### `account.txt`

The first line is your username. For my provider, it's the part of the address
before the @-sign, but it may be the full address or something completely
different for your e-mail provider.

```yml
Usename: youraddress@example.com
Password: YourSecureP4ssw0rd
```

### `recipients.txt`

```text
myfriend1@example.com
myfriend2@example.com
```

### `mail.txt`

```text
Hello friends,

I sent this mail to you using gomail.

Have a nice day :)
```

## FAQ

### Can I trust you that you aren't going to steal my password?

 - Although I never would even try to do something like that, you can't. But
   you can read through the code and compile it yourself, so you don't have to
   trust anyone :)

