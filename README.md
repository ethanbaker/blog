<!--
  Created by: Ethan Baker (contact@ethanbaker.dev)

  Adapted from:
    https://github.com/othneildrew/Best-README-Template/

Here are different preset "variables" that you can search and replace in this template.
`documentation_link`
`path_to_logo`
`path_to_demo`
-->

<div id="top"></div>


<!-- PROJECT SHIELDS/BUTTONS -->
![1.0.0](https://img.shields.io/badge/status-1.0.0-red)
[![Go Coverage](https://github.com/ethanbaker/note/wiki/coverage.svg)](https://raw.githack.com/wiki/ethanbaker/note/coverage.html)
[![License][license-shield]][license-url]
[![Contributors][contributors-shield]][contributors-url]
[![Issues][issues-shield]][issues-url]

[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]


<!-- PROJECT LOGO -->
<br><br><br>
<div align="center">
  <a href="https://github.com/ethanbaker/note">
    <img src="./docs/note-logo.svg" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Note</h3>

  <p align="center">
    Note management CLI tool
  </p>
</div>


<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
        <li>
          <a href="#installation">Autocompletion</a>
          <ul>
            <li><a href="#bash">Bash</a></li>
            <li><a href="#zsh">ZSH</a></li>
            <li><a href="#fish">Fish</a></li>
            <li><a href="#powershell">PowerShell</a></li>
          </ul>
        </li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>


<!-- ABOUT -->
## About

Note is a command-line tool to manage notes. You can create, edit, and
publish markdown notes from any directory with a simple command.

<p align="right">(<a href="#top">back to top</a>)</p>


### Built With

* [Golang](https://golang.org/)
* [Cobra](https://github.com/spf13/cobra)
* [Testify](https://github.com/stretchr/testify)

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

Getting started with Note is simple. Follow these steps to set up the project
before installing.

### Prerequisites

Ensure you have Go installed on your machine. You can download it from
[golang.org](https://golang.org/dl/).

### Installation

In order to install the command-line tool, run the following command:

```sh
go install github.com/ethanbaker/note/cmd/note
```

This will install the `note` command on your machine. You can run the command
by typing `note` in your command line interface.

### Autocompletion

Note does not have default autocompletion features straight away. Instead, Note
can generate an autocompletion script for your preferred command line interface
to enable tab completion. Note delegates to
[Cobra](https://github.com/spf13/cobra) to accomplish this. As of writing this
documentation, Cobra can generate autocompletion scripts for Bash, Fish,
Powershell, and ZSH.

To setup autocompletion, your system must allow for an autocompletion
associated with the given shell script you are generating. After completing
these steps for your shell, Note's autocompletion will be enabled, allowing for
easier and faster command entry.

#### Bash

First, ensure `bash-completion` is installed on your system. For most Linux distributions, install it via your package manager:
- **Debian/Ubuntu**: `sudo apt install bash-completion`
- **Fedora**: `sudo dnf install bash-completion`
- **macOS**: `brew install bash-completion`

Next, you can generate the completion script to bash completion directory. For Debian/Ubuntu and Fedora, you can run:
```bash
note completion bash | sudo tee /etc/bash_completion.d/note > /dev/null
```

For macOS, you can run: 
```bash
[[ -r "$(brew --prefix)/etc/profile.d/bash_completion.sh" ]] && . "$(brew --prefix)/etc/profile.d/bash_completion.sh"
```

If you cannot use `sudo` to add the completion script to `bash_completion.d`, you can source the script from your local shell configuration file (e.g., `~/.bashrc`).
* You can generate the file and store it somewhere locally
    ```bash
    note completion bash > ~/.note_completion
    ```

* Then you can add the following line to your `~/.bashrc` to setup the autocompletion on each shell instance
    ```bash
    source ~/.note-completion
    ```

For these changes to take effect, you must reload your shell.
```bash
source ~/.bashrc
```

#### ZSH

First, you must enable the `compinit` module if not already enabled. Add the following to your `~/.zshrc` file:
```zsh
autoload -Uz compinit
compinit
```

Next, you can generate the completion script:
```zsh
note completion zsh > ~/.note-completion
```

Then, you need to source the script in your `~/.zshrc`:
```zsh
source ~/.note-completion
```

For these changes to take effect, you must reload your shell.
```zsh
source ~/.zshrc
```

#### Fish

Generate the Fish completion script:
```bash
note completion fish > ~/.config/fish/completions/note.fish
```

Because Fish automatically loads completion scripts from the
`~/.config/fish/completions/` directory, all you need to do is restart your shell or open a new Fish session to use the completions.

#### PowerShell

Firstly, generate the PowerShell completion script and save it in a directory:
```powershell
note completion powershell > $HOME\note-completion.ps1
```

Next, add the script to your PowerShell profile (e.g., `$PROFILE`):
```powershell
. $HOME\note-completion.ps1
```

Finally, reload your PowerShell profile for these changes to take effect:
```powershell
. $PROFILE
```

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- USAGE EXAMPLES -->
## Usage

The Note CLI tool offers commands for simple operations on notes:
* `note new`: create and open a new note with the specified title
* `note edit`: edit an existing note
* `note info`: return metadata information about the specified note
* `note list`: list all existing notes
* `note remove`: delete an existing note

You can publish finished notes, or saving those notes to a file with a specified format, by running the command `note publish`. In addition, you can edit default configurations for the Note tool using the command `note config`.

Notes are opened through a shell command of your choosing. This can be configured using the `note config` command. The default editor is set to `vi`, meaning that whenever you create or edit a note, it will open that note using the `vi` editor. Commands that don't involve opening an editor handle other CRUD operations and show associated messages.

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- ROADMAP -->
## Roadmap

- [x] Command Line Functionality
- [ ] Manager Terminal Application 
- [ ] Extended Publishing Formats
    - [ ] HTML

See the [open issues][issues-url] for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- CONTRIBUTING -->
## Contributing

For issues and suggestions, please include as much useful information as possible.
Review the [documentation][documentation-url] and make sure the issue is actually
present or the suggestion is not included. Please share issues/suggestions on the
[issue tracker][issues-url].

For patches and feature additions, please submit them as [pull requests][pulls-url]. 
Please adhere to the [conventional commits][conventional-commits-url]. standard for
commit messaging. In addition, please try to name your git branch according to your
new patch. [These standards][conventional-branches-url] are a great guide you can follow.

You can follow these steps below to create a pull request:

1. Fork the Project
2. Create your Feature Branch (`git checkout -b branch_name`)
3. Commit your Changes (`git commit -m "commit_message"`)
4. Push to the Branch (`git push origin branch_name`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- LICENSE -->
## License

This project uses the Apache 2.0 license. You can find more information in the 
[LICENSE][license-url] file.

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Ethan Baker - contact@ethanbaker.dev - [LinkedIn][linkedin-url]

Project Link: [https://github.com/ethanbaker/note][project-url]

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/ethanbaker/note.svg
[forks-shield]: https://img.shields.io/github/forks/ethanbaker/note.svg
[stars-shield]: https://img.shields.io/github/stars/ethanbaker/note.svg
[issues-shield]: https://img.shields.io/github/issues/ethanbaker/note.svg
[license-shield]: https://img.shields.io/github/license/ethanbaker/note.svg
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?logo=linkedin&colorB=555

[contributors-url]: <https://github.com/ethanbaker/note/graphs/contributors>
[forks-url]: <https://github.com/ethanbaker/note/network/members>
[stars-url]: <https://github.com/ethanbaker/note/stargazers>
[issues-url]: <https://github.com/ethanbaker/note/issues>
[pulls-url]: <https://github.com/ethanbaker/note/pulls>
[license-url]: <https://github.com/ethanbaker/note/blob/master/LICENSE>
[linkedin-url]: <https://linkedin.com/in/ethandbaker>
[project-url]: <https://github.com/ethanbaker/note>

[product-screenshot]: path_to_demo
[documentation-url]: <documentation_link>

[conventional-commits-url]: <https://www.conventionalcommits.org/en/v1.0.0/#summary>
[conventional-branches-url]: <https://docs.microsoft.com/en-us/azure/devops/repos/git/git-branching-guidance?view=azure-devops>