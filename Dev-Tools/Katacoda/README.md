# Katacoda Study Guide

- [Katacoda Study Guide](#katacoda-study-guide)
    - [Scenario Example](#scenario-example)
    - [What is Katacoda?](#what-is-katacoda)
    - [How does it work?](#how-does-it-work)
    - [How do you configure git webhooks?](#how-do-you-configure-git-webhooks)
    - [What layouts can I use for tutorials?](#what-layouts-can-i-use-for-tutorials)
    - [What environments can I use?](#what-environments-can-i-use)
    - [How do you use Katacodas custom markdown to integrate the terminal?](#how-do-you-use-katacodas-custom-markdown-to-integrate-the-terminal)
        - [Allow a code block to be executed](#allow-a-code-block-to-be-executed)
        - [Multiple Hosts](#multiple-hosts)
        - [Multiple Terminals](#multiple-terminals)
        - [Allow a code block to be copied](#allow-a-code-block-to-be-copied)
    - [How do you use Katacodas custom markdown to integrate the editor?](#how-do-you-use-katacodas-custom-markdown-to-integrate-the-editor)
        - [Copy code to the editor](#copy-code-to-the-editor)
    - [Give an example index.json file](#give-an-example-indexjson-file)
    - [Index.json cheatsheet](#indexjson-cheatsheet)
    - [How do you connect to ports and use the web ui?](#how-do-you-connect-to-ports-and-use-the-web-ui)
        - [Example](#example)
        - [Display dashboard tabs](#display-dashboard-tabs)
    - [How do you customize an environment with scripts?](#how-do-you-customize-an-environment-with-scripts)
    - [How do you upload files to an environment?](#how-do-you-upload-files-to-an-environment)
    - [How do you ?](#how-do-you)
    - [How do you ?](#how-do-you)
    - [How do you ?](#how-do-you)
    - [How do you ?](#how-do-you)
    - [How do you ?](#how-do-you)
    - [How do you ?](#how-do-you)
    - [How do you ?](#how-do-you)

## Scenario Example

<details><summary>show</summary>
<p>

[Scenario-Example]https://github.com/katacoda/scenario-example

</p>
</details>


## What is Katacoda?

<details><summary>show</summary>
<p>

* Katacoda is an interactive learning and training platform for software developers. Each student is given access to a new environment without the need to install all the required component by themselves. 
* Katacoda also provides isolation for each student, so they can explore and push the limits of their learning skills without worry about breaking the environment for fellow students.


</p>
</details>

## How does it work?

<details><summary>show</summary>
<p>

* During the training course, the student will be using a hosted environment that is created just for them.
* This environment is not shared with other users of the system. Each environment is limited to a one hour session with a new environment will be created when the user reloads the page.

</p>
</details>

## How do you configure git webhooks?

<details><summary>show</summary>
<p>

1. Create a new public Github Repository for your Katacoda scenarios.
2. Enter Your Git Repository URL
   * You can update via your profile settings.
3. In your Github Repository settings, add a new Github Webhook.
4. Configure the Webhook with the provided payload url and secret

Done.
When changes are made to the repository, they will automatically appear on your profile page.

</p>
</details>

## What layouts can I use for tutorials?

<details><summary>show</summary>
<p>

1. terminal
2. terminal-terminal	
3. editor-terminal
   * Make sure to include data-katacoda-layout="editor-terminal-split" in the html for the embedded code.
4. terminal-iframe
   *  Requires backend.port within index.json

</p>
</details>

## What environments can I use?

<details><summary>show</summary>
<p>

1. docker
2. swarm
3. kubernetes
4. kubernetes-cluster
   * Requires UI layout terminal-terminal
5. coreos
6. ubuntu
7. ansible
8. node6
9.  go
10. c#
11. dotnet
12. java8
13. bash


</p>
</details>

## How do you use Katacodas custom markdown to integrate the terminal?

<details><summary>show</summary>
<p>

### Allow a code block to be executed 
```
`some-command`{{execute}}
```
### Multiple Hosts
```
`some-command`{{execute HOST1}}
`some-command`{{execute HOST2}}
```

### Multiple Terminals
```
`some-command`{{execute T1}}
`some-command`{{execute T2}}
```

### Allow a code block to be copied 
```
`some-command`{{copy}}
```          

</p>
</details>

## How do you use Katacodas custom markdown to integrate the editor?

<details><summary>show</summary>
<p>

### Copy code to the editor

<pre class="file" data-filename="app.js" data-target="replace">var http = require('http');
var requestListener = function (req, res) {
  res.writeHead(200);
  res.end('Hello, World!');
}

var server = http.createServer(requestListener);
server.listen(3000, function() { console.log("Listening on port 3000")});
</pre>
          

<pre class="file" data-target="clipboard">Test</pre>
          

<pre class="file" data-target="regex???">Test</pre>


</p>
</details>

## Give an example index.json file

<details><summary>show</summary>
<p>

```json

{
  "pathwayTitle": "Pathway Title",
  "title": "Scenario Title",
  "description": "Scenario Description",
  "difficulty": "beginner",
  "time": "5-10 minutes",
  "details": {
    "steps": [
      {
        "title": "Step Title",
        "text": "step1.md",
        "answer": "step1-answer.md",
        "verify": "step1-verify.sh",
        "courseData": "run-command-in-background.sh",
        "code": "run-command-in-terminal.sh"
      },
      {
        "title": "Step 2 - Step Title",
        "text": "step2.md"
      }
    ],
    "intro": {
      "text": "intro.md",
      "courseData": "courseBase.sh",
      "credits": "",
      "code": "changecd.sh"
    },
    "finish": {
      "text": "finish.md"
    },
    "assets": {
      "client": [
        {
          "file": "docker-compose.yml",
          "target": "~/"
        }
      ],
      "host01": [
        {
          "file": "config.yml",
          "target": "~/"
        }
      ]
    }
  },
  "files": [
    "app.js"
  ],
  "environment": {
    "showdashboard": true,
    "dashboards": [{"name": "Tab Name", "port": 80}, {"name": "Tab Name", "port": 8080}],
    "uilayout": "terminal",
    "uimessage1": "\u001b[32mYour Interactive Bash Terminal. A safe place to learn and execute commands.\u001b[m\r\n"
  },
  "backend": {
    "imageid": "docker"
  }
}

```

</p>
</details>

## Index.json cheatsheet

<details><summary>show</summary>
<p>

| Titles and Descriptions |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|----------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------|
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
| pathwayTitle                     | Title of the collection of scenarios                                                                                                                                                                                                                                         |                                                                                                                                         |
| title                            | Title the scenario                                                                                                                                                                                                                                                           |                                                                                                                                         |
| description                      | Description of the scenario, displayed on the intro screen                                                                                                                                                                                                                   |                                                                                                                                         |
| difficulty                       | Provide users with a sense of the depth of content, displayed on the intro screen                                                                                                                                                                                            |                                                                                                                                         |
| time                             | Provide users with an estimated time to complete, displayed on the intro screen                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
| Details Node            |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
| steps                            | Details for the scenario steps                                                                                                                                                                                                                                               |                                                                                                                                         |
| intro                            | Details for the intro screen                                                                                                                                                                                                                                                 |                                                                                                                                         |
| finish                           | Details for the finish screen                                                                                                                                                                                                                                                |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
| Steps Node              |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
| title                            | Title for the step.                                                                                                                                                                                                                                                          |                                                                                                                                         |
| text                             | Filename containing the body for the step.                                                                                                                                                                                                                                   |                                                                                                                                         |
| answer                           | Filename containing the answer body for the step.                                                                                                                                                                                                                            |                                                                                                                                         |
| verify                           | Bash script to run to check if the user can proceed. More details here.                                                                                                                                                                                                      |                                                                                                                                         |
| courseData                       | Bash script to run in the background. More details here.                                                                                                                                                                                                                     |                                                                                                                                         |
| code                             | Bash script to run in the foreground. More details here.                                                                                                                                                                                                                     |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
| Intro/Finish Node       |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
| text                             | Filename containing the body for the screen.                                                                                                                                                                                                                                 |                                                                                                                                         |
| credits                          | Display a link on the intro screen, useful for linking to blog post for giving credit.                                                                                                                                                                                       |                                                                                                                                         |
| courseData                       | Bash script to run in the background. More details here.                                                                                                                                                                                                                     |                                                                                                                                         |
| code                             | Bash script to run in the foreground. More details here.                                                                                                                                                                                                                     |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
| Environment             |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
| hideintro                        | Boolean field that control if the intro step is shown to the user. By default it is hidden in the embedded mode.                                                                                                                                                             | Example:"environment": { "uilayout": "terminal", "hideintro": false}                                                                    |
| hidefinish                       | Boolean field that control if the finish step is shown to the user. By default it is hidden in the embedded mode.                                                                                                                                                            | Example:"environment": { "uilayout": "terminal", "hidefinish": true }                                                                   |
| uisettings                       | Especify the format of the files for syntax highlighting in the editor (useful when it doesn't recognize the extension you are using). The suported formats are: reactjs, makefile, dockerfile, dockercompose, csharp, javascript, golang, java and xml.                     | Example:"environment": {"uisettings": "yaml"}                                                                                           |
| icon                             | The icon for the scenario. The list of icons is at the home page. The possible values are: fa-docker, fa-weave, fa-kubernetes, fa-openshift, fa-dcos, fa-tensorflow, fa-runc, fa-coreos, fa-elixir, fa-csharp, fa-fsharp, fa-rlang, fa-golang, fa-java, fa-node and fa-ruby. | Example:"icon": "fa-node"                                                                                                               |
| showdashboard                    | Should Dashboard tabs be shown in UI.                                                                                                                                                                                                                                        |                                                                                                                                         |
| dashboards                       | Easily provide links for accessing dashboard/UI ports running in the environment. When using the terminal-iframe layout, it also show the exposed port of the container, without the need of open a new window. You can specify the name, port and the host identifier.      | To display a tab called App showing the port 8080 from the host host02:"dashboards": [{ "name": "App", "port": 8080, "host": "host02" } |
| uilayout                         | The layout ID provided by Katacoda. More details here.                                                                                                                                                                                                                       |                                                                                                                                         |
| uimessage1                       | Message to display at the top of the interative terminal.                                                                                                                                                                                                                    |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
| Backend                 |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
|                                  |                                                                                                                                                                                                                                                                              |                                                                                                                                         |
| imageid                          | Environment image id provided by Katacoda. More details here.                                                                                                                                                                                                                |                                                                                                                                         |

</p>
</details>

## How do you connect to ports and use the web ui?

<details><summary>show</summary>
<p>

Katacoda allows ports running include the environment to be made available to users. This is designed for displaying dashboards and UI elements running inside a container.

### Example

```

Render port 8500: https://[[HOST_SUBDOMAIN]]-8500-[[KATACODA_HOST]].environments.katacoda.com/

Render port 80: https://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/

Display page allowing user to select port:
https://[[HOST_SUBDOMAIN]]-[[KATACODA_HOST]].environments.katacoda.com/

```

### Display dashboard tabs

```

"environment": {
  "showdashboard": true,
  "dashboards": [{"name": "Display 80", "port": 80}, {"name": "Display 8080", "port": 8080}],
}
          

This style has been deprecated.
"environment": {
    "showdashboard": true,
    "dashboard": "Dashboard"
},
"backend": {
  "port": 80
}

```


</p>
</details>

## How do you customize an environment with scripts?

<details><summary>show</summary>
<p>


Bash scripts defined in courseData will run in the background when the user begins that step. For intro, this means the commands will run when the environment becomes available. For steps, this is when the user starts the step. This is perfect for configuring scenarios and environments to help guide the user.

Bash scripts defined in code work the same as courseData but will run in the foreground.


</p>
</details>

## How do you upload files to an environment?

<details><summary>show</summary>
<p>

Files in a assets directory within the scenario can have files uploaded to the client or host of the environment.

Index.json Example

```json
"details": {
  "intro": {
    "courseData": "courseBase.sh",
  },
  "assets": {
    "client": [
      {
        "file": "docker-compose.yml",
        "target": "~/"
      }
    ],
    "host01": [
      {
        "file": "config.yml",
        "target": "~/"
      }
    ]
  }
}

```    

</p>
</details>


