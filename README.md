# Winremote

Winremote is a simple windows service to remotely sleep/hibernate/restart/shutdown your Windows PC via HTTP GET request.

I have made this to make my [Home Assistant](https://www.home-assistant.io/) able to hibernate my computer.

## Usage
1. `.\winremote.exe install`: install the service
2. Create a new firefall rule if enabled: `New-NetFirewallRule -DisplayName "Winremote" -Direction Inbound -Protocol TCP -LocalPort 8899 -RemoteAddress LocalSubnet -Action Allow`
3. Start the service manually (or it will start on the next boot)

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