# lb | Lockbox CLI

This is the official CLI tool for Lockbox.

## How to Use

- Step 1. Setup your config
  - by default the config file is stored at `$HOME/.lockbox.yaml`
  - you can point to a different config file with the `--config` flag.
  - Please see below for a sample config file
  - NOTE: you can override values that are in the config file in 2 different ways:
    - Environment Variables:
      - `LOCKBOX_USERNAME`
      - `LOCKBOX_PASSWORD`
      - `LOCKBOX_NAMESPACE`

    - CLI Flags
      - `--username`
      - `--password`
      - `--namespace`

- Step 2. Create a new Lockbox
    - `lb init <lockbox>`
    - this will create the lockbox,
      and return a OTP Secret.
    - Add the OTP secret to a MFA device
      manager like Authy, or Google Authenticator.
    - After the secret has been added, you'll
      have access to the MFA codes required to use lb.

- Step 3. Add a Secret to the Lockbox
    - the lockbox file should be in the pwd.
    - `lb set <lockbox-name> --path /some/path --value "Some Secret" --code <MFA>`
    - NOTE: If you see this error message:
        `cannot set value while lockbox is locked`
      it means the mfa code did not work, or that your user information was invalid for some reason.

- Step 3. Get a Secret from the Lockbox
    - the lockbox file should be in the pwd.
    - `lb get <lockbox-name> --path /some/path --code <MFA>`
    - Note: if you specified a namespace, other than the default `main`,
      you'll also have to provide that namespace in order to access the secret.
      You can do that in the config file, with an environement variable `LOCKBOX_NAMESPACE` or with the `--namespace <namespace>` flag.

    - NOTE: If you see this error message:
        `cannot get value while lockbox is locked`
      it means the mfa code did not work for some reason (probably expired), or your user information is inccorrect.

## Config File Example:

Default location is: `$HOME/.lockbox.yaml`
```
username: <USERNAME>
password: <PASSWORD>

namespace: <NAMESPACE>
```