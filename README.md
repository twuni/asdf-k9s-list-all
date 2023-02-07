# GitHub Releases List

This tool lists the names of a GitHub repo's releases in ascending order.
It assumes release names follow the convention `vX.Y.Z`.

The use case is to provide a standalone, cross-platform way of implementing
the `list-all` command for [asdf][asdf] plugins whose managed tools are using
GitHub releases to communicate version updates.

## Installing

Download a [pre-compiled binary][binaries] for your platform and place it
somewhere on your `PATH` (e.g: **~/bin**).

## Usage

```bash
github-releases-list --org=twuni --repo=ajv
```

[asdf]: https://github.com/asdf-vm/asdf
[binaries]: https://releases.twuni.dev/github-releases-list
