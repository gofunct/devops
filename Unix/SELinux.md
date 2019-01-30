# Commands

## Show SELinux labels on files

<details><summary>show</summary>
<p>


    ls -Z

</p>
</details>

## Check if SELinux is on

<details><summary>show</summary>
<p>

    getenforce

</p>
</details>

## Disable temporarily

<details><summary>show</summary>
<p>


    echo 0 >/selinux/enforce

    # or

    setenforce 0

</p>
</details>

## Disable in grub.cfg by adding to kernel options

<details><summary>show</summary>
<p>


    selinux=0

</p>
</details>

## Disable in /etc/selinux/config

<details><summary>show</summary>
<p>

    SELINUX=disabled
    SELINUXTYPE=targeted
    SETLOCALDEFS=0

</p>
</details>

## Install [in Debian](https://wiki.debian.org/SELinux/Setup)

<details><summary>show</summary>
<p>

    apt-get install selinux-basics selinux-policy-default auditd
    selinux-activate
    # Reboot
    check-selinux-installation

</p>
</details>

## Get activation status:

<details><summary>show</summary>
<p>


    sestatus
    getenforce

</p>
</details>

## Install additional profiles

<details><summary>show</summary>
<p>

    semodule -i my_module.pp

</p>
</details>
