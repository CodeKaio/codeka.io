#!/usr/bin/env just --justfile

default: serve

common_options:= "--openBrowser --printI18nWarnings --cleanDestinationDir"

serve:
    hugo server {{common_options}}

draft:
    hugo server {{common_options}} --buildDrafts --buildFuture

non-breakable-spaces:
    sed -i '/^---$/,/^---$/b;s/ \([;?:!]\)/\ \1/g' content/posts/**/*.md