# Chatter

Simple Chatter forum website written in Golang under instructions of [Go Web Programming](https://www.goodreads.com/book/show/27797995-go-web-programming) and inspired by [thesrc](https://github.com/sourcegraph/thesrc).

## Description

![Threads image](https://github.com/williamzion/chatter/blob/master/assets/threads.png)

![Posts image](https://github.com/williamzion/chatter/blob/master/assets/posts.png)

Chatter forum is developed with pure Golang (which is not by any web frameworks or third party routers) and PostgreSQL as datestore.

Users in Chatter can register forum, create threads conversation and post gossips under every thread.

The tour of developing Chatter web application showed how powerful and simple Golang is as compared with other languages. Mostly because of the root of Golang language design and secondly its good and complete standard library packages.

## Installation And Usage

Clone this git repository under your `GOPATH/src` directory:

```bash
git clone git@github.com:williamzion/chatter.git
```

Create a `.env` file under chatter root directory:

```bash
>$ cd chatter
>$ touch .env
```

Type your Postgres credentials as following:

```txt
user=
password=
dbname=
sslmode=disable
```

Then you will need to setup your database tables for this application with file in `chatter/datastore/setup.sql`.

Finally, build and run:

```bash
>$ go build
>$ ./chatter
```

## Credits

- [William](https://github.com/williamzion)

## License

Under MIT license.
