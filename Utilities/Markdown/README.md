# Documentation Study Guide
by: Coleman Word

***

- [Documentation Study Guide](#documentation-study-guide)
    - [Markdown Cheatsheet](#markdown-cheatsheet)
        - [Unsorted Lists](#unsorted-lists)
        - [Numbered Lists](#numbered-lists)
        - [Code Blocks](#code-blocks)
        - [Headings](#headings)
        - [URL's](#urls)
        - [Images](#images)
        - [Text Formatting](#text-formatting)
    - [Flash Card for Code Snippet](#flash-card-for-code-snippet)
    - [Flash Card for Text](#flash-card-for-text)

***

## Markdown Cheatsheet

<details><summary>show</summary>

### Unsorted Lists

```
* item one
* item two
* item three

```

### Numbered Lists

```
1. item one
1. item two
1. item three
1. item four
5. item five

```
### Code Blocks

```

```
insert code here
```

```

### Headings

```

# H1
## H2
### H3

```

### URL's

```

[Link Name](Link URL)

```

### Images

```

![Image Name](Image URL
)
```

### Text Formatting

```

**This is Bold**
*This is italics*
`This is code`

```

</p>
</details>

***

## Flash Card for Code Snippet

<details><summary>show</summary>
<p>

```
*** 

## Title

<details><summary>show</summary>
<p>

```
code here

```

</p>
</details>

```
</p>
</details>

***

## Flash Card for Text

<details><summary>show</summary>
<p>

```
*** 

## Title

* Add Description here

<details><summary>show</summary>
<p>

* Add lists, tips, links, etc here


</p>
</details>

```
</p>
</details>

***

## How do you set a link to an internal document? 

<details><summary>show</summary>
<p>

Relative links are recommended in your Markdown files when linking to internal content.

```markdown
[Link to a header](#awesome-section)
[Link to a file](docs/readme)
```

Absolute links have to be updated whenever the URL changes (e.g., repository renamed, username changed, project forked). Using relative links makes your documentation easily stand on its own.

[*Read more about relative links.*](https://help.github.com/articles/relative-links-in-readmes/)

</p>
</details>


***

## How do you render a pdf file?

<details><summary>show</summary>
<p>

Rendering PDF

GitHub supports rendering PDF:

![PDF](https://cloud.githubusercontent.com/assets/1000669/7492902/f8493160-f42e-11e4-8cea-1cb4f02757e7.png)

[*Read more about rendering PDF.*](https://github.com/blog/1974-pdf-viewing)

</p>
</details>

***

## How do you quickly create licensing for a repo?

<details><summary>show</summary>
<p>

When creating a repository, GitHub gives you the option of adding in a pre-made license:

![License](http://i.imgur.com/Chqj4Fg.png)

You can also add them to existing repositories by creating a new file through the web interface. When the name `LICENSE` is typed in you will get an option to use a template:

![License](http://i.imgur.com/fTjQict.png)

Also works for `.gitignore`.

[*Read more about open source licensing.*](https://help.github.com/articles/open-source-licensing/)


</p>
</details>

***

## How do you create tasks lists in markdown?

<details><summary>show</summary>
<p>

In full Markdown documents **read-only** checklists can now be added using the following syntax:

```
- [ ] Mercury
- [x] Venus
- [x] Earth
  - [x] Moon
- [x] Mars
  - [ ] Deimos
  - [ ] Phobos
```

- [ ] Mercury
- [x] Venus
- [x] Earth
  - [x] Moon
- [x] Mars
  - [ ] Deimos
  - [ ] Phobos

[*Read more about task lists in markdown documents.*](https://github.com/blog/1825-task-lists-in-all-markdown-documents)

</p>
</details>

***

## How do you use syntx highlighting in markdown?

<details><summary>show</summary>
<p>

For example, to syntax highlight Ruby code in your Markdown files write:

    ```ruby
    require 'tabbit'
    table = Tabbit.new('Name', 'Email')
    table.add_row('Tim Green', 'tiimgreen@gmail.com')
    puts table.to_s
    ```

This will produce:

```ruby
require 'tabbit'
table = Tabbit.new('Name', 'Email')
table.add_row('Tim Green', 'tiimgreen@gmail.com')
puts table.to_s
```

GitHub uses [Linguist](https://github.com/github/linguist) to perform language detection and syntax highlighting. You can find out which keywords are valid by perusing the [languages YAML file](https://github.com/github/linguist/blob/master/lib/linguist/languages.yml).

[*Read more about GitHub Flavored Markdown.*](https://help.github.com/articles/github-flavored-markdown/)


</p>
</details>