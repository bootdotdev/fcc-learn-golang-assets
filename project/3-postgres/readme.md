# PostgreSQL

PostgreSQL is a production-ready, open-source database. It's a great choice database for many web applications, and as a back-end engineer, it might be the single most important database to be familiar with.

## How does PostgreSQL work?

Postgres, like most other database technologies, is itself a server. It listens for requests on a port (Postgres' default is `:5432`), and responds to those requests. To interact with Postgres, first you will install the server and start it. Then, you can connect to it using a client like [psql](https://www.postgresql.org/docs/current/app-psql.html#:~:text=psql%20is%20a%20terminal%2Dbased,or%20from%20command%20line%20arguments.) or [PGAdmin](https://www.pgadmin.org/).

## 1. Install

### Mac OS

I recommend using [brew](https://brew.sh/) to install PostgreSQL on Mac OS.

```bash
brew install postgresql
```

### Linux (or WSL)

I recommend using apt to install PostgreSQL on Linux (Ubuntu). Here are the [docs from Microsoft](https://learn.microsoft.com/en-us/windows/wsl/tutorials/wsl-database#install-postgresql). The basic steps are:

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```

## 2. Ensure the installation worked

The `psql` command-line utility is the default client for Postgres. Use it to make sure you're on version 14+ of Postgres:

```bash
psql --version
```

## 3. Start the Postgres server in the background

### Mac OS

```bash
brew services start postgresql
```

### Linux (or WSL)

```bash
sudo service postgresql start
```

## 4. Connect to the server using a client

I'm going to recommend using the PGAdmin client. It's a GUI that makes it easy to interact with Postgres and provides a lot of useful features. If you want to use `psql` on the command line instead, you're welcome to do so.

1. Download [PGAdmin here](https://www.pgadmin.org/).
2. Open PGAdmin and create a new server connection. Here are the connection details you'll need:

### Mac OS (with brew)

* Host: `localhost`
* Port: `5432`
* Username: Your Mac OS username
* Password: *leave this blank*

### Linux (or WSL)

If you're on Linux, you have one more step before you can connect with credentials. On your command line run:

```
sudo passwd postgres
```

Enter a new password *and remember it*. Then restart your shell session.

* Host: `localhost`
* Port: `5432`
* Username: `postgres`
* Password: *the password you created*

## 5. Create a database

A single Postgres server can host multiple databases. In the dropdown menu on the left, open the `Localhost` tab, then right click on "databases" and select "create database".

Name it whatever you like, but you'll need to know the name.

## 6. Query the database

Right click on your database's name in the menu on the left, then select "query tool". You should see a new window open with a text editor. In the text editor, type the following query:

```sql
SELECT version();
```

And click the triangle icon (execute/refresh) to run it. If you see a version number, you're good to go!

*PGAdmin is the "thunder client" of Postgres. It's just a GUI that allows you to run ad-hoc queries against your database.*
