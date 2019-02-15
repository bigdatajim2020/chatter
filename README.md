# Chatter

Simple Chatter forum website written in Golang under instructions of [Go Web Programming](https://www.goodreads.com/book/show/27797995-go-web-programming) and inspired by [thesrc](https://github.com/sourcegraph/thesrc).

## Description

Chatter forum is developed with pure Golang (which is not by any web frameworks or third party routers) and PostgreSQL as datestore.

Users in Chatter can register forum, create threads conversation and post gossips under every thread.

![Threads image](https://github.com/williamzion/chatter/blob/master/assets/threads.png)

![Posts image](https://github.com/williamzion/chatter/blob/master/assets/posts.png)

The tour of developing Chatter web application showed how powerful and simple Golang is as compared with other languages. Mostly because of the root of Golang language design and secondly its good and complete standard library packages.

## Installation

To install Chatter, run the following:

```bash
go get -u github.com/williamzion/chatter
```

## Usage

You may better use chatter root as working directory.

Create a `config.yaml` file under chatter root directory which is the current working directory:

```bash
>$ cd ~/go/src/github.com/williamzion/chatter
>$ touch config.yaml
```

Type your Postgres credentials as following format:

```txt
user: _db-user_
password: _db-password_
dbname: _db-name_
sslmode: disable
```

Then you will need to bootstrap your database tables for this application with file in `chatter/datastore/setup.sql`.

Finally, build and run:

```bash
>$ go build
>$ ./chatter
```

## Credits

- [William](https://github.com/williamzion)

## License

Under MIT license.
