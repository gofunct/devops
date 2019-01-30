# Gerrit Study Guide

## How can you easily checkout changes?

<details><summary>show</summary>
<p>

Changes correspond to branches with the naming schema "gerrit_&lt;change
number>". So you can

    git checkout gerrit_12345

</p>
</details>

## How do you delete changes?

<details><summary>show</summary>
<p>

Sometime you want to delete a change to make it invisible for everyone
(for example when you did commit an unencrypted secret...). This is
possible only via the SQL interface which you can enter with

    ssh <gerrit host>:29418 gerrit gsql

and issue a delete with:

    update changes set status='d' where change_id='<change id>';

</p>
</details>

## How can you search for commit text?

<details><summary>show</summary>
<p>


It is a bit non-obvious to search for text in commit messages. When just
searching "message:" no results are given. This is because "status:"
auto-defaults to a useless values. So you need to specify "status:" with
a useful value:

    status:merged message:critical bugfix
    status:open   message:Patch

</p>
</details>

### How can you solve errors?

<details><summary>show</summary>
<p>


#### Bad pack header

This happens from time to time if you do not garbage collect and hit the
open file limit. To solve: trigger a garbage collection

    ssh <gerrit host>:29418 gerrit gc <project name>

#### Unpack failed: error Missing unknown xxx

Solution here is to abandon the problematic changes. This cannot be done
via the GUI and only by admins

    # Find Gerrit id
    ssh <gerrit hxost>:29418 gerrit query xxx

    # Abandon the change
    ssh <gerrit host>:29418 gerrit gsql -c "update changes set open='N',status='A' where change_id=<id>;"

</p>
</details>
