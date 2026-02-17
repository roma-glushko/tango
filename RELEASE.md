# Release

In order to release a new version of `tango`, follow these steps:

- create a new tag with the version number, e.g. `0.0.9-alpha.1` and push it e.g.

```bash
git tag 1.2.1-rc1
git push --tag
```

This should start a new Github Action pipeline that will build the binaries and upload it to various channels.