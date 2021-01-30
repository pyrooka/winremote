# Winremote

Winremote is a simple windows service to remotely sleep/hibernate/restart/shutdown your Windows PC via HTTP GET request.

I have made this to make my [Home Assistant](https://www.home-assistant.io/) able to hibernate my computer.

## Usage
1. Copy the executable to a safe place. E.g.: `C:\Program Files\Winremote\winremote.exe`. The service will search for the exe in that path every time.
2. `.\winremote.exe install`: install the service   
3. Create a new firefall rule if enabled: `New-NetFirewallRule -DisplayName "Winremote" -Direction Inbound -Protocol TCP -LocalPort 8899 -RemoteAddress LocalSubnet -Action Allow`
4. Start the service manually (or it will start on the next boot)

## TODO
- token auth
- webserver config

## Home Assistant config example
https://www.home-assistant.io/integrations/wake_on_lan/#examples
```yaml
switch:
  - platform: wake_on_lan
    name: PC
    ...
    turn_off:
      service: shell_command.hibernate_PC

shell_command:
  hibernate_PC: 'curl http://192.168.0.2:8899/sleep'
```

## Credits
- service: https://github.com/golang/sys/tree/master/windows/svc/example