#!/usr/bin/env just --justfile

default: serve

common_options:= "--openBrowser --printI18nWarnings"

serve:
    hugo server {{common_options}}

draft:
    hugo server {{common_options}} --buildDrafts

non-breakable-spaces:
    sed -i '/^---$/,/^---$/b;s/ \([;?:!]\)/\ \1/g' content/posts/**/*.md