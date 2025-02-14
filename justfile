serve:
    hugo server --openBrowser

draft:
    hugo server --buildDrafts --openBrowser

non-breakable-spaces:
    sed -i '/^---$/,/^---$/b;s/ \([;?:!]\)/\ \1/g' content/posts/**/*.md