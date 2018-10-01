# Grep Study Guide
by: Coleman Word

- [Grep Study Guide](#grep-study-guide)
    - [How do you use grep and git together?](#how-do-you-use-grep-and-git-together)


## How do you use grep and git together?

<details><summary>show</summary>
<p>


Git Grep will return a list of lines matching a pattern.

Running:
```bash
$ git grep aliases
```
will show all the files containing the string *aliases*.

![git grep aliases](http://i.imgur.com/DL2zpQ9.png)

*Press `q` to quit.*

You can also use multiple flags for more advanced search. For example:

 * `-e` The next parameter is the pattern (e.g., regex)
 * `--and`, `--or` and `--not` Combine multiple patterns.

Use it like this:
```bash
 $ git grep -e pattern --and -e anotherpattern
```

[*Read more about the Git `grep` command.*](http://git-scm.com/docs/git-grep)

</p>
</details>
