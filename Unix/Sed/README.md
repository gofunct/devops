### Command Syntax

<details><summary>show</summary>
<p>

    /abc/p      # Print all with "abc"
    /abc/!p     # Print all without "abc"
    /abc/d      # Delete all with "abc"
    /abc/!d     # Delete all without "abc"

    /start/,/end/!d     # Select a block

    s/abc/def/
    s/abc\(...\)ghi/\1/ # Back references 
                # (see below for correct Shell quoting)

    /abc/{s/def/ghi)}   # Conditional replace

</p>
</details>

### Advanced use of sed

#### In-place Editing

<details><summary>show</summary>
<p>

To edit file use the -i option this safely changes the file contents
without any output redirection needed.

    sed -i 's/abc/ABC/' myfile.txt
    sed -i '/deleteme/d' *

</p>
</details>

#### Drop grep

<details><summary>show</summary>
<p>

Often grep and sed are used together. In all those cases grep can be
dropped. For example

    grep "pattern" file | sed "s/abc/def/"

can be written as

    sed -n "/pattern/p; s/abc/def/"

</p>
</details>

#### Grouping with sed

<details><summary>show</summary>
<p>

Always use single quotes!

    sed 's/^.*\(pattern\).*/\1/'

</p>
</details>

#### Single Quoting Single Quotes

<details><summary>show</summary>
<p>


If you want to do extraction and need a pattern based on single quotes
use \\x27 instead of trying to insert a single quote. For example:

    sed 's/.*var=\x27\([^\x27]*\)\x27.*/\1/'

to extract "some string" from "var='some string'". Or if you don't know
about the quoting, but know there are quotes

    sed 's/.*var=.\([^"\x27]*\)..*/\1/'

</p>
</details>

#### Conditional Replace with sed

<details><summary>show</summary>
<p>

    sed '/conditional pattern/{s/pattern/replacement/g}'

#### Prefix files with a boilerplate using sed

    sed -i '1s/^/# DO NOT TOUCH THIS FILE!\n\n/' *

</p>
</details>

#### Removing Newlines with sed

<details><summary>show</summary>
<p>

The only way to remove new line is this:

    sed ':a;N;$!ba;s/\n//g' file

Check out [this explanation](/Removing-newlines-with-sed) if you want to
know why.

</p>
</details>

#### Selecting Blocks

<details><summary>show</summary>
<p>

    sed '/first line/,/last line/!d' file

</p>
</details>
