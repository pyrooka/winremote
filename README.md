# Winremote

## Usage
- `.\winremote.exe install`: install the service
- `.\winremote.exe uninstall`: uninstall the service

## Notes
- If you have firewall, you have to create a rule to allow incoming TCP connections on the specific port

## TODO
- token auth
- install/uninstall copy the executable to e.g. sys32 and run from there
- webserver config

## Credits
- service: https://github.com/golang/sys/tree/master/windows/svc/example
- message box: https://gist.github.com/NaniteFactory/0bd94e84bbe939cda7201374a0c261fd