# Git Study Guide
by: Coleman Word

***

- [Git Study Guide](#git-study-guide)
    - [Github Resources](#github-resources)
    - [What are some keyboard shortcuts?](#what-are-some-keyboard-shortcuts)
    - [How do you view commit history by author?](#how-do-you-view-commit-history-by-author)
    - [How do you clone a repository?](#how-do-you-clone-a-repository)
    - [How do you compare all branches to another branch?](#how-do-you-compare-all-branches-to-another-branch)
    - [How do you compare branches?](#how-do-you-compare-branches)
    - [How do you compare branches across forked repositories?](#how-do-you-compare-branches-across-forked-repositories)
    - [How do you create a gist?](#how-do-you-create-a-gist)
    - [How do you close issues via commit messages?](#how-do-you-close-issues-via-commit-messages)
    - [How do you cross-link issues?](#how-do-you-cross-link-issues)
    - [How do you Lock Conversations?](#how-do-you-lock-conversations)
    - [How do you use filters?](#how-do-you-use-filters)
    - [How do you use emojis?](#how-do-you-use-emojis)
    - [How do you create task lists in issues and pull requests?](#how-do-you-create-task-lists-in-issues-and-pull-requests)
    - [How do you revert a pull request?](#how-do-you-revert-a-pull-request)
    - [How do you check the differences in a patch or pull request?](#how-do-you-check-the-differences-in-a-patch-or-pull-request)
    - [What is hub?](#what-is-hub)
    - [What is the CONTRIBUTING.md file?](#what-is-the-contributingmd-file)
    - [What is the ISSUE_TEMPLATE file?](#what-is-the-issuetemplate-file)
    - [What is the PULL_REQUEST_TEMPLATE file?](#what-is-the-pullrequesttemplate-file)
    - [What is the Github Student Developer Pack?](#what-is-the-github-student-developer-pack)
    - [How do you view your SSH keys?](#how-do-you-view-your-ssh-keys)
    - [How do you remove all deleted files from the working tree?](#how-do-you-remove-all-deleted-files-from-the-working-tree)
    - [How do you move to the previous branch?](#how-do-you-move-to-the-previous-branch)
    - [How do you check out pull requests?](#how-do-you-check-out-pull-requests)
    - [How do you use an empty commit?](#how-do-you-use-an-empty-commit)
    - [How do you use grep and git together?](#how-do-you-use-grep-and-git-together)
    - [How do you view merged branches?](#how-do-you-view-merged-branches)
    - [How do you use fixup and autosquash when there is something wrong with the previous commit?](#how-do-you-use-fixup-and-autosquash-when-there-is-something-wrong-with-the-previous-commit)
    - [How do you open a web server for browsing local repos?](#how-do-you-open-a-web-server-for-browsing-local-repos)
    - [How do you use aliases to define your own git calls?](#how-do-you-use-aliases-to-define-your-own-git-calls)
    - [How can you use auto-correct?](#how-can-you-use-auto-correct)
    - [How do you](#how-do-you)
    - [How do you](#how-do-you)
    - [How do you](#how-do-you)
    - [How do you](#how-do-you)


***

## Github Resources 

<details><summary>show</summary>
<p>

GitHub Resources
| Title | Link |
| ----- | ---- |
| GitHub Explore | https://github.com/explore |
| GitHub Blog | https://github.com/blog |
| GitHub Help | https://help.github.com/ |
| GitHub Training | https://training.github.com/ |
| GitHub Developer | https://developer.github.com/ |
| Github Education (Free Micro Account and other stuff for students) | https://education.github.com/ |
Official Git Site | http://git-scm.com/ |
| Official Git Video Tutorials | http://git-scm.com/videos |
| Code School Try Git | http://try.github.com/ |
| Introductory Reference & Tutorial for Git | http://gitref.org/ |
| Official Git Tutorial | http://git-scm.com/docs/gittutorial |
| Everyday Git | http://git-scm.com/docs/everyday |
| Git Immersion | http://gitimmersion.com/ |
| Git God | https://github.com/gorosgobe/git-god |
| Git for Computer Scientists | http://eagain.net/articles/git-for-computer-scientists/ |
| Git Magic | http://www-cs-students.stanford.edu/~blynn/gitmagic/ |
| GitHub Training Kit | https://training.github.com/kit/ |
| Git Visualization Playground | http://onlywei.github.io/explain-git-with-d3/#freeplay |
| Learn Git Branching | http://pcottle.github.io/learnGitBranching/ |
| A collection of useful .gitignore templates | https://github.com/github/gitignore |
| Unixorn's git-extra-commands collection of git scripts | https://github.com/unixorn/git-extra-commands |
Linus Torvalds on Git | https://www.youtube.com/watch?v=4XpnKHJAok8 |
| Introduction to Git with Scott Chacon | https://www.youtube.com/watch?v=ZDR433b0HJY |
| Git From the Bits Up | https://www.youtube.com/watch?v=MYP56QJpDr4 |
| Graphs, Hashes, and Compression, Oh My! | https://www.youtube.com/watch?v=ig5E8CcdM9g |
| GitHub Training & Guides | https://www.youtube.com/watch?list=PLg7s6cbtAD15G8lNyoaYDuKZSKyJrgwB-&v=FyfwLX4HAxM |



</p>
</details>

***

## What are some keyboard shortcuts?

<details><summary>show</summary>
<p>

When on a repository page, keyboard shortcuts allow you to navigate easily.

* Pressing t will bring up a file explorer.
* Pressing w will bring up the branch selector.
* Pressing s will focus the search field for the current repository.
* Pressing Backspace to delete the “This repository” pill changes the field to search all of GitHub.
* Pressing l will edit labels on existing Issues.
* Pressing y when looking at a file (e.g., https://github.com/tiimgreen/github-cheat-sheet/blob/master/README.md) will change your URL to one which, in effect, freezes the page you are looking at. If this code changes, you will still be able to see what you saw at that current time.
* To see all of the shortcuts for the current page press ?:

</p>
</details>

***

## How do you view commit history by author?


<details><summary>show</summary>
<p>

To view all commits on a repo by author add `?author={user}` to the URL.

```
https://github.com/rails/rails/commits/master?author=dhh
```

![DHH commit history](http://i.imgur.com/S7AE29b.png)

[*Read more about the differences between commits views.*](https://help.github.com/articles/differences-between-commit-views/)

</p>
</details>

***

## How do you clone a repository?

<details><summary>show</summary>
<p>

When cloning a repository the `.git` can be left off the end.

```bash
$ git clone https://github.com/tiimgreen/github-cheat-sheet
```

[*Read more about the Git `clone` command.*](http://git-scm.com/docs/git-clone)

</p>
</details>

***

## How do you compare all branches to another branch?

<details><summary>show</summary>
<p>


If you go to the repo's [Branches](https://github.com/tiimgreen/github-cheat-sheet/branches) page, next to the Commits button:

```
https://github.com/{user}/{repo}/branches
```

... you would see a list of all branches which are not merged into the main branch.

From here you can access the compare page or delete a branch with a click of a button.

![Compare branches not merged into master in rails/rails repo - https://github.com/rails/rails/branches](http://i.imgur.com/0FEe30z.png)


</p>
</details>

***

## How do you compare branches?


<details><summary>show</summary>
<p>

To use GitHub to compare branches, change the URL to look like this:

```
https://github.com/{user}/{repo}/compare/{range}
```

where `{range} = master...4-1-stable`

For example:

```
https://github.com/rails/rails/compare/master...4-1-stable
```

![Rails branch compare example](http://i.imgur.com/tIRCOsK.png)

`{range}` can be changed to things like:

```
https://github.com/rails/rails/compare/master@{1.day.ago}...master
https://github.com/rails/rails/compare/master@{2014-10-04}...master
```

*Here, dates are in the format `YYYY-MM-DD`*

![Another compare example](http://i.imgur.com/5dtzESz.png)

Branches can also be compared in `diff` and `patch` views:

```
https://github.com/rails/rails/compare/master...4-1-stable.diff
https://github.com/rails/rails/compare/master...4-1-stable.patch
```

[*Read more about comparing commits across time.*](https://help.github.com/articles/comparing-commits-across-time/)

</p>
</details>

***

## How do you compare branches across forked repositories?

<details><summary>show</summary>
<p>

To use GitHub to compare branches across forked repositories, change the URL to look like this:

```
https://github.com/{user}/{repo}/compare/{foreign-user}:{branch}...{own-branch}
```

For example:

```
https://github.com/rails/rails/compare/byroot:master...master
```

![Forked branch compare](http://i.imgur.com/Q1W6qcB.png)


</p>
</details>

***

## How do you create a gist?

<details><summary>show</summary>
<p>

[Gists](https://gist.github.com/) are an easy way to work with small bits of code without creating a fully fledged repository.

![Gist](http://i.imgur.com/VkKI1LC.png?1)

Add `.pibb` to the end of any Gist URL ([like this](https://gist.github.com/tiimgreen/10545817.pibb)) in order to get the *HTML-only* version suitable for embedding in any other site.

Gists can be treated as a repository so they can be cloned like any other:

```bash
$ git clone https://gist.github.com/tiimgreen/10545817
```

![Gists](http://i.imgur.com/BcFzabp.png)

This means you also can modify and push updates to Gists:

```bash
$ git commit
$ git push
Username for 'https://gist.github.com':
Password for 'https://tiimgreen@gist.github.com':
```

However, Gists do not support directories. All files need to be added to the repository root.
[*Read more about creating Gists.*](https://help.github.com/articles/creating-gists/)


</p>
</details>

***

## How do you close issues via commit messages?


<details><summary>show</summary>
<p>

If a particular commit fixes an issue, any of the keywords `fix/fixes/fixed`, `close/closes/closed` or `resolve/resolves/resolved`, followed by the issue number, will close the issue once it is committed to the repository's default branch.

```bash
$ git commit -m "Fix screwup, fixes #12"
```

This closes the issue and references the closing commit.

![Closing Repo](http://i.imgur.com/Uh1gZdx.png)

[*Read more about closing Issues via commit messages.*](https://help.github.com/articles/closing-issues-via-commit-messages/)

</p>
</details>

***

## How do you cross-link issues?

<details><summary>show</summary>
<p>

If you want to link to another issue in the same repository, simply type hash `#` then the issue number, and it will be auto-linked.

To link to an issue in another repository, `{user}/{repo}#ISSUE_NUMBER`, e.g., `tiimgreen/toc#12`.

![Cross-Link Issues](https://camo.githubusercontent.com/447e39ab8d96b553cadc8d31799100190df230a8/68747470733a2f2f6769746875622d696d616765732e73332e616d617a6f6e6177732e636f6d2f626c6f672f323031312f736563726574732f7265666572656e6365732e706e67)

</p>
</details>

***

## How do you Lock Conversations?


<details><summary>show</summary>
<p>

Pull Requests and Issues can now be locked by owners or collaborators of the repo.

![Lock conversation](https://cloud.githubusercontent.com/assets/2723/3221693/bf54dd44-f00d-11e3-8eb6-bb51e825bc2c.png)

This means that users who are not collaborators on the project will no longer be able to comment.

![Comments locked](https://cloud.githubusercontent.com/assets/2723/3221775/d6e513b0-f00e-11e3-9721-2131cb37c906.png)

[*Read more about locking conversations.*](https://github.com/blog/1847-locking-conversations)


</p>
</details>

***

## How do you use filters?

<details><summary>show</summary>
<p>


Both issues and pull requests allow filtering in the user interface.

For the Rails repo: https://github.com/rails/rails/issues, the following filter is built by selecting the label "activerecord":

`is:issue label:activerecord`

But, you can also find all issues that are NOT labeled activerecord:

`is:issue -label:activerecord`

Additionally, this also works for pull requests:

`is:pr -label:activerecord`

Github has tabs for displaying open or closed issues and pull requests but you
can also see merged pull requests.  Just put the following in the filter:

`is:merged`

[*Read more about searching issues.*](https://help.github.com/articles/searching-issues/)

Finally, github now allows you to filter by the Status API's status.

Pull requests with only successful statuses:

`status:success`

[*Read more about searching on the Status API.*](https://github.com/blog/2014-filter-pull-requests-by-status)

</p>
</details>


***

## How do you use emojis?

<details><summary>show</summary>
<p>

Emojis can be added to Pull Requests, Issues, commit messages, repository descriptions, etc. using `:name_of_emoji:`.

The full list of supported Emojis on GitHub can be found at [emoji-cheat-sheet.com](http://www.emoji-cheat-sheet.com/) or [scotch-io/All-Github-Emoji-Icons](https://github.com/scotch-io/All-Github-Emoji-Icons).
A handy emoji search engine can be found at [emoji.muan.co](http://emoji.muan.co/).

The top 5 used Emojis on GitHub are:

1. `:shipit:`
2. `:sparkles:`
3. `:-1:`
4. `:+1:`
5. `:clap:`

</p>
</details>



***

## How do you create task lists in issues and pull requests?

<details><summary>show</summary>
<p>

In Issues and Pull requests check boxes can be added with the following syntax (notice the space):

```
- [ ] Be awesome
- [ ] Prepare dinner
  - [ ] Research recipe
  - [ ] Buy ingredients
  - [ ] Cook recipe
- [ ] Sleep
```

![Task List](http://i.imgur.com/jJBXhsY.png)

When they are clicked, they will be updated in the pure Markdown:

```
- [x] Be awesome
- [ ] Prepare dinner
  - [x] Research recipe
  - [x] Buy ingredients
  - [ ] Cook recipe
- [ ] Sleep
```

[*Read more about task lists.*](https://help.github.com/articles/writing-on-github/#task-lists)

</p>
</details>


***

## How do you revert a pull request?

<details><summary>show</summary>
<p>

After a pull request is merged, you may find it does not help anything or it was a bad decision to merge the pull request.

You can revert it by clicking the **Revert** button on the right side of a commit in the pull request page to create a pull request with reverted changes to this specific pull request.

</p>
</details>


***

## How do you check the differences in a patch or pull request?

<details><summary>show</summary>
<p>

You can get the diff of a Pull Request by adding a `.diff` or `.patch`
extension to the end of the URL. For example:

```
https://github.com/tiimgreen/github-cheat-sheet/pull/15
https://github.com/tiimgreen/github-cheat-sheet/pull/15.diff
https://github.com/tiimgreen/github-cheat-sheet/pull/15.patch
```

The `.diff` extension would give you this in plain text:

```
diff --git a/README.md b/README.md
index 88fcf69..8614873 100644
--- a/README.md
+++ b/README.md
@@ -28,6 +28,7 @@ All the hidden and not hidden features of Git and GitHub. This cheat sheet was i
 - [Merged Branches](#merged-branches)
 - [Quick Licensing](#quick-licensing)
 - [TODO Lists](#todo-lists)
+- [Relative Links](#relative-links)
 - [.gitconfig Recommendations](#gitconfig-recommendations)
     - [Aliases](#aliases)
     - [Auto-correct](#auto-correct)
@@ -381,6 +382,19 @@ When they are clicked, they will be updated in the pure Markdown:
 - [ ] Sleep

(...)
```


</p>
</details>


***

## What is hub?

<details><summary>show</summary>
<p>

[Hub](https://github.com/github/hub) is a command line Git wrapper that gives you extra features and commands that make working with GitHub easier.

This allows you to do things like:

```bash
$ hub clone tiimgreen/toc
```

[*Check out some more cool commands Hub has to offer.*](https://github.com/github/hub#commands)


</p>
</details>


***

## What is the CONTRIBUTING.md file? 

<details><summary>show</summary>
<p>

CONTRIBUTING File
Adding a `CONTRIBUTING` or `CONTRIBUTING.md` file to either the root of your repository or a `.github` directory will add a link to your file when a contributor creates an Issue or opens a Pull Request.

![Contributing Guidelines](https://camo.githubusercontent.com/71995d6b0e620a9ef1ded00a04498241c69dd1bf/68747470733a2f2f6769746875622d696d616765732e73332e616d617a6f6e6177732e636f6d2f736b697463682f6973737565732d32303132303931332d3136323533392e6a7067)

[*Read more about contributing guidelines.*](https://github.com/blog/1184-contributing-guidelines)


</p>
</details>


***

## What is the ISSUE_TEMPLATE file? 

<details><summary>show</summary>
<p>


You can define a template for all new issues opened in your project. The content of this file will pre-populate the new issue box when users create new issues. Add an `ISSUE_TEMPLATE` or `ISSUE_TEMPLATE.md` file to either the root of your repository or a `.github` directory.

[*Read more about issue templates.*](https://github.com/blog/2111-issue-and-pull-request-templates)

[Issue template file generator](https://www.talater.com/open-source-templates/)

![GitHub Issue template](https://cloud.githubusercontent.com/assets/25792/13120859/733479fe-d564-11e5-8a1f-a03f95072f7a.png)


</p>
</details>


***

## What is the PULL_REQUEST_TEMPLATE file?

<details><summary>show</summary>
<p>

You can define a template for all new pull requests opened in your project. The content of this file will pre-populate the text area when users create pull requests. Add a `PULL_REQUEST_TEMPLATE` or `PULL_REQUEST_TEMPLATE.md` file to either the root of your repository or a `.github` directory.

[*Read more about pull request templates.*](https://github.com/blog/2111-issue-and-pull-request-templates)

[Pull request template file generator](https://www.talater.com/open-source-templates/)


</p>
</details>


***

## What is the Github Student Developer Pack?

<details><summary>show</summary>
<p>

If you are a student you will be eligible for the GitHub Student Developer Pack. This gives you free credit, free trials and early access to software that will help you when developing.

![GitHub Student Developer Pack](http://i.imgur.com/9ru3K43.png)

[*Read more about GitHub's Student Developer Pack*](https://education.github.com/pack)

</p>
</details>


***

## How do you view your SSH keys?


<details><summary>show</summary>
<p>

You can get a list of public ssh keys in plain text format by visiting:

```
https://github.com/{user}.keys
```

e.g. [https://github.com/tiimgreen.keys](https://github.com/tiimgreen.keys)

[*Read more about accessing public ssh keys.*](https://changelog.com/github-exposes-public-ssh-keys-for-its-users/)


</p>
</details>


***

## How do you remove all deleted files from the working tree?

<details><summary>show</summary>
<p>

When you delete a lot of files using `/bin/rm` you can use the following command to remove them from the working tree and from the index, eliminating the need to remove each one individually:

```bash
$ git rm $(git ls-files -d)
```

For example:

```bash
$ git status
On branch master
Changes not staged for commit:
	deleted:    a
	deleted:    c

$ git rm $(git ls-files -d)
rm 'a'
rm 'c'

$ git status
On branch master
Changes to be committed:
	deleted:    a
	deleted:    c
```

</p>
</details>

***

## How do you move to the previous branch?

<details><summary>show</summary>
<p>

To move to the previous branch in Git:

```bash
$ git checkout -
# Switched to branch 'master'

$ git checkout -
# Switched to branch 'next'

$ git checkout -
# Switched to branch 'master'
```

[*Read more about Git branching.*](http://git-scm.com/book/en/Git-Branching-Basic-Branching-and-Merging)

</p>
</details>

***

## How do you check out pull requests?

<details><summary>show</summary>
<p>

Pull Requests are special branches on the GitHub repository which can be retrieved locally in several ways:

Retrieve a specific Pull Request and store it temporarily in `FETCH_HEAD` for quickly `diff`-ing or `merge`-ing:

```bash
$ git fetch origin refs/pull/[PR-Number]/head
```

Acquire all Pull Request branches as local remote branches by refspec:

```bash
$ git fetch origin '+refs/pull/*/head:refs/remotes/origin/pr/*'
```

Or setup the remote to fetch Pull Requests automatically by adding these corresponding lines in your repository's `.git/config`:

```
[remote "origin"]
    fetch = +refs/heads/*:refs/remotes/origin/*
    url = git@github.com:tiimgreen/github-cheat-sheet.git
```

```
[remote "origin"]
    fetch = +refs/heads/*:refs/remotes/origin/*
    url = git@github.com:tiimgreen/github-cheat-sheet.git
    fetch = +refs/pull/*/head:refs/remotes/origin/pr/*
```

For Fork-based Pull Request contributions, it's useful to `checkout` a remote branch representing the Pull Request and create a local branch from it:

```bash
$ git checkout pr/42 pr-42
```

Or should you work on more repositories, you can globally configure fetching pull requests in the global git config instead.

```bash
git config --global --add remote.origin.fetch "+refs/pull/*/head:refs/remotes/origin/pr/*"
```

This way, you can use the following short commands in all your repositories:

```bash
git fetch origin
```

```bash
git checkout pr/42
```

[*Read more about checking out pull requests locally.*](https://help.github.com/articles/checking-out-pull-requests-locally/)

</p>
</details>

***

## How do you use an empty commit?

<details><summary>show</summary>
<p>

Commits can be pushed with no code changes by adding `--allow-empty`:

```bash
$ git commit -m "Big-ass commit" --allow-empty
```

Some use-cases for this (that make sense), include:

 - Annotating the start of a new bulk of work or a new feature.
 - Documenting when you make changes to the project that aren't code related.
 - Communicating with people using your repository.
 - The first commit of a repository: `git commit -m "Initial commit" --allow-empty`.


</p>
</details>

***

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

***

## How do you view merged branches?

<details><summary>show</summary>
<p>

Running:

```bash
$ git branch --merged
```

will give you a list of all branches that have been merged into your current branch.

Conversely:

```bash
$ git branch --no-merged
```

will give you a list of branches that have not been merged into your current branch.

[*Read more about the Git `branch` command.*](http://git-scm.com/docs/git-branch)

</p>
</details>

***

## How do you use fixup and autosquash when there is something wrong with the previous commit?

<details><summary>show</summary>
<p>

If there is something wrong with a previous commit (can be one or more from HEAD), for example `abcde`, run the following command after you've amended the problem:
```bash
$ git commit --fixup=abcde
$ git rebase abcde^ --autosquash -i
```
[*Read more about the Git `commit` command.*](http://git-scm.com/docs/git-commit)
[*Read more about the Git `rebase` command.*](http://git-scm.com/docs/git-rebase)

</p>
</details>

***

## How do you open a web server for browsing local repos?

<details><summary>show</summary>
<p>

Use the Git `instaweb` command to instantly browse your working repository in `gitweb`. This command is a simple script to set up `gitweb` and a web server for browsing the local repository.

```bash
$ git instaweb
```

opens:

![Git instaweb](http://i.imgur.com/Dxekmqc.png)

[*Read more about the Git `instaweb` command.*](http://git-scm.com/docs/git-instaweb)

</p>
</details>

***

## How do you use aliases to define your own git calls?

<details><summary>show</summary>
<p>

Aliases are helpers that let you define your own git calls. For example you could set `git a` to run `git add --all`.

To add an alias, either navigate to `~/.gitconfig` and fill it out in the following format:

```
[alias]
  co = checkout
  cm = commit
  p = push
  # Show verbose output about tags, branches or remotes
  tags = tag -l
  branches = branch -a
  remotes = remote -v
```

...or type in the command-line:

```bash
$ git config --global alias.new_alias git_function
```

For example:

```bash
$ git config --global alias.cm commit
```

For an alias with multiple functions use quotes:

```bash
$ git config --global alias.ac 'add -A . && commit'
```

Some useful aliases include:

| Alias | Command | What to Type |
| --- | --- | --- |
| `git cm` | `git commit` | `git config --global alias.cm commit` |
| `git co` | `git checkout` | `git config --global alias.co checkout` |
| `git ac` | `git add . -A` `git commit` | `git config --global alias.ac '!git add -A && git commit'` |
| `git st` | `git status -sb` | `git config --global alias.st 'status -sb'` |
| `git tags` | `git tag -l` | `git config --global alias.tags 'tag -l'` |
| `git branches` | `git branch -a` | `git config --global alias.branches 'branch -a'` |
| `git cleanup` | `git branch --merged \| grep -v '*' \| xargs git branch -d` | `git config --global alias.cleanup "!git branch --merged \| grep -v '*' \| xargs git branch -d"` |
| `git remotes` | `git remote -v` | `git config --global alias.remotes 'remote -v'` |
| `git lg` | `git log --color --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit --` | `git config --global alias.lg "log --color --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit --"` |

*Some Aliases are taken from [@mathiasbynens](https://github.com/mathiasbynens) dotfiles: https://github.com/mathiasbynens/dotfiles/blob/master/.gitconfig*

</p>
</details>

***

## How can you use auto-correct?

<details><summary>show</summary>
<p>

Git gives suggestions for misspelled commands and if auto-correct is enabled the command can be fixed and executed automatically. Auto-correct is enabled by specifying an integer which is the delay in tenths of a second before git will run the corrected command. Zero is the default value where no correcting will take place, and a negative value will run the corrected command with no delay.

For example, if you type `git comit` you will get this:

```bash
$ git comit -m "Message"
# git: 'comit' is not a git command. See 'git --help'.

# Did you mean this?
#   commit
```

Auto-correct can be enabled like this (with a 1.5 second delay):

```bash
$ git config --global help.autocorrect 15
```

So now the command `git comit` will be auto-corrected to `git commit` like this:

```bash
$ git comit -m "Message"
# WARNING: You called a Git command named 'comit', which does not exist.
# Continuing under the assumption that you meant 'commit'
# in 1.5 seconds automatically...
```

The delay before git will rerun the command is so the user has time to abort.


</p>
</details>

***

## How do you 

<details><summary>show</summary>
<p>



</p>
</details>

***

## How do you 

<details><summary>show</summary>
<p>



</p>
</details>

***

## How do you 

<details><summary>show</summary>
<p>



</p>
</details>

***

## How do you 

<details><summary>show</summary>
<p>



</p>
</details>
