# OpenSSH Commands


## Copy Keys

<details><summary>show</summary>
<p>

    ssh-copy-id [-i keyfile] [email protected]


</p>
</details>

## 100% non-interactive SSH

<details><summary>show</summary>
<p>


    ssh -i my_priv_key -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no 
    -o PreferredAuthentications=publickey [email protected] -n "/bin/ls"

</p>
</details>

## Using SSH Agent

<details><summary>show</summary>
<p>


    eval $(ssh-agent)       # Start agent on demand

    ssh-add -l                      # List keys
    ssh-add                         # Add default key
    ssh-add ~/.ssh/id_rsa           # Add specific key
    ssh-add -t 3600 ~/.ssh/id_rsa   # Add with timeout
    ssh-add -D                      # Drop keys

    ssh -A ...          # Enforce agent forwarding

</p>
</details>


## [Transparent
Multi-Hop](http://sshmenu.sourceforge.net/articles/transparent-mulithop.html)

<details><summary>show</summary>
<p>


    ssh host1 -A -t host2 -A -t host3 ...

</p>
</details>

## [How to use a SOCKS
Proxy](http://magicmonster.com/kb/net/ssh/socks_proxy.html)

<details><summary>show</summary>
<p>


On the client start proxy by

    ssh -D <port> <remote host>

</p>
</details>

Extract Public Key from Private Key

<details><summary>show</summary>
<p>

Use ssh-keygen

    ssh-keygen -y -f ~/.ssh/id_rsa > ~/.ssh/id_rsa.pub

</p>
</details>

### ssh\_config


Read the [authorized\_keys
HowTo](http://www.eng.cam.ac.uk/help/jpmg/ssh/authorized_keys_howto.html)
to learn about syntax and options...

#### Per Host Keys

<details><summary>show</summary>
<p>

    Host example.com
    IdentityFile ~/.ssh/example.com_id_rsa

</p>
</details>

#### Agent Forwarding

<details><summary>show</summary>
<p>

[Agent
Forwarding](http://www.unixwiz.net/techtips/ssh-agent-forwarding.html)
explained with pictures! Configured in /etc/ssh\_config with

    Host *
    ForwardAgent yes


</p>
</details>

#### Multiplexing Connections

<details><summary>show</summary>
<p>


This is done using a "ControlMaster". This means the first SSH sessions
connection will be used for all following ones saving you the connection
overhead. **Note:** when you kill the first connection, all connections
will die! Also the first connection won't terminate even if you request
it to.

Create \~/.ssh/tmp before using below snippet

    ControlMaster auto
    ControlPath /home/<user name>/.ssh/tmp/%h_%p_%r

If you are using such an SSH configuration and want a real new
connection add "-S" to the ssh invocation.

</p>
</details>

#### Use Gateway/Jumphost

<details><summary>show</summary>
<p>


You can configure jumphosts using ProxyCommand and netcat:

    Host unreachable_host
      ProxyCommand ssh -e none gateway_host exec nc %h %p

</p>
</details>

#### Automatic Jump Host Proxying

<details><summary>show</summary>
<p>

    Host <your jump host>
      ForwardAgent yes
      Hostname <your jump host>
      User <your user name on jump host>

    # Note the server list can have wild cards, e.g. "webserver-* database*"
    Host <server list>
      ForwardAgent yes
      User <your user name on all these hosts>
      ProxyCommand ssh -q <your jump host> nc -q0 %h 22

</p>
</details>

#### Automatic Port Knocking

<details><summary>show</summary>
<p>


    Host myserver
       User myuser
       Host myserver.com
       ProxyCommand bash -c '/usr/bin/knock %h 1000 2000 3000 4000; sleep 1; exec /bin/nc %h %p'

</p>
</details>

### Troubleshooting


#### Pseudo-terminal will not be allocated...

<details><summary>show</summary>
<p>


This happens when piping shell commands through SSH. Try adding "-T" or
"-t -t" when doing sudo.

</p>
</details>

### Misc


#### [SFTP chroot with
    umask](http://jeff.robbins.ws/articles/setting-the-umask-for-sftp-transactions)

<details><summary>show</summary>
<p>

    How to enforce a umask with SFTP

        Subsystem sftp /usr/libexec/openssh/sftp-server -u 0002

</p>
</details>

#### Parallel SSH on Debian

<details><summary>show</summary>
<p>


        apt-get install pssh

    and use it like this

        pssh -h host_list.txt <command>
        pssh -i -t 60 -h host_list.txt -- <command>   # 60s timeout, list output

</p>
</details>

#### Clustered SSH on Debian

<details><summary>show</summary>
<p>


        apt-get install clusterssh

    and use it like this

        cssh server1 server2

</p>
</details>

#### Vim Remote File Editing

<details><summary>show</summary>
<p>


        vim scp:[email protected]//some/directory/file.txt

</p>
</details>

#### [MonkeySphere](http://web.monkeysphere.info/): Use GPG keys with SSH
    agent

<details><summary>show</summary>
<p>


        monkeysphere subkey-to-ssh-agent -t 3600

</p>
</details>

#### Port Knocking

<details><summary>show</summary>
<p>


Setup server:

    apt-get install knockd iptables-persistent

    # Change sequence numbers in /etc/knockd.conf
    # Default is sequence    = 7000,8000,9000

    # set START_KNOCKD=1 in /etc/default/knockd

    service knockd start

Use from client

    knock <server> <sequence>

e.g.

    knock example.com 7000 8000 9000

</p>
</details>

#### "Secret" Hot Keys

<details><summary>show</summary>
<p>

SSH Escape Key: Pressing "\~?" (directly following a newline) gives a
menu for escape sequences:

    Supported escape sequences:
      ~.  - terminate connection (and any multiplexed sessions)
      ~B  - send a BREAK to the remote system
      ~C  - open a command line
      ~R  - Request rekey (SSH protocol 2 only)
      ~^Z - suspend ssh
      ~#  - list forwarded connections
      ~&  - background ssh (when waiting for connections to terminate)
      ~?  - this message
      ~~  - send the escape character by typing it twice
    (Note that escapes are only recognized immediately after newline.)

</p>
</details>

#### SSHFS

<details><summary>show</summary>
<p>


To mount a remote home dir

     sshfs [email protected]: /mnt/home/user/

Unmount again with

    fuserumount -u /mnt/home/user

</p>
</details>
