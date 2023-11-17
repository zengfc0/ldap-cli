# Ldap-Cli
The openldap command-line tools.

## Config File

```shell
# ~/.ldap.yaml
basedn:
binddn:
bindpasswd:
ldapurl:
```

## Usage

```shell
# get user phoneNumber
ldap-cli get phoneNumber -U <uid>

# update user phoneNumber
ldap-cli update phoneNumber -U <uid> -M <phoneNumber>

# Run 'ldap-cli <command> --help' for more information on a command.
ldap-cli --help
```
