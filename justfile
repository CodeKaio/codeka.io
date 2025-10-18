#!/usr/bin/env just --justfile

default: serve

build_options:="--printI18nWarnings --printPathWarnings --cleanDestinationDir --logLevel debug"
common_options:= "--openBrowser --printI18nWarnings --printPathWarnings --cleanDestinationDir --logLevel debug"

build:
    hugo build {{build_options}}

serve:
    hugo server {{common_options}}

draft:
    hugo server {{common_options}} --buildDrafts --buildFuture

non-breakable-spaces:
    sed -i '/^---$/,/^---$/b;s/ \([;?:!]\)/\ \1/g' content/posts/**/*.md

new date title:
    hugo new content posts/{{date}}-{{title}} -k veille

clever-preview:
    clever deploy --alias preview --force
    clever open --alias preview

clever-deploy:
    clever deploy --alias codeka.io