# arc - a helpful CircleCI & VCS CLI tool [![CI Status](https://circleci.com/gh/hubci/arc.svg?style=shield)](https://app.circleci.com/pipelines/github/hubci/arc) [![Software License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hubci/arc/master/LICENSE)

arc provides small, helpful features for developers using CircleCI, GitHub, and GitLab.
arc is currently considered in alpha so it can change rapidly, including broken changes, before the v1.0 release comes out.
GitLab support has begun to land, starting with the status page.
Overtime, we'll be adding more support until it comes into feature parity with the GitHub features.


## Table of Contents

- [Installation](#installation)
  - [Linux](#linux)
  - [macOS](#macos)
  - [Windows](#windows)
- [Usage](#usage)


## Installation

### Linux

There are a few ways you can install arc on a Linux amd64 or arm64 system.

#### Ubuntu Apt Repository (recommended)
I (Ricardo N Feliciano) run an Apt/Debian repository for a lot of my software, which includes Arc.
The benefit of the Apt repository is that updates are handled by Ubuntu's built-in package manager.

For security reasons, first we install the GPG key for the repository:

```bash
sudo wget "http://pkg.feliciano.tech/ftech-archive-keyring.gpg" -P /usr/share/keyrings/
```

Now we add the repository to the system:

```bash
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/ftech-archive-keyring.gpg] http://pkg.feliciano.tech/ubuntu $(lsb_release -sc) main" | sudo tee /etc/apt/sources.list.d/felicianotech.list
```

Finally, we can install Arc:

```bash
sudo apt update && sudo apt install hubci-arc
```

#### Debian Package (.deb)
You can install arc on an Apt based operating system by downloading the `.deb` file to the desired system.

For graphical systems, you can download it from the [GitHub Releases page][gh-releases].
Many distros allow you to double-click the file to install.
Via terminal, you can do the following:

```bash
wget https://github.com/hubci/arc/releases/download/v0.10.0/arc_0.10.0_amd64.deb
sudo dpkg -i arc_0.10.0_amd64.deb
```

`0.10.0` and `amd64` may need to be replaced with your desired version and CPU architecture respectively.

#### Binary Install
You can download and run the raw arc binary from the [GitHub Releases page][gh-releases] if you don't want to use any package manager.
Simply download the tarball for your OS and architecture and extract the binary to somewhere in your `PATH`.
Here's one way to do this with `curl` and `tar`:

```bash
dlURL="https://github.com/hubci/arc/releases/download/v0.10.0/arc-v0.10.0-linux-amd64.tar.gz"
curl -sSL $dlURL | sudo tar -xz -C /usr/local/bin arc
```

`0.10.0` and `amd64` may need to be replaced with your desired version and CPU architecture respectively.

### macOS

There are two ways you can install arc on a macOS (amd64) system.

#### Brew (recommended)

Installing arc via brew is a simple one-liner:

```bash
brew install hubci/tap/arc
```

#### Binary Install
You can download and run the raw arc binary from the [GitHub Releases page][gh-releases] if you don't want to use Brew.
Simply download the tarball for your OS and architecture and extract the binary to somewhere in your `PATH`.
Here's one way to do this with `curl` and `tar`:

```bash
dlURL="https://github.com/hubci/arc/releases/download/v0.10.0/arc-v0.10.0-macos-amd64.tar.gz"
curl -sSL $dlURL | sudo tar -xz -C /usr/local/bin arc
```

`0.10.0` may need to be replaced with your desired version.

### Windows

arc supports Windows 10 by downloading and installing the binary.
Chocolately support is likely coming in the future.
If there's a Windows package manager you'd like support for (including Chocolately), please open a GitHub Issue and ask for it.

#### Binary Install (exe)
You can download and run the arc executable from the [GitHub Releases page][gh-releases].
Simply download the zip for architecture and extract the exe.


## Usage

Run `arc help` to see all commands available.


## License

This repository is licensed under the MIT license.
The license can be found [here](./LICENSE).



[gh-releases]: https://github.com/hubci/arc/releases
