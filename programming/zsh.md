# Zsh

- Interactive login shell with line editing (zle), code completion (compsys),
  history mechanism, spelling correction, prompt theme or a shell script processor with
  dataflow programming and first-class files, commands, processes and pipes
- Lookup order: 1 `function`, 2 `builtin`, 3 external `command`
- Expansion order:
    - 1 Alias expansion (prefer `function` over `alias`)
    - 2 Process substitution `<(list)`, `>(list)` creates anonymous named pipe
    - 3 Parameter expansion `${param}`, `${array[@]}`
    - 4 Command substitution `$(list)` returns command output rather than status code
    - 5 Arithmetic expansion `((expr))`
    - 7 Brace expansion `a{1,2,3}`, `{1..10}`, `{a,b}{1,2}` string generation
    - 8 Filename expansion `**/*.*` (globbing)
- Builtins
    - `set` options
    - `alias al=cmd` prefer `function` definition over `alias` text substitution
    - `IFS="_ " read a b c d _ <<< "A_B_C D E_F"; echo "^$a, $b, $c, $d$"` read stdin
      into shell variables
    - `echo expr ...`
    - `printf '%c' {a..z} $'\n'`
    - `exec fd< file` opens file for reading from the fd
    - `exec fd> file` opnes file for writing from the fd
    - `exec fd<&-`, `exec fd>&-` closes input / output fd
    - `true`, `false`, `:`
    - `trap func SIG ...`
    - `zmodload mod` loads zsh module
    - `typeset var ...`, `integer var ...`, `float var ...`
- Command -> pipeline -> sublist -> list
    - Simple command
        - `command` `-s` short option `--long` long option `-` arguments
        - `exec cmd` replaces the current shell with the `cmd` by running the `cmd` in
          the current process instread of forking a sub-process. `trap ... EXIT` is not
          invoked
    - Pipeline (implicit redirection)
        - `command` stdout 1 `|` 0 stdin `command2`
        - `command` stdout 1 + stderr 2 `|&` 0 stdin `command2` equivalent to `2>&1 |`
    - Sublist
        - `pipeline && pipeline2` on success runs `pipeline2`
        - `pipeline || pipeline2` on failure runs `pipeline2`
    - List
        - Set of sublists terminated by `\n`, `;` sequencing, `&`, `&|`, `&!` last
          pipeline in the background
- Complex command
    - `if list then list [elif list then list] ... [else list] fi` executes `then` on zero
      exit status
        - `if list { list } [elif list { list }] ... [else { list }]`
    - `for name ... in word ...; do list done` iterates over words
        - `for name ... (word ...) list`
    - `for (( init; stop; next )) do list done` stops iteration when `stop` is zero
        - `for (( init; stop; next )) list`
    - `while / until list do list done` iterates while the `list` is zero / non-zero
        - `while / until list { list }`
    - `case word in pat [| pat] ... list ;; ... esac`
        - `case word { pat [| pat] ... list ;; ...}`
    - `{ list }` executes in the current shell as a job
    - `( list )` executes in a subshell with `trap reset`
    - `{ try-list } always { always-list }` returns the `try-list` exit status after
      execuing the `always-list`
    - `function word { list } [2>&1]` define function with stream redirection
    - `[[ expr ]]` returns zero exit status if the `expr` is true
- Redirection (order matters as file discriptors are pointers, redirection can be
  anywhere in a simple command or can precede or follow a complex command)
    - `< file` open file for reading from stdin
    - `> file` open a file for writing + truncate from stdout
    - `>> file` open a file for writing + append from stdout
    - `&> file`, `&>> file` redirects both stdout and stderr to file + truncate / append
      (`&>` = `> file 2>&1`)
    - `<&fd` duplicates stdin from fd
    - `>&fd` duplicates stdout to fd
    - `fd<&-`, `fd>&-` closes input / output fd
    - `{param}>&1`, `{param}>&-` open / close fd using parameter
        - E. g. `exec {stdout}>&1; echo ok >&$stdout; exec {stdout}>&-`
    - `<&$param`, `>&$param` read / write fd using parameter
        - E. g. `exec {stdin}<&0; read v <&$stdin; echo $v; exec {stdin}<&-`
    - File / stdin literal with parameter substitution and command substitution
        - `<<< "string"` here-string as stdin
        - `<< EOD\n ... \nEOD` here-document as stdin
- Function (is executed like a command with arguments passed as positional parameters in
  the same process)
    - `return` early from a function
    - `autoload func ...` lazy load of functions
    - `zcompile` precompile functions into `*.zwc`
    - `function { list } arg ...` anonymous function is executed immediately at the
      point of definition and provides scope for `local` variables
- Job (is associated with each pipeline)
    - `bg` / `C-z`, `fg %jobid`, `jobs` puts job in background, foreground, lists jobs
- Arithmetic expansion `((v = ..., t = ...))` with assignment, `$((...))` with
  command substitution. Always excplicitly declare type of variables with `integer` and
  `float` builtins
- Conditional expressions
    - File `[[ -efdpLS file, -rwx file ]]`
    - String `[[ -n str, str ==, !=, =~, <, > pat ]]`
        - `setopt RE_MATCH_PCRE`: `$MATCH` whole match, `$match[1]` capture group
    - Expressions `[[ ! expr, expr1 &&, || expr2 ]]`
    - Use arithmetic evaluation `(( ... ))` for numerical comparison
- Quoting
    - `'...'` = literal
    - `"..."` = `${param}`, `$(cmd)`, `$((expr))`, `\escape`
    - `$'...'` = `C \escape`
- Filename generation
    - `setopt EXTENDED_GLOB`
    - `*` any string, `?` any char
    - Char class isngle metch `[...]`, `[^...]` complement
        - `[:alpha:]`, `[:digit:]`, `[:alnum:]`, `[:blank:]`, `[:punct:]`
    - `^pat` pattern complement
    - `pat1~pat2` maches `pat1` excluding `pat2`
    - `pat#` zero or more, `pat##` one or more
    - `(...)` capturing group
    - `(...|...)` alternative
- Parameter expansion `${param}`, `${array[@]}`
    - Unquoted parameters `$param`, `$array[@]` words are not split and empty words are
        elided
    - Quoted parameters `"$param"`, `"${array[@]}"` empty words are not ellided
    - Return default if not set `${v:-return default}`
    - Set to default if not set `${v:=set default}`
    - Raise error if not set `${v:?error message}`
    - Subscripting `${v:offset:length}`
    - Replace `${v/pattern/repl}`, `${v//pattern/repl}`
    - Length `${#v}`

- Atoms, arrays, associative arrays


## curl options

-s silent -S show error -L follow redirects -k insecure
-X request -H header -u user:password -d 'data', @file
-D - dump response header

## Regular expressions

- (?:...) non-capturing group
- (?=...) positive lookahead
- (?!...) negative lookahead
- (?<=...) positive lookbehind
- (?<!...) negative lookbehind
