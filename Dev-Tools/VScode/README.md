# VScode Study Guide
by: Coleman Word

***

- [VScode Study Guide](#vscode-study-guide)
    - [Resources](#resources)
    - [Syntax](#syntax)
    - [How do you migrate keymaps from a different text editor/ide ?](#how-do-you-migrate-keymaps-from-a-different-text-editoride)
        - [Migrating from Atom](#migrating-from-atom)
        - [Migrating from Sublime Text](#migrating-from-sublime-text)
        - [Migrating from Visual Studio](#migrating-from-visual-studio)
        - [Migrating from Intellij IDEA](#migrating-from-intellij-idea)
    - [Productivity Extensions](#productivity-extensions)
        - [Azure Cosmos DB](#azure-cosmos-db)
        - [Azure IoT Toolkit](#azure-iot-toolkit)
        - [Bookmarks](#bookmarks)
        - [Copy Relative Path](#copy-relative-path)
        - [Create tests](#create-tests)
        - [Deploy](#deploy)
        - [Gi](#gi)
        - [Git History](#git-history)
        - [Git Project Manager](#git-project-manager)
        - [GitLink](#gitlink)
        - [GitLens](#gitlens)
        - [Git Indicators](#git-indicators)
        - [GitHub](#github)
        - [GitHub Pull Request Monitor](#github-pull-request-monitor)
        - [Icon Fonts](#icon-fonts)
        - [Kanban](#kanban)
        - [Live Server](#live-server)
        - [Multiple clipboards](#multiple-clipboards)
        - [npm Intellisense](#npm-intellisense)
        - [Partial Diff](#partial-diff)
        - [Paste JSON as Code](#paste-json-as-code)
        - [Path Intellisense](#path-intellisense)
        - [Project Manager](#project-manager)
        - [REST Client](#rest-client)
        - [Settings Sync](#settings-sync)
        - [Todo Tree](#todo-tree)
        - [Yo](#yo)
- [Keyboard Shortcuts](#keyboard-shortcuts)
    - [How do you access the command pallet?](#how-do-you-access-the-command-pallet)
    - [How do you quickly open a file?](#how-do-you-quickly-open-a-file)
- [Command Line Interface](#command-line-interface)
    - [What are some CLI commands you can use?](#what-are-some-cli-commands-you-can-use)
        - [open code with current directory](#open-code-with-current-directory)
        - [open the current directory in the most recently used code window](#open-the-current-directory-in-the-most-recently-used-code-window)
        - [create a new window](#create-a-new-window)
        - [change the language](#change-the-language)
        - [open diff editor](#open-diff-editor)
        - [open file at specific line and column <file:line[:character]>](#open-file-at-specific-line-and-column-filelinecharacter)
        - [see help options](#see-help-options)
        - [disable all extensions](#disable-all-extensions)

***

## Resources

<details><summary>show</summary>
<p>

- [Official website](https://code.visualstudio.com/)
- [Source code](https://github.com/microsoft/vscode) on GitHub
- [Releases (stable channel)](https://code.visualstudio.com/download)
- [Releases (insiders channel)](https://code.visualstudio.com/insiders)
- [Monthly iteration plans](https://github.com/Microsoft/vscode/issues?utf8=%E2%9C%93&q=label%3Aiteration-plan+)


</p>
</details>

***

## Syntax

<details><summary>show</summary>
<p>

Language packages extend the editor with syntax highlighting and/or snippets for a specific language or file format.

- [Arduino](https://marketplace.visualstudio.com/items?itemName=vsciot-vscode.vscode-arduino)
- [Blink](https://marketplace.visualstudio.com/items?itemName=melmass.blink)
- [Bolt](https://marketplace.visualstudio.com/items?itemName=smkamranqadri.vscode-bolt-language)
- [CMake](https://marketplace.visualstudio.com/items?itemName=twxs.cmake)
- [Dart](https://marketplace.visualstudio.com/items?itemName=Dart-Code.dart-code)
- [Dockerfile](https://marketplace.visualstudio.com/items?itemName=PeterJausovec.vscode-docker)
- [EJS](https://marketplace.visualstudio.com/items?itemName=QassimFarid.ejs-language-support)
- [Elixir](https://marketplace.visualstudio.com/items?itemName=mjmcloug.vscode-elixir)
- [Elm](https://marketplace.visualstudio.com/items?itemName=sbrink.elm)
- [Erlang](https://marketplace.visualstudio.com/items?itemName=pgourlain.erlang)
- [F#](https://marketplace.visualstudio.com/items?itemName=Ionide.Ionide-fsharp)
- [Fortran](https://marketplace.visualstudio.com/items?itemName=Gimly81.fortran)
- [Hack(HHVM)](https://marketplace.visualstudio.com/items?itemName=pranayagarwal.vscode-hack)
- [Handlebars](https://marketplace.visualstudio.com/items?itemName=andrejunges.Handlebars)
- [KL](https://marketplace.visualstudio.com/items?itemName=melmass.kl)
- [Kotlin](https://marketplace.visualstudio.com/items?itemName=mathiasfrohlich.Kotlin)
- [LaTeX](https://marketplace.visualstudio.com/items?itemName=torn4dom4n.latex-support)
- [Mason](https://marketplace.visualstudio.com/items?itemName=viatsko.html-mason)
- [openHAB](https://marketplace.visualstudio.com/items?itemName=openhab.openhab)
- [Parser 3](https://marketplace.visualstudio.com/items?itemName=viatsko.parser3)
- [Pascal](https://marketplace.visualstudio.com/items?itemName=alefragnani.pascal), or [OmniPascal](https://marketplace.visualstudio.com/items?itemName=Wosi.omnipascal) (only for Windows)
- [Perl HTML-Template](https://marketplace.visualstudio.com/items?itemName=viatsko.perl-html-template)
- [Protobuf](https://marketplace.visualstudio.com/items?itemName=peterj.proto)
- [Ruby](https://marketplace.visualstudio.com/items?itemName=groksrc.ruby)
- [Scala](https://marketplace.visualstudio.com/items?itemName=itryapitsin.Scala)
- [Stylus](https://marketplace.visualstudio.com/items?itemName=sysoev.language-stylus)
- [Swift](https://marketplace.visualstudio.com/items?itemName=Kasik96.swift)
- [VEX](https://marketplace.visualstudio.com/items?itemName=melmass.vex)
- [Zephir](https://marketplace.visualstudio.com/items?itemName=zephir-lang.zephir)

</p>
</details>


***

## How do you migrate keymaps from a different text editor/ide ?

<details><summary>show</summary>
<p>

The VSCode team provides keymaps from popular editors, making the transition to VSCode almost seamless and easy.

### [Migrating from Atom](https://marketplace.visualstudio.com/items?itemName=ms-vscode.atom-keybindings)

> Popular Atom keybindings for Visual Studio Code

### [Migrating from Sublime Text](https://marketplace.visualstudio.com/items?itemName=ms-vscode.sublime-keybindings)

> Popular Sublime Text keybindings for VS Code.

### [Migrating from Visual Studio](https://marketplace.visualstudio.com/items?itemName=ms-vscode.vs-keybindings)

> Popular Visual Studio keybindings for VS Code.

### [Migrating from Intellij IDEA](https://marketplace.visualstudio.com/items?itemName=k--kato.intellij-idea-keybindings)

> Popular Intellij IDEA keybindings for VS Code.

</p>
</details>

***

## Productivity Extensions

<details><summary>show</summary>
<p>

### [Azure Cosmos DB](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-cosmosdb)

> Browse your database inside the vs code editor

![](https://media.giphy.com/media/fnK9fzP80e7YfO7JAq/giphy.gif)

### [Azure IoT Toolkit](https://marketplace.visualstudio.com/items?itemName=vsciot-vscode.azure-iot-toolkit)

> Everything you need for the Azure IoT development: Interact with Azure IoT Hub, manage devices connected to Azure IoT Hub, and develop with code snippets for Azure IoT Hub

![](https://raw.githubusercontent.com/formulahendry/vscode-azure-iot-toolkit/master/images/device-explorer.png)

### [Bookmarks](https://marketplace.visualstudio.com/items?itemName=alefragnani.Bookmarks)

> Mark lines and jump to them

![](https://raw.githubusercontent.com/alefragnani/vscode-bookmarks/master/images/bookmarks-commands.png)

![](https://raw.githubusercontent.com/alefragnani/vscode-bookmarks/master/images/bookmarks-toggle.png)

### [Copy Relative Path](https://marketplace.visualstudio.com/items?itemName=alexdima.copy-relative-path)

> Copy Relative Path from a File

### [Create tests](https://marketplace.visualstudio.com/items?itemName=hardikmodha.create-tests)

> An extension to quickly generate test files.

![](https://media.giphy.com/media/1iqPhENd8SLd9SggeX/giphy.gif)

### [Deploy](https://marketplace.visualstudio.com/items?itemName=mkloubert.vs-deploy)

> Commands for upload or copy files of a workspace to a destination.

![](https://raw.githubusercontent.com/mkloubert/vs-deploy/master/img/demo.gif)

### [Gi](https://marketplace.visualstudio.com/items?itemName=rubbersheep.gi)
> Generating .gitignore files made easy.

![](https://raw.githubusercontent.com/hasit/vscode-gi/master/assets/gi.gif)

### [Git History](https://marketplace.visualstudio.com/items?itemName=donjayamanne.githistory)

> View git log, file or line History

![](https://raw.githubusercontent.com/DonJayamanne/gitHistoryVSCode/master/images/fileHistoryCommand.gif)

### [Git Project Manager](https://marketplace.visualstudio.com/items?itemName=felipecaputo.git-project-manager)

> Automatically indexes your git projects and lets you easily toggle between them

### [GitLink](https://marketplace.visualstudio.com/items?itemName=qezhu.gitlink)

> GoTo current file's online link in browser and Copy the link in clipboard.

![](https://raw.githubusercontent.com/qinezh/vscode-gitlink/master/images/how_to_use_it.gif)

### [GitLens](https://marketplace.visualstudio.com/items?itemName=eamodio.gitlens)

> Provides Git CodeLens information (most recent commit, # of authors), on-demand inline blame annotations, status bar blame information, file and blame history explorers, and commands to compare changes with the working tree or previous versions.

![](https://raw.githubusercontent.com/eamodio/vscode-git-codelens/master/images/gitlens-preview1.gif)

### [Git Indicators](https://marketplace.visualstudio.com/items?itemName=lamartire.git-indicators)

> Atom like git indicators on active panel

![](https://raw.githubusercontent.com/lamartire/vscode-git-indicators/master/preview/added.png)
![](https://raw.githubusercontent.com/lamartire/vscode-git-indicators/master/preview/removed.png)
![](https://raw.githubusercontent.com/lamartire/vscode-git-indicators/master/preview/modified.png)


### [GitHub](https://marketplace.visualstudio.com/items?itemName=KnisterPeter.vscode-github)

> Provides GitHub workflow support. For example browse project, issues, file (the current line), create and manage pull request. Support for other providers (e.g. gitlab or bitbucket) is planned.

> Have a look at the [README.md](https://github.com/KnisterPeter/vscode-github/blob/master/README.md) on how to get started with the setup for this extension.

### [GitHub Pull Request Monitor](https://marketplace.visualstudio.com/items?itemName=erichbehrens.pull-request-monitor)
> This extension uses the GitHub api to monitor the state of your pull requests and let you know when it's time to merge or if someone requested changes.

![GitHub Pull Request Monitor](https://raw.githubusercontent.com/erichbehrens/pull-request-monitor/master/images/statusBarItems.png)

### [Icon Fonts](https://marketplace.visualstudio.com/items?itemName=idleberg.icon-fonts)

> Snippets for popular icon fonts such as Font Awesome, Ionicons, Glyphicons, Octicons, Material Design Icons and many more!

### [Kanban](https://marketplace.visualstudio.com/items?itemName=mkloubert.vscode-kanban)

![kanban](https://raw.githubusercontent.com/mkloubert/vscode-kanban/master/img/demo1.gif)

> Simple Kanban board for use in Visual Studio Code, with time tracking and Markdown support.

### [Live Server](https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer)

> Launch a development local Server with live reload feature for static & dynamic pages.

![live-server](https://raw.githubusercontent.com/ritwickdey/vscode-live-server/master/images/Screenshot/vscode-live-server-animated-demo.gif)

### [Multiple clipboards](https://marketplace.visualstudio.com/items?itemName=slevesque.vscode-multiclip)

> Override the regular Copy and Cut commands to keep selections in a clipboard ring

### [npm Intellisense](https://marketplace.visualstudio.com/items?itemName=christian-kohler.npm-intellisense)

> Visual Studio Code plugin that autocompletes npm modules in import statements.

![npm-intellisense](https://raw.githubusercontent.com/ChristianKohler/NpmIntellisense/master/images/auto_complete.gif)


### [Partial Diff](https://marketplace.visualstudio.com/items?itemName=ryu1kn.partial-diff)

> Compare (diff) text selections within a file, across different files, or to the clipboard

![Partial Diff](https://raw.githubusercontent.com/ryu1kn/vscode-partial-diff/master/images/public.gif)

### [Paste JSON as Code](https://marketplace.visualstudio.com/items?itemName=quicktype.quicktype)

> Infer the structure of JSON and paste is as types in many programming languages

![Paste JSON as Code](https://raw.githubusercontent.com/quicktype/quicktype-vscode/master/media/demo.gif)

### [Path Intellisense](https://marketplace.visualstudio.com/items?itemName=christian-kohler.path-intellisense)

> Visual Studio Code plugin that autocompletes filenames

![](https://i.giphy.com/iaHeUiDeTUZuo.gif)

### [Project Manager](https://marketplace.visualstudio.com/items?itemName=alefragnani.project-manager)

> Easily switch between projects.

![](https://raw.githubusercontent.com/alefragnani/vscode-project-manager/master/images/project-manager-commands.png)

### [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)

> Allows you to send HTTP request and view the response in Visual Studio Code directly.

![](https://raw.githubusercontent.com/Huachao/vscode-restclient/master/images/usage.gif)

### [Settings Sync](https://marketplace.visualstudio.com/items?itemName=Shan.code-settings-sync)

> Synchronize settings, snippets, themes, file icons, launch, keybindings, workspaces and extensions across multiple machines using Github Gist

![Settings Sync](https://i.imgur.com/QZtaBca.gif)

### [Todo Tree](https://marketplace.visualstudio.com/items?itemName=Gruntfuggly.todo-tree)

> Custom keywords, highlighting, and colors for TODO comments. As well as a sidebar to view all your current tags.

![Todo Tree](https://thumbs.gfycat.com/PowerlessWindyCivet-size_restricted.gif)

### [Yo](https://marketplace.visualstudio.com/items?itemName=samverschueren.yo)

> Scaffold projects using [Yeoman](http://yeoman.io/)

![](https://raw.githubusercontent.com/SamVerschueren/vscode-yo/master/media/yo.gif)


</p>
</details>

***

# Keyboard Shortcuts

![Keyboard Shortcuts(Mac)]https://github.com/gofunct/cloudnative-engineer/blob/master/Assets/cheatsheets/vscode-kb-macos.png?raw=true

## How do you access the command pallet?

<details><summary>show</summary>
<p>

Access all available commands based on your current context.

Keyboard Shortcut: ⇧⌘P

</p>
</details>


## How do you quickly open a file?

<details><summary>show</summary>
<p>

Keyboard Shortcut: ⌘P

Repeat the Quick Open keyboard shortcut to cycle quickly between recently opened files.


</p>
</details>

***

# Command Line Interface

## What are some CLI commands you can use?

<details><summary>show</summary>
<p>

Make sure the VS Code binary is on your path so you can simply type 'code' to launch VS Code. See the platform specific setup topics if VS Code is added to your environment path during installation 

### open code with current directory
```
code .
```

### open the current directory in the most recently used code window

```
code -r .
```

### create a new window

```
code -n
```

### change the language

```
code --locale=es
```

### open diff editor

```
code --diff <file1> <file2>
```

### open file at specific line and column <file:line[:character]>

```
code --goto package.json:10:5
```

### see help options

```
code --help
```

### disable all extensions

```
code --disable-extensions .
```

</p>
</details>

