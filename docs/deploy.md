# GitHub Pages Deployment

The frontend is published from `main` branch `docs/`.

Build locally:

```sh
make build
git add docs
git commit -m "build: publish pages"
git push origin main
```

Rollback by reverting the publishing commit and pushing `main`.

Custom domains require a committed `docs/CNAME` file and DNS configured with
GitHub Pages records.
