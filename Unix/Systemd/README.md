### Commands

    systemctl               # Lists all units
    systemctl --failed          # List failed
    systemctl status
    systemctl status <unit>

    systemctl enable <unit>
    systemctl disable <unit>
    
    systemctl enable --now <unit>  # Enable and start in one go
    systemctl disable --now <unit> # Disable and stop in one go

    systemctl start <unit>
    systemctl stop <unit>

    systemctl mask <unit>
    systemctl unmask <unit>

    systemctl restart <unit>
    systemctl reload <unit>

    systemctl reset-failed [<unit>]

## How do you change unit definitions?

<details><summary>show</summary>
<p>

    systemctl edit <unit>
    systemctl daemon-reload   # after directly modifying/adding config files

Global definitions from /lib/systemd/system will then be overruled by a
new file in /etc/systemd/system

</p>
</details>

## How do you delete units?

<details><summary>show</summary>
<p>


    systemctl stop <service>
    systemctl disable <service>
    rm /etc/systemd/system/<servicename>
    systemctl daemon-reload
    systemctl reset-failed

</p>
</details>

## How do you list processes / containers?

<details><summary>show</summary>
<p>


To help identify cgroup/process relations run

    ps xawf -eo pid,user,cgroup,args

or

    systemd-cgls

</p>
</details>

### Performance

<details><summary>show</summary>
<p>


Print startup time per service

    systemd-analyze blame

</p>
</details>

### Logging

<details><summary>show</summary>
<p>


Log directory is /var/log/journal

    journalctl                       # Print all log entries
    journalctl -b                    # Print everything since boot
    journalctl -e                    # Print last page of the journal
    journalctl -f                    # Print last 10 entries and continues (like tail -f)
    journalctl -u <unit>             # Print for unit only
    journalctl --no-pager            # Non-interactive mode

    journalctl --vacuum-size=100M    # Remove logs until <100MB is used
    journalctl --vacuum-time=2weeks  # Remove logs older 2 weeks

</p>
</details>

### Other

<details><summary>show</summary>
<p>


    hostnamectl
    timedatectl
    localctl
    loginctl

    systemd-detect-virt

</p>
</details>

### Misc

<details><summary>show</summary>
<p>


-   [systemd DBUS API](www.freedesktop.org/wiki/Software/systemd/dbus/)

</p>
</details>
