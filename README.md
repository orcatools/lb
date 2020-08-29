# lb | Lockbox CLI

This is the official CLI tool for Lockbox.

## How to Use

- Step 1. Create a new Lockbox
    - `lb init <lockbox-name>`
    - this will create the lockbox,
      and return a OTP Secret.
    - Add the OTP secret to a MFA device
      manager like Authy, or Google Authenticator.
    - After the secret has been added, you'll
      have access to the MFA codes required to use lb.

- Step 2. Add a Secret to the Lockbox
    - the lockbox file should be in the pwd.
    - `lb set <lockbox-name> --path /some/path --value "Some Secret" --code <MFA>`
    - Optionally, you can also specify a namespace with `--namespace`.
      If no namespace is provided, the default `main` namespace is used.
    
    - NOTE: If you see this error message:
        `cannot set value while lockbox is locked`
      it means the mfa code did not work for some reason (probably expired)

- Step 3. Get a Secret from the Lockbox
    - the lockbox file should be in the pwd.
    - `lb get <lockbox-name> --path /some/path --code <MFA>`
    - Note: if you specified a namespace, other than the default `main`,
      you'll also have to provide that namespace in order to access the secret.
      Do that with the `--namespace <namespace>` flag.

    - NOTE: If you see this error message:
        `cannot get value while lockbox is locked`
      it means the mfa code did not work for some reason (probably expired)
