
# gomail

One day my friend asked me if it's possible to automate sending e-mails, so I
decided it's a great opportunity to practice Go and here it is - a tiny program
that can send e-mails via SMTP.

**It was created for educational purposes only, don't use it to spread spam.**

**You need to have an account on a mail server (e. g. Gmail, Outlook or own
server) to use this program!**

## How to install

### Stable release

Download the executable for your operating system from
https://github.com/alexcoder04/gomail/releases/latest.

### Compiling yourself

```sh
git clone https://github.com/alexcoder04/gomail
cd gomail
make linux # or `make windows`
```

## How to use (important!)

The program is configured with the following files:

### `<USER_CONFIG_DIR>/gomail/settings.yml`

You can use another file by passing the `-s` flag (e. g. `gomail -s mysettings.yml`)

The host address can vary depending on your e-mail provider, most likely it's
something like *mail.example.com* or *smtp.example.com*. The port is also
different for different e-mail providers.

For the username, it is the part of the mail address before the @-sign for my
provider, but it may be the full address (e. g. on Gmail) or something
completely different for your e-mail provider.

```yml
From: youraddress@example.com
Addr: smtp.example.com:587
Host: smtp.example.com
Username: youraddress@example.com
Password: YourSecureP4ssw0rd
Subject: Hello friends!
```

### `<USER_CONFIG_DIR>/gomail/recipients.txt`

You can use another file by passing the `-r` flag (e. g. `gomail -r friendslist.txt`)

```text
myfriend1@example.com
myfriend2@example.com
```

### `./mail.txt`

You can use another file by passing the `-b` flag (e. g. `gomail -b hello.txt`)

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

