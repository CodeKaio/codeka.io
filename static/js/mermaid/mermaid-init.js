const lsTheme = localStorage.getItem("theme");
const hugoTheme = document.body.classList.contains("dark-theme") ? "dark" : null;
const currTheme = lsTheme ? lsTheme : hugoTheme;
const mermaidTheme = currTheme === "dark" ? "dark" : "default";

if (typeof mermaid !== "undefined") {
    mermaid.initialize({
        startOnLoad: true,
        theme: mermaidTheme
    });
}
