# Winremote

Remotely sleep/hibernate/restart/shutdown your Windows PC via HTTP GET request.

I have made this little tool to let my [Home Assistant](https://www.home-assistant.io/) hibernate my computer.

## Usage
- `.\winremote.exe install`: install the service
- `.\winremote.exe uninstall`: uninstall the service

## Note
If you have firewall enabled, you have to create a rule to allow incoming TCP connections on the specific port.

## TODO
- token auth
- install/uninstall copy the executable to e.g. sys32 and run from there
- webserver config

## Home Assistant config example
```yaml
switch:
  - platform: wake_on_lan
    name: "TARGET"
    ...
    turn_off:
      service: shell_command.turn_off_TARGET

shell_command:
  turn_off_TARGET: 'ssh hass@TARGET sudo pm-suspend'
```

## Credits
- service: https://github.com/golang/sys/tree/master/windows/svc/example