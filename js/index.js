import * as monaco from "./node_modules/monaco-editor/esm/vs/editor/edcore.main/index.js";
monaco.editor.create(document.getElementById("monaco-editor"), {
  value: ["function x() {", '\tconsole.log("Hello world!");', "}"].join("\n"),
  language: "javascript",
});
