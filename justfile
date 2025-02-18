#!/usr/bin/env just --justfile

serve:
    hugo server --openBrowser --printI18nWarnings

draft:
    hugo server --buildDrafts --openBrowser

non-breakable-spaces:
    sed -i '/^---$/,/^---$/b;s/ \([;?:!]\)/\ \1/g' content/posts/**/*.md