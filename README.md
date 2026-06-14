CLI tool for managing passwords

# Install

`go get github.com/blihor/parrot`

# Setting password

Parrot initially comes without any password, but all commands except `parrot set`
are disabled. To set new password run:

`parot set <your password>`

After password is set you may provide `--master`(short `-m`) flag for every command.
If you don't, you will be asked to type it in after you ran command.

# Adding new entry

Parrot operates on entries(name, url, username, email, password). To add a new
entry to the vault run:

`parrot add <name> <password> [options]`

> [!NOTE] Only name argument is required. If password wasn't provided it will be generated automatically based on your config file

If you need to associate url, username or email with a new entry, specify them
with `--url`, `--username`, `--password` flags:

`parrot add typeracer --url typeracer.com --username john --email johndoe@gmail.com`

> [!NOTE] You can use any set of additional parameters for an entry. All the 3 are optional

# Editing entry

You can edit any field of an entry by running:

`parrot edit <name> [options]`

Where options are: `--name`(short `-n`), `--username`(short `-u`), `--url`(short `-l`),
`--email`(short `-e`), `--password`(short `-p`)

Example:

`parrot edit oldName --name newName --url newUrl --username newUsername --email newEmail --password NewPassword`

> [!NOTE] All flags are optional, but you should provide at least one

You may opt for not providing new password. In that case you should set `--gen`
flag, which will enable password generator.

# Deleting entry

To delete entry from vault run:

`parrot delete <name>`

# Listing entries

Parrot can list all of your entries or just one by specified name.

`parrot list <name>`

If name argument is omitted parrot will list all entries from the vault.

# Generator

Parrot has standalone command for generating password:

`parrot gen`

It will generate password based on configuration found in your config file. You
can override all the configurations using flags: 

`--length`(short `-l`) for length of the password
`--upper`(short `-u`) for including upper case letters
`--digits`(short `-d`) for including digits
`--special`(short `-s`) for including special characters

This example command will generate password 12 characters long will lower and 
upper case letters, digits and special characters:

`parrot gen -udsl 12`

If config file wasn't found, and you didn't explicitly override generator configuration,
it will use defaults: `-usdl 16`

# Config

Parrot will look for config named `parrot` in following locations:

1. `$XDG_CONFIG_HOME/parrot`
2. `$HOME/.parrot`
3. `/etc/parrot/`

Parrot supports following extensions:

- JSON
- TOML
- YAML
- INI

Config example:

```yaml
---
generator:
  upper: false
  digits: false
  special: true
---
aes:
  time: 1
  mem: 64000
  threads: 4
  keylen: 32
  saltlen: 32
```
